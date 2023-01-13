import React, { useState } from "react";
import reactLogo from "./assets/react.svg";
import "./App.css";

function App() {
  const [active, setActive] = useState(false);

  const onToggleActive = (e: React.MouseEvent) => {
    e.preventDefault();
    setActive(!active);
  };
  return (
    <div className="App">
      <div className="toggle">
        <a onClick={onToggleActive}>
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <div className={`menu ${active ? "active" : ""}`}>
        <ul>
          <li>
            <a href="#">Home</a>
          </li>
          <li>
            <a href="#">About</a>
          </li>
          <li>
            <a href="#">Contact Us</a>
          </li>
        </ul>
      </div>
    </div>
  );
}

export default App;
