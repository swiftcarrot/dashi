import React, { Fragment } from 'react';
import { css, Global } from '@emotion/core';
import GraphiQL from 'graphiql';
import 'graphiql/graphiql.min.css';

function graphQLFetcher(graphQLParams) {
  return fetch('http://localhost:8080', {
    method: 'post',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(graphQLParams),
  }).then((response) => response.json());
}

const GraphiQLPage = () => {
  return (
    <Fragment>
      <Global
        styles={css`
          #graphiql {
            height: 100vh;
          }
        `}
      />
      <GraphiQL fetcher={graphQLFetcher} />
    </Fragment>
  );
};

export default GraphiQLPage;
