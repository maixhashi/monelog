# CSV履歴機能

## 概要

CSV履歴機能は、ユーザーがクレジットカードの明細CSVファイルをアップロードし、管理するための機能です。アップロードされたCSVファイルは年月ごとに整理され、後で参照やダウンロードが可能です。

## 機能一覧

- CSV履歴の一覧取得
- 特定のCSV履歴の詳細取得
- 月別CSV履歴の取得
- CSVファイルのアップロードと保存
- CSV履歴の削除
- CSVファイルのダウンロード

## アーキテクチャ

この機能はClean Architectureに基づいて実装されています。

### モデル (Model)

`backend/model/csv_history.go`にCSV履歴に関連するモデルを定義しています。

- `CSVHistory`: データベースのCSV履歴モデル
- `CSVHistoryResponse`: CSV履歴のレスポンス（ファイルデータなし）
- `CSVHistoryDetailResponse`: CSV履歴の詳細レスポンス（ファイルデータ含む）
- `CSVHistorySaveRequest`: CSV履歴保存リクエスト

### リポジトリ (Repository)

`backend/repository/csv_history_repository.go`にデータベース操作を行うリポジトリを実装しています。

- `GetAllCSVHistories`: すべてのCSV履歴を取得
- `GetCSVHistoryById`: 特定のCSV履歴を取得
- `GetCSVHistoriesByMonth`: 月別のCSV履歴を取得
- `CreateCSVHistory`: CSV履歴を作成
- `DeleteCSVHistory`: CSV履歴を削除

### ユースケース (Usecase)

`backend/usecase/csv_history_usecase.go`にビジネスロジックを実装しています。

- `GetAllCSVHistories`: すべてのCSV履歴を取得
- `GetCSVHistoryById`: 特定のCSV履歴を取得
- `GetCSVHistoriesByMonth`: 月別のCSV履歴を取得
- `SaveCSVHistory`: CSVファイルを保存
- `DeleteCSVHistory`: CSV履歴を削除

### バリデーター (Validator)

`backend/validator/csv_history_validator.go`にリクエストの検証ロジックを実装しています。

- `ValidateCSVHistorySaveRequest`: CSV履歴保存リクエストの検証

### コントローラー (Controller)

`backend/controller/csv_history_controller.go`にHTTPリクエストを処理するコントローラーを実装しています。

- `GetAllCSVHistories`: すべてのCSV履歴を取得するエンドポイント
- `GetCSVHistoryById`: 特定のCSV履歴を取得するエンドポイント
- `GetCSVHistoriesByMonth`: 月別のCSV履歴を取得するエンドポイント
- `SaveCSVHistory`: CSVファイルを保存するエンドポイント
- `DeleteCSVHistory`: CSV履歴を削除するエンドポイント
- `DownloadCSVHistory`: CSVファイルをダウンロードするエンドポイント

## API仕様

### CSV履歴一覧の取得

```
GET /csv-histories
```

#### レスポンス例

```json
[
  {
    "id": 1,
    "file_name": "rakuten_202301.csv",
    "card_type": "rakuten",
    "year": 2023,
    "month": 1,
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
]
```

### 特定のCSV履歴の取得

```
GET /csv-histories/{csvHistoryId}
```

#### レスポンス例

```json
{
  "id": 1,
  "file_name": "rakuten_202301.csv",
  "card_type": "rakuten",
  "file_data": "...",
  "year": 2023,
  "month": 1,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### 月別CSV履歴の取得

```
GET /csv-histories/by-month?year=2023&month=1
```

#### レスポンス例

```json
[
  {
    "id": 1,
    "file_name": "rakuten_202301.csv",
    "card_type": "rakuten",
    "year": 2023,
    "month": 1,
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
]
```

### CSVファイルの保存

```
POST /csv-histories
Content-Type: multipart/form-data
```

#### リクエストパラメータ

- `file`: CSVファイル
- `file_name`: ファイル名
- `card_type`: カード種類（rakuten, mufg, epos）
- `year`: 年
- `month`: 月（1-12）

#### レスポンス例

```json
{
  "id": 1,
  "file_name": "rakuten_202301.csv",
  "card_type": "rakuten",
  "year": 2023,
  "month": 1,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### CSV履歴の削除

```
DELETE /csv-histories/{csvHistoryId}
```

#### レスポンス

- 204 No Content: 削除成功

### CSVファイルのダウンロード

```
GET /csv-histories/{csvHistoryId}/download
```

#### レスポンス

- CSVファイルのバイナリデータ