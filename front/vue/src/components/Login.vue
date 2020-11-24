<template>
  <div class="login userAdmin">
    <form class="form-signin">
      <img class="mb-4" src="" alt="" width="72" height="72" />
      <!-- <h1 class="h3 mb-3 font-weight-normal">Please sign in</h1> -->
      <h1 class="h3 mb-3 font-weight-normal">ログインする</h1>
      <!-- <label for="inputEmail" class="sr-only">Email address</label> -->
      <label for="inputEmail" class="sr-only">ユーザー名</label>
      <!-- <inpuzt type="email" id="inputEmail" class="form-control" placeholder="Email address" required autofocus> -->
      <input
        type="name"
        id="inputName"
        class="form-control"
        placeholder="ユーザー名"
        required
        autofocus
      />
      <!-- <label for="inputPassword" class="sr-only">Password</label> -->
      <label for="inputPassword" class="sr-only">パスワード</label>
      <!-- <input type="password" id="inputPassword" class="form-control" placeholder="Password" required> -->
      <input
        type="password"
        id="inputPassword"
        class="form-control"
        placeholder="パスワード"
        required
      />
      <!-- <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button> -->
      <button class="btn btn-lg btn-primary btn-block" type="submit">
        ログイン</button
      ><br />
      <p>あるいは<router-link to="register">登録する</router-link></p>
    </form>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      msg: "Welcome to Your Vue.js App",
    };
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
