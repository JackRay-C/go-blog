import en from './en.js'
import zh from './zh.js'

import Vue from "vue"
import VueI18n from "vue-i18n"
Vue.use(VueI18n)


let messages =  {
    "en": {
        ...en
    },
    "zh": {
        ...zh
    }
}

let locale = localStorage.getItem('lang') || "zh"

export default new VueI18n({
    locale,
    messages
})
