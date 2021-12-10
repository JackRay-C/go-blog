<template>
  <div class="detail fadeInUp">
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

      <div class="post-tag"><i class="el-icon-view"> </i> {{ post.views }}</div>
      <div class="post-tag">
        <i class="el-icon-star-off"> </i> {{ post.likes }}
      </div>

      <div class="devider"></div>
      <div
        v-for="tag in tags"
        :key="tag.id"
        class="post-tag"
        @click="goTag(tag.id)"
      >
        {{ tag.name }}
      </div>
    </div>
    <div
      class="post-content markdown-body"
      v-highlight
      v-html="post.html_content"
    ></div>
  </div>
</template>

<script>
import { getPost, listPostTags } from "@/api/post.js";
// import "mavon-editor/dist/css/index.css";
// import "github-markdown-css/github-markdown.css";
// import "katex/dist/katex.min.css";

export default {
  name: "Detail",
  data() {
    return {
      id: "",
      post: "",
      tags: [],
    };
  },
  beforeRouteUpdate(to, from, next) {
    this.id = to.params.id;
    this.fetchPost();
    this.fetchTags();
    next();
  },
  mounted() {
    this.id = this.$route.params.id;
    this.fetchPost();
    this.fetchTags();
    window.addEventListener("scroll", this.scrollTop);
  },
  methods: {
    fetchPost() {
      getPost(this.id).then((res) => {
        if (res.code === 200 && res.data) {
          this.post = res.data;
        }
      });
    },
    fetchTags() {
      listPostTags(this.id)
        .then((res) => {
          console.log(res);
          this.tags = res.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    goTag(id) {
      this.$router.push(`/tag/${id}`);
    },
  },
};
</script>

<style lang="scss" scoped>
.detail {
  margin: 0;
  padding: 0;
  font-weight: 400;
  font-size: 16px;
  width: 60%;
  min-height: 573px;
  margin: 60px auto 110px;
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
    margin-right: 8px;
    padding: 0px 10px;
    line-height: 30px;
    font-size: 13px;
    cursor: pointer;
    box-sizing: border-box;
    // border: 1px solid #888888;
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
  // border: 1px solid #888888;
}
.devider {
  display: inline-block;
  width: 1px;
  height: 12px;
  background-color: #c3c6cb;
  margin: 0 8px;
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
  padding-top: 30px;
  color: #606c80 !important;
  font-size: 16px !important;
  line-height: 2em !important;
  box-sizing: border-box;
  font-weight: 400;
  margin: 8px 0 !important;

  h1,
  h2,
  h3,
  h4,
  h5,
  h6 {
    color: #171d26 !important;
    margin-top: 24px;
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
    line-height: 2em;
    white-space: pre-wrap;
    word-break: break-all;
    font-size: 16px;
    box-sizing: border-box;
    font-weight: 400;
    margin: 16px 0;

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