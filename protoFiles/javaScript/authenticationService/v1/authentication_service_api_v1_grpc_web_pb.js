/**
 * @fileoverview gRPC-Web generated client stub for authenticationServiceAPI.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.authenticationServiceAPI = {};
proto.authenticationServiceAPI.v1 = require('./authentication_service_api_v1_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.authenticationServiceAPI.v1.AuthenticationServiceClient =
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
proto.authenticationServiceAPI.v1.AuthenticationServicePromiseClient =
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
 *   !proto.authenticationServiceAPI.v1.NewUserRequest,
 *   !proto.authenticationServiceAPI.v1.LoginAuthResponse>}
 */
const methodDescriptor_AuthenticationService_CreateNewUser = new grpc.web.MethodDescriptor(
  '/authenticationServiceAPI.v1.AuthenticationService/CreateNewUser',
  grpc.web.MethodType.UNARY,
  proto.authenticationServiceAPI.v1.NewUserRequest,
  proto.authenticationServiceAPI.v1.LoginAuthResponse,
  /**
   * @param {!proto.authenticationServiceAPI.v1.NewUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.authenticationServiceAPI.v1.LoginAuthResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authenticationServiceAPI.v1.NewUserRequest,
 *   !proto.authenticationServiceAPI.v1.LoginAuthResponse>}
 */
const methodInfo_AuthenticationService_CreateNewUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authenticationServiceAPI.v1.LoginAuthResponse,
  /**
   * @param {!proto.authenticationServiceAPI.v1.NewUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.authenticationServiceAPI.v1.LoginAuthResponse.deserializeBinary
);


/**
 * @param {!proto.authenticationServiceAPI.v1.NewUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authenticationServiceAPI.v1.LoginAuthResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authenticationServiceAPI.v1.LoginAuthResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authenticationServiceAPI.v1.AuthenticationServiceClient.prototype.createNewUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authenticationServiceAPI.v1.AuthenticationService/CreateNewUser',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_CreateNewUser,
      callback);
};


/**
 * @param {!proto.authenticationServiceAPI.v1.NewUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authenticationServiceAPI.v1.LoginAuthResponse>}
 *     Promise that resolves to the response
 */
proto.authenticationServiceAPI.v1.AuthenticationServicePromiseClient.prototype.createNewUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authenticationServiceAPI.v1.AuthenticationService/CreateNewUser',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_CreateNewUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.authenticationServiceAPI.v1.LoginAuthRequest,
 *   !proto.authenticationServiceAPI.v1.LoginAuthResponse>}
 */
const methodDescriptor_AuthenticationService_LoginAuth = new grpc.web.MethodDescriptor(
  '/authenticationServiceAPI.v1.AuthenticationService/LoginAuth',
  grpc.web.MethodType.UNARY,
  proto.authenticationServiceAPI.v1.LoginAuthRequest,
  proto.authenticationServiceAPI.v1.LoginAuthResponse,
  /**
   * @param {!proto.authenticationServiceAPI.v1.LoginAuthRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.authenticationServiceAPI.v1.LoginAuthResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.authenticationServiceAPI.v1.LoginAuthRequest,
 *   !proto.authenticationServiceAPI.v1.LoginAuthResponse>}
 */
const methodInfo_AuthenticationService_LoginAuth = new grpc.web.AbstractClientBase.MethodInfo(
  proto.authenticationServiceAPI.v1.LoginAuthResponse,
  /**
   * @param {!proto.authenticationServiceAPI.v1.LoginAuthRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.authenticationServiceAPI.v1.LoginAuthResponse.deserializeBinary
);


/**
 * @param {!proto.authenticationServiceAPI.v1.LoginAuthRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.authenticationServiceAPI.v1.LoginAuthResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.authenticationServiceAPI.v1.LoginAuthResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.authenticationServiceAPI.v1.AuthenticationServiceClient.prototype.loginAuth =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/authenticationServiceAPI.v1.AuthenticationService/LoginAuth',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_LoginAuth,
      callback);
};


/**
 * @param {!proto.authenticationServiceAPI.v1.LoginAuthRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.authenticationServiceAPI.v1.LoginAuthResponse>}
 *     Promise that resolves to the response
 */
proto.authenticationServiceAPI.v1.AuthenticationServicePromiseClient.prototype.loginAuth =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/authenticationServiceAPI.v1.AuthenticationService/LoginAuth',
      request,
      metadata || {},
      methodDescriptor_AuthenticationService_LoginAuth);
};


module.exports = proto.authenticationServiceAPI.v1;

