import React, { useState } from 'react';
import { http } from '../lib/axios';

function Home(): JSX.Element {
  const [respText, setRespText] = useState<string | undefined>();
  const handleWhoami = async () => {
    const resp = await http.get<string>('/whoami');
    setRespText(resp.data);
  };
  return (
    <div>
      <button onClick={handleWhoami}>whoami</button>
      {respText && <p>{respText}</p>}
    </div>
  );
}

export default Home;
