<template>
  <div class="subject">
    <div class="header">
      <div class="button-left-group">
        <div class="button" @click="newSubject">
          <i>
            <svg
              t="1590853343243"
              class="icon"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              p-id="2599"
              width="0.9em"
              height="0.9em"
            >
              <path
                d="M956.1240653818753 579.0559211537523h-377.4918801217032v377.55208924331106c0 36.792317387069836-29.839867873102275 66.60264870994936-66.63332128133455 66.60264870994936-36.792317387069836 0-66.63218526017215-29.81033132287954-66.63218526017215-66.60264870994936v-377.55208924331106h-377.49415216402804c-36.822989958455004 0-66.63218526017215-29.81033132287954-66.63218526017215-66.60264870994936 0-36.77868513312089 29.812603365204367-66.60264870994936 66.63218526017215-66.60264870994936h377.4918801217032v-377.55208924331106c0-36.77868513312089 29.839867873102275-66.60264870994936 66.63218526017215-66.60264870994936 36.793453408232246 0 66.63332128133455 29.825099597990917 66.63332128133455 66.60264870994936v377.55322526447344h377.4918801217032c36.822989958455004 0 66.63332128133455 29.822827555666088 66.63332128133455 66.60037666762456-0.0011360211624130007 36.792317387069836-29.81033132287954 66.60037666762456-66.63332128133455 66.60037666762456z"
                fill="#ffffff"
                p-id="2600"
              />
            </svg>
          </i>
          新建专题
        </div>
      </div>
      <div class="button-right-group"></div>
    </div>

    <div class="content">
      <div class="count">
        <p>共 {{ subjects.length }} 个专题</p>
      </div>
      <div class="item">
        <SubjectCard
          v-for="subject in subjects"
          :subject="subject"
          :key="subject.id"
        />
      </div>
    </div>
  </div>
</template>

<script>
import SubjectCard from "@/components/admin/SubjectCard.vue";

import { getSubjects } from "@/api/subject.js";

export default {
  name: "Article",
  components: {
    SubjectCard,
  },
  data() {
    return {
      showDropdown: false,
      page_size: 10,
      total_page: 0,
      page_no: 1,
      total_rows: 0,
      subjects: [],
    };
  },
  mounted() {
    this.fetchSubjects();
   
  },
  methods: {
    fetchSubjects() {
      getSubjects().then((res) => {
        this.subjects = res.data.list;
        this.total_rows = res.data.total_rows;
        this.page_no = res.data.page_no;
        this.page_size = res.data.page_size;
        this.total_page = res.data.total_page;
      });
    },
   
    newSubject() {
      this.$router.push("/admin/subject/new");
    },
  },
};
</script>

<style lang="scss" scoped>
.subject {
  height: 100%;
  width: 100%;
  box-sizing: border-box;
  padding: 64px 100px 0 100px;
  display: flex;
  flex-direction: column;
  overflow: auto;
  // background: #f9f9f9;
}

.header {
  width: 100%;
  height: 100px;
  display: flex;
  flex-direction: row;
  justify-content: flex-start;

  .button-left-group .button {
    min-width: 150px;
    height: 40px;
    outline: none;
    border: 1px solid transparent;
    box-sizing: border-box;
    line-height: 38px;
    text-align: left;
    cursor: pointer;
    font-size: 1.1em;
    letter-spacing: 2px;
    background: #4e6ef2;
    color: #ffffff;
    border-radius: 4px;
    position: relative;
    margin-bottom: 20px;
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

  .count {
    width: 100%;
    display: flex;
    -webkit-box-pack: justify;
    justify-content: space-between;
    margin-bottom: 20px;

    p {
      color: #969798;
      font-size: 15px;
    }
  }

  .item {
    width: 100%;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: flex-start;
    align-items: center;
    transform-style: preserve-3d;
  }

}
</style>