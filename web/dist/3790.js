(self.webpackChunkant_design_pro=self.webpackChunkant_design_pro||[]).push([[3790],{47828:function(){},61859:function(de,I,i){"use strict";i.d(I,{Z:function(){return Jt}});var m=i(22122),S=i(96156),r=i(67294),me=i(94184),te=i.n(me),A=i(42550),h=i(65632),H=function(e,n){var a={};for(var t in e)Object.prototype.hasOwnProperty.call(e,t)&&n.indexOf(t)<0&&(a[t]=e[t]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var l=0,t=Object.getOwnPropertySymbols(e);l<t.length;l++)n.indexOf(t[l])<0&&Object.prototype.propertyIsEnumerable.call(e,t[l])&&(a[t[l]]=e[t[l]]);return a},ge=function(n,a){var t=n.prefixCls,l=n.component,u=l===void 0?"article":l,v=n.className,d=n["aria-label"],T=n.setContentRef,P=n.children,x=H(n,["prefixCls","component","className","aria-label","setContentRef","children"]),N=r.useContext(h.E_),D=N.getPrefixCls,U=N.direction,C=a;T&&(C=(0,A.sQ)(a,T));var j=u,$=D("typography",t),g=te()($,(0,S.Z)({},"".concat($,"-rtl"),U==="rtl"),v);return r.createElement(j,(0,m.Z)({className:g,"aria-label":d,ref:C},x),P)},fe=r.forwardRef(ge);fe.displayName="Typography";var ae=fe,F=ae,f=i(90484),le=i(98423),b=i(28481),K=i(76632),oe=i(50344),ot=i(20640),it=i.n(ot),Ee=i(28991),st={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M257.7 752c2 0 4-.2 6-.5L431.9 722c2-.4 3.9-1.3 5.3-2.8l423.9-423.9a9.96 9.96 0 000-14.1L694.9 114.9c-1.9-1.9-4.4-2.9-7.1-2.9s-5.2 1-7.1 2.9L256.8 538.8c-1.5 1.5-2.4 3.3-2.8 5.3l-29.5 168.2a33.5 33.5 0 009.4 29.8c6.6 6.4 14.9 9.9 23.8 9.9zm67.4-174.4L687.8 215l73.3 73.3-362.7 362.6-88.9 15.7 15.6-89zM880 836H144c-17.7 0-32 14.3-32 32v36c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-36c0-17.7-14.3-32-32-32z"}}]},name:"edit",theme:"outlined"},ct=st,Me=i(27029),Ie=function(n,a){return r.createElement(Me.Z,(0,Ee.Z)((0,Ee.Z)({},n),{},{ref:a,icon:ct}))};Ie.displayName="EditOutlined";var ut=r.forwardRef(Ie),dt=i(79508),ft=i(99165),vt=i(48717),Ce=i(8410),pt=i(42051),Ne=i(34952),De=i(79370),Te=i(61580),$e=i(15105),yt={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M864 170h-60c-4.4 0-8 3.6-8 8v518H310v-73c0-6.7-7.8-10.5-13-6.3l-141.9 112a8 8 0 000 12.6l141.9 112c5.3 4.2 13 .4 13-6.3v-75h498c35.3 0 64-28.7 64-64V178c0-4.4-3.6-8-8-8z"}}]},name:"enter",theme:"outlined"},mt=yt,ke=function(n,a){return r.createElement(Me.Z,(0,Ee.Z)((0,Ee.Z)({},n),{},{ref:a,icon:mt}))};ke.displayName="EnterOutlined";var gt=r.forwardRef(ke),Et=i(94418),Ct=i(96159),ht=function(n){var a=n.prefixCls,t=n["aria-label"],l=n.className,u=n.style,v=n.direction,d=n.maxLength,T=n.autoSize,P=T===void 0?!0:T,x=n.value,N=n.onSave,D=n.onCancel,U=n.onEnd,C=n.component,j=n.enterIcon,$=j===void 0?r.createElement(gt,null):j,g=r.useRef(),w=r.useRef(!1),W=r.useRef(),O=r.useState(x),ne=(0,b.Z)(O,2),J=ne[0],re=ne[1];r.useEffect(function(){re(x)},[x]),r.useEffect(function(){if(g.current&&g.current.resizableTextArea){var y=g.current.resizableTextArea.textArea;y.focus();var E=y.value.length;y.setSelectionRange(E,E)}},[]);var X=function(E){var Z=E.target;re(Z.value.replace(/[\n\r]/g,""))},p=function(){w.current=!0},z=function(){w.current=!1},B=function(E){var Z=E.keyCode;w.current||(W.current=Z)},Y=function(){N(J.trim())},ie=function(E){var Z=E.keyCode,pe=E.ctrlKey,ce=E.altKey,G=E.metaKey,k=E.shiftKey;W.current===Z&&!w.current&&!pe&&!ce&&!G&&!k&&(Z===$e.Z.ENTER?(Y(),U==null||U()):Z===$e.Z.ESC&&D())},_=function(){Y()},se=C?"".concat(a,"-").concat(C):"",R=te()(a,"".concat(a,"-edit-content"),(0,S.Z)({},"".concat(a,"-rtl"),v==="rtl"),l,se);return r.createElement("div",{className:R,style:u},r.createElement(Et.Z,{ref:g,maxLength:d,value:J,onChange:X,onKeyDown:B,onKeyUp:ie,onCompositionStart:p,onCompositionEnd:z,onBlur:_,"aria-label":t,rows:1,autoSize:P}),$!==null?(0,Ct.Tm)($,{className:"".concat(a,"-edit-content-confirm")}):null)},bt=ht;function Pe(e,n){return r.useMemo(function(){var a=!!e;return[a,(0,m.Z)((0,m.Z)({},n),a&&(0,f.Z)(e)==="object"?e:null)]},[e])}var xt=function(e,n){var a=r.useRef(!1);r.useEffect(function(){a.current?e():a.current=!0},n)};function Ae(e){var n=(0,f.Z)(e);return n==="string"||n==="number"}function St(e){var n=0;return e.forEach(function(a){Ae(a)?n+=String(a).length:n+=1}),n}function Ke(e,n){for(var a=0,t=[],l=0;l<e.length;l+=1){if(a===n)return t;var u=e[l],v=Ae(u),d=v?String(u).length:1,T=a+d;if(T>n){var P=n-a;return t.push(String(u).slice(0,P)),t}t.push(u),a=T}return e}var Ot=0,he=1,je=2,we=3,ze=4,Rt=function(n){var a=n.enabledMeasure,t=n.children,l=n.text,u=n.width,v=n.rows,d=n.onEllipsis,T=r.useState([0,0,0]),P=(0,b.Z)(T,2),x=P[0],N=P[1],D=r.useState(Ot),U=(0,b.Z)(D,2),C=U[0],j=U[1],$=(0,b.Z)(x,3),g=$[0],w=$[1],W=$[2],O=r.useState(0),ne=(0,b.Z)(O,2),J=ne[0],re=ne[1],X=r.useRef(null),p=r.useRef(null),z=r.useMemo(function(){return(0,oe.Z)(l)},[l]),B=r.useMemo(function(){return St(z)},[z]),Y=r.useMemo(function(){return!a||C!==we?t(z,!1):t(Ke(z,w),w<B)},[a,C,t,z,w,B]);(0,Ce.Z)(function(){a&&u&&B&&(j(he),N([0,Math.ceil(B/2),B]))},[a,u,l,B,v]),(0,Ce.Z)(function(){var R;C===he&&re(((R=X.current)===null||R===void 0?void 0:R.offsetHeight)||0)},[C]),(0,Ce.Z)(function(){var R,y;if(J){if(C===he){var E=((R=p.current)===null||R===void 0?void 0:R.offsetHeight)||0,Z=v*J;E<=Z?(j(ze),d(!1)):j(je)}else if(C===je)if(g!==W){var pe=((y=p.current)===null||y===void 0?void 0:y.offsetHeight)||0,ce=v*J,G=g,k=W;g===W-1?k=g:pe<=ce?G=w:k=w;var Le=Math.ceil((G+k)/2);N([G,Le,k])}else j(we),d(!0)}},[C,g,W,v,J]);var ie={width:u,whiteSpace:"normal",margin:0,padding:0},_=function(y,E,Z){return r.createElement("span",{"aria-hidden":!0,ref:E,style:(0,m.Z)({position:"fixed",display:"block",left:0,top:0,zIndex:-9999,visibility:"hidden",pointerEvents:"none"},Z)},y)},se=function(y,E){var Z=Ke(z,y);return _(t(Z,!0),E,ie)};return r.createElement(r.Fragment,null,Y,a&&C!==we&&C!==ze&&r.createElement(r.Fragment,null,_("lg",X,{wordBreak:"keep-all",whiteSpace:"nowrap"}),C===he?_(t(z,!1),p,ie):se(w,p)))},Zt=Rt,Tt=function(n){var a=n.title,t=n.enabledEllipsis,l=n.isEllipsis,u=n.children;return!a||!t?u:r.createElement(Te.Z,{title:a,visible:l?void 0:!1},u)},Pt=Tt,wt=function(e,n){var a={};for(var t in e)Object.prototype.hasOwnProperty.call(e,t)&&n.indexOf(t)<0&&(a[t]=e[t]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var l=0,t=Object.getOwnPropertySymbols(e);l<t.length;l++)n.indexOf(t[l])<0&&Object.prototype.propertyIsEnumerable.call(e,t[l])&&(a[t[l]]=e[t[l]]);return a};function Lt(e,n){var a=e.mark,t=e.code,l=e.underline,u=e.delete,v=e.strong,d=e.keyboard,T=e.italic,P=n;function x(N,D){!N||(P=r.createElement(D,{},P))}return x(v,"strong"),x(l,"u"),x(u,"del"),x(t,"code"),x(a,"mark"),x(d,"kbd"),x(T,"i"),P}function be(e,n,a){return e===!0||e===void 0?n:e||a&&n}function Be(e){return Array.isArray(e)?e:[e]}var Mt="...",It=r.forwardRef(function(e,n){var a=e.prefixCls,t=e.className,l=e.style,u=e.type,v=e.disabled,d=e.children,T=e.ellipsis,P=e.editable,x=e.copyable,N=e.component,D=e.title,U=wt(e,["prefixCls","className","style","type","disabled","children","ellipsis","editable","copyable","component","title"]),C=r.useContext(h.E_),j=C.getPrefixCls,$=C.direction,g=(0,pt.E)("Text")[0],w=r.useRef(null),W=r.useRef(null),O=j("typography",a),ne=(0,le.Z)(U,["mark","code","delete","underline","strong","keyboard","italic"]),J=Pe(P),re=(0,b.Z)(J,2),X=re[0],p=re[1],z=(0,K.Z)(!1,{value:p.editing}),B=(0,b.Z)(z,2),Y=B[0],ie=B[1],_=p.triggerType,se=_===void 0?["icon"]:_,R=function(o){var s;o&&((s=p.onStart)===null||s===void 0||s.call(p)),ie(o)};xt(function(){var c;Y||(c=W.current)===null||c===void 0||c.focus()},[Y]);var y=function(o){o==null||o.preventDefault(),R(!0)},E=function(o){var s;(s=p.onChange)===null||s===void 0||s.call(p,o),R(!1)},Z=function(){var o;(o=p.onCancel)===null||o===void 0||o.call(p),R(!1)},pe=Pe(x),ce=(0,b.Z)(pe,2),G=ce[0],k=ce[1],Le=r.useState(!1),Ue=(0,b.Z)(Le,2),Se=Ue[0],We=Ue[1],He=r.useRef(),Fe=function(){clearTimeout(He.current)},Gt=function(o){var s;o==null||o.preventDefault(),o==null||o.stopPropagation(),it()(k.text||String(d)||""),We(!0),Fe(),He.current=setTimeout(function(){We(!1)},3e3),(s=k.onCopy)===null||s===void 0||s.call(k,o)};r.useEffect(function(){return Fe},[]);var Vt=r.useState(!1),Je=(0,b.Z)(Vt,2),Ge=Je[0],Qt=Je[1],Xt=r.useState(!1),Ve=(0,b.Z)(Xt,2),Qe=Ve[0],Yt=Ve[1],_t=r.useState(!1),Xe=(0,b.Z)(_t,2),qt=Xe[0],en=Xe[1],tn=r.useState(!1),Ye=(0,b.Z)(tn,2),_e=Ye[0],nn=Ye[1],rn=r.useState(!1),qe=(0,b.Z)(rn,2),et=qe[0],an=qe[1],ln=Pe(T,{expandable:!1}),tt=(0,b.Z)(ln,2),q=tt[0],L=tt[1],V=q&&!qt,nt=L.rows,ue=nt===void 0?1:nt,Oe=r.useMemo(function(){return!V||L.suffix!==void 0||L.onEllipsis||L.expandable||X||G},[V,L,X,G]);(0,Ce.Z)(function(){q&&!Oe&&(Qt((0,De.G)("webkitLineClamp")),Yt((0,De.G)("textOverflow")))},[Oe,q]);var Q=r.useMemo(function(){return Oe?!1:ue===1?Qe:Ge},[Oe,Qe,Ge]),rt=V&&(Q?et:_e),on=V&&ue===1&&Q,Re=V&&ue>1&&Q,sn=function(o){var s;en(!0),(s=L.onExpand)===null||s===void 0||s.call(L,o)},cn=r.useState(0),at=(0,b.Z)(cn,2),un=at[0],dn=at[1],fn=function(o){var s=o.offsetWidth;dn(s)},vn=function(o){var s;nn(o),_e!==o&&((s=L.onEllipsis)===null||s===void 0||s.call(L,o))};r.useEffect(function(){var c=w.current;if(q&&Q&&c){var o=Re?c.offsetHeight<c.scrollHeight:c.offsetWidth<c.scrollWidth;et!==o&&an(o)}},[q,Q,d,Re]);var Ze=L.tooltip===!0?d:L.tooltip,lt=r.useMemo(function(){var c=function(s){return["string","number"].includes((0,f.Z)(s))};if(!(!q||Q)){if(c(d))return d;if(c(D))return D;if(c(Ze))return Ze}},[q,Q,D,Ze,rt]);if(Y)return r.createElement(bt,{value:typeof d=="string"?d:"",onSave:E,onCancel:Z,onEnd:p.onEnd,prefixCls:O,className:t,style:l,direction:$,component:N,maxLength:p.maxLength,autoSize:p.autoSize,enterIcon:p.enterIcon});var pn=function(){var o=L.expandable,s=L.symbol;if(!o)return null;var M;return s?M=s:M=g.expand,r.createElement("a",{key:"expand",className:"".concat(O,"-expand"),onClick:sn,"aria-label":g.expand},M)},yn=function(){if(!!X){var o=p.icon,s=p.tooltip,M=(0,oe.Z)(s)[0]||g.edit,ee=typeof M=="string"?M:"";return se.includes("icon")?r.createElement(Te.Z,{key:"edit",title:s===!1?"":M},r.createElement(Ne.Z,{ref:W,className:"".concat(O,"-edit"),onClick:y,"aria-label":ee},o||r.createElement(ut,{role:"button"}))):null}},mn=function(){if(!!G){var o=k.tooltips,s=k.icon,M=Be(o),ee=Be(s),ye=Se?be(M[1],g.copied):be(M[0],g.copy),Cn=Se?g.copied:g.copy,hn=typeof ye=="string"?ye:Cn;return r.createElement(Te.Z,{key:"copy",title:ye},r.createElement(Ne.Z,{className:te()("".concat(O,"-copy"),Se&&"".concat(O,"-copy-success")),onClick:Gt,"aria-label":hn},Se?be(ee[1],r.createElement(dt.Z,null),!0):be(ee[0],r.createElement(ft.Z,null),!0)))}},gn=function(o){return[o&&pn(),yn(),mn()]},En=function(o){return[o&&r.createElement("span",{"aria-hidden":!0,key:"ellipsis"},Mt),L.suffix,gn(o)]};return r.createElement(vt.Z,{onResize:fn,disabled:!V||Q},function(c){var o;return r.createElement(Pt,{title:Ze,enabledEllipsis:V,isEllipsis:rt},r.createElement(F,(0,m.Z)({className:te()((o={},(0,S.Z)(o,"".concat(O,"-").concat(u),u),(0,S.Z)(o,"".concat(O,"-disabled"),v),(0,S.Z)(o,"".concat(O,"-ellipsis"),q),(0,S.Z)(o,"".concat(O,"-single-line"),V&&ue===1),(0,S.Z)(o,"".concat(O,"-ellipsis-single-line"),on),(0,S.Z)(o,"".concat(O,"-ellipsis-multiple-line"),Re),o),t),style:(0,m.Z)((0,m.Z)({},l),{WebkitLineClamp:Re?ue:void 0}),component:N,ref:(0,A.sQ)(c,w,n),direction:$,onClick:se.includes("text")?y:null,"aria-label":lt,title:D},ne),r.createElement(Zt,{enabledMeasure:V&&!Q,text:d,rows:ue,width:un,onEllipsis:vn},function(s,M){var ee=s;s.length&&M&&lt&&(ee=r.createElement("span",{key:"show-content","aria-hidden":!0},ee));var ye=Lt(e,r.createElement(r.Fragment,null,ee,En(M)));return ye})))})}),xe=It,Nt=function(e,n){var a={};for(var t in e)Object.prototype.hasOwnProperty.call(e,t)&&n.indexOf(t)<0&&(a[t]=e[t]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var l=0,t=Object.getOwnPropertySymbols(e);l<t.length;l++)n.indexOf(t[l])<0&&Object.prototype.propertyIsEnumerable.call(e,t[l])&&(a[t[l]]=e[t[l]]);return a},Dt=function(n,a){var t=n.ellipsis,l=Nt(n,["ellipsis"]),u=r.useMemo(function(){return t&&(0,f.Z)(t)==="object"?(0,le.Z)(t,["expandable","rows"]):t},[t]);return r.createElement(xe,(0,m.Z)({ref:a},l,{ellipsis:u,component:"span"}))},$t=r.forwardRef(Dt),kt=function(e,n){var a={};for(var t in e)Object.prototype.hasOwnProperty.call(e,t)&&n.indexOf(t)<0&&(a[t]=e[t]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var l=0,t=Object.getOwnPropertySymbols(e);l<t.length;l++)n.indexOf(t[l])<0&&Object.prototype.propertyIsEnumerable.call(e,t[l])&&(a[t[l]]=e[t[l]]);return a},At=function(n,a){var t=n.ellipsis,l=n.rel,u=kt(n,["ellipsis","rel"]),v=r.useRef(null);r.useImperativeHandle(a,function(){return v.current});var d=(0,m.Z)((0,m.Z)({},u),{rel:l===void 0&&u.target==="_blank"?"noopener noreferrer":l});return delete d.navigate,r.createElement(xe,(0,m.Z)({},d,{ref:v,ellipsis:!!t,component:"a"}))},Kt=r.forwardRef(At),jt=i(93355),zt=function(e,n){var a={};for(var t in e)Object.prototype.hasOwnProperty.call(e,t)&&n.indexOf(t)<0&&(a[t]=e[t]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var l=0,t=Object.getOwnPropertySymbols(e);l<t.length;l++)n.indexOf(t[l])<0&&Object.prototype.propertyIsEnumerable.call(e,t[l])&&(a[t[l]]=e[t[l]]);return a},Bt=(0,jt.a)(1,2,3,4,5),Ut=function(n,a){var t=n.level,l=t===void 0?1:t,u=zt(n,["level"]),v;return Bt.indexOf(l)!==-1?v="h".concat(l):v="h1",r.createElement(xe,(0,m.Z)({ref:a},u,{component:v}))},Wt=r.forwardRef(Ut),Ht=function(n,a){return r.createElement(xe,(0,m.Z)({ref:a},n,{component:"div"}))},Ft=r.forwardRef(Ht),ve=F;ve.Text=$t,ve.Link=Kt,ve.Title=Wt,ve.Paragraph=Ft;var Jt=ve},402:function(de,I,i){"use strict";var m=i(38663),S=i.n(m),r=i(47828),me=i.n(r),te=i(22385),A=i(47673)},20640:function(de,I,i){"use strict";var m=i(11742),S={"text/plain":"Text","text/html":"Url",default:"Text"},r="Copy to clipboard: #{key}, Enter";function me(A){var h=(/mac os x/i.test(navigator.userAgent)?"\u2318":"Ctrl")+"+C";return A.replace(/#{\s*key\s*}/g,h)}function te(A,h){var H,ge,fe,ae,F,f,le=!1;h||(h={}),H=h.debug||!1;try{fe=m(),ae=document.createRange(),F=document.getSelection(),f=document.createElement("span"),f.textContent=A,f.style.all="unset",f.style.position="fixed",f.style.top=0,f.style.clip="rect(0, 0, 0, 0)",f.style.whiteSpace="pre",f.style.webkitUserSelect="text",f.style.MozUserSelect="text",f.style.msUserSelect="text",f.style.userSelect="text",f.addEventListener("copy",function(K){if(K.stopPropagation(),h.format)if(K.preventDefault(),typeof K.clipboardData=="undefined"){H&&console.warn("unable to use e.clipboardData"),H&&console.warn("trying IE specific stuff"),window.clipboardData.clearData();var oe=S[h.format]||S.default;window.clipboardData.setData(oe,A)}else K.clipboardData.clearData(),K.clipboardData.setData(h.format,A);h.onCopy&&(K.preventDefault(),h.onCopy(K.clipboardData))}),document.body.appendChild(f),ae.selectNodeContents(f),F.addRange(ae);var b=document.execCommand("copy");if(!b)throw new Error("copy command was unsuccessful");le=!0}catch(K){H&&console.error("unable to copy using execCommand: ",K),H&&console.warn("trying IE specific stuff");try{window.clipboardData.setData(h.format||"text",A),h.onCopy&&h.onCopy(window.clipboardData),le=!0}catch(oe){H&&console.error("unable to copy using clipboardData: ",oe),H&&console.error("falling back to prompt"),ge=me("message"in h?h.message:r),window.prompt(ge,A)}}finally{F&&(typeof F.removeRange=="function"?F.removeRange(ae):F.removeAllRanges()),f&&document.body.removeChild(f),fe()}return le}de.exports=te},11742:function(de){de.exports=function(){var I=document.getSelection();if(!I.rangeCount)return function(){};for(var i=document.activeElement,m=[],S=0;S<I.rangeCount;S++)m.push(I.getRangeAt(S));switch(i.tagName.toUpperCase()){case"INPUT":case"TEXTAREA":i.blur();break;default:i=null;break}return I.removeAllRanges(),function(){I.type==="Caret"&&I.removeAllRanges(),I.rangeCount||m.forEach(function(r){I.addRange(r)}),i&&i.focus()}}}}]);
