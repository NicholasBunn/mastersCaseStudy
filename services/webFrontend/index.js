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
const {
  PropellerLoadRequest,
  PropellerLoadResponse,
} = require("./../../protoFiles/javaScript/propellerMonitorService/v1/propeller_monitor_service_api_v1_pb.js");
const {
  MonitorPropellerServiceClient,
} = require("./../../protoFiles/javaScript/propellerMonitorService/v1/propeller_monitor_service_api_v1_grpc_web_pb.js");

import Chart from "chart.js/auto";
import zoomPlugin from "chartjs-plugin-zoom";

var permissionMapping = {
  admin: [
    "routeAnalysisService",
    "powerTrainService",
    "vesselMotionService",
    "propellerMonitorService",
  ],
  operator: [
    "routeAnalysisService",
    "powerTrainService",
    "vesselMotionService",
    "propellerMonitorService",
  ],
  engineer: [
    "powerTrainService",
    "vesselMotionService",
    "propellerMonitorService",
  ],
  researcher: [
    "routeAnalysisService",
    "powerTrainService",
    "vesselMotionService",
    "propellerMonitorService",
  ],
  guest: ["routeAnalysisService", "powerTrainService", "vesselMotionService"], // This is just to make dev easier
};

var temporalMapping = {
  routeAnalysisService: ["foresight", "insight", "hindsight"],
  powerTrainService: ["foresight"],
  vesselMotionService: ["foresight"],
  propellerMonitorService: ["insight"],
};

// A quick bit of code for the retry interceptor
const RetryInterceptor = function () {
  this.intercept = function (request, invoker) {
    console.log("Started interceptor");
    while (!response) {
      var response = invoker(request);
      console.log(response.a);
    }
    return response;
  };
};

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

  var loginService = new AuthenticationServiceClient(
    "http://localhost:10000",
    null,
    { streamInterceptors: [new RetryInterceptor()] }
  );

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
      var userManagerObject = {
        currentService: null,
        queryID: null,
        userToken: response.getAccessToken(),
        userPermissions: response.getPermissions(),
      };
      var routeManagerObject = {
        exists: false,
        routeInfo: {
          unixTime: [],
          latitude: [],
          longitude: [],
          heading: [],
          propellerPitch: [],
          motorSpeed: [],
          sog: [],
        },
      };

      localStorage.setItem("routeManager", JSON.stringify(routeManagerObject));
      localStorage.setItem("userManager", JSON.stringify(userManagerObject));

      // Hide login div
      toggleDiv("loginBox");

      // Load main interface div
      toggleDiv("mainInterface");

      for (var serviceName of permissionMapping[
        userManagerObject.userPermissions
      ]) {
        // Add buttons to the primary nav bar
        addInputToElement(
          "query" + serviceName + "Primary",
          "primaryNavBar",
          "button",
          "navButton",
          {
            height: "100%",
            width: "200px",
            fontSize: "18px",
            value: serviceName,
            callbackFunction:
              "Load" +
              serviceName.charAt(0).toUpperCase() +
              serviceName.slice(1) +
              "Home",
          }
        );
      }

      LoadHomePage();
    }
  });
};

window.LogMeOut = function () {
  localStorage.clear();

  location.reload();
  return false;
};

// ________LANDING PAGES________

window.LoadLoadingPage = function () {
  clearDiv("ServiceContent");
  clearDiv("consoleOutput");

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
  clearDiv("consoleOutput");

  closeNav();

  var userManagerObject = JSON.parse(localStorage.getItem("userManager"));

  document.getElementById("ServiceContent");
  // Add a welcome message
  // Add a subtitle here to guide the user

  // Load available services into blocks with a nice pic
  // Create new div

  // Add buttons to the service content div
  for (var serviceName of permissionMapping[
    userManagerObject.userPermissions
  ]) {
    addInputToElement(
      "query" + serviceName + "Main",
      "ServiceContent",
      "button",
      "processButton",
      {
        height: "200px",
        width: "300px",
        fontSize: "24px",
        value: serviceName,
        callbackFunction:
          "Load" +
          serviceName.charAt(0).toUpperCase() +
          serviceName.slice(1) +
          "Home",
      }
    );
  }
};

