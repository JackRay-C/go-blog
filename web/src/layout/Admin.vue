<template>
  <div id="admin">
    <!-- 侧边栏 -->
    <Aside class="aside" />

    <div class="main">
      <router-view :key="key" />
    </div>
  </div>
</template>

<script>
import Aside from "@/components/admin/Aside.vue";
export default {
  components: {
    Aside,
  },
  computed: {
    key() {
      return this.$route.name !== undefined
        ? this.$route.name + "?" + new Date().getTime().toString()
        : this.$route + "?" + new Date().getTime().toString();
    },
  },
};
</script>

<style lang="scss">
#admin {
  width: 100vw;
  height: 100vh;
  display: flex;
  box-sizing: border-box;
  flex-direction: row;
  position: fixed;
  top: 0;
  right: 0;
  left: 0;
  bottom: 0;
  height: 100%;
}
.aside {
  position: relative;
  z-index: 800;
  flex: 0 0 300px;
  min-width: 0;
  transform: translate(0);
  border-right: 1px solid #e6e9eb;
  transition: flex-basis 0.25s;
  overflow: hidden;
}
.main {
  position: relative;
  flex-grow: 1;
  display: flex;
  overflow-y: auto;
  overflow-x: hidden;

  box-sizing: border-box;
  background: #f6f7fa;
}
.main .container {
  width: 100%;
  position: relative;
  flex-grow: 1;
  padding: 0 48px 10px 36px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  background: #f6f7fa;
  height: auto;
  overflow: auto;
}

.main .container .container-header {
  margin: 0 -48px;
  padding: 0 48px;
  position: sticky;
  top: 0;
  background: #ffffff;
  z-index: 700;
  border-bottom: 1px solid #edeeef;
}

.main .container .container-header .header-content {
  height: 95px;
  position: relative;
  flex-shrink: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.main .container .container-header .header-content .header-title {
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
}
.main .container .container-header .header-content .header-title a {
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
.main .container .container-header .header-content view-actions {
}



.main .container .container-content {
  flex: 1;
}

.main .container .container-footer {
  height: 60px;
}
</style>
