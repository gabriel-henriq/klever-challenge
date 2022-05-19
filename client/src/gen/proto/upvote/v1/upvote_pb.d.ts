import * as jspb from 'google-protobuf'



export class UpvoteRequest extends jspb.Message {
  getId(): string;
  setId(value: string): UpvoteRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpvoteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpvoteRequest): UpvoteRequest.AsObject;
  static serializeBinaryToWriter(message: UpvoteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpvoteRequest;
  static deserializeBinaryFromReader(message: UpvoteRequest, reader: jspb.BinaryReader): UpvoteRequest;
}

export namespace UpvoteRequest {
  export type AsObject = {
    id: string,
  }
}

export class UpvoteResponse extends jspb.Message {
  getId(): string;
  setId(value: string): UpvoteResponse;

  getTitle(): string;
  setTitle(value: string): UpvoteResponse;

  getAuthor(): string;
  setAuthor(value: string): UpvoteResponse;

  getLikes(): number;
  setLikes(value: number): UpvoteResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpvoteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpvoteResponse): UpvoteResponse.AsObject;
  static serializeBinaryToWriter(message: UpvoteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpvoteResponse;
  static deserializeBinaryFromReader(message: UpvoteResponse, reader: jspb.BinaryReader): UpvoteResponse;
}

export namespace UpvoteResponse {
  export type AsObject = {
    id: string,
    title: string,
    author: string,
    likes: number,
  }
}

export class WatchBookRequest extends jspb.Message {
  getTitle(): string;
  setTitle(value: string): WatchBookRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WatchBookRequest.AsObject;
  static toObject(includeInstance: boolean, msg: WatchBookRequest): WatchBookRequest.AsObject;
  static serializeBinaryToWriter(message: WatchBookRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WatchBookRequest;
  static deserializeBinaryFromReader(message: WatchBookRequest, reader: jspb.BinaryReader): WatchBookRequest;
}

export namespace WatchBookRequest {
  export type AsObject = {
    title: string,
  }
}

export class WatchBookResponse extends jspb.Message {
  getId(): string;
  setId(value: string): WatchBookResponse;

  getTitle(): string;
  setTitle(value: string): WatchBookResponse;

  getAuthor(): string;
  setAuthor(value: string): WatchBookResponse;

  getLikes(): number;
  setLikes(value: number): WatchBookResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WatchBookResponse.AsObject;
  static toObject(includeInstance: boolean, msg: WatchBookResponse): WatchBookResponse.AsObject;
  static serializeBinaryToWriter(message: WatchBookResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WatchBookResponse;
  static deserializeBinaryFromReader(message: WatchBookResponse, reader: jspb.BinaryReader): WatchBookResponse;
}

export namespace WatchBookResponse {
  export type AsObject = {
    id: string,
    title: string,
    author: string,
    likes: number,
  }
}

