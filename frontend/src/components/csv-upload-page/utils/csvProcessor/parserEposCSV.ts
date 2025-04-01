import { CardStatementSummary } from '../../../../types/models/cardStatement';
import { parseDate, safeFormat } from '../dateUtils';
import { ja } from 'date-fns/locale';

export const parseEposCSV = async (csvFile: File): Promise<CardStatementSummary[]> => {
  try {
    const text = await csvFile.text();
    const lines = text.split('\n');
    const summaries: CardStatementSummary[] = [];
    
    // 仮実装: 簡単なサンプルデータを返す
    const sampleDate = new Date();
    
    summaries.push({
      type: '発生',
      statementNo: 1,
      cardType: 'EPOSカード',
      description: 'EPOSカードサンプルデータ',
      useDate: safeFormat(sampleDate, 'yyyy/MM/dd'),
      paymentDate: safeFormat(sampleDate, 'yyyy/MM/dd'),
      paymentMonth: safeFormat(sampleDate, 'yyyy年MM月', { locale: ja }),
      amount: 5000,
      totalChargeAmount: 5000,
      chargeAmount: 5000,
      remainingBalance: 0,
      paymentCount: 1,
      installmentCount: 1,
      annualRate: 0,
      monthlyRate: 0
    });

    console.log('EPOSカードCSV処理完了 - 仮実装');
    return summaries;
  } catch (error) {
    console.error('EPOSカードCSV処理エラー:', error);
    throw new Error('EPOSカードCSVの処理中にエラーが発生しました。');
  }
};
