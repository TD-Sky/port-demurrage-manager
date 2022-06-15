import { request } from '@/utils/service';
import { PutLoad, PostLoad } from '@/models/index';
import { AxiosPromise } from 'axios';

export function get_loads(): AxiosPromise<any> {
    return request({
        url: "/load",
        method: "get",
    });
}

export function put_load(data: PutLoad): AxiosPromise<any> {
    return request({
        url: "/load",
        method: "put",
        data,
    });
}

export function post_loads(loads: PostLoad[]): AxiosPromise<any> {
    return request({
        url: "/load",
        method: "post",
        data: {
            loads,
        },
    });
}

export function delete_load(id: number): AxiosPromise<any> {
    return request({
        url: `/load/${id}`,
        method: "delete",
    });
}
