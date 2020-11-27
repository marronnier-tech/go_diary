<template>
  <div class="goal">
    <div class="top-copy">
      <h1>ゴール！</h1>
      <p>
        あなたのゴールリストを表示します。<br />
        <router-link to="/todo">ゴールとは？</router-link>
      </p>
    </div>
    <ul>
      <li
        class="todo-set todo-goal"
        v-for="goal in goals.GoalObj"
        :key="goal.GoalIndex"
      >
        <h3 class="content">
          <div class="award-line">
            <b-icon icon="award" class="award"></b-icon>{{ goal.Content }}
          </div>
        </h3>
        <div class="achieved-info">
          <p class="last-achieved">
            ゴールまでに実行した日：{{ goal.AchievedCount }}日
          </p>
          <p class="count">ゴール日：{{ goal.GoaledAt }}</p>
        </div>
      </li>
    </ul>
    <br />
    <p>おめでとう！</p>
    <b-button :pressed="false" variant="info" class="add-todo"
      >日課をゴールする!</b-button
    >
    <br />
    ※メンテナンス中
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Goal",
  message: "Not yet",
  data() {
    return {
      goals: [],
    };
  },
  mounted: function () {
    axios.get("/mypage/goal").then((res) => {
      this.goals = res.data.Goal;
      console.log(res);
    });
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  font-size: 1.2em;
}
.award-line {
  margin: 0.5em 0.3em 1.3em 0em;
}
a {
  color: #42b983;
}
.award {
  font-size: 1.4em;
}
</style>
