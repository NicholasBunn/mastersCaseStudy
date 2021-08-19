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
  options['format'] = 'text';

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
  options['format'] = 'text';

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
  options['format'] = 'text';

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
  options['format'] = 'text';

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


module.exports = proto.webGateway.v1;

