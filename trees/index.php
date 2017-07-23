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
let maxComplex = 30;
let handleResponse = (response)=> {
  if (response.status =='success') {
    document.getElementById('page').setAttribute('style',`background-image: url("${response.file}");`);
    document.getElementById('message-container').innerHTML= `<p style="color:green;"> success: <span><a href=${response.file}>file</a></span></p>`;
  }
  else {
    document.getElementById('message-container').innerHTML= `<p style="color:red;"> ${response.message} </p>`;
  }
}//handleResponse

let requestTree= ()=> {
  let length   = document.getElementById('length').value;
  let angle    = document.getElementById('angle').value;
  let branches = document.getElementById('branches').value;
  let depth    = document.getElementById('depth').value;
  if (depth * branches > maxComplex) {
    document.getElementById('message-container').innerHTML= `<p style="color:red;">tree is too complex: depth * branches should be less than or equal to ${maxComplex}</p>`;
    return;
  }

  if (length&&angle&&branches&&depth) {
    document.getElementById('message-container').innerHTML=`<p style="color:green"> pygame is making your tree...</p>`;
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
