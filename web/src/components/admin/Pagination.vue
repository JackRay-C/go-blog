<template>
  <div class="pagination">
    <div class="prev" :class="prevDisable ? 'disable' : ''" @click="prev">
      {{$t("page.prev")}}
    </div>

    <ul class="pagination-item">
      <li
        :class="{ active: current === 1 }"
        v-if="pageCount > 0"
        @click="current = 1"
      >
        1
      </li>
      <li class="more quickprev" v-if="showPrevMore">
        ...
      </li>
      <li
        v-for="i in pagers"
        :key="i"
        @click="go(i)"
        :class="{ active: i === current }"
      >
        {{ i }}
      </li>

      <li class="more quicknext" v-if="showNextMore">
        ...
      </li>
      <li
        :class="{ active: current === pageCount }"
        v-if="pageCount > 1"
        @click="current = pageCount"
      >
        {{ pageCount }}
      </li>
    </ul>

    <div class="next" :class="nextDisable ? 'disable' : ''" @click="next">
      {{$t("page.next")}}
    </div>
  </div>
</template>

<script>
export default {
  name: "Pagination",
  data() {
    return {
      current: this.pageNo
    };
  },
  props: {
    // 页码
    pageNo: {
      type: Number,
      default: 1
    },
    // 每页大小
    pageSize: {
      type: Number,
      default: 10
    },
    // 总页数
    pageCount: {
      type: Number,
      required: true
    },
    // 超过多少页显示省略号
    pagerCount: {
        type: Number,
        default: 10
    }
  },
  computed: {
    prevDisable() {
      return this.current <= 1;
    },
    nextDisable() {
      return this.current >= this.pageCount;
    },
    showPrevMore() {
      let halfPagerCount = (this.pagerCount - 1) / 2;
      if (this.pageCount > halfPagerCount) {
        if (this.current > this.pagerCount - halfPagerCount) {
          return true;
        }
      }
      return false;
    },
    showNextMore() {
      let halfPagerCount = (this.pagerCount - 1) / 2;
      if (this.pageCount > halfPagerCount) {
        if (this.current < this.pageCount - halfPagerCount) {
          return true;
        }
      }
      return false;
    },
    pagers() {
      const pagerCount = this.pagerCount;

      const showPrevMore = this.showPrevMore;
      const showNextMore = this.showNextMore;

      const pageCount = this.pageCount;
      const current = this.current;

      const array = [];

      if (showPrevMore && !showNextMore) {
        const startPage = pageCount - (pagerCount - 2);
        for (let index = startPage; index < pageCount; index++) {
          array.push(index);
        }
      } else if (!showPrevMore && showNextMore) {
        for (let index = 2; index < pagerCount; index++) {
          array.push(index);
        }
      } else if (showNextMore && showPrevMore) {
        const offset = Math.floor(pagerCount / 2) - 1;
        for (let index = current - offset; index <= current + offset; index++) {
          array.push(index);
        }
      } else {
        for (let index = 2; index < pageCount; index++) {
          array.push(index);
        }
      }
      return array;
    },
  },
  methods: {
    prev() {
      if (this.current > 1) {
        this.go(this.current - 1)
      } 
    },
    next() {
      if (this.current < this.pageCount) {
        this.go(this.current + 1)
      } 
    },
    go(index) {
      console.log("go: " + index)
      if(this.current !== index){
        this.current = index
        this.$emit("change", this.current);
      }
    }
  },
  watch: {
    pageNo(newVal) {
      this.current = newVal || 1;
    },
    current(newVal) {
      this.$emit("change", newVal)
    }
  },
};
</script>

<style lang="scss" scoped>
.pagination {
  width: 100%;
  box-sizing: border-box;
  display: flex;
  flex-direction: row;
  align-items: center;

  -webkit-box-orient: vertical;
  transition: all 0.3s;
  margin: 30px auto 20px;

  font-size: 1.2rem;
  font-weight: 400;
  letter-spacing: 0.2px;
  color: #15171a;
  position: relative;
}

.prev,
.next {
  margin: 0 10px;
  padding: 10px 24px;
  cursor: pointer;
  user-select: none;

  &:hover {
    background: #f1f3f4;
  }
}

.pagination-item {
  flex: auto;
  list-style: none;
  display: flex;
  flex-direction: row;
  box-sizing: border-box;
  justify-content: center;
  align-items: center;

  li {
    display: inline-block;

    cursor: pointer;
    border-radius: 3px;

    box-sizing: border-box;
    border: 1px solid transparent;
    text-align: center;
    margin: 0 5px;
    width: 40px;
    height: 40px;
    line-height: 40px;

    &:hover {
      background: #f6f7fa;
    }
    &:focus {
      background: #eaedf7;
    }
  }
}

.active {
  background: #f1f3f4;
  font-weight: 600;
}

.disable {
  color: rgba(0, 0, 0, 0.25);
  cursor: not-allowed;
  &:hover,
  &:focus {
    background: none;
    color: rgba(0, 0, 0, 0.25);
    border: none;
  }
}
</style>
