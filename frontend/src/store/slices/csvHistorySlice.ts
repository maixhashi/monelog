import { StateCreator } from 'zustand';
import { CSVHistoryState } from '../state/csvHistoryState';
import { EditedCSVHistory, CSVHistoryResponse } from '../../types/models/csvHistory';

/**
 * CSV履歴関連のZustandストアスライス
 */
export const createCSVHistorySlice: StateCreator<CSVHistoryState> = (set) => ({
  editedCSVHistory: { id: 0, fileName: '', cardType: 'rakuten' },
  csvHistories: [],
  updateEditedCSVHistory: (payload: EditedCSVHistory) => set({
    editedCSVHistory: payload
  }),
  setCSVHistories: (histories: CSVHistoryResponse[]) => set({
    csvHistories: histories
  }),
  resetEditedCSVHistory: () => set({ 
    editedCSVHistory: { id: 0, fileName: '', cardType: 'rakuten' }
  }),
});
