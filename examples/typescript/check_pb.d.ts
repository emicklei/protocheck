// package: check
// file: check.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_descriptor_pb from "google-protobuf/google/protobuf/descriptor_pb";

export class Check extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getFail(): string;
  setFail(value: string): void;

  getCel(): string;
  setCel(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Check.AsObject;
  static toObject(includeInstance: boolean, msg: Check): Check.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Check, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Check;
  static deserializeBinaryFromReader(message: Check, reader: jspb.BinaryReader): Check;
}

export namespace Check {
  export type AsObject = {
    id: string,
    fail: string,
    cel: string,
  }
}

  export const message: jspb.ExtensionFieldInfo<Check>;

  export const field: jspb.ExtensionFieldInfo<Check>;

