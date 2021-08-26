const {
  LoginAuthRequest,
  LoginAuthResponse,
} = require("./../../protoFiles/javaScript/authenticationService/v1/authentication_service_api_v1_pb.js");
const {
  AuthenticationServiceClient,
} = require("./../../protoFiles/javaScript/authenticationService/v1/authentication_service_api_v1_grpc_web_pb.js");
const {
  AnalysisRequest,
  AnalysisResponse,
} = require("./../../protoFiles/javaScript/routeAnalysisAggregator/v1/route_analysis_aggregator_api_v1_pb.js");
const {
  AnalysisServiceClient,
} = require("./../../protoFiles/javaScript/routeAnalysisAggregator/v1/route_analysis_aggregator_api_v1_grpc_web_pb.js");
const {
  PTEstimateRequest,
  PTEstimateResponse,
} = require("./../../protoFiles/javaScript/powerTrainAggregator/v1/power_train_aggregator_api_v1_pb.js");
const {
  PTEstimateServiceClient,
} = require("./../../protoFiles/javaScript/powerTrainAggregator/v1/power_train_aggregator_api_v1_grpc_web_pb.js");
const {
  VMEstimateRequest,
  VMEstimateResponse,
} = require("./../../protoFiles/javaScript/vesselMotionAggregator/v1/vessel_motion_aggregator_api_v1_pb.js");
const {
  VMEstimateServiceClient,
} = require("./../../protoFiles/javaScript/vesselMotionAggregator/v1/vessel_motion_aggregator_api_v1_grpc_web_pb.js");

import Chart from "chart.js/auto";

// ________LOGIN FUNCTIONS________

window.LogMeIn = function (guest) {
  /* This function sends a login request to the web gateway. On a successful response it stores the user's auth token, hides the login div and shows the main interface div (loading in the available services based on the user permission)
   */

  if (guest) {
    var username = "guest";
    var password = "guestPassword";
  } else {
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
  }

  document.getElementById("username").value = "";
  document.getElementById("password").value = "";

  console.log("Username is: ", username, " and password is: ", password);
  var loginService = new AuthenticationServiceClient("http://localhost:8080");

  var request = new LoginAuthRequest();

  request.setUsername(username);
  request.setPassword(password);

  loginService.loginAuth(request, {}, function (err, response) {
    if (err) {
      console.log(err.code);
      console.log(err.message);
    } else {
      console.log(response.getPermissions());
      console.log(response.getAccessToken());

      // Manager object to make things a lil cleaner :)
      var managerObject = {
        currentService: null,
        queryID: null,
        userToken: response.getAccessToken(),
        routeInfo: {
          exists: false,
          unixTime: [],
          latitude: [],
          longitude: [],
          heading: [],
          propellerPitch: [],
          motorSpeed: [],
          sog: [],
        },
      };

      localStorage.setItem("myManager", JSON.stringify(managerObject));

      // Hide login div
      toggleDiv("loginBox");

      // Load main interface div
      toggleDiv("mainInterface");
    }
  });
};

window.LogMeOut = function () {
  localStorage.clear();

  toggleDiv("mainInterface");
  toggleDiv("loginBox");

  // Clear user details
  console.log("Not implemented yet");
};

// ________LANDING PAGES________

window.LoadLoadingPage = function () {
  clearDiv("ServiceContent");

  var innerSpinner = document.createElement("Div");
  innerSpinner.className = "loaderInner";

  var midSpinner = document.createElement("Div");
  midSpinner.className = "loaderMid";

  var outerSpinner = document.createElement("Div");
  outerSpinner.className = "loaderOuter";

  var loadingText = document.createElement("h1");
  loadingText.className = "SimulationHeading";
  loadingText.appendChild(document.createTextNode("Running simulation"));
  document.getElementById("ServiceContent").appendChild(loadingText);
  document.getElementById("ServiceContent").appendChild(innerSpinner);
  document.getElementById("ServiceContent").appendChild(midSpinner);
  document.getElementById("ServiceContent").appendChild(outerSpinner);
};

window.LoadHomePage = function () {
  // Clear service content div
  clearDiv("ServiceContent");

  closeNav();

  document.getElementById("ServiceContent");
  // Add a welcome message
  // Add a subtitle here to guide the user

  // Load available services into blocks with a nice pic

  // Create new div

  console.log("Not implemented yet");
};

