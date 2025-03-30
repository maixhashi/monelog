import { useMutation, useQueryClient } from '@tanstack/react-query'
import { createTask, updateTask, deleteTask, Task } from '../../api/tasks'

export const useMutateTask = () => {
  const queryClient = useQueryClient()

  const createTaskMutation = useMutation(
    (task: Task) => createTask(task),
    {
      onSuccess: (res) => {
        const previousTasks = queryClient.getQueryData<Task[]>(['tasks']) || []
        queryClient.setQueryData(['tasks'], [...previousTasks, res])
      },
    }
  )

  const updateTaskMutation = useMutation(
    (task: Task) => updateTask(task),
    {
      onSuccess: (res) => {
        const previousTasks = queryClient.getQueryData<Task[]>(['tasks']) || []
        queryClient.setQueryData(
          ['tasks'],
          previousTasks.map((task) => (task.id === res.id ? res : task))
        )
      },
    }
  )

  const deleteTaskMutation = useMutation(
    (id: number) => deleteTask(id),
    {
      onSuccess: (_, variables) => {
        const previousTasks = queryClient.getQueryData<Task[]>(['tasks']) || []
        queryClient.setQueryData(
          ['tasks'],
          previousTasks.filter((task) => task.id !== variables)
        )
      },
    }
  )

  return {
    createTaskMutation,
    updateTaskMutation,
    deleteTaskMutation,
  }
}