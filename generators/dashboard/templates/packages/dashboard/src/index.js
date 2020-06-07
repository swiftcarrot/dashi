import React from 'react';
import { render } from 'react-dom';
import { Layout } from 'layouts';
import Routes from './routes';

const App = () => {
  return (
    <Layout>
      <Routes />
    </Layout>
  );
};

render(<App />, document.getElementById('root'));
