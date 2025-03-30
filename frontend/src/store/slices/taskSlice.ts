import { StateCreator } from 'zustand';
import { TaskState, EditedTask, initialEditedTask } from '../../types/models/task';

/**
 * タスク関連のZustandストアスライス
 */
export const createTaskSlice: StateCreator<TaskState> = (set) => ({
  editedTask: initialEditedTask,
  updateEditedTask: (payload: EditedTask) => set({
    editedTask: payload
  }),
  resetEditedTask: () => set({ 
    editedTask: initialEditedTask 
  }),
});