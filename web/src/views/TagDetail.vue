<template>
  <div class="tagDetail fadeInUp">
    <div
      class="subject-header"
      :style="{ backgroundImage: 'url(' + cover_image + ')' }"
    >
      <div class="subject-header-box">
        <div class="subject-header-title">{{ tag.name }}</div>
      </div>
    </div>


    <div class="content">
      <div class="subject-posts">
        <div class="subject-posts-info">
          <div class="subject-post-info-box">
            <div class="info-item">
              <div class="info-item-number">
                <span class="number">{{posts.length}}</span> 篇
              </div>
              <div class="info-item-text">文 章</div>
            </div>
          </div>
        </div>

        <PostCard v-for="post in posts" :post="post" :key="post.id" />

        <Pagination :pageCount="total_page" :pagerCount="13" @change="handlePageChange" />
        <br />
      </div>
    </div>
  </div>
</template>

<script>
import { getPostByTagId, getTagsById } from "@/api/tag";
import {getFileById} from "@/api/file.js";
import PostCard from "@/components/PostCard.vue";
import Pagination from "@/components/Pagination.vue";

export default {
  name: "TagDetail",
  components: {
    PostCard,
    Pagination,
  },
  data() {
    return {
      tagId: this.$route.params.id,
      page_size: 10,
      total_page: 0,
      page_no: 1,
      total_rows: 0,
      posts: [],
      tag: "",
      cover_image: "",
    };
  },
  beforeRouteUpdate(to, from, next) {
    this.page_no = to.query.page || 1;
    this.fetchPostsByTag();
    next();
  },
  mounted() {
    this.fetchTag();
    this.fetchPostsByTag();
  },
  methods: {
    fetchPostsByTag() {
      getPostByTagId(this.tagId, {
        page_no: this.page_no,
        page_size: this.page_size,
      }).then((res) => {
        if (res.code === 200 && res.data) {
          this.posts = res.data.list;
          this.page_size = res.data.page_size;
          this.page_no = res.data.page_no;
          this.total_page = res.data.total_page;
          this.total_rows = res.data.total_rows;
        }
      });
    },
    fetchTag() {
      getTagsById(this.tagId).then((res) => {
        if (res.code === 200 && res.data) {
          this.tag = res.data;
          getFileById(res.data.cover_image).then(res => {
            this.cover_image = 'http://localhost:8000/' + res.data.access_url
          })
        }
      });
    },
    handlePageChange(current) {
      this.page_no = current;
      this.$router.push({ path: "/tag", query: { page: this.page_no } });
    },
  },
};
</script>

<style lang="scss" scoped>
.tagDetail {
  width: 100%;

  .subject-header {
    height: 360px;
    opacity: 0.9;
    background-repeat: no-repeat;
    background-position: center center;
    background-size: cover;
    overflow: hidden;
    .subject-header-box {
      position: relative;
      display: flex;
      flex-direction: column;
      -webkit-box-pack: center;
      justify-content: center;
      align-items: center;
      width: 1200px;
      height: 100%;
      margin: 0 auto;

      .subject-header-title {
        height: 40px;
        line-height: 40px;
        font-size: 40px;
        font-weight: 500;
        color: #ffffff;
      }

      .subject-header-desc {
        width: 818px;
        margin-top: 40px;
        text-align: center;
        font-size: 18px;
        font-weight: 500;
        color: #ffffff;
      }
    }
  }

    .content {
    width: 100%;
    padding-bottom: 100px;

    .subject-posts {
      width: 60%;
      position: relative;
      margin: 0 auto;

      .subject-posts-info {
        position: relative;
        height: 92px;

        .subject-post-info-box {
          display: flex;
          -webkit-box-align: center;
          align-items: center;
          justify-content: space-around;
          position: absolute;
          top: -40%;
          width: 100%;
          padding: 0 160px 10px;
          background: #ffffff;
          border-radius: 8px;
          box-sizing: border-box;

          .info-item {
            width: auto;
            text-align: center;
            .info-item-number {
              font-size: 12px;
              font-weight: 500;
              color: #888888;
              .number {
                font-size: 37.5px;
                margin-right: 3px;
                color: #4c4c4c;
                font-family: Georgia;
              }
            }
            .info-item-text {
              margin-top: -4px;
              font-size: 16px;
              font-weight: 500;
              color: #888888;
            }
          }
        }
      }
    }
  }


}
</style>