import { CardStatementSummary } from '../../../../types/models/cardStatement';
import { parseRakutenCSV } from './parserRakutenCSV';
import { parseDcCSV } from './parserDcCSV';
import { parseEposCSV } from './parserEposCSV';

// CardType型のみをエクスポート
export type CardType = 'rakuten' | 'epos' | 'dc';

export const processCSVData = async (
  csvFile: File, 
  cardType: CardType = 'rakuten' // デフォルト値を設定して後方互換性を維持
): Promise<CardStatementSummary[]> => {
  try {
    switch (cardType) {
      case 'rakuten':
        return await parseRakutenCSV(csvFile);
      case 'epos':
        return await parseEposCSV(csvFile);
      case 'dc':
        return await parseDcCSV(csvFile);
      default:
        throw new Error('対応していないカード種類です。');
    }
  } catch (error) {
    console.error('CSV処理エラー:', error);
    throw new Error('CSVの処理中にエラーが発生しました。');
  }
};
