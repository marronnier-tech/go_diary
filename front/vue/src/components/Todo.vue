<template>
  <div class="todo">
    <div class="top-copy">
      <h1>みんなの日課</h1>
      同志を探そう！
    </div>
    <ul>
      <li
        class="todo-set"
        v-for="todo in todos.TodoArray"
        :key="todo.TodoIndex"
      >
        <p class="user">
          <b-icon icon="person-circle" class="user-icon"></b-icon
          >{{ todo.User.UserName }}
        </p>
        <h4 class="content">{{ todo.TodoObj.Content }}</h4>
        <div class="achieved-info">
          <p class="count">達成日数：{{ todo.TodoObj.Count }}日</p>
          <p class="last-achieved">
            最終達成日：{{ todo.TodoObj.LastAchieved }}
          </p>
        </div>
        <p class="created-at">{{ todo.TodoObj.CreatedAt }}</p>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Todo",

  data() {
    return {
      todos: [],
    };
  },
  mounted: function () {
    axios.get("/todo").then((res) => {
      this.todos = res.data;
      console.log(res);
    });
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h4 {
  padding-left: 1em;
}

a {
  color: #42b983;
}
</style>
