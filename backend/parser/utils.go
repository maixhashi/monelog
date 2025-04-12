package parser

import (
	"time"
)

// 月利計算
func CalculateMonthlyRate(annualRate float64) float64 {
	return annualRate / 12
}

// 年率の取得
func GetAnnualRate(cardType string, installmentCount int) float64 {
	if cardType == "楽天カード" {
		switch installmentCount {
		case 2:
			return 0
		case 3:
			return 0.1225
		case 5:
			return 0.135
		case 6:
			return 0.1375
		case 10:
			return 0.145
		case 12:
			return 0.1475
		default:
			return 0.15 // 15回以上は15%
		}
	} else if cardType == "MUFG DCカード" {
		switch installmentCount {
		case 3:
			return 0.123
		case 5:
			return 0.135
		case 6:
			return 0.138
		case 10:
			return 0.1452
		case 12:
			return 0.1476
		default:
			return 0.15
		}
	} else if cardType == "eposカード" {
		return 0.15 // eposカードは全て15%
	}
	return 0.15 // デフォルト
}

// カード種類ごとの締め日と支払日を計算する関数
func CalculatePaymentDate(useDate time.Time, cardType string) time.Time {
	var paymentDay int
	var paymentDate time.Time
	
	switch cardType {
	case "MUFG DCカード":
		paymentDay = 10
		// 利用日が当月の10日以前なら当月の10日、それ以降なら翌月の10日
		if useDate.Day() <= paymentDay {
			// 当月の10日
			paymentDate = time.Date(useDate.Year(), useDate.Month(), paymentDay, 0, 0, 0, 0, time.Local)
		} else {
			// 翌月の10日
			nextMonth := AddMonths(useDate, 1)
			paymentDate = time.Date(nextMonth.Year(), nextMonth.Month(), paymentDay, 0, 0, 0, 0, time.Local)
		}
		
	case "楽天カード", "eposカード":
		paymentDay = 27
		// 利用日が当月の27日以前なら当月の27日、それ以降なら翌月の27日
		if useDate.Day() <= paymentDay {
			// 当月の27日
			paymentDate = time.Date(useDate.Year(), useDate.Month(), paymentDay, 0, 0, 0, 0, time.Local)
		} else {
			// 翌月の27日
			nextMonth := AddMonths(useDate, 1)
			paymentDate = time.Date(nextMonth.Year(), nextMonth.Month(), paymentDay, 0, 0, 0, 0, time.Local)
		}
		
	default:
		// デフォルトは翌月の1日
		nextMonth := AddMonths(useDate, 1)
		paymentDate = time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, time.Local)
	}
	
	return paymentDate
}

// 日付を指定された形式でフォーマットする
func FormatDate(date time.Time, layout string) string {
	return date.Format(layout)
}

// 指定された月数を加算する
func AddMonths(date time.Time, months int) time.Time {
	year := date.Year()
	month := int(date.Month()) + months
	
	// 月の調整
	year += (month - 1) / 12
	month = ((month - 1) % 12) + 1
	
	// 日付の調整（月末の場合）
	day := date.Day()
	lastDay := time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.Local).Day()
	if day > lastDay {
		day = lastDay
	}
	
	return time.Date(year, time.Month(month), day, date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())
}

// max関数の実装
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
