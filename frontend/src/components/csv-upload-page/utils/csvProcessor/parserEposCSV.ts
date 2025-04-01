import { CardStatementSummary } from '../../../../types/models/cardStatement';
import { parseDate, safeFormat } from '../dateUtils';
import { ja } from 'date-fns/locale';
import { addMonths, parse, format } from 'date-fns';

interface EposCSVRow {
  type: string;           // 種別（ショッピング、キャッシング、その他）
  useDate: string;        // ご利用年月日
  place: string;          // ご利用場所
  description: string;    // ご利用内容
  amount: number;         // ご利用金額
  paymentType: string;    // 支払区分
  paymentStartMonth: string; // お支払開始月
  note: string;           // 備考
  installmentCount?: number; // 分割回数
  firstPayment?: number;  // 初回支払額
  subsequentPayment?: number; // 2回目以降の支払額
}

// 分割払いの情報を解析する関数
const parseInstallmentInfo = (paymentType: string, note: string): { count: number, firstPayment: number | undefined, subsequentPayment: number | undefined } => {
  const result = { count: 1, firstPayment: undefined as number | undefined, subsequentPayment: undefined as number | undefined };
  
  // 分割回数を抽出
  const countMatch = paymentType.match(/(\d+)回払い/);
  if (countMatch) {
    result.count = parseInt(countMatch[1], 10);
  }
  
  // 備考から初回支払額と2回目以降の支払額を抽出
  if (note) {
    const firstPaymentMatch = note.match(/１回目　([0-9,]+)円/);
    const subsequentPaymentMatch = note.match(/２回目以降　([0-9,]+)円/);
    
    if (firstPaymentMatch) {
      result.firstPayment = parseInt(firstPaymentMatch[1].replace(/,/g, ''), 10);
    }
    
    if (subsequentPaymentMatch) {
      result.subsequentPayment = parseInt(subsequentPaymentMatch[1].replace(/,/g, ''), 10);
    }
  }
  
  return result;
};

// 年率から月利を計算
const calculateMonthlyRate = (annualRate: number): number => {
  return annualRate / 12;
};

// 実質年率を取得
const getAnnualRate = (installmentCount: number): number => {
  // EPOSカードの場合、分割回数に関わらず実質年率は15%
  return 0.15;
};

// 分割払いの総支払額を計算
const calculateTotalPayment = (amount: number, installmentCount: number, monthlyRate: number): number => {
  if (installmentCount <= 1) return amount;
  
  // 分割払いの総支払額計算式: 元金 × {月利 × (1 + 月利)^分割回数 / ((1 + 月利)^分割回数 - 1)} × 分割回数
  const factor = monthlyRate * Math.pow(1 + monthlyRate, installmentCount) / (Math.pow(1 + monthlyRate, installmentCount) - 1);
  return Math.round(amount * factor * installmentCount);
};

// 支払日を計算
const calculatePaymentDate = (useDate: Date): Date => {
  // EPOSカードの場合: 利用日が当月の27日以前なら当月の27日、それ以降なら翌月の27日
  const useDateDay = useDate.getDate();
  const useMonth = useDate.getMonth();
  const useYear = useDate.getFullYear();
  
  if (useDateDay <= 27) {
    return new Date(useYear, useMonth, 27);
  } else {
    return new Date(useYear, useMonth + 1, 27);
  }
};

// 和暦形式の日付を解析する関数
const parseJapaneseDate = (dateStr: string): Date => {
  // 「2023年6月11日」形式の日付を解析
  const match = dateStr.match(/(\d{4})年(\d{1,2})月(\d{1,2})日/);
  if (match) {
    const year = parseInt(match[1], 10);
    const month = parseInt(match[2], 10) - 1; // JavaScriptの月は0-11
    const day = parseInt(match[3], 10);
    return new Date(year, month, day);
  }
  
  // 解析できない場合はparseDate関数に委譲
  return parseDate(dateStr);
};

// CSVの区切り文字を検出する関数
const detectDelimiter = (text: string): string => {
  const possibleDelimiters = [',', '|', '\t'];
  const firstLine = text.split('\n')[0];
  
  for (const delimiter of possibleDelimiters) {
    if (firstLine.includes(delimiter)) {
      return delimiter;
    }
  }
  
  // デフォルトはカンマ
  return ',';
};

