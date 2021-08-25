/**
 * @fileoverview gRPC-Web generated client stub for webGateway.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.webGateway = {};
proto.webGateway.v1 = require('./web_gateway_api_v1_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.webGateway.v1.LoginServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.webGateway.v1.LoginServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.webGateway.v1.LoginRequest,
 *   !proto.webGateway.v1.LoginResponse>}
 */
const methodDescriptor_LoginService_Login = new grpc.web.MethodDescriptor(
  '/webGateway.v1.LoginService/Login',
  grpc.web.MethodType.UNARY,
  proto.webGateway.v1.LoginRequest,
  proto.webGateway.v1.LoginResponse,
  /**
   * @param {!proto.webGateway.v1.LoginRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.webGateway.v1.LoginResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.webGateway.v1.LoginRequest,
 *   !proto.webGateway.v1.LoginResponse>}
 */
const methodInfo_LoginService_Login = new grpc.web.AbstractClientBase.MethodInfo(
  proto.webGateway.v1.LoginResponse,
  /**
   * @param {!proto.webGateway.v1.LoginRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.webGateway.v1.LoginResponse.deserializeBinary
);


/**
 * @param {!proto.webGateway.v1.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.webGateway.v1.LoginResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.webGateway.v1.LoginResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.webGateway.v1.LoginServiceClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/webGateway.v1.LoginService/Login',
      request,
      metadata || {},
      methodDescriptor_LoginService_Login,
      callback);
};


/**
 * @param {!proto.webGateway.v1.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.webGateway.v1.LoginResponse>}
 *     Promise that resolves to the response
 */
proto.webGateway.v1.LoginServicePromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/webGateway.v1.LoginService/Login',
      request,
      metadata || {},
      methodDescriptor_LoginService_Login);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.webGateway.v1.RouteAnalysisAggregatorClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.webGateway.v1.RouteAnalysisAggregatorPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.webGateway.v1.RouteAnalysisRequest,
 *   !proto.webGateway.v1.RouteAnalysisResponse>}
 */
const methodDescriptor_RouteAnalysisAggregator_RouteAnalysis = new grpc.web.MethodDescriptor(
  '/webGateway.v1.RouteAnalysisAggregator/RouteAnalysis',
  grpc.web.MethodType.UNARY,
  proto.webGateway.v1.RouteAnalysisRequest,
  proto.webGateway.v1.RouteAnalysisResponse,
  /**
   * @param {!proto.webGateway.v1.RouteAnalysisRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.webGateway.v1.RouteAnalysisResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.webGateway.v1.RouteAnalysisRequest,
 *   !proto.webGateway.v1.RouteAnalysisResponse>}
 */
const methodInfo_RouteAnalysisAggregator_RouteAnalysis = new grpc.web.AbstractClientBase.MethodInfo(
  proto.webGateway.v1.RouteAnalysisResponse,
  /**
   * @param {!proto.webGateway.v1.RouteAnalysisRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.webGateway.v1.RouteAnalysisResponse.deserializeBinary
);


/**
 * @param {!proto.webGateway.v1.RouteAnalysisRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.webGateway.v1.RouteAnalysisResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.webGateway.v1.RouteAnalysisResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.webGateway.v1.RouteAnalysisAggregatorClient.prototype.routeAnalysis =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/webGateway.v1.RouteAnalysisAggregator/RouteAnalysis',
      request,
      metadata || {},
      methodDescriptor_RouteAnalysisAggregator_RouteAnalysis,
      callback);
};


/**
 * @param {!proto.webGateway.v1.RouteAnalysisRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.webGateway.v1.RouteAnalysisResponse>}
 *     Promise that resolves to the response
 */
proto.webGateway.v1.RouteAnalysisAggregatorPromiseClient.prototype.routeAnalysis =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/webGateway.v1.RouteAnalysisAggregator/RouteAnalysis',
      request,
      metadata || {},
      methodDescriptor_RouteAnalysisAggregator_RouteAnalysis);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.webGateway.v1.RoutePowerAggregatorClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.webGateway.v1.RoutePowerAggregatorPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.webGateway.v1.RoutePowerRequest,
 *   !proto.webGateway.v1.RoutePowerResponse>}
 */
