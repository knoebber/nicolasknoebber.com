const AWS = require('aws-sdk');
const dynamo = new AWS.DynamoDB();
const table = "comment";

exports.handler = (event, context, callback) => {
  const body = JSON.parse(event.body)
  const {comment_name,comment_body:comment_body,number} = body
  var code
  var message = "";
  const dynamo_request = {
    Item: {
     "comment_id": {
       S: "test"
      },
     "comment_body": {
       S: comment_body
      },
     "comment_name": {
       S: comment_name
      }
    },
    ReturnConsumedCapacity: "TOTAL",
    TableName: table
  }
  dynamo.putItem(dynamo_request, (err, data) => {
    if (!err){
      code = 200
      message = `adding: comment_name:${comment_name},comment_body:${comment_body},number:${number}\n`
      // TODO use context to exit and set body here inside async request
    }
    else {
      code = 500
      message = `an error occurred: ${err}\n`
    }
  });

  const response = {
    statusCode: code,
    body: JSON.stringify(message)
  }

  callback(null, response)
};

