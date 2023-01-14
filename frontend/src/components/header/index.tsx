import React, { useState } from "react";
import { message } from "antd";
import type { RcFile } from "antd/es/upload/interface";

import reactLogo from "../../assets/react.svg";
import RegisterForm from "../forms/register-form";
import axios from "../../axios";
import "./index.css";
import LoginForm from "../forms/login-form";
import DogForm from "../forms/dog-form";
import getTokenHeader from "../../utils/token";

const uploadFile = (file: any, signedRequest: string) => {
  return axios
    .put(signedRequest, file, {
      headers: {
        "Content-Type": `image/jpeg`,
      },
    })
    .then((res) => {
      if (res.status === 200)
        return {
          success: true,
          statusCode: res.status,
          message: res.statusText,
        };
      return {
        success: false,
        statusCode: res.status,
        message: res.statusText,
      };
    })
    .catch((err) => {
      message.error(`Could not upload ${file.name}`);
      return {
        success: false,
        statusCode: err.status || 500,
        message: err || "Unknown Error",
      };
    });
};

const Header: React.FC<{ isAuth: boolean, onGetCurrentUser: Function }> = ({ isAuth, onGetCurrentUser }) => {
  const [active, setActive] = useState(false);
  const [register, setRegister] = useState(false);
  const [login, setLogin] = useState(false);
  const [newPost, setNewPost] = useState(false);

  const onToggleActive = (e: React.MouseEvent) => {
    e.preventDefault();
    setActive(!active);
  };

  const registerForm = (
    <RegisterForm
      visible={register}
      onCancel={() => setRegister(false)}
      onRegister={(values: any, cb: any) => {
        axios
          .post("/signin", values)
          .then((res) => {
            message.success("Register successfully, you can login now");
            setRegister(false);
          })
          .catch((err) => console.log(err.response))
          .finally(() => cb());
      }}
      onFinishFailed={() => console.log("finish failed")}
    />
  );

  const loginForm = (
    <LoginForm
      visible={login}
      onCancel={() => setLogin(false)}
      onLogin={(values: any, cb: any) => {
        axios
          .post("/login", values)
          .then((res) => {
            sessionStorage.setItem("token", res.data.access_token);
            message.success("Login successfully");
            setLogin(false);
            onGetCurrentUser();
          })
          .catch((err) => console.log(err.response))
          .finally(() => cb());
      }}
      onFinishFailed={() => console.log("finish failed")}
    />
  );

  const newPostForm = (
    <DogForm
      visible={newPost}
      onCancel={() => setNewPost(false)}
      onNewDogPost={(values: any, cb: any) => {
        const splitImageName = values.image?.file?.name.split(".");
        const image_type = splitImageName[splitImageName.length - 1];
        const serverPostValues = {
          birth_year: parseInt(values.birth_year["$y"]),
          breed: values.breed,
          name: values.name,
          message: values.message,
          image_type,
        };
        axios
          .post("/dog", serverPostValues, {
            headers: {
              Authorization: getTokenHeader(),
            },
          })
          .then((res) => {
            const { dog, upload_url } = res.data;
            const file = values.imageBytes;

            uploadFile(
              file,
              upload_url
            ).then((res) => {
              if (res.success) {
                message.success("Create post successfully");
                setNewPost(false);
              } else {
                axios.delete(`/dog/${dog.id}`, {
                  headers: {
                    Authorization: getTokenHeader(),
                  },
                });
                message.error("Create post failed");
              }
            });
          })
          .catch((err) => console.log(err.response))
          .finally(() => cb());
      }}
      onFinishFailed={() => console.log("finish failed")}
    />
  );
  return (
    <>
      <div className="toggle">
        <a onClick={onToggleActive}>
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <div className="right">
        {!isAuth ? (
          <>
            <a onClick={() => setRegister(!register)}>Register</a>
            <a onClick={() => setLogin(!login)}>Login</a>
          </>
        ) : (
          <a onClick={() => setNewPost(!newPost)}>Post</a>
        )}
        {registerForm}
        {loginForm}
        {newPostForm}
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
    </>
  );
};

export default Header;
