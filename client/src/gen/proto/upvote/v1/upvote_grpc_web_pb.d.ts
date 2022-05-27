import * as grpcWeb from 'grpc-web';

import * as proto_upvote_v1_upvote_pb from '../../../proto/upvote/v1/upvote_pb';


export class UpvoteServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  createBook(
    request: proto_upvote_v1_upvote_pb.CreateBookRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: proto_upvote_v1_upvote_pb.CreateBookResponse) => void
  ): grpcWeb.ClientReadableStream<proto_upvote_v1_upvote_pb.CreateBookResponse>;

  upvote(
    request: proto_upvote_v1_upvote_pb.UpvoteRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: proto_upvote_v1_upvote_pb.UpvoteResponse) => void
  ): grpcWeb.ClientReadableStream<proto_upvote_v1_upvote_pb.UpvoteResponse>;

  downvote(
    request: proto_upvote_v1_upvote_pb.DownvoteRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: proto_upvote_v1_upvote_pb.DownvoteResponse) => void
  ): grpcWeb.ClientReadableStream<proto_upvote_v1_upvote_pb.DownvoteResponse>;

  watchBook(
    request: proto_upvote_v1_upvote_pb.WatchBookRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<proto_upvote_v1_upvote_pb.WatchBookResponse>;

}

export class UpvoteServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  createBook(
    request: proto_upvote_v1_upvote_pb.CreateBookRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<proto_upvote_v1_upvote_pb.CreateBookResponse>;

  upvote(
    request: proto_upvote_v1_upvote_pb.UpvoteRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<proto_upvote_v1_upvote_pb.UpvoteResponse>;

  downvote(
    request: proto_upvote_v1_upvote_pb.DownvoteRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<proto_upvote_v1_upvote_pb.DownvoteResponse>;

  watchBook(
    request: proto_upvote_v1_upvote_pb.WatchBookRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<proto_upvote_v1_upvote_pb.WatchBookResponse>;

}

