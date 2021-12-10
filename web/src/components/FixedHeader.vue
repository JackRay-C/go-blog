<template>
  <header class="fixed-header">
    <div class="box">
      <div class="search-input-block">
        <span class="search-input-prefix">
          <i aria-label="图标: search" class="search-icon">
            <svg
              viewBox="64 64 896 896"
              class
              data-icon="search"
              width="1em"
              height="1em"
              fill="currentColor"
              aria-hidden="true"
              focusable="false"
            >
              <path
                d="M909.6 854.5L649.9 594.8C690.2 542.7 712 479 712 412c0-80.2-31.3-155.4-87.9-212.1-56.6-56.7-132-87.9-212.1-87.9s-155.5 31.3-212.1 87.9C143.2 256.5 112 331.8 112 412c0 80.1 31.3 155.5 87.9 212.1C256.5 680.8 331.8 712 412 712c67 0 130.6-21.8 182.7-62l259.7 259.6a8.2 8.2 0 0 0 11.6 0l43.6-43.5a8.2 8.2 0 0 0 0-11.6zM570.4 570.4C528 612.7 471.8 636 412 636s-116-23.3-158.4-65.6C211.3 528 188 471.8 188 412s23.3-116.1 65.6-158.4C296 211.3 352.2 188 412 188s116.1 23.2 158.4 65.6S636 352.2 636 412s-23.3 116.1-65.6 158.4z"
              />
            </svg>
          </i>
        </span>
        <input
          type="text"
          autocomplete="off"
          class="search-input"
          placeholder="Search..."
          :value="value"
        />
        <span class="search-input-suffix">
            <button type="button" class="search-button"><span>搜索</span></button>
        </span>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: "FixedHeader",
  data() {
    return {
      value: "",

    };
  },
  watch: {
    value(newVal) {
      this.$emit("change", newVal);
    },
  },
};
</script>

<style lang="scss" scoped>
.fixed-header {
  position: fixed;
  top: 0;
  text-align: center;
  margin: auto;
  z-index: 999;

  box-shadow: 0 2px 16px 0 rgba(31, 35, 41, 0.1);
  height: 72px;
  background: #ffffff;
  width: 100%;
  display: block;
}
.box {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  z-index: 1000;
  width: 100%;
}

.search-input-block {
  width: 60%;
  box-sizing: border-box;
  color: #3e4c5b;
  font-size: 15px;
  font-variant: tabular-nums;
  line-height: 1.5;
  list-style-type: none;
  text-align: start;
  display: inline-block;
  position: relative;
}

.search-input-prefix {
    left: 28px;
    font-size: 16px;
    position: absolute;
    top: 50%;
    z-index: 2;
    line-height: 0;
    color: #9ca2a9;
    transform: translateY(-50%);
}
.search-input {
    font-size: 16px;
    padding-left: 62px;
    padding-right: 200px;
    text-align: left;
    height: 50px;
    border: 1px solid #efeff1;
    border-radius: 4px;
    position: relative;
    resize: none;
    box-sizing: border-box;
    list-style: none;
    display: inline-block;
    width: 100%;
    color: #3e4c5b;
    background: #ffffff;
    transition: all .3s;
    outline: none;

    &:hover, &:focus {
        border: 1px solid #4E6EF2;
        background: #ffffff;
        border-right-width: 1px !important;
    }

}

.search-input-suffix {
    right: 0;
    font-size: 16px;
    position: absolute;
    top: 50%;
    z-index: 2;
    line-height: 0;
    color: #9ca2a9;
    transform: translateY(-50%);

    .search-button {
        height: 50px;
        width: 140px;
        font-size: 18px;
        transition: none;
        font-weight: 800;
        border-bottom-left-radius: 0;
        border-top-left-radius: 0;
        box-shadow: none;
        color: #ffffff;
        letter-spacing: 10px;
        background: #4E6EF2;
        border-color: #4E6EF2;
        outline: none;
        text-align: center;
        cursor: pointer;

        &:active {
            outline: none;
        }

        &:hover{
            color: #ffffff;
            background: #4E6EF2;
            border-color: #4E6EF2;
        }

         &:focus {
             color: #ffffff;
             background: #4E6EF2;
             border-color: #4E6EF2;
         }
        
    }
}
.search-icon {
    display: inline-block;
    color: inherit;
    font-style: normal;
    vertical-align: -.125em;
    text-align: center;
    text-rendering: optimizeLegibility;
}
</style>