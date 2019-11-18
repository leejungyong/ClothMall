import { routerRedux } from 'dva/router';
import { stringify } from 'qs';
import { fakeAccountLogin, getFakeCaptcha } from '@/services/api';
import { setAuthority } from '@/utils/authority';
import { getPageQuery } from '@/utils/utils';
import { reloadAuthorized } from '@/utils/Authorized';
import axios from "axios";
import router from 'umi/router';
import qs from "qs";
export default {
  namespace: 'login',

  state: {
    status: undefined,
  },

  effects: {
    *login({ payload }, { call, put }) {
      axios.post('/api2',qs.stringify({phonenum:payload.userName,token:'Jh2044695',cmd:'manageLogin',password:payload.password})).then(res=>{
        if(res.data.success){
              localStorage.setItem("manageid",res.data.data.manageid);
              localStorage.setItem("shopid",res.data.data.shopid);
              res.data.data.shopid==0 &&setAuthority('admin');
              res.data.data.shopid>0  &&setAuthority('user');
              reloadAuthorized();
                router.push({
                  pathname: '/welcome',
                });
             }      
            });
    },

    *getCaptcha({ payload }, { call }) {
      yield call(getFakeCaptcha, payload);
    },

    *logout(_, { put }) {
      yield put({
        type: 'changeLoginStatus',
        payload: {
          status: false,
          currentAuthority: 'guest',
        },
      });
      reloadAuthorized();
      const { redirect } = getPageQuery();
      // redirect
      if (window.location.pathname !== '/user/login' && !redirect) {
        yield put(
          routerRedux.replace({
            pathname: '/user/login',
            search: stringify({
              redirect: window.location.href,
            }),
          })
        );
      }
    },
  },

  reducers: {
    changeLoginStatus(state, { payload }) {
      setAuthority(payload.currentAuthority);
      return {
        ...state,
        status: payload.status,
        type: payload.type,
      };
    },
  },
};
