import React from 'react';
import { Router as DefaultRouter, Route, Switch } from 'react-router-dom';
import dynamic from 'umi/dynamic';
import renderRoutes from 'umi/lib/renderRoutes';
import history from '@tmp/history';
import RendererWrapper0 from '/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/pages/.umi/LocaleWrapper.jsx';
import _dvaDynamic from 'dva/dynamic';

const Router = require('dva/router').routerRedux.ConnectedRouter;

const routes = [
  {
    path: '/user',
    component: __IS_BROWSER
      ? _dvaDynamic({
          component: () =>
            import(/* webpackChunkName: "layouts__UserLayout" */ '../../layouts/UserLayout'),
          LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
            .default,
        })
      : require('../../layouts/UserLayout').default,
    routes: [
      {
        path: '/user',
        redirect: '/user/login',
        exact: true,
      },
      {
        path: '/user/login',
        name: 'login',
        component: __IS_BROWSER
          ? _dvaDynamic({
              app: require('@tmp/dva').getApp(),
              models: () => [
                import(/* webpackChunkName: 'p__User__models__register.js' */ '/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/pages/User/models/register.js').then(
                  m => {
                    return { namespace: 'register', ...m.default };
                  },
                ),
              ],
              component: () =>
                import(/* webpackChunkName: "p__User__Login" */ '../User/Login'),
              LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                .default,
            })
          : require('../User/Login').default,
        exact: true,
      },
      {
        component: () =>
          React.createElement(
            require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/node_modules/umi-build-dev/lib/plugins/404/NotFound.js')
              .default,
            { pagesPath: 'src/pages', hasRoutesInConfig: true },
          ),
      },
    ],
  },
  {
    path: '/',
    component: __IS_BROWSER
      ? _dvaDynamic({
          component: () =>
            import(/* webpackChunkName: "layouts__BasicLayout" */ '../../layouts/BasicLayout'),
          LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
            .default,
        })
      : require('../../layouts/BasicLayout').default,
    Routes: [require('../Authorized').default],
    routes: [
      {
        path: '/',
        redirect: '/welcome',
        exact: true,
      },
      {
        path: '/welcome',
        name: '欢迎页',
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__Welcome__Welcome" */ '../Welcome/Welcome'),
              LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                .default,
            })
          : require('../Welcome/Welcome').default,
        authority: ['admin', 'user'],
        exact: true,
      },
      {
        path: '/shopmanagement/configAccount',
        name: '配置账号',
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__Shopmanagement__ConfigAccount" */ '../Shopmanagement/ConfigAccount'),
              LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                .default,
            })
          : require('../Shopmanagement/ConfigAccount').default,
        authority: ['admin'],
        hideInMenu: true,
        exact: true,
      },
      {
        path: '/shopmanagement/ConfigMachine',
        name: '配置机器',
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__Shopmanagement__ConfigMachine" */ '../Shopmanagement/ConfigMachine'),
              LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                .default,
            })
          : require('../Shopmanagement/ConfigMachine').default,
        authority: ['admin'],
        hideInMenu: true,
        exact: true,
      },
      {
        path: '/scenesMenu',
        name: '模拟场景菜单配置',
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__ScenesMenu__ScenesMenu" */ '../ScenesMenu/ScenesMenu'),
              LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                .default,
            })
          : require('../ScenesMenu/ScenesMenu').default,
        authority: ['admin'],
        exact: true,
      },
      {
        path: '/shopmanagement',
        name: 'shopmanagement',
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__Shopmanagement__ShopManagement" */ '../Shopmanagement/ShopManagement'),
              LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                .default,
            })
          : require('../Shopmanagement/ShopManagement').default,
        authority: ['admin'],
        exact: true,
      },
      {
        path: '/shopmanagement/newShop',
        name: '新增店铺',
        hideInMenu: true,
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__Shopmanagement__NewShop" */ '../Shopmanagement/NewShop'),
              LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                .default,
            })
          : require('../Shopmanagement/NewShop').default,
        authority: ['admin'],
        exact: true,
      },
      {
        path: '/shopmanagement/editShop',
        name: '编辑店铺',
        hideInMenu: true,
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__Shopmanagement__EditShop" */ '../Shopmanagement/EditShop'),
              LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                .default,
            })
          : require('../Shopmanagement/EditShop').default,
        authority: ['admin'],
        exact: true,
      },
      {
        path: '/shopmanagement/editImg',
        name: '编辑广告轮播图',
        hideInMenu: true,
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__Shopmanagement__EditImg" */ '../Shopmanagement/EditImg'),
              LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                .default,
            })
          : require('../Shopmanagement/EditImg').default,
        authority: ['admin'],
        exact: true,
      },
      {
        path: '/goodsmanagement/shopinfo',
        name: 'shopinfo',
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__Goodsmanagement__ShopInfo" */ '../Goodsmanagement/ShopInfo'),
              LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                .default,
            })
          : require('../Goodsmanagement/ShopInfo').default,
        authority: ['user'],
        exact: true,
      },
      {
        path: '/goodsmanagement/productmanage',
        name: 'productmanage',
        authority: ['user'],
        routes: [
          {
            path: '/goodsmanagement/productmanage/list',
            name: 'list',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  component: () =>
                    import(/* webpackChunkName: "p__Goodsmanagement__ProductManage" */ '../Goodsmanagement/ProductManage'),
                  LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                    .default,
                })
              : require('../Goodsmanagement/ProductManage').default,
            exact: true,
          },
          {
            path: '/goodsmanagement/productmanage/newproduct',
            name: 'newproduct',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  component: () =>
                    import(/* webpackChunkName: "p__Goodsmanagement__NewProduct" */ '../Goodsmanagement/NewProduct'),
                  LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                    .default,
                })
              : require('../Goodsmanagement/NewProduct').default,
            hideInMenu: true,
            exact: true,
          },
          {
            path: '/goodsmanagement/productmanage/editgoods',
            name: 'editgoods',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  component: () =>
                    import(/* webpackChunkName: "p__Goodsmanagement__EditGoods" */ '../Goodsmanagement/EditGoods'),
                  LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                    .default,
                })
              : require('../Goodsmanagement/EditGoods').default,
            hideInMenu: true,
            exact: true,
          },
          {
            path: '/goodsmanagement/productmanage/colormanage',
            name: 'colormanage',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  component: () =>
                    import(/* webpackChunkName: "p__Goodsmanagement__ColorManage" */ '../Goodsmanagement/ColorManage'),
                  LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                    .default,
                })
              : require('../Goodsmanagement/ColorManage').default,
            hideInMenu: true,
            exact: true,
          },
          {
            path: '/goodsmanagement/productmanage/hotgoods',
            name: 'hotgoods',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  component: () =>
                    import(/* webpackChunkName: "p__Goodsmanagement__Hotgoods" */ '../Goodsmanagement/Hotgoods'),
                  LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                    .default,
                })
              : require('../Goodsmanagement/Hotgoods').default,
            exact: true,
          },
          {
            component: () =>
              React.createElement(
                require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/node_modules/umi-build-dev/lib/plugins/404/NotFound.js')
                  .default,
                { pagesPath: 'src/pages', hasRoutesInConfig: true },
              ),
          },
        ],
      },
      {
        path: '/goodsmanagement/menumanage',
        name: 'menumanage',
        component: __IS_BROWSER
          ? _dvaDynamic({
              component: () =>
                import(/* webpackChunkName: "p__Goodsmanagement__MenuManage" */ '../Goodsmanagement/MenuManage'),
              LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                .default,
            })
          : require('../Goodsmanagement/MenuManage').default,
        authority: ['user'],
        exact: true,
      },
      {
        name: 'exception',
        icon: 'warning',
        path: '/exception',
        hideInMenu: true,
        routes: [
          {
            path: '/exception/403',
            name: 'not-permission',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  app: require('@tmp/dva').getApp(),
                  models: () => [
                    import(/* webpackChunkName: 'p__Exception__models__error.js' */ '/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/pages/Exception/models/error.js').then(
                      m => {
                        return { namespace: 'error', ...m.default };
                      },
                    ),
                  ],
                  component: () =>
                    import(/* webpackChunkName: "p__Exception__403" */ '../Exception/403'),
                  LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                    .default,
                })
              : require('../Exception/403').default,
            exact: true,
          },
          {
            path: '/exception/404',
            name: 'not-find',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  app: require('@tmp/dva').getApp(),
                  models: () => [
                    import(/* webpackChunkName: 'p__Exception__models__error.js' */ '/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/pages/Exception/models/error.js').then(
                      m => {
                        return { namespace: 'error', ...m.default };
                      },
                    ),
                  ],
                  component: () =>
                    import(/* webpackChunkName: "p__Exception__404" */ '../Exception/404'),
                  LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                    .default,
                })
              : require('../Exception/404').default,
            exact: true,
          },
          {
            path: '/exception/500',
            name: 'server-error',
            component: __IS_BROWSER
              ? _dvaDynamic({
                  app: require('@tmp/dva').getApp(),
                  models: () => [
                    import(/* webpackChunkName: 'p__Exception__models__error.js' */ '/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/pages/Exception/models/error.js').then(
                      m => {
                        return { namespace: 'error', ...m.default };
                      },
                    ),
                  ],
                  component: () =>
                    import(/* webpackChunkName: "p__Exception__500" */ '../Exception/500'),
                  LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                    .default,
                })
              : require('../Exception/500').default,
            hideInMenu: true,
            exact: true,
          },
          {
            path: '/exception/trigger',
            name: 'trigger',
            hideInMenu: true,
            component: __IS_BROWSER
              ? _dvaDynamic({
                  app: require('@tmp/dva').getApp(),
                  models: () => [
                    import(/* webpackChunkName: 'p__Exception__models__error.js' */ '/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/pages/Exception/models/error.js').then(
                      m => {
                        return { namespace: 'error', ...m.default };
                      },
                    ),
                  ],
                  component: () =>
                    import(/* webpackChunkName: "p__Exception__TriggerException" */ '../Exception/TriggerException'),
                  LoadingComponent: require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/src/components/PageLoading/index')
                    .default,
                })
              : require('../Exception/TriggerException').default,
            exact: true,
          },
          {
            component: () =>
              React.createElement(
                require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/node_modules/umi-build-dev/lib/plugins/404/NotFound.js')
                  .default,
                { pagesPath: 'src/pages', hasRoutesInConfig: true },
              ),
          },
        ],
      },
      {
        component: () =>
          React.createElement(
            require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/node_modules/umi-build-dev/lib/plugins/404/NotFound.js')
              .default,
            { pagesPath: 'src/pages', hasRoutesInConfig: true },
          ),
      },
    ],
  },
  {
    component: () =>
      React.createElement(
        require('/Users/xuehuiwan/Documents/1031/ClothMall/frontend/web/node_modules/umi-build-dev/lib/plugins/404/NotFound.js')
          .default,
        { pagesPath: 'src/pages', hasRoutesInConfig: true },
      ),
  },
];
window.g_routes = routes;
const plugins = require('umi/_runtimePlugin');
plugins.applyForEach('patchRoutes', { initialValue: routes });

export { routes };

export default class RouterWrapper extends React.Component {
  unListen() {}

  constructor(props) {
    super(props);

    // route change handler
    function routeChangeHandler(location, action) {
      plugins.applyForEach('onRouteChange', {
        initialValue: {
          routes,
          location,
          action,
        },
      });
    }
    this.unListen = history.listen(routeChangeHandler);
    routeChangeHandler(history.location);
  }

  componentWillUnmount() {
    this.unListen();
  }

  render() {
    const props = this.props || {};
    return (
      <RendererWrapper0>
        <Router history={history}>{renderRoutes(routes, props)}</Router>
      </RendererWrapper0>
    );
  }
}
