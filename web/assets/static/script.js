document.addEventListener('DOMContentLoaded', function () {
    let menu = document.getElementById('popupMenu');
    let avatarImg = document.querySelector('.avatar-img');

    document.addEventListener('click', function (event) {
<<<<<<< HEAD:web/assets/static/script.js
        let isClickInsideAvatar = avatarImg.contains(event.target);
        let isClickInsideMenu = menu.contains(event.target);
=======
        var isClickInsideAvatar = avatarImg.contains(event.target);
        var isClickInsideMenu = menu.contains(event.target);
>>>>>>> Aylan:web/assets/static/home.js

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

function acceptPost(postId) {
    fetch(`/acceptpost?id=${postId}`, {
        method: 'POST'
    }).then(response => {
        if (response.ok) {
            location.reload();
        } else {
            console.error('Failed to accept post');
        }
    }).catch(error => {
        console.error('Error:', error);
    });
}

function deletePost(postId) {
    fetch(`/deletepost?id=${postId}`, {
        method: 'POST'
    }).then(response => {
        if (response.ok) {
            location.reload();
        } else {
            console.error('Failed to delete post');
        }
    }).catch(error => {
        console.error('Error:', error);
    });
}
