<script lang="ts">
  import {allChatMessages, castWsMessagePayload, formatChatMessages, storedHistoryChatMessages} from "./helpers/store";
  import type {chatPayload, wsMessage} from "./helpers/models";
  import IconUser from "./Components/IconUser.svelte";
  import IconMessage from "./Components/IconMessage.svelte";

  let socket = new WebSocket("ws://127.0.0.1:9000/ws")
  console.log("attempting connection...")

  socket.onopen = () => {
    console.log("successfully connected")
    // socket.send("hi from svelte !")
  }

  socket.onclose = event => {
    console.log("socket closed connection : ", event)
    // socket.send("client closed !")
  }

  socket.onerror = error => {
    console.log("socket error : ", error)
  }

  let user: string;
  let msgToSend: string;
  $: displayMessages = formatChatMessages($allChatMessages)

  const onSubmit = (event: any) => {
    if (msgToSend != "") {
      const chatMsg: wsMessage = {
        type: "chat",
        payload: {
          user: user,
          message: msgToSend,
        }
      }
      socket.send(JSON.stringify(chatMsg))
    }

    // remove field value after submit
    event.target.reset();
  }

  socket.onmessage = (receivedMsg) => {
    let msgPayload = castWsMessagePayload(receivedMsg.data)

    // dispatch message payload in its store according to its type
    if (<chatPayload>msgPayload) {
      $allChatMessages = [...$allChatMessages, <chatPayload>msgPayload];
      let chatTextarea = document.getElementById("chatTextarea")!;
      chatTextarea.scrollTop = chatTextarea.scrollHeight;
    }
  }

</script>


<main>
  <div class="container text-center">
    <div class="row align-items-start">
      <div class="col-2">
      </div>
      <div class="col-8">
        <h1>Chat Application</h1>

        <div class="card">
          <div class="card-header">Chat</div>
          <div class="card-body">
            <div class="form-floating">
              <textarea id="chatTextarea" class="form-control" style="height:300px;" readonly placeholder="here messages from others will appear" bind:value={displayMessages}></textarea>
            </div>
          </div>
        </div>

        <br>

        <div class="input-group mb-3">
          <span class="input-group-text"><IconUser/></span>
          <input type="text" class="form-control" placeholder="Username" bind:value={user}>
        </div>

        <form id="chatroom-message" on:submit|preventDefault={onSubmit}>
          <div class="input-group mb-3">
            <span class="input-group-text"><IconMessage/></span>
            <input type="text" class="form-control" placeholder="Message" bind:value={msgToSend}>
          </div>
          <button class="btn btn-primary" type="submit">Send message</button>
        </form>
      </div>
      <div class="col-2">
      </div>
    </div>
  </div>

</main>
