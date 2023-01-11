<template>
    <div class="container">
        <div class="container-header">
            <div class="header-content">
                <h2 class="header-title">
                    <router-link to="/admin/pages">Pages</router-link>
                </h2>
                <section class="view-actions">
                    <div class="contentfilter">
                        
                    </div>
                </section>
            </div>
        </div>
        <div class="container-content">
            pages
        </div>
    </div>
</template>

<script>
export default {
    name: 'pages',
    data() {
        return {
            total_page: 0,
            page_no: 1,
        }
    }
}
</script>

<style lang="scss" scoped>
.container {
    width: 100%;
    position: relative;
    flex-grow: 1;
    padding: 0 48px 48px 48px;
    margin: 0 auto;

    .container-header {
        margin: 0 -48px;
        padding: 0 48px;
        position: sticky;
        top: 0;
        background: #fff;
        z-index: 700;
        border-bottom: 1px solid #edeeef;

        .header-content {
            height: 95px;
            position: relative;
            flex-shrink: 0;
            display: flex;
            justify-content: space-between;
            align-items: center;

            .header-title {
                display: flex;
                align-items: center;
                overflow: hidden;
                margin: -3px 0 0;
                padding: 0;
                text-overflow: ellipsis;
                white-space: nowrap;
                font-size: 3.1rem;
                line-height: 1.3em;
                font-weight: 700;
                letter-spacing: 0;
                min-height: 35px;
                color: #15171a;
            }
        }
    }
}
</style>
