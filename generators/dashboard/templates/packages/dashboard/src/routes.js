import React from 'react';
import { Route, Switch } from 'wouter';
import HomePage from 'pages/home';
import GraphiQLPage from 'pages/graphiql';

const Routes = () => {
  return (
    <Switch>
      <Route path="/" component={HomePage} />
      <Route path="/graphiql" component={GraphiQLPage} />
      <Route path="/:rest*" component={NotFoundPage} />
    </Switch>
  );
};

const NotFoundPage = () => {
  return <div>Page not found</div>;
};

export default Routes;
