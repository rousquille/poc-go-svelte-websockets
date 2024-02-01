import {type Writable, writable} from 'svelte/store';
import type {chatPayload} from "./models";
import {timestampToDate} from "./date";

export const storedHistoryChatMessages = localStorage.getItem("historyChat")

export let allChatMessages: Writable<chatPayload[]>

if (storedHistoryChatMessages === null) {
    allChatMessages = writable([])
} else {
    let oldMessages = JSON.parse(storedHistoryChatMessages) as chatPayload[]
    allChatMessages = writable(oldMessages)
}

allChatMessages.subscribe(value => {
    localStorage.setItem("historyChat", JSON.stringify(value))
    }
)

export function castWsMessagePayload(data: any) {
    const obj = JSON.parse(data)

    if (obj.type === "chat") {
        return obj.payload as chatPayload
    }
}

export function formatChatMessages(data: chatPayload[]): string {
    let formated = data.map(obj => {
        return "| " + timestampToDate(obj.timestamp!) + " | [" + obj.user + "] : " + obj.message; }
    ).join("\n")

    return formated + "\n\n"
}
