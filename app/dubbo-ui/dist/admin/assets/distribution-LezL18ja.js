import{d as A,k as E,b as F,u as U,g as G,n as m,r as S,s as H,f as v,c as o,t as r,e as i,h as n,P as w,o as c,y as u,H as N,I as O,z as g,J as f,F as J,v as L,_ as M}from"./index-sY3lDj5b.js";import{g as Y}from"./service-ZaU-tn05.js";import{f as j}from"./DateUtil-4WtOJDVl.js";import"./request-ouf_oWqW.js";const q={class:"__container_services_tabs_distribution"},K=["onClick"],Q=["onClick"],W=A({__name:"distribution",setup(X){E(e=>({"675d66e3":n(w)}));const C=F(),z=U(),{appContext:{config:{globalProperties:x}}}=G(),I=m(""),D=S([{label:"不指定",value:""},{label:"version=1.0.0",value:"version=1.0.0"},{label:"group=group1",value:"group=group1"},{label:"version=1.0.0,group=group1",value:"version=1.0.0,group=group1"}]);m(D[0].value);const b=m("provider"),P=[{title:"应用名",dataIndex:"appName",width:"20%",customCell:(e,a)=>a===0?{rowSpan:d.value.length}:{rowSpan:0}},{title:"实例数",dataIndex:"instanceNum",width:"15%",customCell:(e,a)=>a===0?{rowSpan:d.value.length}:{rowSpan:0}},{title:"实例名",dataIndex:"instanceName",width:"15%"},{title:"RPC端口",dataIndex:"rpcPort",width:"15%"},{title:"超时时间",dataIndex:"timeOut",width:"10%"},{title:"重试次数",dataIndex:"retryNum",width:"10%"},{title:"标签",dataIndex:"label",width:"15%"}],d=m([]),t=S({total:0,pageSize:10,current:1,pageOffset:0,showTotal:e=>x.$t("searchDomain.total")+": "+e+" "+x.$t("searchDomain.unit")}),k=async()=>{var _;let e={serviceName:(_=z.params)==null?void 0:_.pathId,side:b.value,pageOffset:t.pageOffset,pageSize:t.pageSize};const{data:{list:a,pageInfo:p}}=await Y(e);d.value=a,t.total=p.Total};k();const h=H.debounce(k,300),R=e=>{t.pageSize=e.pageSize||10,t.current=e.current||1,t.pageOffset=(t.current-1)*t.pageSize,h()};return(e,a)=>{const p=i("a-radio-button"),_=i("a-radio-group"),V=i("a-input-search"),y=i("a-flex"),B=i("a-tag"),T=i("a-table");return c(),v("div",q,[o(y,{vertical:""},{default:r(()=>[o(y,{class:"service-filter"},{default:r(()=>[o(_,{value:b.value,"onUpdate:value":a[0]||(a[0]=s=>b.value=s),"button-style":"solid",onClick:n(h)},{default:r(()=>[o(p,{value:"provider"},{default:r(()=>[u("生产者")]),_:1}),o(p,{value:"consumer"},{default:r(()=>[u("消费者")]),_:1})]),_:1},8,["value","onClick"]),o(V,{value:I.value,"onUpdate:value":a[1]||(a[1]=s=>I.value=s),placeholder:"搜索应用，ip，支持前缀搜索",class:"service-filter-input",onSearch:n(h),"enter-button":""},null,8,["value","onSearch"])]),_:1}),o(T,{columns:P,"data-source":d.value,scroll:{y:"45vh"},pagination:t,onChange:R},{bodyCell:r(({column:s,text:l})=>[s.dataIndex==="appName"?(c(),v("span",{key:0,class:"link",onClick:$=>n(C).push("/resources/applications/detail/"+l)},[N("b",null,[o(n(O),{style:{"margin-bottom":"-2px"},icon:"material-symbols:attach-file-rounded"}),u(" "+g(l),1)])],8,K)):f("",!0),s.dataIndex==="instanceName"?(c(),v("span",{key:1,class:"link",onClick:$=>n(C).push("/resources/instances/detail/"+l)},[N("b",null,[o(n(O),{style:{"margin-bottom":"-2px"},icon:"material-symbols:attach-file-rounded"}),u(" "+g(l),1)])],8,Q)):f("",!0),s.dataIndex==="timeOut"?(c(),v(J,{key:2},[u(g(n(j)(l)),1)],64)):f("",!0),s.dataIndex==="label"?(c(),L(B,{key:3,color:n(w)},{default:r(()=>[u(g(l),1)]),_:2},1032,["color"])):f("",!0)]),_:1},8,["data-source","pagination"])]),_:1})])}}}),oe=M(W,[["__scopeId","data-v-bf7e4472"]]);export{oe as default};
