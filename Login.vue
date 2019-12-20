<template>
    <div>
        <card class="form-container" width="400px" height="320px">

            <div>
                <div>登录</div>
                <div>
                    <el-input class="login-input" placeholder="用户名" v-model="username"></el-input>
                    <el-input type="password" class="login-input" placeholder="密码" v-model="password"></el-input>
                    <el-link :underline="false" class="login-link">忘记密码</el-link>
                    <el-button class="login-button" type="primary" @click="login">注册/登录</el-button>

                </div>
            </div>

        </card>
    </div>
</template>

<script>
  //import Card from "@/components/Card";
  import axios from "axios";

  export default {
    name: "Login",
    //components: {Card},
    data() {
      return {
        username: "",
        password: "",
      }
    },
    methods: {
      login() {
        axios.post("http://localhost:8083/login", {
          username: this.username,
          password: this.password,
        }).then(res => {
          if (res.data.data == "err") {
            this.$message.error("登录失败")
          } else {
            this.$message.success("登录成功")
          }
        })

      },
    },
  }
</script>