window.LoadRouteAnalysisServiceHome = function () {
  // Update manager object
  var userManagerObject = JSON.parse(localStorage.getItem("userManager"));

  userManagerObject.currentService = "routeAnalysisService";
  userManagerObject.queryID = "queryRAS";

  localStorage.setItem("userManager", JSON.stringify(userManagerObject));

  closeNav();

  // Clear service content div
  clearDiv("ServiceContent");
  clearDiv("temporals");

  // Create new div
  var div = document.createElement("Div");
  div.id = "contentDiv";
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

  // Add buttons to the service content div
  for (var temporalAspect of temporalMapping[
    userManagerObject.currentService
  ]) {
    console.log(temporalAspect);
    if (temporalAspect === "foresight") {
      var callbackFunc = "LoadRouteInput";
    }
    addInputToElement(
      temporalAspect + "Main",
      "contentDiv",
      "button",
      "processButton",
      {
        height: "200px",
        width: "300px",
        fontSize: "24px",
        value: temporalAspect,
        callbackFunction: callbackFunc,
      }
    );
    addInputToElement(
      temporalAspect + "Secondary",
      "temporals",
      "button",
      "navButton",
      {
        height: "50px",
        width: "100%",
        fontSize: "18px",
        value: temporalAspect,
        callbackFunction: callbackFunc,
      }
    );
  }
};

window.LoadPowerTrainServiceHome = function () {
  // Update manager object
  var userManagerObject = JSON.parse(localStorage.getItem("userManager"));

  userManagerObject.currentService = "powerTrainService";
  userManagerObject.queryID = "queryPTS";

  localStorage.setItem("userManager", JSON.stringify(userManagerObject));

  closeNav();

  // Clear service content div
  clearDiv("ServiceContent");
  clearDiv("temporals");

  // Create new div
  var div = document.createElement("Div");
  div.id = "contentDiv";
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

  // Add buttons to the service content div
  for (var temporalAspect of temporalMapping[
    userManagerObject.currentService
  ]) {
    console.log(temporalAspect);
    if (temporalAspect === "foresight") {
      var callbackFunc = "LoadRouteInput";
    }
    addInputToElement(temporalAspect, "contentDiv", "button", "processButton", {
      height: "200px",
      width: "300px",
      fontSize: "24px",
      value: temporalAspect,
      callbackFunction: callbackFunc,
    });
    addInputToElement(
      temporalAspect + "Secondary",
      "temporals",
      "button",
      "navButton",
      {
        height: "50px",
        width: "100%",
        fontSize: "18px",
        value: temporalAspect,
        callbackFunction: callbackFunc,
      }
    );
  }
};

window.LoadVesselMotionServiceHome = function () {
  // Update manager object
  var userManagerObject = JSON.parse(localStorage.getItem("userManager"));

  userManagerObject.currentService = "vesselMotionService";
  userManagerObject.queryID = "queryVMS";

  localStorage.setItem("userManager", JSON.stringify(userManagerObject));

  closeNav();

  // Clear service content div
  clearDiv("ServiceContent");
  clearDiv("temporals");

  // Create new div
  var div = document.createElement("Div");
  div.id = "contentDiv";
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

  // Add buttons to the service content div
  for (var temporalAspect of temporalMapping[
    userManagerObject.currentService
  ]) {
    console.log(temporalAspect);
    if (temporalAspect === "foresight") {
      var callbackFunc = "LoadRouteInput";
    }
    addInputToElement(temporalAspect, "contentDiv", "button", "processButton", {
      height: "200px",
      width: "300px",
      fontSize: "24px",
      value: temporalAspect,
      callbackFunction: callbackFunc,
    });
    addInputToElement(
      temporalAspect + "Secondary",
      "temporals",
      "button",
      "navButton",
      {
        height: "50px",
        width: "100%",
        fontSize: "18px",
        value: temporalAspect,
        callbackFunction: callbackFunc,
      }
    );
  }
};

