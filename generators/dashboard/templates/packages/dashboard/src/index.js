import './styles.scss';
import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { render } from 'react-dom';
import ApolloClient from 'apollo-boost';
import { ApolloProvider } from '@apollo/react-hooks';
import { ErrorBoundary } from 'components';
import { Layout } from 'layouts';
import { HomePage, GraphiQLPage } from 'pages';

const client = new ApolloClient({
  uri: 'http://localhost:8080/graphql',
});

const App = () => {
  return (
    <ApolloProvider client={client}>
      <BrowserRouter>
        <Layout>
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/graphiql" element={<GraphiQLPage />} />
            <Route path="*" element={<NotFoundPage />} />
          </Routes>
        </Layout>
      </BrowserRouter>
    </ApolloProvider>
  );
};

const NotFoundPage = () => {
  return <div>Page not found</div>;
};

const Root = () => {
  return (
    <ErrorBoundary>
      <App />
    </ErrorBoundary>
  );
};

render(<Root />, document.getElementById('root'));
