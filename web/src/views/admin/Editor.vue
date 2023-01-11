<template>
    <div class="editor" v-loading="loading">
        <div class="flex flex-row">
            <transition name="fade">
                <div class="editor-left">
                    <div class="editor-header">
                        <div class="editor-header-content">
                            <div class="editor-header-back" @click="back">
                                <div class="back-icon">
                                    <i class="el-icon-back"></i>
                                    {{ $t("sidebar.posts") }}
                                </div>
                            </div>
                            <div class="editor-header-info">
                                <div class="editor-header-title">
                                    <input class="title-text" v-model="post.title" />
                                </div>
                                <div class="editor-header-status">
                                    <span class="status"><i class="el-icon-upload"></i> 已保存</span>
                                </div>
                            </div>
                            <div class="editor-header-action" :style="{ 'margin-right': drawer ? '20px' : '100px' }">
                                <!-- <div class="editor-header-action-item">
                                    <div class="save-button" @click="publish">
                                        <span>发表</span>
                                    </div>
                                </div> -->
                                <div class="editor-header-action-item">
                                    <div class="update-button" :class="{ disable: !change }" @click="update">
                                        更新
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="editor-root">
                        <div id="editor" class="editor-content"></div>
                    </div>
                </div>
            </transition>
            <transition name="fade">
                <div class="editor-right" v-show="drawer">
                    <div class="editor-drawer">
                        <div class="setting-menu-header">
                            <span>设置</span>
                        </div>
                        <div class="setting-menu-content">
                            <el-form label-position="top" label-width="80px" :model="post" size="small">
                                <el-form-item>
                                    <label>作者</label>
                                    <el-input disabled v-model="owner.nickname"></el-input>
                                </el-form-item>

                                <el-form-item>
                                    <label>专题</label>
                                    <el-select v-model="post.subject_id" filterable remote
                                        :remote-method="searchSubject" placeholder="请选择" style="width: 100%"
                                        size="small" allow-create>
                                        <el-option v-for="item in subjectOptions" :key="item.id" :label="item.title"
                                            :value="item.id">
                                        </el-option>
                                    </el-select>
                                </el-form-item>
                                <el-form-item>
                                    <label>公开  </label>
                                    <el-switch v-model="post.visibility" :active-value="2" :inactive-value="1" 
                                        active-color="#13ce66" inactive-color="#ff4949">
                                    </el-switch>
                                </el-form-item>

                                <el-form-item>
                                    <label>发布日期
                                    </label>
                                    <!-- <el-row :gutter="20" v-if="post.status===1">
                                        <el-col :span="24">
                                            <el-date-picker type="datetime" placeholder="选择日期" class="border-radius" value-format="timestamp"
                                                v-model="post.created_at" style="width: 100%"></el-date-picker>
                                        </el-col>

                                    </el-row> -->
                                    <el-row :gutter="20">
                                        <el-col :span="24">
                                            <el-date-picker type="datetime" placeholder="选择日期" class="border-radius"
                                                v-model="post.published_at"
                                                format="yyyy-MM-dd HH:mm:ss"
                                                value-format="timestamp" style="width: 100%"></el-date-picker>
                                        </el-col>
                                    </el-row>
                                </el-form-item>

                                <el-form-item>
                                    <label>标签 </label>
                                    <el-row :gutter="20">
                                        <el-col>
                                            <el-select v-model="tags" style="width: 100%" multiple filterable
                                                allow-create placeholder="请选择文章标签" size="small">
                                                <el-option v-for="item in tagOptions" :key="item.value"
                                                    :label="item.name" :value="item.id">
                                                </el-option>
                                            </el-select>
                                        </el-col>
                                    </el-row>
                                </el-form-item>

                                <el-form-item>
                                    <label>描述 </label>
                                    <el-input type="textarea" :rows="10" v-model="post.description"></el-input>
                                </el-form-item>

                            </el-form>
                        </div>
                        <div class="setting-menu-footer">
                            <el-row :gutter="20" style="padding:24px">
                                <el-col :span="16">
                                            <el-button type="primary" class="publish-button" @click="publish">立即发布</el-button>
                                        </el-col>
                                        <el-col :span="8">
                                            <el-button class="publish-delete-button" @click="del">删除</el-button>
                                        </el-col>
                            </el-row>
                                        
                                
                        </div>
                    </div>
                </div>
            </transition>
        </div>
        <button class="setting-toggle" @click="drawer = !drawer">
            <!-- <span class="settings-menu-open"><svg fill="none" viewBox="0 0 24 24"><path d="M21 2.5H3c-.828 0-1.5.608-1.5 1.357v16.286c0 .75.672 1.357 1.5 1.357h18c.828 0 1.5-.608 1.5-1.357V3.857c0-.75-.672-1.357-1.5-1.357zm-4.5 0v19" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path><path d="M16.5 2.5v19h6v-19h-6z" fill="#15171A"></path></svg></span> -->
            <span>
                <svg viewBox="0 0 24 24">
                    <path
                        d="M21 2.5H3c-.828 0-1.5.608-1.5 1.357v16.286c0 .75.672 1.357 1.5 1.357h18c.828 0 1.5-.608 1.5-1.357V3.857c0-.75-.672-1.357-1.5-1.357zm-4.5 0v19">
                    </path>
                    <path v-show="drawer" d="M16.5 2.5v19h6v-19h-6z" fill="#15171A"></path>
                </svg>
            </span>
        </button>

        <!-- <el-drawer :visible.sync="drawer" direction="rtl">
            <div class="setting-form">
                <el-form label-position="top" label-width="80px" :model="post">
                    <el-form-item label="标题">
                        <el-input v-model="post.title"></el-input>
                    </el-form-item>
                    <el-form-item label="标签">
                        <el-input v-model="post.title"></el-input>
                    </el-form-item>
                    <el-form-item label="描述">
                        <el-input type="textarea" v-model="post.dsecription"></el-input>
                    </el-form-item>
                </el-form>
            </div>
        </el-drawer> -->
    </div>
