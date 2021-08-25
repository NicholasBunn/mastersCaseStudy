/**
 * @fileoverview gRPC-Web generated client stub for powerTrainAggregatorAPI.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.powerTrainAggregatorAPI = {};
proto.powerTrainAggregatorAPI.v1 = require('./power_train_aggregator_api_v1_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.powerTrainAggregatorAPI.v1.PTEstimateServiceClient =
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
proto.powerTrainAggregatorAPI.v1.PTEstimateServicePromiseClient =
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
 *   !proto.powerTrainAggregatorAPI.v1.PTEstimateRequest,
 *   !proto.powerTrainAggregatorAPI.v1.PTEstimateResponse>}
 */
const methodDescriptor_PTEstimateService_EstimatePowerTrain = new grpc.web.MethodDescriptor(
  '/powerTrainAggregatorAPI.v1.PTEstimateService/EstimatePowerTrain',
  grpc.web.MethodType.UNARY,
  proto.powerTrainAggregatorAPI.v1.PTEstimateRequest,
  proto.powerTrainAggregatorAPI.v1.PTEstimateResponse,
  /**
   * @param {!proto.powerTrainAggregatorAPI.v1.PTEstimateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerTrainAggregatorAPI.v1.PTEstimateResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.powerTrainAggregatorAPI.v1.PTEstimateRequest,
 *   !proto.powerTrainAggregatorAPI.v1.PTEstimateResponse>}
 */
const methodInfo_PTEstimateService_EstimatePowerTrain = new grpc.web.AbstractClientBase.MethodInfo(
  proto.powerTrainAggregatorAPI.v1.PTEstimateResponse,
  /**
   * @param {!proto.powerTrainAggregatorAPI.v1.PTEstimateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.powerTrainAggregatorAPI.v1.PTEstimateResponse.deserializeBinary
);


/**
 * @param {!proto.powerTrainAggregatorAPI.v1.PTEstimateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.powerTrainAggregatorAPI.v1.PTEstimateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.powerTrainAggregatorAPI.v1.PTEstimateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.powerTrainAggregatorAPI.v1.PTEstimateServiceClient.prototype.estimatePowerTrain =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/powerTrainAggregatorAPI.v1.PTEstimateService/EstimatePowerTrain',
      request,
      metadata || {},
      methodDescriptor_PTEstimateService_EstimatePowerTrain,
      callback);
};


/**
 * @param {!proto.powerTrainAggregatorAPI.v1.PTEstimateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.powerTrainAggregatorAPI.v1.PTEstimateResponse>}
 *     Promise that resolves to the response
 */
proto.powerTrainAggregatorAPI.v1.PTEstimateServicePromiseClient.prototype.estimatePowerTrain =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/powerTrainAggregatorAPI.v1.PTEstimateService/EstimatePowerTrain',
      request,
      metadata || {},
      methodDescriptor_PTEstimateService_EstimatePowerTrain);
};


module.exports = proto.powerTrainAggregatorAPI.v1;

