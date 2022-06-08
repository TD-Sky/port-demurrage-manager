import { request } from '@/utils/service';
import { GetStorage, PutStorage, PostStorage } from '@/models/index';
import { Ref } from "vue";

export function get_stores(receiver: Ref<GetStorage[]>) {
    request({
        url: "/store",
        method: "get",
    }).then((resp) => {
        receiver.value = resp.data.stores
    })
}

export function post_store(data: PostStorage) {
    return request({
        url: "/store",
        method: "post",
        data,
    })
}

export function put_store(data: PutStorage) {
    return request({
        url: "/store",
        method: "put",
        data,
    })
}

export function delete_store(id: number) {
    return request({
        url: `/store/${id}`,
        method: "delete",
    })
}
