const ws = new WebSocket("ws://" + location.host + "/wait")
ws.onmessage = ev => {
    location.href = location.origin + "/play"
}
