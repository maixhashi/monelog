import { create } from 'zustand'
import { State } from './state/index'
import { createUserSlice } from './slices/userSlice'
import { createTaskSlice } from './slices/taskSlice'

const useStore = create<State>((...args) => ({
  ...createUserSlice(...args),
  ...createTaskSlice(...args),
}))

export default useStore