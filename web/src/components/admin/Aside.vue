<template>
  <div class="aside">
    <header class="nav-menu">
      <div class="nav-menu-details">
        <div
          class="nav-menu-icon nav-menu-logo"
          :style="'background-image:url(' + logo + ')'"
        ></div>
        <div class="nav-menu-details-sitetitle">
          Blog Admin
        </div>
      </div>
    </header>
    <section class="nav-body">
      <div class="nav-top">
        <ul class="nav-list nav-main">
          <li>
            <router-link to="/admin/dashboard" active-class="active">
            <svg-icon icon-class="dashboard" />
              Dashboard
            </router-link>
          </li>
        </ul>

        <ul class="nav-list nav-manage">
          <li class="nav-list-new">
            <router-link to="/admin/posts" active-class="active">
            <svg-icon icon-class="edit" />
              Posts
            </router-link>
            <router-link
              to="/admin/edit"
              class="nav-new-post"
              active-class="active"
            >
              <span>
                <svg-icon icon-class="plus" class-name="page_svg__a" />
              </span>
            </router-link>
            <div class="nav-post-container">
              <div class="nav-post-container-child">
                <ul class="nav-view-list">
                  <li>
                    <router-link to="/admin/drafts" active-class="active">
                      <span class="viewname">
                        Drafts
                      </span>
                    </router-link>
                  </li>
                  <li>
                    <router-link to="/admin/published" active-class="active">
                      <span class="viewname">
                        Published
                      </span>
                    </router-link>
                  </li>
                </ul>
              </div>
            </div>
          </li>

          <li v-for="route in routes" :key="route.path">
            <router-link :to="route.path" active-class="active" v-if="route.meta && route.meta.sidebar && route.name!== 'Posts' && route.name!= 'Dashboard'">
              <svg-icon :icon-class="route.meta.icon" :class-name="route.meta.iconClass" />
              {{route.name}}
            </router-link>
          </li>
          <!-- <li>
            <router-link to="/admin/pages" active-class="active">
            <svg-icon icon-class="pages" class-name="page_svg__a" />
              Pages
            </router-link>
          </li>
          <li>
            <router-link to="/admin/tag" active-class="active">
              <svg-icon icon-class="tags"/>
              Tags
            </router-link>
          </li>
          <li>
            <router-link to="/admin/subject" active-class="active">
              <svg-icon icon-class="subjects"/>
              Subjects
            </router-link>
          </li>
          <li>
            <router-link to="/admin/dicts" active-class="active">
              <svg-icon icon-class="dicts"/>
              Dicts
            </router-link>
          </li>
          <li>
            <router-link to="/admin/users" active-class="active">
              <svg-icon icon-class="members" class-name="members_svg__cls-1" />
              Accounts
            </router-link>
          </li> -->
        </ul>

        <ul class="nav-list"></ul>
      </div>

      <div>
        <div class="nav-bottom">
          <div class="nav-bottom-child">
            <div class="nav-bottom-left">
              <div class="nav-trigger">
                <div class="nav-trigger-flex" @click="dropdown1 = !dropdown1">
                  <div
                    class="user-avatar"
                    :style="'background-image: url(' + avatar.host +'' + avatar.access_url + ')'"
                  ></div>
                  <svg-icon icon-class="dropdown" class-name="w3 mr1 fill-darkgrey" />
                </div>
                <div class="nav-dropdown-content" v-if="dropdown1">
                  <ul class="dropdown-menu">
                    <li>
                      <div class="account-menu-header">
                        <div
                          class="user-avatar"
                          :style="'background-image: url(' + avatar.host +'' + avatar.access_url + ')'"
                        ></div>
                        <div class="user-info">
                          <h4 class="user-name">任浩杰</h4>
                          <span class="user-email">renhj@bw30.com</span>
                        </div>
                      </div>
                    </li>
                    <li class="divider"></li>
                    <li>
                      <router-link to="/admin/profile" @click="dropdown1 = !dropdown1" class="dropdown-item"
                        >Your profile</router-link
                      >
                    </li>
                    <li>
                      <router-link to="/admin/logout" @click="logout" class="dropdown-item"
                        >Sign out</router-link
                      >
                    </li>
                  </ul>
                </div>
              </div>
            </div>
            <div class="nav-bottom-right">
              <router-link to="/admin/setting" active-class="active">
                <svg-icon icon-class="settings" class-name="settings_svg__a" />
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import SvgIcon from "@/components/SvgIcon"

