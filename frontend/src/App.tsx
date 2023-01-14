import React, { useEffect, useState } from "react";
import axios from "./axios";
import "./App.css";
import Header from "./components/header";
import getTokenHeader from "./utils/token";
import Feed from "./components/feed";

function App() {
  const [isAuth, setIsAuth] = useState(false);

  useEffect(() => {
    getCurrentUser();
  }, []);

  const getCurrentUser = () => {
    return axios
      .get("/user", {
        headers: {
          Authorization: getTokenHeader(),
        },
      })
      .then((res) => {
        console.log(res);
        setIsAuth(true);
      })
      .catch((err) => console.log(err));
  };

  return (
    <div className="App">
      <Header isAuth={isAuth} onGetCurrentUser={getCurrentUser} />
      <Feed isAuth={isAuth} />
    </div>
  );
}

export default App;
