function validateForm() {
    var password = document.getElementById("password").value;
    var repeatPassword = document.getElementById("repeatpassword").value;

    if (password !== repeatPassword) {
        document.getElementById("error-message").innerText = "Les mots de passe ne correspondent pas.";
        return false;
    }
}