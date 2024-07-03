document.getElementById('comment-report-form').addEventListener('submit', function(event) {
    event.preventDefault();
    var form = event.target;
    var action = event.submitter.value;
    var postID = form.querySelector('[name="postid"]').value;

    if (action === 'comment') {
        form.action = '/addcomment?id=' + postID;
    } else if (action === 'report') {
        form.action = '/reportpost?id=' + postID;
    }
    form.submit();
});

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