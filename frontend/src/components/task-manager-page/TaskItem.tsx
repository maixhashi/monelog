import { FC, memo } from 'react'
import { PencilIcon, TrashIcon } from '@heroicons/react/24/solid'
import useStore from '../../store'
import { Task } from '../../types'
import { useMutateTask } from '../../hooks/mutateHooks/useMutateTask'
import '../../styles/components/task-manager-page/TaskItem.css'

// Taskから必要なプロパティのみを取り出す
// 明示的に型を非オプショナルにする
type TaskItemProps = Required<Pick<Task, 'id' | 'title'>>

const TaskItemMemo: FC<TaskItemProps> = ({
  id,
  title,
}) => {
  const updateTask = useStore((state) => state.updateEditedTask)
  const { deleteTaskMutation } = useMutateTask()
  return (
    <li className="task-item">
      <span className="task-title">{title}</span>
      <div className="task-actions">
        <PencilIcon
          className="task-icon edit"
          onClick={() => {
            updateTask({
              id: id,
              title: title,
            })
          }}
        />
        <TrashIcon
          className="task-icon delete"
          onClick={() => {
            deleteTaskMutation.mutate(id)
          }}
        />
      </div>
    </li>
  )
}
export const TaskItem = memo(TaskItemMemo)