(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d22bd5c"],{f150:function(e,o,t){"use strict";t.r(o);var n=function(){var e=this,o=e.$createElement,t=e._self._c||o;return t("div",{staticClass:"app-container"},[t("el-form",{ref:"form",attrs:{model:e.form,"label-position":e.labelPosition}},[t("el-form-item",{attrs:{label:"企业微信通知开关"}},[t("el-switch",{model:{value:e.form.qywx_is_open,callback:function(o){e.$set(e.form,"qywx_is_open",o)},expression:"form.qywx_is_open"}})],1),t("el-form-item",{attrs:{label:"企业微信通知 webhook"}},[t("el-input",{attrs:{size:"medium"},model:{value:e.form.qywx_webhook,callback:function(o){e.$set(e.form,"qywx_webhook",o)},expression:"form.qywx_webhook"}})],1),t("el-form-item",{attrs:{label:"钉钉通知开关"}},[t("el-switch",{model:{value:e.form.dd_is_open,callback:function(o){e.$set(e.form,"dd_is_open",o)},expression:"form.dd_is_open"}})],1),t("el-form-item",{attrs:{label:"钉钉通知 webhook"}},[t("el-input",{attrs:{size:"medium"},model:{value:e.form.dd_webhook,callback:function(o){e.$set(e.form,"dd_webhook",o)},expression:"form.dd_webhook"}})],1),t("el-form-item",[t("el-button",{attrs:{type:"primary"},on:{click:e.onSubmit}},[e._v("保存")])],1)],1)],1)},r=[],a=t("b775");function i(e){return Object(a["a"])({url:"/notify/Update",method:"post",data:e})}function s(e){return Object(a["a"])({url:"/notify/Get",method:"post",data:e})}var l={data:function(){return{labelPosition:"left",form:{id:"",qywx_is_open:!1,qywx_webhook:"",dd_is_open:!1,dd_webhook:""}}},created:function(){this.fetchData()},methods:{fetchData:function(){var e=this;s().then((function(o){33==o.code?e.form=o.data:e.$message.error(o.data)}))},onSubmit:function(){var e=this;console.log(this.form),i(this.form).then((function(o){33==o.code?e.$message("submit!"):e.$message.error(o.data)}))}}},m=l,f=t("2877"),c=Object(f["a"])(m,n,r,!1,null,null,null);o["default"]=c.exports}}]);