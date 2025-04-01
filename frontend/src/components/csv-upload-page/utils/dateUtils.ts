import { format, addMonths, subMonths, parseISO, endOfMonth, setDate, isAfter } from 'date-fns'
import { ja } from 'date-fns/locale'

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
export const parseDate = (dateStr: string): Date => {
  try {
    // YYYY/MM/DD形式の場合
    if (/^\d{4}\/\d{1,2}\/\d{1,2}$/.test(dateStr)) {
      const [year, month, day] = dateStr.split('/').map(Number)
      const date = new Date(year, month - 1, day)
      if (!isNaN(date.getTime())) return date
    }
    
    // YYYY-MM-DD形式の場合
    if (/^\d{4}-\d{1,2}-\d{1,2}$/.test(dateStr)) {
      const [year, month, day] = dateStr.split('-').map(Number)
      const date = new Date(year, month - 1, day)
      if (!isNaN(date.getTime())) return date
    }
    
    // 標準的なJavaScriptのDateパース
    const date = new Date(dateStr)
    if (!isNaN(date.getTime())) return date
    
    // すべての方法が失敗した場合は現在の日付を返す
    console.warn('日付のパースに失敗しました:', dateStr)
    return new Date()
  } catch (error) {
    console.error('日付パースエラー:', error)
    return new Date()
  }
}

// 安全にフォーマットする関数
export const safeFormat = (date: Date, formatStr: string, options?: any): string => {
  try {
    if (isNaN(date.getTime())) {
      console.warn('無効な日付でフォーマットを試行:', date)
      return 'Invalid Date'
    }
    return format(date, formatStr, options)
  } catch (error) {
    console.error('日付フォーマットエラー:', error)
    return 'Format Error'
  }
}

// カード種類ごとの締め日と支払日を計算する関数
export const calculatePaymentDate = (useDate: Date, cardType: string): Date => {
  let cutoffDay = 0;
  let paymentDay = 0;
  
  // カード種類ごとの締め日と支払日を設定
  if (cardType === 'MUFG DCカード') {
    cutoffDay = 10;
    paymentDay = 10;
  } else if (cardType === '楽天カード') {
    cutoffDay = 27;
    paymentDay = 27;
  } else if (cardType === 'eposカード') {
    cutoffDay = 27;
    paymentDay = 27;
  } else {
    // デフォルト値（楽天カードと同じ）
    cutoffDay = 27;
    paymentDay = 27;
  }
  
  // 前月の末日を取得
  const prevMonthEnd = endOfMonth(subMonths(useDate, 1));
  // 当月の末日を取得
  const currentMonthEnd = endOfMonth(useDate);
  
  // 前月の締め日を計算
  const prevMonthCutoff = setDate(prevMonthEnd, cutoffDay);
  // 当月の締め日を計算
  const currentMonthCutoff = setDate(currentMonthEnd, cutoffDay);
  
  // 利用日が前月の締め日以前なら前月の支払日、そうでなければ当月の支払日
  if (!isAfter(useDate, prevMonthCutoff)) {
    return setDate(prevMonthEnd, paymentDay);
  } else {
    return setDate(currentMonthEnd, paymentDay);
  }
}
