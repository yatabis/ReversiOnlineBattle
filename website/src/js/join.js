const join = () => {
    const gameId = document.getElementById("game-id").value
    if (!gameId) {
        document.getElementById("message").innerText = "合戦名を入力してください！"
        return
    }
    fetch(location.href + "?gameid=" + gameId)
        .then(result => result.text())
        .then(text => {
            if (text === gameId) {
                location.href = location.origin + "/play"
            } else {
                document.getElementById("message").innerText = `「${gameId}」という合戦はまだ始まっていません！`
            }
        })
}
