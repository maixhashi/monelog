import React, { useState, useEffect } from 'react';
import { 
  Box, 
  FormControl, 
  InputLabel, 
  Select, 
  MenuItem, 
  Button, 
  Grid,
  SelectChangeEvent
} from '@mui/material';
import { CardStatementSummary } from '../../types/models/cardStatement';

interface CardStatementFilterProps {
  onFilter: (month: string, cardType: string) => void;
  cardStatements: CardStatementSummary[];
}

export const CardStatementFilter: React.FC<CardStatementFilterProps> = ({ 
  onFilter,
  cardStatements
}) => {
  const [month, setMonth] = useState('');
  const [cardType, setCardType] = useState('');
  const [availableMonths, setAvailableMonths] = useState<string[]>([]);
  const [availableCardTypes, setAvailableCardTypes] = useState<string[]>([]);

  // 利用可能な支払月とカード種類を抽出
  useEffect(() => {
    if (cardStatements.length > 0) {
      // 支払月の一覧を取得（重複を除去）
      const months = Array.from(new Set(
        cardStatements.map(statement => statement.paymentMonth)
      )).sort().reverse(); // 新しい月順にソート
      
      // カード種類の一覧を取得（重複を除去）
      const cardTypes = Array.from(new Set(
        cardStatements.map(statement => statement.cardType)
      )).sort();
      
      setAvailableMonths(months);
      setAvailableCardTypes(cardTypes);
    }
  }, [cardStatements]);

  const handleMonthChange = (event: SelectChangeEvent) => {
    setMonth(event.target.value);
  };

  const handleCardTypeChange = (event: SelectChangeEvent) => {
    setCardType(event.target.value);
  };

  const handleFilter = () => {
    onFilter(month, cardType);
  };

  const handleClear = () => {
    setMonth('');
    setCardType('');
    onFilter('', '');
  };

  return (
    <Box sx={{ p: 2 }}>
      <Grid container spacing={2} alignItems="center">
        <Grid item xs={12} sm={4}>
          <FormControl fullWidth>
            <InputLabel id="month-select-label">支払月</InputLabel>
            <Select
              labelId="month-select-label"
              id="month-select"
              value={month}
              label="支払月"
              onChange={handleMonthChange}
            >
              <MenuItem value="">すべて</MenuItem>
              {availableMonths.map((m) => (
                <MenuItem key={m} value={m}>{m}</MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>
        
        <Grid item xs={12} sm={4}>
          <FormControl fullWidth>
            <InputLabel id="card-type-select-label">カード種類</InputLabel>
            <Select
              labelId="card-type-select-label"
              id="card-type-select"
              value={cardType}
              label="カード種類"
              onChange={handleCardTypeChange}
            >
              <MenuItem value="">すべて</MenuItem>
              {availableCardTypes.map((type) => (
                <MenuItem key={type} value={type}>{type}</MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>
        
        <Grid item xs={12} sm={4}>
          <Box sx={{ display: 'flex', gap: 2 }}>
            <Button 
              variant="contained" 
              color="primary" 
              onClick={handleFilter}
              fullWidth
            >
              フィルター適用
            </Button>
            <Button 
              variant="outlined" 
              onClick={handleClear}
              fullWidth
            >
              クリア
            </Button>
          </Box>
        </Grid>
      </Grid>
    </Box>
  );
};
