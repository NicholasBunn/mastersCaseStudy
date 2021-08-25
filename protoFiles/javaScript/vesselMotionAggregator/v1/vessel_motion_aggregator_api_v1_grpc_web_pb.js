/**
 * @fileoverview gRPC-Web generated client stub for vesselMotionAggregatorAPI.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.vesselMotionAggregatorAPI = {};
proto.vesselMotionAggregatorAPI.v1 = require('./vessel_motion_aggregator_api_v1_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.vesselMotionAggregatorAPI.v1.VMEstimateServiceClient =
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
proto.vesselMotionAggregatorAPI.v1.VMEstimateServicePromiseClient =
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
 *   !proto.vesselMotionAggregatorAPI.v1.VMEstimateRequest,
 *   !proto.vesselMotionAggregatorAPI.v1.VMEstimateResponse>}
 */
const methodDescriptor_VMEstimateService_EstimateVesselMotion = new grpc.web.MethodDescriptor(
  '/vesselMotionAggregatorAPI.v1.VMEstimateService/EstimateVesselMotion',
  grpc.web.MethodType.UNARY,
  proto.vesselMotionAggregatorAPI.v1.VMEstimateRequest,
  proto.vesselMotionAggregatorAPI.v1.VMEstimateResponse,
  /**
   * @param {!proto.vesselMotionAggregatorAPI.v1.VMEstimateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.vesselMotionAggregatorAPI.v1.VMEstimateResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.vesselMotionAggregatorAPI.v1.VMEstimateRequest,
 *   !proto.vesselMotionAggregatorAPI.v1.VMEstimateResponse>}
 */
const methodInfo_VMEstimateService_EstimateVesselMotion = new grpc.web.AbstractClientBase.MethodInfo(
  proto.vesselMotionAggregatorAPI.v1.VMEstimateResponse,
  /**
   * @param {!proto.vesselMotionAggregatorAPI.v1.VMEstimateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.vesselMotionAggregatorAPI.v1.VMEstimateResponse.deserializeBinary
);


/**
 * @param {!proto.vesselMotionAggregatorAPI.v1.VMEstimateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.vesselMotionAggregatorAPI.v1.VMEstimateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.vesselMotionAggregatorAPI.v1.VMEstimateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.vesselMotionAggregatorAPI.v1.VMEstimateServiceClient.prototype.estimateVesselMotion =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/vesselMotionAggregatorAPI.v1.VMEstimateService/EstimateVesselMotion',
      request,
      metadata || {},
      methodDescriptor_VMEstimateService_EstimateVesselMotion,
      callback);
};


/**
 * @param {!proto.vesselMotionAggregatorAPI.v1.VMEstimateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.vesselMotionAggregatorAPI.v1.VMEstimateResponse>}
 *     Promise that resolves to the response
 */
proto.vesselMotionAggregatorAPI.v1.VMEstimateServicePromiseClient.prototype.estimateVesselMotion =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/vesselMotionAggregatorAPI.v1.VMEstimateService/EstimateVesselMotion',
      request,
      metadata || {},
      methodDescriptor_VMEstimateService_EstimateVesselMotion);
};


module.exports = proto.vesselMotionAggregatorAPI.v1;

