function searchItems() {
  // Récupère la valeur de l'input et la met en minuscule
  const filterValue = document
    .getElementById("searchInput")
    .value.toLowerCase();

  // Sélectionne toutes les cartes
  const cards = document.querySelectorAll(".card");

  // Pour chaque carte
  cards.forEach((card) => {
    // Récupère le texte de la face arrière (ou toute autre zone à filtrer)
    const cardText = card.querySelector(".back").textContent.toLowerCase();

    // Vérifie si le texte inclut la valeur de l'input
    if (cardText.includes(filterValue)) {
      card.style.display = "block";
    } else {
      card.style.display = "none";
    }
  });
}

const recentBike = document.getElementById("recent-bike");

recentBike.recentBike.backgroundColor = "red";