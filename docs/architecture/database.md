# データベース設計

## 概要

Monelog のデータベースは MySQL を使用し、GORM を通じてアクセスされます。このドキュメントでは、データベースのスキーマ設計、テーブル構造、およびリレーションシップについて説明します。

## 設計原則

- **正規化**: 適切な正規化によるデータの一貫性確保
- **インデックス最適化**: クエリパフォーマンスのためのインデックス設計
- **外部キー制約**: データの整合性を保証
- **ソフトデリート**: データの論理削除によるデータ履歴の保持

## 構成図

```
+---------------+       +---------------+       +---------------+
|               |       |               |       |               |
|    users      |<----->|    tasks      |       | card_statements|
|               |       |               |       |               |
+---------------+       +---------------+       +---------------+
        |                                               |
        |                                               |
        v                                               v
+---------------+                               +---------------+
|               |                               |               |
|  user_settings |                               | categories    |
|               |                               |               |
+---------------+                               +---------------+
```

## 主要コンポーネント

### users テーブル
ユーザー情報を管理するテーブル
- `id`: プライマリキー (uint)
- `email`: ユーザーのメールアドレス (string, ユニーク)
- `password`: ハッシュ化されたパスワード (string)
- `created_at`: 作成日時 (time.Time)
- `updated_at`: 更新日時 (time.Time)
- `deleted_at`: 削除日時 (time.Time, ソフトデリート用)

### tasks テーブル
ユーザーのタスク情報を管理するテーブル
- `id`: プライマリキー (uint)
- `user_id`: ユーザーID (uint, 外部キー)
- `title`: タスクのタイトル (string)
- `created_at`: 作成日時 (time.Time)
- `updated_at`: 更新日時 (time.Time)
- `deleted_at`: 削除日時 (time.Time, ソフトデリート用)

### card_statements テーブル
クレジットカードの明細情報を管理するテーブル
- `id`: プライマリキー (uint)
- `user_id`: ユーザーID (uint, 外部キー)
- `date`: 取引日 (time.Time)
- `amount`: 金額 (int)
- `description`: 取引内容 (string)
- `category_id`: カテゴリID (uint, 外部キー)
- `created_at`: 作成日時 (time.Time)
- `updated_at`: 更新日時 (time.Time)

### categories テーブル
支出カテゴリを管理するテーブル
- `id`: プライマリキー (uint)
- `name`: カテゴリ名 (string)
- `user_id`: ユーザーID (uint, 外部キー、NULLの場合はシステム定義カテゴリ)
- `created_at`: 作成日時 (time.Time)
- `updated_at`: 更新日時 (time.Time)

## 技術的な詳細

### モデル定義例

```go
// User モデル
type User struct {
    gorm.Model
    Email    string `gorm:"uniqueIndex;not null"`
    Password string `gorm:"not null"`
    Tasks    []Task
}

// Task モデル
type Task struct {
    gorm.Model
    Title  string `gorm:"not null"`
    UserId uint   `gorm:"not null"`
    User   User
}

// ToResponse メソッド例
func (t *Task) ToResponse() TaskResponse {
    return TaskResponse{
        ID:        t.ID,
        Title:     t.Title,
        CreatedAt: t.CreatedAt,
        UpdatedAt: t.UpdatedAt,
    }
}
```

### インデックス
- `users.email`: ログイン時の検索効率化
- `tasks.user_id`: ユーザーごとのタスク検索効率化
- `card_statements.user_id`: ユーザーごとの明細検索効率化
- `card_statements.date`: 日付範囲検索効率化

### 外部キー制約
- `tasks.user_id` → `users.id`
- `card_statements.user_id` → `users.id`
- `card_statements.category_id` → `categories.id`

### マイグレーション
データベースのマイグレーションは GORM のマイグレーション機能を使用して管理されます。

```go
// マイグレーション例
db.AutoMigrate(&User{}, &Task{}, &CardStatement{}, &Category{})
```

## 将来の拡張性

- **パーティショニング**: 大量データに対応するためのテーブルパーティショニング
- **読み書き分離**: 読み取り専用レプリカの導入
- **シャーディング**: データ量増加に対応するためのシャーディング
- **NoSQL統合**: 特定のデータタイプに対するNoSQLデータベースの導入
- **監査ログ**: データ変更の履歴を記録するための監査ログテーブルの追加