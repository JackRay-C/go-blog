<template>
  <div class="dashboard">
    <div class="row">
      <div class="col">
        <div id="post-charts"></div>
      </div>

      <div class="col">
        <div id="views-charts"></div>
      </div>

      <div class="col">
        <div id="month-charts"></div>
      </div>
      <div class="col"></div>
    </div>

    <div class="row">
      <div class="col"></div>
    </div>
  </div>
</template>

<script>
import * as echarts from "echarts";

export default {
  name: "dashboard",
  data() {
    return {
      chart1: "",
      chart2: "",
      chart3: "",
      chartColor: [
        "#409EFF",
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
    };
  },
  mounted() {
    this.initPostCharts();
    this.initViewsCharts();
    this.initMonthCharts();
  },
  methods: {
    initPostCharts() {
      this.chart1 = echarts.init(document.getElementById("post-charts"));
      this.chart1.setOption({
        title: {
          text: "访问量",
        },
        color: this.chartColor,
        xAxis: {
          data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
        },
        yAxis: {},
        series: [
          {
            type: "bar",
            data: [23, 24, 18, 25, 27, 28, 25],
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
        },
        color: this.chartColor,
        tooltip: {
          trigger: "item",
          formatter: "{a} <br/>{b} : {c} ({d}%)",
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
        yAxis: {},
        color: this.chartColor,
        series: [
          {
            type: "scatter",
            data: [220, 182, 191, 234, 290, 330, 310],
          },
        ],
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.dashboard {
  width: 100%;
  height: 100%;
  padding: 30px 30px 60px 30px;
  box-sizing: border-box;
  color: #8af311;

  .row {
    width: 100%;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-around;
    margin-bottom: 20px;

    .col {
      flex: 1;
      height: 260px;
      background: #ffffff;
      border-radius: 8px;
      -webkit-box-orient: vertical;
      box-shadow: 0 8px 24px 0 rgba(54, 97, 174, 0.05);

      &:not(first-child) {
        margin-left: 20px;
      }

      #post-charts {
        padding: 5px 5px;
        width: 100%;
        height: 100%;
      }

      #views-charts {
        padding: 5px 5px;
        width: 100%;
        height: 100%;
      }

      #month-charts {
        padding: 5px 5px;
        width: 100%;
        height: 100%;
      }
    }
  }
}
</style>
