import axios from "axios";
import { useAuthStore } from "@/stores/authStore.ts";

const api = axios.create({
  baseURL: "/api",
  headers: {
    "Content-Type": "application/json",
  },
});

api.interceptors.request.use(
  (config) => {
    const authStore = useAuthStore();

    const userId = authStore.userId;

    if (userId) {
      config.headers["User-ID"] = userId;
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default api;
