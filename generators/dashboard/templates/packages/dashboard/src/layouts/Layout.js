import React from 'react';
import { Sidebar } from 'components';

const Layout = ({ children }) => {
  return (
    <div className="app">
      <Sidebar />
      <div className="content">{children}</div>
    </div>
  );
};

export default Layout;
