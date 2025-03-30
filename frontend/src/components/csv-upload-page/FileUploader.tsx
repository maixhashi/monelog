import React, { useState } from 'react';
import {
  Paper,
  Typography,
  Box,
  Button,
  Divider,
  CircularProgress
} from '@mui/material';
import {
  CloudUpload as CloudUploadIcon,
  Description as DescriptionIcon
} from '@mui/icons-material';
import { styled } from '@mui/material/styles';

// スタイル付きのコンポーネントを作成
const VisuallyHiddenInput = styled('input')({
  clip: 'rect(0 0 0 0)',
  clipPath: 'inset(50%)',
  height: 1,
  overflow: 'hidden',
  position: 'absolute',
  bottom: 0,
  left: 0,
  whiteSpace: 'nowrap',
  width: 1,
});

// カスタムプロパティの型を定義
interface DropZoneProps {
  isDragging?: boolean;
  hasFile?: boolean;
}

// スタイル付きのコンポーネントを作成（型を適切に指定）
const DropZone = styled(Paper, {
  shouldForwardProp: (prop) => prop !== 'isDragging' && prop !== 'hasFile'
})<DropZoneProps>(({ theme, isDragging, hasFile }) => ({
  padding: theme.spacing(3),
  textAlign: 'center',
  cursor: 'pointer',
  border: '2px dashed',
  borderColor: isDragging ? theme.palette.primary.main : 
              hasFile ? theme.palette.success.main : theme.palette.divider,
  backgroundColor: isDragging ? theme.palette.primary.light + '20' : 
                  hasFile ? theme.palette.success.light + '20' : theme.palette.background.default,
  transition: 'all 0.3s ease',
  '&:hover': {
    backgroundColor: isDragging ? theme.palette.primary.light + '30' : 
                    hasFile ? theme.palette.success.light + '30' : theme.palette.action.hover,
  },
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  justifyContent: 'center',
  minHeight: 200
}));

interface FileUploaderProps {
  csvFile: File | null;
  setCsvFile: (file: File | null) => void;
  isProcessing: boolean;
  processCSV: () => void;
}

export const FileUploader: React.FC<FileUploaderProps> = ({ 
  csvFile, 
  setCsvFile, 
  isProcessing, 
  processCSV 
}) => {
  const [isDragging, setIsDragging] = useState(false);
  const [fileName, setFileName] = useState<string>('');

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      const file = event.target.files[0];
      setFileName(file.name);
      setCsvFile(file);
    }
  };

  const handleDragOver = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    e.stopPropagation();
    setIsDragging(true);
  };

  const handleDragLeave = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    e.stopPropagation();
    setIsDragging(false);
  };

  const handleDrop = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    e.stopPropagation();
    setIsDragging(false);
    
    if (e.dataTransfer.files && e.dataTransfer.files.length > 0) {
      const file = e.dataTransfer.files[0];
      if (file.type === 'text/csv' || file.name.endsWith('.csv')) {
        setFileName(file.name);
        setCsvFile(file);
      } else {
        alert('CSVファイルのみアップロードできます。');
      }
    }
  };

  const clearFile = () => {
    setCsvFile(null);
    setFileName('');
  };

  return (
    <Paper elevation={2} sx={{ p: 3, mb: 4 }}>
      <Typography variant="h5" component="h2" gutterBottom sx={{ display: 'flex', alignItems: 'center' }}>
        <DescriptionIcon sx={{ mr: 1 }} color="primary" />
        CSVファイルアップロード
      </Typography>
      <Divider sx={{ my: 2 }} />
      
      <Box sx={{ my: 3 }}>
        <label htmlFor="csv-file-upload">
          <DropZone 
            isDragging={isDragging} 
            hasFile={!!csvFile}
            onDragOver={handleDragOver}
            onDragLeave={handleDragLeave}
            onDrop={handleDrop}
          >
            <CloudUploadIcon sx={{ fontSize: 48, color: csvFile ? 'success.main' : 'primary.main', mb: 2 }} />
            <Typography variant="h6" gutterBottom>
              {csvFile ? 'ファイルを選択しました' : 'CSVファイルをアップロード'}
            </Typography>
            {csvFile ? (
              <Box sx={{ 
                display: 'flex', 
                alignItems: 'center', 
                bgcolor: 'success.light', 
                color: 'success.contrastText',
                px: 2,
                py: 0.5,
                borderRadius: 16
              }}>
                <DescriptionIcon sx={{ mr: 1, fontSize: 16 }} />
                <Typography variant="body2">{fileName}</Typography>
              </Box>
            ) : (
              <Typography variant="body2" color="text.secondary">
                ファイルをドラッグ＆ドロップするか、クリックして選択してください
              </Typography>
            )}
          </DropZone>
        </label>
        <VisuallyHiddenInput
          id="csv-file-upload"
          type="file"
          accept=".csv"
          onChange={handleFileChange}
        />
        
        {csvFile && (
          <Box sx={{ display: 'flex', justifyContent: 'flex-end', mt: 1 }}>
            <Button 
              size="small" 
              color="error" 
              onClick={clearFile}
            >
              ファイルを削除
            </Button>
          </Box>
        )}
      </Box>
      
      <Button
        variant="contained"
        color="primary"
        fullWidth
        size="large"
        startIcon={isProcessing ? <CircularProgress size={20} color="inherit" /> : <DescriptionIcon />}
        onClick={processCSV}
        disabled={!csvFile || isProcessing}
        sx={{ mt: 2 }}
      >
        {isProcessing ? 'CSVを処理中...' : 'CSVを処理する'}
      </Button>
    </Paper>
  );
};
