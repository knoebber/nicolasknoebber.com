const AWS = require('aws-sdk');
const dynamo = new AWS.DynamoDB();
const table = "comment";

exports.handler = (event, context, callback) => {
  const respond = (code,response) => callback(null,
    {
      statusCode:code,
      body: JSON.stringify(response),
      headers:{
        "Access-Control-Allow-Origin" : "*"
      }
    }
  )

  const {
    postNumber,
    commentName,
    commentBody
  } = JSON.parse(event.body);

  if (!( (postNumber === 0 || postNumber) && commentName && commentBody)){
    respond(400,"postNumber, commentName, commentBody must be non empty");
    return;
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
    if (!err) respond(200, newItem);
    else      respond(500,`an error occured: ${err}`);
  });
};
