import { EditedCSVHistory, CSVHistoryResponse } from '../../types/models/csvHistory';

// CSVHistoryState の型定義
export type CSVHistoryState = {
  editedCSVHistory: EditedCSVHistory;
  csvHistories: CSVHistoryResponse[];
  updateEditedCSVHistory: (payload: EditedCSVHistory) => void;
  setCSVHistories: (histories: CSVHistoryResponse[]) => void;
  resetEditedCSVHistory: () => void;
}

// CSVHistoryState の初期状態
export const initialCSVHistoryState: CSVHistoryState = {
  editedCSVHistory: { id: 0, fileName: '', cardType: 'rakuten' },
  csvHistories: [],
  updateEditedCSVHistory: () => {},
  setCSVHistories: () => {},
  resetEditedCSVHistory: () => {},
}
