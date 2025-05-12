import React from 'react'
import { Box, Typography, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, IconButton } from '@mui/material'
import DownloadIcon from '@mui/icons-material/Download'
import DeleteIcon from '@mui/icons-material/Delete'
import { CsvHistoryResponse } from '../../hooks/queryHooks/useQueryCsvHistoriesByMonth'
import { useMutateCsvHistories } from '../../hooks/mutateHooks/useMutateCsvHistories'
import { format } from 'date-fns'
import { cardTypeDisplayNames, CardType } from '../../types/cardType'

type Props = {
  histories?: CsvHistoryResponse[]
  isLoading?: boolean
}

export const CSVHistoryList: React.FC<Props> = ({ histories = [], isLoading = false }) => {
  const { deleteCSVHistoryMutation } = useMutateCsvHistories()

  const handleDownload = (id: number) => {
    window.open(`${process.env.REACT_APP_API_URL}/csv-histories/${id}/download`, '_blank')
  }

  const handleDelete = async (id: number) => {
    if (window.confirm('このCSV履歴を削除しますか？')) {
      try {
        await deleteCSVHistoryMutation.mutateAsync(id)
      } catch (error) {
        console.error('CSV履歴削除エラー:', error)
      }
    }
  }

  if (isLoading) {
    return <Typography>読み込み中...</Typography>
  }

  if (histories.length === 0) {
    return <Typography sx={{ mt: 2 }}>この月のCSV履歴はありません</Typography>
  }

  return (
    <Box sx={{ mt: 4 }}>
      <Typography variant="h6" gutterBottom>
        CSV履歴一覧
      </Typography>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ファイル名</TableCell>
              <TableCell>カード種類</TableCell>
              <TableCell>アップロード日時</TableCell>
              <TableCell align="right">操作</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {histories.map((history) => (
              <TableRow key={history.id}>
                <TableCell>{history.file_name}</TableCell>
                <TableCell>
                  {cardTypeDisplayNames[history.card_type as CardType] || history.card_type}
                </TableCell>
                <TableCell>{format(new Date(history.created_at), 'yyyy/MM/dd HH:mm')}</TableCell>
                <TableCell align="right">
                  <IconButton onClick={() => handleDownload(history.id)} title="ダウンロード">
                    <DownloadIcon />
                  </IconButton>
                  <IconButton onClick={() => handleDelete(history.id)} title="削除">
                    <DeleteIcon />
                  </IconButton>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  )
}