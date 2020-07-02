const p2 = Math.PI * 2
const lineWidth  = 2
const cellSize   = 100
const boardSize  = cellSize * 8 + lineWidth * 9
const baseSize   = lineWidth + cellSize
const diskMargin = 10
const diskRadius = (cellSize - diskMargin) / 2

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
        switch (c) {
            case 1:
                this.ctx.fillStyle = "black"
                break
            case 2:
                this.ctx.fillStyle = "white"
                break
            default:
                return
        }
        const left = x * baseSize + lineWidth
        const top = y * baseSize + lineWidth
        const centerX = left + cellSize / 2
        const centerY = top + cellSize / 2
        this.ctx.beginPath()
        this.ctx.moveTo(centerX + diskRadius, centerY)
        this.ctx.arc(centerX, centerY, diskRadius, 0, p2, true)
        this.ctx.closePath()
        this.ctx.fill()
    }
}

const board = new Board([
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 2, 1, 0, 0, 0],
    [0, 0, 0, 1, 2, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0],
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
    board.updateBoard(JSON.parse(event.data))
    turn = 3 - turn
}

let turn = 1

const put = (x, y) => {
    const data = JSON.stringify({
        gameId: gameId,
        turn: turn,
        point: {
            x: x - 1,
            y: y - 1
        }
    })
    console.log("send: ", data)
    ws.send(data)
}
