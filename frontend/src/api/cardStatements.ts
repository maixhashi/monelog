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
    // 必要なフィールドが全て含まれていることを確認
    const payload = {
      card_statements: cardStatements.map(statement => ({
        // フロントエンドの型（camelCase）からバックエンドの型（snake_case）へ変換
        type: statement.type,
        statement_no: statement.statementNo,
        card_type: statement.cardType || cardType, // cardTypeを使用
        description: statement.description,
        use_date: statement.useDate,
        payment_date: statement.paymentDate,
        payment_month: statement.paymentMonth,
        amount: statement.amount,
        total_charge_amount: statement.totalChargeAmount,
        charge_amount: statement.chargeAmount,
        remaining_balance: statement.remainingBalance,
        payment_count: statement.paymentCount,
        installment_count: statement.installmentCount,
        annual_rate: statement.annualRate,
        monthly_rate: statement.monthlyRate
      })),
      card_type: cardType
    };
    
    console.log('APIに送信するペイロード:', JSON.stringify(payload, null, 2));
    
    const response = await axios.post(
      `${process.env.REACT_APP_API_URL}/card-statements/save`,
      payload,
      {
        headers: {
          'Content-Type': 'application/json',
        },
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
