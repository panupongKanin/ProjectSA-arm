//import { Meteor } from 'meteor/meteor';
import React, { useState } from 'react';

export const LoginForm = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const submit = (e: React.FormEvent) => {
    e.preventDefault();
    //Meteor.loginWithPassword({ email }, password);
  };

  return (
    <form onSubmit={submit} className="login-form">
      <label htmlFor="email">Email</label>

      <input
        type="text"
        placeholder="Email"
        name="email"
        required
        onChange={e => setEmail(e.target.value)}
      />

      <label htmlFor="password">Password</label>

      <input
        type="password"
        placeholder="Password"
        name="password"
        required
        onChange={e => setPassword(e.target.value)}
      />

      <button type="submit">Login</button>
    </form>
  );
};