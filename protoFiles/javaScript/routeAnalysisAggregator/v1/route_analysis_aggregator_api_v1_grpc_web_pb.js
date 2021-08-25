/**
 * @fileoverview gRPC-Web generated client stub for routeAnalysisAggregatorAPI.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.routeAnalysisAggregatorAPI = {};
proto.routeAnalysisAggregatorAPI.v1 = require('./route_analysis_aggregator_api_v1_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.routeAnalysisAggregatorAPI.v1.AnalysisServiceClient =
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
proto.routeAnalysisAggregatorAPI.v1.AnalysisServicePromiseClient =
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
 *   !proto.routeAnalysisAggregatorAPI.v1.AnalysisRequest,
 *   !proto.routeAnalysisAggregatorAPI.v1.AnalysisResponse>}
 */
const methodDescriptor_AnalysisService_AnalyseRoute = new grpc.web.MethodDescriptor(
  '/routeAnalysisAggregatorAPI.v1.AnalysisService/AnalyseRoute',
  grpc.web.MethodType.UNARY,
  proto.routeAnalysisAggregatorAPI.v1.AnalysisRequest,
  proto.routeAnalysisAggregatorAPI.v1.AnalysisResponse,
  /**
   * @param {!proto.routeAnalysisAggregatorAPI.v1.AnalysisRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.routeAnalysisAggregatorAPI.v1.AnalysisResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.routeAnalysisAggregatorAPI.v1.AnalysisRequest,
 *   !proto.routeAnalysisAggregatorAPI.v1.AnalysisResponse>}
 */
const methodInfo_AnalysisService_AnalyseRoute = new grpc.web.AbstractClientBase.MethodInfo(
  proto.routeAnalysisAggregatorAPI.v1.AnalysisResponse,
  /**
   * @param {!proto.routeAnalysisAggregatorAPI.v1.AnalysisRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.routeAnalysisAggregatorAPI.v1.AnalysisResponse.deserializeBinary
);


/**
 * @param {!proto.routeAnalysisAggregatorAPI.v1.AnalysisRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.routeAnalysisAggregatorAPI.v1.AnalysisResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.routeAnalysisAggregatorAPI.v1.AnalysisResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.routeAnalysisAggregatorAPI.v1.AnalysisServiceClient.prototype.analyseRoute =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/routeAnalysisAggregatorAPI.v1.AnalysisService/AnalyseRoute',
      request,
      metadata || {},
      methodDescriptor_AnalysisService_AnalyseRoute,
      callback);
};


/**
 * @param {!proto.routeAnalysisAggregatorAPI.v1.AnalysisRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.routeAnalysisAggregatorAPI.v1.AnalysisResponse>}
 *     Promise that resolves to the response
 */
proto.routeAnalysisAggregatorAPI.v1.AnalysisServicePromiseClient.prototype.analyseRoute =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/routeAnalysisAggregatorAPI.v1.AnalysisService/AnalyseRoute',
      request,
      metadata || {},
      methodDescriptor_AnalysisService_AnalyseRoute);
};


module.exports = proto.routeAnalysisAggregatorAPI.v1;

