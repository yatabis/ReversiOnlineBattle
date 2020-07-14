const gameId = document.cookie.split("; ").filter(s => s.startsWith("GameID"))[0].split("=")[1]
const ws = new WebSocket("ws://" + location.host + "/wait")
ws.onopen = ev => ws.send(gameId)
ws.onmessage = ev => location.href = location.origin + "/play"
