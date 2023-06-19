let emailInput = document.querySelector("#emailInput")
let emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

function CheckEmail(e) {
    let email = e.target.value;

    if (!emailRegex.test(email)) {
        console.log("L'adresse email n'est pas conforme !");
        //TODO prevent submission of form
    }
}


if (localStorage.getItem('cookiesAccepted') == true) {
    document.getElementById('popup').style.display = 'block';
}

// Fonction appelée lorsque l'utilisateur accepte les cookies
function accepterCookies() {
    localStorage.setItem('cookiesAccepted', "true");
    document.getElementById('popup').style.display = 'none';
}

function verifierMotDePasse() {
    var password = document.getElementById('passwordInput').value;
    var passwordRegex = /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/;

    if (passwordRegex.test(password)) {
        alert("Le mot de passe est conforme !");
    } else {
        alert("Le mot de passe n'est pas conforme !");
    }
}
emailInput.addEventListener("change", CheckEmail)
