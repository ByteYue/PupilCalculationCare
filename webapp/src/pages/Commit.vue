<template>
  <div class="practicePage">
    <el-header class="header">
      <el-button @click="backToCreate()">返回主页</el-button>
      <div class="user">
        <img src="../assets/img/头像.jpg" alt="头像" @click="goToUser" />
      </div>
    </el-header>
    <el-main class="main">
      <div id="puzzles">
        <!--v-for循环对象数组-->
      </div>
    </el-main>
    <div class="footer">
      <div class="word">
        芜湖！你得到了
        <p id="score">90</p>
        分！
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data: {},
  methods: {
    backToCreate() {
      //页面跳转
      this.$router.push("/create");
    },
    goToUser() {
      //页面跳转
      this.$router.push("/user");
    },
  },
};
</script>

<style scoped>
.practicePage {
  width: 100%;
  height: 0;
}

.header {
  background-color: rgb(196, 212, 230);
  height: 100px;
}

.header img {
  height: 50px;
  width: 50px;
  position: absolute;
  right: 10px;
  top: 5px;
  border: 2px solid rgb(255, 255, 255);
  border-radius: 10px;
  cursor: pointer;
}

.header button {
  position: absolute;
  top: 10px;
  left: 10px;
}

.main {
  background-color: rgb(255, 255, 255);
  height: 700px;
  width: 100%;
}

.footer {
  background-color: rgb(255, 255, 255);
  color: #000;
  border: 1px solid #000;
  height: 120px;
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
}

.footer .word {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}

#score {
  color: red;
  font-size: larger;
}
</style>