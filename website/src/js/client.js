const p2 = Math.PI * 2
const lineWidth     = 2
const cellSize      = 100
const boardSize     = cellSize * 8 + lineWidth * 9
const baseSize      = lineWidth + cellSize
const diskMargin    = 10
const diskRadius    = cellSize / 2 - diskMargin
const suggestMargin = 40
const suggestRadius = cellSize / 2 - suggestMargin
const realSize   = Math.min(window.innerWidth * 0.8, window.innerHeight * 0.7)
const sizeRatio  = boardSize / realSize

class Board {
    constructor(board) {
        this.board = board
        this.ctx = document.getElementById("canvas").getContext("2d")
        this.updateBoard(this.board)
    }

    initBoard() {
        this.ctx.clearRect(0, 0, boardSize, boardSize)

        this.ctx.fillStyle = "green"
        this.ctx.fillRect(0, 0, boardSize, boardSize)

        this.ctx.strokeStyle = "black"
        this.ctx.lineWidth = lineWidth
        this.ctx.beginPath()
        for (let i = 0; i < 9; i ++) {
            this.ctx.moveTo(0, i * baseSize + lineWidth / 2)
            this.ctx.lineTo(boardSize, i * baseSize + lineWidth / 2)
            this.ctx.moveTo(i * baseSize + lineWidth / 2, 0)
            this.ctx.lineTo(i * baseSize + lineWidth / 2, boardSize)
        }
        this.ctx.stroke()
        this.ctx.closePath()
    }

    updateBoard(board) {
        this.board = board
        this.initBoard()
        for (let y = 0; y < 8; y++) {
            for (let x = 0; x < 8; x++) {
                this.putDisk(x, y, this.board[y][x])
            }
        }
    }

    putDisk(x, y, c) {
        let radius
        switch (c) {
            case 1:
                this.ctx.fillStyle = "black"
                radius = diskRadius
                break
            case 2:
                this.ctx.fillStyle = "white"
                radius = diskRadius
                break
            case 3:
                this.ctx.fillStyle = "gray"
                radius = suggestRadius
                break
            default:
                return
        }
        const left = x * baseSize + lineWidth
        const top = y * baseSize + lineWidth
        const centerX = left + cellSize / 2
        const centerY = top + cellSize / 2
        this.ctx.beginPath()
        this.ctx.moveTo(centerX + radius, centerY)
        this.ctx.arc(centerX, centerY, radius, 0, p2, true)
        this.ctx.closePath()
        this.ctx.fill()
    }
}

const canvas = document.getElementById("canvas")

const board = new Board([
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 3, 0, 0, 0, 0],
    [0, 0, 3, 2, 1, 0, 0, 0],
    [0, 0, 0, 1, 2, 3, 0, 0],
    [0, 0, 0, 0, 3, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0],
])

const host = document.getElementById("host").innerText
const gameId = document.getElementById("game-id").innerText
const ws = new WebSocket("ws://" + host + "/open")
ws.onopen = (event) => console.log("connected.", event)
ws.onclose = (event) => console.log("disconnected.", event)
ws.onerror = (event) => console.log("Error: ", event)
ws.onmessage = (event) => {
    const msg = JSON.parse(event.data)
    switch (msg.type) {
        case "board":
            board.updateBoard(msg.data)
            turn = 3 - turn
    }
}

canvas.addEventListener("click", (event) => {
    const data = JSON.stringify({
        gameId: gameId,
        turn: turn,
        point: {
            x: Math.floor((event.clientX - canvas.offsetLeft) / baseSize * sizeRatio),
            y: Math.floor((event.clientY - canvas.offsetTop) / baseSize * sizeRatio)
        }
    })
    console.log("send: ", data)
    ws.send(data)
})

let turn = 1
