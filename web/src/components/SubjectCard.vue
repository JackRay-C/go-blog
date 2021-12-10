<template>
  <div class="subject-card">
    <div class="subject-info">
      <div class="subject-title" @click="goSubject">{{subject.title}}</div>
      <div class="subject-subtitle">
        {{count}} 篇文章
      </div>
      <div class="subject-desc">
        {{subject.description}}
      </div>
    </div>
    <div class="subject-image" v-show="subject.image">
      <img :src="avatar" class="image" />
    </div>
  </div>
</template>

<script>
import {getSubjectPostCount} from "@/api/subject.js"
import {getFileById} from "@/api/file.js";
export default {
  name: "ArticleBox",
  props:[
    "subject"
  ],
  data(){
    return{
      count: '',
      avatar: ""
    }
  },
  mounted() {
    this.countSubjectPost()
    this.fetchAvatar()
  },
  methods: {
    goSubject() {
      this.$router.push(`/subject/${this.subject.id}`)
    },
    countSubjectPost() {
      getSubjectPostCount(this.subject.id).then(res => {
        if(res.code === 200) {
          this.count = res.data.total_rows;
        }
      })
    },
    fetchAvatar() {
      getFileById(this.subject.image).then(res => {
        if(res.code === 200){
          console.log(res)
          this.avatar = "http://localhost:8000/" +  res.data.access_url
        }
      }).catch(err => {
        console.log(err)
      })
    }
  }
};
</script>

<style lang="scss" scoped>
.subject-card {
  position: relative;
  width: 576px;
  height: 286px;
  padding: 38px 25px 0 26px;
  background: #ffffff;
  border: 10px;
  box-sizing: border-box;
  margin-bottom: 24px;
  border-radius: 8px;
  box-shadow: 0 8px 24px 0 rgba(54, 97, 174, 0.05);
  transition: all 0.5s;
  &:hover {
    box-shadow: 0 6px 28px 0 rgba(24, 52, 117, 0.2);
    transform: scale(1.04, 1.04);
    transition: all 0.8s cubic-bezier(0.075, 0.82, 0.165, 1);
  }

  .subject-info {
    width: 360px;
    margin-left: 12px;
    overflow: hidden;

    .subject-title{
      font-size: 25px;
      font-weight: 700;
      color: #4E6EF2;
      white-space: nowrap;
      text-overflow: ellipsis;
      overflow: hidden;
      cursor: pointer;
      transition: color .2s linear;

    }
    .subject-subtitle {
      font-size: 15px;
      font-weight: 500;
      padding: 10px 0;
      color: #9aa8b6;
    }
    .subject-desc{
      margin-top: 10px;
      line-height: 24px;
      display: -webkit-box;
      line-clamp: 3;
      -webkit-line-clamp: 3;
      -webkit-box-orient: vertical;
      overflow: hidden;
      font-size: 15px;
      font-weight: 500;
      color: #888888;

    }
  }
  .subject-image {
    position: absolute;
    top: 70px;
    right: 26px;
    width: 111px;
    min-width: 111px;
    height: 111px;
    background: #f6f7fb;
    border-radius: 50%;
    overflow: hidden;
    img {
      width: 100%;
      height: 100%;
    }
  }
}
</style>