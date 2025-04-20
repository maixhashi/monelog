import { UserState } from './userState'
import { TaskState } from './taskState'
import { CardStatementState } from './cardStatementState'
import { CSVHistoryState } from './csvHistoryState'

export type State = UserState & TaskState & CardStatementState & CSVHistoryState