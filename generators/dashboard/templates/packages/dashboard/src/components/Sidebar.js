import React from 'react';
import { Link } from 'kit';

const Sidebar = (props) => {
  return (
    <div>
      <Link to="/">Home</Link>
      <Link to="/graphiql">GraphiQL</Link>
    </div>
  );
};

export default Sidebar;
