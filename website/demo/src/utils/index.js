/*
 *常用全局函数
 *
 */

/**
 * new Date().Format('yyyy-MM-dd hh:mm:ss')
 */
Date.prototype.Format = function(fmt) { //author: meizz
  var o = {
      "M+": this.getMonth() + 1, //月份
      "d+": this.getDate(), //日
      "h+": this.getHours(), //小时
      "m+": this.getMinutes(), //分
      "s+": this.getSeconds(), //秒
      "q+": Math.floor((this.getMonth() + 3) / 3), //季度
      "S": this.getMilliseconds() //毫秒
  };
  if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
  for (var k in o)
      if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
  return fmt;
}
/**
* 常用的关于日期的函数
*/
var now = new Date(); //当前日期
var nowDayOfWeek = now.getDay(); //今天本周的第几天
var nowDay = now.getDate(); //当前日
var nowMonth = now.getMonth(); //当前月
var nowYear = now.getFullYear(); //当前年
/**
 * 获得本周的开端日期，以周一为起点
 */
function getWeekStartDate() {
  let weekStartDate = new Date(nowYear, nowMonth, nowDay - nowDayOfWeek + 1);
  //周日需要另外处理
  if (nowDayOfWeek === 0) {
    weekStartDate = new Date(nowYear, nowMonth, nowDay - 6);
  }
  return formatDate(weekStartDate);
}
/**
 * 获得本周的停止日期，以周日为终点
 */
function getWeekEndDate() {
  let weekEndDate = new Date(nowYear, nowMonth, nowDay + (7 - nowDayOfWeek)); 
  //周日需要另外处理
  if (nowDayOfWeek === 0) {
    weekEndDate = new Date(nowYear, nowMonth, nowDay);
  }
  return formatDate(weekEndDate);
}
/** 
* 获得指定时间的周的开端日期
* 参数 N 是指当前周往前[以前](+)、往后[未来](-) N 周
* N = 0 是当前周
*/
function getBeforeWeekStartDate(n) {
  let weekStartDate = new Date(nowYear, nowMonth, nowDay - 6 - nowDayOfWeek - (n -1) * 7);
  //周日需要另外处理
  if (nowDayOfWeek === 0) {
    weekStartDate = new Date(nowYear, nowMonth, nowDay- 6 - n * 7);
  }
  return formatDate(weekStartDate);
}
//获得指定时间的周的停止日期
//参数 N 是指当前周往前[以前](+)、往后[未来](-) N 周
// N = 0 是当前周
function getBeforeWeekEndDate(n) {
  let weekEndDate = new Date(nowYear, nowMonth, nowDay - nowDayOfWeek - (n - 1) * 7);
  //周日需要另外处理
  if (nowDayOfWeek === 0) {
    weekEndDate = new Date(nowYear, nowMonth, nowDay - n * 7);
  }
  return formatDate(weekEndDate);
}
/**
* 获取指定日期在当年的第几周
* 传参请一定要传 date 格式的参数
* 调用 getWeek( new Date('2018-8-30'))  返回35 ，即第35周
*/
function getWeekOfYear(dt) {
  var today = new Date(dt);
  var firstDay = new Date(today.getFullYear(), 0, 1);
  var dayOfWeek = firstDay.getDay();
  var spendDay = 1;
  if (dayOfWeek != 0) {
    spendDay = 7 - dayOfWeek + 1;
  }
  firstDay = new Date(today.getFullYear(), 0, 1 + spendDay);
  //需要+1是因为零分零秒会被算成前一天
  var d = Math.ceil((today.valueOf() - firstDay.valueOf() + 1) / 86400000);
  var result = Math.ceil(d / 7);
  return result + 1;
}
//获得本月的开端日期
function getMonthStartDate() {
  let monthStartDate = new Date(nowYear, nowMonth, 1);
  return formatDate(monthStartDate);
}
//获得本月的停止日期
function getMonthEndDate() {
  let monthEndDate = new Date(nowYear, nowMonth, getMonthDays(nowMonth));
  return formatDate(monthEndDate);
}
//获得上月开端时候
function getLastMonthStartDate() {
  let lastMonthStartDate = new Date(new Date().getFullYear(), new Date().getMonth() - 1, 1)
  return formatDate(lastMonthStartDate);
}
//获得上月停止时候
function getLastMonthEndDate() {
  var date = new Date();
  var day = new Date(date.getFullYear(), date.getMonth(), 0).getDate();
  var lastMonthEndDate = new Date(new Date().getFullYear(), new Date().getMonth() - 1, day);
  return formatDate(lastMonthEndDate);
}
//获得某月的天数
function getMonthDays(myMonth){
  let monthStartDate = new Date(nowYear, myMonth, 1);
  let monthEndDate = new Date(nowYear, myMonth + 1, 1);
  let days = (monthEndDate - monthStartDate)/(1000 * 60 * 60 * 24);
  return days;
}
/*
* 获得指定时间是在当年的第几季度
*/
function getQuarterlyOfYear(dt){
  let month = dt.getMonth()
  return parseInt(month / 3) + 1
}

