package card_statement_test

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCardStatementController_PreviewCSV(t *testing.T) {
	setupCardStatementControllerTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("CSVファイルを正常にプレビューする", func(t *testing.T) {
			// このテストはモックが必要なため、実際のCSVファイルの処理はスキップ
			t.Skip("このテストは実際のCSVファイルが必要なため、スキップします")
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("ファイルが指定されていない場合", func(t *testing.T) {
			// テスト実行
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/card-statements/preview", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// JWTクレームを設定
			setupJWTToken(c, cardStatementTestUser.ID)

			err := cardStatementController.PreviewCSV(c)

			// エラーは返さないが、内部でエラーハンドリングしてステータスコードを返す
			if err != nil {
				t.Errorf("PreviewCSV() unexpected error = %v", err)
			}

			if rec.Code != http.StatusBadRequest {
				t.Errorf("PreviewCSV() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})

		t.Run("card_typeが指定されていない場合", func(t *testing.T) {
			// マルチパートフォームを作成
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			
			// ダミーファイルを追加
			part, err := writer.CreateFormFile("file", "test.csv")
			if err != nil {
				t.Fatalf("Failed to create form file: %v", err)
			}
			part.Write([]byte("dummy,csv,content"))
			writer.Close()

			// テスト実行
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/card-statements/preview", body)
			req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// JWTクレームを設定
			setupJWTToken(c, cardStatementTestUser.ID)

			err = cardStatementController.PreviewCSV(c)

			// エラーは返さないが、内部でエラーハンドリングしてステータスコードを返す
			if err != nil {
				t.Errorf("PreviewCSV() unexpected error = %v", err)
			}

			if rec.Code != http.StatusBadRequest {
				t.Errorf("PreviewCSV() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})

		t.Run("無効なcard_typeが指定された場合", func(t *testing.T) {
			// マルチパートフォームを作成
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			
			// ダミーファイルを追加
			part, err := writer.CreateFormFile("file", "test.csv")
			if err != nil {
				t.Fatalf("Failed to create form file: %v", err)
			}
			part.Write([]byte("dummy,csv,content"))
			
			// 無効なcard_typeを追加
			writer.WriteField("card_type", "invalid_card_type")
			writer.Close()

			// テスト実行
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/card-statements/preview", body)
			req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// JWTクレームを設定
			setupJWTToken(c, cardStatementTestUser.ID)

			err = cardStatementController.PreviewCSV(c)

			// エラーは返さないが、内部でエラーハンドリングしてステータスコードを返す
			if err != nil {
				t.Errorf("PreviewCSV() unexpected error = %v", err)
			}

			if rec.Code != http.StatusInternalServerError {
				t.Errorf("PreviewCSV() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
	})
}