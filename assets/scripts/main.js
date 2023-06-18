function verifierEmail() {
    var email = document.getElementById('emailInput').value;
    var emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    if (emailRegex.test(email)) {
        alert("L'adresse email est conforme !");
    } else {
        alert("L'adresse email n'est pas conforme !");
    }
}

if (!localStorage.getItem('cookiesAccepted')) {
    document.getElementById('popup').style.display = 'block';
}

// Fonction appelée lorsque l'utilisateur accepte les cookies
function accepterCookies() {
    localStorage.setItem('cookiesAccepted', true);
    document.getElementById('popup').style.display = 'none';
}