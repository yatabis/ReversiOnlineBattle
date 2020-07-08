const p2 = Math.PI * 2
const lineWidth     = 2
const cellSize      = 100
const boardSize     = cellSize * 8 + lineWidth * 9
const baseSize      = lineWidth + cellSize
const diskMargin    = 10
const diskRadius    = cellSize / 2 - diskMargin
const suggestMargin = 40
const suggestRadius = cellSize / 2 - suggestMargin
const realSize      = Math.min(window.innerWidth * 0.8, window.innerHeight * 0.7)
const sizeRatio     = boardSize / realSize
const markWidth     = 6

class Board {
    constructor(board) {
        this.board = board
        this.ctx = document.getElementById("board").getContext("2d")
        this.score = document.getElementById("score").getContext("2d")
        this.updateBoard(this.board)
        this.updateScores()
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
        this.circle(this.ctx, centerX, centerY, radius)
    }

    circle(context, x, y, r) {
        context.beginPath()
        context.moveTo(x + r, y)
        context.arc(x, y, r, 0, p2, true)
        context.closePath()
        context.fill()
    }

    initScore() {
        this.score.clearRect(0, 0, boardSize, baseSize)
        this.score.fillStyle = "skyblue"
        this.score.fillRect(0, 0, boardSize, baseSize)
        this.score.strokeStyle = "black"
        this.score.lineWidth = lineWidth
        this.score.beginPath()
        this.score.moveTo(lineWidth / 2, lineWidth / 2)
        this.score.lineTo(boardSize - lineWidth / 2, lineWidth / 2)
        this.score.lineTo(boardSize - lineWidth / 2, baseSize + lineWidth / 2)
        this.score.lineTo(lineWidth / 2, baseSize + lineWidth / 2)
        this.score.closePath()
        this.score.stroke()
    }

    updateScores() {
        const flat = this.board.flat()
        const black = flat.filter(x => x === 1).length
        const white = flat.filter(x => x === 2).length

        this.initScore()
        const centerY = lineWidth + cellSize / 2
        const left = lineWidth + cellSize
        const right = boardSize - lineWidth - cellSize
        this.score.fillStyle = "black"
        this.circle(this.score, left, centerY, diskRadius)
        this.score.fillStyle = "white"
        this.circle(this.score, right, centerY, diskRadius)

        this.score.font = "48px sans-serif"
        this.score.textBaseline = "middle"
        this.score.fillStyle = "black"
        this.score.fillText(black.toString(), baseSize * 1.5 + cellSize / 2, cellSize / 2 + lineWidth)
        this.score.textAlign = "end"
        this.score.fillText(white.toString(), boardSize - baseSize * 1.5 - cellSize / 2, cellSize / 2 + lineWidth)
        this.score.textAlign = "center"
        this.score.fillText("-", boardSize / 2, cellSize / 2 + lineWidth)
        this.turnMark()
    }

    turnMark() {
        const x1 = lineWidth + markWidth / 2 + cellSize / 2
        const x2 = baseSize - markWidth / 2 + cellSize / 2
        const x3 = boardSize - baseSize - markWidth / 2 - cellSize / 2
        const x4 = boardSize - markWidth / 2 - cellSize / 2
        const y1 = lineWidth + markWidth / 2
        const y2 = baseSize - markWidth / 2
        this.score.strokeStyle = "red"
        this.score.lineWidth = markWidth
        this.score.beginPath()
        if (turn === 1) {
            this.score.moveTo(x1, y1)
            this.score.lineTo(x2, y1)
            this.score.lineTo(x2, y2)
            this.score.lineTo(x1, y2)
        } else if (turn === 2) {
            this.score.moveTo(x3, y1)
            this.score.lineTo(x4, y1)
            this.score.lineTo(x4, y2)
            this.score.lineTo(x3, y2)
        }
        this.score.closePath()
        this.score.stroke()
    }
}

const canvas = document.getElementById("board")

let turn = 1

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
ws.onopen = (event) => {
    console.log("connected.", event)
    ws.send(gameId)
}
ws.onclose = (event) => console.log("disconnected.", event)
ws.onerror = (event) => console.log("Error: ", event)
ws.onmessage = (event) => {
    const msg = JSON.parse(event.data)
    switch (msg.type) {
        case "board":
            board.updateBoard(msg.data)
            break
        case "not_your_turn":
            console.log("Not your turn.")
            break
        case "invalid_put":
            console.log("Invalid put.")
            break
        case "turn_change":
            turn = msg.data
            board.updateScores()
            console.log("_黒白"[turn] + "の番です。")
            break
        case "turn_pass":
            turn = msg.data
            board.updateScores()
            console.log("_白黒"[turn] + "はパスです。")
            break
        case "game_end":
            board.updateScores()
            console.log("ゲーム終了")
            break
        default:
            console.log("receive unknown message.")
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
