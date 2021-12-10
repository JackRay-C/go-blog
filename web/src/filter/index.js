
import moment from 'moment';
moment.locale('zh-cn')


const datefmt = function (input) {    
    return new Date(input).toDateString()
}


const momentfmt = function (input, fmtString) {
    return moment(input, fmtString).fromNow()
}



export default {
    datefmt,
    momentfmt
}