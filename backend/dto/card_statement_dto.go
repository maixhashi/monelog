package dto

import (
	"monelog/dto/preview"
	"monelog/dto/statement"
	"monelog/dto/summary"
)

// CardStatementRequest カード明細のCSVアップロードリクエスト
type CardStatementRequest = statement.Request

// CardStatementResponse カード明細のレスポンス
type CardStatementResponse = statement.Response

// CardStatementSummary CSVから解析した明細データ
type CardStatementSummary = summary.Summary

// CardStatementPreviewRequest CSVプレビュー用リクエスト
type CardStatementPreviewRequest = preview.Request

// CardStatementSaveRequest 一時データを保存するリクエスト
type CardStatementSaveRequest = summary.SaveRequest

// CardStatementByMonthRequest 支払月ごとのカード明細取得リクエスト
type CardStatementByMonthRequest = statement.ByMonthRequest