<script setup lang="ts">
import { onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useTaskStore } from "@/stores/taskStore";

const taskStore = useTaskStore();

const { tasks, loading, error } = storeToRefs(taskStore);

onMounted(async () => {
  console.log("Fetching tasks...");
  await taskStore.fetchTasks();
});
</script>

<template>
  <main class="tasks-page">
    <h1>Tasks</h1>

    <section class="tasks-container">
      <div class="task-grid">
        <div v-for="task in tasks" :key="task.id" class="task-card">
          <h3 class="task-title">
            {{ task.title }}
          </h3>
          <p class="task-description">
            {{ task.description }}
          </p>
        </div>
      </div>
    </section>
  </main>
</template>

<style lang="scss" scoped>
.tasks-page {
  margin: 2rem;
}

.tasks-container {
  margin: 2rem;
}

.task-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, max-content));
  gap: 20px;
}

.task-card {
  background-color: rgba($color-text, 10%);
  border: 1px solid $color-primary;
  border-radius: 0.25rem;
  padding: 1rem;
  height: 10rem;
}
</style>
