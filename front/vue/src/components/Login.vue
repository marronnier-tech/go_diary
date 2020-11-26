<template>
  <div class="login userAdmin">
    <form class="form-signin" @submit.prevent="submitForm">
      <!-- <img class="mb-4" src="" alt="" width="72" height="72" /> -->
      <h1 class="h3 mb-3 font-weight-normal">ログインする</h1>
      <label for="name" class="sr-only">ユーザー名</label>
      <input
        v-model="name"
        type="name"
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
        ログイン</button
      ><br />

      <p>あるいは<router-link to="register">登録する</router-link></p>
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
    submitForm() {
      const params = new URLSearchParams();
      params.append("name", this.name);
      params.append("password", this.password);
      axios
        .post("/login", params)
        .then((res) => {
          this.$router.push({ name: "Mypage" });
        })
        .catch((error) => {
          alert("ユーザー名もしくはパスワードが違います");
        });
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/* .login {
  display: -ms-flexbox;
  display: flex;
  -ms-flex-align: center;
  align-items: center;
  padding-top: 40px;
  padding-bottom: 40px;
  background-color: #f5f5f5;
  height: 100%;
} */
</style>
