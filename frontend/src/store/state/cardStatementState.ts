import { EditedCardStatement, CardStatementSummary } from '../../types/models/cardStatement';

// CardStatementState の型定義
export type CardStatementState = {
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

// CardStatementState の初期状態
const currentYear = new Date().getFullYear();
const currentMonth = new Date().getMonth() + 1;

export const initialCardStatementState: CardStatementState = {
  editedCardStatement: { id: 0, csvContent: '' },
  cardStatementSummaries: [],
  selectedYear: currentYear,
  selectedMonth: currentMonth,
  updateEditedCardStatement: () => {},
  setCardStatementSummaries: () => {},
  resetEditedCardStatement: () => {},
  setSelectedYear: () => {},
  setSelectedMonth: () => {},
}
