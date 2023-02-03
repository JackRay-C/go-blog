<!-- 
    usage:

    <select-menu options="selectdata" defaultValue="" placeholder="Page Size: " placeholderKeep onSelect="selectHandler" width="200" height="36" > </select-menu>

    selectdata: ["all", "draft", "published"]

 -->

<template>
    <div class="dropdown">
        <div class="dropdown-header" @blur="onBlur" @click="openSelect()">
            <div class="prefix" v-if="!this.selectedLabel || (this.selectedLabel && placeholderKeep)">
                {{ placeholder }} :
            </div>

            <div class="content" :class="['selected', {'placeholder': !selectedLabel}]" :style="`line-height: ${height-2}px; width: ${width- 37}px; height: ${height - 2}px`" :title="selectedLabel">
                {{ selectedLabel }}
            </div>

            <div class="suffix">
                <i class="el-icon-arrow-ip" v-if="showOptions"></i>
                <i class="el-icon-arrow-down" v-else></i>
            </div>
        </div>
        <transition name="fade">
            <div class="options" v-show="showOptions" @mouseleave="onLeave" :style="`top: ${height + 6}px;`">
                <div v-for="index,item in data" :key="index">
                    <div class="item" @click="selectedItem(index,item)">
                        <div>
                            {{ index }} : {{ item }}
                        </div>
                    </div>
                </div>
            </div>
        </transition>
    </div>
</template>

<script>
export default {
    name: "SelectMenu",
    model:{
        prop: "selectedText",
        event: "model"
    },
    data() {
        return {
            showOptions: false,
            hoverValue: null,
            selectedLabel: null,
        }
    },
    watch:{
        selectedText(){
            this.initSelector();
        },
        default() {
            this.initSelector();
        },
        options() {
            this.initSelector();
        }
    },
    created() {
        this.initSelector();
    },
    props: {
        // 下拉选项
        options: {
            type: Array,
            default: () =>{
                return []
            }
        },
        // 默认值
        default: {
            type: [Number, String, Object],
            default: null
        },
        // 下拉选项文本字段名
        label: {
            type: String,
            default: "label"
        },
        // 下拉选项的值字段名
        text: {
            type: String,
            default: "text"
        },
        // 选中的值
        selected: {
            type: [Number, String, Object],
            default: null
        },
        // 下拉选项宽度
        width: {
            type: Number,
            default: 200
        },
        // 下拉选项高度
        height: {
            type: Number,
            default: 36
        },
        // 默认文字
        placeholder: {
            type: String,
            default: ''
        },
        // 选择之后是否保留默认文字
        placeholderKeep: {
            type: Boolean,
            default: true
        }
    },

    methods: {
        selectedItem: (index, item) => {
            if(this.selected != item) {
                this.selected = item
                this.hoverValue = item
                this.showOptions = false
                this.$emit('model', item)
                this.$emit('onSelect', index, item)
            }
        },
        onBlur() {
            this.showOptions = !this.showOptions
        },
        onEnter(value) {
            this.hoverValue = value
        },
        onLeave() {
            this.hoverValue = null
        },
        openSelect() {
            this.showOptions = !this.showOptions
            // 设置选择项
            if(!this.hoverValue && this.selectedLabel) {
                const target = this.options.find(item => item[this.text] === this.selectedLabel)
                this.hoverValue = target[this.text];
            }
        },
        initSelector() {
            if(this.selectedText) {
                // 如果有selectedText，则设置hover项并查找lable
                this.hoverValue = this.selectedText
                const target = this.options.find(item => item[this.text] === this.selectedText)
                this.selectedLabel = target?target[this.label] : null
            } else {
                if (this.selectedText === '') {
                    this.hoverValue = ''
                    this.selectedLabel = ''
                } else if (this.default){
                    this.hoverValue = this.default[this.text]
                    const target = this.options.find(item => item[this.text] === this.default[this.text])
                    this.selectedLabel = target? target[this.label]:null
                    this.selectedText = target? target[this.text]: null
                    this.$emit('model', target)
                }else {
                    this.hoverValue = null
                    this.selectedLabel = null
                }
            }
        }
    }
}
</script>

<style lang="scss" scoped>

</style>