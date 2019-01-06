const AWS = require('aws-sdk');
const dynamo = new AWS.DynamoDB();
const table = "comment";

exports.handler = (event, context, callback) => {
  const {
    postNumber,
    commentName,
    commentBody
  } = JSON.parse(event.body);

  if (!(postNumber || commentName || commentBody)){
    return respond(callback,400,"postNumber, commentName, commentBody must be non empty");
  }

  const newItem = {
    "time_stamp": {
      N: new Date().getTime().toString()
    },
    "post_number" : {
      N:postNumber.toString()
    },
    "comment_body": {
      S: commentBody
    },
    "comment_name": {
      S: commentName
    }
  };

  const dynamoRequest = {
    Item: newItem,
    ReturnConsumedCapacity: "TOTAL",
    TableName: table
  };

  dynamo.putItem(dynamoRequest, (err, data) => {
    if (!err) respond(callback, 200, newItem);
    else      respond(callback, 500,`an error occured: ${err}`);
  });
};

function respond(callback, code, response){
  let result = {
    statusCode:code,
    body: JSON.stringify(response),
    headers:{
      "Access-Control-Allow-Origin" : "*"
    }
  };

  if (code == 200) result.body = JSON.stringify(response);
  else             result.statusMessage = response;
  callback(null, result);
}
