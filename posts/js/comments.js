const commentAPI = 'https://atmxemymx0.execute-api.us-west-2.amazonaws.com/prod/comments';

// postNum is expected to be included in posts that embed this script.
if (typeof postNum !== 'undefined') fetchComments()

function fetchComments() {
  fetch(`${commentAPI}/${postNum}`)
    .then((response) => parseResponse(response))
    .then((data) => renderComments(data))
    .catch((err) => console.log(err));
}

function parseResponse(response) {
  if (response.ok) return response.json();
  throw new Error(response.statusText);
}

function renderComments(comments){
  const commentDiv = document.createElement('div');
  commentDiv.setAttribute('id', 'comment-section');
  commentDiv.innerHTML = `
      <hr/>
      <h2>Comments</h2>
      <form id="comment-form" onsubmit="postComment(event)">
        <label for="commentName">Name:</label>
        <input name="commentName" maxlength=16 required>
        <label for="commentBody">Comment:</label>
        <textarea name="commentBody" required></textarea>
        <input id="submit-button" type="submit" value="Submit">
      </form>
      <div id="comments">
      </div>`;

  document.getElementById('content').appendChild(commentDiv);
  comments.forEach((c) => renderComment(c));
}

// Strip HTML from user submitted content.
function sanitizeString(s) {
    const temp = document.createElement('div');
    temp.textContent = s;
    return temp.innerHTML;
}

function renderComment({ commentName, commentBody, timestamp }){
  const commentSection = document.getElementById('comments');

  const date = new Date(timestamp);
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();

  const newComment = document.createElement('div');

  newComment.classList.add('comment');
  newComment.innerHTML = `
    <div class="comment-name">
      <strong>${sanitizeString(commentName)}<span class="date">${month}/${day}/${year}</span></strong>
    </div>
    <div class="comment-body">
      ${sanitizeString(commentBody)}
    </div>`;

  commentSection.prepend(newComment);
}

function postComment(e) {
  e.preventDefault();

  const formData = new FormData(e.target);
  const body = { postNumber: postNum };
  [...formData.entries()].forEach(([key, value]) => body[key] = sanitizeString(value));
  document.getElementById('submit-button').disabled = true;
  fetch(commentAPI, {
    method: 'POST',
    body: JSON.stringify(body),
    headers: {
      'Content-Type':'application/json'
    }
  }).then((response) => parseResponse(response))
    .then((data) => renderComment(data))
    .then(() => document.getElementById('submit-button').disabled = false)
    .catch((err) => alert(`falied to post comment: ${err}`))
}
