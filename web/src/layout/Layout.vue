<template>
  <div id="home">
    <Header />
    <router-view />
    <Footer />
  </div>
</template>

<script>
import Header from "@/components/Header";
import Footer from "@/components/Footer.vue";

export default {
  components: {
    Header,
    Footer
  }
};
</script>

<style scoped>
#home {
  /* background: #f6f7fa; */
  background: #F6F7FD;
}
</style>