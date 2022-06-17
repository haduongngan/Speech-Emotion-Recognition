import axios from "axios";

export const AUTH_URL_PYTHON = "http://127.0.0.1:5000";
export const AUTH_URL_GO = "";

export function uploadAudio(path) {
  return axios.post(`${AUTH_URL_PYTHON}/audio_record`, path);
}
