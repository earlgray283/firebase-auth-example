import { getAuth, User } from 'firebase/auth';
import { useEffect, useState } from 'react';
import './App.css';
import SigninWithGithub from './components/auth/github';
import './lib/firebase';
import React from 'react';
import Home from './pages/home';
import { postJson } from './lib/axios';

function App() {
  const [loading, setLoading] = useState(true);
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const auth = getAuth();
  useEffect(() => {
    auth.onAuthStateChanged(async (user) => {
      if (user) {
        const idToken = await user.getIdToken();
        const resp = await postJson('/auth/id-token', { idToken: idToken });
        if (resp.status !== 200) {
          console.error(resp.statusText);
        }
      }
      setCurrentUser(user);
      setLoading(false);
    });
  });
  return (
    <div className='App'>
      {loading ? (
        <p>loading...</p>
      ) : currentUser !== null ? (
        <Home />
      ) : (
        <SigninWithGithub />
      )}
    </div>
  );
}

export default App;
