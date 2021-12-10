<template>

    <div class="post-card">
      <div>
        <h1 class="post-title" v-on:click="toDetail">{{post.title}}</h1>
      </div>
      <div class="post-content">
        <p>
          {{replaceContent(post.description)}}
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
  methods: {
    toDetail() {
      this.$router.push(`/detail/${this.post.id}`);
    },
    replaceContent(content) {
      content = content.replace(/[\r\n]/g, " ");
      content = content.substring(0, 200);

      return content
    }
  }
};
</script>

<style lang="scss" scoped>
.post-card {
  width: 100%;
  height: 200px;
  margin: auto;
  background: #ffffff;
  &:not(:first-child) {
    margin: 24px auto;
  }
  overflow: hidden;

  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  justify-content: space-around;

  border-radius: 8px;
  -webkit-box-orient: vertical;
  box-shadow: 0 8px 24px 0 rgba(54, 97, 174, 0.05);
  padding: 24px 36px;
  transition: all 0.3s;
  &:hover {
    box-shadow: 0 6px 28px 0 rgba(24, 52, 117, 0.2);
    transform: scale(1.04,1.04);
    transition: all .3s cubic-bezier(0.075, 0.82, 0.165, 1);
  }
}
.post-title {
  cursor: pointer;
  font-size: 25px;
  font-weight: 700;
  line-height: 1.6;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
  color: #4E6EF2;
  display: inline-block;
}
.post-content {
  height: 77px;
  font-size: 16px;
  line-height: 1.57;
  letter-spacing: normal;
  color: #8c98a3;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  white-space: pre-line;
  -webkit-line-clamp: 2;
  margin-bottom: 24px;
}
</style>