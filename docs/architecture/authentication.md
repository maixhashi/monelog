# 認証フロー

## 概要

Monelog では、JWT (JSON Web Token) を使用したトークンベースの認証システムを採用しています。このドキュメントでは、ユーザー認証のフローとセキュリティ対策について説明します。

## 設計原則

- **ステートレス認証**: サーバー側でセッション状態を保持しない
- **トークン有効期限**: セキュリティリスク軽減のための適切な有効期限設定
- **CSRF対策**: クロスサイトリクエストフォージェリ対策の実装
- **最小権限の原則**: 必要最小限の権限のみを付与

## 構成図

```
+---------------+       +---------------+       +---------------+
|               |  1    |               |  2    |               |
|  クライアント  |------>|  認証エンドポイント |------>|  JWT生成     |
|               |       |               |       |               |
+---------------+       +---------------+       +---------------+
        ^                                               |
        |                                               |
        |  4                                            |  3
        |                                               v
+---------------+       +---------------+       +---------------+
|               |       |               |       |               |
|  保護されたリソース |<------|  JWT検証     |<------|  トークン返却  |
|               |       |               |       |               |
+---------------+       +---------------+       +---------------+
```

## 主要コンポーネント

### JWT認証
- **トークン生成**: ユーザーIDを含むJWTトークンの生成
- **トークン検証**: リクエスト時のJWTトークン検証
- **ミドルウェア**: 保護されたエンドポイントへのアクセス制御

### CSRF対策
- **CSRFトークン**: 状態変更操作に対するCSRFトークンの要求
- **CSRFトークン生成API**: クライアント用のCSRFトークン取得エンドポイント

### ユーザー認証
- **サインアップ**: 新規ユーザー登録処理
- **ログイン**: 既存ユーザーの認証処理
- **パスワード管理**: ハッシュ化されたパスワードの保存と検証

## 技術的な詳細

### 認証フロー
1. **ユーザーログイン**:
   - クライアントがメールアドレスとパスワードを送信
   - サーバーがユーザーを検証し、パスワードをハッシュ比較
   - 認証成功時にJWTトークンを生成

2. **トークン使用**:
   - クライアントは以降のリクエストにJWTトークンを含める
   - サーバーはリクエストごとにトークンを検証
   - 有効なトークンの場合のみ処理を続行

3. **CSRF保護**:
   - クライアントはCSRFトークンをAPIから取得
   - 状態変更リクエスト時にCSRFトークンをヘッダーに含める
   - サーバーはCSRFトークンを検証

### 実装コード例

```go
// getUserIdFromToken はJWTトークンからユーザーIDを取得する関数
func getUserIdFromToken(c echo.Context) (uint, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))
	return userId, nil
}
```

```typescript
// CSRFトークンを取得する関数
export const getCsrfToken = async (): Promise<string> => {
  const response = await axios.get(`${process.env.REACT_APP_API_URL}/csrf-token`);
  return response.data.csrf_token;
};
```

### テスト例

```go
func setupEchoWithJWT(userId uint) (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	
	// JWTクレームを設定
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = float64(userId)
	
	// コンテキストにトークンを設定
	c.Set("user", token)
	
	return e, c, rec
}
```

## 将来の拡張性

- **OAuth/OpenID Connect**: 外部認証プロバイダとの統合
- **多要素認証**: セキュリティ強化のための多要素認証の導入
- **権限管理**: より詳細な権限管理システムの実装
- **アクセストークン/リフレッシュトークン**: 2種類のトークンによるセキュリティ強化
- **セッションの無効化**: 不審なアクティビティ検出時のセッション強制終了機能