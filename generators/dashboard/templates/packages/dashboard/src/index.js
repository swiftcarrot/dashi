import './styles.scss';
import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { render } from 'react-dom';
import { Layout } from 'layouts';
import HomePage from 'pages/home';
import GraphiQLPage from 'pages/graphiql';

const App = () => {
  return (
    <BrowserRouter>
      <Layout>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/graphiql" element={<GraphiQLPage />} />
          <Route path="*" element={<NotFoundPage />} />
        </Routes>
      </Layout>
    </BrowserRouter>
  );
};

const NotFoundPage = () => {
  return <div>Page not found</div>;
};

render(<App />, document.getElementById('root'));
