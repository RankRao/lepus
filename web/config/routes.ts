export default [
  {
    path: '/user',
    layout: false,
    routes: [
      {
        path: '/user',
        routes: [
          {
            name: 'login',
            path: '/user/login',
            component: './user/login',
          },
        ],
      },
    ],
  },
  { path: '/', redirect: '/portal/' },

  {
    name: 'portal',
    icon: 'ConsoleSqlOutlined',
    path: '/portal/',
    component: './portal/',
  },

  {
    name: 'execute',
    icon: 'ConsoleSqlOutlined',
    path: '/execute/',
    component: './execute/',
  },

  {
    name: 'meta',
    icon: 'database',
    path: '/meta',
    routes: [
      { path: '/meta', redirect: '/meta/dashboard' },
      {
        path: '/meta/dashboard',
        name: 'dashboard',
        component: './meta/dashboard',
        icon: 'BlockOutlined',
      },
      {
        path: '/meta/instance',
        name: 'instance',
        component: './meta/instance',
        icon: 'BlockOutlined',
      },
      {
        path: '/meta/database',
        name: 'database',
        component: './meta/database',
        icon: 'BlockOutlined',
      },
      {
        path: '/meta/table',
        name: 'table',
        component: './meta/table',
        icon: 'BlockOutlined',
      },
      {
        path: '/meta/column',
        name: 'column',
        component: './meta/column',
        icon: 'BlockOutlined',
      },
    ],
  },


  {
    name: 'monitor',
    icon: 'DashboardOutlined',
    path: '/monitor',
    routes: [
      { path: '/monitor', redirect: '/monitor/dashboard' },
      {
        name: 'dashboard',
        icon: 'DashboardOutlined',
        path: '/monitor/dashboard/',
        component: './monitor/dashboard/index',
      },
      {
        name: 'event',
        icon: 'BarChartOutlined',
        path: '/monitor/event/',
        component: './monitor/event/index',
      },
    ]
  },
  // {
  //   name: 'event',
  //   icon: 'FundProjectionScreenOutlined',
  //   path: '/event/',
  //   component: './event/index',
  // },
  {
    name: 'alarm',
    icon: 'WarningOutlined',
    path: '/alarm',
    //component: './Alarm/Index',
    routes: [
      { path: '/alarm', redirect: '/alarm/event' },
      {
        path: '/alarm/event',
        name: 'alarmEvent',
        component: './alarm/event',
        icon: 'MailOutlined',
      },
      {
        path: '/alarm/level',
        name: 'alarmLevel',
        component: './alarm/level',
        icon: 'OrderedListOutlined',
        //access: 'canAdmin',
      },
      {
        path: '/alarm/channel',
        name: 'alarmChannel',
        component: './alarm/channel',
        icon: 'UsergroupAddOutlined',
        //access: 'canAdmin',
      },
      {
        path: '/alarm/rule',
        name: 'alarmRule',
        component: './alarm/rule',
        icon: 'SettingOutlined',
        //access: 'canAdmin',
      },
      {
        path: '/alarm/test',
        name: 'alarmTest',
        component: './alarm/test',
        icon: 'HeartOutlined',
        //access: 'canAdmin',
      },

    ],
  },
  {
    name: 'userManager',
    icon: 'UserOutlined',
    path: '/users/manager',
    component: './UserManager/index',
    access: 'canAdmin',
  },
  {
    name: 'datasource',
    icon: 'SettingOutlined',
    path: '/setting',
    //access: 'canAdmin',
    routes: [
      { path: '/setting', redirect: '/setting/datasource' },
      {
        path: '/setting/idc',
        name: 'idc',
        component: './setting/idc',
        icon: 'CloudServerOutlined',
        //access: 'canAdmin',
      },
      {
        path: '/setting/env',
        name: 'env',
        component: './setting/env',
        icon: 'ChromeOutlined',
        //access: 'canAdmin',
      },
      {
        path: '/setting/datasource_type',
        name: 'datasource_type',
        component: './setting/datasource_type',
        icon: 'CodeSandboxOutlined',
        //access: 'canAdmin',
      },
      {
        path: '/setting/datasource',
        name: 'datasource',
        component: './setting/datasource',
        icon: 'CloudOutlined',
        //access: 'canAdmin',
      },

    ],
  },

  {
    name: 'task',
    icon: 'MenuUnfoldOutlined',
    path: '/task',
    routes: [
      { path: '/task', redirect: '/task/option' },
      {
        path: '/task/option',
        name: 'option',
        component: './task/option',
        icon: 'OrderedListOutlined',
      },
      {
        path: '/task/heartbeat',
        name: 'heartbeat',
        component: './task/heartbeat',
        icon: 'HeatMapOutlined',
      },
    ],
  },
  {
    name: 'support',
    icon: 'BulbOutlined',
    path: '/support/',
    component: './support/index',
  },
  {
    component: './404',
  },
];
