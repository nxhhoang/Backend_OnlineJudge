import logo from './logo.svg';
import './App.css';
import { useEffect, useState } from 'react';

const CLIENT_ID = 'Ov23lix6fJrDjUqWNeFC';

function App() {
  // forward user to the github screen (pass in the client id)
  // user in github and login
  // after that, they got forwarded back to localhost:3000
  // but the url will be localhost:3000/?code=
  // use the code to get access token (can only used once)
  const [renderer, setRerender] = useState(false);
  const [userData, setUserData] = useState({});
  useEffect(() => {
    const queryStr = window.location.search;
    const urlParams = new URLSearchParams(queryStr);
    const codeParam = urlParams.get("code");

    console.log(codeParam);

    if (codeParam && (localStorage.getItem('accessToken') === null)) {
      async function getAccessToken() {
        await fetch(`http://localhost:4000/getAccessToken?code=${codeParam}`, {
          method: "GET"
        }).then((response) => {
          return response.json();
        }).then((data) => {
          console.log(data);
          if (data.access_token) {
            localStorage.setItem("accessToken", data.access_token);
            setRerender(!renderer);
          }
        })
      }
      getAccessToken();
    }
  }, []);

  function loginWithGitHub() {
    window.location.assign("https://github.com/login/oauth/authorize?client_id=" + CLIENT_ID);
  }

  async function getUserData() {
    await fetch(`http://localhost:4000/getUserData`, {
      method: "GET",
      headers: {
        "Authorization": "Bearer " + localStorage.getItem("accessToken")
      }
    }).then((response) => {
      return response.json();
    }).then((data) => {
      console.log(data);
      setUserData(data);
    })
  }

  return (
    <div className="App">
      <header className="App-header">
        {localStorage.getItem("accessToken") ?
          <>
            <h1>
              <button onClick={() => {
                localStorage.removeItem('accessToken');
                setRerender(!renderer);
              }}>
                Logut
              </button>
            </h1>
            <h3>
              <button onClick={getUserData} >
                getData
              </button>
              {Object.keys(userData).length !== 0 ?
                <>
                  <h4>
                    Hey there {userData.login}
                  </h4>
                </>
                :
                <>
                </>
              }
            </h3>
          </>
          :
          <>
            <h3>
              User is not logged in
            </h3>
            <button onClick={loginWithGitHub}>
              Login with github
            </button>
          </>
        }
      </header>
    </div>
  );
}

export default App;
