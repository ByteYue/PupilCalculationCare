<template>
  <div id="app">
    <!-- 路由占位符 -->
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
body {
  width: 100%;
  background-color: rgb(239, 239, 239);
}
.loginPage {
  position: fixed;
  top: 0;
  width: 100%;
}
</style>
