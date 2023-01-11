<template>
  <div class="container">
    <div class="container-header">
      <header class="header-content">
        <h2 class="header-title">
          <router-link to="/admin/posts">Posts</router-link>
        </h2>
        <section class="view-actions">
          <div class="contentfilter">
            <div class="contentfilter-menu" v-for="(dropdown, index) in dropdowns" :key="index"
              :class="{ selected: dropdown.selected.key != 0 }">
              <div class="dropdown-trigger" @click="dropdown.show = !dropdown.show">
                <span class="select-selected-item">
                  <span v-show="dropdown.prev">{{ dropdown.prev }} : </span>
                  <span>{{ dropdown.selected.value }}</span>
                </span>
                <svg viewBox="0 0 26 17">
                  <path d="M.469.18l11.5 13.143L23.469.18" transform="translate(1 2)" stroke-width="3" stroke="#0B0B0A"
                    fill="none" fill-rule="evenodd" stroke-linecap="round" stroke-linejoin="round"></path>
                </svg>
              </div>
              <div class="dropdown-content" v-show="dropdown.show">
                <ul class="dropdown-options">
                  <li class="dropdown-options-item" :class="{
                    selected: dropdown.selected.key == item.key,
                  }" @click="dropdown.selected = item;dropdown.show = !dropdown.show;fetchPosts();" v-for="item in dropdown.options" :key="item.key">
                    {{ item.value }}
                  </li>
                </ul>
              </div>
            </div>
          </div>
          <!-- 新建博客，根据博客的ID跳转到草稿编辑页面 -->
          <div to="/admin/edit" @click="newPost" class="view-action-btn">
            <span>New Post</span>
          </div>
        </section>
      </header>
    </div>

    <el-skeleton :loading="loading" animated :count="6" :throttle="500" class="container-content-flex">
      <template slot="template">
        <el-skeleton-item variant="text" style="margin-top: 32px; height: 60px">
          <section class="content-list">
            <ol class="post-list">
              <li class="post-list-item header">
                <div class="list-header">{{ $t("table.title") }}</div>
                <div class="list-header">{{ $t("table.created") }}</div>
                <div class="list-header">{{ $t("table.status") }}</div>
              </li>
            </ol>
          </section>
        </el-skeleton-item>
      </template>
      <template>
        <div class="container-content">
          <section class="content-list">
            <ol class="post-list">
              <li class="post-list-item header">
                <div class="list-header">{{ $t("table.title") }}</div>
                <div class="list-header">{{ $t("table.created") }}</div>
                <div class="list-header">{{ $t("table.status") }}</div>
              </li>

              <li class="post-list-item" v-for="post in posts" :key="post.id">
                <router-link :to="`/admin/edit/${post.id}`" class="post-list-data post-list-title">
                  <h3 class="title">{{ post.title }}</h3>
                  <p>
                    <span class="meta">
                      by
                      <span class="meta-author">
                        {{ post.user.nickname }}
                      </span>
                      •
                      <span class="meta-date">
                        {{ post.updated_at | momentfmt("YYYY-MM-DD HH:mm:ss") }}
                      </span>
                    </span>
                  </p>
                </router-link>
                <router-link class="post-list-data post-list-created" :to="`/admin/edit/${post.id}`">
                  <div>{{ post.updated_at | datefmt("YYYY-MM-DD") }}</div>
                  <div class="moment">
                    {{ post.updated_at | momentfmt("YYYY-MM-DD HH:mm:ss") }}
                  </div>
                </router-link>
                <router-link class="post-list-data post-list-status" :to="`/admin/edit/${post.id}`">
                  <div class="status">
                    <span v-if="post.status === 1" class="status-draft">{{ $t('post.status.draft') }}</span>
                    <span v-if="post.status === 2" class="status-published">{{ $t('post.status.published') }}</span>
                    <span v-if="post.status === 3" class="status-deleted">{{ $t('post.status.delete') }}</span>
                  </div>
                </router-link>
              </li>
            </ol>
          </section>
        </div>
      </template>
    </el-skeleton>
    <pagination class="container-footer"  :pageNo="parseInt(page_no)" :pageCount="total_page" :pagerCount="13" 
      @change="handlePageChange" />
  </div>
