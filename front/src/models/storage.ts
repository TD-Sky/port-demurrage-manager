export type GetStorage = {
    id: number,
    store_date: string,
    license_plate_number: string,
    stocks: number,
    store_ton: number,
    duration: number,
}

export type PostStorage = {
    store_date: Date,
    license_plate_number: string,
    stocks: number,
}

export type PutStorage = {
    id: number,
    store_date: Date,
    license_plate_number: string,
    stocks: number,
}
