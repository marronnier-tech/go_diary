<template>
  <div id="app">
    <b-navbar toggleable="lg" type="dark" variant="info">
      <b-navbar-brand to="/">ToDotCom</b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item to="/mypage">My ToDo</b-nav-item>
          <b-nav-item to="/goal">My Goaled</b-nav-item>
          <b-nav-item to="/profile">My Profile</b-nav-item>
          <b-nav-item to="/todo">Others ToDo</b-nav-item>
          <b-nav-item to="/logout">Logout</b-nav-item>
          <b-nav-item href="#" disabled>Disabled</b-nav-item>
        </b-navbar-nav>

        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <b-nav-form>
            <b-form-input
              size="sm"
              class="mr-sm-2"
              placeholder="近日公開"
            ></b-form-input>
            <b-button size="sm" class="my-2 my-sm-0" type="submit"
              >検索</b-button
            >
          </b-nav-form>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
    <div class="hello">
      <h1>{{ msg }}</h1>
      <router-link to="/register">register</router-link>
      <br />
    </div>
    <router-view></router-view>
    <footer>
      <h4>ToDotCom ― モチベを上げるSNS</h4>
      運営：<a href="https://github.com/tocchy-tocchy">@Marronnier</a><br />
    </footer>
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
#app {
  background-color: #efefef;
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

footer {
  background-color: #fff;
  text-align: left;
  margin: 100px 0px 0px 0px;
  padding: 30px 0px 20px 30px;
}
</style>
