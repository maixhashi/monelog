import axios from 'axios';
import { CardType } from '../types/cardType';
import { CSVHistoryResponse, CSVHistoryDetailResponse } from '../types/models/csvHistory';

const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

// CSV履歴一覧を取得
export const fetchCSVHistories = async (): Promise<CSVHistoryResponse[]> => {
  const response = await axios.get<CSVHistoryResponse[]>(`${API_URL}/csv-histories`, {
    withCredentials: true,
  });
  return response.data;
};

// 特定のCSV履歴を取得
export const fetchCSVHistoryById = async (id: number): Promise<CSVHistoryDetailResponse> => {
  const response = await axios.get<CSVHistoryDetailResponse>(`${API_URL}/csv-histories/${id}`, {
    withCredentials: true,
  });
  return response.data;
};

// CSV履歴を保存
export const saveCSVHistory = async (file: File, fileName: string, cardType: CardType): Promise<CSVHistoryResponse> => {
  const formData = new FormData();
  formData.append('file', file);
  formData.append('file_name', fileName);
  formData.append('card_type', cardType);

  const response = await axios.post<CSVHistoryResponse>(`${API_URL}/csv-histories`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
    withCredentials: true,
  });
  return response.data;
};

// CSV履歴を削除
export const deleteCSVHistory = async (id: number): Promise<void> => {
  await axios.delete(`${API_URL}/csv-histories/${id}`, {
    withCredentials: true,
  });
};

// CSV履歴からファイルをダウンロード
export const downloadCSVHistory = async (id: number, fileName: string): Promise<void> => {
  const response = await axios.get(`${API_URL}/csv-histories/${id}/download`, {
    responseType: 'blob',
    withCredentials: true,
  });
  
  // ダウンロードリンクを作成して自動クリック
  const url = window.URL.createObjectURL(new Blob([response.data]));
  const link = document.createElement('a');
  link.href = url;
  link.setAttribute('download', fileName);
  document.body.appendChild(link);
  link.click();
  link.remove();
};
