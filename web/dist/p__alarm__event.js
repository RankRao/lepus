(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[4097],{40800:function(Q,A,n){"use strict";n.r(A),n.d(A,{default:function(){return W}});var K=n(49111),v=n(19650),M=n(62350),g=n(75443),k=n(66456),E=n(34424),S=n(28991),o=n(55507),p=n(92137),y=n(28481),H=n(34792),B=n(48086),h=n(67294),w=n(47599),L=n(40418),D=n(3140);function U(l){return F.apply(this,arguments)}function F(){return F=(0,p.Z)((0,o.Z)().mark(function l(d){return(0,o.Z)().wrap(function(i){for(;;)switch(i.prev=i.next){case 0:return i.abrupt("return",(0,D.WY)("/api/v1/alarm/event",{params:d}));case 1:case"end":return i.stop()}},l)})),F.apply(this,arguments)}function b(){return x.apply(this,arguments)}function x(){return x=(0,p.Z)((0,o.Z)().mark(function l(){return(0,o.Z)().wrap(function(a){for(;;)switch(a.prev=a.next){case 0:return a.abrupt("return",(0,D.WY)("/api/v1/alarm/event/analysis",{}));case 1:case"end":return a.stop()}},l)})),x.apply(this,arguments)}function J(){return j.apply(this,arguments)}function j(){return j=_asyncToGenerator(_regeneratorRuntime().mark(function l(){return _regeneratorRuntime().wrap(function(a){for(;;)switch(a.prev=a.next){case 0:return a.abrupt("return",request("/api/fake_analysis_chart_data"));case 1:case"end":return a.stop()}},l)})),j.apply(this,arguments)}var m=n(60331),P=n(85029);function $(l){return Z.apply(this,arguments)}function Z(){return Z=(0,p.Z)((0,o.Z)().mark(function l(d){return(0,o.Z)().wrap(function(i){for(;;)switch(i.prev=i.next){case 0:return i.abrupt("return",(0,P.WY)("/api/v1/alarm/batchUpdateStatus",{method:"PUT",headers:{"Content-Type":"application/json"},data:d}));case 1:case"end":return i.stop()}},l)})),Z.apply(this,arguments)}var t=n(85893),O=function(){(0,h.useEffect)(function(){try{console.info("init page.")}catch(c){B.ZP.error("get event error. ".concat(c))}},[]);var d=(0,h.useState)(),a=(0,y.Z)(d,2),i=a[0],X=a[1],N=(0,h.useState)(!1),T=(0,y.Z)(N,2),q=T[0],_=T[1],I=(0,h.useRef)(),V=(0,h.useState)({}),R=(0,y.Z)(V,2),ee=R[0],Y=R[1],z=function(){var c=(0,p.Z)((0,o.Z)().mark(function u(){var e;return(0,o.Z)().wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return r.prev=0,r.next=4,b();case 4:return e=r.sent,Y(e),r.abrupt("return");case 9:return r.prev=9,r.t0=r.catch(0),r.abrupt("return",{success:!1,msg:r.t0});case 12:case"end":return r.stop()}},u,null,[[0,9]])}));return function(){return c.apply(this,arguments)}}(),G=[{title:"\u544A\u8B66\u6807\u9898",dataIndex:"alarm_title",hideInForm:!0,sorter:!1,search:!1,width:280,copyable:!0,ellipsis:!0,tip:"\u6807\u9898\u8FC7\u957F\u4F1A\u81EA\u52A8\u6536\u7F29",render:function(u,e){switch(e.status){case 0:return(0,t.jsxs)("span",{children:[(0,t.jsx)(m.Z,{color:"volcano",children:"\u672A\u5904\u7406"}),"[",e.alarm_level,"] ",e.alarm_title]});case 1:return(0,t.jsxs)("span",{children:[(0,t.jsx)(m.Z,{color:"blue",children:"\u5904\u7406\u4E2D"}),"[",e.alarm_level,"] ",e.alarm_title]});case 2:return(0,t.jsxs)("span",{children:[(0,t.jsx)(m.Z,{color:"green",children:"\u5DF2\u5B8C\u6210"}),"[",e.alarm_level,"] ",e.alarm_title]});default:return(0,t.jsxs)("span",{children:[(0,t.jsx)(m.Z,{children:"\u672A\u77E5"}),"[",e.alarm_level,"] ",e.alarm_title]})}}},{title:"\u4E8B\u4EF6\u7C7B\u578B",dataIndex:"event_type",hideInForm:!0,sorter:!1,width:100},{title:"\u4E8B\u4EF6\u7EC4",dataIndex:"event_group",hideInForm:!1,sorter:!1,width:85},{title:"\u4E8B\u4EF6\u5B9E\u4F53",dataIndex:"event_entity",hideInForm:!1,sorter:!1,width:160,ellipsis:!0},{title:"\u4E8B\u4EF6\u6807\u7B7E",dataIndex:"event_tag",hideInForm:!1,sorter:!1,width:130},{title:"\u89E6\u53D1\u89C4\u5219",dataIndex:"event_key",hideInForm:!1,sorter:!1,width:140,ellipsis:!0,render:function(u,e){var f=(0,t.jsxs)("span",{children:[e.event_key," [",e.event_value,e.event_unit,e.alarm_rule,e.alarm_value,e.event_unit,"]"]});return f}},{title:"\u90AE\u4EF6",dataIndex:"send_mail",filters:!1,onFilter:!1,valueEnum:{"0":{text:"",status:"Default"},"1":{text:"",status:"Success"},"2":{text:"",status:"Error"}},sorter:!1,search:!1,width:55},{title:"\u544A\u8B66\u65F6\u95F4",dataIndex:"gmt_created",valueType:"dateTime",hideInForm:!0,search:!1,width:180}],C=function(u,e){console.log(u,e),$({ids:u,status:e}).then(function(f){f.success&&I.current.reload()})};return(0,t.jsx)(w.ZP,{children:(0,t.jsx)(L.ZP,{headerTitle:"\u6570\u636E\u5217\u8868",cardBordered:!0,actionRef:I,rowKey:"id",search:{labelWidth:120},request:function(u,e){return U((0,S.Z)((0,S.Z)({},u),{},{sorter:e}))},columns:G,pagination:{pageSize:50},rowSelection:{selections:[E.Z.SELECTION_ALL,E.Z.SELECTION_INVERT],renderCell:function(u,e,f,r){return e.status!==2?r:!1}},tableAlertRender:function(u){var e=u.selectedRows,f=u.onCleanSelected;return(0,t.jsxs)(v.Z,{size:12,children:[(0,t.jsx)("span",{children:e&&(0,t.jsxs)(t.Fragment,{children:["\u5DF2\u9009 ",e.filter(function(r){return r&&(r.status==0||r.status==1)}).length," \u9879 ",(0,t.jsx)("a",{style:{marginLeft:8},onClick:f,children:"\u53D6\u6D88\u9009\u62E9"})]})}),(0,t.jsx)("span",{children:(0,t.jsx)(v.Z,{children:e&&e.filter(function(r){return r&&r.status==0}).length>0&&(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)("strong",{children:e&&"\u672A\u5904\u7406  ".concat(e.filter(function(r){return r&&r.status==0}).length," \u9879")}),(0,t.jsx)("span",{children:"\u64CD\u4F5C\uFF1A"}),(0,t.jsx)(g.Z,{title:"\u662F\u5426\u6807\u8BB0\u4E3A\u5904\u7406\u4E2D \uFF1F",onConfirm:function(){return C(e.filter(function(s){return s&&s.status==0}).map(function(s){return s.id}),1)},okText:"\u662F",cancelText:"\u5426",children:(0,t.jsx)("a",{children:"\u5904\u7406\u4E2D"})}),(0,t.jsx)(g.Z,{title:"\u662F\u5426\u6807\u8BB0\u4E3A\u5DF2\u5B8C\u6210 \uFF1F",onConfirm:function(){return C(e.filter(function(s){return s&&s.status==1}).map(function(s){return s.id}),2)},okText:"\u662F",cancelText:"\u5426",children:(0,t.jsx)("a",{children:"\u5DF2\u5B8C\u6210"})})]})})}),(0,t.jsx)("span",{children:(0,t.jsx)(v.Z,{children:e&&e.filter(function(r){return r&&r.status==1}).length>0&&(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)("strong",{children:e&&"\u5904\u7406\u4E2D ".concat(e.filter(function(r){return r&&r.status==1}).length," \u9879")}),(0,t.jsx)("span",{children:"\u64CD\u4F5C\uFF1A"}),(0,t.jsx)(g.Z,{title:"\u662F\u5426\u6807\u8BB0\u4E3A\u5DF2\u5B8C\u6210 \uFF1F",onConfirm:function(){return C(e.filter(function(s){return s&&s.status==1}).map(function(s){return s.id}),2)},okText:"\u662F",cancelText:"\u5426",children:(0,t.jsx)("a",{children:"\u5DF2\u5B8C\u6210"})})]})})})]})},sticky:!0,tableAlertOptionRender:function(){return""}})})},W=O}}]);
