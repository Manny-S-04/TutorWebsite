function updateStars(input) {
    const stars = document.querySelectorAll('.star');
    const ratingValue = parseInt(input.value);

    stars.forEach((star, index) => {
        if (index < ratingValue) {
            star.style.color = 'gold';  // Fill the selected stars with gold
        } else {
            star.style.color = 'grey';  // Grey out the unselected stars
        }
    });
}
