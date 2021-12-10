<template>
  <div class="containter fadeInUp">
    <div class="content">
      <SubjectCard v-for="subject in subjects" :key="subject.id" :subject="subject" />

      <Pagination :pageCount="total_page" :pagerCount="10" @change="handlePageChange" />

    </div>
    
  </div>
</template>


<script>
import SubjectCard from "@/components/SubjectCard.vue";
import Pagination from "@/components/Pagination.vue";
import {getSubjects} from "@/api/subject.js";

export default {
  name: "Subject",
  components: {
    SubjectCard,
    Pagination
  },
  data() {
    return{
      page_size: 10,
      total_page: 0,
      page_no: 1,
      total_rows: 0,
      subjects: []
    }
  },
  beforeRouteUpdate(to, from, next) {
    this.page_no = to.query.page || 1
    this.fetchSubject()
    next()
  },
  mounted() {
    this.page_no = this.$route.query.page  || 1
    this.fetchSubject()
    window.addEventListener("scroll", this.scrollTop);
  },
  methods: {
    fetchSubject(){
      getSubjects({page_no:this.page_no, page_size: this.page_size}).then(res => {
        console.log(res)
        if(res.code === 200 && res.data) {
          this.subjects = res.data.list;
          this.page_size = res.data.page_size
          this.page_no = res.data.page_no
          this.total_page = res.data.total_page
          this.total_rows = res.data.total_rows
        }
      })
      .catch(err => {
        console.log(err)
      })
    },
    handlePageChange(current) {
      this.page_no = current;
      this.$router.push({path: "/subject", query: {page: this.page_no}})
    },
  }
};
</script>

<style lang="scss" scoped>
.containter {
  width:100%;
  padding-top: 70px;

  .content {
    width: 1200px;
    margin: 0 auto;
    padding-bottom: 64px;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    box-pack: justify;
    -webkit-box-pack: justify;
  }

}
</style>