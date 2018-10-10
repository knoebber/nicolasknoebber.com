#!/usr/bin/bash
cd $HOME/code/personal-website/lambda/post_comment
zip -r post_comment.zip index.js
aws lambda update-function-code --function-name post_comment --zip-file fileb://post_comment.zip
rm post_comment.zip
