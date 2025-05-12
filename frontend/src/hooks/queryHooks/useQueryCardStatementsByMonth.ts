import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { CardStatementSummary } from '../../types/models/cardStatement';

// バックエンドのレスポンスをフロントエンドの型に変換するヘルパー関数
function mapResponseToSummary(item: any): CardStatementSummary & { year: number, month: number } {
  return {
    type: item.type || '',
    statementNo: item.statement_no || 0,
    cardType: item.card_type || '',
    description: item.description || '',
    useDate: item.use_date || '',
    paymentDate: item.payment_date || '',
    paymentMonth: item.payment_month || '',
    amount: item.amount || 0,
    totalChargeAmount: item.total_charge_amount || 0,
    chargeAmount: item.charge_amount || 0,
    remainingBalance: item.remaining_balance || 0,
    paymentCount: item.payment_count || 0,
    installmentCount: item.installment_count || 0,
    annualRate: item.annual_rate || 0,
    monthlyRate: item.monthly_rate || 0,
    id: item.id,
    created_at: item.created_at,
    updated_at: item.updated_at,
    // year/monthはCardStatementSummary型に追加する必要があります
    year: item.year,
    month: item.month
  };
}

function mapResponseToSummaries(items: any[]): (CardStatementSummary & { year: number, month: number })[] {
  return items.map(mapResponseToSummary);
}

export const useQueryCardStatementsByMonth = (year: number, month: number) => {
  return useQuery<(CardStatementSummary & { year: number, month: number })[], Error>({
    queryKey: ['cardStatements', 'byMonth', year, month],
    queryFn: async () => {
      if (!year || !month) return [];
      
      const response = await axios.get(
        `${process.env.REACT_APP_API_URL}/card-statements/by-month`, {
          params: { year, month }
        }
      );
      return mapResponseToSummaries(response.data);
    },
    enabled: !!year && !!month && year > 0 && month >= 1 && month <= 12,
  });
};