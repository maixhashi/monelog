package card_statement_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"monelog/dto"  // dtoパッケージをインポート
	"monelog/model"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWTトークンをセットアップするヘルパー関数
func setupJWTToken(c echo.Context, userId uint) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = float64(userId)
	c.Set("user", token)
}

// JWT認証をモックするヘルパー関数
func setupEchoWithJWT(userId uint) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// JWTクレームを設定
	setupJWTToken(c, userId)
	
	return e, c, rec
}

// リクエストボディ付きのコンテキストを設定するヘルパー関数
func setupEchoWithJWTAndBody(userId uint, method, path, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, bytes.NewReader([]byte(body)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// JWTクレームを設定
	setupJWTToken(c, userId)
	
	return e, c, rec
}

// CardStatementIDパラメータを持つリクエストコンテキストを設定するヘルパー関数
func setupEchoWithCardStatementId(userId uint, cardStatementId uint, method, path, body string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e, c, rec := setupEchoWithJWTAndBody(userId, method, path, body)
	c.SetParamNames("cardStatementId")
	c.SetParamValues(fmt.Sprintf("%d", cardStatementId))
	return e, c, rec
}

// クエリパラメータ付きのコンテキストを設定するヘルパー関数
func setupEchoWithJWTAndQuery(userId uint, method, path string, queryParams map[string]string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, nil)
	
	// クエリパラメータを追加
	q := req.URL.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// JWTクレームを設定
	setupJWTToken(c, userId)
	
	return e, c, rec
}

// マルチパートフォーム付きのコンテキストを設定するヘルパー関数
func setupEchoWithJWTAndMultipartForm(userId uint, method, path string, formData []byte, contentType string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, bytes.NewReader(formData))
	req.Header.Set(echo.HeaderContentType, contentType)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// JWTクレームを設定
	setupJWTToken(c, userId)
	
	return e, c, rec
}

// テスト用CSVファイルを作成するヘルパー関数
func createTestCSVFile(t *testing.T, content string) (*os.File, error) {
	tmpFile, err := os.CreateTemp("", "test-*.csv")
	if err != nil {
		return nil, err
	}
	
	if _, err := tmpFile.WriteString(content); err != nil {
		return nil, err
	}
	
	if err := tmpFile.Close(); err != nil {
		return nil, err
	}
	
	return tmpFile, nil
}

// マルチパートフォームデータを作成するヘルパー関数
func createMultipartFormData(t *testing.T, fieldName, fileName string, fileContent []byte) ([]byte, string) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	
	// ファイルフィールドを追加
	fw, err := w.CreateFormFile(fieldName, fileName)
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	
	if _, err := io.Copy(fw, bytes.NewReader(fileContent)); err != nil {
		t.Fatalf("Failed to copy file content: %v", err)
	}
	
	// カード種類フィールドを追加
	if err := w.WriteField("card_type", "rakuten"); err != nil {
		t.Fatalf("Failed to add card_type field: %v", err)
	}
	
	// 年月フィールドを追加
	if err := w.WriteField("year", "2023"); err != nil {
		t.Fatalf("Failed to add year field: %v", err)
	}
	
	if err := w.WriteField("month", "4"); err != nil {
		t.Fatalf("Failed to add month field: %v", err)
	}
	
	w.Close()
	
	return b.Bytes(), w.FormDataContentType()
}

// テスト用カード明細を作成するヘルパー関数
func createTestCardStatement(cardType, description string, userId uint, year, month int) *model.CardStatement {
	cardStatement := &model.CardStatement{
		Type:              "発生",
		StatementNo:       1,
		CardType:          cardType,
		Description:       description,
		UseDate:           "2023/04/01",
		PaymentDate:       "2023/05/27",
		PaymentMonth:      "2023年05月",
		Amount:            5000,
		TotalChargeAmount: 5000,
		ChargeAmount:      0,
		RemainingBalance:  5000,
		PaymentCount:      0,
		InstallmentCount:  1,
		AnnualRate:        0.0,
		MonthlyRate:       0.0,
		UserId:            userId,
		Year:              year,
		Month:             month,
	}
	cardStatementDB.Create(cardStatement)
	return cardStatement
}

// レスポンスボディをパースするヘルパー関数
func parseCardStatementResponse(t *testing.T, responseBody []byte) dto.CardStatementResponse {
	var response dto.CardStatementResponse
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	return response
}

// 複数カード明細のレスポンスボディをパースするヘルパー関数
func parseCardStatementsResponse(t *testing.T, responseBody []byte) []dto.CardStatementResponse {
	var response []dto.CardStatementResponse
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	return response
}
