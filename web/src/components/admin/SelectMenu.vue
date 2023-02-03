<!-- 
    usage:

    <select-menu options="selectdata" defaultValue="" placeholder="Page Size: " placeholderKeep onSelect="selectHandler" width="200" height="36" > </select-menu>

    selectdata: ["all", "draft", "published"]

 -->

<template>
    <div class="dropdown">
        <div class="dropdown-header" @blur="onBlur" @click="openSelect()">
            <div class="prefix">

            </div>

            <div class="content">

            </div>

            <div class="suffix">

            </div>
        </div>
        <transition name="fade">
            <div class="options" v-show="showOptions" @mouseleave="onLeave" :style="`top: ${height + 6}px;`">
                <div>
                    {{ showOptions }}
                </div>
            </div>
        </transition>
    </div>
</template>

<script>
export default {
    name: "SelectMenu",
    data() {
        return {
            showOptions: false
        }
    },
    props: {
        options: {
            type: Array,
            default: () =>{
                return []
            }
        },
        selected: {
            type: [Number, String, Object],
            default: null
        },
        width: {
            type: Number,
            default: 200
        },
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

}
</script>