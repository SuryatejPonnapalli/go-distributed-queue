import { axiosInstance } from "./AxiosInstance";

export const getRequest = async (url: string) => {
  try {
    const response = await axiosInstance.get(url);
    return { ok: true, status: response.status, data: response.data };
  } catch (error: any) {
    if (error.response) {
      return {
        ok: false,
        status: error.response.status,
        data: error.response.data,
      };
    }
    return { ok: false, status: 500, data: { message: "Network error" } };
  }
};
