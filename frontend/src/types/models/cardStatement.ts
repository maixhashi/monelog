import { definitions } from '../api/generated';

// バックエンドのレスポンス型
export type CardStatementResponse = definitions['model.CardStatementResponse'];

/**
 * 集計結果の型
 * バックエンドのレスポンスと互換性を持たせる
 */
export interface CardStatementSummary {
  type: string;
  statementNo: number;
  cardType: string;
  description: string;
  useDate: string;
  paymentDate: string;
  paymentMonth: string;
  amount: number;
  totalChargeAmount: number;
  chargeAmount: number;
  remainingBalance: number;
  paymentCount: number;
  installmentCount: number;
  annualRate: number;
  monthlyRate: number;
  id?: number;
  created_at?: string;
  updated_at?: string;
};

/**
 * 編集用のカード明細型
 */
export type EditedCardStatement = {
  id: number;
  csvContent: string;
};

/**
 * 編集用カード明細の初期値
 */
export const initialEditedCardStatement: EditedCardStatement = {
  id: 0,
  csvContent: ''
};

/**
 * カード明細状態の型定義（Zustandストア用）
 */
export interface CardStatementState {
  editedCardStatement: EditedCardStatement;
  cardStatementSummaries: CardStatementSummary[];
  selectedYear: number;
  selectedMonth: number;
  updateEditedCardStatement: (payload: EditedCardStatement) => void;
  setCardStatementSummaries: (summaries: CardStatementSummary[]) => void;
  resetEditedCardStatement: () => void;
  setSelectedYear: (year: number) => void;
  setSelectedMonth: (month: number) => void;
}
