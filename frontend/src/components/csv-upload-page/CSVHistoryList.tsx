import React, { useState } from 'react'
import { 
  Box, 
  Typography, 
  List, 
  ListItem, 
  ListItemText, 
  ListItemSecondaryAction, 
  IconButton, 
  Divider,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  CircularProgress,
  Collapse,
  Paper
} from '@mui/material'
import { Delete, Download, History, ExpandMore, ExpandLess } from '@mui/icons-material'
import { useQueryCsvHistories } from '../../hooks/queryHooks/useQueryCsvHistories'
import { useMutateCsvHistories } from '../../hooks/mutateHooks/useMutateCsvHistories'
import { downloadCSVHistory } from '../../api/csvHistories'
import { format } from 'date-fns'
import { ja } from 'date-fns/locale'
import { CSVHistoryResponse } from '../../types/models/csvHistory'
import { CardType } from '../../types/cardType'

export const CSVHistoryList: React.FC = () => {
  const { data: csvHistories, isLoading, isError } = useQueryCsvHistories()
  const { deleteCSVHistoryMutation } = useMutateCsvHistories()
  const [open, setOpen] = useState(false)
  const [selectedHistory, setSelectedHistory] = useState<CSVHistoryResponse | null>(null)
  const [isDownloading, setIsDownloading] = useState(false)
  const [showHistories, setShowHistories] = useState(false)

  const handleDeleteClick = (history: CSVHistoryResponse) => {
    setSelectedHistory(history)
    setOpen(true)
  }

  const handleDeleteConfirm = async () => {
    if (selectedHistory && selectedHistory.id !== undefined) {
      try {
        await deleteCSVHistoryMutation.mutateAsync(selectedHistory.id)
        setOpen(false)
        setSelectedHistory(null)
      } catch (error) {
        console.error('CSV履歴削除エラー:', error)
      }
    }
  }

  const handleDownload = async (history: CSVHistoryResponse) => {
    if (history.id !== undefined && history.file_name) {
      try {
        setIsDownloading(true)
        await downloadCSVHistory(history.id, history.file_name)
      } catch (error) {
        console.error('CSV履歴ダウンロードエラー:', error)
      } finally {
        setIsDownloading(false)
      }
    }
  }

  const formatDate = (dateString: string) => {
    try {
      const date = new Date(dateString)
      return format(date, 'yyyy年MM月dd日 HH:mm', { locale: ja })
    } catch (error) {
      return dateString
    }
  }

  const getCardTypeLabel = (cardType: string): string => {
    const cardTypes: Record<string, string> = {
      'rakuten': '楽天カード',
      'mufg': 'MUFGカード',
      'epos': 'エポスカード'
    }
    return cardTypes[cardType as CardType] || cardType
  }

  if (isLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', mt: 4 }}>
        <CircularProgress />
      </Box>
    )
  }

  if (isError) {
    return (
      <Box sx={{ mt: 4 }}>
        <Typography color="error">CSV履歴の取得中にエラーが発生しました。</Typography>
      </Box>
    )
  }

  return (
    <Paper sx={{ mt: 4, p: 2 }}>
      <Box 
        sx={{ 
          display: 'flex', 
          alignItems: 'center', 
          cursor: 'pointer',
          mb: 1
        }}
        onClick={() => setShowHistories(!showHistories)}
      >
        <History color="primary" sx={{ mr: 1 }} />
        <Typography variant="h6" component="div">
          CSV履歴
        </Typography>
        {showHistories ? <ExpandLess /> : <ExpandMore />}
      </Box>
      
      <Collapse in={showHistories}>
        {csvHistories && csvHistories.length > 0 ? (
          <List>
            {csvHistories.map((history) => (
              <React.Fragment key={history.id}>
                <ListItem>
                  <ListItemText
                    primary={history.file_name}
                    secondary={
                      <>
                        <Typography component="span" variant="body2" color="text.primary">
                          {getCardTypeLabel(history.card_type || '')}
                        </Typography>
                        {` - ${formatDate(history.created_at || '')}`}
                      </>
                    }
                  />
                  <ListItemSecondaryAction>
                    <IconButton 
                      edge="end" 
                      aria-label="download"
                      onClick={() => handleDownload(history)}
                      disabled={isDownloading}
                    >
                      <Download />
                    </IconButton>
                    <IconButton 
                      edge="end" 
                      aria-label="delete"
                      onClick={() => handleDeleteClick(history)}
                      disabled={deleteCSVHistoryMutation.isLoading}
                    >
                      <Delete />
                    </IconButton>
                  </ListItemSecondaryAction>
                </ListItem>
                <Divider />
              </React.Fragment>
            ))}
          </List>
        ) : (
          <Typography variant="body2" color="text.secondary" sx={{ p: 2 }}>
            保存されたCSV履歴はありません。
          </Typography>
        )}
      </Collapse>

      {/* 削除確認ダイアログ */}
      <Dialog
        open={open}
        onClose={() => setOpen(false)}
      >
        <DialogTitle>CSV履歴の削除</DialogTitle>
        <DialogContent>
          <Typography>
            「{selectedHistory?.file_name}」を削除してもよろしいですか？
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpen(false)} disabled={deleteCSVHistoryMutation.isLoading}>
            キャンセル
          </Button>
          <Button 
            onClick={handleDeleteConfirm} 
            color="error"
            disabled={deleteCSVHistoryMutation.isLoading}
          >
            {deleteCSVHistoryMutation.isLoading ? <CircularProgress size={24} /> : '削除'}
          </Button>
        </DialogActions>
      </Dialog>
    </Paper>
  )
}