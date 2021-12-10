<template>
  <div class="subject">
    <div class="container">
      <el-form
        :label-position="labelPosition"
        label-width="80px"
        :model="formLabelAlign"
      >
        <el-form-item label="名称">
          <el-input v-model="formLabelAlign.name"></el-input>
        </el-form-item>
        <el-form-item label="活动区域">
          <el-input v-model="formLabelAlign.region"></el-input>
        </el-form-item>
        <el-form-item label="活动形式">
          <el-input v-model="formLabelAlign.type"></el-input>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: "NewSubject",
  data() {
    return {
      subject: {
        title: "",
      },
      labelPosition: "right",
      formLabelAlign: {
        name: "",
        region: "",
        type: "",
      },
    };
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

  .container {
    width: 70%;
    margin: 0 auto;
    display: flex;
    flex-direction: column;

    .input-item {
      width: 100%;
      display: flex;
      flex-direction: row;

      .input-label {
        width: 120px;
      }
    }
  }
}
</style>
