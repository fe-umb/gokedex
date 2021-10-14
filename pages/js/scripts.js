function makeRequest(url) {
    return new Promise(function (resolve, reject) {
        var xhr = new XMLHttpRequest();
        xhr.open("GET", url);
        xhr.addEventListener('readystatechange', function () {
            if (this.readyState == 4) {
                if (this.status == 200) {
                    resolve(JSON.parse(this.responseText));
                } else {
                    reject(JSON.parse(this.responseText))
                }
            }
        })
        xhr.send();
    });
}

makeRequest("/all")
    .then(
        function (data) {
            addAllPokemons(data);
        }
    );

function addAllPokemons(data) {
    var elemento = "";
    var role = document.querySelector("#cardHolder");
    data.forEach(ele => {
        elemento += `<div class="card">
            <div class="card-header pokemonTitle">
                <p>
                    `+ ele.name.english + ` - #` + ele.id + `
                </p>
            </div>
            <img src="/sprt/`+ getPokemonSprite(ele.id) + `MS.png" class="card-img imgPokemon">
            <div class="card-body">
                <p class="card-text pokemonDescription">Type:</p>
                <div class="divStatus">
                    <p class="pokemonStatus">HP: `+ ele.base.HP + `</p>
                    <p class="pokemonStatus">Attack:`+ ele.base.Attack + `</p>
                    <p class="pokemonStatus">Defense:`+ ele.base.Defense + `</p>
                </div>
            </div>
        </div>`;
    });
    role.innerHTML = elemento;
}

function getPokemonSprite(idPokemon) {
    idPokemon = idPokemon.toString()
    var idReturn = "";

    if (idPokemon.length == 1) {
        idReturn = "00" + idPokemon;
        return idReturn;
    } else if (idPokemon.length == 2) {
        idReturn = "0" + idPokemon;
        return idReturn;
    }

    return idPokemon;

}