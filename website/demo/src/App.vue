<template>
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <Article msg="Article List" />
    <!-- <transition name="fade" mode="out-in">
        <router-view></router-view>
    </transition> -->
  </div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
import Article from './components/Article.vue'

export default {
  name: 'App',
  components: {
    // HelloWorld
    Article
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
