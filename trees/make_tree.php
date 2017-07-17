<?php
error_reporting(E_ALL);
ini_set('display_errors', 1);

$length   = isset($_POST['length'])? $_POST['length']:false;
$angle    = isset($_POST['angle'])? $_POST['angle']:false;
$depth    = isset($_POST['depth'])? $_POST['depth']:false;
$branches = isset($_POST['branches'])? $_POST['branches']:false;
if($length&&$angle&&$depth&&$branches) {//TODO make sure this doesn't let 0s or negatives by
  if($depth * $branches < 30){ //so we don't try to generate something to complex

    $depth    = round($depth);
    $angle    = round($angle);
    $branches = round($branches);
    $length   = round($length);
    $x = 900;
    $y = 900;
    $command = "python /var/www/trees/tree.py $depth $length $angle $branches $x $y";
    var_dump($command);
    $file = shell_exec($command);//TODO fix permissions on python file
    if($file != '') {
      echo json_encode(array('status'=>'success','message'=>'tree was generated','file'=>"$file"));
      }
      else {json_encode(array('status'=>'error','message'=>'there was an error with the script'));}
  }//if
  else {
    echo json_encode(array('status'=>'error','message'=>'tree is too complex'));
  }//else
}//if
else {
    echo json_encode(array('status'=>'error','message'=>'invalid parameters'));
}//else

?>
