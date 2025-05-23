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

  // 年・月を受け取るように修正
  const previewCSVMutation = useMutation({
    mutationFn: ({ file, cardType, year, month }: { file: File, cardType: CardType, year: number, month: number }) => 
      previewCSV(file, cardType, year, month),
    onError: (err: any) => {
      console.error('CSVプレビューエラー:', err);
      throw err;
    },
  });

  // 年・月を受け取るように修正
  const saveCardStatementsMutation = useMutation({
    mutationFn: ({ cardStatements, cardType, year, month }: { cardStatements: CardStatementSummary[], cardType: CardType, year: number, month: number }) => {
      console.log('保存するデータ（変換前）:', JSON.stringify(cardStatements, null, 2));
      console.log('カード種類:', cardType);
      return saveCardStatements(cardStatements, cardType, year, month);
    },
    onSuccess: (data) => {
      console.log('保存成功:', JSON.stringify(data, null, 2));
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
