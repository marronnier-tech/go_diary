<template>
  <div class="todo">
    <div class="topcopy">
      <h1>みんなのToDo</h1>
      同じ目標の仲間を探そう！
    </div>

    <div class="todo-obj" v-for="todo in todos.TodoArray" :key="todo.TodoIbj">
      <p class="user">@{{ todo.User.UserName }}</p>
      <h4 class="content">{{ todo.TodoObj.Content }}</h4>
      <div class="achieved-info">
        <p class="count">ToDoできた日数：{{ todo.TodoObj.Count }}日</p>
        <p class="last-achieved">最終達成日：{{ todo.TodoObj.LastAchieved }}</p>
      </div>
      <p class="created-at">{{ todo.TodoObj.CreatedAt }}</p>
    </div>
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
.todo-obj {
  text-align: left;
  margin: 5% 10%;
  padding: 15px 15px 5px 20px;
  background: #e4ecec;
  border-left: solid 6px #ccdbdb;
  box-shadow: 0 3px 3px rgba(0, 0, 0, 0.22);
}
.topcopy {
  margin: 1% auto 5% auto;
}
h4 {
  font-weight: normal;
  font-size: 1.3em;
  line-height: 0.8em;
  padding-left: 1em;
  margin-bottom: 1em;
}
.achieved-info p {
  line-height: 0.5em;
  font-size: 0.9em;
}
.created-at {
  font-size: 0.8em;
  text-align: right;
}

a {
  color: #42b983;
}
</style>
