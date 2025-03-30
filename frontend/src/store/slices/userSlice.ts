import { StateCreator } from 'zustand';
import { EditedUser, initialEditedUser } from '../types/userTypes';
import { UserState } from '../state/userState';

/**
 * ユーザー関連のZustandストアスライス
 */
export const createUserSlice: StateCreator<UserState> = (set) => ({
  editedUser: initialEditedUser,
  updateEditedUser: (payload: EditedUser) => set({
    editedUser: payload
  }),
  resetEditedUser: () => set({ 
    editedUser: initialEditedUser 
  }),
});
