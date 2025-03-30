import { EditedTask } from '../types/taskTypes';

// TaskState の型定義 - moved here from taskTypes.ts
export type TaskState = {
  editedTask: EditedTask;
  updateEditedTask: (payload: EditedTask) => void;
  resetEditedTask: () => void;
}

// TaskState の初期状態
export const initialTaskState: TaskState = {
  editedTask: { id: 0, title: '' },
  updateEditedTask: () => {},
  resetEditedTask: () => {},
}