export default {
  components: {
    SvgIcon
  },
  computed: {
    ...mapGetters([
      'routes',
      'username',
      'nickname',
      'avatar',
    ])
  },
  mounted() {
    console.log(this.avatar)
  },
  data() {
    return {
      logo:
        "http://localhost:8000/static/uploads/image/37e58b4d0a32bca8c2f3858a678855b1.png",
      // avatar:
      //   "http://localhost:8000/static/uploads/image/3cc5519ea30e020ddf20c082e3149d66.png",
      dropdown1: false,
    };
  },
  methods: {
    logout() {
      this.dropdown1 = !this.dropdown1
      this.$sotre.dispatch('DispatchLogout').then(res=> {
        console.log(res)
        this.$route.push("/login")
      })
    }
  }
};
</script>

<style lang="scss" scoped>
.aside {
  display: flex;
  height: 100%;
  flex-direction: column;
}

.nav-menu {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  height: 96px;
  padding: 32px;
  box-sizing: border-box;

  .nav-menu-details {
    display: flex;
    align-items: center;
    flex-grow: 1;
    padding-right: 10px;
    min-width: 0;
    box-sizing: border-box;

    .nav-menu-icon {
      flex-shrink: 0;
      margin-right: 10px;
      width: 32px;
      height: 32px;
      background-color: transparent;
      background-size: 32px;
      border-radius: 6px;
    }

    .nav-menu-details-sitetitle {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      font-size: 1.5rem;
      line-height: 1.3em;
      font-weight: 700;
      color: #15171a;
    }
  }
}

.nav-body {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  flex-grow: 1;
  overflow-y: auto;
  padding: 0;

  .nav-top {
    box-sizing: border-box;

    .nav-list {
      margin: 32px 0 0;
      padding: 0;
      list-style: none;
      font-size: 1.4rem;
      line-height: 1.6em;

      &:first-of-type {
        margin-top: 0;
      }

      li {
        margin: 0;
        padding: 0;
        position: relative;
        line-height: 1.4em;
      }

      .active {
        position: relative;
        color: #15171a;
        font-weight: 700;
        outline: 0;

        svg {
          fill: #15171a;
          font-weight: 700;
        }
      }

      svg {
        margin-right: 17px;
        width: 16px;
        height: 16px;
        line-height: 1;
        transition: none;
        z-index: 999;

        &:not(:root) {
          overflow: hidden;
        }
      }

      .members_svg__cls-1,
      .page_svg__a {
        fill: none;
        stroke: currentColor;
        stroke-linecap: round;
        stroke-linejoin: round;
        stroke-width: 1.5 px;
      }
    }

    .nav-manage {
      .nav-list-new {
        .nav-new-post {
          opacity: 1;
          position: absolute;
          z-index: 999;
          padding: 10px;
          margin: 0;
          right: 12px;
          top: -11px;
          transition: opacity 0.2s ease;

          span {
            width: 36px;
            height: 36px;
            border-radius: 100%;
            display: flex;
            align-items: center;
            justify-content: center;
            box-sizing: border-box;

            svg {
              width: 18px;
              height: 18px;
              margin-right: 0;
              fill: #4b545d;
            }
          }

          &:hover {
            background: none;

            span {
              background: #f1f3f4;

              svc {
                fill: #15171a;
              }
            }
          }

          .active:hover {
            color: #394047;
            background: #f1f3f4;
            opacity: 1;
          }
        }
      }
    }

    .nav-main {
      margin: 24px 0;
    }
  }

  .nav-bottom {
    box-sizing: border-box;
    margin: 0;
    padding: 32px;
    border-top: 1px solid #ebeef0;

    .nav-bottom-child {
      display: flex;
      justify-content: space-between;
      align-items: center;
      box-sizing: border-box;

      .nav-bottom-left {
        .nav-trigger {
          padding: 4px 8px 4px 4px;
          margin: -4px -8px -4px -4px;
          box-sizing: border-box;
          cursor: pointer;
          outline: 0;

          &:hover {
            background: #ebeef0;
            border-radius: 999px;
            cursor: pointer;
          }

          .nav-trigger-flex {
            display: flex;
            box-sizing: border-box;
            align-items: center;
            flex: 1 1 auto;
            min-width: 0;
            min-height: 0;

            .user-avatar {
              position: relative;
              flex-shrink: 0;
              display: block;
              width: 34px;
              height: 34px;
              margin: 0 8px 0 0;
              background-position: 50%;
              background-size: cover;
              border-radius: 100%;
              border: 1px solid #ebeef0;
            }

            .w3 {
              width: 1.3rem;
              height: 2em;
            }
            .mr1 {
              margin-right: 0.4rem;
            }
          }
        }
      }

      .nav-bottom-right {
        a {
          display: flex;
          align-items: center;
          justify-content: center;
          margin-left: 12px;
          padding: 10px;
          border-radius: 999px;
          width: 40px;
          height: 40px;
          line-height: 1;
          color: #394047;

          &:hover {
            background: #ebeef0;
            border-radius: 999px;
            cursor: pointer;
          }

          svg {
            width: 28px;
            height: 28px;
            fill: #394047;
            line-height: 1;
            transition: none;
          }
        }
        .settings_svg__a {
          fill: none;
          stroke: currentColor;
          stroke-linecap: round;
          stroke-linejoin: round;
          stroke-width: 1.5 px;
        }
      }
    }
  }
}

