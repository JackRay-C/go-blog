<template>
    <div class="editable-div" :style="{'max-width': maxwidth + 'px', 'min-width': minwidth + 'px'}" :class="{'text-hidden': !ischecked, 'text-show': ischecked}" :contenteditable="input" @focus="ischecked=true" @blur="blurF" @input="inputF" v-text="text" @keydown="keydown">
    </div>
</template>

<style scoped lang="scss">
.editable-div {
    display: inline;
    outline: none;
    padding: 0 5px;
    white-space: nowrap;
    border-bottom: 1px solid transparent;

    &:focus {
        border-bottom: 1px solid #ccc;
        border-color: #333;
    }
}

.text-hidden {
    overflow: hidden;
    text-overflow: ellipsis;
}

</style>

<script>
export default {
    name: "Editable",
    data(){
        return {
            ischecked: false,
            text: this.value
        }
    },
    props:{
        input: {
            type: Boolean,
            default: true
        },
        value: {
            type: String,
            default: ""
        },
        maxwidth: {
            type: Number,
            default: 200
        },
        minwidth: {
            type: Number,
            default: 100
        }
    },
    watch: {
        value(){
            if(!this.ischecked) {
                this.text = this.value
            }
        }
    },
    methods:{
        inputF(e){
            const val = e.target.innerHTML
            this.$emit('input', val)
        },
        blurF(e) {
            this.ischecked = false
            this.text = this.value
            e.view.blur()
        },
        keydown(e){
            if(e.keyCode===13) {
                this.blurF(e)
                e.preventDefault()
            }
        }
    }
}
</script>