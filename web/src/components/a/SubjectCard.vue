<template>
  <div class="subject-card">
    <div class="subject-info">
      <div class="subject-title" @click="goSubject">{{subject.title}}</div>
      <div class="subject-subtitle">{{count}} 篇文章</div>
      <div class="subject-desc">{{subject.description}}</div>
    </div>
    <button type="button" class="icon-button" @click="showSubjectSetting = !showSubjectSetting">
      <i class="icon">
        <svg
          t="1594175958278"
          class="icon"
          viewBox="0 0 1260 1024"
          version="1.1"
          xmlns="http://www.w3.org/2000/svg"
          p-id="8704"
          width="200"
          height="200"
          fill="#9aa8b6"
        >
          <path
            d="M385.969231 157.538462h787.692307c23.630769 0 39.384615-15.753846 39.384616-39.384616s-15.753846-39.384615-39.384616-39.384615h-787.692307c-23.630769 0-39.384615 15.753846-39.384616 39.384615s15.753846 39.384615 39.384616 39.384616zM78.769231 196.923077c47.261538 0 78.769231-31.507692 78.769231-78.769231s-31.507692-78.769231-78.769231-78.769231-78.769231 31.507692-78.769231 78.769231 31.507692 78.769231 78.769231 78.769231z m0 393.846154c47.261538 0 78.769231-31.507692 78.769231-78.769231s-31.507692-78.769231-78.769231-78.769231-78.769231 31.507692-78.769231 78.769231 31.507692 78.769231 78.769231 78.769231z m0 393.846154c47.261538 0 78.769231-31.507692 78.769231-78.769231s-31.507692-78.769231-78.769231-78.769231-78.769231 31.507692-78.769231 78.769231 31.507692 78.769231 78.769231 78.769231zM385.969231 551.384615h787.692307c23.630769 0 39.384615-15.753846 39.384616-39.384615s-15.753846-39.384615-39.384616-39.384615h-787.692307c-23.630769 0-39.384615 15.753846-39.384616 39.384615s15.753846 39.384615 39.384616 39.384615z m0 393.846154h787.692307c23.630769 0 39.384615-15.753846 39.384616-39.384615s-15.753846-39.384615-39.384616-39.384616h-787.692307c-23.630769 0-39.384615 15.753846-39.384616 39.384616s15.753846 39.384615 39.384616 39.384615z"
            p-id="8705"
          />
        </svg>
      </i>
    </button>
    <div
      class="setting-popover"
      v-show="showSubjectSetting"
      @mouseleave="showSubjectSetting = false"
    >
      <div class="setting-list">
        <span class="setting-list-item">
          <i class="setting-list-item-icon">
            <svg
              t="1594194305376"
              class="icon"
              viewBox="0 0 1412 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              p-id="17217"
              width="200"
              height="200"
            >
              <path
                d="M1398.287427 422.327684A85.333333 85.333333 0 0 0 1326.935262 385.687382h-109.920904V223.216573a85.333333 85.333333 0 0 0-85.333334-85.333334H621.127352L492.404188 8.677966A32.301318 32.301318 0 0 0 471.191382 0h-385.687382A85.333333 85.333333 0 0 0 0.170667 85.333333v911.186441a5.303202 5.303202 0 0 1 0 2.410546 4.338983 4.338983 0 0 0 0 2.410546h1186.470809a30.372881 30.372881 0 0 0 28.444445-19.284369l192.843691-503.322034a85.333333 85.333333 0 0 0-9.642185-56.406779zM85.504 60.745763H458.174433l129.205273 128.723164a28.444444 28.444444 0 0 0 21.212806 8.677966h523.088512a24.587571 24.587571 0 0 1 24.587571 25.06968V385.687382H289.436203a85.815443 85.815443 0 0 0-80.512241 54.478343L60.916429 829.227872V85.333333a24.587571 24.587571 0 0 1 24.587571-24.58757zM1350.076505 482.109228l-184.165725 482.109228H74.415488l192.843691-501.875706A25.551789 25.551789 0 0 1 289.436203 445.951036h1037.499059a25.551789 25.551789 0 0 1 20.730697 10.606403 26.033898 26.033898 0 0 1 2.410546 25.551789z"
                p-id="17218"
              />
            </svg>
          </i>
          <span class="setting-list-item-text">解散</span>
        </span>
        <span class="setting-list-item">
          <i class="setting-list-item-icon">
            <svg
              t="1594194115119"
              class="icon"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              p-id="13996"
              width="200"
              height="200"
            >
              <path
                d="M505.142004 0.00063a205.673873 205.673873 0 0 0-205.809873 198.642878H84.237263c-28.874982 0-50.104969 6.962996-50.104969 35.905978 0 28.737982 21.024987 39.249976 50.104969 39.249976h18.702989v608.079625c0 94.541942 43.687973 142.052913 136.114916 142.052913h537.904669c109.696933 0 137.752915-55.565966 137.752915-142.053913l-1.364999-608.145625h26.007984c28.874982 0 50.376969-10.307994 50.376969-39.250976 0-28.806982-41.502974-35.632978-70.514957-35.632978H710.609878A205.059874 205.059874 0 0 0 505.142004 0.00163zM371.348087 204.240504c4.777997-76.589953 66.213959-137.546915 136.523916-137.546915 70.173957 0 132.359919 61.093962 136.387916 137.547915H371.348087z m-205.809874 682.82558V273.050462h681.254581v614.085622c0 59.864963-29.693982 66.964959-62.390962 66.964959h-545.209664c-43.549973 0-73.653955-6.962996-73.653955-66.896959z"
                p-id="13997"
              />
              <path
                d="M341.313105 799.486138c18.702988 0 34.129979-14.334991 34.129979-43.959973v-303.289813c0-29.556982-15.289991-42.389974-34.129979-42.389974-18.703988 0-34.130979 12.832992-34.130979 42.389974v303.289813c0 29.351982 15.016991 43.959973 34.129979 43.959973z m157.274903-3.753998c18.703988 0 34.130979-10.649993 34.130979-40.274975V452.305352c0-29.556982-15.289991-42.663974-34.129979-42.663974-18.703988 0-34.131979 13.106992-34.131979 42.663974v303.288813c0 29.351982 15.427991 40.205975 34.130979 40.205975z m165.399898-0.682999c18.702988 0 34.129979-9.965994 34.129979-39.591976V452.305352c0-29.556982-15.289991-42.458974-34.129979-42.458974-18.703988 0-34.131979 12.901992-34.131979 42.459974v303.287813c0 29.351982 14.881991 39.522976 34.130979 39.522976z"
                p-id="13998"
              />
            </svg>
          </i>
          <span class="setting-list-item-text">删除</span>
        </span>
        <span class="setting-list-item">
          <i class="setting-list-item-icon">
            <svg
              t="1594193729140"
              class="icon"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              p-id="12664"
              width="200"
              height="200"
            >
              <path
                d="M925.312 449.728l-71.424-11.2a55.68 55.68 0 0 1-27.072-12.224 50.496 50.496 0 0 1-16.832-21.44 184.128 184.128 0 0 1-8.704-41.344 83.648 83.648 0 0 1 4.608-41.344L848.768 264a47.488 47.488 0 0 0 6.656-17.856 39.296 39.296 0 0 0-1.536-18.88 48.64 48.64 0 0 0-9.152-16.32l-31.616-31.616a38.848 38.848 0 0 0-26.048-11.2 39.808 39.808 0 0 0-27.008 7.168l-58.176 42.88a52.608 52.608 0 0 1-27.584 11.2 44.992 44.992 0 0 1-27.584-4.096 145.92 145.92 0 0 1-21.952-12.8 192.896 192.896 0 0 1-25.472-20.928c-8.512-8.128-13.12-15.296-13.76-21.376l-11.264-71.488a38.848 38.848 0 0 0-13.76-24.448A38.848 38.848 0 0 0 534.464 64H489.6a37.888 37.888 0 0 0-17.92 4.608 44.096 44.096 0 0 0-14.784 12.8 34.816 34.816 0 0 0-7.104 17.28l-11.264 71.488a47.616 47.616 0 0 1-6.144 18.368c-3.392 6.144-7.68 11.584-12.8 16.32s-9.984 7.808-14.72 9.152c-11.584 4.096-25.344 7.04-41.344 8.704-16 1.728-29.824 0.192-41.344-4.608L264 175.232a40.32 40.32 0 0 0-27.008-7.168 39.04 39.04 0 0 0-26.048 11.264l-31.616 31.616a48.96 48.96 0 0 0-9.216 16.32 39.296 39.296 0 0 0-1.536 18.88c1.024 6.464 3.2 12.416 6.656 17.856l42.88 58.176a58.88 58.88 0 0 1 11.264 28.032 43.52 43.52 0 0 1-4.096 27.008c-2.752 6.144-7.04 13.504-12.8 22.016-5.824 8.512-12.8 16.96-20.928 25.472-8.192 8.512-15.296 13.056-21.44 13.76l-71.424 11.264a34.112 34.112 0 0 0-13.248 4.608 39.68 39.68 0 0 0-10.752 9.152 55.936 55.936 0 0 0-7.616 12.224A32.896 32.896 0 0 0 64 489.536V534.4c0 9.536 3.392 18.24 10.24 26.048a39.04 39.04 0 0 0 24.512 13.76l71.424 11.2a48.896 48.896 0 0 1 18.368 6.208 64.76 64.76 0 0 1 16.32 12.8 37.12 37.12 0 0 1 9.152 14.72l2.048 7.168a58.752 58.752 0 0 1 3.072 10.688l2.048 11.712a132.864 132.864 0 0 1 2.048 24.448v11.2a30.656 30.656 0 0 1-5.12 17.344l-42.88 58.176a40.32 40.32 0 0 0-7.168 27.008 39.04 39.04 0 0 0 11.264 26.048l31.616 31.616a38.848 38.848 0 0 0 26.048 11.2A39.936 39.936 0 0 0 264 848.576l58.176-42.88a57.6 57.6 0 0 1 37.76-11.712 37.952 37.952 0 0 1 17.344 4.608c10.88 4.8 23.488 13.12 37.76 25.024 14.272 11.904 22.08 21.952 23.488 30.08l11.264 71.424a35.136 35.136 0 0 0 7.168 17.344 44.8 44.8 0 0 0 14.784 12.8c5.76 3.2 11.712 4.736 17.856 4.736h44.928c2.688 0 5.76-0.512 9.152-1.536a58.048 58.048 0 0 0 9.152-3.584 35.328 35.328 0 0 0 7.616-5.12 44.8 44.8 0 0 0 6.656-7.168 57.408 57.408 0 0 0 5.12-8.192 20.288 20.288 0 0 0 2.048-9.152l11.2-71.424a55.168 55.168 0 0 1 12.224-27.008 50.496 50.496 0 0 1 21.44-16.832c11.584-4.096 25.536-6.976 41.856-8.704 16.32-1.728 29.952-0.192 40.832 4.608L760 848.768a37.888 37.888 0 0 0 13.248 5.632c4.736 1.024 9.536 1.344 14.272 1.024a45.312 45.312 0 0 0 13.76-3.072 32.768 32.768 0 0 0 11.712-7.616l31.616-31.616a39.232 39.232 0 0 0 11.264-26.048 40.256 40.256 0 0 0-7.168-27.008l-42.88-58.176a52.608 52.608 0 0 1-11.2-27.584 44.992 44.992 0 0 1 4.096-27.584 188.8 188.8 0 0 1 33.728-47.424c8.192-8.512 15.296-13.056 21.44-13.76l71.424-11.264a39.168 39.168 0 0 0 24.512-13.76A39.104 39.104 0 0 0 960 534.464V489.6a37.504 37.504 0 0 0-4.608-17.856 43.52 43.52 0 0 0-12.8-14.784 35.584 35.584 0 0 0-17.28-7.232zM512.064 689.536a171.072 171.072 0 0 1-125.504-52.032C351.872 602.816 334.528 560.96 334.528 512s17.344-90.816 52.032-125.504a171.072 171.072 0 0 1 125.504-52.032c48.96 0 90.816 17.344 125.504 52.032C672.32 421.184 689.664 463.04 689.664 512s-17.344 90.816-52.032 125.504a171.328 171.328 0 0 1-125.568 52.032z"
                p-id="12665"
              />
            </svg>
          </i>
          <span class="setting-list-item-text">设置</span>
        </span>
      </div>
    </div>
  </div>
