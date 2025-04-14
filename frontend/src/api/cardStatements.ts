import axios from 'axios';
import { CardStatementSummary } from '../types/models/cardStatement';
import { CardType } from '../types/cardType';

// カード明細一覧を取得
export const fetchCardStatements = async (): Promise<CardStatementSummary[]> => {
  const response = await axios.get(`${process.env.REACT_APP_API_URL}/card-statements`);
  return mapResponseToSummaries(response.data);
};

// 特定のカード明細を取得
export const fetchCardStatementById = async (id: number): Promise<CardStatementSummary> => {
  const response = await axios.get(`${process.env.REACT_APP_API_URL}/card-statements/${id}`);
  return mapResponseToSummary(response.data);
};

// CSVファイルをアップロード
export const uploadCSV = async (file: File, cardType: CardType): Promise<CardStatementSummary[]> => {
  const formData = new FormData();
  formData.append('file', file);
  formData.append('card_type', cardType);
  
  const response = await axios.post(
    `${process.env.REACT_APP_API_URL}/card-statements/upload`, 
    formData, 
    {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    }
  );
  
  return mapResponseToSummaries(response.data);
};

// CSVファイルをプレビュー（保存なし）
export const previewCSV = async (file: File, cardType: CardType): Promise<CardStatementSummary[]> => {
    const formData = new FormData();
    formData.append('file', file);
    formData.append('card_type', cardType);
    
    const response = await axios.post(
      `${process.env.REACT_APP_API_URL}/card-statements/preview`, 
      formData, 
      {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      }
    );
    
    return mapResponseToSummaries(response.data);
  };
  
  // プレビューしたデータを保存
  export const saveCardStatements = async (
    cardStatements: CardStatementSummary[], 
    cardType: CardType
  ): Promise<CardStatementSummary[]> => {
    const response = await axios.post(
      `${process.env.REACT_APP_API_URL}/card-statements/save`,
      {
        card_statements: cardStatements,
        card_type: cardType
      }
    );
    
    return mapResponseToSummaries(response.data);
  };
  
// バックエンドのレスポンスをフロントエンドの型に変換するヘルパー関数
function mapResponseToSummary(item: any): CardStatementSummary {
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
    updated_at: item.updated_at
  };
}

function mapResponseToSummaries(items: any[]): CardStatementSummary[] {
  return items.map(mapResponseToSummary);
}