</template>

<script>
import Vditor from "vditor";
import "@/components/Vditor/css/index.scss";
import { getPost } from "@/api/admin/post";
import api from "@/api/admin/api";
import { deletePost, listPostTags, putPost } from "../../api/admin/post";
import { listSubjects } from "../../api/admin/subject";
import { getUserInfo } from "../../api/admin/user";

export default {
    data() {
        return {
            post: {
                
            },
            switch1: true,
            owner: {},
            tags: [],
            tagOptions: [],
            visibleOptions: [
                { text: "公开", value: 1 },
                { text: "私有", value: 2 },
            ],
            subjectOptions: [],
            editor: "",
            placeholder: "请输入内容",
            loading: false,
            change: false,
            drawer: false,
            post_title_edit: false,
        };
    },
    created() {
        this.fetchPost();
    },
    watch: {
        $route: "fetchPost",
    },
    mounted() { },
    methods: {
        fetchPost() {
            this.loading = true;
            getPost(this.$route.params.id)
                .then((res) => {
                    if (res.code === 200) {
                        console.log(res);
                        this.post = res.data;
                        this.$watch(
                            "post",
                            function () {
                                this.change = true;
                            },
                            { deep: true }
                        );
                        Promise.all([
                            listPostTags(this.post.id),
                            getUserInfo(),
                            listSubjects({ page_no: 1, page_size: 10 }),
                        ])
                            .then((results) => {
                                console.log(results);
                                results[0].data.forEach((t) => {
                                    this.tags.push(t.id);
                                });
                                this.tagOptions = results[0].data;

                                this.owner = results[1].data;
                                console.log(this.owner);

                                this.subjectOptions = results[2].data.list;

                                console.log(this.post.subject_id);
                                console.log(
                                    results[2].data.list.find((item) => {
                                        return item.id === this.post.subject_id;
                                    })
                                );
                                this.post.subject = results[2].data.list.find((item) => {
                                    return item.id === this.post.subject_id;
                                });
                                // this.loading = false
                                this.editor = new Vditor("editor", {
                                    toolbar: [
                                        "|",
                                        {
                                            name: "headings",
                                            hotkey: "⌘H",
                                            icon: '<svg><use xlink:href="#vditor-icon-headings"></use></svg>',
                                            tipPosition: "w",
                                        },
                                        {
                                            name: "quote",
                                            icon: '<svg viewBox="0 0 1024 1024"><path d="M894.6 907.1H605.4c-32.6 0-59-26.4-59-59V608.2l-4-14.9c0-315.9 125.5-485.1 376.5-507.5v59.8C752.7 180.4 711.3 315.8 711.3 442.4v41.2l31.5 12.3h151.8c32.6 0 59 26.4 59 59v293.2c0 32.5-26.4 59-59 59z m-472 0H133.4c-32.6 0-59-26.4-59-59V608.2l-4-14.9c0-315.9 125.5-485.1 376.5-507.5v59.8C280.7 180.4 239.3 315.8 239.3 442.4v41.2l31.5 12.3h151.8c32.6 0 59 26.4 59 59v293.2c0 32.5-26.4 59-59 59z"></path></svg>',
                                            prefix: "> ",
                                            tipPosition: "s",
                                            hotkey: "⌘;",
                                        },
                                        "|",
                                        {
                                            name: "bold",
                                            tipPosition: "s",
                                            hotkey: "⌘B",
                                            icon: '<svg viewBox="0 0 1024 1024"><path d="M707.872 484.64A254.88 254.88 0 0 0 768 320c0-141.152-114.848-256-256-256H192v896h384c141.152 0 256-114.848 256-256a256.096 256.096 0 0 0-124.128-219.36zM384 192h101.504c55.968 0 101.504 57.408 101.504 128s-45.536 128-101.504 128H384V192z m159.008 640H384v-256h159.008c58.464 0 106.016 57.408 106.016 128s-47.552 128-106.016 128z"></path></svg>',
                                            prefix: "**",
                                            suffix: "**",
                                        },
                                        {
                                            name: "outline",
                                            tipPosition: "s",
                                            icon: '<svg viewBox="0 0 1024 1024"><path d="M704 64l128 0 0 416c0 159.072-143.264 288-320 288s-320-128.928-320-288l0-416 128 0 0 416c0 40.16 18.24 78.688 51.36 108.512 36.896 33.216 86.848 51.488 140.64 51.488s103.744-18.304 140.64-51.488c33.12-29.792 51.36-68.352 51.36-108.512l0-416zM192 832l640 0 0 128-640 0z"></path></svg>',
                                        },
                                        {
                                            name: "italic",
                                            tipPosition: "s",
                                            icon: '<svg viewBox="0 0 1024 1024"><path d="M896 64v64h-128L448 896h128v64H128v-64h128L576 128h-128V64z"></path></svg>',
                                        },
                                        {
                                            name: "strike",
                                            hotkey: "⌘D",
                                            icon: '<svg viewBox="0 0 1024 1024"><path d="M1024 512v64h-234.496c27.52 38.496 42.496 82.688 42.496 128 0 70.88-36.672 139.04-100.576 186.976C672.064 935.488 594.144 960 512 960s-160.064-24.512-219.424-69.024C228.64 843.04 192 774.88 192 704h128c0 69.376 87.936 128 192 128s192-58.624 192-128-87.936-128-192-128H0v-64h299.52a385.984 385.984 0 0 1-6.944-5.024C228.64 459.04 192 390.88 192 320s36.672-139.04 100.576-186.976C351.936 88.512 429.856 64 512 64s160.064 24.512 219.424 69.024C795.328 180.96 832 249.12 832 320h-128c0-69.376-87.936-128-192-128s-192 58.624-192 128 87.936 128 192 128c78.976 0 154.048 22.688 212.48 64H1024z"></path></svg>',
                                            tipPosition: "e",
                                            tip: "删除线",
                                            prefix: "~~",
                                            suffix: "~~",
                                        },
                                        {
                                            name: "list",
                                            tipPosition: "se",
                                            prefix: "* ",
                                            hotkey: "⌘L",
                                        },
                                        {
                                            name: "ordered-list",
                                            tipPosition: "se",
                                            prefix: "1. ",
                                            hotkey: "⌘O",
                                        },

                                        "|",
                                        {
                                            name: "link",
                                            icon: '<svg viewBox="0 0 1024 1024"><path d="M440.224 635.776a51.84 51.84 0 0 1-36.768-15.232c-95.136-95.136-95.136-249.92 0-345.056l192-192C641.536 37.408 702.816 12.032 768 12.032s126.432 25.376 172.544 71.456c95.136 95.136 95.136 249.92 0 345.056l-87.776 87.776a51.968 51.968 0 1 1-73.536-73.536l87.776-87.776a140.16 140.16 0 0 0 0-197.984c-26.432-26.432-61.6-40.992-99.008-40.992s-72.544 14.56-99.008 40.992l-192 192a140.16 140.16 0 0 0 0 197.984 51.968 51.968 0 0 1-36.768 88.768z"></path><path d="M256 1012a242.4 242.4 0 0 1-172.544-71.456c-95.136-95.136-95.136-249.92 0-345.056l87.776-87.776a51.968 51.968 0 1 1 73.536 73.536l-87.776 87.776a140.16 140.16 0 0 0 0 197.984c26.432 26.432 61.6 40.992 99.008 40.992s72.544-14.56 99.008-40.992l192-192a140.16 140.16 0 0 0 0-197.984 51.968 51.968 0 1 1 73.536-73.536c95.136 95.136 95.136 249.92 0 345.056l-192 192A242.4 242.4 0 0 1 256 1012z"></path></svg>',
                                            tipPosition: "se",
                                        },
                                        {
                                            name: "upload",
                                            icon: '<svg viewBox="0 0 1024 1024"><path d="M959.877 128l0.123 0.123v767.775l-0.123 0.122H64.102l-0.122-0.122V128.123l0.122-0.123h895.775zM960 64H64C28.795 64 0 92.795 0 128v768c0 35.205 28.795 64 64 64h896c35.205 0 64-28.795 64-64V128c0-35.205-28.795-64-64-64zM832 288.01c0 53.023-42.988 96.01-96.01 96.01s-96.01-42.987-96.01-96.01S682.967 192 735.99 192 832 234.988 832 288.01zM896 832H128V704l224.01-384 256 320h64l224.01-192z"></path></svg>',
                                            tipPosition: "se",
                                        },
                                        "record",
                                        "table",
                                        "|",
                                        {
                                            name: "inline-code",
                                            tipPosition: "se",
                                        },
                                        {
                                            hotkey: "⌘U",
                                            icon: '<svg><use xlink:href="#vditor-icon-code"></use></svg>',
                                            name: "code",
                                            prefix: "```",
                                            suffix: "\n```",
                                            tipPosition: "se",
                                        },
                                        {
                                            icon: '<svg><use xlink:href="#vditor-icon-code-theme"></use></svg>',
                                            name: "code-theme",
                                            tipPosition: "w",
                                        },
                                        {
                                            name: "line",
                                            tipPosition: "se",
                                            prefix: "---",
                                        },
                                        {
                                            name: "emoji",
                                            icon: '<svg viewBox="0 0 1024 1024"><path d="M512 1024C230.4 1024 0 793.6 0 512S230.4 0 512 0s512 230.4 512 512-230.4 512-512 512z m0-102.4c226.742857 0 409.6-182.857143 409.6-409.6S738.742857 102.4 512 102.4 102.4 285.257143 102.4 512s182.857143 409.6 409.6 409.6z m-204.8-358.4h409.6c0 113.371429-91.428571 204.8-204.8 204.8s-204.8-91.428571-204.8-204.8z m0-102.4c-43.885714 0-76.8-32.914286-76.8-76.8s32.914286-76.8 76.8-76.8 76.8 32.914286 76.8 76.8-32.914286 76.8-76.8 76.8z m409.6 0c-43.885714 0-76.8-32.914286-76.8-76.8s32.914286-76.8 76.8-76.8c43.885714 0 76.8 32.914286 76.8 76.8s-32.914286 76.8-76.8 76.8z"></path></svg>',
                                            tipPosition: "se",
                                        },
                                        "|",
                                        {
                                            name: "undo",
                                            icon: '<svg viewBox="0 0 1024 1024"><path d="M512 64A510.272 510.272 0 0 0 149.984 213.984L0.032 64v384h384L240.512 304.48A382.784 382.784 0 0 1 512.032 192c212.064 0 384 171.936 384 384 0 114.688-50.304 217.632-130.016 288l84.672 96a510.72 510.72 0 0 0 173.344-384c0-282.784-229.216-512-512-512z"></path></svg>',
                                            tipPosition: "se",
                                        },
                                        {
                                            name: "redo",
                                            icon: '<svg viewBox="0 0 1024 1024"><path d="M0.00032 576a510.72 510.72 0 0 0 173.344 384l84.672-96A383.136 383.136 0 0 1 128.00032 576C128.00032 363.936 299.93632 192 512.00032 192c106.048 0 202.048 42.976 271.52 112.48L640.00032 448h384V64l-149.984 149.984A510.272 510.272 0 0 0 512.00032 64C229.21632 64 0.00032 293.216 0.00032 576z"></path></svg>',
                                            tipPosition: "se",
                                        },
                                        {
                                            name: "more",
                                            tipPosition: "w",
                                            toolbar: [
                                                "edit-mode",
                                                "content-theme",
                                                "export",
                                                "preview",
                                            ],
                                        },
                                        "|",
                                        // "help",
                                        // "edit-mode",
                                        // "content-theme",
                                        // "export",
                                        // "preview",
                                        // "|",
                                    ],
                                    toolbarConfig: {
                                        pin: true,
                                    },
                                    cache: {
                                        enable: false,
                                    },
                                    // cdn: "https://fastly.jsdelivr.net/npm/vditor",
                                    placeholder: this.placeholder,
                                    tab: "\t",
                                    typewriterMode: true,
                                    mode: "ir",
                                    // input: debounce(this.onChange, 500),
                                    // blur: debounce(this.onBlur, 500),
                                    input: () => {
                                        this.post.markdown_content = this.editor.getValue();
                                        this.post.html_content = this.editor.getHTML();
                                    },
                                    // blur: (s)=> {
                                    //     if(s !== this.post.markdown_content) {
                                    //         this.change = true
                                    //     }else {
                                    //         this.change = false
                                    //     }
                                    // },
                                    esc: () => { },
                                    after: () => {
                                        // this.editor.setValue(this.post.markdown_content);
                                        this.loading = false;
                                        console.log("vditor init success");
                                    },
                                    value: this.post.markdown_content,
                                    preview: {
                                        hljs: {
                                            enable: true,
                                            lineNumber: true,
                                            style: "monokai",
                                        },
                                    },
                                    upload: {
                                        url: process.env.VUE_APP_BASE_URL + api.files,
                                        max: 10 * 1024 * 1024,
                                        linkToImgUrl: "",
                                        linkToImgCallback: (res) => {
                                            console.log(res);
                                        },
                                        linkToImgFormat: (res) => {
                                            console.log(res);
                                        },
                                        headers: {
                                            token: localStorage.getItem("access_token"),
                                        },
                                        withCredentials: true,
                                        fieldName: "file",
                                        multiple: false,
                                        format: (files, res) => {
                                            var resp = JSON.parse(res);
                                            if (resp.code === 200) {
                                                return JSON.stringify({
                                                    msg: "",
                                                    code: 0,
                                                    data: {
                                                        errFiles: [],
                                                        succMap: {
                                                            [files[0].name]: resp.data.access_url,
                                                        },
                                                    },
                                                });
                                            }
                                        },
                                    },
                                });
                            })
                            .catch((err) => {
                                console.log(err);
                                this.loading = false;
                            });
                    } else {
                        this.err = res.message;
                        this.$alert(res.message, "Error " + res.code, {
                            confirmButtonText: "确定",
                            callback: () => {
                                this.$router.go(-1);
                            },
                        });
                    }
                })
                .catch((err) => {
                    console.log(err);
                    this.loading = false;
                    this.$alert(err);
                });

            // listSubjects({page_no: 1, page_size: 10}).then(res => {
            //     console.log(res)
            //     if(res.code === 200) {
            //         this.subjectOptions = res.data
            //     }
            // }).catch(err => {
            //     console.log(err)
            // })
        },
        onSwitch(val){
            console.log(val)
        },
        searchSubject(query) {
            if(query !== '') {
                this.loading = true
                listSubjects({page_no: 1, page_size: 10, search: query}).then(res => {
                    console.log(res)
                    this.loading = false
                    if(res.code === 200) {
                        this.subjectOptions = res.data.list
                    }
                }).catch(err => {
                    console.log(err)
                })
            }
        },
        back() {
            // this.$router.push("/admin/posts")
            if (this.change) {
                this.$confirm("内容已经修改，是否保存？", "", {
                    confirmButtonText: "保存",
                    cancelButtonText: "放弃修改",
                    distinguishCancelAndClose: true,
                })
                    .then(() => {
                        // 更新
                        putPost(this.$route.params.id, this.post)
                            .then((res) => {
                                if (res.code === 200) {
                                    console.log("更新成功");
                                    this.change = false;
                                    console.log(res);
                                    // 跳转到前一个
                                    this.$router.push("/admin/posts");
                                }
                            })
                            .catch((err) => {
                                console.log(err);
                            });
                    })
                    .catch((action) => {
                        console.log(action);
                        this.$router.push("/admin/posts");
                    });
            } else {
                this.$router.push("/admin/posts");
            }
        },
        del(){
            
                this.$confirm("是否确定删除？", "提示", {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    distinguishCancelAndClose: true,
                })
                    .then(() => {
                        // 更新
                        deletePost(this.$route.params.id)
                            .then((res) => {
                                if (res.code === 200) {
                                    console.log("删除成功");
                                    console.log(res);
                                    // 跳转到前一个
                                    this.back()
                                }
                            })
                            .catch((err) => {
                                console.log(err);
                            });
                    })
                    .catch((action) => {
                        console.log(action);
                    });
            
        },
        publish() {
            // 将博客状态由1改为2
            this.post.status = 2;
            putPost(this.$route.params.id, this.post)
                .then((res) => {
                    if (res.code === 200) {
                        console.log("发布成功");
                        this.$notify.success({
                            title: "Success",
                            message: "发布成功. ",
                            offset: 50,
                        });
                    } else{
                        console.log(res)
                    }
                })
                .catch((err) => {
                    console.log("发布失败.");
                    this.$notify.success({
                        title: "发布博客失败",
                        message: err.message,
                        offset: 50,
                    });
                });
        },
        update() {
            if (this.change === true) {
                putPost(this.$route.params.id, this.post)
                    .then((res) => {
                        console.log(res);
                        if (res.code === 200) {
                            console.log("更新成功");
                            this.$notify.success({
                                title: "Success",
                                message: "更新成功. ",
                                offset: 100,
                            });
                            this.change = false;
                        }else{
                        console.log(res)
                    }
                    })
                    .catch((err) => {
                        console.log("更新失败.");
                        this.$notify.success({
                            title: "更新博客失败",
                            message: "更新博客失败: " + err,
                            offset: 100,
                        });
                    });
            }
        },
        saveDraft() {
            // 保存草稿
        },
        parseTimestamp(num){
            return new Date(parseInt(num)*1000)
        }
    },
};
</script>

