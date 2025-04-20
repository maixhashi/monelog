import { create } from 'zustand'
import { State } from './state/index'
import { createUserSlice } from './slices/userSlice'
import { createTaskSlice } from './slices/taskSlice'
import { createCardStatementSlice } from './slices/cardStatementSlice'
import { createCSVHistorySlice } from './slices/csvHistorySlice'

const useStore = create<State>((...args) => ({
  ...createUserSlice(...args),
  ...createTaskSlice(...args),
  ...createCardStatementSlice(...args),
  ...createCSVHistorySlice(...args),
}))

export default useStore