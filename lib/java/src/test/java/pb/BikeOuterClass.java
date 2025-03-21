// Generated by the protocol buffer compiler.  DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: bike.proto
// Protobuf Java Version: 4.29.3

package pb;

public final class BikeOuterClass {
  private BikeOuterClass() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 29,
      /* patch= */ 3,
      /* suffix= */ "",
      BikeOuterClass.class.getName());
  }
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface BikeOrBuilder extends
      // @@protoc_insertion_point(interface_extends:tests.Bike)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>string brand = 1;</code>
     * @return The brand.
     */
    java.lang.String getBrand();
    /**
     * <code>string brand = 1;</code>
     * @return The bytes for brand.
     */
    com.google.protobuf.ByteString
        getBrandBytes();
  }
  /**
   * Protobuf type {@code tests.Bike}
   */
  public static final class Bike extends
      com.google.protobuf.GeneratedMessage implements
      // @@protoc_insertion_point(message_implements:tests.Bike)
      BikeOrBuilder {
  private static final long serialVersionUID = 0L;
    static {
      com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
        com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
        /* major= */ 4,
        /* minor= */ 29,
        /* patch= */ 3,
        /* suffix= */ "",
        Bike.class.getName());
    }
    // Use Bike.newBuilder() to construct.
    private Bike(com.google.protobuf.GeneratedMessage.Builder<?> builder) {
      super(builder);
    }
    private Bike() {
      brand_ = "";
    }

    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return pb.BikeOuterClass.internal_static_tests_Bike_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return pb.BikeOuterClass.internal_static_tests_Bike_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              pb.BikeOuterClass.Bike.class, pb.BikeOuterClass.Bike.Builder.class);
    }

    public static final int BRAND_FIELD_NUMBER = 1;
    @SuppressWarnings("serial")
    private volatile java.lang.Object brand_ = "";
    /**
     * <code>string brand = 1;</code>
     * @return The brand.
     */
    @java.lang.Override
    public java.lang.String getBrand() {
      java.lang.Object ref = brand_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        brand_ = s;
        return s;
      }
    }
    /**
     * <code>string brand = 1;</code>
     * @return The bytes for brand.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getBrandBytes() {
      java.lang.Object ref = brand_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        brand_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    private byte memoizedIsInitialized = -1;
    @java.lang.Override
    public final boolean isInitialized() {
      byte isInitialized = memoizedIsInitialized;
      if (isInitialized == 1) return true;
      if (isInitialized == 0) return false;

      memoizedIsInitialized = 1;
      return true;
    }

    @java.lang.Override
    public void writeTo(com.google.protobuf.CodedOutputStream output)
                        throws java.io.IOException {
      if (!com.google.protobuf.GeneratedMessage.isStringEmpty(brand_)) {
        com.google.protobuf.GeneratedMessage.writeString(output, 1, brand_);
      }
      getUnknownFields().writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (!com.google.protobuf.GeneratedMessage.isStringEmpty(brand_)) {
        size += com.google.protobuf.GeneratedMessage.computeStringSize(1, brand_);
      }
      size += getUnknownFields().getSerializedSize();
      memoizedSize = size;
      return size;
    }

    @java.lang.Override
    public boolean equals(final java.lang.Object obj) {
      if (obj == this) {
       return true;
      }
      if (!(obj instanceof pb.BikeOuterClass.Bike)) {
        return super.equals(obj);
      }
      pb.BikeOuterClass.Bike other = (pb.BikeOuterClass.Bike) obj;

      if (!getBrand()
          .equals(other.getBrand())) return false;
      if (!getUnknownFields().equals(other.getUnknownFields())) return false;
      return true;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptor().hashCode();
      hash = (37 * hash) + BRAND_FIELD_NUMBER;
      hash = (53 * hash) + getBrand().hashCode();
      hash = (29 * hash) + getUnknownFields().hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static pb.BikeOuterClass.Bike parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static pb.BikeOuterClass.Bike parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static pb.BikeOuterClass.Bike parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static pb.BikeOuterClass.Bike parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static pb.BikeOuterClass.Bike parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static pb.BikeOuterClass.Bike parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static pb.BikeOuterClass.Bike parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input);
    }
    public static pb.BikeOuterClass.Bike parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    public static pb.BikeOuterClass.Bike parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseDelimitedWithIOException(PARSER, input);
    }

    public static pb.BikeOuterClass.Bike parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static pb.BikeOuterClass.Bike parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input);
    }
    public static pb.BikeOuterClass.Bike parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    @java.lang.Override
    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(pb.BikeOuterClass.Bike prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
    @java.lang.Override
    public Builder toBuilder() {
      return this == DEFAULT_INSTANCE
          ? new Builder() : new Builder().mergeFrom(this);
    }

    @java.lang.Override
    protected Builder newBuilderForType(
        com.google.protobuf.GeneratedMessage.BuilderParent parent) {
      Builder builder = new Builder(parent);
      return builder;
    }
    /**
     * Protobuf type {@code tests.Bike}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessage.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:tests.Bike)
        pb.BikeOuterClass.BikeOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return pb.BikeOuterClass.internal_static_tests_Bike_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return pb.BikeOuterClass.internal_static_tests_Bike_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                pb.BikeOuterClass.Bike.class, pb.BikeOuterClass.Bike.Builder.class);
      }

      // Construct using pb.BikeOuterClass.Bike.newBuilder()
      private Builder() {

      }

      private Builder(
          com.google.protobuf.GeneratedMessage.BuilderParent parent) {
        super(parent);

      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        bitField0_ = 0;
        brand_ = "";
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return pb.BikeOuterClass.internal_static_tests_Bike_descriptor;
      }

      @java.lang.Override
      public pb.BikeOuterClass.Bike getDefaultInstanceForType() {
        return pb.BikeOuterClass.Bike.getDefaultInstance();
      }

      @java.lang.Override
      public pb.BikeOuterClass.Bike build() {
        pb.BikeOuterClass.Bike result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public pb.BikeOuterClass.Bike buildPartial() {
        pb.BikeOuterClass.Bike result = new pb.BikeOuterClass.Bike(this);
        if (bitField0_ != 0) { buildPartial0(result); }
        onBuilt();
        return result;
      }

      private void buildPartial0(pb.BikeOuterClass.Bike result) {
        int from_bitField0_ = bitField0_;
        if (((from_bitField0_ & 0x00000001) != 0)) {
          result.brand_ = brand_;
        }
      }

      @java.lang.Override
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof pb.BikeOuterClass.Bike) {
          return mergeFrom((pb.BikeOuterClass.Bike)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(pb.BikeOuterClass.Bike other) {
        if (other == pb.BikeOuterClass.Bike.getDefaultInstance()) return this;
        if (!other.getBrand().isEmpty()) {
          brand_ = other.brand_;
          bitField0_ |= 0x00000001;
          onChanged();
        }
        this.mergeUnknownFields(other.getUnknownFields());
        onChanged();
        return this;
      }

      @java.lang.Override
      public final boolean isInitialized() {
        return true;
      }

      @java.lang.Override
      public Builder mergeFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws java.io.IOException {
        if (extensionRegistry == null) {
          throw new java.lang.NullPointerException();
        }
        try {
          boolean done = false;
          while (!done) {
            int tag = input.readTag();
            switch (tag) {
              case 0:
                done = true;
                break;
              case 10: {
                brand_ = input.readStringRequireUtf8();
                bitField0_ |= 0x00000001;
                break;
              } // case 10
              default: {
                if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                  done = true; // was an endgroup tag
                }
                break;
              } // default:
            } // switch (tag)
          } // while (!done)
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          throw e.unwrapIOException();
        } finally {
          onChanged();
        } // finally
        return this;
      }
      private int bitField0_;

      private java.lang.Object brand_ = "";
      /**
       * <code>string brand = 1;</code>
       * @return The brand.
       */
      public java.lang.String getBrand() {
        java.lang.Object ref = brand_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          brand_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string brand = 1;</code>
       * @return The bytes for brand.
       */
      public com.google.protobuf.ByteString
          getBrandBytes() {
        java.lang.Object ref = brand_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          brand_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string brand = 1;</code>
       * @param value The brand to set.
       * @return This builder for chaining.
       */
      public Builder setBrand(
          java.lang.String value) {
        if (value == null) { throw new NullPointerException(); }
        brand_ = value;
        bitField0_ |= 0x00000001;
        onChanged();
        return this;
      }
      /**
       * <code>string brand = 1;</code>
       * @return This builder for chaining.
       */
      public Builder clearBrand() {
        brand_ = getDefaultInstance().getBrand();
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
        return this;
      }
      /**
       * <code>string brand = 1;</code>
       * @param value The bytes for brand to set.
       * @return This builder for chaining.
       */
      public Builder setBrandBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) { throw new NullPointerException(); }
        checkByteStringIsUtf8(value);
        brand_ = value;
        bitField0_ |= 0x00000001;
        onChanged();
        return this;
      }

      // @@protoc_insertion_point(builder_scope:tests.Bike)
    }

    // @@protoc_insertion_point(class_scope:tests.Bike)
    private static final pb.BikeOuterClass.Bike DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new pb.BikeOuterClass.Bike();
    }

    public static pb.BikeOuterClass.Bike getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<Bike>
        PARSER = new com.google.protobuf.AbstractParser<Bike>() {
      @java.lang.Override
      public Bike parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        Builder builder = newBuilder();
        try {
          builder.mergeFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          throw e.setUnfinishedMessage(builder.buildPartial());
        } catch (com.google.protobuf.UninitializedMessageException e) {
          throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
        } catch (java.io.IOException e) {
          throw new com.google.protobuf.InvalidProtocolBufferException(e)
              .setUnfinishedMessage(builder.buildPartial());
        }
        return builder.buildPartial();
      }
    };

    public static com.google.protobuf.Parser<Bike> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<Bike> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public pb.BikeOuterClass.Bike getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_tests_Bike_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_tests_Bike_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\nbike.proto\022\005tests\"\025\n\004Bike\022\r\n\005brand\030\001 \001" +
      "(\tB\006\n\002pbP\000b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_tests_Bike_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_tests_Bike_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_tests_Bike_descriptor,
        new java.lang.String[] { "Brand", });
    descriptor.resolveAllFeaturesImmutable();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
