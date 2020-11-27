<template>
  <div id="app">
    <router-view></router-view>
    <footer>
      <h4><router-link to="/">ToDotCom</router-link></h4>
      <h6>モチベを上げるSNS</h6>
      <br />
      運営：<a href="https://github.com/tocchy-tocchy" target="_blank"
        >@Marronnier</a
      ><br />
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
  margin: 10px 0px 0px 0px;
  padding: 30px 0px 20px 30px;
}

a {
  color: #42b983;
}
</style>
