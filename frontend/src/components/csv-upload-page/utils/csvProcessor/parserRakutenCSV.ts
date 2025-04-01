import { CardStatementSummary } from '../../../../types/models/cardStatement';
import { parseDate, safeFormat, calculatePaymentDate } from '../dateUtils';
import { getAnnualRate, calculateMonthlyRate } from '../cardUtils';
import { addMonths } from 'date-fns';
import { ja } from 'date-fns/locale';

export const parseRakutenCSV = async (csvFile: File): Promise<CardStatementSummary[]> => {
  try {
    const text = await csvFile.text();
    const lines = text.split('\n');
    const summaries: CardStatementSummary[] = [];
    let statementNo = 1;

    // ヘッダー行をスキップして処理
    for (let i = 1; i < lines.length; i++) {
      const line = lines[i].trim();
      if (!line) continue;

      const columns = line.split(',');
      if (columns.length < 10) continue;

      try {
        // CSVの各列を取得
        const useDate = columns[0].replace(/"/g, '');
        const description = columns[1].replace(/"/g, '');
        const user = columns[2].replace(/"/g, '');
        const paymentMethod = columns[3].replace(/"/g, '');
        const amount = parseInt(columns[4].replace(/"/g, '').replace(/,/g, ''), 10) || 0;
        const fee = parseInt(columns[5].replace(/"/g, '').replace(/,/g, ''), 10) || 0;
        const totalAmount = parseInt(columns[6].replace(/"/g, '').replace(/,/g, ''), 10) || 0;
        const paymentMonth = columns[7].replace(/"/g, '');
        const currentMonthPayment = parseInt(columns[8].replace(/"/g, '').replace(/,/g, ''), 10) || 0;

        console.log('処理中の行データ:', {
          useDate,
          description,
          paymentMethod,
          amount,
          totalAmount,
          currentMonthPayment
        });

        // 分割払いの情報を抽出（修正）
        const isInstallment = paymentMethod.includes('分割');
        let installmentCount = 1;
        let currentInstallment = 1;

        if (isInstallment) {
          // 「分割変更12回払い(1回目)」や「分割12回払い(1回目)」などの形式に対応
          const match = paymentMethod.match(/分割(?:変更)?(\d+)回払い\((\d+)回目\)/);
          if (match) {
            installmentCount = parseInt(match[1], 10);
            currentInstallment = parseInt(match[2], 10);
          }
        }

        const cardType = '楽天カード';
        const annualRate = getAnnualRate(cardType, installmentCount);
        const monthlyRate = calculateMonthlyRate(annualRate);
        
        // 利用日をDate型に変換
        const useDateObj = parseDate(useDate);
        
        // 支払日を計算
        const paymentDateObj = calculatePaymentDate(useDateObj, cardType);
        
        // 発生レコードを追加
        summaries.push({
          type: '発生',
          statementNo,
          cardType,
          description,
          useDate: safeFormat(useDateObj, 'yyyy/MM/dd'),
          paymentDate: safeFormat(paymentDateObj, 'yyyy/MM/dd'),
          paymentMonth: safeFormat(paymentDateObj, 'yyyy年MM月', { locale: ja }),
          amount,
          totalChargeAmount: totalAmount,
          chargeAmount: 0,
          remainingBalance: totalAmount,
          paymentCount: 0,
          installmentCount,
          annualRate,
          monthlyRate
        });

        // 分割払いの場合、各回の支払いレコードを生成
        if (isInstallment) {
          let remainingBalance = totalAmount;
          
          for (let j = 1; j <= installmentCount; j++) {
            try {
              // 支払日の計算
              // 分割払いの場合、1回目の支払いは発生レコードの支払月日の1ヶ月後から開始
              // 2回目以降は1回目から1ヶ月ずつ加算
              const installmentPaymentDate = addMonths(paymentDateObj, j);
              
              // 支払金額の計算
              // 現在の支払回数までは実際の支払額を使用し、それ以降は均等に分割
              let paymentAmount = 0;
              
              if (j === currentInstallment) {
                // 現在の支払回数の場合は、CSVから取得した実際の支払額
                paymentAmount = currentMonthPayment;
              } else if (j < currentInstallment) {
                // 過去の支払いは、総額から現在の残高を引いて均等に分配
                const pastPayments = totalAmount - remainingBalance - currentMonthPayment;
                paymentAmount = Math.round(pastPayments / (currentInstallment - 1));
              } else if (j === installmentCount) {
                // 最終回は残額を全て支払う
                paymentAmount = remainingBalance;
              } else {
                // 将来の支払いは均等に分配
                paymentAmount = Math.round(remainingBalance / (installmentCount - j + 1));
              }

              // 残高を更新
              if (j >= currentInstallment) {
                remainingBalance -= paymentAmount;
              }

              summaries.push({
                type: '分割',
                statementNo,
                cardType,
                description,
                useDate: safeFormat(useDateObj, 'yyyy/MM/dd'),
                paymentDate: safeFormat(installmentPaymentDate, 'yyyy/MM/dd'),
                paymentMonth: safeFormat(installmentPaymentDate, 'yyyy年MM月', { locale: ja }),
                amount,
                totalChargeAmount: totalAmount,
                chargeAmount: paymentAmount,
                remainingBalance: Math.max(0, remainingBalance),
                paymentCount: j,
                installmentCount,
                annualRate,
                monthlyRate
              });
            } catch (error) {
              console.error(`分割支払い${j}回目の処理エラー:`, error);
            }
          }
        } else {
          // 一括払いの場合は1回の支払いレコードを生成
          try {
            const paymentDate = paymentDateObj;
            
            summaries.push({
              type: '分割',
              statementNo,
              cardType,
              description,
              useDate: safeFormat(useDateObj, 'yyyy/MM/dd'),
              paymentDate: safeFormat(paymentDate, 'yyyy/MM/dd'),
              paymentMonth: safeFormat(paymentDate, 'yyyy年MM月', { locale: ja }),
              amount,
              totalChargeAmount: totalAmount,
              chargeAmount: totalAmount,
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

    return summaries;
  } catch (error) {
    console.error('楽天カードCSV処理エラー:', error);
    throw new Error('楽天カードCSVの処理中にエラーが発生しました。');
  }
};
