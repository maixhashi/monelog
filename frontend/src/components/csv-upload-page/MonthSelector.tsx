import React from 'react';
import { Box, IconButton, Typography } from '@mui/material';

type Props = {
  month: number;
  setMonth: (month: number) => void;
};

export const MonthSelector: React.FC<Props> = ({ month, setMonth }) => (
  <Box display="flex" justifyContent="center" mb={2} flexWrap="wrap" gap={1}>
    {[...Array(12)].map((_, i) => {
      const m = i + 1;
      const isSelected = month === m;
      return (
        <IconButton
          key={m}
          onClick={() => setMonth(m)}
          sx={{
            borderRadius: '50%',
            width: 48,
            height: 48,
            backgroundColor: isSelected ? 'primary.main' : 'grey.200',
            color: isSelected ? 'white' : 'text.primary',
            border: isSelected ? '2px solid' : '1px solid',
            borderColor: isSelected ? 'primary.dark' : 'grey.300',
            transition: 'background 0.2s, color 0.2s',
            '&:hover': {
              backgroundColor: isSelected ? 'primary.dark' : 'grey.300',
            },
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            p: 0,
          }}
        >
          <Box display="flex" alignItems="center" justifyContent="center">
            <Typography
              component="span"
              sx={{
                fontSize: 22,
                fontWeight: isSelected ? 'bold' : 'normal',
                lineHeight: 1,
              }}
            >
              {m}
            </Typography>
            <Typography
              component="span"
              sx={{
                fontSize: 13,
                lineHeight: 1,
                fontWeight: 'normal',
                letterSpacing: 0,
                ml: 0.2,
              }}
            >
              月
            </Typography>
          </Box>
        </IconButton>
      );
    })}
  </Box>
);