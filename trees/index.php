<?php
$mostRecentTree = shell_exec('ls *.png -t | head -n1');
?>
<style>
html {
  background: url(<?php echo $mostRecentTree?>) no-repeat center center fixed;
  background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  -webkit-background-size: cover;
}
h1 {
  font-family: "Lucida Console", Monaco, monospace
}

ol {
  font:20px "Lucida Console",monospace;
  padding-top:100px;
}

p {
  font:20px "Lucida Console",monospace;
}
</style>

<!DOCTYPE html>
<html id="page">
<head>
<meta content="text/html;charset=utf-8" http-equiv="Content-Type">
<meta content="utf-8" http-equiv="encoding">
<link rel="stylesheet" href="style.css">
</head>
<body>
<h1> Create A Tree </h1>
<ul>
<li> length <input id="length" value=0></input> </li>
<li> branches <input id="branches" value=0></input> </li>
<li> angle <input id="angle" value=0></input> </li>
<li> depth <input id="depth" value=0></input> </li>
<button data-action="tree-submit" onclick="requestTree()"> make me a tree! </button>
<div id="message-container"></div>
</ul>

<script>
let handleResponse = (response)=> {
   if (response.status =='success') {
      document.getElementById('page').setAttribute('style',`background-image: url("${response.file}");`);
      document.getElementById('message-container').innerHTML= `<p style="color:green;"> ${response.message} </p>`;
   }
   else {
      document.getElementById('message-container').innerHTML= `<p style="color:red;"> ${response.message} </p>`;
   }
}//handleResponse

let requestTree= ()=> {
  let length   = 0;
  let angle    = 0;
  let branches = 0;
  let depth    = 0;
  length   = document.getElementById('length').value;
  angle    = document.getElementById('angle').value;
  branches = document.getElementById('branches').value;
  depth    = document.getElementById('depth').value;
  console.log(length);
  console.log(angle);
  console.log(branches);
  console.log(depth);

  if (length&&angle&&branches&&depth) {
    fetch('request_tree.php', {
      method: 'post',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      body: `length=${length}&angle=${angle}&branches=${branches}&depth=${depth}`

      }).then((response) => response.json())
        .then((json) => handleResponse(json));
    }//if
}//requestTree
</script>
</body>
</html>
