import 'graphiql/graphiql.min.css';
import React from 'react';
import GraphiQL from 'graphiql';
import PageHeader from 'components/src/PageHeader';

function graphQLFetcher(graphQLParams) {
  return fetch('http://localhost:8080/graphql', {
    method: 'post',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(graphQLParams),
  }).then((response) => response.json());
}

export default function GraphiQLPage() {
  return (
    <div className="graphiql">
      <PageHeader title="GraphiQL" />
      <GraphiQL fetcher={graphQLFetcher} />
    </div>
  );
}
