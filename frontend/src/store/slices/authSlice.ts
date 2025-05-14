import { StateCreator } from 'zustand';
import { AuthState } from '../state/authState';
import { User } from '../../types/models/user';

export const createAuthSlice: StateCreator<AuthState> = (set) => ({
  isAuthenticated: false,
  user: null,
  isLoading: true,
  setAuth: (isAuthenticated: boolean, user: User | null) => set({
    isAuthenticated,
    user,
    isLoading: false,
  }),
  clearAuth: () => set({
    isAuthenticated: false,
    user: null,
    isLoading: false,
  }),
});