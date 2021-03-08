import React from 'react';
import { Sidebar } from 'components';

export default function Layout({ children }) {
  return (
    <div className="app">
      <Sidebar />
      <div className="content">{children}</div>
    </div>
  );
}
