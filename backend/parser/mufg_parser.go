package parser

import (
	"fmt"
	"math"
	"monelog/dto"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type MufgParser struct{}

func NewMufgParser() ICardStatementParser {
	return &MufgParser{}
}

func (mp *MufgParser) Parse(csvData []byte) ([]dto.CardStatementSummary, error) {
	text := string(csvData)
	lines := strings.Split(text, "\n")
	summaries := []dto.CardStatementSummary{}
	statementNo := 1

	// デバッグ情報を追加
	fmt.Println("処理開始: 行数=", len(lines))

	// 入力データが空かチェック
	if len(csvData) == 0 {
		return nil, fmt.Errorf("CSVデータが空です")
	}

	// ヘッダー行をスキップして処理
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		// デバッグ情報
		fmt.Printf("行 %d: %s\n", i, line)

		// エラーハンドリングを追加
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Recovered from panic in line %d: %v\n", i, r)
				}
			}()
			
			// CSVの行全体を処理
			rawLine := strings.Replace(line, "\"", "", -1)
			
			// 正規表現をより柔軟に修正
			useDateRe := regexp.MustCompile(`(\d{4}[\/\-]\d{1,2}[\/\-]\d{1,2})`)
			paymentInfoRe := regexp.MustCompile(`(\d+)\/(\d+)回目`)
			originalAmountRe := regexp.MustCompile(`当初ご利用金額\s*([0-9,]+)円`)
			remainingBalanceRe := regexp.MustCompile(`お支払残高\s*([0-9,]+)円`)
			
			// 行をカンマで分割して処理
			parts := strings.Split(rawLine, ",")
			
			// 利用日を抽出
			useDateMatch := useDateRe.FindStringSubmatch(rawLine)
			if len(useDateMatch) == 0 {
				fmt.Printf("行 %d: 利用日が見つかりません\n", i)
				return // 無名関数内なので、continueではなくreturnを使用
			}
			useDate := useDateMatch[1]
			
			// 説明を抽出（3番目のフィールドまたは4番目のフィールド）
			description := ""
			if len(parts) >= 4 {
				description = strings.TrimSpace(parts[3])
			} else if len(parts) >= 3 {
				description = strings.TrimSpace(parts[2])
			}
			
			// 分割払いの情報を抽出
			paymentInfoMatch := paymentInfoRe.FindStringSubmatch(rawLine)
			originalAmountMatch := originalAmountRe.FindStringSubmatch(rawLine)
			remainingBalanceMatch := remainingBalanceRe.FindStringSubmatch(rawLine)
			
			isInstallment := strings.Contains(rawLine, "分割") || paymentInfoRe.MatchString(rawLine)
			installmentCount := 1
			currentInstallment := 1
			totalOriginalAmount := 0
			remainingBalance := 0
			currentChargeAmount := 0

			// 新規利用額を抽出（5番目のフィールド）
			if len(parts) >= 5 {
				newAmount, err := strconv.Atoi(strings.Replace(strings.TrimSpace(parts[4]), ",", "", -1))
				if err == nil {
					totalOriginalAmount = newAmount
				}
			}
			
			// 今回請求額を抽出（6番目のフィールド）
			if len(parts) >= 6 {
				amount, err := strconv.Atoi(strings.Replace(strings.TrimSpace(parts[5]), ",", "", -1))
				if err == nil {
					currentChargeAmount = amount
				}
			}

			if isInstallment {
				// 支払回数の情報を抽出 (例: "24/24回目")
				if len(paymentInfoMatch) > 2 {
					currentInstallment, _ = strconv.Atoi(paymentInfoMatch[1])
					installmentCount, _ = strconv.Atoi(paymentInfoMatch[2])
				}

				// 当初利用金額を抽出 (例: "当初ご利用金額 68,261円")
				if len(originalAmountMatch) > 1 {
					totalOriginalAmount, _ = strconv.Atoi(strings.Replace(originalAmountMatch[1], ",", "", -1))
				}

				// お支払残高を抽出 (例: "お支払残高 3,115円")
				if len(remainingBalanceMatch) > 1 {
					remainingBalance, _ = strconv.Atoi(strings.Replace(remainingBalanceMatch[1], ",", "", -1))
				}
			}

			// 金額が0の場合、残高から推測
			if totalOriginalAmount == 0 && remainingBalance > 0 {
				totalOriginalAmount = remainingBalance
			}

			cardType := "MUFG DCカード"
			annualRate := GetAnnualRate(cardType, installmentCount)
			monthlyRate := CalculateMonthlyRate(annualRate)
			
			// 利用日をDate型に変換
			useDateObj, err := time.Parse("2006/1/2", useDate)
			if err != nil {
				// 2006/01/02 形式を試す
				useDateObj, err = time.Parse("2006/01/02", useDate)
				if err != nil {
					// 2006-01-02 形式を試す
					useDateObj, err = time.Parse("2006-01-02", useDate)
					if err != nil {
						// 2006-1-2 形式を試す
						useDateObj, err = time.Parse("2006-1-2", useDate)
						if err != nil {
							fmt.Printf("行 %d: 日付変換エラー: %s\n", i, err)
							return // 無名関数内なので、continueではなくreturnを使用
						}
					}
				}
			}
			
			// 支払日を計算
			paymentDateObj := CalculatePaymentDate(useDateObj, cardType)
			
			// 分割手数料込みの総請求額を計算
			totalChargeAmount := totalOriginalAmount
			if isInstallment && installmentCount > 1 && monthlyRate > 0 {
				numerator := monthlyRate * math.Pow(1 + monthlyRate, float64(installmentCount))
				denominator := math.Pow(1 + monthlyRate, float64(installmentCount)) - 1
				totalChargeAmount = int(math.Round(float64(totalOriginalAmount) * (numerator / denominator) * float64(installmentCount)))
			}

			// デバッグ情報
			fmt.Printf("パース結果: 利用日=%s, 説明=%s, 金額=%d, 分割=%v, 回数=%d/%d\n", 
				useDate, description, totalOriginalAmount, isInstallment, currentInstallment, installmentCount)

			// 発生レコードを追加
			summaries = append(summaries, dto.CardStatementSummary{
				Type:              "発生",
				StatementNo:       statementNo,
				CardType:          cardType,
				Description:       description,
				UseDate:           useDate,
				PaymentDate:       FormatDate(paymentDateObj, "2006/01/02"),
				PaymentMonth:      FormatDate(paymentDateObj, "2006年01月"),
				Amount:            totalOriginalAmount,
				TotalChargeAmount: totalChargeAmount,
				ChargeAmount:      0,
				RemainingBalance:  totalChargeAmount,
				PaymentCount:      0,
				InstallmentCount:  installmentCount,
				AnnualRate:        annualRate,
				MonthlyRate:       monthlyRate,
			})

			// 分割払いの場合、各回の支払いレコードを生成
			if isInstallment {
				// 各回の支払い情報を格納する配列
				installmentPayments := []dto.CardStatementSummary{}
				
				// 1回あたりの均等支払額を計算
				monthlyPayment := totalChargeAmount / installmentCount
				
				// 最終回の調整額を計算（端数処理のため）
				lastPayment := totalChargeAmount - (monthlyPayment * (installmentCount - 1))
				
				// 各回の支払い情報を計算
				for j := 1; j <= installmentCount; j++ {
					// 支払日の計算
					var installmentPaymentDate time.Time
					
					if installmentCount == 1 {
						// 分割回数が1の場合（一括払い）は発生の支払日と同じ
						installmentPaymentDate = paymentDateObj
					} else {
						// 分割回数が2以上の場合
						if j == 1 {
							// 1回目の支払いは発生の支払日の翌月
							installmentPaymentDate = AddMonths(paymentDateObj, 1)
						} else {
							// 2回目以降は1回目から1ヶ月ずつ加算
							installmentPaymentDate = AddMonths(paymentDateObj, j)
						}
					}
					
					// 支払金額の計算（初期値）
					paymentAmount := 0
					
					if j < currentInstallment {
						// 過去の支払い
						if j == installmentCount {
							paymentAmount = lastPayment
						} else {
							paymentAmount = monthlyPayment
						}
					} else if j == currentInstallment {
						// 現在の支払回数の場合は、CSVから取得した実際の支払額
						paymentAmount = currentChargeAmount
					} else if j < installmentCount {
						// 将来の支払い（最終回以外）
						paymentAmount = monthlyPayment
					} else {
						// 最終回は初期値として設定（後で上書きされる）
						paymentAmount = lastPayment
					}

					// 残高計算（初期値）
					calculatedRemainingBalance := 0
					
					if j < currentInstallment {
						// 過去の支払い後の残高
						calculatedRemainingBalance = totalChargeAmount - (monthlyPayment * j)
						if j == installmentCount - 1 {
							calculatedRemainingBalance = lastPayment
						} else if j == installmentCount {
							calculatedRemainingBalance = 0
						}
					} else if j == currentInstallment {
						// 現在の支払い後の残高
						calculatedRemainingBalance = remainingBalance - currentChargeAmount
					} else if j < installmentCount {
						// 将来の支払い後の残高（最終回以外）
						futurePaymentsMade := j - currentInstallment
						calculatedRemainingBalance = remainingBalance - currentChargeAmount - (monthlyPayment * futurePaymentsMade)
					} else {
						// 最終回後の残高は0
						calculatedRemainingBalance = 0
					}
					
					// 残高が負にならないように調整
					calculatedRemainingBalance = Max(0, calculatedRemainingBalance)

					installmentPayments = append(installmentPayments, dto.CardStatementSummary{
						Type:              "分割",
						StatementNo:       statementNo,
						CardType:          cardType,
						Description:       description,
						UseDate:           useDate,
						PaymentDate:       FormatDate(installmentPaymentDate, "2006/01/02"),
						PaymentMonth:      FormatDate(installmentPaymentDate, "2006年01月"),
						Amount:            totalOriginalAmount,
						TotalChargeAmount: totalChargeAmount,
						ChargeAmount:      paymentAmount,
						RemainingBalance:  calculatedRemainingBalance,
						PaymentCount:      j,
						InstallmentCount:  installmentCount,
						AnnualRate:        annualRate,
						MonthlyRate:       monthlyRate,
					})
				}
				
				// 最終回の支払い金額を修正（最終回の前の残高を使用）
				if len(installmentPayments) == installmentCount && installmentCount >= 2 {
					secondLastPayment := installmentPayments[installmentCount - 2]
					lastPayment := &installmentPayments[installmentCount - 1]
					
					// 最終回の支払い金額を、最終回の前の残高に設定
					lastPayment.ChargeAmount = secondLastPayment.RemainingBalance
					// 最終回の残高は0
					lastPayment.RemainingBalance = 0
				}
				
				// 計算済みの支払い情報を追加
				summaries = append(summaries, installmentPayments...)
			} else {
				// 一括払いの場合は1回の支払いレコードを生成
				summaries = append(summaries, dto.CardStatementSummary{
					Type:              "分割",
					StatementNo:       statementNo,
					CardType:          cardType,
					Description:       description,
					UseDate:           useDate,
					PaymentDate:       FormatDate(paymentDateObj, "2006/01/02"),
					PaymentMonth:      FormatDate(paymentDateObj, "2006年01月"),
					Amount:            totalOriginalAmount,
					TotalChargeAmount: totalChargeAmount,
					ChargeAmount:      totalChargeAmount,
					RemainingBalance:  0,
					PaymentCount:      1,
					InstallmentCount:  1,
					AnnualRate:        0,
					MonthlyRate:       0,
				})
			}

			statementNo++
		}()
	}

	// 有効な明細が見つからなかった場合
	if len(summaries) == 0 {
		return nil, fmt.Errorf("有効なカード明細が見つかりませんでした。データ形式を確認してください。")
	}

	return summaries, nil
}
