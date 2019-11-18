import React, { PureComponent } from 'react';
import { connect } from 'dva';
import styles from './welcome.less';
@connect(({ loading }) => ({
  listLoading: loading.effects['list/fetch'],
}))
class welcome extends PureComponent {
  state = {};

  componentDidMount() {}

  render() {
    return <div className={styles.welcome}>
      {/* <div>admin</div> */}
    <div>
    您好
    </div>
    <div>
    欢迎使用布艺商城管理系统
    </div>
    </div>;
  }
}
export default welcome;
