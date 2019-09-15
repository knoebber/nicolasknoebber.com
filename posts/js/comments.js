const savedComments = new Set();

if (postNum) fetchComments()

function fetchComments() {
  fetch('https://l4oejeyzok.execute-api.us-west-2.amazonaws.com/default/get_comments', {
      method: 'POST',
      body: JSON.stringify({ post_number: postNum }),
      headers: {
        'Content-Type':'application/json'
      }
    }).then(response => handleResponse(response))
      .then(data     => handleComments(data))
      .catch(err     => console.log(err));
}

function handleResponse(response) {
  if (response.ok) return response.json();
  throw new Error(response.statusText);
}

function handleComments(comments){
  const commentDiv = document.createElement('div');
  commentDiv.setAttribute('id', 'comment-section');
  commentDiv.innerHTML = `
      <hr/>
      <h2>Comments</h2>
      <form>
        <input id=\"name\" placeholder=\"your name\" maxlength=16 required>
        <textarea id=\"comment-body\" placeholder=\"your comment\" required></textarea>
        <button id=\"submit-button\" type=\"button\" onclick=\"saveComment()\"> Submit </button>
      </form>
      <div id=\"comments\">
      </div>`;
  document
    .getElementById('content')
    .appendChild(commentDiv);

  if (!comments || comments.length === 0) {
    document
      .getElementById('comments')
      .innerHTML("There's nothing here yet")
  } else {
    comments.Items.forEach(c => displayComment(c));
  }
}

function displayComment(comment){
  // Strip the comment of all white space and add it to the comment set.
  savedComments.add(comment.comment_body.S.replace(/\s/g,''));

  const commentSection = document.getElementById('comments');

  const date = new Date(parseInt(comment.time_stamp.N))
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();

  const newComment = document.createElement('div');

  newComment.classList.add("comment");
  newComment.innerHTML = `
    <div class="comment-name">
      <strong>${comment.comment_name.S}<span class="date">${month}/${day}/${year}</span></strong>
    </div>
    <div class="comment-body">
      ${comment.comment_body.S}
    </div>`;

  commentSection.prepend(newComment);
}

function saveComment(){
  const sanitizeString = (s) => {
    const temp = document.createElement('div');
    temp.textContent = s;
    return temp.innerHTML;
  };

  const commentName = sanitizeString(document.getElementById('name').value);
  const commentBody = sanitizeString(document.getElementById('comment-body').value);

  // Check that the comment name/body are not empty, and that the comment body is not a duplicate.
  if (!(commentName || commentBody) || savedComments.has(commentBody.replace(/\s/g,''))) return;

  document.getElementById("submit-button").disabled = true;

  fetch('https://l4oejeyzok.execute-api.us-west-2.amazonaws.com/default/post_comment', {
    method: 'POST',
    body: JSON.stringify({
      postNumber: postNum,
      commentName,
      commentBody,
    }),
    headers: {
      'Content-Type':'application/json'
    }
  }).then(response => response.json())
    .then(data     => displayComment(data))
    .then(()       => document.getElementById("submit-button").disabled = false)
    .catch(err     => console.log(err));
}
