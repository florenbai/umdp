import axios from 'axios';
import qs from 'query-string';

export interface LogRecord {
  id: number;
  channelName: string;
  professionName: string;
  parameters: string;
  templateName: string;
  errMessage: string;
  createdAt: string;
  status: 0 | 1;
  updatedAt: string;
}

export interface LogParams extends Partial<LogRecord> {
  current: number;
  pageSize: number;
}

export interface LogListRes {
  list: LogRecord[];
  total: number;
}

export function queryLogList(params: LogParams) {
  return axios.get<LogListRes>('/api/v1/log/list', {
    // method: 'GET',
    params,
    paramsSerializer: (obj) => {
      // return qs.stringify(obj, { arrayFormat: 'bracket' });
      return qs.stringify(obj);
      // return obj;
    },
  });
}
