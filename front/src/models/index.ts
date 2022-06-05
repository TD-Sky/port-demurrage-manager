export type GetStorage = {
    id: number,
    store_date: string,
    license_plate_number: string,
    stocks: number,
    store_ton: number,
    duration: number,
    fee: number,
}

export type PutStorage = {
    id: number,
    store_date: string,
    license_plate_number: string,
    stocks: number,
    store_ton: number,
}

export type GetLoad = {
    id: number,
    order_number: string,
    load_date: string,
    loads: number,
    load_ton: number,
    business_number: string,
    lading_bill_number: string,
}

export type PutLoad = {
    id: number,
    load_date: string,
    loads: number,
    load_ton: number,
}
