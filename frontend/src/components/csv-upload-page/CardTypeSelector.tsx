import React from 'react';
import { FormControl, InputLabel, Select, MenuItem, SelectChangeEvent } from '@mui/material';
import { CardType, cardTypeDisplayNames } from '../../types/cardType';

interface CardTypeSelectorProps {
  cardType: CardType;
  setCardType: (cardType: CardType) => void;
}

export const CardTypeSelector: React.FC<CardTypeSelectorProps> = ({ cardType, setCardType }) => {
  const handleChange = (event: SelectChangeEvent) => {
    setCardType(event.target.value as CardType);
  };

  return (
    <FormControl fullWidth sx={{ mb: 2 }}>
      <InputLabel id="card-type-select-label">カード種類</InputLabel>
      <Select
        labelId="card-type-select-label"
        id="card-type-select"
        value={cardType}
        label="カード種類"
        onChange={handleChange}
      >
        <MenuItem value="rakuten">{cardTypeDisplayNames.rakuten}</MenuItem>
        <MenuItem value="epos">{cardTypeDisplayNames.epos}</MenuItem>
        <MenuItem value="mufg">{cardTypeDisplayNames.mufg}</MenuItem>
      </Select>
    </FormControl>
  );
};
