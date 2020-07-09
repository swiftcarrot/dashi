import React from 'react';

const PageHeader = ({ title, extra }) => {
  return (
    <div className="page-header">
      <div className="page-header-title">{title}</div>
      {extra}
    </div>
  );
};

export default PageHeader;