window.LoadRouteAnalysisHome = function () {
  // Update manager object
  var managerObject = JSON.parse(localStorage.getItem("myManager"));

  managerObject.currentService = "routeAnalysisService";
  managerObject.queryID = "queryRAS";

  localStorage.setItem("myManager", JSON.stringify(managerObject));

  openNav();

  // Clear service content div
  clearDiv("ServiceContent");

  // Create new div
  var div = document.createElement("Div");
  div.style.width = "95%";
  div.style.height = "95%";

  var heading = document.createElement("h1");
  heading.className = "LandingPageHeading";
  heading.appendChild(
    document.createTextNode("Welcome to the route analysis service")
  );
  var subText = document.createElement("p");
  subText.className = "LandingPageSubText";
  subText.appendChild(
    document.createTextNode(
      "This service will provide you with a summary of the selected route. Please use the navigation bar on the left to analyse a proposed route."
    )
  );

  heading.appendChild(subText);
  div.appendChild(heading);
  document.getElementById("ServiceContent").appendChild(div);
};

window.LoadPowerTrainHome = function () {
  // Update manager object
  var managerObject = JSON.parse(localStorage.getItem("myManager"));

  managerObject.currentService = "powerTrainService";
  managerObject.queryID = "queryPTS";

  localStorage.setItem("myManager", JSON.stringify(managerObject));

  openNav();

  // Clear service content div
  clearDiv("ServiceContent");

  // Create new div
  var div = document.createElement("Div");
  div.style.width = "95%";
  div.style.height = "95%";

  var heading = document.createElement("h1");
  heading.className = "LandingPageHeading";
  heading.appendChild(
    document.createTextNode("Welcome to the power train service")
  );
  var subText = document.createElement("p");
  subText.className = "LandingPageSubText";
  subText.appendChild(
    document.createTextNode(
      "This service will provide you with insight into the power train of your vessel. Please use the navigation bar on the left to select a power train service."
    )
  );

  heading.appendChild(subText);
  div.appendChild(heading);
  document.getElementById("ServiceContent").appendChild(div);
};

window.LoadVesselMotionHome = function () {
  // Update manager object
  var managerObject = JSON.parse(localStorage.getItem("myManager"));

  managerObject.currentService = "vesselMotionService";
  managerObject.queryID = "queryVMS";

  localStorage.setItem("myManager", JSON.stringify(managerObject));

  openNav();

  // Clear service content div
  clearDiv("ServiceContent");

  // Create new div
  var div = document.createElement("Div");
  div.style.width = "95%";
  div.style.height = "95%";

  var heading = document.createElement("h1");
  heading.className = "LandingPageHeading";
  heading.appendChild(
    document.createTextNode("Welcome to the vessel motion service")
  );
  var subText = document.createElement("p");
  subText.className = "LandingPageSubText";
  subText.appendChild(
    document.createTextNode(
      "This service will provide you with insight into the motion of your vessel. Please use the navigation bar on the left to select a motion service."
    )
  );

  heading.appendChild(subText);
  div.appendChild(heading);
  document.getElementById("ServiceContent").appendChild(div);
};

window.LoadComfortHome = function () {
  // Update manager object
  var managerObject = JSON.parse(localStorage.getItem("myManager"));

  managerObject.currentService = "comfortService";
  managerObject.queryID = "queryCS";

  localStorage.setItem("myManager", JSON.stringify(managerObject));

  // Clear service content div
  clearDiv("ServiceContent");

  // Create new div
  var div = document.createElement("Div");
  div.style.width = "95%";
  div.style.height = "95%";

  var heading = document.createElement("h1");
  heading.className = "LandingPageHeading";
  heading.appendChild(
    document.createTextNode("Welcome to the comfort service")
  );
  var subText = document.createElement("p");
  subText.className = "LandingPageSubText";
  subText.appendChild(
    document.createTextNode(
      "This service will provide you with insight into the perceived comfort on board your vessel. Please use the navigation bar on the left to select a comfort service."
    )
  );

  heading.appendChild(subText);
  div.appendChild(heading);
  document.getElementById("ServiceContent").appendChild(div);
};

// ________INPUT PAGES_________

