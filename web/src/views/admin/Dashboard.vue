<template>
  <div class="container">
    <div class="container-header">
      <header class="header-content">
        <h2 class="header-title">
          Dashboard
        </h2>
      </header>
    </div>

    <div class="container-content">
      <section class="container-content-data">
        <el-row :gutter="20">
          <el-col :span="6">
            <div>
              <el-statistic group-separator="," :value="postCount" title="Posts"></el-statistic>
            </div>
          </el-col>
          <el-col :span="6">
            <div>
              <el-statistic group-separator="," :value="subjectCount" title="Subjects"></el-statistic>
            </div>
          </el-col>
          <el-col :span="6">
            <div>
              <el-statistic group-separator="," :value="memberCount" title="Accounts"></el-statistic>
            </div>
          </el-col> 
          <el-col :span="6">
            <div>
              <el-statistic group-separator="," :value="memberCount" title="Tags"></el-statistic>
            </div>
          </el-col> 
        </el-row>
<!-- 
        <div class="dasbboard-box posts">
          <div class="chart-conntainer">
            <div class="chart-summary">
              <h4 class="chart-header">Total Posts</h4>
              <div class="title">
                {{postCount}}
              </div>
            </div>
            <div class="chart-content">
              <div class="chart-box">
                <div id="post-charts" style="height:100%;width:100%;"></div>
              </div>
            </div>
          </div>
        </div>
        <div class="dashboard-box subjects">
          <div class="chart-conntainer">
            <div class="chart-summary small">
              <h4 class="chart-header">Total Subjects</h4>
              <div class="chart-data-container">
                <div class="title">
                  {{subjectCount}}
                </div>
              </div>
            </div>
            <div class="chart-content small">
              <div class="chart-box small">
                <div id="views-charts" style="height:100%;width:100%;"></div>
              </div>
            </div>
          </div>
        </div>

        <div class="dashboard-box accounts">
          <div class="chart-conntainer">
            <div class="chart-summary small">
              <h4 class="chart-header">Total Members</h4>
              <div class="chart-data-container">
                <div class="title">
                  {{memberCount}}
                </div>
              </div>
            </div>
            <div class="chart-content small">
              <div class="chart-box small">
                <div id="month-charts" style="height:100%;width:100%;"></div>
              </div>
            </div>
          </div>
        </div> -->
      </section>
      <section class="dashboard-area mixed">
        <div class="mixed-container">
          <div class="mixed-box dasbboard-box">
            <div class="content">
              <h2>Create New Posts</h2>
              <p>
                Started create your posts for you idea!
              </p>
            </div>
            <div class="footer">
              <router-link to="/admin/edit" active-class="active">
                <span>New</span>
              </router-link>
              <router-link to="/admin/posts" active-class="active">
                <span>Mannager</span>
              </router-link>
            </div>
          </div>
          <div class="mixed-box dasbboard-box">
            <div class="content">
              <h2>Create One Subjects</h2>
              <p>
                Create one subjects for some posts.
              </p>
            </div>
            <div class="footer">
              <router-link to="/admin/subject" active-class="active">
                <span>New</span>
              </router-link>
              <router-link to="/admin/subject" active-class="active">
                <span>Mannager</span>
              </router-link>
            </div>
          </div>
          <div class="mixed-box dasbboard-box">
            <div class="content">
              <h2>Create Accounts</h2>
              <p>
                Create one accounts
              </p>
            </div>
            <div class="footer">
              <router-link to="/admin/subject" active-class="active">
                <span>New</span>
              </router-link>
              <router-link to="/admin/subject" active-class="active">
                <span>Mannager</span>
              </router-link>
            </div>
          </div>
        </div>
      </section>
      <section class="dashboard-area activity">
        <div class="dasbboard-box recent-posts">
          <div class="header">
            <h2>Recente Posts</h2>
          </div>
          <div class="content">
            <router-link v-for="p in recent_posts" :key="p.id" :to="'/admin/edit/'+p.id">
              <h2>{{p.title}}</h2>
              <span class="date">{{p.created_at | datefmt("YYYY-MM-DD") }}</span>
            </router-link>
          </div>
          <div class="footer">
            <router-link to="/admin/article">
               See More →
            </router-link>
          </div>
        </div>
        <div class="dashboard-box draft"></div>
      </section>
    </div>
    <div class="container-footer">
    </div>
  </div>
</template>

<script>
import * as echarts from "echarts";
import {listPosts} from "@/api/admin/post"
import {listSubjects} from "@/api/admin/subject"
import {listUsers} from "@/api/admin/user"

