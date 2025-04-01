import { StateCreator } from 'zustand';
import { CardStatementState, EditedCardStatement, CardStatementSummary, initialEditedCardStatement } from '../../types/models/cardStatement';

/**
 * カード明細関連のZustandストアスライス
 */
export const createCardStatementSlice: StateCreator<CardStatementState> = (set) => ({
  editedCardStatement: initialEditedCardStatement,
  cardStatementSummaries: [],
  updateEditedCardStatement: (payload: EditedCardStatement) => set({
    editedCardStatement: payload
  }),
  setCardStatementSummaries: (summaries: CardStatementSummary[]) => set({
    cardStatementSummaries: summaries
  }),
  resetEditedCardStatement: () => set({ 
    editedCardStatement: initialEditedCardStatement 
  }),
});
