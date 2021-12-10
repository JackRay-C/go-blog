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
      reload: this.reload
    }
  },
  data() {
    return {
      default_layout: 'empty',
      isRouterAlive: true
    }
  },
  computed: {
    layout() {
      return (this.$route.meta.layout || this.default_layout  ) + "-layout"
    }
  },
  methods: {
    reload() {
      this.isRouterAlive = false;
      this.$nextTick(function () {
        this.isRouterAlive = true
      })
    }
  }
};
</script>

<style lang="scss">
html,
body {
  font-family: -apple-system,BlinkMacSystemFont,Segoe UI,Roboto,Oxygen,Ubuntu,Droid Sans,Helvetica Neue,sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  // color: #3e4c5b;
  font-size: 62.5%;
  text-rendering: optimizeLegibility;
  font-variant: tabular-nums;
  line-height: 1.65;
  letter-spacing: .2px;
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
