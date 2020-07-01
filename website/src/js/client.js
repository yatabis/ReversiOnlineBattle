const initCanvas = (ctx) => {
    ctx.fillStyle = "green"
    ctx.fillRect(0, 0, 418, 418)
}

window.addEventListener("load", () => {
    const canvas = document.getElementById("canvas")
    const ctx = canvas.getContext("2d")
    initCanvas(ctx)
})

