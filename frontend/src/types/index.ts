// モデル関連の型をすべてエクスポート
export * from './models';

// API関連の型をエクスポート
export * from './api/generated';

export type CsrfToken = {
  csrf_token: string
}
