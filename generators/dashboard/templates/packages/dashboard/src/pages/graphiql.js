import React from 'react';
import GraphiQL from 'graphiql';
import 'graphiql/graphiql.min.css';

function graphQLFetcher(graphQLParams) {
  return fetch('http://localhost:8080/graphql', {
    method: 'post',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(graphQLParams),
  }).then((response) => response.json());
}

const GraphiQLPage = () => {
  return (
    <div className="graphiql">
      <GraphiQL fetcher={graphQLFetcher} />
    </div>
  );
};

export default GraphiQLPage;
