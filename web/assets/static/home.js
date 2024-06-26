document.addEventListener('DOMContentLoaded', function () {
    var menu = document.getElementById('popupMenu');
    var avatarImg = document.querySelector('.avatar-img');

    document.addEventListener('click', function (event) {
        // Vérifier si le clic n'est pas sur l'avatar et pas sur le menu
        var isClickInsideAvatar = avatarImg.contains(event.target);
        var isClickInsideMenu = menu.contains(event.target);

        if (!isClickInsideAvatar && !isClickInsideMenu && menu.style.display === 'block') {
            menu.style.display = 'none';
        }
    });

    // Fonction pour basculer la visibilité du menu
    avatarImg.addEventListener('click', function (event) {
        toggleMenu();
        event.stopPropagation(); // Empêcher l'événement de se propager plus loin
    });

    function toggleMenu() {
        if (menu.style.display === 'none' || menu.style.display === '') {
            menu.style.display = 'block';
        } else {
            menu.style.display = 'none';
        }
    }
});
