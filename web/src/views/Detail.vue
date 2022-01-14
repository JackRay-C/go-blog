<template>
  <div class="detail fadeInUp">
    <el-skeleton
      style="width: 100%"
      :loading="loading"
      animated
      :throttle="500"
    >
      <template slot="template" v-if="loading">
        <el-skeleton-item variant="h1" style="width: 40%" />
        <div
          style="
            display: flex;
            align-items: center;
            justify-item: space-between;
            margin-top: 16px;
            height: 16px;
          "
        >
          <el-skeleton-item
            variant="text"
            style="width: 80px; margin-right: 16px; height: 30px"
          />

          <el-skeleton-item
            variant="text"
            style="width: 80px; margin-right: 16px; height: 30px"
          />

          <el-skeleton-item
            variant="text"
            style="width: 80px; margin-right: 16px; height: 30px"
          />
        </div>
        <el-skeleton-item variant="text" style="100%" />
        <el-skeleton-item variant="text" style="100%" />
        <el-skeleton-item variant="text" style="100%" />
        <el-skeleton-item variant="text" style="100%" />
        <el-skeleton-item variant="text" style="100%" />
        <el-skeleton-item variant="text" style="100%" />
      </template>
      <template v-if="!loading">
        <div class="post-title">
          <span class="post-title-text">
            <h1>{{ post.title }}</h1>
          </span>
        </div>
        <div class="post-info">
          <span class="post-created" v-if="post.updated_at"
            >更新于:
            {{ post.updated_at | momentfmt("YYYY-MM-DD HH:mm:ss") }}
            &nbsp;&nbsp;</span
          >
          <span class="post-created" v-else
            >发表于:
            {{ post.created_at | momentfmt("YYYY-MM-DD HH:mm:ss") }}
            &nbsp;&nbsp;</span
          >
          <div class="devider"></div>
          <div
            v-for="tag in tags"
            :key="tag.id"
            class="post-tag"
            @click="goTag(tag.id)"
          >
            {{ tag.name }}
          </div>
          <div class="devider"></div>
          <div class="post-tag">阅读量： {{ post.views }}</div>
        </div>
        <!-- <div
          class="post-content markdown-body"
          v-highlight
          v-html="post.html_content"
        ></div> -->
        <div id="post-content" class="post-content markdown-body"></div>
      </template>
    </el-skeleton>
  </div>
</template>

<script>
import VditorPreview from "vditor/dist/method.min";
import "@/views/admin/vditor/index.scss";
import { getPost } from "@/api/web/post.js";

export default {
  name: "Detail",
  data() {
    return {
      id: "",
      post: {
        title: "",
        markdown_content: "",
        html_content: "",
        created_at: "",
        updated_at: "",
        views: 0,
        likes: 0,
      },
      tags: [],
      loading: true,
    };
  },
  // created() {
  //   this.fetchPost();
  // },
  watch: {
    $route: "fetchPost",
    loading: "preview"
  },
  mounted() {
    this.fetchPost();
  },
  methods: {
    fetchPost() {
      this.post = null;
      this.loading = true;
      getPost(this.$route.params.id)
        .then((res) => {
          console.log(res);
          if (res.code === 200 && res.data) {
            this.post = res.data;
            this.loading = false;
            console.log("获取数据完成");
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
    preview() {
      console.log(document)
      console.log()
      console.log(document.getElementById("post-content"));
      VditorPreview.preview(
        document.getElementById("post-content"),
        this.post.markdown_content,
        {
          mode: "light",
          anchor: 0,
          hljs: {
            enable: true,
            lineNumber: true,
            style: "monokai",
          },
          speech: {
            enable: false,
          },
        }
      );
    },
    goTag(id) {
      this.$router.push(`/tag/${id}`);
    },
  },
};
</script>

<style lang="scss" scoped>
.detail {
  font-weight: 400;
  font-size: 16px;
  width: 60%;
  min-height: 573px;
  margin: 70px auto 110px;
  padding-left: 32px;
  box-sizing: border-box;
}

.post-title {
  align-items: center;
  margin-bottom: 20px;
  position: relative;
  font-size: 30px;
  font-weight: 800;
  color: #171d26;
}
.post-title-text {
  vertical-align: middle;
  font-size: 1.2rem;
}
.post-info {
  display: flex;
  align-items: center;
  color: #9aa8b6;
  font-size: 16px;
  height: 40px;
  margin-top: 20px;

  .post-created {
    height: 32px;
    // background: #f6f7fa;
    color: #888888;
    border-radius: 4px;
    margin: 0 9px 0 0;
    padding: 0px 10px 0 0;
    line-height: 30px;
    font-size: 13px;
    cursor: pointer;
    box-sizing: border-box;
  }
}
.post-info .post-tag {
  height: 32px;
  // background: #f6f7fa;
  color: #888888;
  border-radius: 4px;
  margin: 0 9px;
  padding: 0px 10px;
  line-height: 30px;
  font-size: 13px;
  cursor: pointer;
  box-sizing: border-box;

  &:hover {
    color: #0c64e9;
    text-decoration: #171d26;
  }
}
.devider {
  display: inline-block;
  width: 1px;
  height: 12px;
  background-color: #c3c6cb;
}

.fadeInUp {
  animation: fadeInUp 0.5s cubic-bezier(0.075, 0.82, 0.165, 1);
}

@keyframes fadeInUp {
  0% {
    opacity: 0;
    transform: translate(0, 15%, 0);
    -webkit-transform: translate3d(0, 15%, 0);
    -ms-transform: translate3d(0, 15%, 0);
  }

  100% {
    opacity: 1;
    -ms-transform: none;
    -webkit-transform: none;
    transform: none;
  }
}
</style>

<style lang="scss">
.markdown-body {
  // padding-top: 30px;
  color: #606c80 !important;
  font-size: 16px !important;
  line-height: 2em !important;
  box-sizing: border-box;
  font-weight: 400;
  margin: 8px 0 !important;

  img {
    width: 100%;
    height: 100%;
    padding: 0;
    margin: 0;
  }

  h1,
  h2,
  h3,
  h4,
  h5,
  h6 {
    color: #171d26 !important;

    &:not(first-child) {
      margin-top: 24px;
    }
  }

  h1 {
    line-height: 2;
    margin-bottom: 20px !important;
  }
  h2 {
    line-height: 2;
    font-weight: 800;
    font-size: 25px;
    margin-bottom: 20px !important;
  }

  p {
    line-height: 2;
    white-space: pre-wrap;
    color: #171d26;
    word-break: break-all;
    font-size: 16px;
    box-sizing: border-box;
    font-weight: 400;
    margin: 24px 0;

    strong {
      color: #464952;
      font-weight: 500;
    }
  }
  pre {
    padding: 0;
    margin: 24px auto;
    white-space: pre-wrap;
    word-break: break-all;
  }

  code {
    font-size: 16px;
    margin: 0;
    white-space: pre-wrap;
    word-break: break-all;
  }

  ul {
    font-size: 16px;
    font-weight: 400;
    line-height: 2em;

    li {
      margin: 8px 0;
    }
  }
}
.hljs {
  padding: 15px 15px;
}
</style>
