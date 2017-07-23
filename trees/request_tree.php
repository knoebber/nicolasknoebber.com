<?php
error_reporting(E_ALL);
ini_set('display_errors', 1);

function checkVar($param) {
  return (isset($_POST[$param]) && is_numeric($_POST[$param]) && $_POST[$param] > 0);
}

$requestRandom = isset($_POST['getRandom']);

$length   = checkVar('length')? $_POST['length']:false;
$angle    = checkVar('angle')? $_POST['angle']:false;
$depth    = checkVar('depth')? $_POST['depth']:false;
$branches = checkVar('branches')? $_POST['branches']:false;

if ($requestRandom) {
  $file = exec('ls *.png | shuf | head -n 1');
  echo json_encode(array('status'=>'success','file'=>$file));
}
else {
  if($length&&$angle&&$depth&&$branches) {
    if($depth * $branches <= 30){ //so we don't try to generate something to complex
      $depth    = round($depth);
      $angle    = round($angle);
      $branches = round($branches);
      $length   = round($length);
      $x = 900;
      $y = 900;
      $command = "/var/www/trees/tree.py $depth $length $angle $branches $x $y 2>&1";
      $file = exec($command);
      if($file != '') {
        echo json_encode(array('status'=>'success','message'=>'tree was generated','file'=>"$file"));
      }
      else {
        echo json_encode(array('status'=>'error','message'=>'there was an error with the script'));
      }
    }//if
    else {
      echo json_encode(array('status'=>'error','message'=>'tree is too complex'));
    }
  }//if
  else {
      echo json_encode(array('status'=>'error','message'=>'invalid parameters'));
  }
}//else
?>
