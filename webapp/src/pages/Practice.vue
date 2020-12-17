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
        <div class="testname">练习题</div>
      </div>
      <div class="puzzles">
        <el-table :data="practiceData" style="width: 100%">
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
  </div>
</template>

<script>
export default {
  props: [],
  data() {
    return {
      show: false,
      lastPoint: 100,
      //练习题数据
      practiceData: [],
      tmppracticeData: [
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
          Expression: "8*7+27",
          Answer: "83",
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
    submitAnswer() {
      //提交练习
      //this.$router.push("/commit");
      var tablelength = this.practiceData.length;
      var rightNum = 0;
      for (var i = 0; i < tablelength; i++) {
        this.compareAnswer(i);
        if (this.practiceData[i].Judge === "√") rightNum++;
      }
      this.lastPoint = (this.lastPoint * rightNum) / tablelength;
      //展现结果
      this.ColShow();
      //传送数据
      this.$http({
        url: "/commit",
        method: "post",
        data: this.practiceData,
      }).then((res) => {
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
      /*const res = await this.$http.get("/practice");
      if (res.status !== 200)
        return this.$message.error("获取题目数据时出现错误!");
      console.log(res);
      //放入practiceData
      this.expressions = res.data;*/
      this.$http.get("/practice").then((res) => {
        console.log(res);
        if (res.status !== 200)
          return this.$message.error("获取题目数据时出现错误!");
        this.practiceData = res.data.expressions;
      });
      //渲染表格
      consol.log("practice data");
      console.log(this.practiceData);
    },
    compareAnswer(index) {
      console.log("compare Answer");
      var tablelength = this.practiceData.length;
      let tmpobj = this.practiceData[index];
      let ans = tmpobj.Answer;
      let res = tmpobj.Input;
      if (ans === res) {
        tmpobj.Judge = "√";
        tmpobj.Point = (100 / tablelength).toFixed(1);
        this.$set(this.practiceData, index, tmpobj);
      } else {
        tmpobj.Judge = "×";
        tmpobj.Point = 0;
        this.$set(this.practiceData, index, tmpobj);
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