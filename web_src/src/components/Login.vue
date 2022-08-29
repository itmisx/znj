<template>
  <div class="main">
    <el-card v-loading="loading" class="box-card">
      <div
        style="
          text-align: center;
          margin-top: 30px;
          color: #25b864;
          font-size: 32px;
        "
      >
        <i class="iconfont icon-edit" style="font-size: 36px"
          ><span style="font-size: 32px">{{ productName }}</span></i
        >
      </div>
      <div
        style="
          text-align: center;
          margin-top: 10px;
          font-size: 16px;
          color: #999999;
        "
      >
        <span>记笔记更简单</span>
      </div>
      <el-form style="margin: 10px 40px 5px">
        <el-input
          placeholder="账号"
          v-model="user"
          prefix-icon="iconfont icon-user"
          size="small"
          @keyup.enter.native="login"
          class="form-input"
        >
        </el-input>
        <el-input
          placeholder="密码"
          v-model="password"
          show-password
          prefix-icon="iconfont icon-key"
          size="small"
          @keyup.enter.native="login"
          class="form-input"
        >
          <el-button slot="append" style="width: 80px" type="text"
            >忘记密码?</el-button
          >
        </el-input>
        <el-row>
          <el-button
            type="primary"
            style="width: 100%"
            size="small"
            @click="login"
            class="form-input"
            >登录</el-button
          >
        </el-row>
      </el-form>
    </el-card>
    <div class="bottom">
      <span>©znj</span>
    </div>
  </div>
</template>

<script>
import * as productConfig from "../config/product";
export default {
  name: "Login",
  data() {
    return {
      user: "",
      password: "",
      loading: false,
      productName: productConfig.name,
    };
  },
  methods: {
    login() {
      let data = {
        user: this.user,
        password: this.password,
      };
      this.loading = true;
      this.Request.signin(data)
        .then((response) => {
          if (response.data.code === 0) {
            /*保存token*/

            localStorage.setItem("Authorization", response.data.data.token);
            this.$router.push({
              path: "/main",
            });
          } else {
            this.$message(response.data.msg);
          }
          this.loading = false;
        })
        .catch((error) => {
          this.loading = false;
        });
    },
  },
  mounted() {
    this.Request.isSignedIn().then((response) => {
      if (response.data.code === 0) {
        this.$router.push({
          path: "/main/todo-list",
        });
      }
    });
  },
};
</script>

<style scoped>
.main {
  height: 100%;
  width: 100%;
}
.box-card {
  height: 430px;
  width: 430px;
  margin: 100px auto;
}
.form-input {
  margin-top: 10px;
}
.bottom {
  width: 100%;
  height: 30px;
  position: absolute;
  bottom: 0px;
  text-align: center;
  font-size: 12px;
}
</style>
