package csv_history_test

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestSaveCSVHistory(t *testing.T) {
	setupCSVHistoryControllerTest()
	
	t.Run("正常系", func(t *testing.T) {
		t.Run("有効なCSVファイルをアップロードする場合", func(t *testing.T) {
			// マルチパートフォームを作成
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			
			// ファイルフィールドを追加
			fileField, err := writer.CreateFormFile("file", "test.csv")
			if err != nil {
				t.Fatalf("ファイルフィールドの作成に失敗しました: %v", err)
			}
			
			// CSVデータを書き込む
			fileField.Write([]byte("header1,header2\nvalue1,value2"))
			
			// その他のフォームフィールドを追加
			writer.WriteField("file_name", "test.csv")
			writer.WriteField("card_type", "rakuten")
			writer.WriteField("year", "2023")
			writer.WriteField("month", "1")
			
			// フォームを閉じる
			writer.Close()
			
			// リクエストとレスポンスを作成
			req := httptest.NewRequest(http.MethodPost, "/csv-histories", body)
			req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			setUserContext(c, csvHistoryTestUser)
			
			// コントローラーを実行
			err = csvHistoryController.SaveCSVHistory(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			assertStatusCode(t, rec, http.StatusCreated)
			
			// レスポンスを検証
			var response map[string]interface{}
			if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
				t.Fatalf("レスポンスのJSONパースに失敗しました: %v", err)
			}
			
			// IDが返されていることを確認
			if _, exists := response["id"]; !exists {
				t.Errorf("レスポンスにIDフィールドがありません: %v", response)
			}
			
			// ファイル名が正しいことを確認
			if fileName, exists := response["file_name"]; !exists || fileName != "test.csv" {
				t.Errorf("ファイル名が一致しません: got=%v, want=test.csv", fileName)
			}
		})
	})
	
	t.Run("異常系", func(t *testing.T) {
		t.Run("ファイルが提供されていない場合", func(t *testing.T) {
			// マルチパートフォームを作成（ファイルなし）
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			
			// その他のフォームフィールドを追加
			writer.WriteField("file_name", "test.csv")
			writer.WriteField("card_type", "rakuten")
			writer.WriteField("year", "2023")
			writer.WriteField("month", "1")
			
			// フォームを閉じる
			writer.Close()
			
			// リクエストとレスポンスを作成
			// リクエストとレスポンスを作成
			req := httptest.NewRequest(http.MethodPost, "/csv-histories", body)
			req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			setUserContext(c, csvHistoryTestUser)
			
			// コントローラーを実行
			err := csvHistoryController.SaveCSVHistory(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			validateErrorResponse(t, rec, http.StatusBadRequest)
		})
		
		t.Run("必須フィールドが不足している場合", func(t *testing.T) {
			// マルチパートフォームを作成
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			
			// ファイルフィールドを追加
			fileField, err := writer.CreateFormFile("file", "test.csv")
			if err != nil {
				t.Fatalf("ファイルフィールドの作成に失敗しました: %v", err)
			}
			
			// CSVデータを書き込む
			fileField.Write([]byte("header1,header2\nvalue1,value2"))
			
			// 一部のフィールドを省略
			writer.WriteField("file_name", "test.csv")
			// card_typeを省略
			writer.WriteField("year", "2023")
			writer.WriteField("month", "1")
			
			// フォームを閉じる
			writer.Close()
			
			// リクエストとレスポンスを作成
			req := httptest.NewRequest(http.MethodPost, "/csv-histories", body)
			req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			setUserContext(c, csvHistoryTestUser)
			
			// コントローラーを実行
			err = csvHistoryController.SaveCSVHistory(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			validateErrorResponse(t, rec, http.StatusBadRequest)
		})
		
		t.Run("無効なカードタイプの場合", func(t *testing.T) {
			// マルチパートフォームを作成
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			
			// ファイルフィールドを追加
			fileField, err := writer.CreateFormFile("file", "test.csv")
			if err != nil {
				t.Fatalf("ファイルフィールドの作成に失敗しました: %v", err)
			}
			
			// CSVデータを書き込む
			fileField.Write([]byte("header1,header2\nvalue1,value2"))
			
			// 無効なカードタイプを設定
			writer.WriteField("file_name", "test.csv")
			writer.WriteField("card_type", "invalid_card")
			writer.WriteField("year", "2023")
			writer.WriteField("month", "1")
			
			// フォームを閉じる
			writer.Close()
			
			// リクエストとレスポンスを作成
			req := httptest.NewRequest(http.MethodPost, "/csv-histories", body)
			req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			setUserContext(c, csvHistoryTestUser)
			
			// コントローラーを実行
			err = csvHistoryController.SaveCSVHistory(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			// 実際のステータスコードを確認
			if rec.Code != http.StatusBadRequest && rec.Code != http.StatusInternalServerError {
				t.Errorf("予期しないステータスコード: got=%d, want=%d or %d", 
					rec.Code, http.StatusBadRequest, http.StatusInternalServerError)
			}
		})
		
		t.Run("無効な年月の場合", func(t *testing.T) {
			// マルチパートフォームを作成
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			
			// ファイルフィールドを追加
			fileField, err := writer.CreateFormFile("file", "test.csv")
			if err != nil {
				t.Fatalf("ファイルフィールドの作成に失敗しました: %v", err)
			}
			
			// CSVデータを書き込む
			fileField.Write([]byte("header1,header2\nvalue1,value2"))
			
			// 無効な月を設定
			writer.WriteField("file_name", "test.csv")
			writer.WriteField("card_type", "rakuten")
			writer.WriteField("year", "2023")
			writer.WriteField("month", "13") // 無効な月
			
			// フォームを閉じる
			writer.Close()
			
			// リクエストとレスポンスを作成
			req := httptest.NewRequest(http.MethodPost, "/csv-histories", body)
			req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			setUserContext(c, csvHistoryTestUser)
			
			// コントローラーを実行
			err = csvHistoryController.SaveCSVHistory(c)
			
			// アサーション
			if err != nil {
				t.Fatalf("コントローラーの実行に失敗しました: %v", err)
			}
			
			// 実際のステータスコードを確認
			if rec.Code != http.StatusBadRequest && rec.Code != http.StatusInternalServerError {
				t.Errorf("予期しないステータスコード: got=%d, want=%d or %d", 
					rec.Code, http.StatusBadRequest, http.StatusInternalServerError)
			}
		})
	})
}
