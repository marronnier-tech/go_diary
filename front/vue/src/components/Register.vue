<template>
  <div class="login userAdmin">
    <form class="form-signin" @submit.prevent="registerForm">
      <!-- <img class="mb-4" src="" alt="" width="72" height="72" /> -->
      <h1 class="h3 mb-3 font-weight-normal">登録する</h1>
      <label for="name" class="sr-only">ユーザー名</label>
      <input
        v-model="name"
        id="name"
        class="form-control"
        placeholder="ユーザー名"
        required
        autofocus
      />
      <label for="password" class="sr-only">パスワード</label>
      <input
        v-model="password"
        type="password"
        id="password"
        class="form-control"
        placeholder="パスワード"
        required
      />

      <button class="btn btn-lg btn-primary btn-block" type="submit">
        登録する</button
      ><br />
      <p>登録済みの方は<router-link to="login">ログイン</router-link></p>
    </form>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      name: "",
      password: "",
    };
  },
  methods: {
    registerForm() {
      const params = new URLSearchParams();
      params.append("name", this.name);
      params.append("password", this.password);
      axios.post("/register", params).then((res) => {
        this.$router.push({ name: "Mypage" });
      });
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
a {
  color: #42b983;
}
</style>