.nav-list a {
  display: flex;
  align-items: center;
  color: #40474f;
  transition: none;
  font-weight: 400;
  padding: 7px 32px 7px 39px;
  font-size: 1.45rem;
  margin: 0;
  border-radius: 0;
  box-sizing: border-box;

  &:hover {
    transition: background 0.1s, color 0.1s;
    text-decoration: none;
    outline: 0;
  }

  &:not(.active):hover {
    color: #394047;
    background: #f1f3f4;
    opacity: 1;
  }

  &:hover svg,
  &:not(.active):hover svg {
    fill: #15171a;
    font-weight: 700;
  }
}

.nav-post-container {
  box-sizing: border-box;
  position: relative;
  overflow: hidden;
  transform: translateY(0);

  .nav-post-container-child {
    top: 0px;
    left: 0px;
    overflow: hidden;
    box-sizing: border-box;

    .nav-view-list {
      padding: 0;
      margin: 0 0 22px;
      list-style: none;
      font-size: 1.4rem;
      line-height: 1.6em;
      li {
        margin: 0;
        padding: 0;
        line-height: 1.4em;

        .active {
          position: relative;
          color: #15171a;
          font-weight: 700;
          outline: 0;

          svg {
            fill: #15171a;
            font-weight: 700;
          }
        }
        a {
          position: relative;
          padding-left: 74px;

          .viewname {
            display: inline-block;
            font-weight: inherit;
            max-width: 160 px;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
          }
        }
      }
    }
  }
}

.nav-dropdown-content {
  position: absolute;
  width: auto;
  background: #ffffff;
  z-index: 1000;
  left: 28px;
  top: 871px;

  .dropdown-menu {
    position: absolute;
    top: 100%;
    left: 0;
    float: left;
    margin: 10px 0 0;
    padding: 6px 0;
    min-width: 290px;
    background: #ffffff;
    background-clip: padding-box;
    border-radius: 3px;
    box-shadow: 0 0 7px rgb(0 0 0 / 8%), 0 2.1px 2.2px -5px rgb(0 0 0 / 1%),
      0 5.1px 5.3px -5px rgb(0 0 0 / 2%), 0 9.5px 10px -5px rgb(0 0 0 / 2%),
      0 17px 17.9px -5px rgb(0 0 0 / 2%), 0 31.8px 33.4px -5px rgb(0 0 0 / 3%),
      0 76px 80px -5px rgb(0 0 0 / 4%);
    top: -324px;
    left: -13px;
    list-style: none;
    text-align: left;
    text-transform: none;
    font-size: 1.4rem;
    font-weight: 400;

    li {
      margin: 0;

      a {
        font-size: 1.4rem;
        margin: 0;
        width: unset;
        padding: 8px 24px 9px;
        display: flex;
        align-items: center;
        clear: both;
        color: #394047;
        text-align: left;
        white-space: nowrap;
        line-height: 1.4em;
        font-weight: 400;
        transition: none;

        &:hover {
          background: rgba(237, 238, 238, 0.6);
          color: #394047;
          text-decoration: none;
          outline: 0;
        }
      }
    }

    .divider {
      overflow: hidden;
      margin: 8px 0;
      height: 1px;
      background: #edeeef;
    }
  }
}

.account-menu-header {
  position: relative;
  display: flex;
  align-items: center;
  padding: 12px 24px;

  .user-avatar {
    position: relative;
    flex-shrink: 0;
    display: block;
    width: 44px;
    height: 44px;
    flex-basis: 44px;
    margin: 0;
    padding: 0;
    background-position: 50%;
    background-size: cover;
    border-radius: 100%;
    border: 1px solid #ebeef0;
  }
}
</style>
