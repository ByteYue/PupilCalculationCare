<template>
  <div class="loginPage">
    <vue-particles
      v-show="!(path == '/User')"
      color="#555"
      :particleOpacity="0.5"
      :particlesNumber="100"
      shapeType="circle"
      :particleSize="4"
      linesColor="#555"
      :linesWidth="1"
      :lineLinked="true"
      :lineOpacity="0.4"
    />
    <div class="login_box">
      <div class="logo_box">
        <img src="../assets/img/logo.png" alt="logo" />
      </div>
      <el-row :gutter="1">
        <el-form
          ref="loginFormRef"
          :model="loginForm"
          :rules="loginFormRules"
          size="medium"
          label-width="100px"
        >
          <el-col :span="15">
            <el-form-item label="账号" prop="account">
              <el-input
                v-model="loginForm.account"
                placeholder="请输入账号"
                clearable
                :style="{ width: '100%' }"
              >
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="15">
            <el-form-item label="密码" prop="password">
              <el-input
                v-model="loginForm.password"
                placeholder="请输入密码"
                clearable
                show-password
                :style="{ width: '100%' }"
              ></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="20">
            <el-row type="flex" justify="center" align="middle">
              <el-col :span="7">
                <el-form-item label-width="0" prop="btnlogin">
                  <el-button
                    type="primary"
                    icon="el-icon-sunny"
                    size="medium"
                    @click="tmpLogin"
                  >
                    登陆
                  </el-button>
                </el-form-item>
              </el-col>
              <el-col :span="7">
                <el-form-item label-width="0" prop="btnreset">
                  <el-button type="info" size="medium" @click="resetLoginForm">
                    重置
                  </el-button>
                </el-form-item>
              </el-col>
            </el-row>
          </el-col>
        </el-form>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  components: {},
  props: [],
  data() {
    return {
      // 登陆表单的数据绑定对象
      loginForm: {
        account: "admin",
        password: "12345",
      },
      loginFormRules: {
        account: [
          {
            required: true,
            message: "请输入账号,长度为5~12个字符",
            trigger: "blur",
          },
          {
            min: 5,
            max: 12,
          },
        ],
        password: [
          {
            required: true,
            message: "请输入密码,长度为5~12个字符",
            trigger: "blur",
          },
          {
            min: 5,
            max: 12,
          },
        ],
      },
    };
  },
  computed: {},
  watch: {},
  created() {},
  mounted() {},
  methods: {
    tmpLogin() {
      //弹出信息
      this.$message.success("登陆成功！");
      //console.log(this.loginForm);
      //页面跳转
      this.$router.push("/login/create");
    },
    submitLoginForm() {
      this.$refs.loginFormRef.validate(async (valid) => {
        if (!valid) return;
        //提交表单,这里需要http
        const { data: res } = await this.$http.post(
          "/api/xxx/xxx",
          this.loginForm
        );
        if (res.meta.status !== 200) return alert("登录失败！");
        this.$message.success("登陆成功！");
        console.log(res);
        //保存token
        window.sessionStorage.setItem("token", res.data.token);
        //页面跳转
        this.$router.push("/login/create");
      });
    },
    resetLoginForm() {
      this.$refs["loginFormRef"].resetFields();
    },
  },
};
</script>
<style scoped>
.loginPage {
  height: 100%;
}
.login_box {
  width: 600px;
  height: 300px;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}

.logo_box {
  height: 100px;
  width: 100px;
  position: absolute;
  top: -60%;
  left: 20%;
}
</style>
