import { StateCreator } from 'zustand';
import { CardStatementState, EditedCardStatement, CardStatementSummary, initialEditedCardStatement } from '../../types/models/cardStatement';

const currentYear = new Date().getFullYear();
const currentMonth = new Date().getMonth() + 1;

/**
 * カード明細関連のZustandストアスライス
 */
export const createCardStatementSlice: StateCreator<CardStatementState> = (set) => ({
  editedCardStatement: initialEditedCardStatement,
  cardStatementSummaries: [],
  selectedYear: currentYear,
  selectedMonth: currentMonth,
  updateEditedCardStatement: (payload: EditedCardStatement) => set({
    editedCardStatement: payload
  }),
  setCardStatementSummaries: (summaries: CardStatementSummary[]) => set({
    cardStatementSummaries: summaries
  }),
  resetEditedCardStatement: () => set({ 
    editedCardStatement: initialEditedCardStatement 
  }),
  setSelectedYear: (year: number) => set({ selectedYear: year }),
  setSelectedMonth: (month: number) => set({ selectedMonth: month }),
});
