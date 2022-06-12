<template>
  <div id="todo-item-container">
    <span id="todo-text">{{todo.text}}</span>
    <input id="todo-done" type="checkbox" :checked="todo.done" />
  </div>
</template>

<script>
export default {
  name: 'TodoItem',
  props: {
    todo: {
      type: Object,
    },
  },
};
</script>

<style scoped>

</style>
