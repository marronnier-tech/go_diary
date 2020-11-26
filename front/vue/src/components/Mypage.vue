<template>
  <div class="mypage">
    <div class="top-copy">
      <h1>Hello, {{ user.UserName }}!</h1>
      <p>あなたの今の日課はこちらです。</p>
    </div>
    <ul>
      <li
        class="todo-set todo-yet"
        v-for="todo in todosYet"
        :key="todo.TodoIndex"
      >
        <div class="my-todoset">
          <h4 class="content content-yet">
            {{ todo.Content }}
          </h4>
          <b-button variant="info" class="yatta" v-on:click="Yatta(todo.TodoID)"
            >Yatta!</b-button
          >
        </div>
        <div class="achieved-info">
          <p class="count">実行日数：{{ todo.Count }}日</p>
          <p class="last-achieved">最終実行日：{{ todo.LastAchieved }}</p>
          <p class="created-at">{{ todo.CreatedAt }}</p>
        </div>
      </li>
      <li
        class="todo-set todo-get"
        v-for="todo in todosGet"
        :key="todo.TodoIndex"
      >
        <div class="my-todoset">
          <h4 class="content content-set">
            <b-icon icon="bookmark-star-fill" class="check"></b-icon
            >{{ todo.Content }}
          </h4>
          <b-icon
            icon="arrow-return-left"
            class="unget"
            v-on:click="unGet(todo.TodoID)"
            :pressed="false"
          ></b-icon>
        </div>
        <div class="achieved-info">
          <p class="count">実行日数：{{ todo.Count }}日</p>
          <p class="last-achieved">最終実行日：{{ todo.LastAchieved }}</p>
          <p class="created-at">{{ todo.CreatedAt }}</p>
        </div>
      </li>
    </ul>
    <br />
    今日も一日頑張りましょう！<br />
    <br />
    <div class="todo-add">
      <h2>日課を追加する</h2>
      <div class="add-form">
        <b-form @submit.prevent="AddTodo">
          <b-form-input
            id="content"
            placeholder="新しい日課"
            v-model="content"
          ></b-form-input>
          <b-button variant="info" type="submit">追加</b-button>
        </b-form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Todo",
  message: "Not yet",
  data() {
    return {
      user: "",
      todosYet: [],
      todosGet: [],
      content: "",
    };
  },
  mounted: function () {
    this.GetAndSep();
  },
  methods: {
    GetAndSep() {
      axios.get("/mypage").then((res) => {
        this.user = res.data.Todo.User;
        /* ↓あとでまとめる↓ */
        this.todosYet = res.data.Todo.TodoObj.filter(
          (t) => t.TodayAchieved === false
        );
        this.todosGet = res.data.Todo.TodoObj.filter(
          (t) => t.TodayAchieved === true
        );
      });
    },
    AddTodo() {
      const params = new URLSearchParams();
      params.append("content", this.content);
      axios
        .post("/mypage", params)
        .then((postres) => {
          this.GetAndSep();
        })
        .catch((error) => {
          alert("同じ日課が既に登録されています");
          console.log(error);
        });
    },
    Yatta(todo_id) {
      axios.post("mypage/" + todo_id + "/today").then((res) => {
        this.GetAndSep();
      });
    },
    unGet(todo_id) {
      axios.delete("mypage/" + todo_id + "/today").then((res) => {
        this.GetAndSep();
      });
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h4 {
  font-weight: normal;
  padding-top: 0.5em;
  margin-right: 1em;
}
.content {
  margin-bottom: 1.2em;
}
.content-yet {
  padding-left: 1.2em;
}
.achieved-info {
  padding-left: 1em;
}
a {
  color: #42b983;
}
.my-todoset {
  display: flex;
  justify-content: space-between;
}
.yatta {
  margin-right: 1em;
  padding: 0.5em;
  height: 2.5em;
  letter-spacing: 0.05em;
}
.check {
  color: #e46c86;
  margin-right: 0.5em;
}
.todo-add {
  margin-top: 4em;
  text-align: left;
}
.add-form {
  margin: 0.2em 2em 0.2em 2em;
}
h2 {
  font-size: 1.5em;
  margin-bottom: 0.8em;
}
</style>
