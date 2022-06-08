import { request } from '@/utils/service';
import { GetLoad, PutLoad, PostLoad } from '@/models/index';
import { Ref } from "vue";

export function get_loads(receiver: Ref<GetLoad[]>) {
    request({
        url: "/load",
        method: "get",
    }).then((resp) => {
        receiver.value = resp.data.loads
    })
}

export function put_load(data: PutLoad) {
    return request({
        url: "/load",
        method: "put",
        data,
    })
}

export function post_loads(loads: PostLoad[]) {
    return request({
        url: "/load",
        method: "post",
        data: {
            loads,
        },
    })
}

export function delete_load(id: number) {
    return request({
        url: `/load/${id}`,
        method: "delete",
    })
}
