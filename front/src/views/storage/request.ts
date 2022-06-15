import { request } from '@/utils/service';
import { PutStorage, PostStorage } from '@/models/index';
import { AxiosPromise } from 'axios';

export function get_stores(): AxiosPromise<any> {
    return request({
        url: "/store",
        method: "get",
    });
}

export function post_store(data: PostStorage): AxiosPromise<any> {
    return request({
        url: "/store",
        method: "post",
        data,
    });
}

export function put_store(data: PutStorage): AxiosPromise<any> {
    return request({
        url: "/store",
        method: "put",
        data,
    });
}

export function delete_store(id: number): AxiosPromise<any> {
    return request({
        url: `/store/${id}`,
        method: "delete",
    });
}
