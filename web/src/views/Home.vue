<template>
  <div class="home">
    <div class="header fadeInUp">
      <div class="background" :style="{backgroundImage: 'url(' + banner_image + ')'}">
        <div class="title">
          <div class="text">{{bannerTitle}}</div>
        </div>
      </div>

      <div class="search">
        <span class="search-block">
          <span class="search-input-prefix">
            <i aria-label="图标：Search" class="search-icon">
              <svg
                viewBox="64 64 896 896"
                class
                data-icon="search"
                width="1em"
                height="1em"
                fill="currentColor"
                aria-hidden="true"
                focusable="false"
              >
                <path
                  d="M909.6 854.5L649.9 594.8C690.2 542.7 712 479 712 412c0-80.2-31.3-155.4-87.9-212.1-56.6-56.7-132-87.9-212.1-87.9s-155.5 31.3-212.1 87.9C143.2 256.5 112 331.8 112 412c0 80.1 31.3 155.5 87.9 212.1C256.5 680.8 331.8 712 412 712c67 0 130.6-21.8 182.7-62l259.7 259.6a8.2 8.2 0 0 0 11.6 0l43.6-43.5a8.2 8.2 0 0 0 0-11.6zM570.4 570.4C528 612.7 471.8 636 412 636s-116-23.3-158.4-65.6C211.3 528 188 471.8 188 412s23.3-116.1 65.6-158.4C296 211.3 352.2 188 412 188s116.1 23.2 158.4 65.6S636 352.2 636 412s-23.3 116.1-65.6 158.4z"
                />
              </svg>
            </i>
          </span>
          <input
            type="text"
            autocomplete="off"
            class="input"
            placeholder="Search..."
            v-model="value"
          />
          <span class="search-input-suffix">
            <button type="button" class="search-button" @click="search">
              <span>搜索</span>
            </button>
          </span>
        </span>
      </div>
    </div>

    <div class="content fadeInUp">
      <PostCard v-for="post,index in posts" :key="index" :post="post" />

      <Pagination :pageCount="total_page" :pagerCount="10" @change="handlePageChange" />
      <br />
    </div>

    <FixedHeader v-show="fixedHeaderVisibility" />
  </div>
</template>

<script>
import PostCard from "@/components/PostCard.vue";
import Pagination from "@/components/Pagination.vue";
import FixedHeader from "@/components/FixedHeader.vue";
import { listPosts } from "@/api/web/post.js";

export default {
  name: "Home",
  components: {
    PostCard,
    Pagination,
    FixedHeader
  },
  data() {
    return {
      value: "",
      page_size: 10,
      total_page: 0,
      page_no: 1,
      total_rows: 0,
      posts: [],
      bannerTitle: "无名天地之始，有名万物之母",
      fixedHeaderVisibility: false,
      banner_image: "http://localhost:8000/static/uploads/image/0cc175b9c0f1b6a831c399e269772661.jpg"
    };
  },
  beforeRouteUpdate(to, from, next) {
    this.page_no = to.query.page || 1
    this.fetchPageData()
    next()
  },
  mounted() {
    this.page_no = this.$route.query.page  || 1
    this.fetchPageData();
    window.addEventListener("scroll", this.scrollTop);
  },
  methods: {
    fetchPageData() {
      listPosts({page_no: this.page_no, page_size: this.page_size})
        .then(res => {
          console.log(res)
          if (res.code === 200 && res.data) {
            this.posts = res.data.list;
            this.total_rows = res.data.total_rows;
            this.page_no = res.data.page_no;
            this.page_size = res.data.page_size;
            this.total_page = res.data.total_page;
          }
        })
        .catch(err => {
          console.log(err);
        });
    },
    search(){
      console.log(this.value)

      let searchText = this.value

      console.log(searchText) 
      listPosts({page_no: this.page_no, page_size: this.page_size, search: searchText})
        .then(res => {
          console.log(searchText)
          console.log(res)
          if (res.code === 200 && res.data) {
            this.posts = res.data.list;
            this.total_rows = res.data.total_rows;
            this.page_no = res.data.page_no;
            this.page_size = res.data.page_size;
            this.total_page = res.data.total_page; 
            this.value = searchText
          }
        })
        .catch(err => {
          console.log(err);
        });
    },
    handlePageChange(current) {
      this.page_no = current;
      this.$router.push({path: "/", query: {page: this.page_no}})
    },
    scrollTop() {
      var scrollTop =
        window.pageYOffset ||
        document.documentElement.scrollTop ||
        document.body.scrollTop;
      var clientHeight = document.querySelector(".search").clientHeight;
      var offsetTop =
        document.querySelector(".search").offsetTop + clientHeight;

      if (scrollTop > offsetTop) {
        this.fixedHeaderVisibility = true;
      } else {
        this.fixedHeaderVisibility = false;
      }
    }
  },
  destroyed() {
    window.removeEventListener("scroll", this.scrollTop);
  }
};
</script>

