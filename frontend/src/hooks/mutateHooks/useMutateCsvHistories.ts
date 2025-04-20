import { useMutation, useQueryClient } from '@tanstack/react-query';
import { saveCSVHistory, deleteCSVHistory } from '../../api/csvHistories';
import { CardType } from '../../types/cardType';

export const useMutateCsvHistories = () => {
  const queryClient = useQueryClient();

  // CSV履歴保存ミューテーション
  const saveCSVHistoryMutation = useMutation({
    mutationFn: ({ file, fileName, cardType }: { file: File, fileName: string, cardType: CardType }) => 
      saveCSVHistory(file, fileName, cardType),
    onSuccess: () => {
      // 成功時にCSV履歴一覧を再取得
      queryClient.invalidateQueries({ queryKey: ['csvHistories'] });
    },
    onError: (err: any) => {
      console.error('CSV履歴保存エラー:', err);
      throw err;
    },
  });

  // CSV履歴削除ミューテーション
  const deleteCSVHistoryMutation = useMutation({
    mutationFn: (id: number) => deleteCSVHistory(id),
    onSuccess: () => {
      // 成功時にCSV履歴一覧を再取得
      queryClient.invalidateQueries({ queryKey: ['csvHistories'] });
    },
    onError: (err: any) => {
      console.error('CSV履歴削除エラー:', err);
      throw err;
    },
  });

  return { 
    saveCSVHistoryMutation,
    deleteCSVHistoryMutation
  };
};
