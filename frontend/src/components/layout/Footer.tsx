import React from 'react';
import { Box, Container, Typography, Link as MuiLink } from '@mui/material';
import { Link } from 'react-router-dom';
import { useAuth } from '../../hooks/useAuth';

export const Footer: React.FC = () => {
  const { isAuthenticated } = useAuth();
  const currentYear = new Date().getFullYear();

  return (
    <Box
      component="footer"
      sx={{
        py: 3,
        px: 2,
        mt: 'auto',
        backgroundColor: (theme) => theme.palette.grey[200],
      }}
    >
      <Container maxWidth="lg">
        <Box sx={{ display: 'flex', justifyContent: 'space-between', flexWrap: 'wrap' }}>
          <Box>
            <Typography variant="h6" color="text.primary" gutterBottom>
              Monelog
            </Typography>
            <Typography variant="body2" color="text.secondary">
              家計簿管理アプリケーション
            </Typography>
          </Box>
          
          <Box>
            <Typography variant="h6" color="text.primary" gutterBottom>
              リンク
            </Typography>
            <Box sx={{ display: 'flex', flexDirection: 'column' }}>
              <MuiLink component={Link} to="/" color="inherit" underline="hover">
                ホーム
              </MuiLink>
              
              {isAuthenticated ? (
                <>
                  <MuiLink component={Link} to="/task-manager" color="inherit" underline="hover">
                    タスク管理
                  </MuiLink>
                  <MuiLink component={Link} to="/csv-upload" color="inherit" underline="hover">
                    CSV管理
                  </MuiLink>
                </>
              ) : (
                <>
                  <MuiLink component={Link} to="/login" color="inherit" underline="hover">
                    ログイン
                  </MuiLink>
                  <MuiLink component={Link} to="/signup" color="inherit" underline="hover">
                    新規登録
                  </MuiLink>
                </>
              )}
            </Box>
          </Box>
        </Box>
        
        <Box mt={3}>
          <Typography variant="body2" color="text.secondary" align="center">
            {'Copyright © '}
            <MuiLink component={Link} to="/" color="inherit">
              Monelog
            </MuiLink>{' '}
            {currentYear}
          </Typography>
        </Box>
      </Container>
    </Box>
  );
};