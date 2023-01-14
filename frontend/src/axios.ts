import axios from "axios";
import config from "./config";

const instance = axios.create({
  baseURL: config.BACKEND_HOST,
  timeout: 2000,
});

export default instance;
