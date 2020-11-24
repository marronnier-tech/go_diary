<template>
  <div class="register userAdmin">
    <form class="form-signin">
      <img class="mb-4" src="" alt="" width="72" height="72" />
      <!-- <h1 class="h3 mb-3 font-weight-normal">Please sign in</h1> -->
      <h1 class="h3 mb-3 font-weight-normal">登録する</h1>
      <!-- <label for="inputEmail" class="sr-only">Email address</label> -->
      <label for="inputName" class="sr-only">ユーザー名</label>
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
        登録する</button
      ><br />
      <p>登録済みの方は<router-link to="login">ログイン</router-link></p>
    </form>
  </div>
</template>

<script>
export default {
  name: "Register",
  data() {
    return {
      msg: "Welcome to Your Vue.js App",
    };
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
