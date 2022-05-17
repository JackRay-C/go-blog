<template>
    <div class="about fadeInUp markdown-content">
      <div class="post-title">
        <span class="post-title-text">
          <h1>关于我</h1>
        </span>
      </div>

      <div class="post-content vditor-reset markdown-body" v-highlight>
        <h2>个人信息</h2>
        <p>1、姓名： 任浩杰</p>
        <p>2、性别： 男</p>
        <p>3、邮箱： 18435175817@163.com</p>

        <h2>技能</h2>
        <p>1、熟练掌握Java、Python的使用</p>
        <p>2、熟练掌握Mysql、Redis数据库的使用及原理</p>
        <p>3、熟练使用Go、Monogdb、Hadoop、Zookeeper、RabiitMQ、Kafka</p>
        <p>4、了解过Spark、Hive、Pytorch的使用</p>

        <h2>职业经历</h2>

        <p v-html="'1、北京北纬通信科技股份有限公司 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2017.07-202007\
          职位： 信息安全工程师\
          工作职责：\
          1、公司内部应用渗透测试\
          2、公司及部门内部应用开发'">
          
        </p>
      </div>
    </div>

</template>

<script>
export default {
  name: "Detail",
  data() {
    return {};
  }
};
</script>

<style lang="scss" scoped>
.about {
  font-weight: 400;
  font-size: 16px;
  width: 60%;
  min-height: 573px;
  margin: 70px auto 110px;
  padding-left: 32px;
  box-sizing: border-box;
}

.post-title {
  align-items: center;
  margin-bottom: 13px;
  position: relative;
  font-size: 30px;
  font-weight: 800;
  // color: #3e4c5b;
}
.post-title-text {
  vertical-align: middle;
  font-size: 21px;
}

.post-info .post-tag {
  cursor: pointer;
}
.devider {
  display: inline-block;
  width: 1px;
  height: 12px;
  background-color: #c3c6cb;
  margin: 0 8px;
}
// .post-content {
//   line-height: 2em;
//   white-space: pre-line;
//   word-break: break-word;
//   font-size: 16px;
//   color: #3e4c5b;
//   box-sizing: border-box;
//   font-weight: 400;

//   h1,
//   h2,
//   h3,
//   h4,
//   h5,
//   h6 {
//     color: #616569;
//     line-height: 2.5;

//     &:first-child {
//       margin-top: 50px;
//     }
//   }
//   h1 {
//     font-weight: 800;
//   }
//   h2 {
//     font-weight: 800;
//     font-size: 25px;
//   }
//   p {
//     line-height: 2em;
//     white-space: pre-line;
//     word-break: break-word;
//     word-wrap: break-word;
//     font-size: 16px;
//     color: #616569;
//     box-sizing: border-box;
//     font-weight: 400;
//     margin: 8px 0;
//   }

//   pre[class*="language-"] {
//     line-height: 1.5;
//     padding: 0;
//     margin: 1rem 0;
//     font-size: 16px;
//     overflow: auto;
//     user-select: contain;
//     -webkit-user-select: text;
//     -moz-user-select: text;
//     tab-size: 4;
//     hyphens: none;
//     display: block;
//     position: relative;
//     z-index: 1;
//     code {
//       padding: 2rem 1rem;
//     }
//   }
// }
</style>