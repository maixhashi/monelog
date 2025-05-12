import { useState, useCallback } from 'react';
import { Container, Alert, Snackbar, Button, Box, CircularProgress } from '@mui/material';
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
  CSVHistoryList,
  YearPagination,
  MonthSelector
} from '../components/csv-upload-page';
import { useMutateCardStatements } from '../hooks/mutateHooks/useMutateCardStatements';
import { useQueryCardStatementsByMonth } from '../hooks/queryHooks/useQueryCardStatementsByMonth';
import { useQueryCsvHistoriesByMonth } from '../hooks/queryHooks/useQueryCsvHistoriesByMonth';

export const CsvUploadPage = () => {
  const { 
    setCardStatementSummaries,
    cardStatementSummaries,
    selectedYear,
    selectedMonth,
    setSelectedYear,
    setSelectedMonth
  } = useStore();
  const [csvFile, setCsvFile] = useState<File | null>(null);
  const [isProcessing, setIsProcessing] = useState(false);
  const [isSaving, setIsSaving] = useState(false);
  const [showInstructions, setShowInstructions] = useState(false);
  const [cardType, setCardType] = useState<CardType>('rakuten');
  const [error, setError] = useState<string | null>(null);
  const [previewData, setPreviewData] = useState<boolean>(false);

  // APIミューテーションフックを使用
  const { previewCSVMutation, saveCardStatementsMutation } = useMutateCardStatements();
  
  // 選択中の年月のカード明細を取得
  const { 
    data: monthlyCardStatements = [], 
    isLoading: isCardStatementsLoading 
  } = useQueryCardStatementsByMonth(selectedYear, selectedMonth);
  
  // 選択中の年月のCSV履歴を取得
  const { 
    data: monthlyCsvHistories = [], 
    isLoading: isCsvHistoriesLoading 
  } = useQueryCsvHistoriesByMonth(selectedYear, selectedMonth);

  // カード明細をソート（statementNo, paymentCount順）
  const sortedCardStatements = [...monthlyCardStatements].sort((a, b) => {
    if (a.statementNo !== b.statementNo) return a.statementNo - b.statementNo;
    return a.paymentCount - b.paymentCount;
  });

  // CSV履歴をソート（作成日時の降順）
  const sortedCsvHistories = [...monthlyCsvHistories].sort((a, b) => 
    new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
  );

  // CSVをアップロードしてプレビュー（データベースには保存しない）
  const processCSV = useCallback(async () => {
    if (!csvFile) return;

    setIsProcessing(true);
    setError(null);
    
    try {
      // 年・月を渡す
      const result = await previewCSVMutation.mutateAsync({
        file: csvFile,
        cardType: cardType,
        year: selectedYear,
        month: selectedMonth,
      });
      
      setCardStatementSummaries(result);
      setPreviewData(true);
    } catch (error: any) {
      console.error('CSVプレビューエラー:', error);
      setError(error.message || 'CSVのプレビュー中にエラーが発生しました。');
    } finally {
      setIsProcessing(false);
    }
  }, [csvFile, cardType, selectedYear, selectedMonth, setCardStatementSummaries, previewCSVMutation]);

  // プレビューデータをデータベースに保存
  const saveData = useCallback(async () => {
    if (cardStatementSummaries.length === 0) return;

    setIsSaving(true);
    setError(null);
    
    try {
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

      // 年・月を渡す
      const result = await saveCardStatementsMutation.mutateAsync({
        cardStatements: mappedStatements,
        cardType: cardType,
        year: selectedYear,
        month: selectedMonth,
      });
      
      setPreviewData(false);
    } catch (error: any) {
      console.error('データ保存エラー:', error);
      setError(error.message || 'データの保存中にエラーが発生しました。');
    } finally {
      setIsSaving(false);
    }
  }, [cardStatementSummaries, cardType, selectedYear, selectedMonth, saveCardStatementsMutation]);

  const clearResults = useCallback(() => {
    setCsvFile(null);
    setCardStatementSummaries([]);
    setError(null);
    setPreviewData(false);
  }, [setCardStatementSummaries]);

  return (
    <Container maxWidth="lg" sx={{ py: 4 }}>
      {/* 年・月セレクターを追加 */}
      <YearPagination year={selectedYear} setYear={setSelectedYear} />
      <MonthSelector month={selectedMonth} setMonth={setSelectedMonth} />
      
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
      {csvFile && (
        <CSVHistorySaver
          file={csvFile}
          year={selectedYear}
          month={selectedMonth}
          cardType={cardType}
        />
      )}
      
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
        
        {/* プレビューデータがある場合はそれを表示、なければ月別データを表示 */}
        {previewData ? (
          <ResultsTable 
            cardStatementSummaries={cardStatementSummaries}
            clearResults={clearResults}
            isPreviewData={previewData}
          />
        ) : (
          <>
            {isCardStatementsLoading ? (
              <Box sx={{ display: 'flex', justifyContent: 'center', my: 4 }}>
                <CircularProgress />
              </Box>
            ) : (
              <ResultsTable 
                cardStatementSummaries={sortedCardStatements}
                clearResults={clearResults}
                isPreviewData={false}
              />
            )}
          </>
        )}
      </Box>
      
      {/* CSV履歴一覧コンポーネント */}
      {isCsvHistoriesLoading ? (
        <Box sx={{ display: 'flex', justifyContent: 'center', my: 4 }}>
          <CircularProgress />
        </Box>
      ) : (
        <CSVHistoryList histories={sortedCsvHistories} isLoading={isCsvHistoriesLoading} />
      )}
      
      <Footer />
    </Container>
  );
};
