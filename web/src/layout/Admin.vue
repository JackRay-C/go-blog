<template>
  <div id="admin">
    <!-- 侧边栏 -->
    <Aside class="aside" />

    <div class="main">
      <router-view />
    </div>
  </div>
</template>

<script>
import Aside from "@/components/admin/Aside.vue";
export default {
  components: {
    Aside
  }
};
</script>


<style lang="scss" scoped>
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
  flex: 0 0 320px;
  min-width: 0;
  transform: translate(0);
  border-right: 1px solid #e6e9eb;
  transition: flex-basis .25s;
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
</style>