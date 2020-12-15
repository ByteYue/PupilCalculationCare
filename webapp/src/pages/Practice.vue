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
        <el-button @click="start()">开始答题</el-button>
        <div class="testname">
          <div id="testName">练习题</div>
        </div>
      </div>
      <div class="puzzles">
        <el-table
          :data="puzzleData"
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
          <el-table-column prop="Input" width="100" align="center" label="输入">
            <el-input></el-input
          ></el-table-column>
          <el-table-column
            prop="Answer"
            width="200"
            align="center"
            label="答案"
          >
          </el-table-column>
          <el-table-column prop="Judge" width="200" align="center" label="对错">
          </el-table-column>
          <el-table-column prop="Point" width="100" align="center" label="得分">
          </el-table-column>
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
      puzzleData: [
        {
          id: 1,
          Expression: "400-5+8",
          Answer: "403",
          Judge: "√",
          Point: 2,
        },
        {
          id: 2,
          Expression: "40*7+9",
          Answer: "289",
          Judge: "X",
          Point: 0,
        },
      ],
    };
  },
  methods: {
    backToCreate() {
      //页面跳转
      this.$router.push("/login/create");
    },
    goToUser() {
      //页面跳转
      this.$router.push("/login/user");
    },
    submitAnswer() {
      //提交练习
      this.$router.push("/login/create/practice/practicecommit");
    },
    start() {
      //使用get的数据渲染表格
      /*
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
        }*/
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
  font-size: 4ex;
  display: flex;
  position: absolute;
  left: 50%;
  width: 100%;
}
.main #puzzles {
  color: black;
  width: 100%;
}
.main .puzzle {
  margin: 5px;
  width: 100%;
  display: flex;
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

.footer button {
  position: absolute;
  right: 10px;
  bottom: 10px;
}
</style>