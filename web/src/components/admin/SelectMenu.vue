<!-- 
    usage:

    <select-menu options="selectdata" defaultValue="" placeholder="Page Size: " placeholderKeep onSelect="selectHandler" width="200" height="36" > </select-menu>

    selectdata: ["all", "draft", "published"]

 -->

<template>
    <div class="select-menu" :style="`height: ${height}px`">
        <div class="header" tabindex="1" @blur="onBlur" @click="openSelect()" :class="[{focus: this.selectedText!=this.default},'u-select-wrap', 'hover']">
            <div class="prefix"  v-if="!this.selectedLabel || (this.selectedLabel && placeholderKeep)">
                {{ placeholder }} :
            </div>

            <div class="content" :class="['selected', {'placeholder': !selectedLabel}]" :title="selectedLabel">
                {{ selectedLabel }} 
            </div>

            <div class="suffix">
                <i class="el-icon-arrow-up" v-if="showOptions" > </i>
                <i class="el-icon-arrow-down" v-if="!showOptions"> </i>
            </div>
        </div>
        <transition name="fade">
            <div class="options" v-show="showOptions" :style="`top: ${height + 6}px;width: ${width}px`">
                <div class="options-content"   @mouseleave="onLeave" tabindex="2">
                    <div v-for="(item, index) in options" :key="index" class="item" @click="selectedItem(index,item)" :class="[{'option-selected': item[label] === selectedLabel, 'option-hover': item[text]===hoverValue}]" :title="item[label]" @mouseenter="onEnter(item)">
                        <div>
                            {{ item[text] }} : {{ item[label] }}
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
        selectedText: {
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
            default: false
        }
    },

    methods: {
        selectedItem(index, item){
            if(this.selectedLabel != item[this.label]) {
                this.selectedLabel = item[this.label]
                this.hoverValue = item[this.text]
                this.showOptions = false
                console.log(item)
                this.$emit('model', item[this.text])
                this.$emit('select', index, item)
            }
        },
        onBlur() {
            console.log("onblur")
            if (this.showOptions) {
                this.showOptions = !this.showOptions
            }
        },
        onEnter(value) {
            this.hoverValue = value[this.text]
        },
        onLeave() {
            this.hoverValue = null
            this.showOptions = false
        },
        openSelect() {
            this.showOptions = !this.showOptions
            // 设置选择项
            if(!this.hoverValue && this.selectedLabel) {
                const target = this.options.find(item => item[this.label] === this.selectedLabel)
                console.log(target)
                this.hoverValue = target[this.text];
            }
        },
        initSelector() {
            console.log(this.selectedText === this.default)
            if(this.selectedText) {
                this.hoverValue = this.selectedText
                const target = this.options.find(item => item[this.text] === this.selectedText)
                this.selectedLabel = target ? target[this.label] : null
            } else {
                if (this.default){
                    this.hoverValue = this.default[this.text]
                    const target = this.options.find(item => item[this.text] === this.default)
                    this.selectedLabel = target? target[this.label]:null
                    this.$emit('model', target[this.text])
                }else {
                    this.hoverValue = null
                    this.selectedLabel = null
                }
            }
            // if(this.default >= 0) {
            //     const target = this.options.find(item => item[this.text] === this.default)
            //     this.hoverValue = this.default
            //     this.selectedLabel = target[this.label]
            //     this.$emit('model', target)
            // } else {
            //     if(this.selectedText >= 0) {
            //         this.hoverValue = this.selectedText
            //         const target = this.options.find(item => item[this.text] === this.selectedText)
            //         this.selectedLabel = target?target[this.label] : null
            //     } else {
            //         this.hoverValue = null
            //         this.selectedLabel = null
            //         this.selected = null
            //     }
            // }
            // if(this.selectedText) {
            //     // 如果有selectedText，则设置hover项并查找lable
            //     this.hoverValue = this.selectedText
            //     const target = this.options.find(item => item[this.text] === this.selectedText)
            //     this.selectedLabel = target?target[this.label] : null
            // } else {
            //     console.log(this.default)
            //     console.log(this.selectedText)
            //     if (this.selectedText === '') {
            //         this.hoverValue = ''
            //         this.selectedLabel = ''
            //     } else if (this.default >= 0){
            //         this.hoverValue = this.default[this.text]
            //         const target = this.options.find(item => item[this.text] === this.default)
            //         this.selectedLabel = target? target[this.label]:null
            //         this.selectedText = target? target[this.text]: null
            //         this.$emit('model', target)
            //     }else {
            //         this.hoverValue = null
            //         this.selectedLabel = null
            //     }
            // }
        }
    }
}
</script>

<style lang="scss" scoped>
.select-menu {
    position: relative;
    display: inline-block;
    font-size: 1.35rem;
    font-weight: 400;
    color: #394047;
    letter-spacing: 0.2px;
}


.header {
    position: relative;
    display: inline-flex;
    align-content: space-around;
    background: #ffffff;
    font-size: 1.35rem;
    font-weight: 400;
    color: #394047;
    letter-spacing: 0.2px;
    padding: 6px 12px;
    margin-right: 8px;
    outline: none;
    border: 1px solid transparent;
    border-radius: 2px;
    line-height: 1.75;
    user-select: none;

    .prefix{
        display: inline-block;
        margin: 0 4px;
    }
    .content {
        display: inline-block;
        overflow: hidden;
        text-overflow: ellipsis;
        transition: all 0.2s ease;
        margin: 0 4px;
    }

    .suffix {
        display: inline-block;
        margin: 0 4px;
    }
}

.u-select-wrap {
    cursor: pointer;
    transition: all 0.2s ease;
}


.options {
    margin-top: 6px;
    padding: 6px 0;
    border: none !important;
    font-size: 1.35rem;
    box-shadow: 0 0 0 1px rgb(0 0 0 / 4%), 0 7px 20px -5px rgb(0 0 0 / 15%);
    border-radius: 5px;
    position: absolute;
    z-index: 9999;
    box-sizing: border-box;
    background: #fff;
    line-height: 1.75;
    overflow: hidden;
    color: inherit;
    height: auto;

    .option-content {
        position: relative;
        display: block;
        overflow-x: hidden;
        max-height: 50px;
        overflow-y: auto;
        box-sizing: border-box;
        list-style: none;
        margin: 0;
        padding: 0;
        user-select: none;
    }
   
    .item {
        position: relative;
        display: block;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        color: #394047;
        background: transparent;
        margin-bottom: 0;
        padding: 6px 14px;
        cursor: pointer;
        line-height: 1.35em;
        line-height: 1.4em;

        &.option-selected {
          font-weight: 700;
        }

        &:not(.selected):hover {
          background: #f4f5f5;
        }
        .option-hover {
            background: #f4f5f5;
        }

    }
}

.hover {
    &:hover {
        font-weight: 600;
        background: #f1f3f4;
        border: 1px solid #f1f3f4;
    }
}
.focus {
    color: #30cf43;
    font-weight: 600;
    background: #f1f3f4;
    border: 1px solid #f1f3f4;
}

.fade-enter-active, .fade-leave-active {
    transition: all .3s;
}

.fade-enter, .fade-leave-to {
    opacity: 0;

}
</style>