export default {
  data() {
    return {
      chart1: "",
      chart2: "",
      chart3: "",
      chartColor: [
        "#cbdccd",
        "#2f4554",
        "#61a0a8",
        "#d48265",
        "#91c7ae",
        "#749f83",
        "#ca8622",
        "#bda29a",
        "#6e7074",
        "#546570",
        "#c4ccd3",
      ],
      postCount: 0,
      memberCount: 0,
      subjectCount: 0,
      recent_posts: ''
    };
  },
  mounted() {
    // this.initPostCharts();
    // this.initViewsCharts();
    // this.initMonthCharts();
    this.getPostCount();
    this.getSubjectCount();
    this.getMemberCount()
  },
  methods: {
    initPostCharts() {
      this.chart1 = echarts.init(document.getElementById("post-charts"));
      this.chart1.setOption({
        color: this.chartColor,
        xAxis: {
          data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
        },
        yAxis: {
          axisLine: {
            show: true,
          },
          axisLabel: {
            show: true,
          },
        },
        tooltip: {
          trigger: "item",
          formatter: "{b} <br/>{c} ",
        },
        series: [
          {
            type: "bar",
            data: [23, 24, 18, 25, 27, 28, 155],
          },
        ],
      });
    },
    initViewsCharts() {
      this.chart2 = echarts.init(document.getElementById("views-charts"));
      this.chart2.setOption({
        xAxis: {
          type: "category",
          data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
        },
        yAxis: {
          type: "value",

          axisLine: {
            show: true,
          },
          axisLabel: {
            show: false,
          },
        },
        color: this.chartColor,
        tooltip: {
          trigger: "item",
          formatter: "{b} <br/>{c} ",
        },
        series: [
          {
            data: [220, 182, 191, 234, 290, 330, 310],
            type: "line",
            areaStyle: {},
            smooth: true,
          },
        ],
      });
    },
    initMonthCharts() {
      this.chart3 = echarts.init(document.getElementById("month-charts"));
      this.chart3.setOption({
        xAxis: {
          data: ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"],
        },
        yAxis: {
          show: true,
          axisLine: {
            show: true,
          },
          axisLabel: {
            show: false,
          },
        },
        tooltip: {
          trigger: "item",
          formatter: "{b} <br/>{c} ",
        },
        color: this.chartColor,
        series: [
          {
            type: "scatter",
            data: [220, 182, 191, 234, 290, 330, 310],
          },
        ],
      });
    },
    getPostCount(){
      listPosts({page_no: 1, page_size: 5}).then(res => {
         if(res.code === 200){
           this.recent_posts = res.data.list
           this.postCount = res.data.total_rows
        } else {
           console.log(res)
        }
       
      }).catch(err => {
        console.log(err)
      })
    },
    getSubjectCount(){
      listSubjects({page_no: 1, page_size: 10}).then(res => {
        if(res.code === 200){
           this.subjectCount = res.data.total_rows
        } else {
           console.log(res)
        }
      
      }).catch(err => {
        console.log(err)
      })
    },
    getMemberCount(){
      listUsers({page_no: 1, page_size: 10}).then(res => {
        if(res.code === 200){
           this.memberCount = res.data.total_rows
        } else {
           console.log(res)
        }
      }).catch(err => {
        alert(err)
      })
    }
  },
};
</script>

<style lang="scss" scoped>
.container-content {
  display: grid;
  grid-template-columns: 2fr 1fr;
  grid-gap: 1.8vw;
  position: relative;
  flex-grow: 1;
  padding-top: 32px;
  padding-bottom: 32px;
}

.dashboard-area {
  display: felx;
  flex-direction: column;
}
.dashboard-area.charts {
  display: grid;
  grid-template-columns: 2fr 1fr;
  grid-template-rows: 1fr 1fr;
  border: 1px solid #edeef0;
  border-radius: 3px;
  grid-column-gap: 1.8vw;
  grid-row-gap: 5px;
  align-items: stretch;
  padding: 16px 0 0;
  flex-direction: column;
  background: #ffffff;
  grid-column: 1/3;
}

.container-content-data {
  background: #fff;
  padding: 24px 0;
  border: 1px solid #edeef0;
  border-radius: 3px;
  grid-column: 1/3;
  display: grid;
  flex-direction: column;
  grid-column-gap: 1.8vw;
}
.dasbboard-box {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  border: 1px solid #ebeef0;
  border-radius: 3px;
  padding: 28px;
  margin-bottom: 1.0vw;
  // box-shadow: 0 8px 24px 0 rgba(54, 97, 174, 0.05);
}

.dashboard-area.charts .posts {
  grid-column: 1/2;
  grid-row: 1/4;
  border: none;
  border-radius: 0;
  margin: 0;
  padding: 0 0 0 28px;
}

.box-summary {
  flex: 1;
  display: flex;
}

.dashboard-area.charts .subjects {
  grid-column: 2/3;
  grid-row: 1/2;
  padding: 0 28px;
  height: 75px;
}
.dashboard-area.charts .accounts {
  grid-column: 2/3;
  grid-row: 2/3;
  padding: 0 28px;
  height: 75px;
}

