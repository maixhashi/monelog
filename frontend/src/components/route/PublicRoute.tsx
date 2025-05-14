import React from 'react';
import { Navigate, Outlet } from 'react-router-dom';
import { useAuth } from '../../hooks/useAuth';
import { CircularProgress, Box } from '@mui/material';

// 未認証ユーザーのみアクセス可能なルート（ログイン・サインアップなど）
export const PublicRoute: React.FC = () => {
  const { isAuthenticated, isLoading } = useAuth();

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <CircularProgress />
      </Box>
    );
  }

  // 認証済みの場合はダッシュボードにリダイレクト
  if (isAuthenticated) {
    return <Navigate to="/task-manager" replace />;
  }

  return <Outlet />;
};