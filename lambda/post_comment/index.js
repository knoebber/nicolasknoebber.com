exports.handler = (event, context, callback) => {
  const body = JSON.parse(event.body)
  const {comment_name,comment_body:comment_body,number} = body;
  const message = `comment_name:${comment_name},comment_body:${comment_body},number:${number}`
  const response = {
    statusCode: 200,
    body: JSON.stringify(message)
  };
  callback(null, response);
};
