
document.addEventListener("DOMContentLoaded", () => {
    const scrollElements = document.querySelectorAll(".scroll-reveal");

    const elementInView = (el, offset = 100) => {
        const elementTop = el.getBoundingClientRect().top;
        return elementTop <= (window.innerHeight || document.documentElement.clientHeight) - offset;
    };

    const displayScrollElement = (el) => {
        el.classList.add("visible");
    };

    const hideScrollElement = (el) => {
        el.classList.remove("visible");
    };

    const handleScrollAnimation = () => {
        scrollElements.forEach((el) => {
            if (elementInView(el)) {
                displayScrollElement(el);
            } else {
                hideScrollElement(el);
            }
        });
    };

    window.addEventListener("scroll", () => {
        handleScrollAnimation();
    });

    // Initial call
    handleScrollAnimation();
});

document.addEventListener('DOMContentLoaded', function () {
    const cartButton = document.querySelector('.cart-button');
    const cartDropdown = document.querySelector('.cart-dropdown');
    const closeCartBtn = document.querySelector('.close-cart-btn');

    // Ouvrir le menu panier lorsque l'on clique sur le bouton du panier
    cartButton.addEventListener('click', function () {
        cartDropdown.classList.add('active');
    });

    // Fermer le menu panier lorsque l'on clique sur le bouton "X"
    closeCartBtn.addEventListener('click', function () {
        cartDropdown.classList.remove('active');
    });
});
