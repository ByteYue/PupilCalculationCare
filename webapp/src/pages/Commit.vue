<template>
  <div class="practicePage">
    <el-header class="header">
      <el-button @click="backToCreate()">返回主页</el-button>
      <div class="user">
        <img src="../assets/img/头像.jpg" alt="头像" @click="goToUser" />
      </div>
    </el-header>
    <el-main class="main">
      <el-main class="main">
        <div class="pheader">
          <el-button @click="start()">开始答题</el-button>
          <div class="testname">
            <div id="testName">练习题</div>
          </div>
        </div>
        <div class="puzzles">
          <el-table
            :data="expressions"
            style="width: 100%"
            :row-class-name="puzzleRow"
          >
            <el-table-column prop="id" width="200" align="right" label="序号">
            </el-table-column>
            <el-table-column
              prop="Expression"
              width="200"
              align="center"
              label="题目"
            >
            </el-table-column>
            <el-table-column
              prop="Input"
              width="100"
              align="center"
              label="输入"
            >
              <el-input></el-input
            ></el-table-column>
            <div style="display: inline-block" v-if="show">
              <el-table-column
                prop="Answer"
                width="200"
                align="center"
                label="答案"
              >
              </el-table-column>
              <el-table-column
                prop="Judge"
                width="200"
                align="center"
                label="对错"
              >
              </el-table-column>
              <el-table-column
                prop="Point"
                width="100"
                align="center"
                label="得分"
              >
              </el-table-column>
            </div>
          </el-table>
        </div>
      </el-main>
    </el-main>
    <div class="footer"></div>
  </div>
</template>

<script>
export default {
  data: {
    //练习题数据
    expressions: [
      {
        Expression: "11+11*6-11",
        Answer: "66",
      },
      {
        Expression: "11+6+3+10",
        Answer: "30",
      },
      {
        Expression: "18*7/1+20",
        Answer: "146",
      },
      {
        Expression: "1+3+10/2",
        Answer: "9",
      },
      {
        Expression: "15-5-1+12",
        Answer: "21",
      },
      {
        Expression: "15*10+11+17",
        Answer: "178",
      },
      {
        Expression: "12+7*4-9",
        Answer: "31",
      },
      {
        Expression: "13*17-12*6",
        Answer: "149",
      },
      {
        Expression: "19+6+12-14",
        Answer: "23",
      },
      {
        Expression: "6*3/6+15",
        Answer: "18",
      },
    ],
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