export const parseEposCSV = async (csvFile: File): Promise<CardStatementSummary[]> => {
  try {
    const text = await csvFile.text();
    
    // CSVの区切り文字を検出
    const delimiter = detectDelimiter(text);
    console.log(`検出された区切り文字: "${delimiter}"`);
    
    const lines = text.split('\n');
    const summaries: CardStatementSummary[] = [];
    let statementNo = 1;
    
    // CSVの内容をデバッグ出力
    console.log(`CSVファイル内容（先頭5行）:`);
    for (let i = 0; i < Math.min(5, lines.length); i++) {
      console.log(`行${i+1}: ${lines[i]}`);
    }
    
    // ヘッダー行をスキップ（最初の行がヘッダーと仮定）
    for (let i = 1; i < lines.length; i++) {
      try {
        const line = lines[i].trim();
        if (!line) continue;
        
        // 行の内容をデバッグ出力
        console.log(`処理中の行${i+1}: ${line}`);
        
        const columns = line.split(delimiter);
        console.log(`列数: ${columns.length}, 列内容:`, columns);
        
        if (columns.length < 7) {
          console.warn(`列数が不足しています: ${columns.length}, 行: ${i+1}`);
          continue;
        }
        
        const row: EposCSVRow = {
          type: columns[0].trim(),
          useDate: columns[1].trim(),
          place: columns[2].trim(),
          description: columns[3].trim(),
          amount: parseInt(columns[4].replace(/,/g, '').trim(), 10) || 0,
          paymentType: columns[5].trim(),
          paymentStartMonth: columns[6].trim(),
          note: columns.length > 7 ? columns.slice(7).join(' ').trim() : ''
        };
        
        console.log(`解析された行データ:`, row);
        
        // 日付形式の検証（和暦形式も許容）
        if (!row.useDate.match(/\d{4}[\/年]\d{1,2}[\/月]\d{1,2}[日]?/)) {
          console.warn(`無効な日付形式をスキップします: ${row.useDate}, 行: ${i+1}`);
          continue;
        }
        
        // 分割払い情報を解析
        const installmentInfo = parseInstallmentInfo(row.paymentType, row.note);
        const installmentCount = installmentInfo.count;
        
        // 利用日をDate型に変換
        let useDate;
        try {
          // 和暦形式の日付を解析
          useDate = parseJapaneseDate(row.useDate);
          if (isNaN(useDate.getTime())) {
            console.warn(`無効な日付をスキップします: ${row.useDate}, 行: ${i+1}`);
            continue;
          }
        } catch (e) {
          console.warn(`日付の解析に失敗しました: ${row.useDate}, 行: ${i+1}, エラー:`, e);
          continue;
        }
        
        // 実質年率と月利を計算
        const annualRate = getAnnualRate(installmentCount);
        const monthlyRate = calculateMonthlyRate(annualRate);
        
        // 総支払額を計算
        const totalChargeAmount = calculateTotalPayment(row.amount, installmentCount, monthlyRate);
        
        // 初回支払日を計算
        const firstPaymentDate = calculatePaymentDate(useDate);
        
        // 発生レコードを作成
        summaries.push({
          type: '発生',
          statementNo,
          cardType: 'EPOSカード',
          description: row.place,
          useDate: safeFormat(useDate, 'yyyy/MM/dd'),
          paymentDate: safeFormat(firstPaymentDate, 'yyyy/MM/dd'),
          paymentMonth: safeFormat(firstPaymentDate, 'yyyy年MM月', { locale: ja }),
          amount: row.amount,
          totalChargeAmount,
          chargeAmount: 0, // 発生時点では請求額は0
          remainingBalance: totalChargeAmount,
          paymentCount: 0,
          installmentCount,
          annualRate,
          monthlyRate
        });
        
        // 分割払いの場合、各回の支払いレコードを作成
        if (installmentCount > 1) {
          let remainingBalance = totalChargeAmount;
          
          for (let j = 1; j <= installmentCount; j++) {
            // 支払日は初回支払日から1ヶ月ずつ増加
            const paymentDate = addMonths(firstPaymentDate, j);
            
            // 各回の支払額を計算
            let chargeAmount;
            if (j === 1 && installmentInfo.firstPayment) {
              chargeAmount = installmentInfo.firstPayment;
            } else if (installmentInfo.subsequentPayment) {
              chargeAmount = installmentInfo.subsequentPayment;
            } else {
              // 均等払いの場合
              chargeAmount = j === installmentCount 
                ? remainingBalance // 最終回は残額全て
                : Math.round(totalChargeAmount / installmentCount);
            }
            
            remainingBalance -= chargeAmount;
            
            // 小数点以下の端数調整（最終回）
            if (j === installmentCount && remainingBalance !== 0) {
              chargeAmount += remainingBalance;
              remainingBalance = 0;
            }
            
            summaries.push({
              type: '分割',
              statementNo,
              cardType: 'EPOSカード',
              description: row.place,
              useDate: safeFormat(useDate, 'yyyy/MM/dd'),
              paymentDate: safeFormat(paymentDate, 'yyyy/MM/dd'),
              paymentMonth: safeFormat(paymentDate, 'yyyy年MM月', { locale: ja }),
              amount: row.amount,
              totalChargeAmount,
              chargeAmount,
              remainingBalance,
              paymentCount: j,
              installmentCount,
              annualRate,
              monthlyRate
            });
          }
        } else {
          // 一括払いの場合は、発生と同じ支払日で支払いレコードを作成
          summaries.push({
            type: '分割',
            statementNo,
            cardType: 'EPOSカード',
            description: row.place,
            useDate: safeFormat(useDate, 'yyyy/MM/dd'),
            paymentDate: safeFormat(firstPaymentDate, 'yyyy/MM/dd'),
            paymentMonth: safeFormat(firstPaymentDate, 'yyyy年MM月', { locale: ja }),
            amount: row.amount,
            totalChargeAmount,
            chargeAmount: totalChargeAmount,
            remainingBalance: 0,
            paymentCount: 1,
            installmentCount: 1,
            annualRate,
            monthlyRate
          });
        }
        
        statementNo++;
      } catch (error) {
        console.error(`行 ${i+1} の処理中にエラーが発生しました:`, error);
        // エラーが発生しても処理を続行
        continue;
      }
    }

    console.log(`EPOSカードCSV処理完了: ${statementNo - 1}件の明細を処理しました`);
    return summaries;
  } catch (error) {
    console.error('EPOSカードCSV処理エラー:', error);
    throw new Error('EPOSカードCSVの処理中にエラーが発生しました。');
  }
};
