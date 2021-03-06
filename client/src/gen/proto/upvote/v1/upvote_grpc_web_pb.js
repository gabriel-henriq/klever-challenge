/**
 * @fileoverview gRPC-Web generated client stub for upvote.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.upvote = {};
proto.upvote.v1 = require('./upvote_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.upvote.v1.UpvoteServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

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
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.upvote.v1.UpvoteServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

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
 *   !proto.upvote.v1.CreateBookRequest,
 *   !proto.upvote.v1.CreateBookResponse>}
 */
const methodDescriptor_UpvoteService_CreateBook = new grpc.web.MethodDescriptor(
  '/upvote.v1.UpvoteService/CreateBook',
  grpc.web.MethodType.UNARY,
  proto.upvote.v1.CreateBookRequest,
  proto.upvote.v1.CreateBookResponse,
  /**
   * @param {!proto.upvote.v1.CreateBookRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.upvote.v1.CreateBookResponse.deserializeBinary
);


/**
 * @param {!proto.upvote.v1.CreateBookRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.upvote.v1.CreateBookResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.upvote.v1.CreateBookResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.upvote.v1.UpvoteServiceClient.prototype.createBook =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/upvote.v1.UpvoteService/CreateBook',
      request,
      metadata || {},
      methodDescriptor_UpvoteService_CreateBook,
      callback);
};


/**
 * @param {!proto.upvote.v1.CreateBookRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.upvote.v1.CreateBookResponse>}
 *     Promise that resolves to the response
 */
proto.upvote.v1.UpvoteServicePromiseClient.prototype.createBook =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/upvote.v1.UpvoteService/CreateBook',
      request,
      metadata || {},
      methodDescriptor_UpvoteService_CreateBook);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.upvote.v1.UpvoteRequest,
 *   !proto.upvote.v1.UpvoteResponse>}
 */
const methodDescriptor_UpvoteService_Upvote = new grpc.web.MethodDescriptor(
  '/upvote.v1.UpvoteService/Upvote',
  grpc.web.MethodType.UNARY,
  proto.upvote.v1.UpvoteRequest,
  proto.upvote.v1.UpvoteResponse,
  /**
   * @param {!proto.upvote.v1.UpvoteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.upvote.v1.UpvoteResponse.deserializeBinary
);


/**
 * @param {!proto.upvote.v1.UpvoteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.upvote.v1.UpvoteResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.upvote.v1.UpvoteResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.upvote.v1.UpvoteServiceClient.prototype.upvote =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/upvote.v1.UpvoteService/Upvote',
      request,
      metadata || {},
      methodDescriptor_UpvoteService_Upvote,
      callback);
};


/**
 * @param {!proto.upvote.v1.UpvoteRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.upvote.v1.UpvoteResponse>}
 *     Promise that resolves to the response
 */
proto.upvote.v1.UpvoteServicePromiseClient.prototype.upvote =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/upvote.v1.UpvoteService/Upvote',
      request,
      metadata || {},
      methodDescriptor_UpvoteService_Upvote);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.upvote.v1.DownvoteRequest,
 *   !proto.upvote.v1.DownvoteResponse>}
 */
const methodDescriptor_UpvoteService_Downvote = new grpc.web.MethodDescriptor(
  '/upvote.v1.UpvoteService/Downvote',
  grpc.web.MethodType.UNARY,
  proto.upvote.v1.DownvoteRequest,
  proto.upvote.v1.DownvoteResponse,
  /**
   * @param {!proto.upvote.v1.DownvoteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.upvote.v1.DownvoteResponse.deserializeBinary
);


/**
 * @param {!proto.upvote.v1.DownvoteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.upvote.v1.DownvoteResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.upvote.v1.DownvoteResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.upvote.v1.UpvoteServiceClient.prototype.downvote =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/upvote.v1.UpvoteService/Downvote',
      request,
      metadata || {},
      methodDescriptor_UpvoteService_Downvote,
      callback);
};


/**
 * @param {!proto.upvote.v1.DownvoteRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.upvote.v1.DownvoteResponse>}
 *     Promise that resolves to the response
 */
proto.upvote.v1.UpvoteServicePromiseClient.prototype.downvote =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/upvote.v1.UpvoteService/Downvote',
      request,
      metadata || {},
      methodDescriptor_UpvoteService_Downvote);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.upvote.v1.WatchBookRequest,
 *   !proto.upvote.v1.WatchBookResponse>}
 */
const methodDescriptor_UpvoteService_WatchBook = new grpc.web.MethodDescriptor(
  '/upvote.v1.UpvoteService/WatchBook',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.upvote.v1.WatchBookRequest,
  proto.upvote.v1.WatchBookResponse,
  /**
   * @param {!proto.upvote.v1.WatchBookRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.upvote.v1.WatchBookResponse.deserializeBinary
);


/**
 * @param {!proto.upvote.v1.WatchBookRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.upvote.v1.WatchBookResponse>}
 *     The XHR Node Readable Stream
 */
proto.upvote.v1.UpvoteServiceClient.prototype.watchBook =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/upvote.v1.UpvoteService/WatchBook',
      request,
      metadata || {},
      methodDescriptor_UpvoteService_WatchBook);
};


/**
 * @param {!proto.upvote.v1.WatchBookRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.upvote.v1.WatchBookResponse>}
 *     The XHR Node Readable Stream
 */
proto.upvote.v1.UpvoteServicePromiseClient.prototype.watchBook =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/upvote.v1.UpvoteService/WatchBook',
      request,
      metadata || {},
      methodDescriptor_UpvoteService_WatchBook);
};


module.exports = proto.upvote.v1;