const methodDescriptor_RoutePowerAggregator_RoutePower = new grpc.web.MethodDescriptor(
  '/webGateway.v1.RoutePowerAggregator/RoutePower',
  grpc.web.MethodType.UNARY,
  proto.webGateway.v1.RoutePowerRequest,
  proto.webGateway.v1.RoutePowerResponse,
  /**
   * @param {!proto.webGateway.v1.RoutePowerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.webGateway.v1.RoutePowerResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.webGateway.v1.RoutePowerRequest,
 *   !proto.webGateway.v1.RoutePowerResponse>}
 */
const methodInfo_RoutePowerAggregator_RoutePower = new grpc.web.AbstractClientBase.MethodInfo(
  proto.webGateway.v1.RoutePowerResponse,
  /**
   * @param {!proto.webGateway.v1.RoutePowerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.webGateway.v1.RoutePowerResponse.deserializeBinary
);


/**
 * @param {!proto.webGateway.v1.RoutePowerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.webGateway.v1.RoutePowerResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.webGateway.v1.RoutePowerResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.webGateway.v1.RoutePowerAggregatorClient.prototype.routePower =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/webGateway.v1.RoutePowerAggregator/RoutePower',
      request,
      metadata || {},
      methodDescriptor_RoutePowerAggregator_RoutePower,
      callback);
};


/**
 * @param {!proto.webGateway.v1.RoutePowerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.webGateway.v1.RoutePowerResponse>}
 *     Promise that resolves to the response
 */
proto.webGateway.v1.RoutePowerAggregatorPromiseClient.prototype.routePower =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/webGateway.v1.RoutePowerAggregator/RoutePower',
      request,
      metadata || {},
      methodDescriptor_RoutePowerAggregator_RoutePower);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.webGateway.v1.RouteMotionAggregatorClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.webGateway.v1.RouteMotionAggregatorPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.webGateway.v1.RouteMotionRequest,
 *   !proto.webGateway.v1.RouteMotionResponse>}
 */
const methodDescriptor_RouteMotionAggregator_RouteMotion = new grpc.web.MethodDescriptor(
  '/webGateway.v1.RouteMotionAggregator/RouteMotion',
  grpc.web.MethodType.UNARY,
  proto.webGateway.v1.RouteMotionRequest,
  proto.webGateway.v1.RouteMotionResponse,
  /**
   * @param {!proto.webGateway.v1.RouteMotionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.webGateway.v1.RouteMotionResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.webGateway.v1.RouteMotionRequest,
 *   !proto.webGateway.v1.RouteMotionResponse>}
 */
const methodInfo_RouteMotionAggregator_RouteMotion = new grpc.web.AbstractClientBase.MethodInfo(
  proto.webGateway.v1.RouteMotionResponse,
  /**
   * @param {!proto.webGateway.v1.RouteMotionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.webGateway.v1.RouteMotionResponse.deserializeBinary
);


/**
 * @param {!proto.webGateway.v1.RouteMotionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.webGateway.v1.RouteMotionResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.webGateway.v1.RouteMotionResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.webGateway.v1.RouteMotionAggregatorClient.prototype.routeMotion =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/webGateway.v1.RouteMotionAggregator/RouteMotion',
      request,
      metadata || {},
      methodDescriptor_RouteMotionAggregator_RouteMotion,
      callback);
};


/**
 * @param {!proto.webGateway.v1.RouteMotionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.webGateway.v1.RouteMotionResponse>}
 *     Promise that resolves to the response
 */
proto.webGateway.v1.RouteMotionAggregatorPromiseClient.prototype.routeMotion =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/webGateway.v1.RouteMotionAggregator/RouteMotion',
      request,
      metadata || {},
      methodDescriptor_RouteMotionAggregator_RouteMotion);
};


module.exports = proto.webGateway.v1;

