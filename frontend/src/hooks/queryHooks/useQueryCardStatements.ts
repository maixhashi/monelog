import { useQuery } from '@tanstack/react-query';
import { fetchCardStatements } from '../../api/cardStatements';
import { CardStatementSummary } from '../../types/models/cardStatement';

export const useQueryCardStatements = () => {
  return useQuery<CardStatementSummary[], Error>({
    queryKey: ['cardStatements'],
    queryFn: fetchCardStatements,
    staleTime: Infinity,
  });
};
