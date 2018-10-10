#!/usr/bin/bash
KEY=$(cat api-key)
curl -X POST -H "x-api-key: $KEY" -H "Content-Type: application/json" -d '{"contents":"lol a comment"}' https://cx00ooxpol.execute-api.us-west-2.amazonaws.com/default/post_comment | python -mjson.tool
