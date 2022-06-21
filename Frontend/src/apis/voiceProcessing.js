import axios from "axios";

export const AUTH_URL_PYTHON = "http://127.0.0.1:5000";
export const AUTH_URL_GO = "https://spser.herokuapp.com/api/v1";

export function uploadAudio(path) {
  return axios.post(`${AUTH_URL_PYTHON}/audio_record`, path);
}
export function uploadFirstVoice(path) {
  return axios.post(`${AUTH_URL_PYTHON}/firstcallanalystic`, path);
}
export function UploadData(query) {
  return axios.post(`${AUTH_URL_GO}/call/create`, query);
}
export function getHistoryByPhone(param) {
  return axios.put(`${AUTH_URL_GO}/customer/calls`, param);
}
export function getAllCalls() {
  return axios.get(`${AUTH_URL_GO}/customer/all`);
}