window.LoadPropellerMonitorServiceHome = function () {
  // Update manager object
  var userManagerObject = JSON.parse(localStorage.getItem("userManager"));

  userManagerObject.currentService = "propellerMonitorService";
  userManagerObject.queryID = "queryPMS";

  localStorage.setItem("userManager", JSON.stringify(userManagerObject));

  // Clear service content div
  clearDiv("ServiceContent");
  clearDiv("temporals");

  // Create new div
  var div = document.createElement("Div");
  div.id = "contentDiv";
  div.style.width = "95%";
  div.style.height = "95%";

  var heading = document.createElement("h1");
  heading.className = "LandingPageHeading";
  heading.appendChild(
    document.createTextNode("Welcome to the propeller monitor service")
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

  // Add buttons to the service content div
  for (var temporalAspect of temporalMapping[
    userManagerObject.currentService
  ]) {
    console.log(temporalAspect);
    if (temporalAspect === "insight") {
      var callbackFunc = "LoadPropMonDisplay";
    }
    addInputToElement(temporalAspect, "contentDiv", "button", "processButton", {
      height: "200px",
      width: "300px",
      fontSize: "24px",
      value: temporalAspect,
      callbackFunction: callbackFunc,
    });
    addInputToElement(
      temporalAspect + "Secondary",
      "temporals",
      "button",
      "navButton",
      {
        height: "50px",
        width: "100%",
        fontSize: "18px",
        value: temporalAspect,
        callbackFunction: callbackFunc,
      }
    );
  }
};

// ________INPUT PAGES_________

window.LoadRouteInput = function () {
  var routeManagerObject = JSON.parse(localStorage.getItem("routeManager"));
  var userManagerObject = JSON.parse(localStorage.getItem("userManager"));

  // If a route has already been entered, make the request for that route
  if (routeManagerObject.exists) {
    window[userManagerObject.queryID]();
  } else {
    // Clear service content div
    clearDiv("ServiceContent");
    clearDiv("consoleOutput");

    openNav();

    // Create new div
    var div = document.createElement("Div");
    div.id = "loadRoute";
    div.style.width = "95%";
    div.style.height = "95%";

    // Add div to service content div
    document.getElementById("ServiceContent").appendChild(div);

    // Create input elements
    addInputToElement("unixTimeIn", "loadRoute", "text", "textInput", {
      height: "40px",
      width: "200px",
      fontSize: "16px",
      placeholder: "Unix time",
    });
    addInputToElement("latitudeIn", "loadRoute", "text", "textInput", {
      height: "40px",
      width: "200px",
      fontSize: "16px",
      placeholder: "Latitude",
    });
    addInputToElement("longitudeIn", "loadRoute", "text", "textInput", {
      height: "40px",
      width: "200px",
      fontSize: "16px",
      placeholder: "Longitude",
    });
    addInputToElement("headingIn", "loadRoute", "text", "textInput", {
      height: "40px",
      width: "200px",
      fontSize: "16px",
      placeholder: "Heading",
    });
    addInputToElement("propPitchIn", "loadRoute", "text", "textInput", {
      height: "40px",
      width: "200px",
      fontSize: "16px",
      placeholder: "Propeller Pitch",
    });
    addInputToElement("motorSpeedIn", "loadRoute", "text", "textInput", {
      height: "40px",
      width: "200px",
      fontSize: "16px",
      placeholder: "Motor Speed",
    });
    addInputToElement("sogIn", "loadRoute", "text", "textInput", {
      height: "40px",
      width: "200px",
      fontSize: "16px",
      placeholder: "SOG",
    });

    addInputToElement(
      "query" + userManagerObject.currentService,
      "loadRoute",
      "button",
      "processButton",
      {
        height: "40px",
        width: "200px",
        fontSize: "16px",
        value: "Submit",
        callbackFunction: userManagerObject.queryID,
      }
    );
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
  clearDiv("consoleOutput");

  // Create new div
  var div = document.createElement("Div");
  div.id = "contentDiv";
  div.style.width = "95%";
  div.style.height = "95%";

  var canvas = document.createElement("canvas");
  canvas.style.width = "95%";
  canvas.style.height = "95%";
  canvas.id = "chartID";

  div.appendChild(canvas);

  document.getElementById("ServiceContent").appendChild(div);

  addInputToElement("backButton", "contentDiv", "button", "systemButton", {
    height: "40px",
    width: "60px",
    fontSize: "18px",
    value: "Back",
    callbackFunction: "LoadPowerTrainServiceHome",
  });

  addInputToElement("newRoute", "contentDiv", "button", "systemButton", {
    height: "40px",
    width: "60px",
    fontSize: "18px",
    value: "New Route",
    callbackFunction: "AddNewRoute",
  });

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
  console.log(responseObject);

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
        x: {
          grid: {
            display: true,
            drawBorder: true,
            drawOnChartArea: true,
            drawTicks: true,
            color: "black",
          },
          ticks: {
            color: "black",
          },
        },
        y1: {
          type: "linear",
          display: true,
          position: "left",
          grid: {
            display: true,
            drawBorder: true,
            drawOnChartArea: true,
            drawTicks: true,
            color: "black",
          },
          ticks: {
            color: "black",
            fontColor: "black",
            callback: function (value, index, values) {
              return value + "kW";
            },
          },
        },
        y2: {
          type: "linear",
          display: true,
          position: "right",
          ticks: {
            color: "black",
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
  clearDiv("consoleOutput");

  // Create new div
  var div = document.createElement("Div");
  div.id = "contentDiv";
  div.style.width = "95%";
  div.style.height = "95%";

  var canvas = document.createElement("canvas");
  canvas.style.width = "95%";
  canvas.style.height = "95%";
  canvas.id = "chartID";

  div.appendChild(canvas);

  document.getElementById("ServiceContent").appendChild(div);

  addInputToElement("backButton", "contentDiv", "button", "systemButton", {
    height: "40px",
    width: "60px",
    fontSize: "18px",
    value: "Back",
    callbackFunction: "LoadPowerTrainServiceHome",
  });

  addInputToElement("newRoute", "contentDiv", "button", "systemButton", {
    height: "40px",
    width: "60px",
    fontSize: "18px",
    value: "New Route",
    callbackFunction: "AddNewRoute",
  });

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
        x: {
          grid: {
            display: true,
            drawBorder: true,
            drawOnChartArea: true,
            drawTicks: true,
            color: "black",
          },
          ticks: {
            color: "black",
          },
        },
        y: {
          grid: {
            display: true,
            drawBorder: true,
            drawOnChartArea: true,
            drawTicks: true,
            color: "black",
          },
          ticks: {
            color: "black",
            fontColor: "black",
            callback: function (value, index, values) {
              return value + "m/s^2";
            },
          },
        },
      },
    },
  });
};

