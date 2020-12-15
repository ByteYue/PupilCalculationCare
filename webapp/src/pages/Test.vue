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
    <el-main class="main">
      <div class="pheader">
        <el-button @click="start()">开始答题</el-button>
        <!-- 计时器 -->
        <div class="timer">
          剩余时间：
          <div id="testTime">00:00</div>
        </div>
        <div class="testname">
          <div id="testName">测试卷一</div>
        </div>
        <div class="testid">
          id:
          <div id="testId"></div>
        </div>
      </div>
      <!-- 测试题 -->
      <div class="puzzles">
        <el-table
          id="puzzleTable"
          :data="puzzleData"
          style="width: 100%"
          :row-class-name="puzzleRow"
        >
          <el-table-column prop="id" width="100" label="序号">
          </el-table-column>
          <el-table-column prop="expression" width="300" label="题目">
          </el-table-column>
          <el-table-column prop="input" label="输入"> </el-table-column>
          <el-table-column prop="answer" label="答案"> </el-table-column>
          <el-table-column prop="judge" label="对错"> </el-table-column>
          <el-table-column width="400px"> </el-table-column>
        </el-table>
      </div>
    </el-main>
    <el-footer class="footer">
      <el-button @click="submitAnswer()">提交</el-button>
    </el-footer>
  </div>
</template>

<script>
export default {
  data() {
    return {
      testTime: {
        min: 5,
        sec: 0,
        now: "",
        end: "",
      },
    };
  },
  methods: {
    start() {
      //定义结束时间戳
      //我这里300000是5分钟倒计时，在当前时间戳的基础上加了300000毫秒
      this.end = new Date().getTime + 300000;
      //调用倒计时方法
      this.backDate();
      //使用get的数据渲染表格
      var xhr = new XMLHttpRequest();
      xhr.open("GET", "", true);
      xhr.onload = function () {
        if ((xhr.status >= 200 && xhr.status < 300) || xhr.status == 304) {
          var data = xhr.responseText;
          var p = JSON.parse(data);
          var table_node = document.getElementById("table");
          // 创建 tbody节点
          var tbody_node = document.createElement("tbody");
          // 将 tbody 节点插入到表格中
          table_node.appendChild(tbody_node);
          for (var key in person) {
            // 循环插入所有的行到 tbody 中
            tbody_node.insertRow(key);
            person[key].gender = person[key].gender == 1 ? "男" : "女";
            for (var attribute in person[key]) {
              //  在tbody中的行中插入一个单元格
              var td_node = tbody_node.rows[key].insertCell(-1);
              // 创建一个文本节点
              var text_node = document.createTextNode(person[key][attribute]);
              // console.log(attribute.gender);
              td_node.appendChild(text_node);
            }
          }
        } else {
          alert("Request was unsuccessful :" + xhr.status);
          console.log(xhr.statusText);
        }
      };
    },
    backToCreate() {
      //页面跳转
      this.$router.push("/login/create");
    },
    goToUser() {
      //页面跳转
      this.$router.push("/login/user");
    },
    backDate() {
      // 获取当前时间
      this.now = new Date().getTime();
      if (this.now >= this.end) {
        this.min = 0;
        this.sec = 0;
        this.submitAnswer();
        return;
      }
      // 用结束时间戳减去当前时间
      const msec = this.end - this.now;
      let min = parseInt(msec / 1000 / 60); //算出分钟数
      let sec = parseInt((msec / 1000) % 60); //算出秒数
      // 给数据赋值
      this.min = min > 9 ? min : "0" + min;
      this.sec = sec > 9 ? sec : "0" + sec;

      // 使用定时器 然后使用递归 让每一次函数能调用自己达到倒计时效果
      setTimeout(function () {
        this.BackDate();
      }, 1000);
    },
    submitAnswer() {
      this.$router.push("/login/create/practice/testcommit");
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

.pheader .timer {
  position: absolute;
  left: 30%;
  color: red;
}

.pheader .testname {
  font-size: 4ex;
  display: flex;
  position: absolute;
  left: 50%;
  width: 100%;
}

.pheader .testid {
  margin-left: 100px;
  position: absolute;
  left: 70%;
  top: 100px;
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