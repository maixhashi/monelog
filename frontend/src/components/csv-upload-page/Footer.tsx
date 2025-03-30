import React from 'react';
import { Box, Typography } from '@mui/material';

export const Footer: React.FC = () => {
  return (
    <Box sx={{ mt: 8, textAlign: 'center', color: 'text.secondary', fontSize: '0.875rem' }}>
      <Typography variant="body2" color="text.secondary">
        © 2023 MoneyLog - クレジットカード明細分析ツール
      </Typography>
      <Typography variant="caption" display="block" sx={{ mt: 1 }}>
        このツールは個人情報を保存しません。すべての処理はブラウザ上で行われます。
      </Typography>
    </Box>
  );
};
