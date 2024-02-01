import moment from "moment";

export function timestampToDate(int: number) {
    return moment.unix(int).format("YYYY/MM/DD hh:mm:ss")
}

export function timestampToDateFromString(str: string) {
    return moment.unix(Number(str)).format("YYYY/MM/DD hh:mm:ss")
}
