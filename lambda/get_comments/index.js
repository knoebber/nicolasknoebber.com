const AWS = require('aws-sdk');
const dynamo = new AWS.DynamoDB();
const table = "comment";

exports.handler = (event, context, callback) => {
  const dynamoRequest = {
    ExpressionAttributeValues: {
      ":v1": { N: JSON.parse(event.body).post_number.toString() }
    },
    ScanIndexForward: true,
    KeyConditionExpression: "post_number = :v1",
    TableName: table
  };

  const respond = (code,message) => {
    callback(null, {
      headers: {
        "Access-Control-Allow-Origin" : "*",
      },
      statusCode: code,
      body: JSON.stringify(message)
    });
  };

  dynamo.query(dynamoRequest, function(err, data) {
    if (!err) respond(200,data)
    else      respond(500, `an error occured: ${err.stack}`)
  })
};
