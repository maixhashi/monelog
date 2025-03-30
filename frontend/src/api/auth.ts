import axios from 'axios';

export interface UserCredential {
  email: string;
  password: string;
}

export interface User {
  id?: number;
  email: string;
}

// バックエンドのエンドポイントに合わせて修正
export const register = async (credential: UserCredential): Promise<User> => {
  const response = await axios.post(`${process.env.REACT_APP_API_URL}/signup`, credential);
  return response.data;
};

export const login = async (credential: UserCredential): Promise<User> => {
  const response = await axios.post(`${process.env.REACT_APP_API_URL}/login`, credential);
  return response.data;
};

export const logout = async (): Promise<void> => {
  await axios.post(`${process.env.REACT_APP_API_URL}/logout`);
};

// CSRFトークン取得用の関数を追加
export const getCsrfToken = async (): Promise<string> => {
  const response = await axios.get(`${process.env.REACT_APP_API_URL}/csrf-token`);
  return response.data.csrf_token;
};

// getCurrentUser関数はバックエンドに対応するエンドポイントがないため、
// 必要に応じて追加するか、または削除を検討
export const getCurrentUser = async (): Promise<User> => {
  // バックエンドにエンドポイントがない場合は、
  // ログイン時に保存したユーザー情報をローカルストレージから取得するなどの対応が必要
  throw new Error('Not implemented: getCurrentUser endpoint is not available in backend');
};
