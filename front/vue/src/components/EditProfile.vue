<template>
  <div class="profile">
    <div class="top-copy">
      <h1>プロフィール</h1>
      <p>公開したくない情報は空欄でかまいません。<br /></p>
    </div>
    <div class="whole-prof">
      <div class="prof-item">
        <h2>ユーザー名</h2>
        <p>{{ UserInfo.Name }}</p>
      </div>

      <div class="prof-item">
        <h2>ハンドルネーム</h2>
        <p>{{ UserInfo.HN }}</p>
      </div>
      <div class="prof-item">
        <h2>目標</h2>
        <p>{{ UserInfo.FinalGoal }}</p>
      </div>
      <div class="prof-item">
        <h2>自由プロフィール</h2>
        <p>{{ UserInfo.Profile }}</p>
      </div>
      <div class="prof-item">
        <h2>Twitter</h2>
        <p>{{ UserInfo.Twitter }}</p>
      </div>
      <div class="prof-item">
        <h2>Instagram</h2>
        <p>{{ UserInfo.Instagram }}</p>
      </div>
      <div class="prof-item">
        <h2>Facebook</h2>
        <p>{{ UserInfo.Facebook }}</p>
      </div>
      <div class="prof-item">
        <h2>Github</h2>
        <p>{{ UserInfo.Github }}</p>
      </div>
      <div class="prof-item">
        <h2>その他URL</h2>
        <p>{{ UserInfo.URL }}</p>
      </div>
      <b-button :pressed="false" variant="info" class="edit-profile"
        >編集する</b-button
      >
    </div>
  </div>
</template>

<script>
export default {
  name: "Profile",
  data() {
    return {
      UserInfo: {
        Name: "meow",
        HN: "こねこ",
        FinalGoal: "世界一の猫",
        Profile: "にゃー",
        Twitter: "neco",
        Instagram: "necco",
        Facebook: "necco",
        Github: "neoneo",
        URL: "http://cat.com/",
      },
    };
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.whole-prof {
  text-align: left;
  margin: 5% 0% 0% 0%;
}
.prof-item {
  margin-bottom: 3em;
}
.prof-item p {
  margin-left: 1.5em;
  margin-top: 0.8em;
}
h2 {
  font-weight: normal;
  font-size: 1.3em;
  border-bottom: dashed 1px #4fa4b3;
  padding: 0em 0em 0.2em 0.5em;
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
