import { definitions } from '../api/generated';

/**
 * タスクモデルの型定義
 * 
 * このファイルではタスク関連のすべての型を定義します。
 * APIから自動生成された型を基に、アプリケーション固有の型を定義します。
 */

// APIから生成されたタスクのレスポンス型
export type TaskResponse = definitions["model.TaskResponse"];

// APIから生成されたタスクのリクエスト型
export type TaskRequest = definitions["model.TaskRequest"];

/**
 * 編集用のタスク型
 * APIのレスポンス型から必要なプロパティを取り出し、必須にしたもの
 */
export type EditedTask = {
  id: NonNullable<TaskResponse["id"]>;
  title: NonNullable<TaskResponse["title"]>;
};

/**
 * 編集用タスクの初期値
 */
export const initialEditedTask: EditedTask = {
  id: 0,
  title: ''
};

/**
 * タスク状態の型定義（Zustandストア用）
 */
export interface TaskState {
  editedTask: EditedTask;
  updateEditedTask: (payload: EditedTask) => void;
  resetEditedTask: () => void;
}