window.LoadRouteInput = function () {
  var managerObject = JSON.parse(localStorage.getItem("myManager"));

  // If a route has already been entered, make the request for that route
  if (managerObject.routeInfo.exists) {
    switch (managerObject.queryID) {
      case "queryRAS":
        console.log("Calling RAS");
        queryRAS();
        break;
      case "queryPTS":
        queryPTS();
        break;
      case "queryVMS":
        queryVMS();
        break;
      case "queryCS":
        queryCS();
        break;
    }
  } else {
    // Clear service content div
    clearDiv("ServiceContent");

    // Create new div
    var div = document.createElement("Div");
    div.id = "loadRoute";
    div.style.width = "95%";
    div.style.height = "95%";

    // Add div to service content div
    document.getElementById("ServiceContent").appendChild(div);

    // Create input elements
    createTextInput("Unix time", "unixTimeIn", "loadRoute");
    createTextInput("Latitude", "latitudeIn", "loadRoute");
    createTextInput("Longitude", "longitudeIn", "loadRoute");
    createTextInput("Heading", "headingIn", "loadRoute");
    createTextInput("Propeller Pitch", "propPitchIn", "loadRoute");
    createTextInput("Motor Speed", "motorSpeedIn", "loadRoute");
    createTextInput("SOG", "sogIn", "loadRoute");

    var managerObject = JSON.parse(localStorage.getItem("myManager"));

    createButtonInput(
      "Query",
      "query" + managerObject.currentService,
      "loadRoute"
    );

    // Add event to Query button
    var sendQueryButton = document.getElementById(
      "query" + managerObject.currentService
    );

    var func = null;

    switch (managerObject.queryID) {
      case "queryRAS":
        func = queryRAS;
        break;
      case "queryPTS":
        func = queryPTS;
        break;
      case "queryVMS":
        func = queryVMS;
        break;
      case "queryCS":
        func = queryCS;
        break;
    }

    sendQueryButton.onclick = func;
  }
};

// ________DISPLAY PAGES_______

window.LoadRouteAnalysisDisplay = function (responseObject) {
  openNav();

  console.log(responseObject);
};

window.LoadPowerTrainDisplay = function (responseObject) {
  openNav();

  // Clear service content div
  clearDiv("ServiceContent");

  // Create new div
  var div = document.createElement("Div");
  div.style.width = "95%";
  div.style.height = "95%";

  var canvas = document.createElement("canvas");
  canvas.style.width = "95%";
  canvas.style.height = "95%";
  canvas.id = "chartID";

  div.appendChild(canvas);

  document.getElementById("ServiceContent").appendChild(div);

  console.log(responseObject);

  var chartData = {
    labels: responseObject.time,
    datasets: [
      {
        data: responseObject.powerEstimate,
        label: "Power Consumption",
        borderColor: "rgb(0, 0, 128)",
        backgroundColor: "rgba(0, 0, 128, 0.5)",

        fill: false,
        order: 1,
        type: "line",
        yAxisID: "y1",
      },
      {
        data: responseObject.costEstimate,
        label: "Cumulative Cost",
        borderColor: "rgb(255, 205, 86)",
        backgroundColor: "rgba(255, 205, 86, 0.5)",

        fill: false,
        order: 2,
        yAxisID: "y2",
      },
    ],
  };

  var myChart = new Chart(canvas, {
    type: "bar",
    data: chartData,
    options: {
      interaction: {
        mode: "index",
        intersect: false,
      },
      title: {
        display: true,
        text: "Cost and Power over Time",
      },
      hover: {
        mode: "nearest",
        intersect: true,
      },
      scales: {
        xAxes: [
          {
            display: true,
            gridLines: {
              display: true,
              color: "black",
            },
            ticks: {
              fontColor: "black",
            },
          },
        ],
        y1: {
          type: "linear",
          display: true,
          position: "left",
          gridLines: {
            display: true,
            color: "black",
          },
          ticks: {
            fontColor: "black",
            callback: function (value, index, values) {
              return value + "MW";
            },
          },
        },
        y2: {
          type: "linear",
          display: true,
          position: "right",
          ticks: {
            fontColor: "black",
            callback: function (value, index, values) {
              return "R" + value;
            },
          },
          grid: {
            drawOnChartArea: false, // only want the grid lines for one axis to show up
          },
        },
      },
    },
  });
};

