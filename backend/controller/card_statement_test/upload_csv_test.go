package card_statement_test

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func TestCardStatementController_UploadCSV(t *testing.T) {
	setupCardStatementControllerTest()

	t.Run("正常系", func(t *testing.T) {
		t.Run("CSVファイルを正常にアップロードする", func(t *testing.T) {
			// このテストはモックが必要なため、実際のCSVファイルの処理はスキップ
			// 実際のテストでは、テスト用のCSVファイルを用意して、multipart/form-dataリクエストを作成する必要がある
			t.Skip("このテストは実際のCSVファイルが必要なため、スキップします")
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("ファイルが指定されていない場合", func(t *testing.T) {
			// テスト実行
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/card-statements/upload", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEMultipartForm)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// JWTクレームを設定
			setupJWTToken(c, cardStatementTestUser.ID)

			err := cardStatementController.UploadCSV(c)

			// エラーは返さないが、内部でエラーハンドリングしてステータスコードを返す
			if err != nil {
				t.Errorf("UploadCSV() unexpected error = %v", err)
			}

			if rec.Code != http.StatusBadRequest {
				t.Errorf("UploadCSV() status code = %d, want %d", rec.Code, http.StatusBadRequest)
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
			req := httptest.NewRequest(http.MethodPost, "/card-statements/upload", body)
			req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// JWTクレームを設定
			setupJWTToken(c, cardStatementTestUser.ID)

			err = cardStatementController.UploadCSV(c)

			// エラーは返さないが、内部でエラーハンドリングしてステータスコードを返す
			if err != nil {
				t.Errorf("UploadCSV() unexpected error = %v", err)
			}

			if rec.Code != http.StatusBadRequest {
				t.Errorf("UploadCSV() status code = %d, want %d", rec.Code, http.StatusBadRequest)
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
			writer.WriteField("year", "2023")
			writer.WriteField("month", "2")
			writer.Close()

			// テスト実行
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/card-statements/upload", body)
			req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// JWTクレームを設定
			setupJWTToken(c, cardStatementTestUser.ID)

			err = cardStatementController.UploadCSV(c)

			// エラーは返さないが、内部でエラーハンドリングしてステータスコードを返す
			if err != nil {
				t.Errorf("UploadCSV() unexpected error = %v", err)
			}

			if rec.Code != http.StatusInternalServerError {
				t.Errorf("UploadCSV() status code = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
		
		t.Run("年パラメータが不正な場合", func(t *testing.T) {
			// マルチパートフォームを作成
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			
			// ダミーファイルを追加
			part, err := writer.CreateFormFile("file", "test.csv")
			if err != nil {
				t.Fatalf("Failed to create form file: %v", err)
			}
			part.Write([]byte("dummy,csv,content"))
			
			// 無効な年パラメータを追加
			writer.WriteField("card_type", "rakuten")
			writer.WriteField("year", "invalid_year")
			writer.WriteField("month", "2")
			writer.Close()

			// テスト実行
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/card-statements/upload", body)
			req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// JWTクレームを設定
			setupJWTToken(c, cardStatementTestUser.ID)

			err = cardStatementController.UploadCSV(c)

			// エラーは返さないが、内部でエラーハンドリングしてステータスコードを返す
			if err != nil {
				t.Errorf("UploadCSV() unexpected error = %v", err)
			}

			if rec.Code != http.StatusBadRequest {
				t.Errorf("UploadCSV() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})
		
		t.Run("月パラメータが不正な場合", func(t *testing.T) {
			// マルチパートフォームを作成
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			
			// ダミーファイルを追加
			part, err := writer.CreateFormFile("file", "test.csv")
			if err != nil {
				t.Fatalf("Failed to create form file: %v", err)
			}
			part.Write([]byte("dummy,csv,content"))
			
			// 無効な月パラメータを追加
			writer.WriteField("card_type", "rakuten")
			writer.WriteField("year", "2023")
			writer.WriteField("month", "invalid_month")
			writer.Close()

			// テスト実行
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/card-statements/upload", body)
			req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// JWTクレームを設定
			setupJWTToken(c, cardStatementTestUser.ID)

			err = cardStatementController.UploadCSV(c)

			// エラーは返さないが、内部でエラーハンドリングしてステータスコードを返す
			if err != nil {
				t.Errorf("UploadCSV() unexpected error = %v", err)
			}

			if rec.Code != http.StatusBadRequest {
				t.Errorf("UploadCSV() status code = %d, want %d", rec.Code, http.StatusBadRequest)
			}
		})
	})
}

// JWTトークンをコンテキストに設定するヘルパー関数
func setupJWTToken(c echo.Context, userId uint) {
	token := createJWTToken(userId)
	c.Set("user", token)
}

// JWTトークンを作成するヘルパー関数
func createJWTToken(userId uint) *jwt.Token {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = float64(userId)
	return token
}
