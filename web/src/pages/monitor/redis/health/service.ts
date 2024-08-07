import { request } from '@@/plugin-request/request';
import { stringify } from 'qs';

export async function queryHealthList(params?: string) {
  return request(`/api/v1/performance/redis/health?${stringify(params)}`);
}
