# カード明細CSV解析機能

## 概要

この機能は、ユーザーがクレジットカードの明細CSVファイルをアップロードし、システムがそれを解析して保存する機能です。現在、楽天カード、三菱UFJカード、エポスカードの3種類のCSVフォーマットに対応しています。また、CSVデータのプレビュー機能や支払月ごとの明細取得機能も提供しています。

## アーキテクチャ

この機能はClean Architectureのパターンに従って実装されています。

### レイヤー構成

1. **Controller層** - HTTPリクエストの受け取りとレスポンスの返却
2. **Usecase層** - ビジネスロジックの実装
3. **Repository層** - データベースとのやり取り
4. **Validator層** - リクエストデータの検証
5. **Model層** - データモデルの定義

## 実装詳細

### Model層 (`backend/model/card_statement.go`)

カード明細に関する以下のデータモデルを定義しています：

- `CardStatement` - データベースに保存されるカード明細のモデル
- `CardStatementRequest` - CSVアップロードのリクエストモデル
- `CardStatementResponse` - APIレスポンスのモデル
- `CardStatementSummary` - CSVから解析した明細データの中間モデル
- `CardStatementPreviewRequest` - CSVプレビュー用リクエストモデル
- `CardStatementSaveRequest` - 一時データを保存するリクエストモデル
- `CardStatementByMonthRequest` - 支払月ごとのカード明細取得リクエストモデル

また、モデル間の変換メソッドも提供しています：
- `ToResponse()` - `CardStatement`から`CardStatementResponse`への変換
- `ToModel()` - `CardStatementSummary`から`CardStatement`への変換

### Repository層 (`backend/repository/card_statement_repository.go`)

データベースとのやり取りを担当するレイヤーです。以下の機能を提供しています：

- `GetAllCardStatements` - ユーザーのすべてのカード明細を取得
- `GetCardStatementById` - 特定のカード明細を取得
- `CreateCardStatement` - 単一のカード明細を作成
- `CreateCardStatements` - 複数のカード明細を一括作成
- `DeleteCardStatements` - ユーザーのカード明細を削除
- `GetCardStatementsByMonth` - 指定された年月の支払いに関するカード明細を取得

### Validator層 (`backend/validator/card_statement_validator.go`)

リクエストデータの検証を行うレイヤーです。以下の検証を実装しています：

- `ValidateCardStatementRequest` - カード明細リクエストの検証
  - `CardType`が必須であること
  - `CardType`が許可されている値（rakuten, mufg, epos）のいずれかであること
- `ValidateCardStatementPreviewRequest` - CSVプレビューリクエストの検証
- `ValidateCardStatementSaveRequest` - カード明細保存リクエストの検証
- `ValidateCardStatementByMonthRequest` - 支払月ごとのカード明細取得リクエストの検証
  - `Year`が必須であること
  - `Month`が必須であり、1から12の間であること

### Usecase層 (`backend/usecase/card_statement_usecase.go`)

ビジネスロジックを実装するレイヤーです。以下の機能を提供しています：

- `GetAllCardStatements` - ユーザーのすべてのカード明細を取得
- `GetCardStatementById` - 特定のカード明細を取得
- `ProcessCSV` - CSVファイルを処理して解析・保存
  1. リクエストの検証
  2. CSVファイルの読み込み
  3. カード種類に応じたパーサーの取得
  4. CSVの解析
  5. 既存データの削除
  6. 解析結果のデータベースへの保存
  7. レスポンスの作成
- `PreviewCSV` - CSVファイルを解析してプレビューデータを返す（DBには保存しない）
- `SaveCardStatements` - プレビューしたカード明細データをDBに保存
- `GetCardStatementsByMonth` - 指定された年月の支払いに関するカード明細を取得

### Controller層 (`backend/controller/card_statement_controller.go`)

HTTPリクエストを受け取り、適切なUsecaseの処理を呼び出し、結果をクライアントに返すレイヤーです。以下のエンドポイントを提供しています：

- `GET /card-statements` - ユーザーのすべてのカード明細を取得
- `GET /card-statements/{cardStatementId}` - 特定のカード明細を取得
- `POST /card-statements/upload` - CSVファイルをアップロードして解析・保存
- `POST /card-statements/preview` - CSVファイルをアップロードして解析（プレビュー、保存なし）
- `POST /card-statements/save` - プレビューしたカード明細データを保存
- `GET /card-statements/by-month` - 支払月ごとのカード明細を取得

## 使用方法

### CSVファイルのアップロード

```
POST /card-statements/upload
```

**リクエストパラメータ**:
- `file` (multipart/form-data): CSVファイル
- `card_type` (form-data): カード種類 (rakuten, mufg, epos)
- `year` (form-data): 年 (例: 2023)
- `month` (form-data): 月 (1-12)

**レスポンス**:
- 成功時: 201 Created と解析されたカード明細の配列
- 失敗時: エラーメッセージを含む適切なHTTPステータスコード

### CSVファイルのプレビュー

```
POST /card-statements/preview
```

**リクエストパラメータ**:
- `file` (multipart/form-data): CSVファイル
- `card_type` (form-data): カード種類 (rakuten, mufg, epos)

**レスポンス**:
- 成功時: 200 OK と解析されたカード明細の配列（DBには保存されない）
- 失敗時: エラーメッセージを含む適切なHTTPステータスコード

### プレビューしたカード明細の保存

```
POST /card-statements/save
```

**リクエストボディ**:
```json
{
  "card_statements": [/* カード明細の配列 */],
  "card_type": "rakuten"
}
```

**レスポンス**:
- 成功時: 201 Created と保存されたカード明細の配列
- 失敗時: エラーメッセージを含む適切なHTTPステータスコード

### すべてのカード明細の取得

```
GET /card-statements
```

**レスポンス**:
- 成功時: 200 OK とカード明細の配列
- 失敗時: エラーメッセージを含む適切なHTTPステータスコード

### 特定のカード明細の取得

```
GET /card-statements/{cardStatementId}
```

**パスパラメータ**:
- `cardStatementId`: 取得するカード明細のID

**レスポンス**:
- 成功時: 200 OK とカード明細のオブジェクト
- 失敗時: エラーメッセージを含む適切なHTTPステータスコード

### 支払月ごとのカード明細の取得

```
GET /card-statements/by-month?year=2023&month=4
```

**クエリパラメータ**:
- `year`: 年 (例: 2023)
- `month`: 月 (1-12)

**レスポンス**:
- 成功時: 200 OK と指定された支払月のカード明細の配列
- 失敗時: エラーメッセージを含む適切なHTTPステータスコード
