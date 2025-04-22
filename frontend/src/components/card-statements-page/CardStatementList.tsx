import React, { useState } from 'react';
import { 
  Table, 
  TableBody, 
  TableCell, 
  TableContainer, 
  TableHead, 
  TableRow, 
  Paper, 
  TablePagination,
  Typography,
  Box,
  Chip,
  Tooltip
} from '@mui/material';
import { CardStatementSummary } from '../../types/models/cardStatement';

interface CardStatementListProps {
  cardStatements: CardStatementSummary[];
  filterMonth: string;
  filterCardType: string;
}

export const CardStatementList: React.FC<CardStatementListProps> = ({ 
  cardStatements,
  filterMonth,
  filterCardType
}) => {
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);

  // ページネーション処理
  const handleChangePage = (event: unknown, newPage: number) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  // 合計金額の計算
  const totalAmount = cardStatements.reduce((sum, statement) => sum + statement.amount, 0);

  // 日付のフォーマット関数
  const formatDate = (dateString: string) => {
    if (!dateString) return '';
    try {
      const date = new Date(dateString);
      return date.toLocaleString('ja-JP');
    } catch (e) {
      return dateString;
    }
  };

  return (
    <Box>
      <Box sx={{ mb: 2, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <Typography variant="h6">
          {filterMonth && <Chip label={`支払月: ${filterMonth}`} sx={{ mr: 1 }} />}
          {filterCardType && <Chip label={`カード: ${filterCardType}`} />}
        </Typography>
        <Typography variant="h6">
          合計金額: {totalAmount.toLocaleString()}円
        </Typography>
      </Box>
      
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="カード明細テーブル" size="small">
          <TableHead>
            <TableRow>
              <TableCell>種別</TableCell>
              <TableCell>明細No</TableCell>
              <TableCell>カード種類</TableCell>
              <TableCell>説明</TableCell>
              <TableCell>利用日</TableCell>
              <TableCell>支払日</TableCell>
              <TableCell>支払月</TableCell>
              <TableCell align="right">金額</TableCell>
              <TableCell align="right">総請求額</TableCell>
              <TableCell align="right">手数料</TableCell>
              <TableCell align="right">残高</TableCell>
              <TableCell align="right">支払回数</TableCell>
              <TableCell align="right">分割回数</TableCell>
              <TableCell align="right">年率(%)</TableCell>
              <TableCell align="right">月率(%)</TableCell>
              <TableCell>登録日時</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {cardStatements
              .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
              .map((statement, index) => (
                <TableRow
                  key={statement.id || index}
                  sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                  <TableCell>{statement.type}</TableCell>
                  <TableCell>{statement.statementNo}</TableCell>
                  <TableCell>{statement.cardType}</TableCell>
                  <Tooltip title={statement.description} arrow>
                    <TableCell sx={{ maxWidth: 150, overflow: 'hidden', textOverflow: 'ellipsis', whiteSpace: 'nowrap' }}>
                      {statement.description}
                    </TableCell>
                  </Tooltip>
                  <TableCell>{statement.useDate}</TableCell>
                  <TableCell>{statement.paymentDate}</TableCell>
                  <TableCell>{statement.paymentMonth}</TableCell>
                  <TableCell align="right">{statement.amount.toLocaleString()}</TableCell>
                  <TableCell align="right">{statement.totalChargeAmount.toLocaleString()}</TableCell>
                  <TableCell align="right">{statement.chargeAmount.toLocaleString()}</TableCell>
                  <TableCell align="right">{statement.remainingBalance.toLocaleString()}</TableCell>
                  <TableCell align="right">{statement.paymentCount}</TableCell>
                  <TableCell align="right">{statement.installmentCount}</TableCell>
                  <TableCell align="right">{(statement.annualRate * 100).toFixed(2)}</TableCell>
                  <TableCell align="right">{(statement.monthlyRate * 100).toFixed(4)}</TableCell>
                  <TableCell>{formatDate(statement.created_at || '')}</TableCell>
                </TableRow>
              ))}
          </TableBody>
        </Table>
      </TableContainer>
      
      <TablePagination
        rowsPerPageOptions={[5, 10, 25, 50]}
        component="div"
        count={cardStatements.length}
        rowsPerPage={rowsPerPage}
        page={page}
        onPageChange={handleChangePage}
        onRowsPerPageChange={handleChangeRowsPerPage}
        labelRowsPerPage="表示件数:"
        labelDisplayedRows={({ from, to, count }) => `${from}-${to} / ${count}`}
      />
    </Box>
  );
};
