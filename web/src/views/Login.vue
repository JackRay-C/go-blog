<template>
  <div class="login fadeInUp">
    <div class="liner"></div>
    <div class="login-form">
      <!-- 标题 -->
      <div class="login-title">
        <div class="login-title-text">我的博客登录</div>
      </div>

      <!-- 用户名 -->
      <div class="input-block">
        <div class="input-prefix">
          <!-- <svg
            viewBox="0 0 1024 1024"
            class
            width="1.1em"
            height="1.1em"
            fill="currentColor"
            aria-hidden="true"
            focusable="false"
          >
            <path
              d="M512 512c-140.8 0-256-115.2-256-256s115.2-256 256-256 256 115.2 256 256-115.2 256-256 256z m512 512H0c0-281.6 230.4-512 512-512s512 230.4 512 512z"
              p-id="3188"
              fill="#555555"
            />
          </svg> -->
          <svg-icon icon-class="username" class-name="" />
        </div>
        <input
          type="text"
          class="input"
          v-bind:class="{ error: uError }"
          v-model="username"
          placeholder="用户名或密码"
        />
        <div class="username-input-suffix"></div>
        <div class="error-text" v-if="uError">
          <p>{{ uError }}</p>
        </div>
      </div>
      <!-- 密码 -->
      <div class="input-block">
        <div class="input-prefix">
          <i aria-label="图标：password" class="search-icon">
            <svg-icon icon-class="lock" class-name="" />
          </i>
        </div>
        <input
          type="password"
          class="input"
          v-model="password"
          v-bind:class="{ error: pError }"
          placeholder="密码"
        />
        <div class="password-input-suffix"></div>
      </div>

      <!-- 登录按钮 -->
      <div class="submit-block">
        <div class="submit-button-group">
          <div class="submit-button-block">
            <button class="submit-button" @click="login">登录</button>
          </div>

          <div class="submit-button-block">
            <button class="submit-button register-button">注册</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import SvgIcon from "@/components/SvgIcon";
// import { webRoutes } from "../router";
// import {
//   getPermissionByAccessToken,
//   getRolesByAccessToken,
//   login,
// } from "../api/web/login";

export default {
  components: {
    SvgIcon,
  },
  name: "Login",
  data() {
    return {
      username: "",
      uError: "",
      password: "",
      pError: "",
    };
  },
  watch: {
    $route: {
      handler: function (route) {
        const query = route.query;
        if (query) {
          this.redirect = query.redirect;
          this.otherQuery = this.getOtherQuery(query);
        }
      },
      immediate: true,
    },
  },
  methods: {
    async login() {
      // 1、登录获取token，存储到localstorage
      let data = { username: this.username.trim(), password: this.password.trim() };
      this.$store.dispatch('DispatchLogin', data).then(()=> {
        this.$router.push({
        path: this.redirect || "/admin/dashboard",
        query: this.otherQuery,
      });
      }).catch(err => {
        this.$notify.error({ title: "失败", message: err });
        return
      })
      // let login_resp = await login(data);
      // if (login_resp.code !== 200) {
      //   this.$notify.error({ title: "失败", message: login_resp.message });
      //   return 
      // } 
      // let { access_token, refresh_token, expire } = login_resp.data;
      // localStorage.setItem("access_token", access_token);
      // localStorage.setItem("refresh_token", refresh_token);
      // localStorage.setItem("expire", expire);

      // // 2、通过accessToken获取角色列表
      // let roles_resp = await getRolesByAccessToken(access_token);
      // if (roles_resp.code !== 200) {
      //   this.$notify.error({ title: "失败", message: roles_resp.message });
      //   return 
      // }
      // localStorage.setItem("roles", JSON.stringify(roles_resp.data));

      // // 3、通过accessToken获取权限列表
      // let permissions_resp = await getPermissionByAccessToken(access_token);
      // if (permissions_resp.code !== 200) {
      //   this.$notify.error({title: "失败",message: permissions_resp.message});
      //   return
      // }
      // localStorage.setItem("permissions", JSON.stringify(permissions_resp.data));

      // 跳转首页或者redirect
      // this.$router.push({
      //   path: this.redirect || "/admin/dashboard",
      //   query: this.otherQuery,
      // });
    },
    getOtherQuery(query) {
      return Object.keys(query).reduce((acc, cur) => {
        if (cur !== "redirect") {
          acc[cur] = query[cur];
        }
        return acc;
      }, {});
    },
    goHome() {
      this.$route.push("/");
    },
  },
};
</script>

