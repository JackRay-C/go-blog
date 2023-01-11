<template>
    <div class="container">
        <div class="container-header">
            <div class="header-content">
                <h2 class="header-title">
                    <router-link to="/admin/tags">Tags</router-link>
                </h2>
                <section class="view-actions">
                    <div class="contentfilter">
                        
                    </div>
                </section>
            </div>
        </div>
        subjects
    </div>
</template>

<script>
export default {
    name: 'Tags',
    data() {
        return {
            total_page: 0,
            page_no: 1,
        }
    }
}
</script>

<style lang="scss" scoped>
</style>
