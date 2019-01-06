const AWS = require('aws-sdk');
const dynamo = new AWS.DynamoDB();
const table = "comment";

exports.handler = (event, context, callback) => {
  const {
    postNumber,
    commentName,
    commentBody
  } = JSON.parse(event.body);

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

function respond(callback, code, body){
  const response = {
    statusCode:code,
    headers:{
      "Access-Control-Allow-Origin" : "*"
    }
  };

  if (code == 200) response.body = JSON.stringify(body);
  else             response.body = JSON.stringify({'message':body});
  callback(null, response);
}
