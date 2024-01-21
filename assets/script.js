function confirmerSoumission() {

    var confirmation = confirm("Êtes-vous sûr de vouloir créer ce personnage ?");

    return confirmation;
}


function displayImage() {
    // Get selected values
    var Sexe = document.getElementById("sexe").value;
    var HairColor = document.getElementById("couleurcheveux").value;
    var Team = document.getElementById("equipe").value;

    // Build the image source based on the selected values
    var imagePath = "static/images/" + Sexe + "_" + Team + "_" + HairColor + ".png";

    // Set the image source
    document.getElementById("characterImage").src = imagePath;
}
