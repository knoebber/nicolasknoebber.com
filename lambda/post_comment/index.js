exports.handler = (event, context, callback) => {
  const body = event.body;
  const response =
  {
    statusCode: 200,
    body: JSON.stringify(body)
  };
    callback(null, response);
};
