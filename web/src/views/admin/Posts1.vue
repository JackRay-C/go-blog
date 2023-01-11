<template>
    <div class="posts-container">
        <div class="posts-container-header">
            <header class="header-content">
                <h2 class="header-title">
                    <router-link to="/admin/posts">Posts</router-link>
                </h2>
            </header>
        </div>
        <br />

        <el-skeleton :loading="loading" animated :count="6" :throttle="500" class="container-content-flex">
            <template>
                <el-table :data="posts"  style="width: 100%">
                    <el-table-column type="selection" width="55"> </el-table-column>
                    <el-table-column prop="title" label="标题" >
                    </el-table-column>
                    <el-table-column label="日期" width="250">
                        <template slot-scope="scope">
                            <div>
                                {{ scope.row.created_at | momentfmt("YYYY-MM-DD HH:mm:ss") }}
                            </div>
                            
                        </template>
                    </el-table-column>
                    <el-table-column prop="user.nickname" label="作者" width="180"> </el-table-column>
                    <el-table-column prop="status" label="状态" width="180"> </el-table-column>
                    <el-table-column label="操作" width="180">
      <template slot-scope="scope">
        <el-button
          size="mini"
          plain
          @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
        <el-button
          size="mini"
          plain
          @click="handleDelete(scope.$index, scope.row)">删除</el-button>
      </template>
    </el-table-column>
                </el-table>
            </template>
        </el-skeleton>

        <div class="posts-container-footer">
            <!-- <el-pagination  :page-size="20" :pager-count="11" width="100%" layout="prev, pager, next" :total="1000"> -->
                <pagination class="container-footer" :pageCount="total_page" :pagerCount="13" :current="parseInt(page_no)"
      @change="handlePageChange" />
        <!-- </el-pagination> -->
        </div>
        
    </div>
</template>

<script>
    import Pagination from "@/components/admin/Pagination.vue";
import { listPosts } from "@/api/admin/post.js";
import { getUserById } from "@/api/admin/user.js";
import { addPost } from "../../api/admin/post";

export default {
    components:{
        Pagination
    },
    data() {
        return {
            total_page: 0,
            page_no: 1,
            total_rows: 0,
            posts: [],
            loading: true,
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
                        { key: 0, value: 7 },
                        { key: 1, value: 12 },
                        { key: 2, value: 36 },
                        { key: 3, value: 48 },
                    ],
                    prev: "Page size",
                    default: 0,
                    selected: { key: 0, value: 10 },
                },
            ],
        };
    },
    beforeRouteUpdate(to, from, next) {
        this.page_no = to.query.page || 1;
        this.fetchPosts();
        next();
    },
    mounted() {
        this.page_no = this.$route.query.page || 1;
        console.log(this.page_no);
        this.fetchPosts();
    },
    methods: {
        async fetchPosts() {
            let res = await listPosts({
                page_no: this.page_no,
                page_size: this.dropdowns[2].selected.value,
            });

            if (res.code === 200) {
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
            this.$router.push({
                path: "/admin/posts",
                query: { page: this.page_no },
            });
        },
        newPost() {
            // 新建空博客，获取ID，跳转到编辑页面
            let post = {
                title: "新建博客",
                markdown_content: "",
                status: 2,
                visibility: 1,
            };

            addPost(post)
                .then((res) => {
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

<style lang="scss">
.posts-container {
    width: 100%;
    position: relative;
    flex-grow: 1;
    padding: 0 48px 10px 36px;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    // background: #f7f8fa;

    background: #f6f7fa;
    .posts-container-header {
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

    .container-content-flex {
        flex: 1;
        margin-top: 10px;
    }

    .posts-container-footer {
        height: 60px;
        display: flex;
        flex-direction: column;
        align-content: center;
        justify-items: center;
        justify-content: center;
        background: #ffffff;
        margin-bottom: 40px;
    }
}

.pagination {
  width: 100%;
  background: #ffffff;
  border-bottom: 1px solid #e6e9eb;
  min-height: 68px;
}
</style>
