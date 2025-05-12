import { useQuery } from '@tanstack/react-query';
import axios from 'axios';

export interface CsvHistoryResponse {
  id: number;
  file_name: string;
  card_type: string;
  created_at: string;
  updated_at: string;
  year: number;
  month: number;
}

export const useQueryCsvHistoriesByMonth = (year: number, month: number) => {
  return useQuery<CsvHistoryResponse[], Error>({
    queryKey: ['csvHistories', 'byMonth', year, month],
    queryFn: async () => {
      if (!year || !month) return [];
      
      const response = await axios.get(
        `${process.env.REACT_APP_API_URL}/csv-histories/by-month`, {
          params: { year, month }
        }
      );
      return response.data;
    },
    enabled: !!year && !!month && year > 0 && month >= 1 && month <= 12,
  });
};