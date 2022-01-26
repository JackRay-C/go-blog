<template>
  <div class="edit" v-loading="loading">
    <div class="header">
      <div class="edit-header-left">
        <div class="back">
          <el-button
            class="back-text"
            @click="goback"
            type="text"
            icon="el-icon-arrow-left"
            >返回</el-button
          >
        </div>
        <div class="title">
          <span class="title-text" v-show="!post_title_edit">{{
            post.title
          }}</span>

          <el-input
            class="title-text"
            v-show="post_title_edit"
            v-model="post.title"
            @blur="post_title_edit = !post_title_edit"
          ></el-input>

          <i
            :class="{
              'el-icon-edit': !post_title_edit,
              'el-icon-check': post_title_edit,
            }"
            @click="post_title_edit = !post_title_edit"
          ></i>

          <span v-if="post.id" class="date-text"
            >创建于{{
              post.created_at | momentfmt("YYYY-MM-DD HH:mm:ss")
            }}</span
          >
        </div>
      </div>

      <div class="edit-header-right">
        <el-button-group class="button-right">
          <el-button size="medium" plain @click="stagedPost"
            >保存草稿</el-button
          >
          <el-button size="medium" plain @click="commitPost">发布博客</el-button>
          <el-button size="medium" plain @click="drawer = !drawer"
            >博客设置</el-button
          >
        </el-button-group>
        <!-- <el-avatar
          class="avatar"
          size="medium"
          v-if="post.user"
          :src="post.user.avatar.host + post.user.avatar.access_url"
        ></el-avatar> -->

        <!-- <el-avatar
          class="avatar"
          size="medium"
          src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"
        ></el-avatar> -->
      </div>
    </div>

    <el-drawer title="我是标题" :visible.sync="drawer" direction="rtl">
      <span>我来啦!</span>
    </el-drawer>

    <div class="content">
      <div id="editor" class="editor-content"></div>
    </div>
  </div>
</template>

<script>
import Vditor from "vditor";
import "@/components/Vditor/css/index.scss";
import debounce from "lodash.debounce";
import { getPost, addPost, putPost } from "@/api/admin/post";
import api from "@/api/admin/api";

