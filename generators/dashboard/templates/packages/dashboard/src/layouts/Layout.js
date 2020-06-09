import React, { Fragment } from 'react';
import { Sidebar } from 'components';

const Layout = ({ children }) => {
  return (
    <Fragment>
      <Sidebar />
      <div className="content">{children}</div>
    </Fragment>
  );
};

export default Layout;