.chart-conntainer {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;

  .chart-summary {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: space-between;
    margin-top: 10px;

    .chart-header {
      display: flex;
      align-items: center;
      font-size: 1.1rem;
      //   text-transform: uppercase;
      font-weight: 500;
      letter-spacing: 0.5px;
      padding: 0;
      color: #15171a;
      margin: 3px 0 4px;
    }
    .title {
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 3.0rem;
      line-height: 1;
      font-weight: 600;
      color: #15171a;
      letter-spacing: -0.1px;
      white-space: nowrap;
    }
  }

  .chart-content {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    justify-content: flex-start;

    .chart-box {
      position: relative;
      display: flex;
      justify-content: center;
      height: 320px;
      width: 40vw;
      margin-right: -50px;
      margin-top: -35px;
      margin-bottom: -20px;
    }
  }
}

.chart-summary.small {
  min-width: 12px;
  margin-top: 0;
}
.small {
  .chart-header {
    margin: 3px 0 4px;
  }
}
.small .title {
  font-size: 2.8rem;
  letter-spacing: -0.1px;
}
.chart-content.small {
  margin-bottom: -6px;
  flex-grow: 1;
}
.chart-box.small {
  position: relative;
  width: 14vw !important;
  min-width: 165px;
  height: 180px !important;
  padding-top: 0;
  margin-right: -6px !important;
  transform: scale(0.9);
  transform-origin: top right;
}

.mixed {
  grid-column: 1/2;
  border: none;
  background: transparent;
  border-radius: unset;
  padding: 0;
  align-items: stretch;
  justify-content: flex-start;

  .mixed-container {
    grid-template-columns: 1fr 1fr;
    grid-gap: 1.8vw;
    display: grid;
    width: 100%;

    .mixed-box {
      background: #fff;
    }
  }

  .mixed-container .content {
    color: #394047;

    h2 {
      font-size: 1.65rem;
      font-weight: 600;
      line-height: 1.4em;
      margin-bottom: 8px;
      color: #15171a;
    }
    p {
      margin-bottom: 12px;
    }
  }
  .mixed-container .footer {
    display: flex;
    align-items: center;
    flex-wrap: nowrap;

    a {
      color: #394047;
      fill: #394047;
      border: 1px solid #dee3e7;
      background: none;
      display: inline-block;
      outline: none;
      text-decoration: none;
      user-select: none;
      font-weight: 500;
      border-radius: 3px;
      transition: all 0.2s ease;
      transition-property: color, border-color, background, width, height,
        box-shadow;
      margin-top: 0.8rem;
      margin-right: 0.8rem;

      &:hover {
        color: #15171a;
        background: #ebeef0;
        border-color: #ced4d9;
        transition: background 0.1s, color 0.1s;
        outline: 0;
      }

      span {
        display: block;
        overflow: hidden;
        padding: 0 14px;
        height: 34px;
        font-size: 1.35rem;
        line-height: 34px;
        text-align: center;
        letter-spacing: 0.2px;
        border-radius: 3px;
        white-space: nowrap;
        text-overflow: ellipsis;
      }
    }
  }
}

.dashboard-area.activity {
  grid-column: 2/3;
  display: flex;
  flex-direction: column;

  .dasbboard-box {
    background: #ffffff;
    border: 1px solid #ebeef0;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    border-radius: 3px;
    padding: 28px;
    margin-bottom: 1.8vw;
  }
}

.dasbboard-box.recent-posts {
  .header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-bottom: 1px solid #ebeef0;
    padding-bottom: 10px;

    h2 {
    //   display: flex;
    //   align-items: center;
    //   font-size: 1.1rem;
    //   font-weight: 500;
    //   letter-spacing: 0.2px;
    //   margin: -4px 0 4px;
    //   padding: 0;
    //   color: #15171a;
    font-size: 1.65rem;
      font-weight: 600;
      line-height: 1.4em;
      margin-bottom: 8px;
      color: #15171a;
    }
  }
  .content {
    color: #394047;

    a {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
      color: #394047;
      margin: 12px 0;
      padding: 0 0 12px;
      border-bottom: 1px solid #ebeef0;
      text-decoration: none;

      &:hover {
        text-decoration: none;
        color: #30cf43;
        transition: background 0.1s, color 0.1s;
        outline: 0;
        h2 {
            color: #30cf43;
        }
      }

      h2  {
          margin-bottom: 0;
          font-size: 1.65rem;
          font-weight: 600;
          line-height: 1.4em;
          color: #15171a;
          margin: 0 0 0.3em;
          text-rendering: optimizeLegibility;
      }
      span {
          font-size: 1.3rem;
          color: #7c8b9a;
      }
    }
  }
  .footer {
      margin-bottom: -12px;
      display: flex;
      align-items: center;
      flex-wrap: nowrap;

      a {
          color: #30cf43;
          outline: none;
          text-decoration: none;

          &:hover {
              transition: background .1s, color .1s;
          }
      }
  }
}
</style>