window.LoadVesselMotionDisplay = function (responseObject) {
  openNav();

  // Clear service content div
  clearDiv("ServiceContent");

  // Create new div
  var div = document.createElement("Div");
  div.style.width = "95%";
  div.style.height = "95%";

  var canvas = document.createElement("canvas");
  canvas.style.width = "95%";
  canvas.style.height = "95%";
  canvas.id = "chartID";

  div.appendChild(canvas);

  document.getElementById("ServiceContent").appendChild(div);

  console.log(responseObject);

  var chartData = {
    labels: responseObject.time,
    datasets: [
      {
        data: responseObject.xAxisVibration,
        label: "X-Axis Acceleration",
        borderColor: "rgb(0, 0, 128)",
        backgroundColor: "rgba(0, 0, 128, 0.5)",

        fill: false,
        type: "line",
      },
      {
        data: responseObject.yAxisVibration,
        label: "Y-Axis Acceleration",
        borderColor: "rgb(0, 128, 0)",
        backgroundColor: "rgba(0, 128, 0, 0.5)",

        fill: false,
        type: "line",
      },
      {
        data: responseObject.zAxisVibration,
        label: "Z-Axis Acceleration",
        borderColor: "rgb(128, 0, 0)",
        backgroundColor: "rgba(128, 0, 0, 0.5)",

        fill: false,
        type: "line",
      },
    ],
  };

  var myChart = new Chart(canvas, {
    type: "line",
    data: chartData,
    options: {
      interaction: {
        mode: "index",
        intersect: false,
      },
      title: {
        display: true,
        text: "Bridge Acceleration over Time",
      },
      hover: {
        mode: "nearest",
        intersect: true,
      },
      scales: {
        xAxes: [
          {
            display: true,
            gridLines: {
              display: true,
              color: "black",
            },
            ticks: {
              fontColor: "black",
            },
          },
        ],
        yAxes: [
          {
            display: true,
            position: "left",
            gridLines: {
              display: true,
              color: "black",
            },
            ticks: {
              fontColor: "black",
              callback: function (value, index, values) {
                return value + "m/s^2";
              },
            },
          },
        ],
      },
    },
  });
};

window.LoadComfortDisplay = function () {
  console.log("Not implemented yet");
};

// ________QUERY FUNCTIONS_________

window.queryRAS = function () {
  console.log("Received RAS");

  LoadLoadingPage();

  var managerObject = JSON.parse(localStorage.getItem("myManager"));

  if (!managerObject.routeInfo.exists) {
    updateManagerRoute(managerObject);
  }

  console.log(managerObject);

  var rasService = new AnalysisServiceClient("http://localhost:8080");

  var request = new AnalysisRequest();
  request.setUnixTimeList(managerObject.routeInfo.unixTime);
  request.setLatitudeList(managerObject.routeInfo.latitude);
  request.setLongitudeList(managerObject.routeInfo.longitude);
  request.setHeadingList(managerObject.routeInfo.heading);
  request.setPropPitchList(managerObject.routeInfo.propellerPitch);
  request.setMotorSpeedList(managerObject.routeInfo.motorSpeed);
  request.setSogList(managerObject.routeInfo.sog);

  console.log("Request: ", request);

  var metadata = { authorisation: managerObject.userToken };
  rasService.analyseRoute(request, metadata, function (err, response) {
    if (err) {
      console.log(err.code);
      console.log(err.message);
      document.getElementById("consoleOutput").innerHTML =
        "Error " + err.code + ": " + err.message;

      LoadRouteAnalysisHome();
    } else {
      var responseObject = {
        time: response.getUnixTimeList(),
        averagePower: response.getAveragePower(),
        totalCost: response.getTotalCost(),
        AverageRMSX: response.getAverageRmsX(),
        AverageRMSY: response.getAverageRmsY(),
        AverageRMSZ: response.getAverageRmsZ(),
      };

      console.log("Response: ", response);

      LoadRouteAnalysisDisplay(responseObject);
    }
  });
};

