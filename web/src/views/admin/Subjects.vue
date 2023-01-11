<template>
  <div class="container">
    <div class="container-header">
      <div class="header-content">
        <h2 class="header-title">
          <router-link to="/admin/subjects">Subjects</router-link>
        </h2>
        <section class="view-actions">
          <div class="contentfilter"></div>
        </section>
      </div>
    </div>

    <div class="container-content">
      <div class="mixed-container">
        <el-card class="mixed-box">
            subject1
        </el-card>
        <el-card class="mixed-box" >
            subject2
        </el-card>
        <el-card class="mixed-box" shadow="never">
            subject3
        </el-card>
      </div>
    </div>

    <div class="container-footer"></div>
  </div>
</template>

<script>
export default {
  name: "subjects",
  data() {
    return {
      total_page: 0,
      page_no: 1,
    };
  },
};
</script>

<style lang="scss" scoped>
.mixed-container {
  display: flex;
  flex-direction: row;
}

.mixed-box {
    background: #fff;
    padding: 28px;
    width: 33%;
    max-height: 72px;
    margin: 24px 12px;
}
</style>
