import { useState, useEffect } from 'react';
import { Container, Typography, Box, Paper, CircularProgress, Alert } from '@mui/material';
import { useQueryCardStatements } from '../hooks/queryHooks/useQueryCardStatements';
import { CardStatementList } from '../components/card-statements-page/CardStatementList';
import { CardStatementFilter } from '../components/card-statements-page/CardStatementFilter';
import { CardStatementSummary } from '../types/models/cardStatement';

export const CardStatementsPage = () => {
  const { data: cardStatements, isLoading, error } = useQueryCardStatements();
  const [filteredStatements, setFilteredStatements] = useState<CardStatementSummary[]>([]);
  const [filterMonth, setFilterMonth] = useState<string>('');
  const [filterCardType, setFilterCardType] = useState<string>('');

  // データが取得できたらフィルタリング用のステートを更新
  useEffect(() => {
    if (cardStatements) {
      setFilteredStatements(cardStatements);
    }
  }, [cardStatements]);

  // フィルタリング処理
  const handleFilter = (month: string, cardType: string) => {
    setFilterMonth(month);
    setFilterCardType(cardType);
    
    if (!cardStatements) return;
    
    let filtered = [...cardStatements];
    
    if (month) {
      filtered = filtered.filter(statement => statement.paymentMonth.includes(month));
    }
    
    if (cardType) {
      filtered = filtered.filter(statement => statement.cardType === cardType);
    }
    
    setFilteredStatements(filtered);
  };

  return (
    <Container maxWidth="lg" sx={{ py: 4 }}>
      <Typography variant="h4" component="h1" gutterBottom>
        カード明細一覧
      </Typography>
      
      <Paper sx={{ p: 2, mb: 3 }}>
        <CardStatementFilter 
          onFilter={handleFilter}
          cardStatements={cardStatements || []}
        />
      </Paper>
      
      <Box sx={{ mt: 3 }}>
        {isLoading ? (
          <Box sx={{ display: 'flex', justifyContent: 'center', p: 3 }}>
            <CircularProgress />
          </Box>
        ) : error ? (
          <Alert severity="error">
            データの取得中にエラーが発生しました: {error.message}
          </Alert>
        ) : filteredStatements.length === 0 ? (
          <Alert severity="info">
            表示するカード明細データがありません。CSVアップロードページからデータを登録してください。
          </Alert>
        ) : (
          <CardStatementList 
            cardStatements={filteredStatements} 
            filterMonth={filterMonth}
            filterCardType={filterCardType}
          />
        )}
      </Box>
    </Container>
  );
};
