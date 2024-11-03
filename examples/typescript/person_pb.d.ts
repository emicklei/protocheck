// package: golang
// file: person.proto

import * as jspb from "google-protobuf";
import * as check_pb from "./check_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Person extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  hasDescription(): boolean;
  clearDescription(): void;
  getDescription(): string;
  setDescription(value: string): void;

  hasBirthDate(): boolean;
  clearBirthDate(): void;
  getBirthDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setBirthDate(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Person.AsObject;
  static toObject(includeInstance: boolean, msg: Person): Person.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Person, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Person;
  static deserializeBinaryFromReader(message: Person, reader: jspb.BinaryReader): Person;
}

export namespace Person {
  export type AsObject = {
    name: string,
    description: string,
    birthDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Pet extends jspb.Message {
  getKind(): string;
  setKind(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Pet.AsObject;
  static toObject(includeInstance: boolean, msg: Pet): Pet.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Pet, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Pet;
  static deserializeBinaryFromReader(message: Pet, reader: jspb.BinaryReader): Pet;
}

export namespace Pet {
  export type AsObject = {
    kind: string,
  }
}

