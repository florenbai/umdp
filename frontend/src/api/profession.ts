import axios from 'axios';
import qs from 'query-string';
import { HttpResponse } from './interceptor';

export interface ProfessionRecord {
  id: number;
  token: string;
  professionName: string;
  channels: string;
  createdAt: string;
  updatedAt: string;
}

export interface ProfessionForm {
  professionName: string;
  token: string;
  channels: number[];
}

export interface ProfessionChannelRecord {
  id: number;
  channelName: string;
  channelTag: string;
}

export interface ProfessionParams extends Partial<ProfessionRecord> {
  current: number;
  pageSize: number;
}

export interface ProfessionListRes {
  list: ProfessionRecord[];
  total: number;
}

export function queryProfessionList(params: ProfessionParams) {
  return axios.get<ProfessionListRes>('/api/v1/profession/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}

export function submitProfessionForm(data: ProfessionForm) {
  return axios.post<HttpResponse>('/api/v1/profession', data);
}

export function deleteProfession(id: number) {
  return axios.delete(`/api/v1/profession/${id}`);
}

export function getProfessionById(id: string) {
  return axios.get<ProfessionRecord>(`/api/v1/profession/${id}`);
}

export function submitEditProfessionForm(id: string, data: ProfessionForm) {
  return axios.put<HttpResponse>(`/api/v1/profession/${id}`, data);
}

export function getAllProfession() {
  return axios.get<HttpResponse>(`/api/v1/profession/all`);
}

export function getProfessionChannels(id: number) {
  return axios.get<ProfessionChannelRecord[]>(
    `/api/v1/profession/channels/${id}`
  );
}

export function getProfessionChannelsMap(id: number) {
  return axios.get(`/api/v1/profession/channels-map/${id}`);
}
