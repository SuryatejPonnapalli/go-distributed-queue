import { axiosInstance } from "./AxiosInstance";

export const postRequest = async (url: string, body: any) => {
  try {
    const response = await axiosInstance.post(url, body);
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
