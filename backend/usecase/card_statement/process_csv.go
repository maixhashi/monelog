package card_statement

import (
	"bytes"
	"io"
	"mime/multipart"
	"monelog/model"
	"monelog/parser"
)

func (csu *cardStatementUsecase) ProcessCSV(file *multipart.FileHeader, request model.CardStatementRequest) ([]model.CardStatementResponse, error) {
	if err := csu.csv.ValidateCardStatementRequest(request); err != nil {
		return nil, err
	}

	// ファイルを開く
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// ファイルの内容を読み込む
	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, src); err != nil {
		return nil, err
	}

	// カード種類に応じたパーサーを取得
	cardParser, err := parser.GetParser(request.CardType)
	if err != nil {
		return nil, err
	}

	// CSVを解析
	summaries, err := cardParser.Parse(buf.Bytes())
	if err != nil {
		return nil, err
	}

	// 既存のデータを削除
	if err := csu.csr.DeleteCardStatements(request.UserId); err != nil {
		return nil, err
	}

	// 解析結果をデータベースに保存
	cardStatements := make([]model.CardStatement, len(summaries))
	for i, summary := range summaries {
		cardStatements[i] = summary.ToModel(request.UserId, request.Year, request.Month)
	}

	if err := csu.csr.CreateCardStatements(cardStatements); err != nil {
		return nil, err
	}

	// レスポンスを作成
	responses := make([]model.CardStatementResponse, len(cardStatements))
	for i, cardStatement := range cardStatements {
		responses[i] = cardStatement.ToResponse()
	}

	return responses, nil
}