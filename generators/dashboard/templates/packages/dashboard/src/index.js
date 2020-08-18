import './styles.scss';
import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { render } from 'react-dom';
import {
  ApolloProvider,
  ApolloClient,
  InMemoryCache,
  createHttpLink,
} from '@apollo/client';
import { setContext } from '@apollo/client/link/context';
import { ErrorBoundary } from 'components';
import { Layout } from 'src/layouts';
import { HomePage, GraphiQLPage } from 'src/pages';

const ENDPOINT =
  process.env.NODE_ENV === 'development'
    ? 'http://localhost:8080'
    : 'https://api.yousite.com/graphql';

const httpLink = createHttpLink({
  uri: ENDPOINT + '/graphql',
});

const authLink = setContext((_, { headers }) => {
  const token = localStorage.getItem('token');
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : '',
    },
  };
});

const client = new ApolloClient({
  link: authLink.concat(httpLink),
  cache: new InMemoryCache(),
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
