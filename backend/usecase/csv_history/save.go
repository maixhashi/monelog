package csv_history

import (
	"bytes"
	"io"
	"mime/multipart"
	"monelog/model"
)

func (chu *csvHistoryUsecase) SaveCSVHistory(file *multipart.FileHeader, request model.CSVHistorySaveRequest) (model.CSVHistoryResponse, error) {
	if err := chu.chv.ValidateCSVHistorySaveRequest(request); err != nil {
		return model.CSVHistoryResponse{}, err
	}

	// ファイルを開く
	src, err := file.Open()
	if err != nil {
		return model.CSVHistoryResponse{}, err
	}
	defer src.Close()

	// ファイルの内容を読み込む
	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, src); err != nil {
		return model.CSVHistoryResponse{}, err
	}

	// CSVHistoryモデルを作成
	csvHistory := model.CSVHistory{
		FileName: request.FileName,
		CardType: request.CardType,
		FileData: buf.Bytes(),
		UserId:   request.UserId,
		Year:     request.Year,   // 年を保存
		Month:    request.Month,  // 月を保存
	}

	// データベースに保存
	if err := chu.chr.CreateCSVHistory(&csvHistory); err != nil {
		return model.CSVHistoryResponse{}, err
	}

	return csvHistory.ToResponse(), nil
}