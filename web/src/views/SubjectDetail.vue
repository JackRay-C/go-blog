<template>
  <div class="subject-detail">
    <el-skeleton style="width:100%" :loading="loading1 && loading2" animated :throttle="2000">
      <template slot="template">
        <el-skeleton-item variant="image" style="height: 360px;">
          <div class="subject-header-box">
            <el-skeleton-item variant="text" class="subject-header-title" />
            <el-skeleton-item variant="text" class="subject-header-desc" />
          </div>
        </el-skeleton-item>
        <div class="content">
          <div class="subject-posts">
            <div class="subject-posts-info">
              <el-skeleton-item variant="rect" style="height:92px;position: absolute;top: -40%;width: 100%;background:#ffffff;" />
            </div>


            <el-skeleton-item variant="rect" v-for="i in 3" :key="i"  style="width: 100%;height: 200px;margin:24px auto;" />
            
            <br />
          </div>
        </div>
      </template>

      <template v-if="!loading1 && !loading2">
        <div
          class="subject-header"
          :style="{
            backgroundImage:
              'url(' +
              subject.cover_image.host +
              subject.cover_image.access_url +
              ')',
          }"
        >
          <div class="subject-header-box">
            <div class="subject-header-title">{{ subject.title }}</div>
            <div class="subject-header-desc">{{ subject.description }}</div>
          </div>
        </div>
        <div class="content">
          <div class="subject-posts">
            <div class="subject-posts-info">
              <div class="subject-post-info-box">
                <div class="info-item">
                  <div class="info-item-number">
                    <span class="number">{{ posts.length }}</span> 篇
                  </div>
                  <div class="info-item-text">文 章</div>
                </div>
                <div class="info-item">
                  <div class="info-item-number">
                    <span class="number">{{ subject.views }}</span> 次
                  </div>
                  <div class="info-item-text">阅 读</div>
                </div>
              </div>
            </div>

            <PostCard v-for="post in posts" :post="post" :key="post.id" />

            <Pagination
              :pageCount="total_page"
              :pagerCount="13"
              @change="handlePageChange"
            />
            <br />
          </div>
        </div>
      </template>
    </el-skeleton>
  </div>
</template>

<script>
import PostCard from "@/components/PostCard.vue";
import Pagination from "@/components/Pagination.vue";
import { getSubjectById, getPostBySubjectId } from "@/api/web/subject.js";

export default {
  components: {
    PostCard,
    Pagination,
  },
  data() {
    return {
      id: "",
      page_no: 1,
      page_size: 10,
      total_rows: 0,
      total_page: 0,
      posts: [],
      subject: {},
      loading1: true,
      loading2: true
    };
  },
  beforeRouteUpdate(to, from, next) {
    this.page_no = to.query.page - 1;
    this.id = to.params.id;
    this.fetchSubjectId();
    this.fetchSubjectPosts();
    next();
  },
  mounted() {
    this.id = this.$route.params.id;
    this.page_no = this.$route.query.page || 1;
    this.fetchSubjectId();
    this.fetchSubjectPosts();
  },
  methods: {
    fetchSubjectId() {
      getSubjectById(this.id)
        .then((res) => {
          this.loading1 = false
          this.subject = res.data;
          // getFileById(res.data.cover_image).then(res => {
          //   this.cover_image = 'http://localhost:8000/' + res.data.access_url
          // })
        })
        .catch((err) => {
          console.log(err);
        });
    },
    fetchSubjectPosts() {
      getPostBySubjectId(this.id, this.page_no, this.page_size)
        .then((res) => {
          console.log(res)
          this.loading2 = false
          if (res.code === 200 && res.data) {
            this.posts = res.data.list;
            this.page_no = res.data.page_no;
            this.page_size = res.data.page_size;
            this.total_page = res.data.total_page;
            this.total_rows = res.data.total_rows;
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
    handlePageChange(current) {
      this.pageNo = current;
      this.$router.push({ path: "", query: { page: this.pageNo } });
    },
  },
};
</script>

<style lang="scss" scoped>
.subject-detail {
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
