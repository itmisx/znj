import axios from "axios";
import { Message } from "element-ui";

export default function fetch(options) {
  var instance = axios.create({
    baseURL: domain,
    timeout: 5000,
    withCredentials: true,
  });
  // 登录凭证
  if (localStorage.getItem("Authorization")) {
    axios.defaults.headers.common["Authorization"] = localStorage.getItem(
      "Authorization"
    );
  }
  // 添加请求拦截器
  instance.interceptors.request.use(
    (config) => {
      // 语言
      const lang = localStorage.getItem("lang");
      switch (lang) {
        case "zh":
          axios.defaults.headers.common["Accept-Language"] = "zh";
          break;
        case "en":
          axios.defaults.headers.common["Accept-Language"] = "en";
          break;
        default:
          break;
      }
      return config;
    },
    (error) => {
      return Promise.reject(error);
    }
  );

  // 添加响应拦截器
  instance.interceptors.response.use(
    function(response) {
      if (response.data.code == 200000) {
        let href = window.location.href;
        if (href.indexOf("/login") == -1) {
          window.location.href = "/#login";
          location.reload();
        }
        return response;
      } else {
        return response;
      }
    },
    function(error) {
      let err = error.data ? error.data : error;
      Message({
        message: err,
      });
      return Promise.reject(error);
    }
  );
  return instance(options);
}
