export default [
  // user
  {
    path: '/user',
    component: '../layouts/UserLayout',
    routes: [
      { path: '/user', redirect: '/user/login' },
      { path: '/user/login', name: 'login', component: './User/Login' },
    ],
  },
  // app
  {
    path: '/',
    component: '../layouts/BasicLayout',
    Routes: ['src/pages/Authorized'],
    routes: [
      // 开始页面\欢迎页
      {
        path: '/',
        redirect: '/welcome',
      },
      {
        path: '/welcome',
        name: '欢迎页',
        component: './Welcome/Welcome',
        authority: ['admin','user'],
      },
      // 账号配置
      {
        path: '/shopmanagement/configAccount',
        name: '配置账号',
        component: './Shopmanagement/ConfigAccount',
        authority: ['admin'],
        hideInMenu: true
      },
      // 机器配置
      {
        path: '/shopmanagement/ConfigMachine',
        name: '配置机器',
        component: './Shopmanagement/ConfigMachine',
        authority: ['admin'],
        hideInMenu: true
      },
      // 模拟场景菜单配置
      {
        path: '/scenesMenu',
        name: '模拟场景菜单配置',
        component: './ScenesMenu/ScenesMenu',
        authority: ['admin'],
      },
      // 店铺管理
      {
        path: '/shopmanagement',
        name: 'shopmanagement',
        component: './Shopmanagement/ShopManagement',
        authority: ['admin'],
      },
        // 店铺管理、新增店铺
        {
          path: '/shopmanagement/newShop',
          name: '新增店铺',
          hideInMenu: true,
          component: './Shopmanagement/NewShop',
          authority: ['admin'],
        },
           // 店铺管理、编辑店铺
           {
            path: '/shopmanagement/editShop',
            name: '编辑店铺',
            hideInMenu: true,
            component: './Shopmanagement/EditShop',
            authority: ['admin'],
          },
             // 店铺管理、编辑广告轮播图
             {
              path: '/shopmanagement/editImg',
              name: '编辑广告轮播图',
              hideInMenu: true,
              component: './Shopmanagement/EditImg',
              authority: ['admin'],
            },
      // 产品管理
      {
        path: '/goodsmanagement/shopinfo',
        name: 'shopinfo',
        component: './Goodsmanagement/ShopInfo',
        authority: ['user'],
      },
      {
        path: '/goodsmanagement/productmanage',
        name: 'productmanage',
        authority: ['user'],
        routes: [
          {
            path: '/goodsmanagement/productmanage/list',
            name: 'list',
            component: './Goodsmanagement/ProductManage',
          },
          {
            path: '/goodsmanagement/productmanage/newproduct',
            name: 'newproduct',
            component: './Goodsmanagement/NewProduct',
            hideInMenu: true
          },
          {
            path:'/goodsmanagement/productmanage/editgoods',
            name:'editgoods',
            component:'./Goodsmanagement/EditGoods',
            hideInMenu: true
          },
          {
            path:'/goodsmanagement/productmanage/colormanage',
            name:'colormanage',
            component:'./Goodsmanagement/ColorManage',
            hideInMenu: true
          },
          {
            path: '/goodsmanagement/productmanage/hotgoods',
            name: 'hotgoods',
            component: './Goodsmanagement/Hotgoods'
          }
        ]
      },
     
      {
        path: '/goodsmanagement/menumanage',
        name: 'menumanage',
        component: './Goodsmanagement/MenuManage',
        authority: ['user'],
      },

      // {
      //   path: '/goodsmanagement',
      //   name: 'goodsmanagement',
      //   icon: 'dashboard',
      //   // component: './Welcome/Welcome',
      //   authority: ['user'],
      //   routes: [
      //     {
      //       path: '/goodsmanagement/productmanage',
      //       name: 'productmanage',
      //       // component:'./Goodsmanagement/ProductManage',
      //       routes: [
      //         {
      //           path: '/goodsmanagement/productmanage/list',
      //           name: 'list',
      //           component: './Goodsmanagement/ProductManage',
      //         },
      //         {
      //           path: '/goodsmanagement/productmanage/newproduct',
      //           name: 'newproduct',
      //           component: './Goodsmanagement/NewProduct',
      //           hideInMenu: true
      //         },
      //         {
      //           path:'/goodsmanagement/productmanage/editgoods',
      //           name:'editgoods',
      //           component:'./Goodsmanagement/EditGoods',
      //           hideInMenu: true
      //         },
      //         {
      //           path:'/goodsmanagement/productmanage/colormanage',
      //           name:'colormanage',
      //           component:'./Goodsmanagement/ColorManage',
      //           hideInMenu: true
      //         },
      //         {
      //           path: '/goodsmanagement/productmanage/hotgoods',
      //           name: 'hotgoods',
      //           component: './Goodsmanagement/Hotgoods'
      //         }
      //       ]
      //     },
      //     {
      //       path: '/goodsmanagement/shopinfo',
      //       name: 'shopinfo',
      //       component: './Goodsmanagement/ShopInfo'
      //     },
      //     {
      //       path: '/goodsmanagement/menumanage',
      //       name: 'menumanage',
      //       component: './Goodsmanagement/MenuManage'
      //     }
      //   ]
      // },
      // 以下为脚手架生成的
      // dashboard
      // {
      //   path: '/dashboard',
      //   name: 'dashboard',
      //   icon: 'dashboard',
      //   routes: [
      //     {
      //       path: '/dashboard/analysis',
      //       name: 'analysis',
      //       component: './Dashboard/Analysis',
      //     },
      //     {
      //       path: '/dashboard/monitor',
      //       name: 'monitor',
      //       component: './Dashboard/Monitor',
      //     },
      //     {
      //       path: '/dashboard/workplace',
      //       name: 'workplace',
      //       component: './Dashboard/Workplace',
      //     },
      //   ],
      // },
      // // forms
      // {
      //   path: '/form',
      //   icon: 'form',
      //   name: 'form',
      //   routes: [
      //     {
      //       path: '/form/basic-form',
      //       name: 'basicform',
      //       component: './Forms/BasicForm',
      //     },
      //     {
      //       path: '/form/step-form',
      //       name: 'stepform',
      //       component: './Forms/StepForm',
      //       hideChildrenInMenu: true,
      //       routes: [
      //         {
      //           path: '/form/step-form',
      //           redirect: '/form/step-form/info',
      //         },
      //         {
      //           path: '/form/step-form/info',
      //           name: 'info',
      //           component: './Forms/StepForm/Step1',
      //         },
      //         {
      //           path: '/form/step-form/confirm',
      //           name: 'confirm',
      //           component: './Forms/StepForm/Step2',
      //         },
      //         {
      //           path: '/form/step-form/result',
      //           name: 'result',
      //           component: './Forms/StepForm/Step3',
      //         },
      //       ],
      //     },
      //     {
      //       path: '/form/advanced-form',
      //       name: 'advancedform',
      //       authority: ['admin'],
      //       component: './Forms/AdvancedForm',
      //     },
      //   ],
      // },
      // // list
      // {
      //   path: '/list',
      //   icon: 'table',
      //   name: 'list',
      //   routes: [
      //     {
      //       path: '/list/table-list',
      //       name: 'searchtable',
      //       component: './List/TableList',
      //     },
      //     {
      //       path: '/list/basic-list',
      //       name: 'basiclist',
      //       component: './List/BasicList',
      //     },
      //     {
      //       path: '/list/card-list',
      //       name: 'cardlist',
      //       component: './List/CardList',
      //     },
      //     {
      //       path: '/list/search',
      //       name: 'searchlist',
      //       component: './List/List',
      //       routes: [
      //         {
      //           path: '/list/search/articles',
      //           name: 'articles',
      //           component: './List/Articles',
      //           hideInMenu: true
      //         },
      //         {
      //           path: '/list/search/projects',
      //           name: 'projects',
      //           component: './List/Projects',
      //           hideInMenu: true
      //         },
      //         {
      //           path: '/list/search/applications',
      //           name: 'applications',
      //           component: './List/Applications',
      //           hideInMenu: true
      //         },
      //       ],
      //     },
      //   ],
      // },
      // {
      //   path: '/profile',
      //   name: 'profile',
      //   icon: 'profile',
      //   routes: [
      //     // profile
      //     {
      //       path: '/profile/basic',
      //       name: 'basic',
      //       component: './Profile/BasicProfile',
      //     },
      //     {
      //       path: '/profile/basic/:id',
      //       hideInMenu: true,
      //       component: './Profile/BasicProfile',
      //     },
      //     {
      //       path: '/profile/advanced',
      //       name: 'advanced',
      //       authority: ['admin'],
      //       component: './Profile/AdvancedProfile',
      //     },
      //   ],
      // },
      // {
      //   name: 'result',
      //   icon: 'check-circle-o',
      //   path: '/result',
      //   routes: [
      //     // result
      //     {
      //       path: '/result/success',
      //       name: 'success',
      //       component: './Result/Success',
      //     },
      //     { path: '/result/fail', name: 'fail', component: './Result/Error' },
      //   ],
      // },
      {
        name: 'exception',
        icon: 'warning',
        path: '/exception',
        hideInMenu: true,
        routes: [
          // exception
          {
            path: '/exception/403',
            name: 'not-permission',
            component: './Exception/403',
          },
          {
            path: '/exception/404',
            name: 'not-find',
            component: './Exception/404',
          },
          {
            path: '/exception/500',
            name: 'server-error',
            component: './Exception/500',
              hideInMenu: true,
          },
          {
            path: '/exception/trigger',
            name: 'trigger',
            hideInMenu: true,
            component: './Exception/TriggerException',
          },
        ],
      },
      // {
      //   name: 'account',
      //   icon: 'user',
      //   path: '/account',
      //   routes: [
      //     {
      //       path: '/account/center',
      //       name: 'center',
      //       component: './Account/Center/Center',
      //       routes: [
      //         {
      //           path: '/account/center',
      //           redirect: '/account/center/articles',
      //         },
      //         {
      //           path: '/account/center/articles',
      //           component: './Account/Center/Articles',
      //         },
      //         {
      //           path: '/account/center/applications',
      //           component: './Account/Center/Applications',
      //         },
      //         {
      //           path: '/account/center/projects',
      //           component: './Account/Center/Projects',
      //         },
      //       ],
      //     },
      //     {
      //       path: '/account/settings',
      //       name: 'settings',
      //       component: './Account/Settings/Info',
      //       routes: [
      //         {
      //           path: '/account/settings',
      //           redirect: '/account/settings/base',
      //         },
      //         {
      //           path: '/account/settings/base',
      //           component: './Account/Settings/BaseView',
      //         },
      //         {
      //           path: '/account/settings/security',
      //           component: './Account/Settings/SecurityView',
      //         },
      //         {
      //           path: '/account/settings/binding',
      //           component: './Account/Settings/BindingView',
      //         },
      //         {
      //           path: '/account/settings/notification',
      //           component: './Account/Settings/NotificationView',
      //         },
      //       ],
      //     },
      //   ],
      // },
      // //  editor
      // {
      //   name: 'editor',
      //   icon: 'highlight',
      //   path: '/editor',
      //   routes: [
      //     {
      //       path: '/editor/flow',
      //       name: 'flow',
      //       component: './Editor/GGEditor/Flow',
      //     },
      //     {
      //       path: '/editor/mind',
      //       name: 'mind',
      //       component: './Editor/GGEditor/Mind',
      //     },
      //     {
      //       path: '/editor/koni',
      //       name: 'koni',
      //       component: './Editor/GGEditor/Koni',
      //     },
      //   ],
      // },
    ],
  },
];
