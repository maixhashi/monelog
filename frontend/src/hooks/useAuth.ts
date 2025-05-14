import { useEffect } from 'react';
import { useQuery } from '@tanstack/react-query';
import { verifyAuth } from '../api/auth';
import useStore from '../store';
import { User } from '../types/models/user';

export const useAuth = () => {
  const { isAuthenticated, user, isLoading, setAuth, clearAuth } = useStore();

  // 認証状態を確認するクエリ
  const { data, isLoading: isAuthLoading, refetch } = useQuery({
    queryKey: ['auth-verify'],
    queryFn: verifyAuth,
    retry: false,
    refetchOnWindowFocus: false,
  });

  // 認証状態が変わったらストアを更新
  useEffect(() => {
    if (!isAuthLoading && data) {
      if (data.authenticated && data.user) {
        // バックエンドのユーザー型をフロントエンドのユーザー型に変換
        const frontendUser: User = {
          id: data.user.id || 0, // id が undefined の場合は 0 を設定
          email: data.user.email
        };
        setAuth(true, frontendUser);
      } else {
        clearAuth();
      }
    }
  }, [data, isAuthLoading, setAuth, clearAuth]);

  return {
    isAuthenticated,
    user,
    isLoading: isLoading || isAuthLoading,
    refetchAuth: refetch,
  };
};
