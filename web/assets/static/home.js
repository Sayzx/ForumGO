document.addEventListener('DOMContentLoaded', function () {
    var menu = document.getElementById('popupMenu');
    var avatarImg = document.querySelector('.avatar-img');

    document.addEventListener('click', function (event) {
        var isClickInsideAvatar = avatarImg.contains(event.target);
        var isClickInsideMenu = menu.contains(event.target);

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
