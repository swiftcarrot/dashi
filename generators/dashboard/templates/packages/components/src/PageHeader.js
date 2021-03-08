import React from 'react';

export default function PageHeader({ title, extra }) {
  return (
    <div className="page-header">
      <div className="page-header-title">{title}</div>
      {extra}
    </div>
  );
}