<style lang="scss" scoped>
.editor {
    width: 100%;
    height: 100%;
    background: #fff;
    margin: 0 auto;
    min-height: 100vh;
}

.flex {
    display: flex;
}

.flex-row {
    flex-direction: row;
}

.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.03s ease;
}

.fade-enter,
.fade-leave-to {
    opacity: 0;
}

button {
    padding: 0;
    outline: none;
    border: none;
    background: transparent;
    box-shadow: none;
    line-height: inherit;
}

.setting-toggle {
    -webkit-appearance: button;
    position: absolute;
    top: 12px;
    right: 28px;
    z-index: 9999;
    margin-right: 0 !important;
    display: inline-block;
    outline: none;
    text-decoration: none;
    user-select: none;
    fill: #000;
    background: #fff;
    font-weight: 500;
    border-radius: 3px;
    transition: all 0.2s ease;
    transition-property: color, border-color, background, width, height,
        box-shadow;
    cursor: pointer;

    &:hover {
        color: #394047;
        background: #ebeef0;
    }

    & span {
        color: #394047;
        display: flex;
        align-items: center;
        justify-content: center;
        overflow: hidden;
        padding: 0 14px;
        height: 33px;
        font-size: 1.4rem;
        line-height: 33px;
        text-align: center;
        letter-spacing: 0.2px;
        border-radius: 3px;
        white-space: nowrap;
        text-overflow: ellipsis;

        & svg {
            height: 16px;
            width: 16px;
            margin: 0;
            fill: none;

            path {
                stroke: #15171a;
                stroke-width: 1.5;
                stroke-linecap: round;
                stroke-linejoin: round;
            }
        }
    }
}

