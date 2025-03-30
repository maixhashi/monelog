import axios from 'axios';

export interface Task {
  id?: number;
  title: string;
}

export const fetchTasks = async (): Promise<Task[]> => {
  const response = await axios.get(`${process.env.REACT_APP_API_URL}/tasks`);
  return response.data;
};

export const fetchTask = async (id: number): Promise<Task> => {
  const response = await axios.get(`${process.env.REACT_APP_API_URL}/tasks/${id}`);
  return response.data;
};

export const createTask = async (task: Task): Promise<Task> => {
  const response = await axios.post(`${process.env.REACT_APP_API_URL}/tasks`, task);
  return response.data;
};

export const updateTask = async (task: Task): Promise<Task> => {
  const response = await axios.put(`${process.env.REACT_APP_API_URL}/tasks/${task.id}`, task);
  return response.data;
};

export const deleteTask = async (id: number): Promise<void> => {
  await axios.delete(`${process.env.REACT_APP_API_URL}/tasks/${id}`);
};
