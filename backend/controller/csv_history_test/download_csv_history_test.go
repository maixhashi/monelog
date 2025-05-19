package csv_history_test

import (
	"fmt"
	"monelog/model"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestDownloadCSVHistory(t *testing.T) {
    setupCSVHistoryControllerTest()
    
    t.Run("正常系", func(t *testing.T) {
        t.Run("自分のCSV履歴をダウンロードする場合", func(t *testing.T) {
            // テスト用のCSV履歴を作成
            csvHistory := createTestCSVHistory(t, "test_download.csv", csvHistoryTestUser.ID, 2023, 1)
            
            // リクエストとレスポンスを作成
            rec, c := createRequestResponse(http.MethodGet, fmt.Sprintf("/csv-histories/%d/download", csvHistory.ID))
            setUserContext(c, csvHistoryTestUser)
            c.SetParamNames("csvHistoryId")
            c.SetParamValues(fmt.Sprintf("%d", csvHistory.ID))
            
            // コントローラーを実行
            err := csvHistoryController.DownloadCSVHistory(c)
            
            // アサーション
            if err != nil {
                t.Fatalf("コントローラーの実行に失敗しました: %v", err)
            }
            
            assertStatusCode(t, rec, http.StatusOK)
            
            // Content-Dispositionヘッダーを確認
            contentDisposition := rec.Header().Get(echo.HeaderContentDisposition)
            if contentDisposition == "" {
                t.Errorf("Content-Dispositionヘッダーがありません")
            }
            
            // Content-Typeヘッダーを確認
            contentType := rec.Header().Get(echo.HeaderContentType)
            if contentType != "text/csv" {
                t.Errorf("Content-Typeが一致しません: got=%s, want=text/csv", contentType)
            }
            
            // レスポンスボディを確認
            if len(rec.Body.Bytes()) == 0 {
                t.Errorf("レスポンスボディが空です")
            }
        })
    })
    
    t.Run("異常系", func(t *testing.T) {
        t.Run("存在しないCSV履歴をダウンロードしようとした場合", func(t *testing.T) {
            // リクエストとレスポンスを作成
            rec, c := createRequestResponse(http.MethodGet, fmt.Sprintf("/csv-histories/%d/download", nonExistentCSVHistoryID))
            setUserContext(c, csvHistoryTestUser)
            c.SetParamNames("csvHistoryId")
            c.SetParamValues(fmt.Sprintf("%d", nonExistentCSVHistoryID))
            
            // コントローラーを実行
            err := csvHistoryController.DownloadCSVHistory(c)
            
            // アサーション
            if err != nil {
                t.Fatalf("コントローラーの実行に失敗しました: %v", err)
            }
            
            // 実際のステータスコードを確認
            if rec.Code != http.StatusNotFound && rec.Code != http.StatusInternalServerError {
                t.Errorf("予期しないステータスコード: got=%d, want=%d or %d", 
                    rec.Code, http.StatusNotFound, http.StatusInternalServerError)
            }
        })
        
        t.Run("他のユーザーのCSV履歴をダウンロードしようとした場合", func(t *testing.T) {
            // 別のユーザーのCSV履歴を作成
            otherCsvHistory := createTestCSVHistory(t, "other_download.csv", csvHistoryOtherUser.ID, 2023, 1)
            
            // リクエストとレスポンスを作成
            rec, c := createRequestResponse(http.MethodGet, fmt.Sprintf("/csv-histories/%d/download", otherCsvHistory.ID))
            setUserContext(c, csvHistoryTestUser)
            c.SetParamNames("csvHistoryId")
            c.SetParamValues(fmt.Sprintf("%d", otherCsvHistory.ID))
            
            // コントローラーを実行
            err := csvHistoryController.DownloadCSVHistory(c)
            
            // アサーション
            if err != nil {
                t.Fatalf("コントローラーの実行に失敗しました: %v", err)
            }
            
            // 実際のステータスコードを確認
            if rec.Code != http.StatusForbidden && 
               rec.Code != http.StatusInternalServerError && 
               rec.Code != http.StatusNotFound {
                t.Errorf("予期しないステータスコード: got=%d, want=%d, %d or %d", 
                    rec.Code, http.StatusForbidden, http.StatusInternalServerError, http.StatusNotFound)
            }
        })
        
        t.Run("無効なIDパラメータの場合", func(t *testing.T) {
            // リクエストとレスポンスを作成
            rec, c := createRequestResponse(http.MethodGet, "/csv-histories/invalid/download")
            setUserContext(c, csvHistoryTestUser)
            c.SetParamNames("csvHistoryId")
            c.SetParamValues("invalid")
            
            // コントローラーを実行
            err := csvHistoryController.DownloadCSVHistory(c)
            
            // アサーション
            if err != nil {
                t.Fatalf("コントローラーの実行に失敗しました: %v", err)
            }
            
            assertStatusCode(t, rec, http.StatusBadRequest)
        })
        
        t.Run("CSV履歴のファイルデータが空の場合", func(t *testing.T) {
            // 空のファイルデータを持つCSV履歴を作成
            emptyDataCsvHistory := model.CSVHistory{
                FileName: "empty.csv",
                CardType: "rakuten",
                FileData: []byte{}, // 空のデータ
                Year:     2023,
                Month:    1,
                UserId:   csvHistoryTestUser.ID,
            }
            
            result := csvHistoryDB.Create(&emptyDataCsvHistory)
            if result.Error != nil {
                t.Fatalf("テストCSV履歴の作成に失敗しました: %v", result.Error)
            }
            
            // リクエストとレスポンスを作成
            rec, c := createRequestResponse(http.MethodGet, fmt.Sprintf("/csv-histories/%d/download", emptyDataCsvHistory.ID))
            setUserContext(c, csvHistoryTestUser)
            c.SetParamNames("csvHistoryId")
            c.SetParamValues(fmt.Sprintf("%d", emptyDataCsvHistory.ID))
            
            // コントローラーを実行
            err := csvHistoryController.DownloadCSVHistory(c)
            
            // アサーション
            if err != nil {
                t.Fatalf("コントローラーの実行に失敗しました: %v", err)
            }
            
            // 空のファイルでも200 OKが返されることを確認
            assertStatusCode(t, rec, http.StatusOK)
            
            // Content-Dispositionヘッダーを確認
            contentDisposition := rec.Header().Get(echo.HeaderContentDisposition)
            if contentDisposition == "" {
                t.Errorf("Content-Dispositionヘッダーがありません")
            }
            
            // Content-Typeヘッダーを確認
            contentType := rec.Header().Get(echo.HeaderContentType)
            if contentType != "text/csv" {
                t.Errorf("Content-Typeが一致しません: got=%s, want=text/csv", contentType)
            }
            
            // レスポンスボディが空であることを確認
            if len(rec.Body.Bytes()) != 0 {
                t.Errorf("レスポンスボディが空ではありません: %s", rec.Body.String())
            }
        })
    })
}
