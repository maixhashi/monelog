import { User } from '../../types/models/user';

export type AuthState = {
  isAuthenticated: boolean;
  user: User | null;
  isLoading: boolean;
  setAuth: (isAuthenticated: boolean, user: User | null) => void;
  clearAuth: () => void;
};

export const initialAuthState: AuthState = {
  isAuthenticated: false,
  user: null,
  isLoading: true,
  setAuth: () => {},
  clearAuth: () => {},
};