const AWS = require('aws-sdk');
const dynamo = new AWS.DynamoDB();
const table = "comment";

exports.handler = (event, context, callback) => {
  //TODO get comment by primary key
  /*
  const post_number = JSON.parse(event.body).post_number;
  const dynamo_request = {
    Item: {
      "time_stamp": {
        S: now.toString()
      },
      "post_number" : {
        S:post_number.toString()
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
      statusCode: code,
      body: JSON.stringify(message)
    });
  };
  dynamo.putItem(dynamo_request, (err, data) => {
    if (!err)
      respond(200,`name:${comment_name},body:${comment_body},for post:${post_number}`);
    else
      respond(500,`an error occured ${err}`);
  });
  */
};
