import { PRODUCTION_CONFIG } from "./config.production";
import { DEV_CONFIG } from "./config.dev";

const { NODE_ENV } = import.meta.env;
let env = DEV_CONFIG;
if (NODE_ENV === "production") env = PRODUCTION_CONFIG;

export default {
  ...env
};
