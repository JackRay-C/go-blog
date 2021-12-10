<template>
  <div class="container">
    <div class="container-header">
      <header class="header-content">
        <h2 class="header-title">
          <router-link to="/admin/posts">Posts</router-link>
        </h2>
        <section class="view-actions">
          <div class="contentfilter">
            <div
              class="contentfilter-menu"
              v-for="(dropdown, index) in dropdowns"
              :key="index"
              :class="{ selected: dropdown.selected.key != 0 }"
            >
              <div
                class="dropdown-trigger"
                @click="dropdown.show = !dropdown.show"
              >
                <span class="select-selected-item">
                  <span v-show="dropdown.prev">{{ dropdown.prev }} : </span>
                  <span>{{ dropdown.selected.value }}</span>
                </span>
                <svg viewBox="0 0 26 17">
                  <path
                    d="M.469.18l11.5 13.143L23.469.18"
                    transform="translate(1 2)"
                    stroke-width="3"
                    stroke="#0B0B0A"
                    fill="none"
                    fill-rule="evenodd"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  ></path>
                </svg>
              </div>
              <div class="dropdown-content" v-show="dropdown.show">
                <ul class="dropdown-options">
                  <li
                    class="dropdown-options-item"
                    :class="{
                      selected: dropdown.selected.key == item.key,
                    }"
                    @click="
                      dropdown.selected = item;
                      dropdown.show = !dropdown.show;
                      fetchPosts()
                    "
                    v-for="item in dropdown.options"
                    :key="item.key"
                  >
                    {{ item.value }}
                  </li>
                </ul>
              </div>
            </div>
          </div>
          <router-link to="/admin/edit" class="view-action-btn">
            <span>New Post</span>
          </router-link>
        </section>
      </header>
    </div>

    <div class="container-content">
      <section class="content-list">
        <ol class="post-list">
          <li class="post-list-item header">
            <div class="list-header">
              Title
            </div>
            <div class="list-header">
              Created
            </div>
            <div class="list-header">
              Status
            </div>
          </li>
     
          <li class="post-list-item" v-for="post in posts" :key="post.id">
            <router-link
              :to="`/admin/edit/${post.id}`"
              class="post-list-data post-list-title"
            >
              <h3 class="title">{{ post.title }}</h3>
              <p>
                <span class="meta">
                  by
                  <span class="meta-author">
                    {{ post.user_id }}
                  </span>
                  •
                  <span class="meta-date">
                    {{ post.updated_at | momentfmt("YYYY-MM-DD HH:mm:ss") }}
                  </span>
                </span>
              </p>
            </router-link>
            <router-link
              class="post-list-data post-list-created"
              :to="`/admin/edit/${post.id}`"
            >
              <div>{{ post.created_at | datefmt("YYYY-MM-DD") }}</div>
              <div class="moment">
                {{ post.created_at | momentfmt("YYYY-MM-DD HH:mm:ss") }}
              </div>
            </router-link>
            <router-link
              class="post-list-data post-list-status"
              :to="`/admin/edit/${post.id}`"
            >
              <div class="status">
                  
                <span v-if="post.status === 2" class="status-published">published</span>
                <span v-else class="status-draft">draft</span>
              </div>
            </router-link>
          </li>
        </ol>
        <pagination
          :pageCount="total_page"
          :pagerCount="13"
          :current="parseInt(page_no)"
          @change="handlePageChange"
        />
      </section>
    </div>
  </div>
</template>

<script>
import Pagination from "@/components/admin/Pagination.vue";
import { listPosts } from "@/api/admin/post.js";

