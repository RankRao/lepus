(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[3308],{56621:function(A){A.exports={pre:"pre___34IB6"}},26909:function(A,R,n){"use strict";n.r(R);var Vt=n(57338),Fe=n(25084),Jt=n(66456),be=n(34424),$t=n(57663),l=n(71577),Xt=n(49111),B=n(19650),qt=n(17462),p=n(76772),ea=n(54421),W=n(38272),ta=n(13062),U=n(71230),aa=n(89032),O=n(15746),na=n(58024),S=n(39144),sa=n(43358),T=n(34041),K=n(28991),ua=n(34792),E=n(48086),oa=n(9715),h=n(37657),c=n(28481),je=n(92570),Ie=n(40504),Ae=n(57206),Re=n(34707),Q=n(26001),We=n(80893),Ue=n(91582),i=n(67294),Ke=n(30381),Qe=n.n(Ke),we=n(56621),ze=n.n(we),ke=n(74981),Ne=n(49332),ra=n.n(Ne),He=n(24203),la=n.n(He),Ge=n(82679),ia=n.n(Ge),Ye=n(15500),Ve=n(89899),ca=n.n(Ve),w=n(74295),_a=n.n(w),z=n(93162),da=n.n(z),e=n(85893),Je=[{title:"\u6536\u85CF\u65F6\u95F4",dataIndex:"gmt_created"},{title:"\u6536\u85CF\u5185\u5BB9",dataIndex:"content",copyable:!0,tip:"\u70B9\u51FB\u590D\u5236\u56FE\u6807\u53EF\u4EE5\u590D\u5236\u5B8C\u6574SQL"}],$e=function(){var Xe=h.Z.useForm(),qe=(0,c.Z)(Xe,1),f=qe[0],et=(0,i.useState)({datasource:"",database:"",table:"",sql:""}),k=(0,c.Z)(et,2),ha=k[0],tt=k[1],at=(0,i.useState)([{id:0,cluster_name:""}]),N=(0,c.Z)(at,2),H=N[0],nt=N[1],st=(0,i.useState)([]),G=(0,c.Z)(st,2),Y=G[0],ut=G[1],ot=(0,i.useState)([]),V=(0,c.Z)(ot,2),J=V[0],x=V[1],rt=(0,i.useState)([]),$=(0,c.Z)(rt,2),g=$[0],Z=$[1],lt=(0,i.useState)([]),X=(0,c.Z)(lt,2),it=X[0],q=X[1],ct=(0,i.useState)(!1),ee=(0,c.Z)(ct,2),_t=ee[0],te=ee[1],dt=(0,i.useState)(""),ae=(0,c.Z)(dt,2),o=ae[0],Et=ae[1],ht=(0,i.useState)(""),ne=(0,c.Z)(ht,2),D=ne[0],ft=ne[1],Dt=(0,i.useState)(""),se=(0,c.Z)(Dt,2),d=se[0],L=se[1],mt=(0,i.useState)(""),ue=(0,c.Z)(mt,2),F=ue[0],b=ue[1],Ct=(0,i.useState)(""),oe=(0,c.Z)(Ct,2),C=oe[0],v=oe[1],vt=(0,i.useState)(!1),re=(0,c.Z)(vt,2),yt=re[0],P=re[1],pt=(0,i.useState)(0),le=(0,c.Z)(pt,2),ie=le[0],ce=le[1],St=(0,i.useState)(),_e=(0,c.Z)(St,2),de=_e[0],Ee=_e[1],Pt=(0,i.useState)(),he=(0,c.Z)(Pt,2),fe=he[0],De=he[1],Mt=(0,i.useState)(!1),me=(0,c.Z)(Mt,2),j=me[0],Ce=me[1],Bt=(0,i.useState)(""),ve=(0,c.Z)(Bt,2),M=ve[0],ye=ve[1],Ot=(0,i.useState)(0),pe=(0,c.Z)(Ot,2),Tt=pe[0],Se=pe[1],xt=(0,i.useState)({chineseName:"",username:""}),Pe=(0,c.Z)(xt,2),fa=Pe[0],gt=Pe[1],Zt=(0,i.useState)(""),Me=(0,c.Z)(Zt,2),Da=Me[0],Lt=Me[1],Ft=i.createRef();(0,i.useEffect)(function(){var a=Qe()().format("YYYYMMDD");Lt(a),fetch("/api/v1/currentUser").then(function(t){return t.json()}).then(function(t){gt(t.data)}).catch(function(t){console.log("Fetch current userinfo failed",t)}),fetch("/api/v1/query/datasource_type").then(function(t){return t.json()}).then(function(t){nt(t.data);var s={};t.data.forEach(function(r){s[r.id]=r.name})}).catch(function(t){console.log("Fetch type list failed",t)})},[]);var bt=function(t){x([]),Z([]),L(""),b(""),v(""),f.setFieldsValue({datasource:"",database:"",table:"",sql:""});var s=f.getFieldsValue(),r=s.type;Et(t),fetch("/api/v1/query/datasource?type="+r).then(function(u){return u.json()}).then(function(u){return ut(u.data)}).catch(function(u){console.log("fetch datasource list failed",u)})},jt=function(t){x([]),Z([]),L(""),b(""),v(""),f.setFieldsValue({database:"",table:"",sql:""}),ft(t),fetch("/api/v1/query/database?datasource="+t+"&type="+o).then(function(s){return s.json()}).then(function(s){return x(s.data)}).catch(function(s){console.log("fetch database list failed",s)})},Be=function(t){L(t),v(""),f.setFieldsValue({table:"",sql:""}),fetch("/api/v1/query/table?datasource="+D+"&database="+t+"&type="+o).then(function(s){return s.json()}).then(function(s){return s.data==null?[]:s.data}).then(function(s){return Z(s)}).catch(function(s){console.log("fetch table list failed",s)})},It=function(t){At(t)},At=function(t){b(t);var s="";(o=="MySQL"||o=="TiDB"||o=="Doris"||o=="MariaDB"||o=="GreatSQL"||o=="OceanBase"||o=="ClickHouse"||o=="PostgreSQL")&&(s="select * from "+t+" limit 100"),o=="Oracle"&&(s="select * from "+d+"."+t+" where rownum<=100"),o=="SQLServer"&&(s="select top 100 * from "+t),o=="MongoDB"&&(s="select.from('"+t+"').where('_id','!=','').limit(100)"),v(s),f.setFieldsValue({sql:s})},Rt=function(t,s){var r=s.map(function(u){return{name:u.table_name,value:u.table_name,score:100,meta:""}});console.log(r),t.completers.push({getCompletions:function(m,ge,Ze,Le,I){I(null,r)}})},Wt=function(t){f.setFieldsValue({sql:t}),v(t)},Ut=function(){if(o=="Redis"){E.ZP.warning("Redis\u6570\u636E\u6E90\u4E0D\u652F\u6301\u8BE5\u529F\u80FD");return}if(o==""||d==""||C==""){E.ZP.warning("\u6570\u636E\u6E90/\u6570\u636E\u5E93/SQL\u4E0D\u5B8C\u6574\uFF0C\u65E0\u6CD5\u683C\u5F0F\u5316SQL");return}v((0,Ye.WU)(C))},Kt=function(){if(o==""||D==""||C==""){E.ZP.warning("\u6570\u636E\u6E90/SQL\u4E0D\u5B8C\u6574\uFF0C\u65E0\u6CD5\u6536\u85CFSQL");return}var t=new Headers,s={datasource_type:o,datasource:D,database_name:d,content:C};t.append("Content-Type","application/json"),fetch("/api/v1/favorite/list",{method:"post",headers:t,body:JSON.stringify(s)}).then(function(r){return r.json()}).then(function(r){r.success==!0?E.ZP.success("\u52A0\u5165\u6536\u85CF\u5939\u6210\u529F."):E.ZP.success("\u52A0\u5165\u6536\u85CF\u5939\u5931\u8D25.")}).catch(function(r){console.log("fetch data failed",r)})},Qt=function(){if(o==""||D==""){E.ZP.warning("\u9009\u62E9\u6570\u636E\u6E90\u540E\u624D\u80FD\u6253\u5F00\u6536\u85CF\u5939");return}fetch("/api/v1/favorite/list?datasource="+D+"&datasource_type="+o+"&database="+d).then(function(t){return t.json()}).then(function(t){return q(t.data==null?[]:t.data)}).catch(function(t){console.log("fetch favorite list failed",t)}),te(!0)},Oe=function(){q([]),te(!1)},wt=function(t){console.info(t),P(!0);var s=(0,K.Z)((0,K.Z)({},t),{},{query_type:"execute"}),r=new Headers;r.append("Content-Type","application/json"),fetch("/api/v1/query/doQuery",{method:"post",headers:r,body:JSON.stringify(s)}).then(function(u){return u.json()}).then(function(u){return console.info(u.data),P(!1),Ce(u.success),ye(u.msg),Ee(u.data),De(u.columns),ce(u.total),Se(u.times)}).catch(function(u){console.log("fetch data failed",u)})},Te=function(t){var s={datasource_type:t.type,datasource:t.datasource,database:t.database,table:t.table,sql:t.sql};tt(s),wt(s)},xe=function(t){console.info(t),E.ZP.error("\u6267\u884C\u67E5\u8BE2\u672A\u5B8C\u6210.")},_=function(t){if(t!="doExplain"&&(F==""||F==null)){E.ZP.error("\u8BF7\u5148\u70B9\u51FB\u5DE6\u4FA7\u8868\u540D\u79F0\u9009\u62E9\u8868.");return}P(!0);var s={datasource_type:o,datasource:D,database:d,table:F,sql:C,query_type:t},r=new Headers;r.append("Content-Type","application/json"),fetch("/api/v1/query/doQuery",{method:"post",headers:r,body:JSON.stringify(s)}).then(function(u){return u.json()}).then(function(u){return P(!1),Ce(u.success),ye(u.msg),Ee(u.data),De(u.columns),ce(u.total),Se(u.times)}).catch(function(u){console.log("fetch data failed",u),E.ZP.error("\u6267\u884C\u67E5\u8BE2\u5931\u8D25")})},zt=function(t){return t.map(function(s){var r={header:s.title,key:s.dataIndex,width:s.width/5||20};return r})},kt=function(t,s){t.xlsx.writeBuffer().then(function(r){var u=new Blob([r],{type:""});(0,z.saveAs)(u,s)})},Nt=function(){var t=new w.Workbook,s=t.addWorksheet("\u6570\u636E\u7ED3\u679C");s.properties.defaultRowHeight=20,s.columns=zt(fe);var r=s.addRows(de);r==null||r.forEach(function(y){y.font={size:11,name:"\u5B8B\u4F53"},y.alignment={vertical:"middle",horizontal:"left",wrapText:!1}});var u=s.getRow(1);u.eachCell(function(y,ma){y.fill={type:"pattern",pattern:"solid",fgColor:{argb:"0099CC"}},y.font={bold:!0,italic:!1,size:11,name:"\u5B8B\u4F53",color:{argb:"FFFFFF"}},y.alignment={vertical:"middle",horizontal:"center",wrapText:!1}});var m=new Date,ge=m.getFullYear().toString(),Ze=(m.getMonth()+1).toString(),Le=m.getDate().toString(),I=m.getHours().toString(),Ht=m.getMinutes().toString(),Gt=m.getSeconds().toString(),Yt=o+"-"+ge+Ze+Le+I+Ht+Gt+".xlsx";kt(t,Yt)};return(0,e.jsxs)(We.ZP,{children:[(0,e.jsx)(U.Z,{style:{marginTop:"10px"},children:(0,e.jsx)(O.Z,{span:24,children:(0,e.jsx)(S.Z,{children:(0,e.jsxs)(h.Z,{style:{marginTop:0},form:f,onFinish:Te,onFinishFailed:xe,initialValues:{},name:"sqlForm",layout:"inline",children:[(0,e.jsx)(h.Z.Item,{name:"type",label:"\u6570\u636E\u6E90\u7C7B\u578B",rules:[{required:!0,message:"\u8BF7\u9009\u62E9\u6570\u636E\u6E90\u7C7B\u578B"}],children:(0,e.jsx)(T.Z,{showSearch:!0,style:{width:240},placeholder:"\u8BF7\u9009\u62E9\u6570\u636E\u6E90\u7C7B\u578B",onChange:function(t){bt(t)},children:H&&H.map(function(a){return(0,e.jsx)(Option,{value:a.name,children:a.name},a.name)})})}),(0,e.jsx)(h.Z.Item,{name:"datasource",label:"\u6570\u636E\u6E90",rules:[{required:!0,message:"\u8BF7\u9009\u62E9\u6570\u636E\u6E90"}],children:(0,e.jsx)(T.Z,{showSearch:!0,style:{width:320},placeholder:"\u8BF7\u9009\u62E9\u6570\u636E\u6E90",value:D,onChange:function(t){jt(t)},children:Y&&Y.map(function(a){return(0,e.jsxs)(Option,{value:a.host+":"+a.port,children:[a.name,"[",a.status==1?"\u53EF\u7528":"\u4E0D\u53EF\u7528","] "]},a.host+":"+a.port)})})}),o!="Redis"&&(0,e.jsx)(h.Z.Item,{name:"database",label:"\u6570\u636E\u5E93",rules:[{required:!0,message:"\u8BF7\u9009\u62E9\u6570\u636E\u5E93"}],children:(0,e.jsx)(T.Z,{showSearch:!0,style:{width:240},placeholder:"\u8BF7\u9009\u62E9\u6570\u636E\u5E93",value:d,onChange:function(t){Be(t)},children:J&&J.map(function(a){return(0,e.jsx)(Option,{value:a.database_name,children:a.database_name},a.database_name)})})})]})})})}),(0,e.jsxs)(U.Z,{children:[o!="Redis"&&(0,e.jsx)(O.Z,{span:4,children:(0,e.jsx)(S.Z,{size:"small",title:"\u6570\u636E\u8868",extra:(0,e.jsx)("a",{href:"javascript:void(0)",onClick:function(t){return Be(d)},children:"\u5237\u65B0"}),style:{width:"100%",height:"750px",overflow:"auto"},children:(0,e.jsx)(W.ZP,{size:"small",dataSource:g,renderItem:g!=null&&function(a){return(0,e.jsx)(W.ZP.Item,{children:(0,e.jsxs)("a",{href:"javascript:void(0)",onClick:function(s){return It(a.table_name)},children:[(0,e.jsx)(je.Z,{})," ",a.table_name]})})}})})}),(0,e.jsxs)(O.Z,{span:20,children:[(0,e.jsxs)(S.Z,{children:[d&&d.length>0&&(0,e.jsx)(p.Z,{message:"\u5F53\u524D\u67E5\u8BE2\u5F15\u64CE\uFF1A"+o+"\uFF0C\u6570\u636E\u5E93: "+d,type:"info",showIcon:!0,closable:!0}),o=="Redis"&&(0,e.jsx)(B.Z,{direction:"vertical",children:(0,e.jsx)(p.Z,{message:"\u8BF7\u9009\u62E9\u67E5\u8BE2\u6570\u636E\u6E90\uFF0C\u518D\u8F93\u5165\u547D\u4EE4\uFF0C\u5F53\u524D\u652F\u6301\u7684\u547D\u4EE4\u6709\uFF1ARANDOMKEY\u3001EXISTS\u3001TYPE\u3001TTL\u3001GET\u3001HLEN\u3001HKEYS\u3001HGET\u3001HGETALL\u3001LLEN\u3001LINDEX\u3001LRANGE\u3001SCARD\u3001SMEMBERS\u3001SISMEMBER\u3001ZCARD\u3001ZCOUNT\u3001ZRANGE",type:"info",showIcon:!0,closable:!0})}),(0,e.jsxs)(h.Z,{style:{marginTop:8},form:f,onFinish:Te,onFinishFailed:xe,initialValues:{},name:"sqlForm",layout:"horizontal",children:[(0,e.jsxs)(h.Z.Item,{name:"sql",rules:[{required:!0,message:"\u8BF7\u8F93\u5165SQL\u67E5\u8BE2\u547D\u4EE4"}],children:[(0,e.jsx)(ke.ZP,{ref:Ft,placeholder:"\u8BF7\u8F93\u5165\u6267\u884C\u7684SQL\u547D\u4EE4",mode:"mysql",theme:"textmate",name:"blah2",fontSize:14,showPrintMargin:!0,showGutter:!0,highlightActiveLine:!0,style:{width:"100%",height:"200px",border:"1px solid #ccc"},value:C,editorProps:{$blockScrolling:!1},onChange:function(t){return Wt(t)},onLoad:function(t){return Rt(t,g)},setOptions:{useWorker:!1,enableBasicAutocompletion:!0,enableLiveAutocompletion:!0,enableSnippets:!0,showLineNumbers:!0,tabSize:1}}),(0,e.jsx)(l.Z,{htmlType:"button",type:"dashed",icon:(0,e.jsx)(Ie.Z,{}),size:"small",onClick:function(){return Ut()},children:"\u683C\u5F0F\u5316SQL\u8BED\u53E5"}),(0,e.jsx)(l.Z,{htmlType:"button",type:"dashed",icon:(0,e.jsx)(Ae.Z,{}),size:"small",onClick:function(){return Kt()},children:"\u52A0\u5165\u6536\u85CF\u5939"}),(0,e.jsx)(l.Z,{htmlType:"button",type:"dashed",icon:(0,e.jsx)(Re.Z,{}),size:"small",onClick:function(){return Qt()},children:"\u6253\u5F00\u6536\u85CF\u5939"})]}),(0,e.jsx)(h.Z.Item,{wrapperCol:{offset:0,span:16},children:(0,e.jsxs)(B.Z,{children:[(0,e.jsx)(l.Z,{type:"primary",htmlType:"submit",icon:(0,e.jsx)(Q.Z,{}),children:"\u6267\u884C\u8BED\u53E5"}),(o=="MySQL"||o=="TiDB"||o=="Doris"||o=="MariaDB"||o=="GreatSQL"||o=="OceanBase")&&(0,e.jsxs)(e.Fragment,{children:[(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("doExplain")},children:"\u67E5\u770B\u6267\u884C\u8BA1\u5212"}),(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showIndex")},children:"\u67E5\u770B\u8868\u7D22\u5F15"}),(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showColumn")},children:"\u67E5\u770B\u8868\u7ED3\u6784"}),(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showCreate")},children:"\u67E5\u770B\u5EFA\u8868\u8BED\u53E5"}),(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showTableSize")},children:"\u67E5\u770B\u8868\u5BB9\u91CF"})]}),o=="Oracle"&&(0,e.jsxs)(e.Fragment,{children:[(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("doExplain")},children:"\u67E5\u770B\u6267\u884C\u8BA1\u5212"}),(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showIndex")},children:"\u67E5\u770B\u8868\u7D22\u5F15"}),(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showColumn")},children:"\u67E5\u770B\u8868\u7ED3\u6784"}),(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showCreate")},children:"\u67E5\u770B\u5EFA\u8868\u8BED\u53E5"}),(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showTableSize")},children:"\u67E5\u770B\u8868\u5BB9\u91CF"})]}),o=="PostgreSQL"&&(0,e.jsxs)(e.Fragment,{children:[(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("doExplain")},children:"\u67E5\u770B\u6267\u884C\u8BA1\u5212"}),(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showIndex")},children:"\u67E5\u770B\u8868\u7D22\u5F15"}),(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showColumn")},children:"\u67E5\u770B\u8868\u7ED3\u6784"}),(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showTableSize")},children:"\u67E5\u770B\u8868\u5BB9\u91CF"})]}),o=="ClickHouse"&&(0,e.jsxs)(e.Fragment,{children:[(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showColumn")},children:"\u67E5\u770B\u8868\u7ED3\u6784"}),(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showCreate")},children:"\u67E5\u770B\u5EFA\u8868\u8BED\u53E5"}),(0,e.jsx)(l.Z,{type:"default",htmlType:"button",onClick:function(){return _("showTableSize")},children:"\u67E5\u770B\u8868\u5BB9\u91CF"})]})]})})]})]}),(0,e.jsxs)(S.Z,{children:[j==!1&&M!=""&&(0,e.jsx)(p.Z,{type:"error",message:"\u6267\u884C\u5931\u8D25\uFF1A"+M,banner:!0}),j==!0&&M!=""&&(0,e.jsx)(p.Z,{type:"success",message:"\u6267\u884C\u6210\u529F\uFF0C\u8017\u65F6\uFF1A"+Tt+"\u6BEB\u79D2,"+M,banner:!0}),j==!0&&ie>=0&&(0,e.jsxs)("div",{style:{whiteSpace:"pre-wrap",marginTop:"10px"},onCopy:function(t){return handleCopy(t)},children:[(0,e.jsxs)("div",{style:{width:"100%",float:"right",marginBottom:"10px"},children:["\u67E5\u8BE2\u5230"+ie+"\u6761\u6570\u636E"," ",(0,e.jsx)(l.Z,{icon:(0,e.jsx)(Q.Z,{}),onClick:Nt,children:"\u67E5\u8BE2\u7ED3\u679C\u5BFC\u51FAExcel"})]}),(0,e.jsx)(be.Z,{bordered:!0,loading:yt,scroll:{scrollToFirstRowOnChange:!0,x:100},className:ze().tableStyle,dataSource:de,columns:fe,size:"small"})]})]}),(0,e.jsx)(p.Z,{type:"info",message:"\u652F\u6301MySQL/MariaDB/GreatSQL/TiDB/Doris/OceanBase/ClickHouse/Oracle/PostgreSQL/SQLServer/MongoDB/Redis\u6570\u636E\u5E93\u7684\u6570\u636E\u67E5\u8BE2/\u6570\u636E\u53D8\u66F4/\u7ED3\u6784\u53D8\u66F4\uFF0C\u529F\u80FD\u548C\u6743\u9650\u8BF7\u53C2\u9605\u5B98\u65B9\u6587\u6863\u3002",banner:!0,closable:!0})]})]}),(0,e.jsx)(Fe.Z,{title:"SQL\u6536\u85CF\u5939",placement:"right",width:800,onClose:Oe,visible:_t,extra:(0,e.jsx)(B.Z,{children:(0,e.jsx)(l.Z,{onClick:Oe,children:"\u5173\u95ED"})}),children:(0,e.jsx)(Ue.Z,{rowKey:"id",search:!1,dataSource:it,columns:Je,size:"middle"})})]})};R.default=$e}}]);