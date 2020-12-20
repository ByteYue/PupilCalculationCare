<template>
  <div class="practicePage">
    <el-header class="header">
      <el-popconfirm icon="el-icon-info" @click="backToCreate()">
        <el-button slot="reference">返回主页</el-button>
      </el-popconfirm>
      <div class="user">
        <img src="../assets/img/头像.jpg" alt="头像" @click="goToUser" />
      </div>
    </el-header>
    <el-main class="main"> 这是用户主页 </el-main>
  </div>
</template>

<script>
export default {
  data: {
    user: [],
  },
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
  opacity: 1;
  height: 700px;
  width: 100%;
}

.footer {
  background-color: rgb(196, 212, 230);
  height: 80px;
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
}

.footer button {
  position: absolute;
  right: 10px;
  bottom: 10px;
}
</style>