import { useState, useCallback } from 'react';
import { Container, Alert, Snackbar } from '@mui/material';
import useStore from '../store';
import { CardType } from '../types/cardType';
import { 
  Header, 
  Instructions, 
  FileUploader, 
  ResultsTable, 
  Footer,
  CsvPreview
} from '../components/csv-upload-page';
import { useMutateCardStatements } from '../hooks/mutateHooks/useMutateCardStatements';

export const CsvUploadPage = () => {
  const { setCardStatementSummaries, cardStatementSummaries } = useStore();
  const [csvFile, setCsvFile] = useState<File | null>(null);
  const [isProcessing, setIsProcessing] = useState(false);
  const [showInstructions, setShowInstructions] = useState(false);
  const [cardType, setCardType] = useState<CardType>('rakuten');
  const [error, setError] = useState<string | null>(null);

  // APIミューテーションフックを使用
  const { uploadCSVMutation } = useMutateCardStatements();

  // CSVをアップロードして処理
  const processCSV = useCallback(async () => {
    if (!csvFile) return;

    setIsProcessing(true);
    setError(null);
    
    try {
      // バックエンドAPIを使用してCSVをアップロード・処理
      const result = await uploadCSVMutation.mutateAsync({
        file: csvFile,
        cardType: cardType
      });
      
      // 結果をストアに保存
      setCardStatementSummaries(result);
    } catch (error: any) {
      console.error('CSV処理エラー:', error);
      setError(error.message || 'CSVの処理中にエラーが発生しました。');
    } finally {
      setIsProcessing(false);
    }
  }, [csvFile, cardType, setCardStatementSummaries, uploadCSVMutation]);

  const clearResults = useCallback(() => {
    setCsvFile(null);
    setCardStatementSummaries([]);
    setError(null);
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
      
      {/* CSVプレビューコンポーネント */}
      <CsvPreview file={csvFile} maxRows={10} />
      
      {/* エラー表示 */}
      <Snackbar 
        open={!!error} 
        autoHideDuration={6000} 
        onClose={() => setError(null)}
      >
        <Alert severity="error" onClose={() => setError(null)}>
          {error}
        </Alert>
      </Snackbar>
      
      <ResultsTable 
        cardStatementSummaries={cardStatementSummaries}
        clearResults={clearResults}
      />
      
      <Footer />
    </Container>
  );
};
