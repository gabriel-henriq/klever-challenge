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

export class ListBooksRequest extends jspb.Message {
  getTitle(): string;
  setTitle(value: string): ListBooksRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListBooksRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListBooksRequest): ListBooksRequest.AsObject;
  static serializeBinaryToWriter(message: ListBooksRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListBooksRequest;
  static deserializeBinaryFromReader(message: ListBooksRequest, reader: jspb.BinaryReader): ListBooksRequest;
}

export namespace ListBooksRequest {
  export type AsObject = {
    title: string,
  }
}

export class ListBooksResponse extends jspb.Message {
  getId(): string;
  setId(value: string): ListBooksResponse;

  getTitle(): string;
  setTitle(value: string): ListBooksResponse;

  getAuthor(): string;
  setAuthor(value: string): ListBooksResponse;

  getLikes(): number;
  setLikes(value: number): ListBooksResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListBooksResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListBooksResponse): ListBooksResponse.AsObject;
  static serializeBinaryToWriter(message: ListBooksResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListBooksResponse;
  static deserializeBinaryFromReader(message: ListBooksResponse, reader: jspb.BinaryReader): ListBooksResponse;
}

export namespace ListBooksResponse {
  export type AsObject = {
    id: string,
    title: string,
    author: string,
    likes: number,
  }
}

