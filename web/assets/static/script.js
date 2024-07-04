document.addEventListener('DOMContentLoaded', function () {
    let menu = document.getElementById('popupMenu');
    let avatarImg = document.querySelector('.avatar-img');

    document.addEventListener('click', function (event) {
        let isClickInsideAvatar = avatarImg.contains(event.target);
        let isClickInsideMenu = menu.contains(event.target);

        if (!isClickInsideAvatar && !isClickInsideMenu && menu.style.display === 'block') {
            menu.style.display = 'none';
        }
    });

    avatarImg.addEventListener('click', function (event) {
        toggleMenu();
        event.stopPropagation();
    });

    function toggleMenu() {
        if (menu.style.display === 'none' || menu.style.display === '') {
            menu.style.display = 'block';
        } else {
            menu.style.display = 'none';
        }
    }
});