window.LoadPropMonDisplay = function () {
  var userManagerObject = JSON.parse(localStorage.getItem("userManager"));

  openNav();

  // Clear service content div
  clearDiv("ServiceContent");
  clearDiv("consoleOutput");

  // Create new div
  var div = document.createElement("Div");
  div.id = "contentDiv";
  div.style.width = "95%";
  div.style.height = "95%";

  var canvas = document.createElement("canvas");
  canvas.style.width = "95%";
  canvas.style.height = "95%";
  canvas.id = "chartID";

  div.appendChild(canvas);

  document.getElementById("ServiceContent").appendChild(div);

  var responseObject = {
    QIce: [{ x: null, y: null }],
    QMotor: [{ x: null, y: null }],
    QProp: [{ x: null, y: null }],
    QIceThreshold: [{ x: null, y: null }],
    QMotorThreshold: [{ x: null, y: null }],
    QPropThreshold: [{ x: null, y: null }],
  };

  var chartData = {
    datasets: [
      {
        data: responseObject.QIce,
        label: "Ice-Induced Torque",
        borderColor: "rgb(0, 0, 128)",
        backgroundColor: "rgba(0, 0, 128, 0.5)",
        showLine: true,
      },
      {
        data: responseObject.QMotor,
        label: "Motor Torque",
        borderColor: "rgb(0, 128, 0)",
        backgroundColor: "rgba(0, 128, 0, 0.5)",
        showLine: true,
      },
      {
        data: responseObject.QProp,
        label: "Propeller Torque",
        borderColor: "rgb(128, 0, 0)",
        backgroundColor: "rgba(128, 0, 0, 0.5)",
        showLine: true,
      },
    ],
  };

  var myChart = new Chart(canvas, {
    type: "scatter",
    data: chartData,
    options: {
      plugins: {
        zoom: {
          pan: {
            enabled: true,
            mode: "xy",
            modifierKey: "ctrl",
          },
          zoom: {
            wheel: {
              enabled: true,
            },
            drag: {
              enabled: true,
            },
            mode: "xy",
          },
        },
      },
      interaction: {
        mode: "index",
        intersect: false,
      },
      title: {
        display: true,
        text: "Propeller Torques",
      },
      hover: {
        mode: "nearest",
        intersect: true,
      },
      scales: {
        x: {
          grid: {
            display: true,
            drawBorder: true,
            drawOnChartArea: true,
            drawTicks: true,
            color: "black",
          },
          ticks: {
            color: "black",
          },
        },
        y: {
          grid: {
            display: true,
            drawBorder: true,
            drawOnChartArea: true,
            drawTicks: true,
            color: "black",
          },
          ticks: {
            color: "black",
          },
        },
      },
    },
  });
  Chart.register(zoomPlugin);

  var propMonService = new MonitorPropellerServiceClient(
    "http://localhost:10000",
    null,
    { streamInterceptors: [new RetryInterceptor()] }
  );

  var request = new PropellerLoadRequest();

  var metadata = { authorisation: userManagerObject.userToken };

  var stream = propMonService.estimatePropellerLoad(request, metadata);

  stream.on("data", function (response) {
    // Update response object
    console.log(myChart.data);
    // for (const [index, element] of a.entries())
    console.log("NEW DATA :)");
    for (var i = 0; i < response.getUnixTimeList().length; i++) {
      responseObject.QIce.push({
        x: response.getUnixTimeList()[i],
        y: response.getIceLoadList()[i],
      });
      myChart.data.datasets[0].data.push(responseObject.QIce);
      responseObject.QMotor.push({
        x: response.getUnixTimeList()[i],
        y: response.getMotorTorqueList()[i],
      });
      myChart.data.datasets[1].data.push(responseObject.QMotor);
      responseObject.QProp.push({
        x: response.getUnixTimeList()[i],
        y: response.getPropellerTorqueList()[i],
      });
      myChart.data.datasets[2].data.push(responseObject.QProp);
    }

    // Update with new data
    console.log(myChart.data);
    myChart.update();
  });

  stream.on("status", function (status) {
    console.log(status.code);
    console.log(status.details);
    console.log(status.metadata);
  });

  stream.on("end", function (end) {
    console.log("Server-side timeout due to no new data.");
  });

  addInputToElement("backButton", "contentDiv", "button", "systemButton", {
    height: "40px",
    width: "60px",
    fontSize: "18px",
    value: "Back",
    callbackFunction: "stream.cancel()",
  });
};

