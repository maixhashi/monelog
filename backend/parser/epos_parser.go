package parser

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"math"
	"monelog/dto"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type EposParser struct{}

func NewEposParser() ICardStatementParser {
	return &EposParser{}
}

func (ep *EposParser) Parse(csvData []byte) ([]dto.CardStatementSummary, error) {
	// CSVの区切り文字を検出
	text := string(csvData)
	delimiter := ep.detectDelimiter(text)
	
	// CSVリーダーを設定
	reader := csv.NewReader(bytes.NewReader(csvData))
	reader.Comma = rune(delimiter[0])
	
	// 行を読み込む
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	summaries := []dto.CardStatementSummary{}
	statementNo := 1

	// ヘッダー行をスキップして処理
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if len(line) < 7 {
			continue
		}

		// CSVの各列を取得
		useDate := strings.TrimSpace(line[1])
		place := strings.TrimSpace(line[2])
		
		amount, err := strconv.Atoi(strings.Replace(strings.TrimSpace(line[4]), ",", "", -1))
		if err != nil {
			amount = 0
		}
		
		paymentType := strings.TrimSpace(line[5])
		note := ""
		if len(line) > 7 {
			note = strings.Join(line[7:], " ")
		}

		// 日付形式の検証（和暦形式も許容）
		if !ep.isValidDate(useDate) {
			continue
		}
		
		// 分割払い情報を解析
		installmentInfo := ep.parseInstallmentInfo(paymentType, note)
		installmentCount := installmentInfo.count
		
		// 利用日をDate型に変換
		useDateObj, err := ep.parseJapaneseDate(useDate)
		if err != nil {
			continue
		}
		
		cardType := "eposカード"
		annualRate := GetAnnualRate(cardType, installmentCount)
		monthlyRate := CalculateMonthlyRate(annualRate)
		
		// 総支払額を計算
		totalChargeAmount := ep.calculateTotalPayment(amount, installmentCount, monthlyRate)
		
		// 初回支払日を計算
		firstPaymentDate := ep.calculatePaymentDate(useDateObj)
		
		// 発生レコードを作成
		summaries = append(summaries, dto.CardStatementSummary{
			Type:              "発生",
			StatementNo:       statementNo,
			CardType:          cardType,
			Description:       place,
			UseDate:           FormatDate(useDateObj, "2006/01/02"),
			PaymentDate:       FormatDate(firstPaymentDate, "2006/01/02"),
			PaymentMonth:      FormatDate(firstPaymentDate, "2006年01月"),
			Amount:            amount,
			TotalChargeAmount: totalChargeAmount,
			ChargeAmount:      0,
			RemainingBalance:  totalChargeAmount,
			PaymentCount:      0,
			InstallmentCount:  installmentCount,
			AnnualRate:        annualRate,
			MonthlyRate:       monthlyRate,
		})
		
		// 分割払いの場合、各回の支払いレコードを作成
		if installmentCount > 1 {
			remainingBalance := totalChargeAmount
			
			for j := 1; j <= installmentCount; j++ {
				// 支払日は初回支払日から1ヶ月ずつ増加
				paymentDate := AddMonths(firstPaymentDate, j)
				
				// 各回の支払額を計算
				var chargeAmount int
				if j == 1 && installmentInfo.firstPayment > 0 {
					chargeAmount = installmentInfo.firstPayment
				} else if installmentInfo.subsequentPayment > 0 {
					chargeAmount = installmentInfo.subsequentPayment
				} else {
					// 均等払いの場合
					if j == installmentCount {
						chargeAmount = remainingBalance // 最終回は残額全て
					} else {
						chargeAmount = int(math.Round(float64(totalChargeAmount) / float64(installmentCount)))
					}
				}
				
				remainingBalance -= chargeAmount
				
				// 小数点以下の端数調整（最終回）
				if j == installmentCount && remainingBalance != 0 {
					chargeAmount += remainingBalance
					remainingBalance = 0
				}
				
				summaries = append(summaries, dto.CardStatementSummary{
					Type:              "分割",
					StatementNo:       statementNo,
					CardType:          cardType,
					Description:       place,
					UseDate:           FormatDate(useDateObj, "2006/01/02"),
					PaymentDate:       FormatDate(paymentDate, "2006/01/02"),
					PaymentMonth:      FormatDate(paymentDate, "2006年01月"),
					Amount:            amount,
					TotalChargeAmount: totalChargeAmount,
					ChargeAmount:      chargeAmount,
					RemainingBalance:  remainingBalance,
					PaymentCount:      j,
					InstallmentCount:  installmentCount,
					AnnualRate:        annualRate,
					MonthlyRate:       monthlyRate,
				})
			}
		} else {
			// 一括払いの場合は、発生と同じ支払日で支払いレコードを作成
			summaries = append(summaries, dto.CardStatementSummary{
				Type:              "分割",
				StatementNo:       statementNo,
				CardType:          cardType,
				Description:       place,
				UseDate:           FormatDate(useDateObj, "2006/01/02"),
				PaymentDate:       FormatDate(firstPaymentDate, "2006/01/02"),
				PaymentMonth:      FormatDate(firstPaymentDate, "2006年01月"),
				Amount:            amount,
				TotalChargeAmount: totalChargeAmount,
				ChargeAmount:      totalChargeAmount,
				RemainingBalance:  0,
				PaymentCount:      1,
				InstallmentCount:  1,
				AnnualRate:        annualRate,
				MonthlyRate:       monthlyRate,
			})
		}
		
		statementNo++
	}

	return summaries, nil
}

