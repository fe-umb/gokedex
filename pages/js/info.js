var idPokemon = window.location.href;
idPokemon = idPokemon.split("info/")
idPokemon = idPokemon[1];
var urlAPI = "/pokemon/" + idPokemon;

makeRequest(urlAPI)
    .then(
        function (data) {
            // General info
            document.getElementById("titlePokemon").innerText = data.name.english + " #" + data.id;
            document.getElementById("imgPokemon").src = "/imgs/" + getPokemonImage(data.id) + ".png";
            document.getElementById("typeDetail").innerHTML = getPokemonType(data.type);

            // Status
            document.getElementById("englishName").innerText = data.name.english;
            document.getElementById("japaneseName").innerText = data.name.japanese;
            document.getElementById("chineseName").innerText = data.name.chinese;
            document.getElementById("frenchName").innerText = data.name.french;
            document.getElementById("hp").innerText = data.base.HP;
            document.getElementById("attack").innerText = data.base.Attack;
            document.getElementById("defense").innerText = data.base.Defense;
            document.getElementById("spAttack").innerText = data.base["Sp. Attack"];
            document.getElementById("spDefense").innerText = data.base["Sp. Defense"];
            document.getElementById("speed").innerText = data.base.Speed;
        }
    );