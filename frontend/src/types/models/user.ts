// ユーザー関連の型定義

// ユーザーモデル
export type User = {
  id: number;
  email: string;
};

// ログイン・サインアップ用のリクエスト型
export type UserCredential = {
  email: string;
  password: string;
};

// 編集中のユーザー情報（Zustand用）
export type EditedUser = {
  id: number;
  email: string;
};

// 初期状態
export const initialEditedUser: EditedUser = {
  id: 0,
  email: '',
};
