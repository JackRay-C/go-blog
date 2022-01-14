<template>
  <div class="tag fadeInUp">
    <div class="title">
      <h1>标签</h1>
    </div>
    <div class="wordCount">
      <div
        v-for="tag in tags"
        :key="tag.id"
        class="tagCard"
        @click="goTag(tag.id)"
      >
        {{ tag.name }}
      </div>

      <Pagination
      :pageCount="total_page"
      :pagerCount="10"
      @change="handlePageChange"
    />

    </div>
  </div>
</template>
<script>
import Pagination from "@/components/Pagination.vue";

import { getTags } from "@/api/web/tag.js";

export default {
  name: "Tag",
  components: {
    Pagination,
  },
  data() {
    return {
      page_size: 100,
      total_page: 0,
      page_no: 1,
      total_rows: 0,
      tags: [],
    };
  },
  beforeRouteUpdate(to, from, next) {
    this.page_no = to.query.page || 1;
    this.fetchTags();
    next();
  },
  mounted() {
    this.page_no = this.$route.query.page || 1;
    this.fetchTags();
    window.addEventListener("scroll", this.scrollTop);
  },
  methods: {
    fetchTags() {
      getTags({ page_no: this.page_no, page_size: this.page_size }).then(
        (res) => {
          if (res.code === 200 && res.data) {
            this.tags = res.data.list;
            this.page_size = res.data.page_size;
            this.page_no = res.data.page_no;
            this.total_page = res.data.total_page;
            this.total_rows = res.data.total_rows;
          }
        }
      );
    },
    goTag(id) {
      this.$router.push(`/tag/${id}`);
    },
    handlePageChange(current) {
      this.page_no = current;
      this.$router.push({ path: "/tag", query: { page: this.page_no } });
    },
  },
};
</script>

<style lang="scss" scoped>
.tag {
  width: 100%;
  font-weight: 400;
  font-size: 16px;
  min-height: 490px;
  margin: 50px auto 110px;
  padding-left: 32px;
  box-sizing: border-box;
}
.title {
  width: 60%;
  margin: 0 auto 20px;

  text-align: center;
}
.wordCount {
  width: 60%;
  margin: 0 auto;
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: flex-start;
  align-items: center;

  .tagCard {
    min-width: 130px;
    height: 50px;
    background: #ffffff;
    border-radius: 8px;
    margin: 15px 10px;
    padding: 10px 24px;
    line-height: 30px;
    text-align: center;
    box-sizing: border-box;
    box-shadow: 0 8px 24px 0 rgba(54, 97, 174, 0.097);
    transition: all 0.3s;
    cursor: pointer;

    &:hover {
      box-shadow: 0 6px 28px 0 rgba(24, 52, 117, 0.2);
      transform: scale(1.07, 1.07);
      transition: all 0.8s cubic-bezier(0.075, 0.82, 0.165, 0.9);
    }
  }
}
</style>

