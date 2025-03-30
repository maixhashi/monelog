import { useQuery } from '@tanstack/react-query'
import { fetchTasks, Task } from '../../api/tasks'

export const useQueryTasks = () => {
  return useQuery<Task[], Error>({
    queryKey: ['tasks'],
    queryFn: fetchTasks,
    staleTime: Infinity,
  })
}