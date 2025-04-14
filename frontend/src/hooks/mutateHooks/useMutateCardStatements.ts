import { useMutation, useQueryClient } from '@tanstack/react-query';
import { uploadCSV, previewCSV, saveCardStatements } from '../../api/cardStatements';
import { CardType } from '../../types/cardType';
import { CardStatementSummary } from '../../types/models/cardStatement';

export const useMutateCardStatements = () => {
  const queryClient = useQueryClient();

  // 既存のミューテーション（後方互換性のために残す）
  const uploadCSVMutation = useMutation({
    mutationFn: ({ file, cardType }: { file: File, cardType: CardType }) => 
      uploadCSV(file, cardType),
    onSuccess: (data) => {
      queryClient.setQueryData(['cardStatements'], data);
    },
    onError: (err: any) => {
      console.error('CSV処理エラー:', err);
      throw err;
    },
  });

  // 新しいプレビューミューテーション
  const previewCSVMutation = useMutation({
    mutationFn: ({ file, cardType }: { file: File, cardType: CardType }) => 
      previewCSV(file, cardType),
    onError: (err: any) => {
      console.error('CSVプレビューエラー:', err);
      throw err;
    },
  });

  // 保存ミューテーション
  const saveCardStatementsMutation = useMutation({
    mutationFn: ({ cardStatements, cardType }: { cardStatements: CardStatementSummary[], cardType: CardType }) => 
      saveCardStatements(cardStatements, cardType),
    onSuccess: (data) => {
      queryClient.setQueryData(['cardStatements'], data);
    },
    onError: (err: any) => {
      console.error('カード明細保存エラー:', err);
      throw err;
    },
  });

  return { 
    uploadCSVMutation, 
    previewCSVMutation,
    saveCardStatementsMutation
  };
};