export default {  
  name: "editor",
  data() {
    return {
      editor: "",
      post_title_edit: false,
      placeholder: "请输入内容",
      drawer: false,
      img_ids: [],
      loading: false,
      post: {},
    };
  },
  created() {
    this.initEditor();
    this.fetchPost();
  },
  watch: {
    $route: "fetchPost",
  },
  computed: {},
  methods: {
    initEditor() {
      this.$nextTick(() => {
        this.editor = new Vditor("editor", {
          toolbar: [
            "|",
            {
              name: "headings",
              hotkey: "⌘H",
              icon: '<svg><use xlink:href="#vditor-icon-headings"></use></svg>',
              tipPosition: "w",
            },
            {
              name: "quote",
              icon: '<svg viewBox="0 0 1024 1024"><path d="M894.6 907.1H605.4c-32.6 0-59-26.4-59-59V608.2l-4-14.9c0-315.9 125.5-485.1 376.5-507.5v59.8C752.7 180.4 711.3 315.8 711.3 442.4v41.2l31.5 12.3h151.8c32.6 0 59 26.4 59 59v293.2c0 32.5-26.4 59-59 59z m-472 0H133.4c-32.6 0-59-26.4-59-59V608.2l-4-14.9c0-315.9 125.5-485.1 376.5-507.5v59.8C280.7 180.4 239.3 315.8 239.3 442.4v41.2l31.5 12.3h151.8c32.6 0 59 26.4 59 59v293.2c0 32.5-26.4 59-59 59z"></path></svg>',
              prefix: "> ",
              tipPosition: "s",
              hotkey: "⌘;",
            },
            "|",
            {
              name: "bold",
              tipPosition: "s",
              hotkey: "⌘B",
              icon: '<svg viewBox="0 0 1024 1024"><path d="M707.872 484.64A254.88 254.88 0 0 0 768 320c0-141.152-114.848-256-256-256H192v896h384c141.152 0 256-114.848 256-256a256.096 256.096 0 0 0-124.128-219.36zM384 192h101.504c55.968 0 101.504 57.408 101.504 128s-45.536 128-101.504 128H384V192z m159.008 640H384v-256h159.008c58.464 0 106.016 57.408 106.016 128s-47.552 128-106.016 128z"></path></svg>',
              prefix: "**",
              suffix: "**",
            },
            {
              name: "outline",
              tipPosition: "s",
              icon: '<svg viewBox="0 0 1024 1024"><path d="M704 64l128 0 0 416c0 159.072-143.264 288-320 288s-320-128.928-320-288l0-416 128 0 0 416c0 40.16 18.24 78.688 51.36 108.512 36.896 33.216 86.848 51.488 140.64 51.488s103.744-18.304 140.64-51.488c33.12-29.792 51.36-68.352 51.36-108.512l0-416zM192 832l640 0 0 128-640 0z"></path></svg>',
            },
            {
              name: "italic",
              tipPosition: "s",
              icon: '<svg viewBox="0 0 1024 1024"><path d="M896 64v64h-128L448 896h128v64H128v-64h128L576 128h-128V64z"></path></svg>',
            },
            {
              name: "strike",
              hotkey: "⌘D",
              icon: '<svg viewBox="0 0 1024 1024"><path d="M1024 512v64h-234.496c27.52 38.496 42.496 82.688 42.496 128 0 70.88-36.672 139.04-100.576 186.976C672.064 935.488 594.144 960 512 960s-160.064-24.512-219.424-69.024C228.64 843.04 192 774.88 192 704h128c0 69.376 87.936 128 192 128s192-58.624 192-128-87.936-128-192-128H0v-64h299.52a385.984 385.984 0 0 1-6.944-5.024C228.64 459.04 192 390.88 192 320s36.672-139.04 100.576-186.976C351.936 88.512 429.856 64 512 64s160.064 24.512 219.424 69.024C795.328 180.96 832 249.12 832 320h-128c0-69.376-87.936-128-192-128s-192 58.624-192 128 87.936 128 192 128c78.976 0 154.048 22.688 212.48 64H1024z"></path></svg>',
              tipPosition: "e",
              tip: "删除线",
              prefix: "~~",
              suffix: "~~",
            },
            {
              name: "list",
              tipPosition: "se",
              prefix: "* ",
              hotkey: "⌘L",
            },
            {
              name: "ordered-list",
              tipPosition: "se",
              prefix: "1. ",
              hotkey: "⌘O",
            },

            "|",
            {
              name: "link",
              icon: '<svg viewBox="0 0 1024 1024"><path d="M440.224 635.776a51.84 51.84 0 0 1-36.768-15.232c-95.136-95.136-95.136-249.92 0-345.056l192-192C641.536 37.408 702.816 12.032 768 12.032s126.432 25.376 172.544 71.456c95.136 95.136 95.136 249.92 0 345.056l-87.776 87.776a51.968 51.968 0 1 1-73.536-73.536l87.776-87.776a140.16 140.16 0 0 0 0-197.984c-26.432-26.432-61.6-40.992-99.008-40.992s-72.544 14.56-99.008 40.992l-192 192a140.16 140.16 0 0 0 0 197.984 51.968 51.968 0 0 1-36.768 88.768z"></path><path d="M256 1012a242.4 242.4 0 0 1-172.544-71.456c-95.136-95.136-95.136-249.92 0-345.056l87.776-87.776a51.968 51.968 0 1 1 73.536 73.536l-87.776 87.776a140.16 140.16 0 0 0 0 197.984c26.432 26.432 61.6 40.992 99.008 40.992s72.544-14.56 99.008-40.992l192-192a140.16 140.16 0 0 0 0-197.984 51.968 51.968 0 1 1 73.536-73.536c95.136 95.136 95.136 249.92 0 345.056l-192 192A242.4 242.4 0 0 1 256 1012z"></path></svg>',
              tipPosition: "se",
            },
            {
              name: "upload",
              icon: '<svg viewBox="0 0 1024 1024"><path d="M959.877 128l0.123 0.123v767.775l-0.123 0.122H64.102l-0.122-0.122V128.123l0.122-0.123h895.775zM960 64H64C28.795 64 0 92.795 0 128v768c0 35.205 28.795 64 64 64h896c35.205 0 64-28.795 64-64V128c0-35.205-28.795-64-64-64zM832 288.01c0 53.023-42.988 96.01-96.01 96.01s-96.01-42.987-96.01-96.01S682.967 192 735.99 192 832 234.988 832 288.01zM896 832H128V704l224.01-384 256 320h64l224.01-192z"></path></svg>',
              tipPosition: "se",
            },
            "record",
            "table",
            "|",
            {
              name: "inline-code",
              tipPosition: "se",
            },
            {
              hotkey: "⌘U",
              icon: '<svg><use xlink:href="#vditor-icon-code"></use></svg>',
              name: "code",
              prefix: "```",
              suffix: "\n```",
              tipPosition: "se",
            },
            {
              icon: '<svg><use xlink:href="#vditor-icon-code-theme"></use></svg>',
              name: "code-theme",
              tipPosition: "w",
            },
            {
              name: "line",
              tipPosition: "se",
              prefix: "---",
            },
            {
              name: "emoji",
              icon: '<svg viewBox="0 0 1024 1024"><path d="M512 1024C230.4 1024 0 793.6 0 512S230.4 0 512 0s512 230.4 512 512-230.4 512-512 512z m0-102.4c226.742857 0 409.6-182.857143 409.6-409.6S738.742857 102.4 512 102.4 102.4 285.257143 102.4 512s182.857143 409.6 409.6 409.6z m-204.8-358.4h409.6c0 113.371429-91.428571 204.8-204.8 204.8s-204.8-91.428571-204.8-204.8z m0-102.4c-43.885714 0-76.8-32.914286-76.8-76.8s32.914286-76.8 76.8-76.8 76.8 32.914286 76.8 76.8-32.914286 76.8-76.8 76.8z m409.6 0c-43.885714 0-76.8-32.914286-76.8-76.8s32.914286-76.8 76.8-76.8c43.885714 0 76.8 32.914286 76.8 76.8s-32.914286 76.8-76.8 76.8z"></path></svg>',
              tipPosition: "se",
            },
            "|",
            {
              name: "undo",
              icon: '<svg viewBox="0 0 1024 1024"><path d="M512 64A510.272 510.272 0 0 0 149.984 213.984L0.032 64v384h384L240.512 304.48A382.784 382.784 0 0 1 512.032 192c212.064 0 384 171.936 384 384 0 114.688-50.304 217.632-130.016 288l84.672 96a510.72 510.72 0 0 0 173.344-384c0-282.784-229.216-512-512-512z"></path></svg>',
              tipPosition: "se",
            },
            {
              name: "redo",
              icon: '<svg viewBox="0 0 1024 1024"><path d="M0.00032 576a510.72 510.72 0 0 0 173.344 384l84.672-96A383.136 383.136 0 0 1 128.00032 576C128.00032 363.936 299.93632 192 512.00032 192c106.048 0 202.048 42.976 271.52 112.48L640.00032 448h384V64l-149.984 149.984A510.272 510.272 0 0 0 512.00032 64C229.21632 64 0.00032 293.216 0.00032 576z"></path></svg>',
              tipPosition: "se",
            },
            {
              name: "more",
              tipPosition: "w",
              toolbar: ["edit-mode", "content-theme", "export", "preview"],
            },
            "|",
            // "help",
            // "edit-mode",
            // "content-theme",
            // "export",
            // "preview",
            // "|",
          ],
          toolbarConfig: {
            pin: true,
          },
          cache: {
            enable: false,
          },
          placeholder: this.placeholder,
          tab: "\t",
          typewriterMode: true,
          mode: "ir",
          input: debounce(this.onChange, 500),
          blur: debounce(this.onBlur, 500),
          esc: () => {},
          after: () => {
            this.editor.setValue(this.post.markdown_content);
            this.loading = false;
            console.log("vditor init success");
          },
          preview: {
            hljs: {
              enable: true,
              lineNumber: true,
              style: "monokai",
            },
          },
          upload: {
            url: process.env.VUE_APP_BASE_URL + api.files,
            max: 10 * 1024 * 1024,
            linkToImgUrl: "",
            linkToImgCallback: (res) => {
              console.log(res);
            },
            linkToImgFormat: (res) => {
              console.log(res);
            },
            headers: {
              token: localStorage.getItem("token"),
            },
            withCredentials: true,
            fieldName: "file",
            multiple: false,
            format: (files, res) => {
              var resp = JSON.parse(res);
              if (resp.code === 200) {
                return JSON.stringify({
                  msg: "",
                  code: 0,
                  data: {
                    errFiles: [],
                    succMap: {
                      [files[0].name]:
                        process.env.VUE_APP_BASE_URL +
                        `/${resp.data.access_url}`,
                    },
                  },
                });
              }
            },
          },
        });
      });
    },
    fetchPost() {
      if (this.$route.params.id && /^\d+$/.test(this.$route.params.id)) {
        this.pullPost();
      } else {
        this.initPost();
      }
    },
    pushPost() {
      if (this.$route.params.id && /^\d+$/.test(this.$route.params.id)) {
        // 如果有route参数，update
        putPost(this.$route.params.id, this.post)
          .then((res) => {
            if (res.code === 200) {
              this.post = res.data
              this.$notify.success({title: 'Success', message: '发布博客成功. '})
            }
          })
          .catch((err) => {
            this.$notify.success({title: 'Failed', message: '发布博客失败: ' + err})
            console.log(err);
          });
      } else {
        console.log("id: " + this.$route.params.id)
      }
    },
    onChange() {
      this.post.html_content = this.editor.getHTML();
      this.post.markdown_content = this.editor.getValue();
      this.post.description = this.editor.substring(0, 300);
    },
    onBlur(val) {
      // 保存 或保存到草稿箱,如果是新建 跳转到编辑页面
      this.description = val.substring(0, 300);
      console.log("onBlur", val.substring(0, 300));
      this.pushPost();
    },
    initPost() {
      // 初始化博客
      this.$loading()
      this.post = {
        title: "新建博客",
        markdown_content: "",
        status: 1,
        visibility: 1,
        tags: [],
        created_at: new Date().toJSON(),
      };
      addPost({ ...this.post })
        .then((res) => {
          console.log(res);
          if (res.code === 200) {
            this.post.id = res.data.id;
            this.$loading.close()
            this.$router.push("/admin/edit/" + res.data.id);
          } else {
            this.$notify({
              title: "Error " + res.code,
              message: res.message,
            });
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
    pullPost() {
      // 请求博客
      this.loading = true
      getPost(this.$route.params.id).then((res) => {
        this.loading = false;
        if (res.code === 200) {
          this.post = res.data;
        } else {
          this.$alert(res.message, "Error " + res.code, {
            confirmButtonText: "确定",
            callback: () => {
              this.$router.go(-1);
            },
          });
        }
      });
    },
    stagedPost() {
      // 保存草稿
      this.post.status = 1;
      this.pushPost();
    },
    commitPost() {
      // 发布博客
      this.post.status = 2;
      this.post.visibility = 2;
      this.pushPost();
    },
    goback() {
      localStorage.setItem("draft", this.getEditorPost);
      this.$router.back(-1);
    },
  },
};
</script>

<style lang="scss" scoped>
.edit {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
}

.header {
  width: 100%;
  height: 56px;
  background: #ffffff;
  position: relative;
  z-index: 1000;

  display: flex;
  flex-direction: row;
  align-items: center;

  .edit-header-left {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    flex: 1;
    height: 56px;

    .back {
      margin-left: 36px;
      margin-right: 20px;
      width: 56px;

      .back-text {
        font-size: 16px;
        color: #606266;
      }
    }

    .title {
      display: flex;
      flex-direction: row;
      align-items: center;
      font-size: 16px;
      color: #606266;
      min-width: 180px;
      max-width: 500px;

      .title-text {
        overflow: hidden;
        -o-text-overflow: ellipsis;
        -webkit-text-overflow: ellipsis;
        white-space: nowrap;
        text-overflow: ellipsis;
        margin-right: 15px;
      }

      .date-text {
        color: #d6d3d6;
        font-size: 14px;
        margin-left: 20px;
      }
    }
  }

  .edit-header-right {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: flex-end;
    margin-right: 36px;
    height: 56px;

    .button-right {
      margin: 0 20px;
    }
  }
}

.content {
  width: 100%;
  flex: 1;
  background-color: rgb(245, 245, 245);

  .editor-content {
    height: 100%;
    // vditor scss
    .vditor-toolbar {
      line-height: 1.5;
      background: #ffffff !important;
    }
  }
}
</style>
