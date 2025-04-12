package parser

import (
	"monelog/model"
	"fmt"
)

// ICardStatementParser カード明細CSVパーサーのインターフェース
type ICardStatementParser interface {
	Parse(csvData []byte) ([]model.CardStatementSummary, error)
}

// GetParser カード種類に応じたパーサーを返す
func GetParser(cardType string) (ICardStatementParser, error) {
	switch cardType {
	case "rakuten":
		return NewRakutenParser(), nil
	case "mufg":
		return NewMufgParser(), nil
	case "epos":
		return NewEposParser(), nil
	default:
		return nil, fmt.Errorf("unsupported card type: %s", cardType)
	}
}
