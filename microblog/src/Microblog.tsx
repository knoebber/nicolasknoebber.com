import React from 'react';
import {
  BrowserRouter as Router,
  Route,
  Switch,
} from 'react-router-dom';
import Feed from './components/feed';
import Post from './components/post';
import Admin from './components/admin';

export default function Microblog() {
  return (
    <Router basename="/microblog">
      <Switch>
        <Route exact path="/post/:id">
          <Post />
        </Route>
        <Route exact path="/admin">
          <Admin />
        </Route>
        <Route path="*">
          <Feed />
        </Route>
      </Switch>
    </Router>
  );
}
