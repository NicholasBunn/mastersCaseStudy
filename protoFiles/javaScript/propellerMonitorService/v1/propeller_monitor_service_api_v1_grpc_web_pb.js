/**
 * @fileoverview gRPC-Web generated client stub for propellerMonitorService.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.propellerMonitorService = {};
proto.propellerMonitorService.v1 = require('./propeller_monitor_service_api_v1_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.propellerMonitorService.v1.MonitorPropellerServiceClient =
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
proto.propellerMonitorService.v1.MonitorPropellerServicePromiseClient =
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
 *   !proto.propellerMonitorService.v1.PropellerLoadRequest,
 *   !proto.propellerMonitorService.v1.PropellerLoadResponse>}
 */
const methodDescriptor_MonitorPropellerService_EstimatePropellerLoad = new grpc.web.MethodDescriptor(
  '/propellerMonitorService.v1.MonitorPropellerService/EstimatePropellerLoad',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.propellerMonitorService.v1.PropellerLoadRequest,
  proto.propellerMonitorService.v1.PropellerLoadResponse,
  /**
   * @param {!proto.propellerMonitorService.v1.PropellerLoadRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.propellerMonitorService.v1.PropellerLoadResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.propellerMonitorService.v1.PropellerLoadRequest,
 *   !proto.propellerMonitorService.v1.PropellerLoadResponse>}
 */
const methodInfo_MonitorPropellerService_EstimatePropellerLoad = new grpc.web.AbstractClientBase.MethodInfo(
  proto.propellerMonitorService.v1.PropellerLoadResponse,
  /**
   * @param {!proto.propellerMonitorService.v1.PropellerLoadRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.propellerMonitorService.v1.PropellerLoadResponse.deserializeBinary
);


/**
 * @param {!proto.propellerMonitorService.v1.PropellerLoadRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.propellerMonitorService.v1.PropellerLoadResponse>}
 *     The XHR Node Readable Stream
 */
proto.propellerMonitorService.v1.MonitorPropellerServiceClient.prototype.estimatePropellerLoad =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/propellerMonitorService.v1.MonitorPropellerService/EstimatePropellerLoad',
      request,
      metadata || {},
      methodDescriptor_MonitorPropellerService_EstimatePropellerLoad);
};


/**
 * @param {!proto.propellerMonitorService.v1.PropellerLoadRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.propellerMonitorService.v1.PropellerLoadResponse>}
 *     The XHR Node Readable Stream
 */
proto.propellerMonitorService.v1.MonitorPropellerServicePromiseClient.prototype.estimatePropellerLoad =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/propellerMonitorService.v1.MonitorPropellerService/EstimatePropellerLoad',
      request,
      metadata || {},
      methodDescriptor_MonitorPropellerService_EstimatePropellerLoad);
};


module.exports = proto.propellerMonitorService.v1;

