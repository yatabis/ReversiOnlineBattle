const join = () => {
    const gameId = document.getElementById("game-id").value
    fetch(location.href + "?gameid=" + gameId)
        .then(result => result.text())
        .then(text => {
            if (text === gameId) {
                location.href = location.origin + "/play"
            }
        })
}
