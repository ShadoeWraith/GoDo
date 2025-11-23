<template>
  <div class="not-found-container">
    <h1 class="status-code">404</h1>
    <h2 class="message">Page Not Found</h2>
    <p>
      Oops! We can't find the page you're looking for. It might have been moved
      or deleted.
    </p>

    <router-link to="/" class="home-link">
      Go back to the homepage
    </router-link>
  </div>
</template>

<script setup>
import { onMounted } from "vue";
onMounted(() => {
  document.title = "404 - Not Found";
});
</script>

<style lang="scss" scoped>
.not-found-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  min-height: 100vh;
  padding: 2rem;
  background-color: lighten($color-background, 5%);
  color: $color-text;
}

.status-code {
  font-size: 8rem;
  font-weight: bold;
  color: $color-primary;
  margin-bottom: 0;
}

.message {
  font-size: 2rem;
  margin-top: 0;
  margin-bottom: 1rem;
}

.home-link {
  margin-top: 1.5rem;
  padding: 0.75rem 1.5rem;
  border: 1px solid $color-primary;
  border-radius: 4px;
  background-color: $color-primary;
  color: white;
  text-decoration: none;
  transition: background-color 0.3s;

  &:hover {
    background-color: darken($color-primary, 10%);
  }
}
</style>
