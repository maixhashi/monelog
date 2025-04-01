// 月利計算
export const calculateMonthlyRate = (annualRate: number): number => {
  return annualRate / 12
}

// 年率の取得
export const getAnnualRate = (cardType: string, installmentCount: number): number => {
  if (cardType === '楽天カード') {
    if (installmentCount === 2) return 0
    if (installmentCount === 3) return 0.1225
    if (installmentCount === 5) return 0.135
    if (installmentCount === 6) return 0.1375
    if (installmentCount === 10) return 0.145
    if (installmentCount === 12) return 0.1475
    return 0.15 // 15回以上は15%
  }
  // 他のカード会社の場合はここに追加
  return 0.15 // デフォルト
}

// 分割払いの各回の支払額を計算
export const calculateInstallmentPayment = (amount: number, totalAmount: number, installmentCount: number, annualRate: number): number => {
  const monthlyRate = calculateMonthlyRate(annualRate)
  // 最終回の調整用
  if (installmentCount === 1) {
    return Math.round(totalAmount)
  }
  // 通常の分割計算（元利均等返済方式）
  return Math.round(amount * (monthlyRate * Math.pow(1 + monthlyRate, installmentCount)) / (Math.pow(1 + monthlyRate, installmentCount) - 1))
}
