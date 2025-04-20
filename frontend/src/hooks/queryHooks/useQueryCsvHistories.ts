import { useQuery } from '@tanstack/react-query';
import { fetchCSVHistories, fetchCSVHistoryById } from '../../api/csvHistories';
import { CSVHistoryResponse, CSVHistoryDetailResponse } from '../../types/models/csvHistory';

// CSV履歴一覧を取得するフック
export const useQueryCsvHistories = () => {
  return useQuery<CSVHistoryResponse[], Error>({
    queryKey: ['csvHistories'],
    queryFn: fetchCSVHistories,
  });
};

// 特定のCSV履歴を取得するフック
export const useQueryCsvHistoryById = (id: number) => {
  return useQuery<CSVHistoryDetailResponse, Error>({
    queryKey: ['csvHistory', id],
    queryFn: () => fetchCSVHistoryById(id),
    enabled: !!id, // idがある場合のみクエリを実行
  });
};
