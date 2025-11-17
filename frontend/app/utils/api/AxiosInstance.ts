import axios from "axios";

const backendUrl = import.meta.env.VITE_BACKEND_URL;

export const axiosInstance = axios.create({
  baseURL: backendUrl,
  withCredentials: true,
});

const authExcluded = ["/users/login", "/users/register"];

axiosInstance.interceptors.response.use(
  (res) => res,
  (error) => {
    const url = error?.config?.url;
    const status = error?.response?.status;

    if (authExcluded.some((route) => url?.includes(route))) {
      return Promise.reject(error);
    }

    if (status === 401) {
      window.location.href = "/login";
    }

    return Promise.reject(error);
  }
);
