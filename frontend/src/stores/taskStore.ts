import { defineStore } from "pinia";
import api from "@/utils/axios";

type Task = {
  id: string;
  title: string;
  userId: number;
  description: string;
  completed: boolean;
  dueDate: Date | null;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date | null;
};

type TaskState = {
  tasks: Task[];
  loading: boolean;
  error: unknown;
};

export const useTaskStore = defineStore("task", {
  state: (): TaskState => ({
    tasks: [],
    loading: false,
    error: null,
  }),
  actions: {
    async fetchTasks() {
      this.loading = true;
      try {
        const res = await api.get("/tasks");
        this.tasks = res.data;
      } catch (err) {
        this.error = err;
      } finally {
        this.loading = false;
      }
    },
  },
});
