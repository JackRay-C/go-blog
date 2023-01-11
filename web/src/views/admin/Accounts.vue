<template>
  <div class="container">
    <div class="container-header">
      <div class="header-content">
        <h2 class="header-title">
          <router-link to="/admin/accounts">Accounts</router-link>
        </h2>
        <section class="view-actions">
          <div class="contentfilter"></div>
        </section>
      </div>
    </div>

    <div class="container-content">
      accounts

      <!-- <pagination /> -->
    </div>

    <div class="container-footer"></div>
  </div>
</template>

<script>
import Pagination from "@/components/admin/Pagination.vue";
import { listUsers } from "@/api/admin/user.js";

export default {
  components: { Pagination },
  name: "accounts",
  data() {
    return {
      total_page: 0,
      page_no: 1,
      page_size: 10,
    };
  },
  mounted() {
    this.fetchAccounts();
  },
  methods: {
    fetchAccounts() {
      listUsers({ page_no: this.page_no, page_size: this.page_size })
        .then((res) => {
          console.log(res);
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>

<style lang="scss" scoped></style>
