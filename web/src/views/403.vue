<template>
  <div class="container">
    <div class="a">
       <p>403 - Forbiden .</p>
      <p>Access to this resource on the server is denied. </p>
      <div class="goback" @click="goback">go back </div> 
    </div>
     
  </div>
</template>

<script>
export default {
  data() {
    return  {
      url: ''
    }
  },
  created(){
    this.getUrl()
  },
  watch: {
    '$route': 'getUrl'
  },
  methods: {
    getUrl() {
      console.log(this.$route)
      this.url = this.$route.path
    },
    goback() {
      this.$router.go(-1)
    }
  }
}
</script>

<style lang="scss" scoped>
.container {
  width: 100%;
  height: 100%;
  background: #fff;
  padding: 15px;
  color: #222;

  .a {
    margin: 7% auto 0;
    max-width: 390px;
    min-height: 180px;
    padding: 30px 0 15px;


    p {
      margin: 11px 0 22px;
      overflow: hidden;
      font-size: 18px;
    }


    .goback {
      display: block;
      color: #30cf43;
      font-size: 16px;
      cursor: pointer;
  
    }
  }

}
</style>
