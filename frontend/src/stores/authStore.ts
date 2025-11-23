import { defineStore } from "pinia";
import { ref, computed } from "vue";

interface AuthState {
  userId: string | null;
  token: string | null;
}

export const useAuthStore = defineStore("auth", () => {
  const testUserID = 1;

  const userId = ref<string | number | null>(
    localStorage.getItem("userID") || testUserID
  );
  const token = ref<string | null>(localStorage.getItem("token") || null);

  const isAuthenticated = computed(() => !!userId.value && !!token.value);

  const login = async (data: { id: string; token: string }): Promise<void> => {
    userId.value = data.id;
    token.value = data.token;

    localStorage.setItem("userID", data.id);
    localStorage.setItem("token", data.token);
  };

  const logout = () => {
    userId.value = null;
    token.value = null;

    localStorage.removeItem("userID");
    localStorage.removeItem("token");
  };

  const initializeAuth = () => {
    if (isAuthenticated.value) {
      console.log("User session found and loaded.");
    }
  };

  return {
    userId,
    token,

    isAuthenticated,

    login,
    logout,
    initializeAuth,
  };
});
