import "./styles.scss";
import React from "react";
import { render } from "react-dom";
import {
  ApolloProvider,
  ApolloClient,
  InMemoryCache,
  createHttpLink,
} from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import { ErrorBoundary } from "components";
import Pages from "@swiftcarrot/react-router/src/loader!src/pages";

const ENDPOINT =
  process.env.NODE_ENV === "development"
    ? "http://localhost:8080"
    : "https://api.yousite.com";

const httpLink = createHttpLink({
  uri: ENDPOINT + "/graphql",
});

const authLink = setContext((_, { headers }) => {
  const token = localStorage.getItem("token");
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : "",
    },
  };
});

const client = new ApolloClient({
  link: authLink.concat(httpLink),
  cache: new InMemoryCache(),
  defaultOptions: {
    watchQuery: {
      fetchPolicy: "network-only",
      errorPolicy: "none",
    },
    query: {
      fetchPolicy: "network-only",
      errorPolicy: "none",
    },
  },
});

function Root() {
  return (
    <ErrorBoundary>
      <ApolloProvider client={client}>
        <Pages />
      </ApolloProvider>
    </ErrorBoundary>
  );
}

render(<Root />, document.getElementById("root"));
