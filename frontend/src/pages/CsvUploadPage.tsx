import { useState, useCallback } from 'react';
import { Container, Alert, Snackbar, Button, Box } from '@mui/material';
import useStore from '../store';
import { CardType } from '../types/cardType';
import { 
  Header, 
  Instructions, 
  FileUploader, 
  ResultsTable, 
  Footer,
  CsvPreview,
  CSVHistorySaver,
  CSVHistoryList
} from '../components/csv-upload-page';
import { useMutateCardStatements } from '../hooks/mutateHooks/useMutateCardStatements';

export const CsvUploadPage = () => {
  const { setCardStatementSummaries, cardStatementSummaries } = useStore();
  const [csvFile, setCsvFile] = useState<File | null>(null);
  const [isProcessing, setIsProcessing] = useState(false);
  const [isSaving, setIsSaving] = useState(false);
  const [showInstructions, setShowInstructions] = useState(false);
  const [cardType, setCardType] = useState<CardType>('rakuten');
  const [error, setError] = useState<string | null>(null);
  const [previewData, setPreviewData] = useState<boolean>(false);

  // APIミューテーションフックを使用
  const { previewCSVMutation, saveCardStatementsMutation } = useMutateCardStatements();

  // CSVをアップロードしてプレビュー（データベースには保存しない）
  const processCSV = useCallback(async () => {
    if (!csvFile) return;

    setIsProcessing(true);
    setError(null);
    
    try {
      // バックエンドAPIを使用してCSVをプレビュー（DBには保存しない）
      const result = await previewCSVMutation.mutateAsync({
        file: csvFile,
        cardType: cardType
      });
      
      // 結果をストアに保存（一時的なプレビューデータ）
      setCardStatementSummaries(result);
      setPreviewData(true);
    } catch (error: any) {
      console.error('CSVプレビューエラー:', error);
      setError(error.message || 'CSVのプレビュー中にエラーが発生しました。');
    } finally {
      setIsProcessing(false);
    }
  }, [csvFile, cardType, setCardStatementSummaries, previewCSVMutation]);

  // プレビューデータをデータベースに保存
  const saveData = useCallback(async () => {
    if (cardStatementSummaries.length === 0) return;

    setIsSaving(true);
    setError(null);
    
    try {
      // すべての必要なフィールドを明示的にマッピング（キャメルケースを使用）
      const mappedStatements = cardStatementSummaries.map(statement => ({
        type: statement.type || "発生",
        statementNo: statement.statementNo || 0,
        cardType: cardType,
        description: statement.description || "",
        useDate: statement.useDate || "",
        paymentDate: statement.paymentDate || "",
        paymentMonth: statement.paymentMonth || "",
        amount: Number(statement.amount) || 0,
        totalChargeAmount: Number(statement.totalChargeAmount) || 0,
        chargeAmount: Number(statement.chargeAmount) || 0,
        remainingBalance: Number(statement.remainingBalance) || 0,
        paymentCount: Number(statement.paymentCount) || 0,
        installmentCount: Number(statement.installmentCount) || 0,
        annualRate: Number(statement.annualRate) || 0,
        monthlyRate: Number(statement.monthlyRate) || 0,
      }));
      
      console.log('送信するデータ:', JSON.stringify(mappedStatements, null, 2));
      
      const result = await saveCardStatementsMutation.mutateAsync({
        cardStatements: mappedStatements,
        cardType: cardType
      });
      
      // 保存完了後の処理
      setPreviewData(false);
    } catch (error: any) {
      console.error('データ保存エラー:', error);
      setError(error.message || 'データの保存中にエラーが発生しました。');
    } finally {
      setIsSaving(false);
    }
  }, [cardStatementSummaries, cardType, saveCardStatementsMutation]);

  const clearResults = useCallback(() => {
    setCsvFile(null);
    setCardStatementSummaries([]);
    setError(null);
    setPreviewData(false);
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
      
      {/* CSV履歴保存コンポーネント */}
      {csvFile && <CSVHistorySaver file={csvFile} />}
      
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
      
      {/* 結果テーブルと保存ボタン */}
      <Box>
        {previewData && cardStatementSummaries.length > 0 && (
          <Box sx={{ mt: 2, mb: 2, display: 'flex', justifyContent: 'flex-end' }}>
            <Button 
              variant="contained" 
              color="primary" 
              onClick={saveData}
              disabled={isSaving}
            >
              {isSaving ? '保存中...' : 'データを保存する'}
            </Button>
          </Box>
        )}
        
        <ResultsTable 
          cardStatementSummaries={cardStatementSummaries}
          clearResults={clearResults}
          isPreviewData={previewData}
        />
      </Box>
      
      {/* CSV履歴一覧コンポーネント */}
      <CSVHistoryList />
      
      <Footer />
    </Container>
  );
};
