import { CardStatementSummary } from '../../../../types/models/cardStatement';
import { parseDate, safeFormat, calculatePaymentDate } from '../dateUtils';
import { getAnnualRate, calculateMonthlyRate } from '../cardUtils';
import { addMonths } from 'date-fns';
import { ja } from 'date-fns/locale';

export const parseDcCSV = async (csvFile: File): Promise<CardStatementSummary[]> => {
  try {
    const text = await csvFile.text();
    const lines = text.split('\n');
    const summaries: CardStatementSummary[] = [];
    let statementNo = 1;

    // ヘッダー行をスキップして処理
    for (let i = 1; i < lines.length; i++) {
      const line = lines[i].trim();
      if (!line) continue;

      try {
        // CSVの行全体を処理
        const rawLine = line.replace(/"/g, '');
        
        // 正規表現を使って必要な情報を抽出
        const useDateMatch = rawLine.match(/^(\d{4}\/\d{1,2}\/\d{1,2})/);
        const userMatch = rawLine.match(/^[\d\/]+\s*,\s*([^,]+)/);
        const paymentMethodMatch = rawLine.match(/^[\d\/]+\s*,[^,]+\s*,\s*([^,]+)/);
        const descriptionMatch = rawLine.match(/^[\d\/]+\s*,[^,]+\s*,[^,]+\s*,\s*([^,]+)/);
        const newAmountMatch = rawLine.match(/^[\d\/]+\s*,[^,]+\s*,[^,]+\s*,[^,]+\s*,\s*([0-9,]+)/);
        const currentChargeAmountMatch = rawLine.match(/^[\d\/]+\s*,[^,]+\s*,[^,]+\s*,[^,]+\s*,[^,]+\s*,\s*([0-9,]+)/);
        
        // 支払回数と備考欄の情報を抽出
        const paymentInfoMatch = rawLine.match(/(\d+)\/(\d+)回目/);
        const originalAmountMatch = rawLine.match(/当初ご利用金額\s*([0-9,]+)円/);
        const remainingBalanceMatch = rawLine.match(/お支払残高\s*([0-9,]+)円/);
        
        if (!useDateMatch || !descriptionMatch) {
          console.error(`行${i}のフォーマットが不正です:`, rawLine);
          continue;
        }
        
        const useDate = useDateMatch[1];
        const user = userMatch ? userMatch[1].trim() : '';
        const paymentMethod = paymentMethodMatch ? paymentMethodMatch[1].trim() : '';
        const description = descriptionMatch ? descriptionMatch[1].trim() : '';
        const newAmount = newAmountMatch ? parseInt(newAmountMatch[1].replace(/,/g, ''), 10) || 0 : 0;
        const currentChargeAmount = currentChargeAmountMatch ? parseInt(currentChargeAmountMatch[1].replace(/,/g, ''), 10) || 0 : 0;
        
        console.log('処理中の行データ:', {
          useDate,
          user,
          paymentMethod,
          description,
          newAmount,
          currentChargeAmount
        });

        // 分割払いの情報を抽出
        const isInstallment = paymentMethod.includes('分割');
        let installmentCount = 1;
        let currentInstallment = 1;
        let totalOriginalAmount = newAmount;
        let remainingBalance = 0;

        if (isInstallment) {
          // 支払回数の情報を抽出 (例: "24/24回目")
          if (paymentInfoMatch) {
            currentInstallment = parseInt(paymentInfoMatch[1], 10);
            installmentCount = parseInt(paymentInfoMatch[2], 10);
          }

          // 当初利用金額を抽出 (例: "当初ご利用金額 68,261円")
          if (originalAmountMatch) {
            totalOriginalAmount = parseInt(originalAmountMatch[1].replace(/,/g, ''), 10);
          }

          // お支払残高を抽出 (例: "お支払残高 3,115円")
          if (remainingBalanceMatch) {
            remainingBalance = parseInt(remainingBalanceMatch[1].replace(/,/g, ''), 10);
          }
        }

        const cardType = 'MUFG DCカード';
        const annualRate = getAnnualRate(cardType, installmentCount);
        const monthlyRate = calculateMonthlyRate(annualRate);
        
        // 利用日をDate型に変換
        const useDateObj = parseDate(useDate);
        
        // 支払日を計算
        const paymentDateObj = calculatePaymentDate(useDateObj, cardType);
        
        // 分割手数料込みの総請求額を計算
        // 計算式: 元金 * (月利 * (1 + 月利)^分割回数) / ((1 + 月利)^分割回数 - 1) * 分割回数
        let totalChargeAmount = totalOriginalAmount;
        if (isInstallment && installmentCount > 1 && monthlyRate > 0) {
          const numerator = monthlyRate * Math.pow(1 + monthlyRate, installmentCount);
          const denominator = Math.pow(1 + monthlyRate, installmentCount) - 1;
          totalChargeAmount = Math.round(totalOriginalAmount * (numerator / denominator) * installmentCount);
        }

        // 発生レコードを追加
        summaries.push({
          type: '発生',
          statementNo,
          cardType,
          description,
          useDate: safeFormat(useDateObj, 'yyyy/MM/dd'),
          paymentDate: safeFormat(paymentDateObj, 'yyyy/MM/dd'),
          paymentMonth: safeFormat(paymentDateObj, 'yyyy年MM月', { locale: ja }),
          amount: totalOriginalAmount,
          totalChargeAmount,
          chargeAmount: 0,
          remainingBalance: totalChargeAmount,
          paymentCount: 0,
          installmentCount,
          annualRate,
          monthlyRate
        });

        // 分割払いの場合、各回の支払いレコードを生成
        if (isInstallment) {
          // 各回の支払い情報を格納する配列
          const installmentPayments: CardStatementSummary[] = [];
          
          // 1回あたりの均等支払額を計算
          const monthlyPayment = Math.floor(totalChargeAmount / installmentCount);
          
          // 最終回の調整額を計算（端数処理のため）
          const lastPayment = totalChargeAmount - (monthlyPayment * (installmentCount - 1));
          
          // 各回の支払い情報を計算
          for (let j = 1; j <= installmentCount; j++) {
            try {
              // 支払日の計算
              let installmentPaymentDate;
              
              if (installmentCount === 1) {
                // 分割回数が1の場合（一括払い）は発生の支払日と同じ
                installmentPaymentDate = paymentDateObj;
              } else {
                // 分割回数が2以上の場合
                if (j === 1) {
                  // 1回目の支払いは発生の支払日の翌月
                  installmentPaymentDate = addMonths(paymentDateObj, 1);
                } else {
                  // 2回目以降は1回目から1ヶ月ずつ加算
                  // 1回目が発生の支払日の翌月なので、j回目は発生の支払日からj月後
                  installmentPaymentDate = addMonths(paymentDateObj, j);
                }
              }
              
              // 支払金額の計算（初期値）
              let paymentAmount = 0;
              
              if (j < currentInstallment) {
                // 過去の支払い
                paymentAmount = j === installmentCount ? lastPayment : monthlyPayment;
              } else if (j === currentInstallment) {
                // 現在の支払回数の場合は、CSVから取得した実際の支払額
                paymentAmount = currentChargeAmount;
              } else if (j < installmentCount) {
                // 将来の支払い（最終回以外）
                paymentAmount = monthlyPayment;
              } else {
                // 最終回は初期値として設定（後で上書きされる）
                paymentAmount = lastPayment;
              }

              // 残高計算（初期値）
              let calculatedRemainingBalance = 0;
              
              if (j < currentInstallment) {
                // 過去の支払い後の残高
                calculatedRemainingBalance = totalChargeAmount - (monthlyPayment * j);
                if (j === installmentCount - 1) {
                  calculatedRemainingBalance = lastPayment;
                } else if (j === installmentCount) {
                  calculatedRemainingBalance = 0;
                }
              } else if (j === currentInstallment) {
                // 現在の支払い後の残高
                calculatedRemainingBalance = remainingBalance - currentChargeAmount;
              } else if (j < installmentCount) {
                // 将来の支払い後の残高（最終回以外）
                const futurePaymentsMade = j - currentInstallment;
                calculatedRemainingBalance = remainingBalance - currentChargeAmount - (monthlyPayment * futurePaymentsMade);
              } else {
                // 最終回後の残高は0
                calculatedRemainingBalance = 0;
              }
              
              // 残高が負にならないように調整
              calculatedRemainingBalance = Math.max(0, calculatedRemainingBalance);

              installmentPayments.push({
                type: '分割',
                statementNo,
                cardType,
                description,
                useDate: safeFormat(useDateObj, 'yyyy/MM/dd'),
                paymentDate: safeFormat(installmentPaymentDate, 'yyyy/MM/dd'),
                paymentMonth: safeFormat(installmentPaymentDate, 'yyyy年MM月', { locale: ja }),
                amount: totalOriginalAmount,
                totalChargeAmount,
                chargeAmount: paymentAmount,
                remainingBalance: calculatedRemainingBalance,
                paymentCount: j,
                installmentCount,
                annualRate,
                monthlyRate
              });
            } catch (error) {
              console.error(`分割支払い${j}回目の処理エラー:`, error);
            }
          }
          
          // 最終回の支払い金額を修正（最終回の前の残高を使用）
          if (installmentPayments.length === installmentCount) {
            const secondLastPayment = installmentPayments[installmentCount - 2];
            const lastPayment = installmentPayments[installmentCount - 1];
            
            if (secondLastPayment && lastPayment) {
              // 最終回の支払い金額を、最終回の前の残高に設定
              lastPayment.chargeAmount = secondLastPayment.remainingBalance;
              // 最終回の残高は0
              lastPayment.remainingBalance = 0;
            }
          }
          
          // 計算済みの支払い情報を追加
          installmentPayments.forEach(payment => {
            summaries.push(payment);
          });
        } else {
          // 一括払いの場合は1回の支払いレコードを生成
          try {
            summaries.push({
              type: '分割',
              statementNo,
              cardType,
              description,
              useDate: safeFormat(useDateObj, 'yyyy/MM/dd'),
              paymentDate: safeFormat(paymentDateObj, 'yyyy/MM/dd'),
              paymentMonth: safeFormat(paymentDateObj, 'yyyy年MM月', { locale: ja }),
              amount: totalOriginalAmount,
              totalChargeAmount,
              chargeAmount: totalChargeAmount,
              remainingBalance: 0,
              paymentCount: 1,
              installmentCount: 1,
              annualRate: 0,
              monthlyRate: 0
            });
          } catch (error) {
            console.error('一括支払いの処理エラー:', error);
          }
        }

        statementNo++;
      } catch (error) {
        console.error(`行${i}の処理エラー:`, error);
        continue; // エラーが発生した行はスキップして次の行へ
      }
    }

    console.log('DCカードCSV処理完了');
    return summaries;
  } catch (error) {
    console.error('DCカードCSV処理エラー:', error);
    throw new Error('DCカードCSVの処理中にエラーが発生しました。');
  }
};
