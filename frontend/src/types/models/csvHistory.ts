import { definitions } from '../api/generated';

// バックエンドのレスポンス型
export type CSVHistoryResponse = definitions['model.CSVHistoryResponse'];
export type CSVHistoryDetailResponse = definitions['model.CSVHistoryDetailResponse'];

/**
 * 編集用のCSV履歴型
 */
export type EditedCSVHistory = {
  id: number;
  fileName: string;
  cardType: string;
};

/**
 * 編集用CSV履歴の初期値
 */
export const initialEditedCSVHistory: EditedCSVHistory = {
  id: 0,
  fileName: '',
  cardType: 'rakuten',
};

/**
 * CSV履歴状態の型定義（Zustandストア用）
 */
export interface CSVHistoryState {
  editedCSVHistory: EditedCSVHistory;
  csvHistories: CSVHistoryResponse[];
  updateEditedCSVHistory: (payload: EditedCSVHistory) => void;
  setCSVHistories: (histories: CSVHistoryResponse[]) => void;
  resetEditedCSVHistory: () => void;
}
