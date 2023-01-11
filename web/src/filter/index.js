
import moment from 'moment';
moment.locale('zh-cn')


const datefmt = function (input) {    
    let n = input.toString().length || 0
    if(!isNaN(parseFloat(input)) && isFinite(input) ) {
        if(n === 13 ){
            return moment.unix(input/1000).format("YYYY年MM月DD日")
        }
        if(n === 10) {
            return moment.unix(input).format("YYYY年MM月DD日")
        }
    }
    return input
}


const momentfmt = function (input, fmtString) {
    let n = input.toString().length || 0
    if(!isNaN(parseFloat(input)) && isFinite(input) ) {
        if(n === 13 ){
            return moment.unix(input/1000).fromNow()
        }
        if(n === 10) {
            return moment.unix(input).fromNow()
        }
    }
    return moment(input, fmtString).fromNow()
}

export default {
    datefmt,
    momentfmt
}