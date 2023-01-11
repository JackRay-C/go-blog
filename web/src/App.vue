<template>
  <div id="app">
    <component :is="layout">
      <router-view v-if="isRouterAlive" />
    </component>
  </div>
</template>

<script>
export default {
  components: {},
  provide() {
    return {
      reload: this.reload,
    };
  },
  data() {
    return {
      default_layout: "empty",
      isRouterAlive: true,
    };
  },
  computed: {
    layout() {
      return (this.$route.meta.layout || this.default_layout) + "-layout";
    },
  },
  methods: {
    reload() {
      this.isRouterAlive = false;
      this.$nextTick(function () {
        this.isRouterAlive = true;
      });
    },
  },
};
</script>

<style lang="scss">
html,
body {
  font-family: "pingfang SC", "helvetica neue", arial, "hiragino sans gb",
    "microsoft yahei ui", "microsoft yahei", simsun, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  // color: #3e4c5b;
  font-size: 62.5%;
  text-rendering: optimizeLegibility;
  font-variant: tabular-nums;
  line-height: 1.65;
  letter-spacing: 0.2px;
  font-weight: 400;
  background: #fff;
  font-feature-settings: "tnum";
  padding: 0;
  margin: 0;
}
* {
  margin: 0;
  padding: 0;
  font-weight: 500;
}
a {
  text-decoration: none;
  &:hover,
  &:active,
  &:focus,
  &:visited {
    text-decoration: none;
  }
}

div {
  display: block;
}

::selection {
  color: inherit;
  background: #e1eaff;
}

::-webkit-scrollbar {
  width: 0px; /*对垂直流动条有效*/
}

.fadeInUp {
  animation: fadeInUp 0.5s cubic-bezier(0.75, 0.82, 0.165, 1);
}

@keyframes fadeInUp {
  0% {
    opacity: 0;
    transform: translate(0, 10%, 0);
    -webkit-transform: translate3d(0, 10%, 0);
    -ms-transform: translate3d(0, 10%, 0);
  }

  100% {
    opacity: 1;
    -ms-transform: none;
    -webkit-transform: none;
    transform: none;
  }
}
</style>

<style lang="scss">
.markdown-body {
  font-family: "pingfang SC", "helvetica neue", arial, "hiragino sans gb",
    "microsoft yahei ui", "microsoft yahei", simsun, sans-serif;
  font-size: 15px;
  color: #555555;

  * {
    box-sizing: border-box;
    list-style: inherit;
  }

  h1:first-child,
  h2:first-child,
  h3:first-child,
  h4:first-child,
  h5:first-child {
    margin-top: 30px;
  }
  img:first-child {
    margin-top: 30px;
  }

  h1,
  h2,
  h3,
  h4,
  h5 {
    color: #0a0a0a;
    font-weight: 700;
    margin-top: 36px;
    margin-bottom: 16px;
  }

  h1 {
    margin-top: 20px;
    margin-bottom: 16px;
    font-size: 25px;
    line-height: 2;
    font-weight: 600;
  }
  h2 {
    line-height: 2;
    font-size: 23px;
  }
  h3 {
    margin-top: 15px;
    margin-bottom: 15px;
    font-size: 20px;
  }
  h4,
  h5 {
    font-size: 15px;
    margin-top: 15px;
    margin-bottom: 15px;
  }

  p {
    font-size: 15px;
    line-height: 2.5;
    word-wrap: break-all;
    text-size-adjust: none;
    margin-top: 0px;
    margin-bottom: 20px;
    white-space: pre-wrap;
    color: #555555;

    strong {
      color: #464952;
      font-weight: 500;
    }
  }

  img {
    width: 100%;
    height: 100%;
    padding: 0;
    margin: 0;
  }

  pre {
    padding: 0;
    margin: 24px auto;
    white-space: pre-wrap;
    word-break: break-all;
  }

  code {
    font-size: 16px;
    margin: 0;
    white-space: pre-wrap;
    word-break: break-all;
  }

  ol,
  ul {
    font-size: 16px;
    font-weight: 400;
    line-height: 3;
    margin: 20px 0;

    li {
      p {
        margin: 0;
        padding: 0;
      }
      line-height: 2.5;
      font-size: 15px;
      color: #555555;
    }
  }

  table {
    border-collapse: collapse;
    empty-cells: show;
    margin-bottom: 20px;
    overflow: auto;
    border-spacing: 0;
    // display: block;
    word-break: keep-all;
    width: 100%;

    tr {
      background-color: #e5f3f3;
      border: 1px solid #2f2f31;
      color: #555555;
      text-align: center;
    }

    td,
    th {
      padding: 10px 13px;
      border: 1px solid #e1e1eb;
      word-break: normal;
      white-space: nowrap;

      &:first-child::after {
        content: "";
        display: inline-block;
        vertical-align: top;
        min-height: 24px;
      }
    }

    th {
      font-weight: 600;
      color: #303030;
    }

    tbody tr:nth-child(2n-1) {
      background-color: #fff;
    }
  }
}
</style>
