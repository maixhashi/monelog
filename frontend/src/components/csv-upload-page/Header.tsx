import React from 'react';
import { 
  Typography, 
  Box, 
  Button 
} from '@mui/material';
import {
  KeyboardArrowDown as KeyboardArrowDownIcon,
  KeyboardArrowUp as KeyboardArrowUpIcon
} from '@mui/icons-material';

interface HeaderProps {
  showInstructions: boolean;
  setShowInstructions: (show: boolean) => void;
}

export const Header: React.FC<HeaderProps> = ({ showInstructions, setShowInstructions }) => {
  return (
    <Box sx={{ mb: 4, textAlign: 'center' }}>
      <Typography variant="h3" component="h1" gutterBottom sx={{ fontWeight: 'bold' }}>
        カード明細CSV集計ツール
      </Typography>
      <Typography variant="subtitle1" color="text.secondary" sx={{ maxWidth: 700, mx: 'auto', mb: 2 }}>
        クレジットカードの明細CSVをアップロードして、分割払いの支払いスケジュールを自動計算します。
      </Typography>
      <Button 
        variant="text" 
        color="primary"
        onClick={() => setShowInstructions(!showInstructions)}
        endIcon={showInstructions ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
        size="small"
      >
        {showInstructions ? '使い方を隠す' : '使い方を見る'}
      </Button>
    </Box>
  );
};