<style lang="scss" scoped>
.home {
  margin: 0;
  padding: 0;
  font-weight: 400;
  font-size: 16px;
  background: #f6f7fa;
}

.header {
  width: 100%;
  margin-bottom: 50px;
  position: relative;
}
.header .background {
  width: 100%;
  height: 360px;
  background-size: cover;
  background-repeat: no-repeat;
  background-position: 50% 50%;
  background-image: url("../assets/pc-background-1.png");
}
.header .background .title {
  color: #ffffff;
  width: 650px;
  height: 100%;
  text-align: center;
  margin: auto;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.header .background .title .text {
  height: 60px;
  color: #ffffff;
  font-size: 48px;
  font-weight: 700;
  line-height: 1.25;
  font-family: Gilroy, -apple-system, BlinkMacSystemFont, Segoe UI, PingFang SC,
    Hiragino Sans GB, Microsoft YaHei, Helvetica Neue, Helvetica, Arial,
    sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol;
}
.header .search {
  height: 60px;
  text-align: center;
  margin-top: -25px;
  box-sizing: border-box;
}

.header .search .search-block {
  width: 60%;
  height: 100%;
  border-radius: 4px;
  box-shadow: 0 8px 24px 0 rgba(54, 97, 174, 0.05);
  font-size: 16px;
  color: #3e4c5b;
  position: relative;
  display: inline-block;
  font-variant: tabular-nums;
  list-style: none;
  text-align: start;
  box-sizing: border-box;
}

.header .search .search-block .search-input-prefix {
  left: 16px;
  font-size: 18px;
  position: absolute;
  top: 50%;
  z-index: 2;
  line-height: 0;
  color: #9ca2a9;
  transform: translateY(-50%);
}

.header .search .search-block .input {
  border: none;
  padding-left: 48px;
  padding-right: 220px;
  min-height: 100%;
  position: relative;
  font-size: 18px;
  text-align: left;
  resize: none;
  box-sizing: border-box;
  position: relative;
  display: inline-block;
  width: 100%;
  height: 32px;
  color: #3e4c5b;
  line-height: 1.5;
  border-radius: 4px;
  transition: all 0.3s;
  caret-color: #4e6ef2;
  overflow: visible;
  border: 1px solid rgba(0, 0, 0, 0);

  &:hover,
  &:focus {
    outline: none;
    border: 1px solid #4e6ef2;
    border-right-width: 1px !important;
  }
}

.header .search .search-block .search-input-suffix {
  top: 0;
  right: 0;
  bottom: 0;
  transform: none;
  font-size: 16px;
  position: absolute;
  z-index: 2;
  line-height: 0;
  color: #9ca2a9;

  .search-button {
    font-weight: 800;
    position: relative;
    width: 180px;
    height: 100%;
    border-radius: 3px;
    background: #4e6ef2;
    font-size: 18px;
    letter-spacing: 10px;
    color: #ffffff;
    border-bottom-left-radius: 0;
    border-top-left-radius: 0;

    outline: 0;
    line-height: 1.499;
    text-align: center;
    border: 1px solid #4e6ef2;
    cursor: pointer;
    transition: padding 0.3s cubic-bezier(0.645, 0.045, 0.355, 1),
      background-color 0.3s, border-color 0.3s;
    user-select: none;
    touch-action: manipulation;

    &:hover {
      background-color: #4e6ef2;
      border-color: #4e6ef2;
    }
    &:focus {
      background-color: #4e6ef2;
      border-color: #4e6ef2;
    }
  }
}

.content {
  padding: 0;
  width: 60%;

  margin: 50px auto;
  background: #f6f7fa;
}
</style>
