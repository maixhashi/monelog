import { parse, format, isValid, addMonths, getDate, setDate, endOfMonth, Locale } from 'date-fns';
import { ja } from 'date-fns/locale';

// 日付文字列をシリアル値に変換（Excel形式）
export const dateToSerial = (dateStr: string): number => {
  try {
    const date = new Date(dateStr)
    if (isNaN(date.getTime())) {
      console.warn('無効な日付文字列:', dateStr)
      return 0
    }
    return Math.floor((date.getTime() / 86400000) + 25569)
  } catch (error) {
    console.error('日付変換エラー:', error)
    return 0
  }
}

// シリアル値を日付文字列に変換（Excel形式）
export const serialToDate = (serial: number): Date => {
  return new Date((serial - 25569) * 86400000)
}

// 日付文字列をDate型に変換する関数
export const parseDate = (dateString: string): Date => {
  // 日付文字列をDate型に変換
  const parsedDate = parse(dateString, 'yyyy/MM/dd', new Date());
  if (!isValid(parsedDate)) {
    throw new Error(`Invalid date format: ${dateString}`);
  }
  return parsedDate;
}

// 安全にフォーマットする関数
export const safeFormat = (date: Date, formatStr: string, options?: { locale: Locale }): string => {
  try {
    return format(date, formatStr, options);
  } catch (error) {
    console.error('Date formatting error:', error);
    return '';
  }
}

// カード種類ごとの締め日と支払日を計算する関数
export const calculatePaymentDate = (useDate: Date, cardType: string): Date => {
  let paymentDay: number;
  let paymentDate: Date;
  
  switch (cardType) {
    case 'MUFG DCカード':
      paymentDay = 10;
      // 利用日が当月の10日以前なら当月の10日、それ以降なら翌月の10日
      if (getDate(useDate) <= paymentDay) {
        // 当月の10日
        paymentDate = setDate(useDate, paymentDay);
      } else {
        // 翌月の10日
        paymentDate = setDate(addMonths(useDate, 1), paymentDay);
      }
      break;
      
    case '楽天カード':
    case 'eposカード':
      paymentDay = 27;
      // 利用日が当月の27日以前なら当月の27日、それ以降なら翌月の27日
      if (getDate(useDate) <= paymentDay) {
        // 当月の27日
        paymentDate = setDate(useDate, paymentDay);
      } else {
        // 翌月の27日
        paymentDate = setDate(addMonths(useDate, 1), paymentDay);
      }
      break;
      
    default:
      // デフォルトは翌月の1日
      paymentDate = setDate(addMonths(useDate, 1), 1);
  }
  
  return paymentDate;
}
