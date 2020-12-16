<template>
  <div class="loginPage">
    <vue-particles
      color="#fff"
      :particleOpacity="0.7"
      :particlesNumber="60"
      shapeType="circle"
      :particleSize="4"
      linesColor="#fff"
      :linesWidth="1"
      :lineLinked="true"
      :lineOpacity="0.4"
      :linesDistance="150"
      :moveSpeed="2"
      :hoverEffect="true"
      hoverMode="grab"
      :clickEffect="true"
      clickMode="push"
    />
    <div class="login_box">
      <div class="logo_box">
        <img src="../assets/img/logo.png" alt="logo" />
      </div>
      <div class="slogen_box">
        <img src="../assets/img/slogen.png" alt="" />
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
            <el-form-item label="账号" prop="owner">
              <el-input
                v-model="loginForm.owner"
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
                    @click="Login()"
                  >
                    登陆
                  </el-button>
                </el-form-item>
              </el-col>
              <el-col :span="7">
                <el-form-item label-width="0" prop="btnreset">
                  <el-button
                    type="info"
                    size="medium"
                    @click="resetLoginForm()"
                  >
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
        owner: "",
        password: "",
      },
      loginFormRules: {
        owner: [
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
      this.$router.push("/create");
    },
    Login() {
      this.$refs.loginFormRef.validate(async (valid) => {
        if (!valid) return;
        //提交表单,这里需要http
        //const result = this.$http.post("/create", this.loginForm);
        //console.log(result);
        const { data: res } = await this.$http.post("/login", this.loginForm);
        console.log(res);

        if (res.meta.status !== 200) return alert("登录失败！");
        this.$message.success("登陆成功！");

        //保存token和owner
        window.sessionStorage.setItem("token", res.data.token);
        window.sessionStorage.setItem("owner", this.loginForm.owner);
        //页面跳转
        this.$router.push("/create");
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
  background-image: url("../assets/img/login背景.jpg");
  background-repeat: no-repeat;
  background-attachment: scroll;
  background-size: 100% 100%;
}
.login_box {
  width: 600px;
  height: 300px;
  position: absolute;
  left: 53%;
  top: 50%;
  transform: translate(-50%, -50%);
}

.slogen_box img {
  width: 330px;
  height: 100px;
  position: absolute;
  top: -40%;
  left: 33%;
}

.logo_box {
  height: 100px;
  width: 100px;
  position: absolute;
  top: -65%;
  left: 0;
}
</style>