export default {
  components: {
    pagination: Pagination,
  },
  computed: {},
  data() {
    return {
      total_page: 0,
      page_no: 1,
      total_rows: 0,
      posts: [],
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
          onSelected: ()=> {
            
          }
        },
        {
          show: false,
          options: [
            { key: 0, value: "All subjects" },
            { key: 1, value: "Go" },
            { key: 2, value: "Python" },
          ],
          default: 0,
          selected: { key: 0, value: "All subjects" },
        },
        {
          show: false,
          options: [
            { key: 0, value: "All authors" },
            { key: 1, value: "aaa" },
            { key: 2, value: "bbb" },
          ],
          default: 0,
          selected: { key: 0, value: "All authors" },
        },
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
            { key: 0, value: 7 },
            { key: 1, value: 12 },
            { key: 2, value: 36 },
            { key: 3, value: 48 },
          ],
          prev: "Page size",
          default: 0,
          selected: { key: 0, value: 7 },
        },
      ],
    };
  },
  beforeRouteUpdate(to, from, next) {
    this.page_no = to.query.page || 1
    this.fetchPosts();
    next()
  },
  mounted() {
    this.page_no = this.$route.query.page  || 1
    this.fetchPosts();
  },
  methods: {
    fetchPosts() {
      listPosts({ page_no: this.page_no, page_size: this.dropdowns[4].selected.value }).then(
        (res) => {
            console.log(res)
          this.posts = res.data.list;
          this.total_rows = res.data.total_rows;
          this.page_no = res.data.page_no;
          this.total_page = res.data.total_page;
        }
      );
    },
    handlePageChange(current) {
      this.page_no = current;
      this.$router.push({path: "/admin/posts", query: {page: this.page_no}})
    },
  },
};
</script>

<style lang="scss">
.container {
  width: 100%;
  position: relative;
  flex-grow: 1;
  padding: 0 48px 48px 48px;
  margin: 0 auto;

  .container-header {
    margin: 0 -48px;
    padding: 0 48px;
    position: sticky;
    top: 0;
    background: #ffffff;
    z-index: 700;
    border-bottom: 1px solid #edeeef;
  

    .header-content {
      height: 95px;
      position: relative;
      flex-shrink: 0;
      display: flex;
      justify-content: space-between;
      align-items: center;

      .header-title {
        display: flex;
        align-items: center;
        overflow: hidden;
        margin: -3px 0 0;
        padding: 0;
        text-overflow: ellipsis;
        white-space: nowrap;
        font-size: 3.1rem;
        line-height: 1.3em;
        font-weight: 700;
        letter-spacing: 0;
        min-height: 35px;
        color: #15171a;

        a {
          display: flex;
          align-items: center;
          overflow: hidden;
          margin: -3px 0 0;
          padding: 0;
          text-overflow: ellipsis;
          white-space: nowrap;
          font-size: 3.1rem;
          line-height: 1.3em;
          font-weight: 700;
          letter-spacing: 0;
          min-height: 35px;
          color: #15171a;
          outline: 0;
        }
      }
    }
  }
}
</style>

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
  padding-bottom: 32px;
  margin-top: 32px;

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
        background: #f1f3f4;
      }
    }

    a {
      color: #15171a;
    }

    .post-list-title {
      h3 {
        margin: 0 0 10px;
        font-weight: 600;
        font-size: 1.6rem;
        color: #15171a;
        line-height: 1.3em;
        text-rendering: optimizeLegibility;
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
      .list-header {
        border-bottom: 1px solid #e6e9eb;
        // border-top: 1px solid #e6e9eb;
        font-size: 1.1rem;
        font-weight: 500;
        letter-spacing: 0.1px;
        color: #15171a;
        padding: 16px 20px;
        line-height: 1.2em;
        text-transform: uppercase;
        white-space: nowrap;
        display: table-cell;
        vertical-align: middle;
      }
    }
  }
}

.post-list-data.post-list-status {
  .status {
    display: flex;
    align-items: center;

    .status-draft {
      color: #fb2d8d;
      background: #f1f3f4;
      display: flex;
      justify-content: center;
      align-items: center;
      height: 20px;
      padding: 0 9px;
      font-size: 1.2rem;
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
      color: #37e227;
      background: #f1f3f4;
      text-transform: uppercase;
      font-size: 1.2rem;
      font-weight: 500;
      border-radius: 999px;
      display: flex;
      justify-content: center;
      align-items: center;
      height: 20px;
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
