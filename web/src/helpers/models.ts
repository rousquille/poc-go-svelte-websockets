
export type wsMessage = {
    type: string,
    payload: any
}

export type chatPayload = {
    user: string,
    message: string,
    timestamp?: number
}
