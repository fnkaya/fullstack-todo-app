<template>
  <div id="todo-list-container">
    <TodoItem v-for="todo of todos"
              :key="todo.id"
              :todo="todo"
    />
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import TodoItem from '@/components/TodoItem.vue';
import { FETCH_TODOS } from '@/store/actions';

export default {
  name: 'TodoList',
  components: { TodoItem },
  computed: {
    ...mapGetters(['todos']),
  },
  created() {
    this.$store.dispatch(FETCH_TODOS);
  },
};
</script>

<style scoped>

</style>