// ________QUERY FUNCTIONS_________

window.queryRAS = function () {
  console.log("Received RAS");

  var routeManagerObject = JSON.parse(localStorage.getItem("routeManager"));
  var userManagerObject = JSON.parse(localStorage.getItem("userManager"));

  if (!routeManagerObject.exists) {
    updateManagerRoute(routeManagerObject);
  }

  LoadLoadingPage();

  var rasService = new AnalysisServiceClient("http://localhost:10000", null, {
    streamInterceptors: [new RetryInterceptor()],
  });

  var request = new AnalysisRequest();
  request.setUnixTimeList(routeManagerObject.routeInfo.unixTime);
  request.setLatitudeList(routeManagerObject.routeInfo.latitude);
  request.setLongitudeList(routeManagerObject.routeInfo.longitude);
  request.setHeadingList(routeManagerObject.routeInfo.heading);
  request.setPropPitchList(routeManagerObject.routeInfo.propellerPitch);
  request.setMotorSpeedList(routeManagerObject.routeInfo.motorSpeed);
  request.setSogList(routeManagerObject.routeInfo.sog);

  console.log("Request: ", request);

  var metadata = { authorisation: userManagerObject.userToken };
  rasService.analyseRoute(request, metadata, function (err, response) {
    if (err) {
      console.log(err.code);
      console.log(err.message);
      document.getElementById("consoleOutput").textContent =
        "Error " + err.code + ": " + err.message;

      LoadRouteAnalysisServiceHome();
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

// window.queryPTS = function () {
//   var userManagerObject = JSON.parse(localStorage.getItem("userManager"));

//   if (!userManagerObject.routeInfo.exists) {
//     updateManagerRoute(userManagerObject);
//   }
//   console.log(userManagerObject);

//   LoadLoadingPage();

//   var responseObject = {
//     time: [
//       new Date(1630080753 * 1000).toLocaleString(),
//       new Date(1630084353 * 1000).toLocaleString(),
//       new Date(1630087953 * 1000).toLocaleString(),
//       new Date(1630095153 * 1000).toLocaleString(),
//       new Date(1630098753 * 1000).toLocaleString(),
//     ],
//     powerEstimate: [3200, 3100, 2983, 3251, 3653],
//     costEstimate: [1000, 2430, 3135, 4636, 6253],
//     totalCost: [150000],
//   };

//   console.log(responseObject);

//   LoadPowerTrainDisplay(responseObject);
// };

window.queryPTS = function () {
  var routeManagerObject = JSON.parse(localStorage.getItem("routeManager"));
  var userManagerObject = JSON.parse(localStorage.getItem("userManager"));

  if (!routeManagerObject.exists) {
    updateManagerRoute(routeManagerObject);
  }

  LoadLoadingPage();

  var ptsService = new PTEstimateServiceClient("http://localhost:10000", null, {
    streamInterceptors: [new RetryInterceptor()],
  });

  var request = new PTEstimateRequest();
  request.setUnixTimeList(routeManagerObject.routeInfo.unixTime);
  request.setLatitudeList(routeManagerObject.routeInfo.latitude);
  request.setLongitudeList(routeManagerObject.routeInfo.longitude);
  request.setHeadingList(routeManagerObject.routeInfo.heading);
  request.setPropPitchList(routeManagerObject.routeInfo.propellerPitch);
  request.setMotorSpeedList(routeManagerObject.routeInfo.motorSpeed);
  request.setSogList(routeManagerObject.routeInfo.sog);

  console.log("Request: ", request);

  var metadata = { authorisation: userManagerObject.userToken };
  ptsService.estimatePowerTrain(request, metadata, function (err, response) {
    if (err) {
      console.log(err.code);
      console.log(err.message);
      document.getElementById("consoleOutput").textContent =
        "Error " + err.code + ": " + err.message;

      LoadPowerTrainServiceHome();
    } else {
      var responseObject = {
        time: response.getUnixTimeList(),
        powerEstimate: response.getPowerEstimateList(),
        costEstimate: response.getCostEstimateList(),
        totalCost: response.getTotalCost(),
      };

      console.log("Response: ", response);

      LoadPowerTrainDisplay(responseObject);
    }
  });
};

// window.queryVMS = function () {
//   var userManagerObject = JSON.parse(localStorage.getItem("userManager"));

//   if (!userManagerObject.routeInfo.exists) {
//     updateManagerRoute(userManagerObject);
//   }
//   console.log(userManagerObject);

//   LoadLoadingPage();

//   var responseObject = {
//     time: [
//       new Date(1630080753 * 1000).toLocaleString(),
//       new Date(1630084353 * 1000).toLocaleString(),
//       new Date(1630087953 * 1000).toLocaleString(),
//       new Date(1630095153 * 1000).toLocaleString(),
//       new Date(1630098753 * 1000).toLocaleString(),
//     ],
//     xAxisVibration: [0.00032524, 0.00024432, 0.00005324, 0.0000432, 0.0005325],
//     yAxisVibration: [
//       0.0000312434, 0.00005324, 0.000024325, 0.00006325, 0.00002356,
//     ],
//     zAxisVibration: [
//       0.0005324, 0.0003245, 0.00003525, 0.0006436, 0.00003524235,
//     ],
//   };

//   LoadVesselMotionDisplay(responseObject);

//   console.log("Unimplemented");
// };

window.queryVMS = function () {
  var routeManagerObject = JSON.parse(localStorage.getItem("routeManager"));
  var userManagerObject = JSON.parse(localStorage.getItem("userManager"));

  if (!routeManagerObject.exists) {
    updateManagerRoute(routeManagerObject);
  }
  console.log(userManagerObject);

  LoadLoadingPage();

  var vmsService = new VMEstimateServiceClient("http://localhost:10000", null, {
    streamInterceptors: [new RetryInterceptor()],
  });

  var request = new VMEstimateRequest();
  request.setUnixTimeList(routeManagerObject.routeInfo.unixTime);
  request.setLatitudeList(routeManagerObject.routeInfo.latitude);
  request.setLongitudeList(routeManagerObject.routeInfo.longitude);
  request.setHeadingList(routeManagerObject.routeInfo.heading);
  request.setPropPitchList(routeManagerObject.routeInfo.propellerPitch);
  request.setMotorSpeedList(routeManagerObject.routeInfo.motorSpeed);
  request.setSogList(routeManagerObject.routeInfo.sog);

  console.log("Request: ", request);

  var metadata = { authorisation: userManagerObject.userToken };
  vmsService.estimateVesselMotion(request, metadata, function (err, response) {
    if (err) {
      console.log(err.code);
      console.log(err.message);
      document.getElementById("consoleOutput").textContent =
        "Error " + err.code + ": " + err.message;

      LoadVesselMotionServiceHome();
    } else {
      var responseObject = {
        // new Date(1630080753 * 1000).toLocaleString(),
        time: response.getUnixTimeList(),
        xAxisVibration: response.getAccelerationEstimateXList(),
        yAxisVibration: response.getAccelerationEstimateYList(),
        zAxisVibration: response.getAccelerationEstimateZList(),
      };

      console.log("Response: ", response);

      LoadVesselMotionDisplay(responseObject);
    }
  });
};

window.queryPMS = function () {
  console.log("UNIMPLEMENTED");
};

// ________SUPPORTING FUNCTIONS________

window.AddNewRoute = function () {
  var routeManagerObject = JSON.parse(localStorage.getItem("routeManager"));
  routeManagerObject.exists = false;
  localStorage.setItem("routeManager", JSON.stringify(routeManagerObject));

  LoadRouteInput();
};

function addInputToElement(
  elementID,
  parentElement,
  elementType,
  elementClass,
  elementInfo
) {
  var newElement = document.createElement("input");

  newElement.id = elementID;
  newElement.type = elementType;
  newElement.className = elementClass;
  newElement.style.height = elementInfo.height;
  newElement.style.width = elementInfo.width;
  newElement.style.fontSize = elementInfo.fontSize;

  if (elementType === "button") {
    newElement.value = elementInfo.value;
    // Only try to add a callback function if one has been provided
    if (elementInfo.callbackFunction) {
      newElement.onclick = window[elementInfo.callbackFunction];
    }
  } else if (elementType === "text") {
    newElement.placeholder = elementInfo.placeholder;
  }

  document.getElementById(parentElement).appendChild(newElement);
}

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

function updateManagerRoute(routeManagerObject) {
  routeManagerObject.routeInfo.unixTime = document
    .getElementById("unixTimeIn")
    .value.split(",")
    .map((x) => +x);
  routeManagerObject.routeInfo.latitude = document
    .getElementById("latitudeIn")
    .value.split(",")
    .map((x) => +x);
  routeManagerObject.routeInfo.longitude = document
    .getElementById("longitudeIn")
    .value.split(",")
    .map((x) => +x);
  routeManagerObject.routeInfo.heading = document
    .getElementById("headingIn")
    .value.split(",")
    .map((x) => +x);
  routeManagerObject.routeInfo.propellerPitch = document
    .getElementById("propPitchIn")
    .value.split(",")
    .map((x) => +x);
  routeManagerObject.routeInfo.motorSpeed = document
    .getElementById("motorSpeedIn")
    .value.split(",")
    .map((x) => +x);
  routeManagerObject.routeInfo.sog = document
    .getElementById("sogIn")
    .value.split(",")
    .map((x) => +x);
  routeManagerObject.exists = true;

  localStorage.setItem("routeManager", JSON.stringify(routeManagerObject));
}
