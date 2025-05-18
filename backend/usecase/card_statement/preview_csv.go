package card_statement

import (
	"bytes"
	"io"
	"mime/multipart"
	"monelog/model"
	"monelog/parser"
)

func (csu *cardStatementUsecase) PreviewCSV(file *multipart.FileHeader, request model.CardStatementPreviewRequest) ([]model.CardStatementResponse, error) {
	if err := csu.csv.ValidateCardStatementPreviewRequest(request); err != nil {
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

	// レスポンスを作成（DBには保存しない）
	responses := make([]model.CardStatementResponse, len(summaries))
	for i, summary := range summaries {
		cardStatement := summary.ToModel(request.UserId, request.Year, request.Month)
		
		// リクエストから年月を取得して設定
		if request.Year > 0 {
			cardStatement.Year = request.Year
		}
		if request.Month > 0 {
			cardStatement.Month = request.Month
		}
		
		responses[i] = cardStatement.ToResponse()
	}

	return responses, nil
}