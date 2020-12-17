<template>
  <div class="createBox">
    <div class="particles">
      <vue-particles
        color="#555"
        :particleOpacity="0.2"
        :particlesNumber="100"
        shapeType="circle"
        :particleSize="3"
        linesColor="#555"
        :linesWidth="1"
        :lineLinked="true"
        :lineOpacity="0.1"
      />
    </div>
    <div class="formBox">
      <!-- 练习题表单 -->
      <div class="form1">
        <el-form
          ref="createPracticeFormRef"
          :model="createPracticeFormData"
          :rules="createPracticeRules"
          size="medium"
          label-width="100px"
          label-position="left"
        >
          <el-col :span="12">
            <el-row>
              <el-col :span="24">
                <el-form-item label="出现符号:" prop="symbol">
                  <el-checkbox-group
                    v-model="createPracticeFormData.symbol"
                    :min="1"
                    :max="4"
                    size="medium"
                  >
                    <el-checkbox
                      v-for="(item, index) in symbolOptions"
                      :key="index"
                      :label="item.value"
                      :disabled="item.disabled"
                      >{{ item.label }}</el-checkbox
                    >
                  </el-checkbox-group>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item
                  label-width="110px"
                  label="操作数个数:"
                  prop="dnum"
                >
                  <el-select
                    v-model="createPracticeFormData.dnum"
                    placeholder="请选择:"
                    clearable
                    :style="{ width: '100%' }"
                  >
                    <el-option
                      v-for="(item, index) in dnumOptions"
                      :key="index"
                      :label="item.label"
                      :value="item.value"
                      :disabled="item.disabled"
                    ></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item
                  label-width="106px"
                  label="操作数范围:"
                  prop="dsize"
                >
                  <el-select
                    v-model="createPracticeFormData.dsize"
                    placeholder="请选择:"
                    clearable
                    :style="{ width: '100%' }"
                  >
                    <el-option
                      v-for="(item, index) in dsizeOptions"
                      :key="index"
                      :label="item.label"
                      :value="item.value"
                      :disabled="item.disabled"
                    ></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item
                  label-width="110px"
                  label="题目数量:"
                  prop="puznum"
                >
                  <el-select
                    v-model="createPracticeFormData.puznum"
                    placeholder="请选择:题目数量"
                    clearable
                    :style="{ width: '100%' }"
                  >
                    <el-option
                      v-for="(item, index) in puznumOptions"
                      :key="index"
                      :label="item.label"
                      :value="item.value"
                      :disabled="item.disabled"
                    ></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item
                  label-width="110px"
                  label="答案范围:"
                  prop="anssize"
                >
                  <el-select
                    v-model="createPracticeFormData.anssize"
                    placeholder="请选择:"
                    clearable
                    :style="{ width: '100%' }"
                  >
                    <el-option
                      v-for="(item, index) in anssizeOptions"
                      :key="index"
                      :label="item.label"
                      :value="item.value"
                      :disabled="item.disabled"
                    ></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
          </el-col>
        </el-form>
        <el-button
          type="primary"
          icon="el-icon-reading"
          size="medium"
          style="position: relative"
          @click="CreatePractice()"
        >
          确认生成练习题
        </el-button>
      </div>
      <div class="form2">
        <!-- 测试题表单 -->
        <el-form
          ref="createTestFormRef"
          :model="createTestFormData"
          :rules="createTestRules"
          size="medium"
          label-width="100px"
          label-position="left"
        >
          <el-col :span="12">
            <el-row>
              <el-col :span="20">
                <el-form-item label="出现符号:" prop="symbols">
                  <el-checkbox-group
                    v-model="createTestFormData.symbols"
                    :min="1"
                    :max="4"
                    size="medium"
                  >
                    <el-checkbox
                      v-for="(item, index) in symbolsOptions"
                      :key="index"
                      :label="item.value"
                      :disabled="item.disabled"
                      >{{ item.label }}</el-checkbox
                    >
                  </el-checkbox-group>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item
                  label-width="110px"
                  label="操作数个数:"
                  prop="dnum"
                >
                  <el-select
                    v-model="createTestFormData.dnum"
                    placeholder="请选择:"
                    clearable
                    :style="{ width: '100%' }"
                  >
                    <el-option
                      v-for="(item, index) in dnumOptions"
                      :key="index"
                      :label="item.label"
                      :value="item.value"
                      :disabled="item.disabled"
                    ></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item
                  label-width="106px"
                  label="操作数范围:"
                  prop="dsize"
                >
                  <el-select
                    v-model="createTestFormData.dsize"
                    placeholder="请选择:"
                    clearable
                    :style="{ width: '100%' }"
                  >
                    <el-option
                      v-for="(item, index) in dsizeOptions"
                      :key="index"
                      :label="item.label"
                      :value="item.value"
                      :disabled="item.disabled"
                    ></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item
                  label-width="110px"
                  label="题目数量:"
                  prop="puznum"
                >
                  <el-select
                    v-model="createTestFormData.puznum"
                    placeholder="请选择:"
                    clearable
                    :style="{ width: '100%' }"
                  >
                    <el-option
                      v-for="(item, index) in puznumOptions"
                      :key="index"
                      :label="item.label"
                      :value="item.value"
                      :disabled="item.disabled"
                    ></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item
                  label-width="110px"
                  label="答案范围:"
                  prop="anssize"
                >
                  <el-select
                    v-model="createTestFormData.anssize"
                    placeholder="请选择:"
                    clearable
                    :style="{ width: '100%' }"
                  >
                    <el-option
                      v-for="(item, index) in anssizeOptions"
                      :key="index"
                      :label="item.label"
                      :value="item.value"
                      :disabled="item.disabled"
                    ></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="测试时长:" prop="testtime">
                  <el-select
                    v-model="createTestFormData.testtime"
                    placeholder="请选择测试时长:"
                    clearable
                    :style="{ width: '100%' }"
                  >
                    <el-option
                      v-for="(item, index) in testtimeOptions"
                      :key="index"
                      :label="item.label"
                      :value="item.value"
                      :disabled="item.disabled"
                    ></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="测试卷名:" prop="testname">
                  <el-input
                    v-model="createTestFormData.testname"
                    placeholder="请输入卷名"
                    clearable
                    :style="{ width: '100%' }"
                  >
                  </el-input>
                </el-form-item>
              </el-col>
            </el-row>
          </el-col>
        </el-form>
        <el-button
          type="primary"
          icon="el-icon-receiving"
          size="medium"
          style="position: relative"
          @click="CreateTest()"
        >
          确认生成测试题
        </el-button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  components: {},
  props: [],
  data() {
    return {
      //练习题表单数据
      createPracticeFormData: {
        symbol: [1],
        dnum: 2,
        dsize: 10,
        puznum: 10,
        anssize: 10,
      },
      createPracticeRules: {
        symbol: [
          {
            required: true,
            type: "array",
            message: "请至少选择一个出现符号:",
            trigger: "change",
          },
        ],
        dnum: [
          {
            required: true,
            message: "请选择:",
            trigger: "change",
          },
        ],
        dsize: [
          {
            required: true,
            message: "请选择:",
            trigger: "change",
          },
        ],
        puznum: [
          {
            required: true,
            message: "请选择:题目数量",
            trigger: "change",
          },
        ],
        anssize: [
          {
            required: true,
            message: "请选择:",
            trigger: "change",
          },
        ],
      },
      symbolOptions: [
        {
          label: "+",
          value: 1,
        },
        {
          label: "-",
          value: 2,
        },
        {
          label: "×",
          value: 3,
        },
        {
          label: "÷",
          value: 4,
        },
      ],
      dnumOptions: [
        {
          label: "2",
          value: 2,
        },
        {
          label: "3",
          value: 3,
        },
        {
          label: "4",
          value: 4,
        },
      ],
      dsizeOptions: [
        {
          label: "0~10",
          value: 10,
        },
        {
          label: "0~20",
          value: 20,
        },
        {
          label: "0~100",
          value: 100,
        },
      ],
      puznumOptions: [
        {
          label: "10",
          value: 10,
        },
        {
          label: "20",
          value: 20,
        },
        {
          label: "50",
          value: 50,
        },
        {
          label: "100",
          value: 100,
        },
        {
          label: "300",
          value: 300,
        },
      ],
      anssizeOptions: [
        {
          label: "0~10",
          value: 10,
        },
        {
          label: "0~20",
          value: 20,
        },
        {
          label: "0~100",
          value: 100,
        },
        {
          label: "0~1000",
          value: 1000,
        },
      ],
      //测试卷卷表单数据
      createTestFormData: {
        symbols: [1],
        dnum: 2,
        dsize: 10,
        puznum: 10,
        anssize: 10,
        testtime: 5,
        testname: "测试卷",
      },
      createTestRules: {
        symbols: [
          {
            required: true,
            type: "array",
            message: "请至少选择一个出现符号:",
            trigger: "change",
          },
        ],
        dnum: [
          {
            required: true,
            message: "请选择:",
            trigger: "change",
          },
        ],
        dsize: [
          {
            required: true,
            message: "请选择:",
            trigger: "change",
          },
        ],
        puznum: [
          {
            required: true,
            message: "请选择:",
            trigger: "change",
          },
        ],
        anssize: [
          {
            required: true,
            message: "请选择:",
            trigger: "change",
          },
        ],
        testtime: [
          {
            required: true,
            message: "请选择测试时长:",
            trigger: "change",
          },
        ],
        testname: [
          {
            required: true,
            message: "请输入卷名",
            trigger: "blur",
          },
        ],
      },
      symbolsOptions: [
        {
          label: "+",
          value: 1,
        },
        {
          label: "-",
          value: 2,
        },
        {
          label: "×",
          value: 3,
        },
        {
          label: "÷",
          value: 4,
        },
      ],
      dnumOptions: [
        {
          label: "2",
          value: 2,
        },
        {
          label: "3",
          value: 3,
        },
        {
          label: "4",
          value: 4,
        },
      ],
      dsizeOptions: [
        {
          label: "0~10",
          value: 10,
        },
        {
          label: "0~20",
          value: 20,
        },
        {
          label: "0~100",
          value: 100,
        },
      ],
      puznumOptions: [
        {
          label: "10",
          value: 10,
        },
        {
          label: "20",
          value: 20,
        },
        {
          label: "50",
          value: 50,
        },
        {
          label: "100",
          value: 100,
        },
        {
          label: "300",
          value: 300,
        },
      ],
      anssizeOptions: [
        {
          label: "0~10",
          value: 10,
        },
        {
          label: "0~20",
          value: 20,
        },
        {
          label: "0~100",
          value: 100,
        },
        {
          label: "0~1000",
          value: 1000,
        },
      ],
      testtimeOptions: [
        {
          label: "5",
          value: 5,
        },
        {
          label: "10",
          value: 10,
        },
        {
          label: "20",
          value: 20,
        },
        {
          label: "30",
          value: 30,
        },
        {
          label: "45",
          value: 45,
        },
        {
          label: "60",
          value: 60,
        },
      ],
    };
  },
  computed: {},
  watch: {},
  created() {},
  mounted() {},
  methods: {
    CreatePractice() {
      this.$refs.createPracticeFormRef.validate(async (valid) => {
        if (!valid) return;
        //提交表单,这里需要http
        const res = await this.$http.post(
          "/practice",
          this.$qs.stringify(this.createPracticeFormData)
        );
        if (res.status !== 200) return alert("生成错误！");

        this.$message.success("生成练习题...");
        console.log(res);
        //页面跳转
        this.$router.push("/practice");
      });
    },
    CreateTest() {
      this.$refs.createTestFormRef.validate(async (valid) => {
        if (!valid) return;
        //提交表单,这里需要http
        const res = await this.$http.post(
          "/test",
          this.$qs.stringify(this.createTestFormData)
        );
        if (res.status !== 200) return alert("生成错误！");
        //弹出信息
        this.$message.success("生成测试题...");
        window.sessionStorage.setItem(
          "testtime",
          this.createTestFormData.testtime
        );
        window.sessionStorage.setItem(
          "testname",
          this.createTestFormData.testname
        );
        //页面跳转
        this.$router.push("/test");
      });
    },
    submitForm() {
      this.$refs["createPuzzleFormRef"].validate((valid) => {
        if (!valid) return;
        // TODO 提交表单
      });
    },
    resetForm() {
      this.$refs["createPuzzleFormRef"].resetFields();
    },
  },
};
</script>

<style scoped>
.createBox {
  width: 100%;
}
.createBox .particles {
  height: 0;
}
.formBox {
  display: flex;
  height: 100%;
}

.formBox .form1 {
  position: absolute;
  margin-top: 5%;
  margin-left: 15%;
  width: 45%;
}

.formBox .form2 {
  position: absolute;
  margin-top: 5%;
  margin-left: 60%;
  width: 45%;
}

.formBox button {
  width: 53%;
  cursor: pointer;
}
</style>
