package testutils

import "monelog/model"

// GenerateValidTitle は最大文字数制限内の有効なタイトルを生成します
func GenerateValidTitle() string {
    return "ValidTitle" // model.TaskTitleMaxLength以下の長さ
}

// GenerateInvalidTitle は最大文字数制限を超える無効なタイトルを生成します
func GenerateInvalidTitle() string {
    invalidTitle := ""
    for i := 0; i <= model.TaskTitleMaxLength; i++ {
        invalidTitle += "x"
    }
    return invalidTitle
}
