import { UserState } from '../state/userState';
import { TaskState } from '../state/taskState';

// 全体のアプリケーション状態の型を更新
export type State = UserState & TaskState