// CSVの区切り文字を検出する関数
func (ep *EposParser) detectDelimiter(text string) string {
	possibleDelimiters := []string{",", "|", "\t"}
	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		return ","
	}
	
	firstLine := lines[0]
	
	for _, delimiter := range possibleDelimiters {
		if strings.Contains(firstLine, delimiter) {
			return delimiter
		}
	}
	
	// デフォルトはカンマ
	return ","
}

// 分割払いの情報を解析する関数
type installmentInfo struct {
	count             int
	firstPayment      int
	subsequentPayment int
}

func (ep *EposParser) parseInstallmentInfo(paymentType string, note string) installmentInfo {
	result := installmentInfo{count: 1, firstPayment: 0, subsequentPayment: 0}
	
	// 分割回数を抽出
	re := regexp.MustCompile(`(\d+)回払い`)
	matches := re.FindStringSubmatch(paymentType)
	if len(matches) > 1 {
		count, err := strconv.Atoi(matches[1])
		if err == nil {
			result.count = count
		}
	}
	
	// 備考から初回支払額と2回目以降の支払額を抽出
	if note != "" {
		firstPaymentRe := regexp.MustCompile(`１回目　([0-9,]+)円`)
		subsequentPaymentRe := regexp.MustCompile(`２回目以降　([0-9,]+)円`)
		
		firstPaymentMatches := firstPaymentRe.FindStringSubmatch(note)
		if len(firstPaymentMatches) > 1 {
			amount, err := strconv.Atoi(strings.Replace(firstPaymentMatches[1], ",", "", -1))
			if err == nil {
				result.firstPayment = amount
			}
		}
		
		subsequentPaymentMatches := subsequentPaymentRe.FindStringSubmatch(note)
		if len(subsequentPaymentMatches) > 1 {
			amount, err := strconv.Atoi(strings.Replace(subsequentPaymentMatches[1], ",", "", -1))
			if err == nil {
				result.subsequentPayment = amount
			}
		}
	}
	
	return result
}

// 分割払いの総支払額を計算
func (ep *EposParser) calculateTotalPayment(amount int, installmentCount int, monthlyRate float64) int {
	if installmentCount <= 1 {
		return amount
	}
	
	// 分割払いの総支払額計算式: 元金 × {月利 × (1 + 月利)^分割回数 / ((1 + 月利)^分割回数 - 1)} × 分割回数
	factor := monthlyRate * math.Pow(1 + monthlyRate, float64(installmentCount)) / (math.Pow(1 + monthlyRate, float64(installmentCount)) - 1)
	return int(math.Round(float64(amount) * factor * float64(installmentCount)))
}

// 支払日を計算
func (ep *EposParser) calculatePaymentDate(useDate time.Time) time.Time {
	// EPOSカードの場合: 利用日が当月の27日以前なら当月の27日、それ以降なら翌月の27日
	useDateDay := useDate.Day()
	useMonth := useDate.Month()
	useYear := useDate.Year()
	
	if useDateDay <= 27 {
		return time.Date(useYear, useMonth, 27, 0, 0, 0, 0, time.Local)
	} else {
		return time.Date(useYear, time.Month(int(useMonth) + 1), 27, 0, 0, 0, 0, time.Local)
	}
}

// 和暦形式の日付を解析する関数
func (ep *EposParser) parseJapaneseDate(dateStr string) (time.Time, error) {
	// 「2023年6月11日」形式の日付を解析
	re := regexp.MustCompile(`(\d{4})年(\d{1,2})月(\d{1,2})日`)
	matches := re.FindStringSubmatch(dateStr)
	if len(matches) > 3 {
		year, _ := strconv.Atoi(matches[1])
		month, _ := strconv.Atoi(matches[2])
		day, _ := strconv.Atoi(matches[3])
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil
	}
	
	// 「2023/6/11」形式の日付を解析
	t, err := time.Parse("2006/1/2", dateStr)
	if err == nil {
		return t, nil
	}
	
	// 「2023/06/11」形式の日付を解析
	t, err = time.Parse("2006/01/02", dateStr)
	if err == nil {
		return t, nil
	}
	
	return time.Time{}, fmt.Errorf("invalid date format: %s", dateStr)
}

// 日付形式が有効かどうかを確認する関数
func (ep *EposParser) isValidDate(dateStr string) bool {
	// 「2023年6月11日」形式をチェック
	if regexp.MustCompile(`\d{4}年\d{1,2}月\d{1,2}日`).MatchString(dateStr) {
		return true
	}
	
	// 「2023/6/11」または「2023/06/11」形式をチェック
	if regexp.MustCompile(`\d{4}/\d{1,2}/\d{1,2}`).MatchString(dateStr) {
		return true
	}
	
	return false
}