</template>

<script>
import Pagination from "@/components/admin/Pagination.vue";
import { listPosts } from "@/api/admin/post.js";
import { getUserById } from "@/api/admin/user.js";
import { addPost } from "../../api/admin/post";

export default {
  components: {
    pagination: Pagination,
  },

  data() {
    return {
      total_page: 1,
      page_no: 1,
      total_rows: 0,
      posts: [],
      loading: true,
      status: 1,
      dropdowns: [
        {
          show: false,
          options: [
            { key: 0, value: "All posts" },
            { key: 1, value: "Drafts" },
            { key: 2, value: "Published" },
          ],
          default: 0,
          selected: { key: 0, value: "All posts" },
          onSelected: () => { },
        },
        // {
        //   show: false,
        //   options: [
        //     { key: 0, value: "All subjects" },
        //     { key: 1, value: "Go" },
        //     { key: 2, value: "Python" },
        //   ],
        //   default: 0,
        //   selected: { key: 0, value: "All subjects" },
        // },
        // {
        //   show: false,
        //   options: [
        //     { key: 0, value: "All authors" },
        //     { key: 1, value: "aaa" },
        //     { key: 2, value: "bbb" },
        //   ],
        //   default: 0,
        //   selected: { key: 0, value: "All authors" },
        // },
        {
          show: false,
          options: [
            { key: 0, value: "Newest" },
            { key: 1, value: "Oldest" },
            { key: 2, value: "Updated" },
          ],
          prev: "Sort by",
          default: 0,
          selected: { key: 0, value: "Newest" },
        },
        {
          show: false,
          options: [
            { key: 0, value: 8 },
            { key: 1, value: 16 },
            { key: 2, value: 24 },
            { key: 3, value: 32 },
            { key: 4, value: 40 },
          ],
          prev: "Page size",
          default: 0,
          selected: { key: 0, value: 8 },
        },
      ],
    };
  },
  computed: {},
  watch: {
        "$route.path": (newRoute, oldRoute)=>{
          console.log(newRoute)
          console.log(oldRoute)
        }
  },
  beforeRouteUpdate(to, from, next) {
    this.page_no = to.query.page || 1;
    this.fetchPosts();
    next();
  },
  mounted() {
    if(this.$route.name === "drafts") {
      this.status = 1
    } else if (this.$route.name === "published") {
      this.status = 2
    } else {
      this.status = 0
    }
    this.page_no = this.$route.query.page || 1;
    this.fetchPosts();
  },
  methods: {
    async fetchPosts() {
      let res = await listPosts({
        page_no: this.page_no,
        status: this.status,
        page_size: this.dropdowns[2].selected.value
      });
      
      if (res.code === 200) {
        console.log(res)

        let posts = res.data.list;

        for (let i = 0; i < posts.length; i++) {
          let response = await getUserById(posts[i].user_id);
          posts[i].user = response.data;
        }
        this.posts = posts;
        this.loading = false;
        this.total_rows = res.data.total_rows;
        this.page_no = res.data.page_no;
        this.total_page = res.data.total_page;
      } else {
        this.loading = false;
        this.$notification({
          message: res.message,
        });
      }
    },
    handlePageChange(current) {
      this.page_no = current;
      console.log(current)
      this.$router.push({
        // path: "/admin/posts",
        path: this.$route.path,
        query: { page: this.page_no },
      });
    },
    newPost(){
      // 新建空博客，获取ID，跳转到编辑页面
      let post = {
        "title": "新建博客", // 博客标题
        "markdown_content": "", // 博客内容
        "status": 1, // 草稿
        "visibility": 2, // 公开
      }; 
      
      addPost(post).then((res) => {
          console.log(res);
          if (res.code === 200) {
            this.$router.push("/admin/new/" + res.data.id);
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
    getAuthor(post) {
      getUserById(post.user_id).then((res) => {
        console.log(res);
        return res.data.nickname;
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.view-actions {
  flex-shrink: 0;
  display: flex;
  align-items: center;

  .contentfilter {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin: 0 24px;
    padding: 0;
    list-style: none;

    .contentfilter-menu {
      display: block;
      box-sizing: border-box;

      .dropdown-trigger {
        background: #ffffff;
        font-size: 1.35rem;
        font-weight: 400;
        color: #394047;
        letter-spacing: 0.2px;
        padding: 6px 12px;
        margin-right: 8px;
        outline: none;
        border: 1px solid transparent;
        border-radius: 2px;
        white-space: nowrap;
        transition: all 0.25s ease;
        overflow: hidden;
        position: relative;
        text-overflow: ellipsis;
        line-height: 1.75;
        min-height: 1.75em;
        user-select: none;

        .select-selected-item {
          margin-left: 0;
          box-sizing: border-box;
        }

        svg {
          height: 4px;
          width: 6.11px;
          margin-left: 4px;
          margin-top: -2px;
          vertical-align: middle;

          &:not(:root) {
            overflow: hidden;
          }
        }

        &:after {
          content: "";
          display: table;
          clear: both;
        }
      }
    }
  }

  .selected.contentfilter-menu {
    .dropdown-trigger {
      color: #30cf43;
      font-weight: 600;
      background: #f1f3f4;
      border: 1px solid #f1f3f4;
    }
  }
}

.view-action-btn {
  display: flex;
  align-items: center;

  padding: 0 14px;
  height: 34px;
  font-size: 1.35rem;
  line-height: 34px;
  text-align: center;
  letter-spacing: 0.2px;
  border-radius: 3px;
  white-space: nowrap;
  text-overflow: ellipsis;
  cursor: pointer;
  color: #ffffff;
  background: #15171a;
  font-weight: 500;
  outline: none;
  text-decoration: none;
  user-select: none;
  transition: all 0.2s ease;

  &:hover {
    text-decoration: none;
    color: #fff;
    transition: background 0.1s, color 0.1s;
  }
}

.container-content {
  position: relative;
  flex-grow: 1;
  padding-top: 0;
  // padding-bottom: 32px;
  padding-top: 20px;
  // margin-bottom: 32px;
  // margin-top: 20px;
  // background: #fff;
  overflow-y: scroll;

  .content-list {
    display: flex;
    flex-direction: column;
    height: 100%;

    .post-list {
      flex: 1;
    }
  }

  .post-list {
    display: table;
    border-collapse: collapse;
    list-style: none;
    background: #fff;
    width: 100%;
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    // border: 1px solid #ebeef0;
    border-radius: 3px;
    border-bottom: none;

    .post-list-item {
      display: table-row;
      box-sizing: border-box;
      position: relative;

      .post-list-data {
        display: table-cell;
        vertical-align: middle;
        padding: 16px 20px;
        font-size: 1.3rem;
        text-decoration: none;
      }
    }

    li {
      margin-bottom: 10px;
      line-height: 1.4em;
      border-bottom: 1px solid #e6e9eb;

      &:not(.header):hover {
        background: #fafafa;
      }
    }

    a {
      color: #15171a;
    }

    .post-list-title {
      width: 45%;

      h3 {
        margin: 0 0 10px;
        font-weight: 400;
        font-size: 14px;
        color: #15171a;
        line-height: 1.3em;
        text-rendering: optimizeLegibility;
        width: 300px;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
      }

      p {
        margin: 2px 0 0;
        font-size: 1.3rem;
        overflow-wrap: break-word;
        color: #626d79;
        box-sizing: border-box;

        .meta {
          margin: 4px 0 0;
          font-size: 1.3rem;
          box-sizing: border-box;
          color: #abb4be;

          .meta-author {
            color: #8e9cac;
            font-weight: 500;
          }

          .meta-date {
            position: relative;

            &:before {
              visibility: hidden;
              opacity: 0;
              pointer-events: none;
              transition: all 0.2s ease;
              z-index: 9999;
              position: absolute;
              bottom: calc(100% + 4px);
              left: 50%;
              white-space: nowrap;
              border-radius: 3px;
              background-color: #394047;
              color: #fff;
              content: attr(data-tooltip);
              text-align: center;
              font-size: 1.3rem;
              font-weight: 400;
              line-height: 1.4em;
              letter-spacing: 0.2px;
              text-transform: none;
              transform: translate(-50%, 5px);
            }
          }
        }
      }
    }

    .header {
      height: 60px;

      .list-header {
        border-bottom: 1px solid #e6e9eb;
        // // border-top: 1px solid #e6e9eb;
        font-size: 17px;
        font-weight: 500;
        letter-spacing: 0.1px;
        color: #0f1011;
        // height: 30px;
        padding: 0px 20px;
        line-height: 60px;
        text-transform: uppercase;
        white-space: nowrap;
        display: table-cell;
        vertical-align: middle;
      }
    }
  }
}

.post-list-data.post-list-status {
  width: 20%;
  .status {
    display: flex;
    align-items: center;
    height: 30px;

    .status-draft {
      // color: #0c0c0c;
      // background: #f1f3f4;
      display: flex;
      justify-content: center;
      align-items: center;
      height: 30px;
      padding: 0 9px;
      line-height: 1em;
      font-weight: 500;
      letter-spacing: 0.2px;
      text-align: center;
      text-decoration: none;
      white-space: nowrap;
      user-select: none;
      border-radius: 999px;
      text-transform: uppercase;
    }

    .status-published {
      // color: #0b0c0b;
      // color: #565058;
      // background: #f2f6f8;
      text-transform: uppercase;
      font-weight: 600;
      border-radius: 999px;
      display: flex;
      justify-content: center;
      align-items: center;
      height: 30px;
      padding: 0 9px;
      white-space: nowrap;
    }

    .status-deleted {
      // color: #f10000;
      // background: #f2f6f8;
      text-transform: uppercase;
      font-weight: 600;
      border-radius: 999px;
      display: flex;
      justify-content: center;
      align-items: center;
      height: 30px;
      padding: 0 9px;
      white-space: nowrap;
    }
  }
}

.post-list-data.post-list-created {
  display: table-cell;

  font-size: 1.3rem;
  white-space: nowrap;

  div {
    margin: 0 0 10px;
  }

  .moment {
    color: #abb4be;
    margin: 2px 0 0;
  }
}

.pagination {
  width: 100%;
  background: #ffffff;
  border-bottom: 1px solid #e6e9eb;
  min-height: 68px;
}

.dropdown-content {
  width: 180px;
  margin-top: 6px;
  padding: 6px 0;
  border: none !important;
  font-size: 1.35rem;
  box-shadow: 0 0 0 1px rgb(0 0 0 / 4%), 0 7px 20px -5px rgb(0 0 0 / 15%);
  border-radius: 5px;
  position: absolute;
  z-index: 1000;
  box-sizing: border-box;
  background: #fff;
  line-height: 1.75;
  overflow: hidden;
  color: inherit;
  top: 64.5px;

  .dropdown-options {
    overflow-x: hidden;
    max-height: 50vh;
    overflow-y: auto;
    box-sizing: border-box;
    list-style: none;
    margin: 0;
    padding: 0;
    user-select: none;

    .dropdown-options-item {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      color: #394047;
      background: transparent;
      margin-bottom: 0;
      padding: 6px 14px;
      cursor: pointer;
      line-height: 1.35em;
      line-height: 1.4em;

      &.selected {
        font-weight: 700;
      }

      &:not(.selected):hover {
        background: #f4f5f5;
      }
    }
  }
}
</style>