</template>

<script>
import { getSubjectPostCount } from "@/api/subject.js";

import {getFileById} from "@/api/file.js";
export default {
  name: "SubjectCard",
  props: ["subject"],
  data() {
    return {
      count: "",
      showSubjectSetting: false
    };
  },
  mounted() {
    this.countSubjectPost();
  },
  methods: {
    goSubject() {
      this.$router.push(`/admin/subject/${this.subject.subjectId}/post`);
    },
    countSubjectPost() {
      getSubjectPostCount(this.subject.id).then(res => {
        if(res.code === 200) {
          this.count = res.data.total_rows;
        }
      })
    },
    fetchAvatar() {
      getFileById(this.subject.image).then(res => {
        if(res.code === 200){
          console.log(res)
          this.avatar = "http://localhost:8000/" +  res.data.access_url
        }
      }).catch(err => {
        console.log(err)
      })
    }
  }
};
</script>

<style lang="scss" scoped>
.subject-card {
  position: relative;
  width: 46%;
  height: 276px;
  padding: 38px 25px 0 26px;
  background: #ffffff;
  border: 10px;
  box-sizing: border-box;
  margin-bottom: 25px;
  margin-right: 50px;
  border-radius: 8px;
  box-shadow: 0 8px 24px 0 rgba(54, 97, 174, 0.05);
  transition: all 0.5s;
  &:hover {
    box-shadow: 0 6px 28px 0 rgba(24, 52, 117, 0.2);
    // transform 会影响弹框组件的层级，因此在此不使用
    // transform: scale(1.04, 1.04);
    transition: all 0.8s cubic-bezier(0.075, 0.82, 0.165, 1);
  }

  .subject-card-header {
    height: 70px;
    position: absolute;
    top: 0;
    right: 0;
    padding: 13px 25px 0 26px;
    display: flex;
    justify-content: flex-end;
    flex-direction: row;
  }

  .subject-info {
    width: 360px;
    margin-left: 12px;
    overflow: hidden;

    .subject-title {
      font-size: 25px;
      font-weight: 700;
      color: #4e6ef2;
      white-space: nowrap;
      text-overflow: ellipsis;
      overflow: hidden;
      cursor: pointer;
      transition: color 0.2s linear;
    }
    .subject-subtitle {
      font-size: 15px;
      font-weight: 500;
      padding: 10px 0;
      color: #9aa8b6;
    }
    .subject-desc {
      margin-top: 10px;
      line-height: 24px;
      display: -webkit-box;
      line-clamp: 3;
      -webkit-line-clamp: 3;
      -webkit-box-orient: vertical;
      overflow: hidden;
      font-size: 15px;
      font-weight: 500;
      color: #888888;
    }
  }
  .icon-button {
    position: absolute;
    top: 38px;
    right: 25px;
    width: 40px;
    height: 40px;
    z-index: 1000;
    border: none;
    border-radius: 50%;
    cursor: pointer;
    vertical-align: middle;
    transition-duration: 0.28s;
    background-color: transparent;
    transition-property: transofrm, opacity, background-color;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
    user-select: none;
    transition: transform 0.2s cubic-bezier(0.4, 0, 0.2, 1),
      background-color 0.2s cubic-bezier(0.4, 0, 0.2, 1),
      left 0.2s cubic-bezier(0.4, 0, 0.2, 1),
      right 0.2s cubic-bezier(0.4, 0, 0.2, 1),
      box-shadow 0.28s cubic-bezier(0.4, 0, 0.2, 1),
      max-width 0.25s cubic-bezier(0.4, 0, 0.2, 1),
      width 0.25s cubic-bezier(0.4, 0, 0.2, 1);

    color: #1c1f22;
    box-sizing: border-box;
    &:active,
    &:focus {
      border-style: none;
      outline: none;
    }
    &:hover {
      background: #e5e5e5;
      color: #222222;
      border-radius: 50%;
    }

    .icon {
      width: 24px;
      height: 24px;
      border: none;
      pointer-events: none;
    }
  }
  /** 专题设置弹出框 */
  .setting-popover {
    position: absolute;
    min-width: 150px;
    border-radius: 8px;
    background: #ffffff;
    top: 40px;
    left: 90%;
    box-shadow: 0 6px 28px 0 rgba(24, 52, 117, 0.2);
    z-index: 2000;
    user-select: none;
    text-align: left;
    transform: scale(1.04, 1.04);
    transition: all 0.3s cubic-bezier(0.075, 0.82, 0.165, 1);

    .setting-list {
      margin: 0;
      padding: 10px 0;
      background: #ffffff;
      color: #222222;

      .setting-list-item {
        text-overflow: ellipsis;
        position: relative;
        overflow: hidden;
        white-space: nowrap;
        display: block;
        min-width: 100%;
        height: 40px;
        line-height: 30px;
        padding: 5px 20px;
        color: #888888;
        text-decoration: none;
        box-sizing: border-box;

        &:hover {
          background: #f6f7fa;
        }

        .setting-list-item-icon {
          font-size: 16px;
          padding: 5px 0;
          line-height: 30px;
          color: #222222;
          visibility: visible;
          display: inline;
          vertical-align: middle;

          .icon {
            width: 16px;
            height: 16px;
            fill: #888888;
          }
        }

        .setting-list-item-text {
          margin: 0 10px;
        }
      }
    }
  }
}
</style>