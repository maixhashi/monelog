import React from 'react';
import { Box, Button, Typography, Paper, CircularProgress, Grid } from '@mui/material';
import CloudUploadIcon from '@mui/icons-material/CloudUpload';
import { CardType } from './utils/csvProcessor/index';
import { CardTypeSelector } from './CardTypeSelector';

interface FileUploaderProps {
  csvFile: File | null;
  setCsvFile: (file: File | null) => void;
  isProcessing: boolean;
  processCSV: () => void;
  cardType: CardType;
  setCardType: (cardType: CardType) => void;
}

export const FileUploader: React.FC<FileUploaderProps> = ({
  csvFile,
  setCsvFile,
  isProcessing,
  processCSV,
  cardType,
  setCardType
}) => {
  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      setCsvFile(event.target.files[0]);
    }
  };

  return (
    <Paper elevation={3} sx={{ p: 3, mb: 4 }}>
      <Typography variant="h6" gutterBottom>
        カード明細CSVアップロード
      </Typography>
      
      <CardTypeSelector cardType={cardType} setCardType={setCardType} />
      
      <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center', mb: 2 }}>
        <input
          accept=".csv"
          id="csv-file-upload"
          type="file"
          onChange={handleFileChange}
          style={{ display: 'none' }}
        />
        <label htmlFor="csv-file-upload">
          <Button
            variant="contained"
            component="span"
            startIcon={<CloudUploadIcon />}
            sx={{ mb: 2 }}
          >
            CSVファイルを選択
          </Button>
        </label>
        
        {csvFile && (
          <Typography variant="body2" color="textSecondary">
            選択されたファイル: {csvFile.name}
          </Typography>
        )}
      </Box>
      
      <Grid container justifyContent="center">
        <Button
          variant="contained"
          color="primary"
          onClick={processCSV}
          disabled={!csvFile || isProcessing}
          sx={{ minWidth: 200 }}
        >
          {isProcessing ? (
            <>
              <CircularProgress size={24} sx={{ mr: 1, color: 'white' }} />
              処理中...
            </>
          ) : (
            'CSVを処理'
          )}
        </Button>
      </Grid>
    </Paper>
  );
};
