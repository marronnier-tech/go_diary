<template>
  <div class="qa-all">
    <h1>Q &amp; A</h1>
    <ul>
      <li class="qanda" v-for="(item, index) in QandA" :key="QandA.index">
        <p class="q">Q{{ index + 1 }}. {{ item.Q }}</p>
        <p class="a">{{ item.A }}</p>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "Top",
  data() {
    return {
      /* そのうちDBにいれる　*/
      QandA: [
        {
          Q: "ToDotCom（とぅどっとこむ）とは？",
          A:
            "三日坊主になりがちなあなたを支えるSNSです。近日中にフォロー機能やメッセージ機能を実装します。",
        },
        {
          Q: "パスワードを忘れてしまいました……",
          A:
            "問い合わせ先やメールアドレスの認証機能をまだ作っていません。申し訳ありませんが、頑張って思い出してください。",
        },
        {
          Q: "ゴールとは？",
          A:
            "ToDotCom独自の機能です。満足いくまで成し遂げた状況を想定しています。「日課をやり切って目標を達成した」「試験に合格した」など、ご自由にお使いください。",
        },
        {
          Q: "あなた（運営）はだあれ？",
          A: "しがない学生が個人で運営しています。",
        },
      ],
    };
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.q {
  text-align: left;
  border-bottom: dashed 2px #6ca9bb;
  padding-left: 0.5em;
}
.a {
  text-align: left;
  text-indent: 1em;
}
.qanda {
  margin: 4em 0em;
}
</style>
