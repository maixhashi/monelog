import React from 'react';
import {
  Box,
  Paper,
  Typography,
  Button
} from '@mui/material';
import {
  Description as DescriptionIcon
} from '@mui/icons-material';

import { CardStatementSummary } from '../../types/models/cardStatement';

interface ResultsTableProps {
  cardStatementSummaries: CardStatementSummary[];
  clearResults: () => void;
  isPreviewData?: boolean; // 追加: プレビューデータかどうかを示すフラグ
}

export const ResultsTable: React.FC<ResultsTableProps> = ({ 
  cardStatementSummaries, 
  clearResults,
  isPreviewData = false // デフォルト値を設定
}) => {
  if (cardStatementSummaries.length === 0) {
    return null;
  }

  return (
    <Box sx={{ mt: 4, animation: 'fadeIn 0.5s ease-in-out' }}>
      <Paper elevation={2} sx={{ p: 3, mb: 2 }}>
        <Typography variant="h5" component="h2" gutterBottom>
          {isPreviewData ? 'プレビュー結果' : '集計結果'}
        </Typography>
        <Typography variant="body2" color="text.secondary">
          {cardStatementSummaries.length}件のレコードが{isPreviewData ? 'プレビュー' : '処理'}されました。
          {isPreviewData && '「データを保存する」ボタンをクリックすると、データベースに保存されます。'}
          分割払いの支払いスケジュールは以下の通りです。
        </Typography>
      </Paper>
      
      <Paper elevation={3} sx={{ overflow: 'auto' }}>
        <Box sx={{ minWidth: '100%', overflowX: 'auto' }}>
          <table style={{ width: '100%', borderCollapse: 'collapse' }}>
            <thead style={{ backgroundColor: '#f5f5f5' }}>
              <tr>
                <th style={{ padding: '12px 8px', textAlign: 'left', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>種別</th>
                <th style={{ padding: '12px 8px', textAlign: 'left', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>明細No</th>
                <th style={{ padding: '12px 8px', textAlign: 'left', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>カード種類</th>
                <th style={{ padding: '12px 8px', textAlign: 'left', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>利用店名・商品名</th>
                <th style={{ padding: '12px 8px', textAlign: 'left', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>利用日</th>
                <th style={{ padding: '12px 8px', textAlign: 'left', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>支払日</th>
                <th style={{ padding: '12px 8px', textAlign: 'left', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>支払月</th>
                <th style={{ padding: '12px 8px', textAlign: 'right', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>利用金額</th>
                <th style={{ padding: '12px 8px', textAlign: 'right', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>請求金額（分割手数料込）</th>
                <th style={{ padding: '12px 8px', textAlign: 'right', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>請求金額</th>
                <th style={{ padding: '12px 8px', textAlign: 'right', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>請求金額残高</th>
                <th style={{ padding: '12px 8px', textAlign: 'center', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>支払回数</th>
                <th style={{ padding: '12px 8px', textAlign: 'center', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>分割回数</th>
                <th style={{ padding: '12px 8px', textAlign: 'right', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>実質年率</th>
                <th style={{ padding: '12px 8px', textAlign: 'right', fontSize: '0.75rem', fontWeight: 500, color: '#666', textTransform: 'uppercase' }}>月利</th>
              </tr>
            </thead>
            <tbody>
              {cardStatementSummaries.map((summary, index) => (
                <tr 
                  key={index} 
                  style={{ 
                    backgroundColor: summary.type === '発生' ? '#EBF4FF' : 'white',
                    borderBottom: '1px solid #eee'
                  }}
                >
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem', fontWeight: summary.type === '発生' ? 500 : 400 }}>{summary.type}</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem' }}>{summary.statementNo}</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem' }}>{summary.cardType}</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem' }}>{summary.description}</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem' }}>{summary.useDate}</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem' }}>{summary.paymentDate}</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem' }}>{summary.paymentMonth}</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem', textAlign: 'right' }}>{summary.amount.toLocaleString()}</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem', textAlign: 'right' }}>{summary.totalChargeAmount.toLocaleString()}</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem', textAlign: 'right' }}>{summary.chargeAmount.toLocaleString()}</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem', textAlign: 'right' }}>{summary.remainingBalance.toLocaleString()}</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem', textAlign: 'center' }}>{summary.paymentCount}</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem', textAlign: 'center' }}>{summary.installmentCount}</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem', textAlign: 'right' }}>{(summary.annualRate * 100).toFixed(2)}%</td>
                  <td style={{ padding: '10px 8px', fontSize: '0.875rem', textAlign: 'right' }}>{(summary.monthlyRate * 100).toFixed(4)}%</td>
                </tr>
              ))}
            </tbody>
          </table>
        </Box>
      </Paper>
      
      <Box sx={{ mt: 3, display: 'flex', justifyContent: 'flex-end' }}>
        <Button 
          variant="outlined" 
          color="primary" 
          startIcon={<DescriptionIcon />}
          sx={{ mr: 2 }}
        >
          CSVダウンロード
        </Button>
        <Button 
          variant="outlined" 
          color="secondary" 
          onClick={clearResults}
        >
          クリア
        </Button>
      </Box>
    </Box>
  );
};
