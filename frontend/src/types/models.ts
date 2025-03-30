import { definitions } from './api/generated';

// APIの型定義から派生した型

// TaskResponseの型
export type Task = definitions["model.TaskResponse"];

// TaskRequestの型
export type TaskRequest = definitions["model.TaskRequest"];

// 編集用のタスク型
export type EditedTask = {
  id: NonNullable<definitions["model.TaskResponse"]["id"]>;
  title: NonNullable<definitions["model.TaskResponse"]["title"]>;
};

// 初期値の定義
export const initialEditedTask: EditedTask = {
  id: 0,
  title: ''
};

// Zustandのストア用の型定義
export interface TaskState {
  editedTask: EditedTask;
  updateEditedTask: (payload: EditedTask) => void;
  resetEditedTask: () => void;
}

// 他の型定義...
