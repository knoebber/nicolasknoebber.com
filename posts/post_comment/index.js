const AWS = require('aws-sdk');
const dynamo = new AWS.DynamoDB();
const table = "comment";

exports.handler = (event, context, callback) => {
  const {post_number,comment_name,comment_body:comment_body} = JSON.parse(event.body);

  const dynamo_request = {
    Item: {
      "time_stamp": {
        N: new Date().getTime().toString()
      },
      "post_number" : {
        N:post_number.toString()
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
  };

  const respond = (code,message) => {
    callback(null, {
      headers: {
        "Access-Control-Allow-Origin" : "*",
      },
      statusCode: code,
      body: JSON.stringify({
          "status"  : code,
          "message" : message
        })
    });
  };

  dynamo.putItem(dynamo_request, (err, data) => {
    if (!err) respond(200, `added comment for post ${post_number}! name:${comment_name}, body:${comment_body}`);
    else      respond(500,`an error occured ${err}`);
  });
};
