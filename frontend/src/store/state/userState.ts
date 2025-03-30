import { EditedUser } from '../types/userTypes';

// UserState の型定義 - moved here from userTypes.ts
export type UserState = {
  editedUser: EditedUser;
  updateEditedUser: (payload: EditedUser) => void;
  resetEditedUser: () => void;
}

// UserState の初期状態
export const initialUserState: UserState = {
  editedUser: { id: 0, name: '' },
  updateEditedUser: () => {},
  resetEditedUser: () => {},
}