.editor-left {
    flex: 1;
}

.editor-right {
    display: block;
    background: #fff;
    z-index: 999;
    height: 100vh;
    min-width: 300px;
    overflow-x: visible;
    overflow-y: auto;
    border-left: 1px solid #e6e9eb;
    box-sizing: border-box;
}

.editor-drawer {
    position: relative;
    top: auto;
    right: auto;
    bottom: auto;
    height: 100vh;
    width: 100%;
    width: 380px;
    overflow: auto;
    background: #fff;
    opacity: 1;
    display: flex;
    flex-direction: column;
    align-content: space-around;
    justify-content: space-around;
}

.setting-menu-header {
    height: 56px;
    position: flexed;
    display: flex;
    width: 100%;
    max-width: 380px;
    padding: 20px 24px 24px;
    letter-spacing: 0.2px;
    justify-content: space-between;
    align-items: center;
    z-index: 1;
    color: #394047;
    background: #fff;

    & span {
        margin: 0;
        font-size: 1.4rem;
        font-weight: 500;
        line-height: 1.3;
    }
}

.setting-menu-content {
    // margin-top: 20px;
    flex: 1;
    padding: 20px 28px 33px 24px;
}
.setting-menu-footer {
    height: 100px;
}

.editor-header {
    // border-bottom: 1px solid #f4f5f5;
    position: relative;
    height: 56px;
    background: #fff;
    z-index: 999;
}

