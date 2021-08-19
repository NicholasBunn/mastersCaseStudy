const {
  LoginRequest,
  LoginResponse,
} = require("./../../protoFiles/javaScript/webGateway/v1/web_gateway_api_v1_pb.js");
const {
  LoginServiceClient,
} = require("./../../protoFiles/javaScript/webGateway/v1/web_gateway_api_v1_grpc_web_pb.js");

window.LogMeIn = function () {
  var username = document.getElementById("username").value;
  var password = document.getElementById("password").value;

  console.log("Username is: ", username, " and password is: ", password);
  var loginService = new LoginServiceClient("http://localhost:8080");

  var request = new LoginRequest();

  request.setUsername(username);
  request.setPassword(password);

  loginService.login(request, {}, function (err, res) {
    console.log("error:", err, " response:", res);

    // if (err) {
    //   console.log(err);
    //   console.log(err.code);
    //   console.log(err.message);
    // }
    // console.log(response.permissions);
  });
};

function LogMeOut() {
  console.log("Not implemented yet");
}

function LoadHomePage() {
  console.log("Not implemented yet");
}

function LoadRouteAnalysisService() {
  console.log("Not implemented yet");
}

function LoadPowerTrainService() {
  console.log("Not implemented yet");
}

function LoadVesselMotionService() {
  console.log("Not implemented yet");
}

function LoadComfortService() {
  console.log("Not implemented yet");
}
