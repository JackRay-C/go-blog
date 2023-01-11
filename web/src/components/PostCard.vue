<template>

    <div class="post-card">
      <div>
        <h1 class="post-title" v-on:click="toDetail">{{post.title}}</h1>
      </div>
      <div class="post-content">
        <p v-if="post.description">
          {{replaceContent(post.description)}}
        </p>
        <p v-else>
          {{replaceContent(post.markdown_content)}}
        </p>
      </div>
    </div>

</template>

<script>

export default {
  name: "PostCard",
  props: {
    post: {
      type: Object,
      required: true
    }
  },
  mounted(){
    
  },
  computed: {
  
  },
  methods: {
    toDetail() {
      this.$router.push(`/detail/${this.post.id}`);
    },
    replaceContent(content) {
      content = content.replace(/[\r\n]/g, " ");
      content = content.substring(0, 150);
      console.log(content)

      return content
    }
  }
};
</script>

<style lang="scss" scoped>
.post-card {
  width: 100%;
  height: 350px;
  margin: auto;
  background: #ffffff;
  &:not(:first-child) {
    margin: 18px auto;
  }

  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  justify-content: space-around;

  border-radius: 8px;
  -webkit-box-orient: vertical;
  box-shadow: 0 8px 24px 0 rgba(54, 97, 174, 0.05);
  padding: 24px 36px;
  transition: all 0.5s;
  
  &:hover {
    box-shadow: 0 6px 28px 0 rgba(24, 52, 117, 0.2);
    transform: scale(1.04,1.04);
    transition: all .4s cubic-bezier(0.075, 0.82, 0.165, 1);
  }
}
.post-title {
  height: 77px;
  cursor: pointer;
  font-size: 24px;
  font-weight: 500;
  line-height: 77px;
  text-align: justify;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
  color: #121127;
  // color: #2c23d6;
  display: inline-block;
}
.post-content {
  font-size: 17px;
  flex: 1;
  
  p {
    font-size: 15px;
    font-weight: 500;
    text-align: justify;
    line-height: 2.5;
    color: #121115;
    overflow: hidden;
    text-overflow: ellipsis;
    margin: 0px 0 28px 0;
  }
}
</style>