.border-radius {
    border-radius: 1px;
}

.editor-header-content {
    width: 100%;
    position: relative;
    height: 56px;
    display: flex;
    -webkit-box-flex: 1;
    flex: 1 1 auto;
}

.editor-header-back {
    display: flex;
    width: 100px;
    // border-right: 1px solid #e6eaea;
    box-sizing: border-box;
    box-align: center;
    margin-left: 1.2rem;
    align-items: center;
    -webkit-box-pack: center;
    justify-content: center;
    // height: 33px;
}

.back-icon {
    display: flex;
    align-items: center;

    padding: 0 18px;
    height: 33px;
    font-size: 1.4rem;
    line-height: 33px;
    text-align: center;
    letter-spacing: 2px;
    border-radius: 3px;
    white-space: nowrap;
    text-overflow: ellipsis;
    cursor: pointer;
    color: #394047;
    background: #ffffff;
    font-weight: 500;
    outline: none;
    text-decoration: none;
    user-select: none;
    transition: all 0.2s ease;

    &:hover {
        text-decoration: none;
        background: #ebeef0;
        color: #394047;
        transition: background 0.1s, color 0.1s;
    }
}

.editor-header-info {
    display: flex;
    align-items: center;
    height: 56px;
    // width: 40%;
}

.editor-header-title {
    display: flex;
    box-align: center;
    align-items: center;
    -webkit-box-pack: center;
    justify-content: center;
    height: 33px;

    .title-text {
        border: none;
        text-decoration: none;
        outline: none;
        cursor: pointer;
        font-weight: 500;
        font-size: 1.5rem;
        color: #394047;
        background: #ffffff;
        letter-spacing: 0.2px;
        display: inline-block;
        white-space: nowrap;
        border: 1px solid transparent;
        padding: 5px 8px;
        margin: 0 15px 0 0;

        &:focus {
            border: none;
            // outline: 1px solid rgb(179, 180, 179);
            outline: none;
            border-radius: 4px;
            border: 1px solid #dddedf;
        }
    }
}

