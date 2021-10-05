const RetryInterceptor = function () {
  this.intercept = function (request, invoker) {
    console.log("Started interceptor");
    while (!response) {
      var response = invoker(request);
      console.log("Response: ", response);
    }
    return response;
  };
};