//字段拼接
function substrDataStr(data) {
  let strArr = [];
  strArr.push(data.substr(0, 4))
  strArr.push(data.substr(4, 2))
  strArr.push(data.substr(6, 2))
  return strArr
}
//格局化日期：yyyy-MM-dd
function formatDate(date) {
  let myyear = date.getFullYear();
  let mymonth = date.getMonth() + 1;
  let myweekday = date.getDate();
  if (mymonth < 10) {
    mymonth = "0" + mymonth;
  }
  if (myweekday < 10) {
    myweekday = "0" + myweekday;
  }
  return (myyear + "-" + mymonth + "-" + myweekday);
}
//获取今天之前(-)、之后(+)的多少天的日期
function getDateStr(AddDayCount) {
  let dd = new Date();
  dd.setDate(dd.getDate() + AddDayCount); //获取AddDayCount天后的日期
  let y = dd.getFullYear();
  let m = dd.getMonth() + 1;
  let d = dd.getDate();
  if (m < 10) {
    m = "0" + m;
  }
  if (d < 10) {
    d = "0" + d;
  }
  return y + "-" + m + "-" + d;
}

/**
* 对象的深度拷贝通用方法
* 使用方式：var newObj = objDeepCopy(oldObj)
* 参考https://www.cnblogs.com/jiangzilong/p/6513552.html
*/
var objDeepCopy = function (source) {
  var sourceCopy = source instanceof Array ? [] : {};
  for (var item in source) {
      sourceCopy[item] = typeof source[item] === 'object' ? objDeepCopy(source[item]) : source[item];
  }
  return sourceCopy;
}

/**
 * 检测是否空对象
 */
function isEmptyObject(obj)
{
 　　for(var name in obj)
 　　{
 　　　　if(obj.hasOwnProperty(name))
 　　　　{
　　　　　　return false;//返回false，不为空对象
 　　　　}
 　　}
 　　return true;//返回true，为空对象
}

function formatNumber (n) {
  const str = n.toString()
  return str[1] ? str : `0${str}`
}

export function formatTime (date) {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()

  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()

  const t1 = [year, month, day].map(formatNumber).join('/')
  const t2 = [hour, minute, second].map(formatNumber).join(':')

  return `${t1} ${t2}`
}

export function getCurrentPageUrl() {
  const pages = getCurrentPages()
  const currentPage = pages[pages.length - 1]
  const url = `/${currentPage.route}`
  return url
}

export function getCurrentPageUrlWithArgs() {
  const pages = getCurrentPages()
  const currentPage = pages[pages.length - 1]
  const url = currentPage.route
  const options = currentPage.options
  let urlWithArgs = `/${url}?`
  for (let key in options) {
    const value = options[key]
    urlWithArgs += `${key}=${value}&`
  }
  urlWithArgs = urlWithArgs.substring(0, urlWithArgs.length - 1)
  return urlWithArgs
}

/**
 * 常用输入校验方法
 * 创建者: kevin
 * 创建于 2019年10月19日
 */

/**
 * 限定输入字符串长度不少于N
 */
function minString(str,n) {
  if(str.length >= n) {
      return true
  } else {
      return false
  }
}

/**
 *格式化头像
 */
function formatHeadImg(imgSrc) { //默认头像
  let url = 'https:/m.zxoa.com.cn/agent/web'
  if (!imgSrc) {
    return '/static/images/head.png';
  }
  if (imgSrc.indexOf('.png') != -1) {
      imgSrc = url+imgSrc
  } else if (imgSrc.indexOf('.jpg') != -1) {
      imgSrc = url+imgSrc
  }
  return imgSrc;
}


export default {
  formatNumber,
  formatTime,
  getCurrentPageUrl,
  getCurrentPageUrlWithArgs,
  objDeepCopy,
  isEmptyObject,
  getWeekStartDate,
  getWeekEndDate,
  getWeekOfYear,
  getBeforeWeekStartDate,
  getBeforeWeekEndDate,
  getMonthStartDate,
  getMonthEndDate,
  getLastMonthStartDate,
  getLastMonthEndDate,
  getQuarterlyOfYear,
  substrDataStr,
  formatDate,
  getDateStr,
  minString,
  formatHeadImg
}