.editor-header-status {
    display: flex;
    height: 20px;
    // margin-left: 16px;
    align-items: center;
    justify-content: center;
    font-size: 12px;
    color: #8a8f8d;
}

.editor-header-action {
    // position: absolute;
    display: flex;
    align-items: center;
    justify-content: flex-end;
    // width: 100px;
    flex: 1;
    height: 56px;
}

.editor-header-action-item {
    padding-left: 8px;
}

.editor-header-action-item-more {
    // display: flex;
    position: relative;
    // -webkit-box-align: center;
    // align-items: center;
    // justify-content: center;
    // margin-left: 20px;
}

.more-icon {
    display: flex;
    align-items: center;

    padding: 0 20px;
    height: 34px;
    font-size: 1.35rem;
    line-height: 34px;
    text-align: center;
    letter-spacing: 0.2px;
    border-radius: 3px;
    white-space: nowrap;
    text-overflow: ellipsis;
    cursor: pointer;
    font-weight: 500;
    outline: none;
    text-decoration: none;
    user-select: none;
    transition: all 0.2s ease;

    &:hover {
        text-decoration: none;
        background: #ebeef0;
        color: #394047;
        transition: background 0.1s, color 0.1s;
    }
}

.save-button {
    display: flex;
    align-items: center;

    padding: 0 20px;
    height: 34px;
    font-size: 1.35rem;
    line-height: 34px;
    text-align: center;
    letter-spacing: 0.2px;
    border-radius: 3px;
    white-space: nowrap;
    text-overflow: ellipsis;
    cursor: pointer;
    color: #394047;
    background: #fff;
    font-weight: 500;
    outline: none;
    text-decoration: none;
    user-select: none;
    transition: all 0.2s ease;

    &:hover {
        text-decoration: none;
        color: #394047;
        background: #ebeef0;
        transition: background 0.1s, color 0.1s;
    }
}

