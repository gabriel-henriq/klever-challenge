import * as grpcWeb from 'grpc-web';

import * as proto_upvote_v1_upvote_pb from '../../../proto/upvote/v1/upvote_pb';


export class UpvoteServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  listBooks(
    request: proto_upvote_v1_upvote_pb.ListBooksRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<proto_upvote_v1_upvote_pb.ListBooksResponse>;

  upvote(
    request: proto_upvote_v1_upvote_pb.UpvoteRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: proto_upvote_v1_upvote_pb.UpvoteResponse) => void
  ): grpcWeb.ClientReadableStream<proto_upvote_v1_upvote_pb.UpvoteResponse>;

}

export class UpvoteServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  listBooks(
    request: proto_upvote_v1_upvote_pb.ListBooksRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<proto_upvote_v1_upvote_pb.ListBooksResponse>;

  upvote(
    request: proto_upvote_v1_upvote_pb.UpvoteRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<proto_upvote_v1_upvote_pb.UpvoteResponse>;

}

