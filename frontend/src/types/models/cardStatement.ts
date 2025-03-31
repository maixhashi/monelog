import { definitions } from '../api/generated';

/**
 * カード明細の入力データ型
 */
export type CardStatementInput = {
  useDate: string;
  storeName: string;
  user: string;
  paymentMethod: string;
  amount: number;
  fee: number;
  totalAmount: number;
  paymentMonth: string;
  currentMonthPayment: number;
  nextMonthBalance: number;
  futurePayments: number;
};

/**
 * 集計結果の型
 */
export interface CardStatementSummary {
  type: string;
  statementNo: number; // number型であることを確認
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
  updateEditedCardStatement: (payload: EditedCardStatement) => void;
  setCardStatementSummaries: (summaries: CardStatementSummary[]) => void;
  resetEditedCardStatement: () => void;
}
