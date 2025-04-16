import axios from 'axios';
import qs from 'query-string';
import { HttpResponse } from './interceptor';

export interface WechatConfig {
  corpid: string;
  secret: string;
  agentid: string;
  corpsecret: string;
  token?: string;
  encodingAESKey?: string;
}

export interface WechatBotConfig {
  key: string;
}

export interface DingTalkConfig {
  appKey: string;
  appSecret: string;
  agentId: string;
}

export interface FeishuConfig {
  appId: string;
  appSecret: string;
}

export interface YmrtConfig {
  appId: string;
  appSecret: string;
  templateId: string;
}

export interface EmailConfig {
  username: string;
  password: string;
  smtpServer: string;
  smtpPort: string;
}

export interface AliyunSmsConfig {
  accessKeyId: string;
  accessKeySecret: string;
  regionId: string;
  endpoint: string;
}
export interface ChannelRecord {
  id: number;
  channelName: string;
  channelTag: string;
  channelConfig:
  | WechatConfig
  | EmailConfig
  | AliyunSmsConfig
  | DingTalkConfig
  | YmrtConfig;
  channelStatus: 0 | 1;
  createdAt: string;
  updatedAt: string;
}

export interface ChannelForm {
  channelName: string;
  channelTag: string;
  channelConfig: any;
  channelStatus: 0 | 1;
}

export interface ChannelParams extends Partial<ChannelRecord> {
  current: number;
  pageSize: number;
}

export interface ChannelListRes {
  list: ChannelRecord[];
  total: number;
}

export function queryChannelList(params: ChannelParams) {
  return axios.get<ChannelListRes>('/api/v1/channel/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}

export function submitChannelForm(data: ChannelForm) {
  return axios.post<HttpResponse>('/api/v1/channel', data);
}

export function deleteChannel(id: number) {
  return axios.delete(`/api/v1/channel/${id}`);
}

export function getChannelById(id: string) {
  return axios.get<ChannelRecord>(`/api/v1/channel/${id}`);
}

export function submitEditChannelForm(id: string, data: ChannelForm) {
  return axios.put<HttpResponse>(`/api/v1/channel/${id}`, data);
}

export function getAllChannel() {
  return axios.get<ChannelListRes>(`/api/v1/channel/all`);
}
