import React from 'react';
import {
  browserSessionPersistence,
  getAuth,
  GithubAuthProvider,
  signInWithRedirect,
} from 'firebase/auth';

function SigninWithGithub(): JSX.Element {
  const provider = new GithubAuthProvider();
  const auth = getAuth();

  auth.setPersistence(browserSessionPersistence);
  const signin = async () => {
    try {
      await signInWithRedirect(auth, provider);
    } catch (e: unknown) {
      console.error(e);
    }
  };

  return <button onClick={signin}>Sign in with Github</button>;
}

export default SigninWithGithub;
