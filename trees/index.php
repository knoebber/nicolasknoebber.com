<?php
$mostRecentTree = shell_exec('ls *.png -t | head -n1');
$info = array_map(function($e){return substr($e,1);},explode('_',$mostRecentTree));
$initLength   = $info[1];
$initDepth    = $info[2];
$initBranches = $info[3];
$initAngle    = $info[4];
$initWidth    = $info[5][0];
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
</head>
<body>
<h1> Create A Tree </h1>
<ul>
  <li> length
    <input type="range" id="len" value="<?php echo $initLength ?>" min="25" max="200"
    oninput="lenOut.value = len.value">
    <output id="lenOut"><?php echo $initLength ?></output>
  </li>
  <li> branches
    <input type="range" id="branches" value="<?php echo $initBranches?>" min="1" max="15"
    oninput="branchesOut.value = branches.value">
    <output id="branchesOut"><?php echo $initBranches?></output>
  </li>
  <li> angle
    <input type="range" id="angle" value="<?php echo $initAngle?>" min="1" max="360"
    oninput="angleOut.value = angle.value">
    <output id="angleOut"><?php echo $initAngle?></output>
  </li>
  <li> depth
    <input type="range" id="depth" value="<?php echo $initDepth?>" min="1" max="15"
    oninput="depthOut.value = depth.value">
    <output id="depthOut"><?php echo $initDepth?></output>
  </li>
  <li> width
    <input type="range" id="thickness" value="<?php echo $initWidth?>" min="1" max="10"
    oninput="thicknessOut.value = thickness.value">
    <output id="thicknessOut"><?php echo $initWidth?></output>
  </li>
</ul>
<span>
  <button data-action="tree-submit" onclick="requestTree()"> make me a tree! </button>
  <button data-action="pull-random" onclick="requestRandom()"> view random </button>
</span>
<div id="message-container"></div>

<script>
let maxComplex = 30;
let handleResponse = (response) => {
  console.log(response);
  if (response.status =='success') {
    document.getElementById('page')
    .setAttribute('style',`background-image: url("${response.file}");`);
    document.getElementById('message-container')
    .innerHTML= `<p style="color:green;"> success:
                   <span>
                     <a href=${response.file}>file</a>
                   </span>
                 </p>`;

    ids = ['len','depth','branches','angle','thickness'];
    info = response.file.slice(5,-4)
                        .split('_')
                        .map(s=>s.slice(1))
                        .map(e => ({'id':ids.shift(),'val':e}));

    for (i of info) document.getElementById(i.id).value=i.val;
    for (i of info) document.getElementById(i.id+'Out').innerHTML=i.val;
  }//if response.status=='success'

  else {
    document.getElementById('message-container')
    .innerHTML= `<p style="color:red;"> ${response.message} </p>`;
  }
}//handleResponse

let requestTree = () => {
  //TODO refactor
  let length   = document.getElementById('lenOut').value;
  let angle    = document.getElementById('angleOut').value;
  let branches = document.getElementById('branchesOut').value;
  let depth    = document.getElementById('depthOut').value;
  let width    = document.getElementById('thicknessOut').value;
  if (depth * branches > maxComplex) {
    document.getElementById('message-container').innerHTML= `
    <p style="color:red;">
      tree is too complex: depth * branches should be less than or equal to ${maxComplex}
    </p>`;
    return;
  }

  if (length&&angle&&branches&&depth&&width) {
    document.getElementById('message-container').innerHTML=`
    <p style="color:green"> pygame is making your tree...</p>`;
    fetch('request_tree.php', {
      method: 'post',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      body: `length=${length}&angle=${angle}&branches=${branches}&depth=${depth}&width=${width}`

      }).then((response) => response.json())
        .then((json) => handleResponse(json));
    }//if
}//requestTree

let requestRandom = () => {
  fetch('request_tree.php', {
      method: 'post',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      body: 'getRandom=true'

      }).then((response) => response.json())
        .then((json) => handleResponse(json));
}//requestRandom
</script>
</body>
</html>
