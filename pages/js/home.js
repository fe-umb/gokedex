makeRequest("/all")
    .then(
        function (data) {
            addAllPokemons(data);
        }
    );

FilterPokemon();

function addAllPokemons(data) {
    var elemento = "";
    var role = document.querySelector("#cardHolder");
    data.forEach(ele => {
        elemento += `<div class="card" data-name="` + ele.name.english + `">
            <div class="card-header pokemonTitle">
                <p>
                    `+ ele.name.english + ` - #` + ele.id + `
                </p>
            </div>
            <a class="imgHover" href="/info/`+ (ele.id) + `" target="_blank">
            <img src="/sprt/`+ getPokemonImage(ele.id) + `MS.png" class="card-img imgPokemon">
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

function FilterPokemon() {
    var searchBox = document.querySelector("#searchBox");
    searchBox.addEventListener('keyup', function () {
        var text = this.value;
        var allCards = document.querySelectorAll("[data-name]");

        allCards.forEach(ele => {
            if (ele.getAttribute("data-name").toLowerCase().includes(text.toLowerCase())) {
                ShowPokemon(ele);
            }
            else {
                HidePokemon(ele);
            }
        })
    });
}

function ShowPokemon(element) {
    element.style.display = "flex";
}

function HidePokemon(element) {
    element.style.display = "none";
}