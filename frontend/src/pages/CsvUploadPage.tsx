import { useState, useCallback } from 'react';
import { Container } from '@mui/material';
import useStore from '../store';
import { processCSVData, CardType } from '../components/csv-upload-page/utils/csvProcessor/index';
import { 
  Header, 
  Instructions, 
  FileUploader, 
  ResultsTable, 
  Footer,
  CsvPreview
} from '../components/csv-upload-page';

export const CsvUploadPage = () => {
  const { setCardStatementSummaries, cardStatementSummaries } = useStore();
  const [csvFile, setCsvFile] = useState<File | null>(null);
  const [isProcessing, setIsProcessing] = useState(false);
  const [showInstructions, setShowInstructions] = useState(false);
  const [cardType, setCardType] = useState<CardType>('rakuten');

  // CSVを処理して集計データを生成
  const processCSV = useCallback(async () => {
    if (!csvFile) return;

    setIsProcessing(true);
    
    try {
      const summaries = await processCSVData(csvFile, cardType);
      setCardStatementSummaries(summaries);
    } catch (error) {
      console.error('CSV処理エラー:', error);
      alert('CSVの処理中にエラーが発生しました。');
    } finally {
      setIsProcessing(false);
    }
  }, [csvFile, cardType, setCardStatementSummaries]);

  const clearResults = useCallback(() => {
    setCsvFile(null);
    setCardStatementSummaries([]);
  }, [setCardStatementSummaries]);

  return (
    <Container maxWidth="lg" sx={{ py: 4 }}>
      <Header 
        showInstructions={showInstructions} 
        setShowInstructions={setShowInstructions} 
      />
      
      <Instructions showInstructions={showInstructions} />
      
      <FileUploader 
        csvFile={csvFile}
        setCsvFile={setCsvFile}
        isProcessing={isProcessing}
        processCSV={processCSV}
        cardType={cardType}
        setCardType={setCardType}
      />
      
      {/* CSVプレビューコンポーネントを追加 */}
      <CsvPreview file={csvFile} maxRows={10} />
      
      <ResultsTable 
        cardStatementSummaries={cardStatementSummaries}
        clearResults={clearResults}
      />
      
      <Footer />
    </Container>
  );
};
