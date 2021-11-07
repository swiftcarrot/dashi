import React from "react";
import Sidebar from "src/components/Sidebar";

export default function Layout({ children }) {
  return (
    <div className="app">
      <Sidebar />
      <div className="content">{children}</div>
    </div>
  );
}
