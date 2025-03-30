import React from 'react';
import {
  Paper,
  Typography,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Collapse,
  Alert
} from '@mui/material';
import {
  LooksOne as LooksOneIcon,
  LooksTwo as LooksTwoIcon,
  Looks3 as Looks3Icon,
  Looks4 as Looks4Icon,
  Warning as WarningIcon
} from '@mui/icons-material';

interface InstructionsProps {
  showInstructions: boolean;
}

export const Instructions: React.FC<InstructionsProps> = ({ showInstructions }) => {
  return (
    <Collapse in={showInstructions}>
      <Paper sx={{ p: 3, mb: 4, maxWidth: 800, mx: 'auto' }}>
        <Typography variant="h6" gutterBottom>使い方</Typography>
        <List>
          <ListItem>
            <ListItemIcon><LooksOneIcon color="primary" /></ListItemIcon>
            <ListItemText primary="クレジットカード会社のウェブサイトから明細データをCSV形式でダウンロードします" />
          </ListItem>
          <ListItem>
            <ListItemIcon><LooksTwoIcon color="primary" /></ListItemIcon>
            <ListItemText primary="下のアップロードエリアにCSVファイルをドラッグ＆ドロップするか、クリックして選択します" />
          </ListItem>
          <ListItem>
            <ListItemIcon><Looks3Icon color="primary" /></ListItemIcon>
            <ListItemText primary="「CSVを処理する」ボタンをクリックして、データを分析します" />
          </ListItem>
          <ListItem>
            <ListItemIcon><Looks4Icon color="primary" /></ListItemIcon>
            <ListItemText primary="分割払いの支払いスケジュールが表示されます" />
          </ListItem>
        </List>
        <Alert severity="info" icon={<WarningIcon />} sx={{ mt: 2 }}>
          <Typography variant="body2">
            <strong>対応フォーマット:</strong> 現在は楽天カードの明細CSVフォーマットに対応しています。
          </Typography>
        </Alert>
      </Paper>
    </Collapse>
  );
};