.update-button {
    display: flex;
    align-items: center;
    padding: 0 20px;
    height: 33px;
    font-size: 1.4rem;
    line-height: 33px;
    text-align: center;
    letter-spacing: 2px;
    border-radius: 3px;
    white-space: nowrap;
    text-overflow: ellipsis;
    cursor: pointer;
    color: #394047;
    background: #ffffff;
    font-weight: 500;
    outline: none;
    text-decoration: none;
    user-select: none;
    transition: all 0.2s ease;

    &:hover {
        text-decoration: none;
        color: #394047;
        background: #ebeef0;
        transition: background 0.1s, color 0.1s;
    }
}

.update-button.disable {
    // background: #e4e6ea;
    color: #767a7d;
    cursor: not-allowed;
}

.publish-button {
    width: 100%;
    color: #f7f7f7;
    background: #394047;
    border: 1px solid #ccc;
    height: 35px;

    &:hover{
        background: #394047;
        color: #fff;
        border: 1px solid #394047;
    }
}

.publish-delete-button {
    width: 100%;
    background: #f64949;
    color: #fff;
    border: 1px solid #ccc;
    height: 35px;
    &:hover{
        background: #f64949;
        border: 1px solid #ccc;
        color: #fff;
    }
}

.editor-root {
    width: 100%;
    flex: 1;
    background: #f3f5f7;
}
</style>
