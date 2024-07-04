function likePost(postID) {
    fetch('/likepost', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: 'postid=' + postID,
    }).then(response => location.reload());
}

function dislikePost(postID) {
    fetch('/dislikepost', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: 'postid=' + postID,
    }).then(response => location.reload());
}