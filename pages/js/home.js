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
            <a class="imgHover" href="/info/`+ (ele.id) + `" target="_blank">
            <img src="/sprt/`+ getPokemonSprite(ele.id) + `MS.png" class="card-img imgPokemon">
            </a>
            <div class="card-body">
                <div class="pokemonType">
                `+ getPokemonType(ele.type) + `
                </div>
                <div class="divStatus">
                    <p class="pokemonStatus"><span class="iconify" data-icon="mdi:cards-heart-outline" style="color: #ff6e5a;"></span>`+ ele.base.HP + `</p>
                    <p class="pokemonStatus"><span class="iconify" data-icon="mdi:sword" data-rotate="180deg" style="color: #ff6e5a;"></span>`+ ele.base.Attack + `</p>
                    <p class="pokemonStatus"><span class="iconify" data-icon="mdi:shield-outline" style="color: #ff6e5a;"></span>`+ ele.base.Defense + `</p>
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

function getPokemonType(types) {
    var html = "";

    types.forEach(ele => {
        switch (ele) {
            case "Fighting":
                html += `<div class="typeBox fighting">Fighting</div>`
                break;

            case "Flying":
                html += `<div class="typeBox flying">Flying</div>`
                break;

            case "Poison":
                html += `<div class="typeBox poison">Poison</div>`
                break;

            case "Ground":
                html += `<div class="typeBox ground">Ground</div>`
                break;

            case "Rock":
                html += `<div class="typeBox rock">Rock</div>`
                break;

            case "Bug":
                html += `<div class="typeBox bug">Bug</div>`
                break;

            case "Ghost":
                html += `<div class="typeBox ghost">Ghost</div>`
                break;

            case "Steel":
                html += `<div class="typeBox steel">Steel</div>`
                break;

            case "Fire":
                html += `<div class="typeBox fire">Fire</div>`
                break;

            case "Water":
                html += `<div class="typeBox water">Water</div>`
                break;

            case "Grass":
                html += `<div class="typeBox grass">Grass</div>`
                break;

            case "Electric":
                html += `<div class="typeBox electric">Electric</div>`
                break;

            case "Psychic":
                html += `<div class="typeBox psychic">Psychic</div>`
                break;

            case "Ice":
                html += `<div class="typeBox ice">Ice</div>`
                break;

            case "Dragon":
                html += `<div class="typeBox dragon">Dragon</div>`
                break;

            case "Dark":
                html += `<div class="typeBox dark">Dark</div>`
                break;

            case "Fairy":
                html += `<div class="typeBox fairy">Fairy</div>`
                break;

            default:
                html += `<div class="typeBox normal">Normal</div>`
                break;
        }
    })

    return html
}