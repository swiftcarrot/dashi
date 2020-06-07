import React from 'react';
import { Sidebar } from 'components';
import { Container } from 'kit';

const Layout = ({ children }) => {
  return (
    <div>
      <Sidebar />
      <Container fluid>{children}</Container>
    </div>
  );
};

export default Layout;
