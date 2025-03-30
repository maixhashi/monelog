import { useState, useCallback } from 'react';
import { Container } from '@mui/material';
import useStore from '../store';
import { processCSVData } from '../components/csv-upload-page/utils/csvProcessor';
import { 
  Header, 
  Instructions, 
  FileUploader, 
  ResultsTable, 
  Footer 
} from '../components/csv-upload-page';

export const CsvUploadPage = () => {
  const { setCardStatementSummaries, cardStatementSummaries } = useStore();
  const [csvFile, setCsvFile] = useState<File | null>(null);
  const [isProcessing, setIsProcessing] = useState(false);
  const [showInstructions, setShowInstructions] = useState(false);

  // CSVを処理して集計データを生成
  const processCSV = useCallback(async () => {
    if (!csvFile) return;

    setIsProcessing(true);
    
    try {
      const summaries = await processCSVData(csvFile);
      setCardStatementSummaries(summaries);
    } catch (error) {
      console.error('CSV処理エラー:', error);
      alert('CSVの処理中にエラーが発生しました。');
    } finally {
      setIsProcessing(false);
    }
  }, [csvFile, setCardStatementSummaries]);

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
      />
      
      <ResultsTable 
        cardStatementSummaries={cardStatementSummaries}
        clearResults={clearResults}
      />
      
      <Footer />
    </Container>
  );
};
