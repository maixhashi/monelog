import { create } from 'zustand'
import { State } from './state/index'
import { createUserSlice } from './slices/userSlice'
import { createTaskSlice } from './slices/taskSlice'
import { createCardStatementSlice } from './slices/cardStatementSlice'
import { createCSVHistorySlice } from './slices/csvHistorySlice'
import { createAuthSlice } from './slices/authSlice'
import { AuthState } from './state/authState'

// ステートの型を拡張
export type RootState = State & AuthState

const useStore = create<RootState>((...args) => ({
  ...createUserSlice(...args),
  ...createTaskSlice(...args),
  ...createCardStatementSlice(...args),
  ...createCSVHistorySlice(...args),
  ...createAuthSlice(...args),
}))

export default useStore