window.queryPTS = function () {
  LoadLoadingPage();

  var managerObject = JSON.parse(localStorage.getItem("myManager"));

  if (!managerObject.routeInfo.exists) {
    updateManagerRoute(managerObject);
  }
  console.log(managerObject);

  var responseObject = {
    time: [
      new Date(1630080753 * 1000).toLocaleString(),
      new Date(1630084353 * 1000).toLocaleString(),
      new Date(1630087953 * 1000).toLocaleString(),
      new Date(1630095153 * 1000).toLocaleString(),
      new Date(1630098753 * 1000).toLocaleString(),
    ],
    powerEstimate: [3200, 3100, 2983, 3251, 3653],
    costEstimate: [1000, 2430, 3135, 4636, 6253],
    totalCost: [150000],
  };

  console.log(responseObject);

  LoadPowerTrainDisplay(responseObject);
};

// window.queryPTS = function () {
//   LoadLoadingPage();

//   var managerObject = JSON.parse(localStorage.getItem("myManager"));

//   if (!managerObject.routeInfo.exists) {
//     updateManagerRoute(managerObject);
//   }
//   console.log(managerObject);

//   var ptsService = new PTEstimateServiceClient("http://localhost:8080");

//   var request = new PTEstimateRequest();
//   request.setUnixTimeList(managerObject.routeInfo.unixTime);
//   request.setLatitudeList(managerObject.routeInfo.latitude);
//   request.setLongitudeList(managerObject.routeInfo.longitude);
//   request.setHeadingList(managerObject.routeInfo.heading);
//   request.setPropPitchList(managerObject.routeInfo.propellerPitch);
//   request.setMotorSpeedList(managerObject.routeInfo.motorSpeed);
//   request.setSogList(managerObject.routeInfo.sog);

//   console.log("Request: ", request);

//   var metadata = { authorisation: managerObject.userToken };
//   ptsService.estimatePowerTrain(request, metadata, function (err, response) {
//     if (err) {
//       console.log(err.code);
//       console.log(err.message);
//       document.getElementById("consoleOutput").innerHTML =
//         "Error " + err.code + ": " + err.message;

// LoadPowerTrainHome();
//     } else {
//       var responseObject = {
//         time: response.getUnixTimeList(),
//         powerEstimate: response.getPowerEstimateList(),
//         costEstimate: response.getCostEstimateList(),
//         totalCost: response.getTotalCost(),
//       };

//       console.log("Response: ", response);

//       LoadPowerTrainDisplay(responseObject);
//     }
//   });
// };

window.queryVMS = function () {
  LoadLoadingPage();

  var managerObject = JSON.parse(localStorage.getItem("myManager"));

  if (!managerObject.routeInfo.exists) {
    updateManagerRoute(managerObject);
  }
  console.log(managerObject);

  var responseObject = {
    time: [
      new Date(1630080753 * 1000).toLocaleString(),
      new Date(1630084353 * 1000).toLocaleString(),
      new Date(1630087953 * 1000).toLocaleString(),
      new Date(1630095153 * 1000).toLocaleString(),
      new Date(1630098753 * 1000).toLocaleString(),
    ],
    xAxisVibration: [0.00032524, 0.00024432, 0.00005324, 0.0000432, 0.0005325],
    yAxisVibration: [
      0.0000312434, 0.00005324, 0.000024325, 0.00006325, 0.00002356,
    ],
    zAxisVibration: [
      0.0005324, 0.0003245, 0.00003525, 0.0006436, 0.00003524235,
    ],
  };

  LoadVesselMotionDisplay(responseObject);

  console.log("Unimplemented");
};

// window.queryVMS = function () {
// LoadLoadingPage();

//   var managerObject = JSON.parse(localStorage.getItem("myManager"));

//   if (!managerObject.routeInfo.exists) {
//     updateManagerRoute(managerObject);
//   }
//   console.log(managerObject);

//   var vmsService = new VMEstimateServiceClient("http://localhost:8080");

//   var request = new VMEstimateRequest();
//   request.setUnixTimeList(managerObject.routeInfo.unixTime);
//   request.setLatitudeList(managerObject.routeInfo.latitude);
//   request.setLongitudeList(managerObject.routeInfo.longitude);
//   request.setHeadingList(managerObject.routeInfo.heading);
//   request.setPropPitchList(managerObject.routeInfo.propellerPitch);
//   request.setMotorSpeedList(managerObject.routeInfo.motorSpeed);
//   request.setSogList(managerObject.routeInfo.sog);

//   console.log("Request: ", request);