<style lang="scss" scoped>
.login {
  font-weight: 400;
  font-size: 16px;
  min-height: 573px;
  box-sizing: border-box;
  margin: 120px auto;
  background: #ffffff;
  width: 600px;
  padding: 32px 24px;
  box-shadow: 0 6px 28px 0 rgba(24, 52, 117, 0.2);

  border-top: 2px solid #4e6ef2;
}

.login-form {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  width: 80%;
  height: 100%;
  margin: auto;
}
.login-title {
  width: 100%;
  height: 50px;
  text-align: center;
  align-items: center;
  margin-top: 32px;
}

.login-title-text {
  font-size: 25px;
  line-height: 0;
}

.input-block {
  width: 100%;

  height: 50px;
  box-sizing: border-box;
  border-radius: 4px;
  color: #3e4c5b;
  position: relative;
  display: inline-block;
  font-variant: tabular-nums;
  list-style: none;
  text-align: start;
  box-sizing: border-box;
  &:not(:first-child) {
    margin-top: 32px;
  }

  .input-prefix {
    left: 18px;
    font-size: 16px;
    position: absolute;
    top: 50%;
    z-index: 2;
    line-height: 0;
    color: #9ca2a9;
    transform: translateY(-55%);
  }

  .input-suffix {
    top: 0;
    right: 0;
    bottom: 0;
    transform: none;
    font-size: 16px;
    position: absolute;
    z-index: 2;
    line-height: 0;
    color: #9ca2a9;
  }

  .input {
    padding-left: 52px;
    padding-right: 10px;
    min-height: 100%;
    position: relative;
    font-size: 18px;
    text-align: left;
    resize: none;
    box-sizing: border-box;
    position: relative;
    display: inline-block;
    width: 100%;
    color: #3e4c5b;
    border-radius: 4px;
    transition: all 0.3s;
    caret-color: #4e6ef2;
    overflow: visible;
    border: 1px solid #dddddd;

    &:hover,
    &:focus {
      outline: none;
      border: 1px solid #4e6ef2;
      border-right-width: 1px !important;
    }

    &::-webkit-input-placeholder {
      color: #aaaaaa;
    }
  }
  .error {
    border: 1px solid red;
  }

  .error-text {
    color: red;
    font-size: 15px;
    margin-top: 5px;
    margin-bottom: 0;
  }
}

.submit-block {
  width: 100%;
  margin-top: 80px;
}

.submit-button-group {
  width: 100%;
  display: flex;
  flex-direction: column;

  .submit-button-block {
    width: 100%;
    &:not(:first-child) {
      margin-top: 32px;
    }
  }
  .submit-button {
    width: 100%;
    height: 50px;
    font-size: 18px;
    margin: auto;
    background: mix(#ffffff, #4e6ef2, 15%);
    color: #ffffff;
    border: 1px solid transparent;
    cursor: pointer;
    border-radius: 4px;

    &:hover {
      outline: none;
      background: #4e6ef2;
      color: #ffffff;
    }
  }

  .register-button {
    border-radius: 4px;
    background: #f6f7fa;
    border: 1px solid transparent;
    text-align: center;
    outline: none;
    transition: 0.1s;
    outline: none;
    font-size: 18px;
    color: #555555;
    &:hover,
    &:focus {
      outline: none;
      background: mix(#ffffff, #4e6ef2, 15%);
      color: #ffffff;
    }
  }
}
</style>
