import React from 'react';

export default function Feed() {
  return (
    <ul>
      {[
        'use react-router for /login, /new, /:post_id',
        'Create api gateway',
        'figure out auth',
        'basic workflow: no auth, submit post and picture, display it',
      ].map((task) => <li>{task}</li>)}
    </ul>
  );
}
