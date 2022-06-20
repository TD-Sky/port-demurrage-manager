export type GetLoad = {
    id: number,
    order_number: string,
    load_date: string,
    loads: number,
    load_ton: number,
    business_number: string,
    lading_bill_number: string,
    fee: number,
}

export type PostLoad = {
    load_date: Date,
    loads: number,
}

export type PutLoad = {
    id: number,
    load_date: Date,
    loads: number,
}

export type PostShippingOrder = {
    company_code: string,
    loads: PostLoad[],
}
