import './styles.scss';
import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { render } from 'react-dom';
import ApolloClient from 'apollo-boost';
import { ApolloProvider } from '@apollo/react-hooks';
import { ErrorBoundary } from 'components';
import { Layout } from 'layouts';
import { HomePage, GraphiQLPage } from 'pages';

const ENDPOINT =
  process.env.NODE_ENV === 'development'
    ? 'http://localhost:8080'
    : 'https://api.yousite.com/graphql';

const client = new ApolloClient({
  uri: ENDPOINT + '/graphql',
  request: (operation) => {
    const token = localStorage.getItem('token');
    operation.setContext({
      headers: {
        authorization: token ? `Bearer ${token}` : '',
      },
    });
  },
});

// TODO: defaultOptions in constructor not working
client.defaultOptions = {
  watchQuery: {
    fetchPolicy: 'network-only',
    errorPolicy: 'none',
  },

  query: {
    fetchPolicy: 'network-only',
    errorPolicy: 'none',
  },
};

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
