<template>
  <div class="detail fadeInUp">
    <div v-if="!loading">
      <div class="post-title">
        <span class="post-title-text">
          <h1>{{ post.title }}</h1>
        </span>
      </div>
      <div class="post-info">
        <span class="post-created" v-if="post.updated_at"
          >更新于:
          {{ post.updated_at | momentfmt("YYYY-MM-DD HH:mm:ss") }}
          </span
        >
        <div class="devider"></div>
        <span class="post-created" v-if="post.published_at"
          >发表于:
          {{ post.published_at | momentfmt("YYYY-MM-DD HH:mm:ss") }}
       </span
        >
        <div class="devider"></div>
        <div
          v-for="tag in tags"
          :key="tag.id"
          class="post-tag"
          @click="goTag(tag.id)"
        >
          {{ tag.name }}
          <div class="devider"></div>
        </div>
        
        <div class="post-created">阅读量： {{ post.views }}</div>
      </div>

      <div
        id="post-content markdown-body"
        v-html="html"
      ></div>


    </div>


  
  </div>
</template>

<script>
// import VditorPreview from "vditor/dist/method.min";
// import "@/components/Vditor/css/index.scss";
// import "github-markdown-css/github-markdown.css";
import "katex/dist/katex.min.css";
import 'highlight.js/styles/monokai-sublime.css';

import MarkdownIt from 'markdown-it'
import hljs from 'markdown-it-highlightjs'
import latex from 'markdown-it-katex'
import { getPost } from "@/api/web/post.js";

export default {
  name: "Detail",
  data() {
    return {
      id: "",
      head: null,
      repository: null,
      history: null,
      tags: [],
      loading: false,
    };
  },
  created() {
    this.md = new MarkdownIt()
    this.md.use(hljs).use(latex)
  },
  watch: {
    $route: "fetchPost",
  },
  mounted() {
    this.fetchPost();
  },
  computed: {
    html: function() {
      let res = this.md.render(this.post.markdown_content)

      return '<div class="vditor-reset markdown-body ">' + res + '</div>'
    }
  },
  methods: {
    fetchPost() {
      this.post = null;
      this.loading = true;
      this.load = this.$loading();
      getPost(this.$route.params.id)
        .then((res) => {
          if (res.code === 200 && res.data) {
            console.log(res.data)
            this.post = res.data
            this.loading = false;
            this.load.close()
          }
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
  font-weight: 400;
  font-size: 16px;
  width: 60%;
  min-height: 573px;
  margin: 50px auto 110px;
  padding-left: 32px;
  box-sizing: border-box;
}

.post-title {
  align-items: center;
  margin-bottom: 20px;
  position: relative;
  font-weight: 800;
  color: #171d26;
}
.post-title-text {
  line-height: 1;
  font-size: 16px;
  text-overflow: ellipsis;
}
.post-info {
  display: flex;
  align-items: center;
  color: #888888;
  font-size: 13px;
  height: 32px;
  margin-top: 20px;
  display: flex;

  .post-created {
    height: 32px;
    border-radius: 4px;
    line-height: 30px;
    box-sizing: border-box;
  }
}
.post-info .post-tag {
  height: 32px;
  line-height: 30px;
  box-sizing: border-box;
  cursor: pointer;

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
  margin: 0 10px;
}

.fadeInUp {
  animation: fadeInUp 0.5s cubic-bezier(0.075, 0.82, 0.165, 1);
}
</style>

// <style lang="scss">
// .markdown-body {
//   // padding-top: 30px;
//   color: #606c80 !important;
//   font-size: 16px !important;
//   line-height: 2em !important;
//   box-sizing: border-box;
//   font-weight: 400;
//   margin: 8px 0 !important;

//   // .katex .vlist>span {
//   //   top: -1.263em !important;
//   // }
//   // .katex .vlist>span>span {
//   //   font-size: 0.8em;
//   // }


//   img {
//     width: 100%;
//     height: 100%;
//     padding: 0;
//     margin: 0;
//   }

//   h1,
//   h2,
//   h3,
//   h4,
//   h5,
//   h6 {
//     color: #171d26 !important;

//     &:not(first-child) {
//       margin-top: 24px;
//     }
//   }

//   h1 {
//     line-height: 2;
//     margin-bottom: 20px !important;
//   }
//   h2 {
//     line-height: 2;
//     font-weight: 800;
//     font-size: 25px;
//     margin-bottom: 20px !important;
//   }

//   p {
//     line-height: 2;
//     white-space: pre-wrap;
//     color: #171d26;
//     word-break: break-all;
//     font-size: 16px;
//     box-sizing: border-box;
//     font-weight: 400;
//     margin: 24px 0;

//     strong {
//       color: #464952;
//       font-weight: 500;
//     }
//   }
//   pre {
//     padding: 0;
//     margin: 24px auto;
//     white-space: pre-wrap;
//     word-break: break-all;
//   }

//   code {
//     font-size: 16px;
//     margin: 0;
//     white-space: pre-wrap;
//     word-break: break-all;
//   }

//   ul {
//     font-size: 16px;
//     font-weight: 400;
//     line-height: 2em;

//     li {
//       margin: 8px 0;
//     }
//   }
// }
// .hljs {
//   padding: 15px 15px;
// }

// </style>
