import React, { useState, useEffect } from 'react';
import { 
  Paper, 
  Table, 
  TableBody, 
  TableCell, 
  TableContainer, 
  TableHead, 
  TableRow,
  Typography,
  Box,
  Collapse
} from '@mui/material';

interface CsvPreviewProps {
  file: File | null;
  maxRows?: number;
}

export const CsvPreview: React.FC<CsvPreviewProps> = ({ file, maxRows = 5 }) => {
  const [previewData, setPreviewData] = useState<string[][]>([]);
  const [error, setError] = useState<string | null>(null);
  const [expanded, setExpanded] = useState(true);

  useEffect(() => {
    if (!file) {
      setPreviewData([]);
      setError(null);
      return;
    }

    const reader = new FileReader();
    
    reader.onload = (e) => {
      try {
        const text = e.target?.result as string;
        const lines = text.split('\n');
        const data = lines
          .filter(line => line.trim() !== '')
          .slice(0, maxRows)
          .map(line => line.split(','));
        
        setPreviewData(data);
        setError(null);
      } catch (err) {
        setError('CSVファイルの解析中にエラーが発生しました');
        setPreviewData([]);
      }
    };

    reader.onerror = () => {
      setError('ファイルの読み込み中にエラーが発生しました');
      setPreviewData([]);
    };

    reader.readAsText(file);
  }, [file, maxRows]);

  if (!file) return null;

  return (
    <Paper 
      elevation={2} 
      sx={{ 
        mt: 3, 
        mb: 3, 
        p: 2,
        cursor: 'pointer'
      }}
      onClick={() => setExpanded(!expanded)}
    >
      <Typography variant="h6" gutterBottom>
        CSVプレビュー {expanded ? '▼' : '▶'} 
        <Typography variant="caption" sx={{ ml: 1 }}>
          (最初の{maxRows}行のみ表示 - クリックで{expanded ? '折りたたむ' : '展開'})
        </Typography>
      </Typography>

      <Collapse in={expanded}>
        {error ? (
          <Typography color="error">{error}</Typography>
        ) : previewData.length > 0 ? (
          <TableContainer component={Box} sx={{ maxHeight: 300 }}>
            <Table size="small" stickyHeader>
              <TableHead>
                <TableRow>
                  {previewData[0].map((cell, index) => (
                    <TableCell key={index} sx={{ fontWeight: 'bold' }}>
                      {cell}
                    </TableCell>
                  ))}
                </TableRow>
              </TableHead>
              <TableBody>
                {previewData.slice(1).map((row, rowIndex) => (
                  <TableRow key={rowIndex}>
                    {row.map((cell, cellIndex) => (
                      <TableCell key={cellIndex}>{cell}</TableCell>
                    ))}
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        ) : (
          <Typography>プレビューするデータがありません</Typography>
        )}
      </Collapse>
    </Paper>
  );
};
