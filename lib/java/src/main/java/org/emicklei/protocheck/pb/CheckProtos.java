// Generated by the protocol buffer compiler.  DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: check.proto
// Protobuf Java Version: 4.29.3

package org.emicklei.protocheck.pb;

public final class CheckProtos {
  private CheckProtos() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 29,
      /* patch= */ 3,
      /* suffix= */ "",
      CheckProtos.class.getName());
  }
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
    registry.add(org.emicklei.protocheck.pb.CheckProtos.message);
    registry.add(org.emicklei.protocheck.pb.CheckProtos.field);
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public static final int MESSAGE_FIELD_NUMBER = 20030916;
  /**
   * <code>extend .google.protobuf.MessageOptions { ... }</code>
   */
  public static final
    com.google.protobuf.GeneratedMessage.GeneratedExtension<
      com.google.protobuf.DescriptorProtos.MessageOptions,
      java.util.List<org.emicklei.protocheck.pb.Check>> message = com.google.protobuf.GeneratedMessage
          .newFileScopedGeneratedExtension(
        org.emicklei.protocheck.pb.Check.class,
        org.emicklei.protocheck.pb.Check.getDefaultInstance());
  public static final int FIELD_FIELD_NUMBER = 20030916;
  /**
   * <code>extend .google.protobuf.FieldOptions { ... }</code>
   */
  public static final
    com.google.protobuf.GeneratedMessage.GeneratedExtension<
      com.google.protobuf.DescriptorProtos.FieldOptions,
      java.util.List<org.emicklei.protocheck.pb.Check>> field = com.google.protobuf.GeneratedMessage
          .newFileScopedGeneratedExtension(
        org.emicklei.protocheck.pb.Check.class,
        org.emicklei.protocheck.pb.Check.getDefaultInstance());
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_check_Check_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_check_Check_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_check_CheckError_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_check_CheckError_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\013check.proto\022\005check\032 google/protobuf/de" +
      "scriptor.proto\".\n\005Check\022\n\n\002id\030\001 \001(\t\022\014\n\004f" +
      "ail\030\002 \001(\t\022\013\n\003cel\030\003 \001(\t\"4\n\nCheckError\022\n\n\002" +
      "id\030\001 \001(\t\022\014\n\004fail\030\002 \001(\t\022\014\n\004path\030\003 \001(\t:A\n\007" +
      "message\022\037.google.protobuf.MessageOptions" +
      "\030\304\313\306\t \003(\0132\014.check.Check:=\n\005field\022\035.googl" +
      "e.protobuf.FieldOptions\030\304\313\306\t \003(\0132\014.check" +
      ".CheckBV\n\032org.emicklei.protocheck.pbB\013Ch" +
      "eckProtosP\001Z)github.com/emicklei/protoch" +
      "eck;protocheckb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.DescriptorProtos.getDescriptor(),
        });
    internal_static_check_Check_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_check_Check_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_check_Check_descriptor,
        new java.lang.String[] { "Id", "Fail", "Cel", });
    internal_static_check_CheckError_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_check_CheckError_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_check_CheckError_descriptor,
        new java.lang.String[] { "Id", "Fail", "Path", });
    message.internalInit(descriptor.getExtensions().get(0));
    field.internalInit(descriptor.getExtensions().get(1));
    descriptor.resolveAllFeaturesImmutable();
    com.google.protobuf.DescriptorProtos.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
