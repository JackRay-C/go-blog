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
              <svg viewBox="0 0 24 24">
                <path
                  d="M22.272 23.247a.981.981 0 00.978-.978V9.747a1.181 1.181 0 00-.377-.8L12 .747l-10.873 8.2a1.181 1.181 0 00-.377.8v12.522a.981.981 0 00.978.978z"
                  fill="none"
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="1.5"
                ></path>
              </svg>
              Dashboard
            </router-link>
          </li>
        </ul>

        <ul class="nav-list nav-manage">
          <li class="nav-list-new">
            <router-link to="/admin/posts" active-class="active">
              <svg
                t="1636681531011"
                class="icon"
                viewBox="0 0 1024 1024"
                version="1.1"
                xmlns="http://www.w3.org/2000/svg"
                p-id="1190"
                width="200"
                height="200"
              >
                <path
                  d="M1006.592 925.39392c0 54.311936-44.859392 98.60608-100.196352 98.60608H100.196352C45.207552 1024 0 979.994624 0 925.710336V235.505664C0 181.192704 45.125632 137.216 100.79232 137.216h539.76576v49.265664H100.79232c-27.92448 0-50.46272 21.972992-50.46272 49.024v690.204672c0 26.954752 22.550528 49.024 49.866752 49.024h806.199296c27.416576 0 49.866752-22.116352 49.866752-49.340416v-626.944H1006.592v626.944z m0 0c0 54.311936-44.859392 98.60608-100.196352 98.60608H100.196352C45.207552 1024 0 979.994624 0 925.710336V235.505664C0 181.192704 45.125632 137.216 100.79232 137.216h539.76576v49.265664H100.79232c-27.92448 0-50.46272 21.972992-50.46272 49.024v690.204672c0 26.954752 22.550528 49.024 49.866752 49.024h806.199296c27.416576 0 49.866752-22.116352 49.866752-49.340416v-626.944H1006.592v626.944z"
                  p-id="1191"
                ></path>
                <path
                  d="M516.465664 517.764096c19.206144-6.542336 50.36032-24.948736 64.864256-38.288384l378.635264-348.248064c19.523584-17.956864 19.503104-47.152128-0.088064-65.272832l-5.225472-4.83328c-20.186112-18.670592-52.790272-18.783232-72.92928-0.260096L503.044096 409.147392c-13.686784 12.589056-31.880192 40.728576-37.430272 57.837568l-24.835072 76.560384 75.685888-25.781248zM989.96224 27.654144l5.225472 4.83328c39.458816 36.49536 39.602176 95.865856-0.011264 132.3008L616.539136 513.03424c-19.570688 18.000896-56.854528 40.067072-82.895872 48.93696l-123.001856 41.897984c-26.315776 8.964096-40.638464-4.017152-32.60416-28.78464l39.502848-121.772032c7.969792-24.57088 30.553088-59.567104 50.295808-77.725696L846.512128 27.301888c39.836672-36.63872 103.842816-36.28032 143.449088 0.352256z"
                  p-id="1192"
                ></path>
              </svg>
              Posts
            </router-link>
            <router-link
              to="/admin/edit"
              class="nav-new-post"
              active-class="active"
            >
              <span>
                <svg viewBox="0 0 24 24">
                  <path
                    d="M12 1v22m11-11H1"
                    fill="none"
                    stroke="#000"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  ></path>
                </svg>
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
          <li>
            <router-link to="/admin/pages" active-class="active">
              <svg viewBox="0 0 24 24">
                <path
                  class="page_svg__a"
                  d="M16.5 21.513a1.5 1.5 0 01-1.9 1.446L4.1 20.042A1.5 1.5 0 013 18.6V2.487a1.5 1.5 0 011.9-1.446l10.5 3.391a1.5 1.5 0 011.1 1.445z"
                ></path>
                <path
                  class="page_svg__a"
                  d="M4.5.987h15a1.5 1.5 0 011.5 1.5v15.75a1.5 1.5 0 01-1.5 1.5h-3"
                ></path>
              </svg>
              Pages
            </router-link>
          </li>
          <li>
            <router-link to="/admin/tag" active-class="active">
              <svg viewBox="0 0 24 24">
                <path
                  d="M1.061 2.56v6.257a3 3 0 00.878 2.121L13.5 22.5a1.5 1.5 0 002.121 0l6.879-6.88a1.5 1.5 0 000-2.121L10.939 1.938a3 3 0 00-2.121-.878H2.561a1.5 1.5 0 00-1.5 1.5z"
                  fill="none"
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="1.5"
                ></path>
              </svg>
              Tags
            </router-link>
          </li>
          <li>
            <router-link to="/admin/subject" active-class="active">
              <svg
                t="1636704534949"
                class="icon"
                viewBox="0 0 1024 1024"
                version="1.1"
                xmlns="http://www.w3.org/2000/svg"
                p-id="16114"
                width="200"
                height="200"
              >
                <path
                  d="M95.744 1008.128c-17.408 0-31.744-13.824-32.256-31.744V49.152c0.512-18.432 15.36-33.28 34.304-33.28h829.44c8.704 0 17.408 3.584 23.552 10.24 6.144 6.656 9.728 14.848 9.728 24.064v925.696c0 3.584-0.512 7.168-2.048 11.776-4.608 12.288-16.384 19.968-30.208 19.968-4.096 0-8.192-0.512-11.776-2.048l-356.864-137.728c-14.848-5.632-31.232-8.704-47.104-8.704-16.384 0-32.256 3.072-47.104 8.704l-357.376 137.728H107.52c-3.072 1.536-7.168 2.56-10.752 2.56h-1.024z m44.544-941.056c-13.824 0-24.576 10.752-24.576 24.576v845.824L482.816 803.84c9.728-3.584 19.968-5.632 29.696-5.632 10.24 0 19.968 2.048 29.696 5.632l367.104 133.632V92.16c0-13.824-10.752-24.576-24.576-24.576H140.288z"
                  p-id="16115"
                ></path>
                <path
                  d="M305.152 510.976c0-15.872 12.8-28.16 28.16-28.16h356.864c15.872 0 28.16 12.8 28.16 28.16 0 15.872-12.8 28.16-28.16 28.16H333.824c-15.872 0.512-28.672-12.288-28.672-28.16z m0-206.848c0-15.872 12.8-28.16 28.16-28.16h356.864c15.872 0 28.16 12.8 28.16 28.16 0 15.872-12.8 28.16-28.16 28.16H333.824c-15.872 0-28.672-12.288-28.672-28.16z"
                  p-id="16116"
                ></path>
              </svg>
              Subjects
            </router-link>
          </li>
          <li>
            <router-link to="/admin/dicts" active-class="active">
              <svg
                t="1636705394328"
                class="icon"
                viewBox="0 0 1024 1024"
                version="1.1"
                xmlns="http://www.w3.org/2000/svg"
                p-id="17857"
                width="200"
                height="200"
              >
                <path
                  d="M917.344 50.464h-491.84c-88.992 0-288.16-35.728-331.04 82.48-20.288 55.888 30.224 153.808 93.744 153.808h729.152l-24-24v594.752c0 14.944 1.072 57.232 0.24 60.704-11.632 47.264-59.184 36.464-93.696 36.464H170.112c-57.84 0-34.08-124.8-34.08-160.112V168v-3.808c0-30.96-48-30.96-48 0v590.8c0 63.008-14.192 143.744 6.416 204.512 19.84 58.496 114.8 43.152 160.064 43.152h605.76c76.112 0 81.104-82.432 81.104-135.536V513.312 262.736c0-13.088-10.928-24-24-24H192.8c-36.4 0-70.608-69.168-46.144-100.464 30.224-38.672 66.384-39.824 103.2-39.824H917.36c30.944 0.016 30.944-47.984-0.016-47.984z"
                  p-id="17858"
                ></path>
                <path
                  d="M190.656 185.264H914.4c30.96 0 30.96-48 0-48H190.656c-30.96 0-30.96 48 0 48zM222.56 262.752h48v709.312h-48z"
                  p-id="17859"
                ></path>
              </svg>
              Dicts
            </router-link>
          </li>
          <li>
            <router-link to="/admin/users" active-class="active">
              <svg id="members_svg__Regular" viewBox="0 0 24 24">
                <circle
                  class="members_svg__cls-1"
                  cx="7.5"
                  cy="7.875"
                  r="4.125"
                ></circle>
                <path
                  class="members_svg__cls-1"
                  d="M.75 20.25a6.75 6.75 0 0113.5 0"
                ></path>
                <circle
                  class="members_svg__cls-1"
                  cx="17.727"
                  cy="10.125"
                  r="3.375"
                ></circle>
                <path
                  class="members_svg__cls-1"
                  d="M15.813 15.068a5.526 5.526 0 017.437 5.182"
                ></path>
              </svg>
              Accounts
            </router-link>
          </li>
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
                    :style="'background-image: url(' + avatar + ')'"
                  ></div>
                  <svg viewBox="0 0 26 24" class="w3 mr1 fill-darkgrey">
                    <path
                      clip-rule="evenodd"
                      d="M1.043 6.604a1 1 0 011.414 0L13 17.146 23.543 6.604a1 1 0 011.414 1.414l-10.72 10.719a1.75 1.75 0 01-2.474 0L1.042 8.018a1 1 0 010-1.414zm11.78 10.72v-.001zm.355 0v-.001z"
                    ></path>
                  </svg>
                </div>
                <div class="nav-dropdown-content" v-if="dropdown1">
                  <ul class="dropdown-menu">
                    <li>
                      <div class="account-menu-header">
                        <div
                          class="user-avatar"
                          :style="'background-image: url(' + avatar + ')'"
                        ></div>
                        <div class="user-info">
                          <h4 class="user-name">任浩杰</h4>
                          <span class="user-email">renhj@bw30.com</span>
                        </div>
                      </div>
                    </li>
                    <li class="divider"></li>
                    <li>
                      <router-link to="/admin/profile" class="dropdown-item"
                        >Your profile</router-link
                      >
                    </li>
                    <li>
                      <router-link to="/admin/profile" class="dropdown-item"
                        >Sign out</router-link
                      >
                    </li>
                  </ul>
                </div>
              </div>
            </div>
            <div class="nav-bottom-right">
              <router-link to="/admin/setting" active-class="active">
                <svg viewBox="0 0 24 24">
                  <path
                    class="settings_svg__a"
                    d="M10.546 2.438a1.957 1.957 0 002.908 0L14.4 1.4a1.959 1.959 0 013.41 1.413l-.071 1.4a1.958 1.958 0 002.051 2.054l1.4-.071a1.959 1.959 0 011.41 3.41l-1.042.94a1.96 1.96 0 000 2.909l1.042.94a1.959 1.959 0 01-1.413 3.41l-1.4-.071a1.958 1.958 0 00-2.056 2.056l.071 1.4A1.959 1.959 0 0114.4 22.6l-.941-1.041a1.959 1.959 0 00-2.908 0L9.606 22.6A1.959 1.959 0 016.2 21.192l.072-1.4a1.958 1.958 0 00-2.056-2.056l-1.4.071A1.958 1.958 0 011.4 14.4l1.041-.94a1.96 1.96 0 000-2.909L1.4 9.606A1.958 1.958 0 012.809 6.2l1.4.071a1.958 1.958 0 002.058-2.06L6.2 2.81A1.959 1.959 0 019.606 1.4z"
                  ></path>
                  <circle
                    class="settings_svg__a"
                    cx="12"
                    cy="12.001"
                    r="4.5"
                  ></circle>
                </svg>
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  data() {
    return {
      logo:
        "http://localhost:8000/static/uploads/image/37e58b4d0a32bca8c2f3858a678855b1.png",
      avatar:
        "http://localhost:8000/static/uploads/image/3cc5519ea30e020ddf20c082e3149d66.png",
      dropdown1: false,
    };
  },
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
              width: 1.2rem;
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
