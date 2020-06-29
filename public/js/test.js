function makeAjaxCall() {
    fetch('/api/fruits')
    .then(function(response) {
        return response.text();
    })
    .then(function(text) {
        alert(`we called out to the api via ajax and got this response => ${text}`);
    });
}