<template>
  <div class="article">
    <div class="article-card">
      <div class="article-header">
        <div class="article-header-left">
          <el-button icon="el-icon-plus" @click="newPost">新建</el-button>
          <el-button icon="el-icon-delete">删除</el-button>
        </div>
        <div class="article-header-right">
          <el-input
            placeholder="请输入内容"
            v-model="input3"
            class="input-with-select"
          >
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </div>
      </div>

      <div class="article-list">
        <el-table
          ref="multipleTable"
          :data="posts"
          tooltip-effect="dark"
          style="width: 100%"
          :cell-class-name="cellStyle"
          @selection-change="handleSelectionChange"
          :header-cell-class-name="headerCellStyle"
        >
          <el-table-column type="selection" width="55"></el-table-column>
          <el-table-column label="标题" show-overflow-tooltip>
            <template slot-scope="scope">{{ scope.row.title }}</template>
          </el-table-column>
          <el-table-column
            label="创建时间"
            width="250"
            resizable
            show-overflow-tooltip
          >
            <template slot-scope="scope">{{
              scope.row.created_at | datefmt("YYYY-MM-DD HH:mm:ss")
            }}</template>
          </el-table-column>
          <el-table-column
            label="更新时间"
            width="250"
            resizable
            show-overflow-tooltip
          >
            <template slot-scope="scope">{{
              scope.row.updated_at | datefmt("YYYY-MM-DD HH:mm:ss")
            }}</template>
          </el-table-column>
          <el-table-column
            prop="public"
            effect="plain"
            label="发布"
            width="200"
            resizable
          >
            <template slot-scope="scope">
              <el-tag size="medium" effect="plain">
                {{ scope.row.public | dictfmt }}</el-tag
              >
            </template>
          </el-table-column>
          <el-table-column label="操作" resizable align="center" width="200">
            <template slot-scope="scope">
              <el-button
                type="text"
                size="mini"
                @click="handleEdit(scope.$index, scope.row)"
                style="margin-right: 12px"
                >编辑</el-button
              >
              <el-popover placement="top" width="130" trigger="hover">
                <div>
                  <p style="margin-bottom: 12px">确定删除吗？</p>
                  <el-button
                    type="primary"
                    size="mini"
                    @click="handleDelete(scope.$index, scope.row)"
                    >确定</el-button
                  >
                </div>
                <el-button type="text" size="mini" slot="reference"
                  >删除</el-button
                >
              </el-popover>
            </template>
          </el-table-column>
        </el-table>
        <el-pagination
          class="pagination"
          @current-change="handleCurrentChange"
          :current-page="page_no"
          :page-sizes="[10, 20, 30, 50]"
          :page-size="page_size"
          background
          layout="total, prev, pager, next, jumper"
          :total="total_rows"
        >
        </el-pagination>
      </div>
    </div>
  </div>
</template>

<script>
import { listPosts, deletePost } from "@/api/post.js";
import { listDicts } from "@/api/dict";

