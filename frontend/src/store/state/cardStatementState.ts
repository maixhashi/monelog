import { EditedCardStatement, CardStatementSummary } from '../../types/models/cardStatement';

// CardStatementState の型定義
export type CardStatementState = {
  editedCardStatement: EditedCardStatement;
  cardStatementSummaries: CardStatementSummary[];
  updateEditedCardStatement: (payload: EditedCardStatement) => void;
  setCardStatementSummaries: (summaries: CardStatementSummary[]) => void;
  resetEditedCardStatement: () => void;
}

// CardStatementState の初期状態
export const initialCardStatementState: CardStatementState = {
  editedCardStatement: { id: 0, csvContent: '' },
  cardStatementSummaries: [],
  updateEditedCardStatement: () => {},
  setCardStatementSummaries: () => {},
  resetEditedCardStatement: () => {},
}
