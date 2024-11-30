// Placeholder for JavaScript file.

document.addEventListener("DOMContentLoaded", function() {
    // Sélection des éléments HTML
    const hamburgerMenu = document.getElementById("hamburger-menu");
    const navbarMenu = document.getElementById("navbar-menu");

    // Vérifie si les éléments existent
    if (hamburgerMenu && navbarMenu) {
        // Ajoute un événement au clic sur le bouton hamburger
        hamburgerMenu.addEventListener("click", function() {
            // Bascule la classe "open" pour afficher ou masquer le menu
            navbarMenu.classList.toggle("open");

            // Met à jour l'attribut aria-expanded pour l'accessibilité
            const isOpen = navbarMenu.classList.contains("open");
            hamburgerMenu.setAttribute("aria-expanded", isOpen);
        });
    }
});
