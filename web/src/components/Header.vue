<template>
  <header class="header">
    <nav class="nav">
      <div class="nav-left">
        <router-link to="/">
          <span class="title">{{title}}</span>
        </router-link>
      </div>

      <ul class="nav-mid">
        <!-- <li>
          <router-link to="/category" active-class="active">分类</router-link>
        </li> -->
        <li>
          <router-link to="/subject" active-class="active">专题</router-link>
        </li>
        <li>
          <router-link to="/tag" active-class="active">标签</router-link>
        </li>
        <li>
          <router-link to="/about" active-class="active">关于我</router-link>
        </li>
      </ul>

      <div class="nav-right">
        <div class="login-button">
          <button @click="login">登录</button>
        </div>
      </div>
    </nav>
  </header>
</template>

<script>
export default {
  name: "Header",
  data() {
    return {
      title: "任浩杰的个人博客"
    };
  },
  methods: {
    login(){
      this.$router.push("/login")
    }
  }
};
</script>

<style lang="scss" scoped>
.header {
  width: 100%;
  height: 72px;
  box-shadow: 0 2px 20px 0 rgba(31, 35, 41, 0.1);
  background: #ffffff;
  z-index: 1;
  font-family: "pingfang SC", "helvetica neue", arial, "hiragino sans gb",
    "microsoft yahei ui", "microsoft yahei", simsun, sans-serif;
}
.header .nav {
  display: flex;
  align-items: center;
  margin: 0 40px;
  height: 72px;
  justify-content: space-between;

  .nav-left {
    justify-content: flex-start;

    .title {
      font-weight: 800;
      font-size: 20px;
      color: #333;
      cursor: pointer;
      transition: transform 0.4s;
    }
  }

  .nav-mid {
    display: flex;
    flex-direction: row;
    font-size: 16px;
    line-height: 1;
    list-style: none;
    justify-content: flex-end;
    flex: auto;
    padding-right: 100px;

    li {
      display: inline-block;
      cursor: pointer;
      box-sizing: border-box;
      border: 1px solid transparent;
      text-align: center;
      margin: 0 5px;
      min-width: 100px;
      font-weight: 500;
      height: 50px;
      line-height: 50px;
      vertical-align: bottom;

      a {
        display: block;
        height: 100%;
        width: 100%;
        color: #1c1f22;
        box-sizing: border-box;
        &:hover {
          background: #f6f7fa;
        }
      }
      .active {
        color: #4E6EF2;
      }
    }
  }
  .nav-right {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    button {
      min-width: 80px;
      font-size: 16px;
      display: inline-block;
      // line-height: 1;
      white-space: nowrap;
      cursor: pointer;
      background: #ffffff;

      border: 1px solid #4E6EF2;
      color: #4E6EF2;
      text-align: center;
      outline: none;
      transition: 0.1s;
      font-weight: 500;
      padding: 8px 30px;
      letter-spacing: 2px;
   
      &:hover {
        border: 1px solid #4E6EF2;
        color: #ffffff;
        background: mix(#ffffff, #4E6EF2, 10%);
      }
    }
  }
}
</style>