var _this;
export default {
  name: "Article",
  components: {},
  data() {
    return {
      showDropdown: false,
      page_size: 10,
      total_page: 0,
      page_no: 1,
      total_rows: 0,
      posts: [],
      cellStyle: "cellStyle",
      headerCellStyle: "headerCellStyle",
      dicts: [],
      visible: false,
    };
  },
  filters: {
    dictfmt: function(code) {
      if (_this.dicts.find((item) => item.code == code) != undefined) {
        return _this.dicts.find((item) => item.code == code).value;
      } else {
        return "undefined";
      }
    },
  },
  beforeCreate() {
    _this = this;
  },
  mounted() {
    this.fetchPosts();
    this.fetchDicts();
  },
  methods: {
    fetchPosts() {
      listPosts({ page_no: this.page_no, page_size: this.page_size }).then(
        (res) => {
          this.posts = res.data.list;
          this.total_rows = res.data.total_rows;
          this.page_no = res.data.page_no;
          this.page_size = res.data.page_size;
          this.total_page = res.data.total_page;
        }
      );
    },
    fetchDicts() {
      listDicts({ page_no: 0, page_size: 10, name: "public" })
        .then((res) => {
          if (res.code === 200) {
            res.data.list.forEach((e) => {
              this.dicts.push(e);
            });
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
    newPost() {
      this.$router.push("/admin/edit");
    },
    newDraft() {
      this.$router.push("/admin/edit");
    },
    handleEdit(index, row) {
      this.$router.push("/admin/edit/" + row.id);
    },
    handleDelete(index, row) {
      deletePost(row.id)
        .then((res) => {
          console.log(res);
          this.fetchPosts();
        })
        .catch((err) => {
          console.log(err);
        });
    },
    handleCurrentChange(current) {
      this.page_no = current;
      this.fetchPosts();
    },
  },
};
</script>

<style lang="scss" scoped>
.article {
  height: 100%;
  width: 100%;
  box-sizing: border-box;
  padding: 30px 30px 30px 30px;
  display: flex;
  flex-direction: column;

  .article-card {
    box-sizing: border-box;
    width: 100%;
    height: 100%;
    -webkit-box-orient: vertical;
    box-shadow: 0 8px 24px 0 rgba(54, 97, 174, 0.05);
    background: #fff;

    .article-header {
      height: 100px;
      display: flex;
      justify-content: center;
      align-items: center;
      padding: 0 20px;
      // border-bottom: 1px #ecebeb solid;
      .article-header-left {
        width: 40%;
        box-sizing: border-box;
      }
      .article-header-right {
        width: 60%;
        box-sizing: border-box;
      }
    }

    .article-list {
      padding: 0 20px;
      margin-top: 20px;
      -webkit-box-orient: vertical;

      .pagination {
        padding: 18px 25px 0 25px;
        text-align: right;
      }
    }
  }
}

.header {
  width: 100%;
  height: 46px;
  margin: 30px 0 30px 0;
  display: flex;
  flex-direction: row;

  .button-left-group .button {
    text-align: center;
    min-width: 180px;
    height: 46px;
    outline: none;
    border: 1px solid transparent;
    box-sizing: border-box;
    line-height: 46px;
    text-align: left;
    cursor: pointer;
    font-size: 1.1em;
    // letter-spacing: 2px;
    background: #4e6ef2;
    color: #ffffff;
    border-radius: 4px;
    position: relative;
    text-align: center;

    .icon {
      font-size: 15px;
      vertical-align: -1%;
      margin: 0px auto;
    }

    &:hover {
      border: 1px solid #4e6ef2;
    }
  }
}

.content {
  width: 100%;
  height: 80%;
  background: #ffffff;
  border-radius: 4px;
  box-shadow: 0 6px 28px 28px rgba(24, 52, 117, 0.03);
  position: relative;

  .item {
    width: 100%;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: flex-start;
    align-items: center;
    transform-style: preserve-3d;
  }

  .pagination {
    padding: 18px 25px 0 25px;
    text-align: right;
  }
}
</style>

<style lang="scss">
.cellStyle {
  padding-left: 28px !important;
}
.headerCellStyle {
  padding-left: 28px !important;
  color: #3e4c5b;
  /* font-size: 15px; */
}

.post-dropdown {
  list-style: none;
  margin-top: 3px;
  padding: 0;
  border-radius: 4px;
  padding: 10px 0;
  background: #ffffff;
  box-shadow: 0 4px 12px 0 rgba(0, 0, 0, 0.15);
  text-align: left;
  user-select: none;
  z-index: 2000;

  li {
    list-style: none;
    box-sizing: border-box;
    line-height: 30px;
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: pre;

    .menu-item {
      padding: 10px;
      line-height: 24px;
      color: #555555;
      transition: 0.2s;
      font-size: 16px;
      box-sizing: border-box;
      width: 100%;
      text-decoration: none;
      display: flex;
      flex-direction: row;
      justify-content: space-around;
      align-items: center;

      p {
        display: inline;
        font-size: 16px;
        line-height: 22px;
        font-weight: 500;
        box-sizing: border-box;
        text-align: left;
      }

      &:hover {
        background: #f4f5fa;
        color: #4e6ef2;
        cursor: pointer;
      }
    }
  }
}

.button-popper-class-1 {
  padding: 0 !important;
}
</style>
