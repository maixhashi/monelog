import { useMutation, useQueryClient } from '@tanstack/react-query';
import { uploadCSV } from '../../api/cardStatements';
import { CardType } from '../../types/cardType';

export const useMutateCardStatements = () => {
  const queryClient = useQueryClient();

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

  return { uploadCSVMutation };
};
