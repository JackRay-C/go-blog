<template>
  <aside class="aside">
    <div class="header">
      <div class="circle">
        <img src="/images/avatar.png" alt @click="home" />
        <div class="nickname">{{ nickname }}</div>
      </div>
    </div>

    <div class="aside-nav">
      <ul class="nav-list">
        <li class="nav-list-item">
          <router-link to="/admin/dashboard" active-class="active"
            >首页</router-link
          >
        </li>
        <li class="nav-list-item">
          <router-link to="/admin/article" active-class="active"
            >博客</router-link
          >
        </li>
        <!-- <li class="nav-list-item">
          <router-link to="/admin/draft" active-class="active"
            >草稿箱</router-link
          >
        </li> -->
        <li class="nav-list-item">
          <router-link to="/admin/subject" active-class="active"
            >专题</router-link
          >
        </li>
        <li class="nav-list-item">
          <router-link to="/admin/tag" active-class="active">标签</router-link>
        </li>

        <li class="nav-list-item">
          <el-popover placement="right" width="200" trigger="hover" popper-class="nav-popper">
            <div>
              <ul class="nav-list">
                <li class="nav-list-item">
                  <router-link to="/admin/users" active-class="active"
                    >用户管理</router-link
                  >
                </li>
                <li class="nav-list-item">
                  <router-link to="/admin/roles" active-class="active"
                    >角色管理</router-link
                  >
                </li>
                <li class="nav-list-item">
                  <router-link to="/admin/menus" active-class="active"
                    >菜单设置</router-link
                  >
                </li>
                <li class="nav-list-item">
                  <router-link to="/admin/routes" active-class="active"
                    >路由设置</router-link
                  >
                </li>
                <li class="nav-list-item">
                  <router-link to="/admin/setting" active-class="active"
                    >字典管理</router-link
                  >
                </li>
              </ul>
            </div>

            <a slot="reference" active-class="active">设置</a>
          </el-popover>
        </li>
      </ul>
    </div>

    <div class="logout">退出</div>
  </aside>
</template>

<script>
export default {
  data() {
    return {
      nickname: "任浩杰",
    };
  },
  methods: {
    home() {
      this.$router.push("/admin/dashboard");
    },
  },
};
</script>


<style lang="scss" scoped>
.aside {
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  // box-shadow: 0 0px 20px 0 rgba(31, 35, 41, 0.13);
  box-shadow: 0 0 20px 0 rgba(31, 35, 41, 0.03);
  background: #ffffff;
  position: relative;
}

.header {
  width: 100%;
  height: 60px;
  margin: 20px auto;
  text-align: center;
}

.circle {
  width: 80px;
  height: 80px;
  border-radius: 25px;
  margin: auto;
  color: #555555;

  img {
    width: 60px;
    height: 60px;
  }
  .nickname {
    overflow: hidden;
    font-size: 16px;
    text-align: center;
    width: 60px;
    margin: 8px auto;
    text-overflow: ellipsis;
    -webkit-line-clamp: 1;
    white-space: nowrap;
  }
}

.aside-nav {
  width: 100%;
  margin-top: 30px;
  font-size: 16px;
}

.nav-list {
  text-align: center;
  list-style: none;
  padding: 0;

  .nav-list-item {
    width: 100%;
    height: 56px;
    line-height: 56px;
    margin: 5px auto;
    cursor: pointer;

    a {
      color: inherit;
      width: 100%;
      height: 100%;
      // padding-left: 40px;
      display: block;
      box-sizing: border-box;
      transition: all 0.3s cubic-bezier(0.075, 0.82, 0.165, 1);
      position: relative;

      &:hover {
        background: #f4f5fa;
        color: #4e6ef2;
      }

      &::after {
        content: "";
        display: block;
        position: absolute;
        height: 0;
        width: 3px;
      }
    }

    .active {
      color: #4e6ef2;
      background: #f4f5fa;
      &::after {
        top: 0;
        right: 0;
        height: 100%;
        background-color: #4e6ef2;
        transition: all 0.2s ease-in-out;
        animation: fadeInCenter 0.5s ease-in-out;
      }
    }
  }
}
.logout {
  position: absolute;
  bottom: 40px;
  width: 100%;
  height: 50px;
  line-height: 50px;
  text-align: center;
  cursor: pointer;
}

.logout-text {
  width: 80%;
  height: 100%;
  margin: auto;
  background: #ffffff;
  font-size: 16px;

  &:hover {
    background: #f6f7fa;
  }
}

.fadeInCenter {
  animation: fadeInCenter 0.5s ease-in-out;
}

@keyframes fadeInCenter {
  0% {
    transform: translateY(-50%);
    top: 50%;
  }

  100% {
    transform: translateY(0);
    top: 0;
  }
}
</style>


<style lang="scss">
.nav-popper{
  padding: 0 !important;
}
</style>