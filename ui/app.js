function calculatePacks() {
    const orderSize = document.getElementById('orderSize').value;
    fetch(`/calculate-packs?orderSize=${orderSize}`)
        .then(response => response.json())
        .then(data => {
            document.getElementById('result').textContent = JSON.stringify(data, null, 2);
        });
}
