const playerId = document.cookie.split("; ").filter(s => s.startsWith("PlayerID"))[0].split("=")[1]
const ws = new WebSocket("ws://" + location.host + "/wait")
ws.onopen = ev => ws.send(playerId)
ws.onmessage = ev => location.href = location.origin + "/play"
