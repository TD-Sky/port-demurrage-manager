import { request } from '@/utils/service';
import { Warehouse } from '@/models/warehouse';
import { AxiosPromise } from 'axios';

export function get_warehouses(): AxiosPromise<any> {
    return request({
        url: "/warehouse",
        method: "get",
    });
}

export function post_warehouse(data: Warehouse): AxiosPromise<any> {
    return request({
        url: "/warehouse",
        method: "post",
        data,
    });
}

export function put_warehouse(data: Warehouse): AxiosPromise<any> {
    return request({
        url: "/warehouse",
        method: "put",
        data,
    });
}

export function delete_warehouse(house_name: string): AxiosPromise<any> {
    return request({
        url: `/warehouse/${house_name}`,
        method: "delete",
    });
}
