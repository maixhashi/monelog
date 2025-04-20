import React, { useState } from 'react';
import { Button, Box, TextField, Dialog, DialogTitle, DialogContent, DialogActions, Typography, CircularProgress } from '@mui/material';
import { CardTypeSelector } from './CardTypeSelector';
import { CardType } from '../../types/cardType';
import { useMutateCsvHistories } from '../../hooks/mutateHooks/useMutateCsvHistories';

interface CSVHistorySaverProps {
  file: File | null;
}

export const CSVHistorySaver: React.FC<CSVHistorySaverProps> = ({ file }) => {
  const [open, setOpen] = useState(false);
  const [fileName, setFileName] = useState('');
  const [cardType, setCardType] = useState<CardType>('rakuten');
  const { saveCSVHistoryMutation } = useMutateCsvHistories();
  const isLoading = saveCSVHistoryMutation.isLoading;

  const handleOpen = () => {
    if (file) {
      setFileName(file.name); // デフォルトでファイル名をセット
      setOpen(true);
    }
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleSave = async () => {
    if (!file) return;

    try {
      await saveCSVHistoryMutation.mutateAsync({
        file,
        fileName,
        cardType
      });
      handleClose();
    } catch (error) {
      console.error('CSV履歴保存エラー:', error);
    }
  };

  return (
    <>
      <Button 
        variant="outlined" 
        color="primary" 
        onClick={handleOpen}
        disabled={!file}
        sx={{ mt: 2, mb: 2 }}
      >
        CSV履歴として保存
      </Button>

      <Dialog open={open} onClose={handleClose} maxWidth="sm" fullWidth>
        <DialogTitle>CSV履歴として保存</DialogTitle>
        <DialogContent>
          <Box sx={{ mt: 2 }}>
            <Typography variant="body2" color="text.secondary" gutterBottom>
              CSVファイルを履歴として保存します。後で再利用することができます。
            </Typography>
            
            <TextField
              fullWidth
              label="ファイル名"
              value={fileName}
              onChange={(e) => setFileName(e.target.value)}
              margin="normal"
              required
            />
            
            <CardTypeSelector 
              cardType={cardType}
              setCardType={setCardType}
            />
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} disabled={isLoading}>
            キャンセル
          </Button>
          <Button 
            onClick={handleSave} 
            variant="contained" 
            color="primary"
            disabled={!fileName || isLoading}
          >
            {isLoading ? <CircularProgress size={24} /> : '保存'}
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};
