package parser

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"monelog/model"
	"strconv"
	"strings"
	"time"
)

type RakutenParser struct{}

func NewRakutenParser() ICardStatementParser {
	return &RakutenParser{}
}

func (rp *RakutenParser) Parse(csvData []byte) ([]model.CardStatementSummary, error) {
	// CSVデータをUTF-8に変換（必要に応じて）
	csvReader := csv.NewReader(bytes.NewReader(csvData))
	
	// 行を読み込む
	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	summaries := []model.CardStatementSummary{}
	statementNo := 1

	// ヘッダー行をスキップして処理
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if len(line) < 9 {
			continue
		}

		// CSVの各列を取得
		useDate := strings.TrimSpace(line[0])
		description := strings.TrimSpace(line[1])
		paymentMethod := strings.TrimSpace(line[3])
		
		amount, err := strconv.Atoi(strings.Replace(strings.TrimSpace(line[4]), ",", "", -1))
		if err != nil {
			amount = 0
		}
		
		totalAmount, err := strconv.Atoi(strings.Replace(strings.TrimSpace(line[6]), ",", "", -1))
		if err != nil {
			totalAmount = 0
		}
		
		currentMonthPayment, err := strconv.Atoi(strings.Replace(strings.TrimSpace(line[8]), ",", "", -1))
		if err != nil {
			currentMonthPayment = 0
		}

		// 分割払いの情報を抽出
		isInstallment := strings.Contains(paymentMethod, "分割")
		installmentCount := 1
		currentInstallment := 1

		if isInstallment {
			// 「分割変更12回払い(1回目)」や「分割12回払い(1回目)」などの形式に対応
			fmt.Sscanf(paymentMethod, "分割%d回払い(%d回目)", &installmentCount, &currentInstallment)
			if installmentCount == 1 {
				fmt.Sscanf(paymentMethod, "分割変更%d回払い(%d回目)", &installmentCount, &currentInstallment)
			}
		}

		cardType := "楽天カード"
		annualRate := GetAnnualRate(cardType, installmentCount)
		monthlyRate := CalculateMonthlyRate(annualRate)
		
		// 利用日をパース
		useDateObj, err := time.Parse("2006/01/02", useDate)
		if err != nil {
			continue
		}
		
		// 支払日を計算
		paymentDateObj := CalculatePaymentDate(useDateObj, cardType)
		
		// 発生レコードを追加
		summary := model.CardStatementSummary{
			Type:              "発生",
			StatementNo:       statementNo,
			CardType:          cardType,
			Description:       description,
			UseDate:           useDate,
			PaymentDate:       FormatDate(paymentDateObj, "2006/01/02"),
			PaymentMonth:      FormatDate(paymentDateObj, "2006年01月"),
			Amount:            amount,
			TotalChargeAmount: totalAmount,
			ChargeAmount:      0,
			RemainingBalance:  totalAmount,
			PaymentCount:      0,
			InstallmentCount:  installmentCount,
			AnnualRate:        annualRate,
			MonthlyRate:       monthlyRate,
		}
		
		summaries = append(summaries, summary)

		// 分割払いの場合、各回の支払いレコードを生成
		if isInstallment {
			remainingBalance := totalAmount
			
			for j := 1; j <= installmentCount; j++ {
				// 支払日の計算
				// 分割払いの場合、1回目の支払いは発生レコードの支払月日の1ヶ月後から開始
				// 2回目以降は1回目から1ヶ月ずつ加算
				installmentPaymentDate := AddMonths(paymentDateObj, j)
				
				// 支払金額の計算
				var paymentAmount int
				
				if j == currentInstallment {
					// 現在の支払回数の場合は、CSVから取得した実際の支払額
					paymentAmount = currentMonthPayment
				} else if j < currentInstallment {
					// 過去の支払いは、総額から現在の残高を引いて均等に分配
					pastPayments := totalAmount - remainingBalance - currentMonthPayment
					if currentInstallment > 1 {
						paymentAmount = int(float64(pastPayments) / float64(currentInstallment-1))
					}
				} else if j == installmentCount {
					// 最終回は残額を全て支払う
					paymentAmount = remainingBalance
				} else {
					// 将来の支払いは均等に分配
					paymentAmount = int(float64(remainingBalance) / float64(installmentCount-j+1))
				}

				// 残高を更新
				if j >= currentInstallment {
					remainingBalance -= paymentAmount
				}

				installmentSummary := model.CardStatementSummary{
					Type:              "分割",
					StatementNo:       statementNo,
					CardType:          cardType,
					Description:       description,
					UseDate:           useDate,
					PaymentDate:       FormatDate(installmentPaymentDate, "2006/01/02"),
					PaymentMonth:      FormatDate(installmentPaymentDate, "2006年01月"),
					Amount:            amount,
					TotalChargeAmount: totalAmount,
					ChargeAmount:      paymentAmount,
					RemainingBalance:  Max(0, remainingBalance),
					PaymentCount:      j,
					InstallmentCount:  installmentCount,
					AnnualRate:        annualRate,
					MonthlyRate:       monthlyRate,
				}
				
				summaries = append(summaries, installmentSummary)
			}
		} else {
			// 一括払いの場合は1回の支払いレコードを生成
			installmentSummary := model.CardStatementSummary{
				Type:              "分割",
				StatementNo:       statementNo,
				CardType:          cardType,
				Description:       description,
				UseDate:           useDate,
				PaymentDate:       FormatDate(paymentDateObj, "2006/01/02"),
				PaymentMonth:      FormatDate(paymentDateObj, "2006年01月"),
				Amount:            amount,
				TotalChargeAmount: totalAmount,
				ChargeAmount:      totalAmount,
				RemainingBalance:  0,
				PaymentCount:      1,
				InstallmentCount:  1,
				AnnualRate:        0,
				MonthlyRate:       0,
			}
			
			summaries = append(summaries, installmentSummary)
		}

		statementNo++
	}

	return summaries, nil
}
