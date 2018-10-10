#!/usr/bin/bash
KEY=$(cat api-key)
curl -X POST -H "x-api-key: $KEY" -H "Content-Type: application/json" -d '{"contents":"lol a comment"}' https://5p33qip6cg.execute-api.us-east-1.amazonaws.com/default/savePost
