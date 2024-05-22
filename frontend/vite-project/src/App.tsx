import React, { useState } from 'react';
import './App.css';

interface ShortenedURLResponse {
  shortenedUrl: string;
}

function App(): JSX.Element {
  const [originalUrl, setOriginalUrl] = useState('');
  const [shortenedUrl, setShortenedUrl] = useState('');
  const [error, setError] = useState<string>('');
  const [shake, setShake] = useState(false);

  // Validate URL using regular expression
  const urlRegex = /^(ftp|http|https):\/\/[^ "]+$/


  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (originalUrl.trim() === '') {
      setError('Enter URL');
      setShake(true);
      setTimeout(() => setShake(false), 500);
      return;
    }

    if (!urlRegex.test(originalUrl)) {
      setError('Enter a valid URL')
      setShake(true)
      setTimeout(() => setShake(false), 500)
      return
    }

    try {
      const response = await fetch('http://localhost:8080/shorten', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ url: originalUrl }),
      });

      if (!response.ok) {
        throw new Error('Failed to shorten URL');
      }

      const data: ShortenedURLResponse = await response.json();
      setShortenedUrl(data.shortenedUrl);
      setError('');
    } catch (error: any) {
      setError(error.message);
    }
  };

  return (
    <div>
      <h1>URL Shortener</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="url"
          value={originalUrl}
          onChange={(event) => setOriginalUrl(event.target.value)}
          placeholder="Enter URL to shorten"
          className={shake ? 'shake' : ''}
        />
        <button type="submit">Shorten</button>
      </form>
      {shortenedUrl && <p>Shortened URL: {shortenedUrl}</p>}
      {error && <p className="error">{error}</p>}
    </div>
  );
}

export default App;
