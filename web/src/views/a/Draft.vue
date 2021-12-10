<template>
  <div class="draft">
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
        <p>共 {{postCount}} 篇文章</p>
      </div>

      <div class="subject-content">
        <SubjectCard v-for="subject in subjects" :subject="subject" :key="subject.subjectId" />
      </div>

    </div>

  </div>
</template>

<script>

import SubjectCard from "@/components/admin/SubjectCard.vue";

import { getSubjects } from "@/api/subject.js";
import { getPostsCount } from "@/api/post.js";

export default {
  name: "Draft",
  components: {

    SubjectCard
  },
  data() {
    return {
      postCount: 0,
      subjects: []
    };
  },
  mounted() {
    this.fetchSubjects();
    this.fetchPostsCount();
  },
  methods: {
    fetchSubjects() {
      getSubjects().then(res => {
        this.subjects = res.data;
      });
    },
    fetchPostsCount() {
      getPostsCount().then(res => {
        this.postCount = res.data;
      });
    },
    newSubject(){
      this.$router.push("/admin/newSubject")
    }
  }
};
</script>

<style lang="scss" scoped>
.draft {
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
    margin-bottom: 50px;
    text-align: center;

    .icon {
      font-size: 15px;
      vertical-align: -1%;
      margin: 0px auto;
    }

    &:hover {
      border: 1px solid #4e6ef2;
    }

    .dropdown {
      list-style: none;
      margin-top: 3px;
      padding: 0;
      position: absolute;
      top: 40px;
      left: 0px;
      min-width: 150px;
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
  }
}

.content {
  width: 100%;

  .count {
    width: 100%;
    display: flex;
    -webkit-box-pack: justify;
    justify-content: space-between;
    margin-bottom: 50px;

    p {
      color: #969798;
      font-size: 15px;
    }
  }

  .subject-content {
    width: 100%;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: flex-start;
    align-items: center;
    transform-style: preserve-3d;
  }

  .post-content {
    width: 100%;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: flex-start;
    align-items: center;
  }
}
</style>