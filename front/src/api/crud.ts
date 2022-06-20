import { request } from '@/utils/service';
import { AxiosPromise } from 'axios';

export function get_dataset(url: string): AxiosPromise<any> {
    return request({
        url,
        method: "get",
    });
}

export function post_data<T>(url: string, data: T): AxiosPromise<ResponseBody> {
    return request({
        url,
        method: "post",
        data,
    });
}

export function put_data<T>(url: string, data: T): AxiosPromise<ResponseBody> {
    return request({
        url,
        method: "put",
        data,
    });
}

export function delete_data<I>(prefix: string, pkey: I): AxiosPromise<ResponseBody> {
    return request({
        url: `${prefix}/${pkey}`,
        method: "delete",
    });
}