//   var metadata = { authorisation: managerObject.userToken };
//   vmsService.estimateVesselMotion(request, metadata, function (err, response) {
//     if (err) {
//       console.log(err.code);
//       console.log(err.message);
//       document.getElementById("consoleOutput").innerHTML =
//         "Error " + err.code + ": " + err.message;

// LoadVesselMotionHome();

//     } else {
//       var responseObject = {
//         // new Date(1630080753 * 1000).toLocaleString(),
//         time: response.getUnixTimeList(),
//         xAxisVibration: response.getAccelerationEstimateXList(),
//         yAxisVibration: response.getAccelerationEstimateYList(),
//         zAxisVibration: response.getAccelerationEstimateZList(),
//       };

//       console.log("Response: ", response);

//       LoadVesselMotionDisplay(responseObject);
//     }
//   });
// };

window.queryCS = function () {
  var managerObject = JSON.parse(localStorage.getItem("myManager"));

  if (!managerObject.routeInfo.exists) {
    updateManagerRoute(managerObject);
  }
  console.log("Unimplemented");
};

// ________SUPPORTING FUNCTIONS________

function toggleDiv(divID) {
  var x = document.getElementById(divID);
  if (x.style.display === "none") {
    x.style.display = "block";
  } else {
    x.style.display = "none";
  }
}

function clearDiv(divID) {
  /* This function removes all children from a div, clearing it to a blank state
   */

  // Get a div object
  var div = document.getElementById(divID);

  // Iterate through the div until there are no more children in it (basically popping a FIFO stack)
  while (div.firstChild) {
    // Remove the first child (popping the first element of a stack)
    div.removeChild(div.firstChild);
  }
}

function openNav() {
  document.getElementById("secondaryNavBar").style.display = "block";
  document.getElementById("SystemResponse").style.width = "90%";
  document.getElementById("SystemResponse").style.left = "10%";
  document.getElementById("ServiceContent").style.width = "82%";
  document.getElementById("ServiceContent").style.left = "13%";
}

function closeNav() {
  document.getElementById("secondaryNavBar").style.display = "none";
  document.getElementById("SystemResponse").style.width = "100%";
  document.getElementById("SystemResponse").style.left = "0%";
  document.getElementById("ServiceContent").style.width = "92%";
  document.getElementById("ServiceContent").style.left = "3%";
}

function createTextInput(placeholderName, id, parentDiv) {
  var elementName = document.createElement("input");

  elementName.placeholder = placeholderName;
  elementName.className = "textInput";
  elementName.type = "text";
  elementName.id = id;
  elementName.style.border = null;
  elementName.style.borderBottom = "thin solid white";
  elementName.style.background = "transparent";
  elementName.style.outline = null;
  elementName.style.height = "40px";
  elementName.style.color = "black";
  elementName.style.fontSize = "16px";

  // Add input elements to parent div
  document.getElementById(parentDiv).appendChild(elementName);
}

function createButtonInput(elementName, idName, parentDiv) {
  var elementName = document.createElement("input");

  elementName.type = "submit";
  elementName.id = idName;
  elementName.className = "processButton";
  elementName.style.height = "40px";
  elementName.style.fontSize = "18x";

  // Add input elements to parent div
  document.getElementById(parentDiv).appendChild(elementName);
}

function updateManagerRoute(managerObject) {
  managerObject.routeInfo.unixTime = document
    .getElementById("unixTimeIn")
    .value.split(",")
    .map((x) => +x);
  managerObject.routeInfo.latitude = document
    .getElementById("latitudeIn")
    .value.split(",")
    .map((x) => +x);
  managerObject.routeInfo.longitude = document
    .getElementById("longitudeIn")
    .value.split(",")
    .map((x) => +x);
  managerObject.routeInfo.heading = document
    .getElementById("headingIn")
    .value.split(",")
    .map((x) => +x);
  managerObject.routeInfo.propellerPitch = document
    .getElementById("propPitchIn")
    .value.split(",")
    .map((x) => +x);
  managerObject.routeInfo.motorSpeed = document
    .getElementById("motorSpeedIn")
    .value.split(",")
    .map((x) => +x);
  managerObject.routeInfo.sog = document
    .getElementById("sogIn")
    .value.split(",")
    .map((x) => +x);
  managerObject.routeInfo.exists = true;

  localStorage.setItem("myManager", JSON.stringify(managerObject));
}
