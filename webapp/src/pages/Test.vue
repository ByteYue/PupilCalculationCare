<template>
  <div class="practicePage">
    <el-header class="header">
      <el-popconfirm
        icon="el-icon-info"
        @confirm="backToCreate()"
        title="注意：返回主页后会清空当前进度！"
      >
        <el-button slot="reference">返回主页</el-button>
      </el-popconfirm>
      <div class="user">
        <img src="../assets/img/头像.jpg" alt="头像" @click="goToUser()" />
      </div>
    </el-header>
    <!--练习题目-->
    <el-main class="main">
      <div class="pheader">
        <div class="timer">{{ min }}:{{ sec }}</div>
        <div class="testname">测试卷</div>
      </div>
      <div class="puzzles">
        <el-table :data="testData" style="width: 100%">
          <el-table-column prop="id" width="200" align="right" label="序号">
            <template slot-scope="scope">
              <span v-if="scope.row.id">{{ scope.row.id }}</span>
            </template>
          </el-table-column>
          <el-table-column
            prop="Expression"
            width="200"
            align="center"
            label="题目"
          >
          </el-table-column>
          <el-table-column prop="Input" width="100" align="center" label="输入">
            <template slot-scope="scope">
              <el-input
                v-model="scope.row.Input"
                :data-price="scope.row.result"
                type="text"
              ></el-input>
            </template>
          </el-table-column>
          <el-table-column
            prop="Answer"
            width="200"
            align="center"
            label="答案"
            v-if="show"
          >
          </el-table-column>
          <el-table-column
            prop="Judge"
            width="200"
            align="center"
            label="对错"
            v-if="show"
          >
          </el-table-column>
          <el-table-column
            prop="Point"
            width="100"
            align="center"
            label="得分"
            v-if="show"
          >
          </el-table-column>
        </el-table>
      </div>
    </el-main>
    <el-footer class="footer">
      <el-button @click="submitAnswer()">提交</el-button>
    </el-footer>
    <div class="points" v-if="show">
      <img src="../assets/img/成绩单.png" alt="" />
      <div class="word">
        <div v-if="lastPoint === 100">恭喜你拿了满分！你真是个小天才！</div>
        <div v-else-if="lastPoint > 90">你做的非常棒！小心不要粗心哦~</div>
        <div v-else-if="lastPoint > 80">还差一点点，加油！</div>
        <div v-else-if="lastPoint > 60">多多练习，你会更棒的！</div>
        <div v-else>脚踏实地，选择低一点的难度重新开始吧！</div>
      </div>
      <div class="score">
        {{ lastPoint }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: [],
  data() {
    return {
      show: false,
      lastPoint: 100,
      end: "",
      now: "",
      min: 5,
      sec: 0,
      //练习题数据
      testData: [],
      tmptestData: [
        {
          id: 1,
          Expression: "2+8+8+20",
          Answer: "38",
          Judge: "",
          Point: 0,
        },
        {
          id: 2,
          Expression: "3+15+18+5",
          Answer: "41",
          Judge: "",
          Point: 0,
        },
        {
          id: 3,
          Expression: "3+15+18+5",
          Answer: "41",
          Judge: "",
          Point: 0,
        },
        {
          id: 4,
          Expression: "3+15+18+5",
          Answer: "41",
          Judge: "",
          Point: 0,
        },
        {
          id: 5,
          Expression: "3+15+18+5",
          Answer: "41",
          Judge: "",
          Point: 0,
        },
        {
          id: 6,
          Expression: "3+15+18+5",
          Answer: "41",
          Judge: "",
          Point: 0,
        },
        {
          id: 7,
          Expression: "3+15+18+5",
          Answer: "41",
          Judge: "",
          Point: 0,
        },
        {
          id: 8,
          Expression: "3+15+18+5",
          Answer: "41",
          Judge: "",
          Point: 0,
        },
        {
          id: 9,
          Expression: "3+15+18+5",
          Answer: "41",
          Judge: "",
          Point: 0,
        },
        {
          id: 10,
          Expression: "3+15+18+5",
          Answer: "41",
          Judge: "",
          Point: 0,
        },
      ],
    };
  },
  mounted() {
    this.start();
    this.timerStart();
  },
  methods: {
    timerStart() {
      this.min = JSON.parse(sessionStorage.getItem("testtime"));
      //this.testName = JSON.parse(sessionStorage.getItem("testname"));
      // 当前时间戳
      this.now = Date.parse(new Date());
      // 目标日期时间戳
      this.end = this.now + this.min * 60 * 1000 + this.sec * 1000;
      this.countTime();
      return;
    },
    countTime() {
      if (this.min === 0 && this.sec === 0) {
        this.submitAnswer();
        return;
      }
      // 当前时间戳
      this.now = Date.parse(new Date());
      // 相差的毫秒数
      var msec = this.end - this.now;
      // 计算分秒数
      this.min = parseInt((msec / 60 / 1000) % 60);
      this.sec = parseInt((msec / 1000) % 60);
      // 个位数前补零
      this.min = this.min > 9 ? this.min : "0" + this.min;
      this.sec = this.sec > 9 ? this.sec : "0" + this.sec;
      // 控制台打印
      //console.log(`${this.min}分钟 ${this.sec}秒`);
      // **错误记录：一秒后递归(这里的this指向指的windows)
      setTimeout(() => {
        this.countTime();
      }, 1000);
    },
    backToCreate() {
      //页面跳转
      this.$router.push("/create");
    },
    goToUser() {
      //页面跳转
      this.$router.push("/user");
    },
    submitAnswer() {
      //提交练习
      //this.$router.push("/commit");
      var tablelength = this.testData.length;
      var rightNum = 0;
      for (var i = 0; i < tablelength; i++) {
        this.compareAnswer(i);
        if (this.testData[i].Judge === "√") rightNum++;
      }
      this.lastPoint = ((this.lastPoint * rightNum) / tablelength).toFixed(1);
      //展现成绩单
      this.ColShow();
      //传送数据
      this.$http.post("/commit", this.testData).then((res) => {
        console.log(res.data);
        if (res.meta.status !== 200) return alert("提交数据错误！");
        //弹出信息
        this.$message.success("提交成功");
      });
    },
    ColShow() {
      this.show = true;
    },
    start() {
      //async start() {
      //用get获取数据
      /*const res = await this.$http.get("/test");
      if (res.status !== 200)
        return this.$message.error("获取题目数据时出现错误!");
      console.log(res);
      //放入practiceData
      this.testData = res.data;*/
      this.$http.get("/test").then((res) => {
        console.log(res);
        if (res.status !== 200)
          return this.$message.error("获取题目数据时出现错误!");
        this.testData = res.data.expressions;
      });
      //渲染表格
      console.log("test data");
      console.log(this.testData);
    },
    compareAnswer(index) {
      console.log("compare Answer");
      var tablelength = this.testData.length;
      let tmpobj = this.testData[index];
      let ans = tmpobj.Answer;
      let res = tmpobj.Input;
      if (ans === res) {
        tmpobj.Judge = "√";
        tmpobj.Point = (100 / tablelength).toFixed(1);
        this.$set(this.testData, index, tmpobj);
      } else {
        tmpobj.Judge = "×";
        tmpobj.Point = 0;
        this.$set(this.testData, index, tmpobj);
      }
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
.main .pheader {
  display: flex;
  width: 200px;
  color: black;
}

.pheader .testname {
  font-size: 30px;
  display: flex;
  position: absolute;
  left: 30rem;
  top: 60px;
  width: 100%;
}

.main .puzzles {
  margin-top: 30px;
  margin-left: 0;
  color: black;
  width: 100%;
}
.main .puzzle {
  width: 100%;
  display: flex;
}

.timer {
  position: absolute;
  z-index: 1;
  color: red;
  font-size: 20px;
}
.puzzles textarea {
  width: 100px;
}

.footer {
  background-color: rgb(196, 212, 230);
  height: 80px;
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
}

.points {
  background-color: rgb(255, 255, 255);
  color: #000;
  border: 1px solid rgb(51, 48, 48);
  height: 150px;
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  background-image: url("../assets/img/成绩单背景.jpg");
  background-repeat: no-repeat;
  background-attachment: fixed;
}

.points .word {
  position: absolute;
  font-size: 18px;
  font-family: "Franklin Gothic Medium", "Arial Narrow", Arial, sans-serif;
  left: 50%;
  top: 60%;
  transform: translate(-50%, -50%);
}

.points img {
  position: absolute;
  width: 280px;
  height: 100%;
  left: 10%;
  top: 50%;
  transform: translate(-50%, -50%);
}

.points .score {
  color: red;
  font-style: oblique;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
    Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
  font-size: 40px;
  position: absolute;
  left: 50%;
  transform: translate(-50%);
}

.footer button {
  position: absolute;
  right: 10px;
  bottom: 10px;
}
</style>