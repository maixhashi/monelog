import React from 'react';
import { IconButton, Typography, Box } from '@mui/material';
import ArrowBackIosIcon from '@mui/icons-material/ArrowBackIos';
import ArrowForwardIosIcon from '@mui/icons-material/ArrowForwardIos';

type Props = {
  year: number;
  setYear: (year: number) => void;
};

export const YearPagination: React.FC<Props> = ({ year, setYear }) => (
  <Box display="flex" alignItems="center" justifyContent="center" mb={2}>
    <IconButton onClick={() => setYear(year - 1)}>
      <ArrowBackIosIcon />
    </IconButton>
    <Typography variant="h6" sx={{ mx: 2 }}>{year}年</Typography>
    <IconButton onClick={() => setYear(year + 1)}>
      <ArrowForwardIosIcon />
    </IconButton>
  </Box>
);