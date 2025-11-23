import {
  createRouter,
  createWebHistory,
  type RouteRecordRaw,
} from "vue-router";

import Home from "@/views/Home.vue";

const routes: Array<RouteRecordRaw> = [
  { path: "/", name: "Home", component: Home },
  {
    path: "/:catchAll(.*)",
    name: "NotFound",
    component: () => import("@/views/NotFound.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
