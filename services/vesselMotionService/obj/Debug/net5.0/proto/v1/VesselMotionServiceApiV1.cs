// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: proto/v1/vessel_motion_service_api_v1.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace vesselMotionService {

  /// <summary>Holder for reflection information generated from proto/v1/vessel_motion_service_api_v1.proto</summary>
  public static partial class VesselMotionServiceApiV1Reflection {

    #region Descriptor
    /// <summary>File descriptor for proto/v1/vessel_motion_service_api_v1.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static VesselMotionServiceApiV1Reflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Citwcm90by92MS92ZXNzZWxfbW90aW9uX3NlcnZpY2VfYXBpX3YxLnByb3Rv",
            "EhZ2ZXNzZWxNb3Rpb25TZXJ2aWNlLnYxIs8BChVNb3Rpb25Fc3RpbWF0ZVJl",
            "cXVlc3QSHQoVcG9ydF9wcm9wX21vdG9yX3Bvd2VyGAEgAygCEhsKE3dpbmRf",
            "c3BlZWRfcmVsYXRpdmUYAiADKAISEAoIbGF0aXR1ZGUYAyADKAISDwoHSGVh",
            "ZGluZxgEIAMoAhITCgt3YXZlX2hlaWdodBgFIAMoAhJCCg5xdWVyeV9sb2Nh",
            "dGlvbhgGIAEoDjIqLnZlc3NlbE1vdGlvblNlcnZpY2UudjEuTG9jYXRpb25P",
            "blNoaXBFbnVtIhcKFU1vdGlvblRyYWNraW5nUmVxdWVzdCImCg5Nb3Rpb25S",
            "ZXNwb25zZRIUCgxhY2NlbGVyYXRpb24YASADKAIiVgoYTW90aW9uRXZhbHVh",
            "dGlvblJlc3BvbnNlEh0KFWFjY2VsZXJhdGlvbl9lc3RpbWF0ZRgBIAMoAhIb",
            "ChNhY2NlbGVyYXRpb25fYWN0dWFsGAIgAygCKjAKEkxvY2F0aW9uT25TaGlw",
            "RW51bRILCgdVTktOT1dOEAASDQoJU09NRVdIRVJFEAEy5AIKE3Zlc3NlbE1v",
            "dGlvblNlcnZpY2USZwoOTW90aW9uRXN0aW1hdGUSLS52ZXNzZWxNb3Rpb25T",
            "ZXJ2aWNlLnYxLk1vdGlvbkVzdGltYXRlUmVxdWVzdBomLnZlc3NlbE1vdGlv",
            "blNlcnZpY2UudjEuTW90aW9uUmVzcG9uc2USZwoOTW90aW9uVHJhY2tpbmcS",
            "LS52ZXNzZWxNb3Rpb25TZXJ2aWNlLnYxLk1vdGlvblRyYWNraW5nUmVxdWVz",
            "dBomLnZlc3NlbE1vdGlvblNlcnZpY2UudjEuTW90aW9uUmVzcG9uc2USewoY",
            "TW90aW9uRXN0aW1hdGVFdmFsdWF0aW9uEi0udmVzc2VsTW90aW9uU2Vydmlj",
            "ZS52MS5Nb3Rpb25Fc3RpbWF0ZVJlcXVlc3QaMC52ZXNzZWxNb3Rpb25TZXJ2",
            "aWNlLnYxLk1vdGlvbkV2YWx1YXRpb25SZXNwb25zZUIWqgITdmVzc2VsTW90",
            "aW9uU2VydmljZWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::vesselMotionService.LocationOnShipEnum), }, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::vesselMotionService.MotionEstimateRequest), global::vesselMotionService.MotionEstimateRequest.Parser, new[]{ "PortPropMotorPower", "WindSpeedRelative", "Latitude", "Heading", "WaveHeight", "QueryLocation" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::vesselMotionService.MotionTrackingRequest), global::vesselMotionService.MotionTrackingRequest.Parser, null, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::vesselMotionService.MotionResponse), global::vesselMotionService.MotionResponse.Parser, new[]{ "Acceleration" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::vesselMotionService.MotionEvaluationResponse), global::vesselMotionService.MotionEvaluationResponse.Parser, new[]{ "AccelerationEstimate", "AccelerationActual" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Enums
  public enum LocationOnShipEnum {
    [pbr::OriginalName("UNKNOWN")] Unknown = 0,
    [pbr::OriginalName("SOMEWHERE")] Somewhere = 1,
  }

  #endregion

  #region Messages
  public sealed partial class MotionEstimateRequest : pb::IMessage<MotionEstimateRequest>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<MotionEstimateRequest> _parser = new pb::MessageParser<MotionEstimateRequest>(() => new MotionEstimateRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<MotionEstimateRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::vesselMotionService.VesselMotionServiceApiV1Reflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public MotionEstimateRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public MotionEstimateRequest(MotionEstimateRequest other) : this() {
      portPropMotorPower_ = other.portPropMotorPower_.Clone();
      windSpeedRelative_ = other.windSpeedRelative_.Clone();
      latitude_ = other.latitude_.Clone();
      heading_ = other.heading_.Clone();
      waveHeight_ = other.waveHeight_.Clone();
      queryLocation_ = other.queryLocation_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public MotionEstimateRequest Clone() {
      return new MotionEstimateRequest(this);
    }

    /// <summary>Field number for the "port_prop_motor_power" field.</summary>
    public const int PortPropMotorPowerFieldNumber = 1;
    private static readonly pb::FieldCodec<float> _repeated_portPropMotorPower_codec
        = pb::FieldCodec.ForFloat(10);
    private readonly pbc::RepeatedField<float> portPropMotorPower_ = new pbc::RepeatedField<float>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<float> PortPropMotorPower {
      get { return portPropMotorPower_; }
    }

    /// <summary>Field number for the "wind_speed_relative" field.</summary>
    public const int WindSpeedRelativeFieldNumber = 2;
    private static readonly pb::FieldCodec<float> _repeated_windSpeedRelative_codec
        = pb::FieldCodec.ForFloat(18);
    private readonly pbc::RepeatedField<float> windSpeedRelative_ = new pbc::RepeatedField<float>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<float> WindSpeedRelative {
      get { return windSpeedRelative_; }
    }

    /// <summary>Field number for the "latitude" field.</summary>
    public const int LatitudeFieldNumber = 3;
    private static readonly pb::FieldCodec<float> _repeated_latitude_codec
        = pb::FieldCodec.ForFloat(26);
    private readonly pbc::RepeatedField<float> latitude_ = new pbc::RepeatedField<float>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<float> Latitude {
      get { return latitude_; }
    }

    /// <summary>Field number for the "Heading" field.</summary>
    public const int HeadingFieldNumber = 4;
    private static readonly pb::FieldCodec<float> _repeated_heading_codec
        = pb::FieldCodec.ForFloat(34);
    private readonly pbc::RepeatedField<float> heading_ = new pbc::RepeatedField<float>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<float> Heading {
      get { return heading_; }
    }

    /// <summary>Field number for the "wave_height" field.</summary>
    public const int WaveHeightFieldNumber = 5;
    private static readonly pb::FieldCodec<float> _repeated_waveHeight_codec
        = pb::FieldCodec.ForFloat(42);
    private readonly pbc::RepeatedField<float> waveHeight_ = new pbc::RepeatedField<float>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<float> WaveHeight {
      get { return waveHeight_; }
    }

    /// <summary>Field number for the "query_location" field.</summary>
    public const int QueryLocationFieldNumber = 6;
    private global::vesselMotionService.LocationOnShipEnum queryLocation_ = global::vesselMotionService.LocationOnShipEnum.Unknown;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::vesselMotionService.LocationOnShipEnum QueryLocation {
      get { return queryLocation_; }
      set {
        queryLocation_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as MotionEstimateRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(MotionEstimateRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!portPropMotorPower_.Equals(other.portPropMotorPower_)) return false;
      if(!windSpeedRelative_.Equals(other.windSpeedRelative_)) return false;
      if(!latitude_.Equals(other.latitude_)) return false;
      if(!heading_.Equals(other.heading_)) return false;
      if(!waveHeight_.Equals(other.waveHeight_)) return false;
      if (QueryLocation != other.QueryLocation) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= portPropMotorPower_.GetHashCode();
      hash ^= windSpeedRelative_.GetHashCode();
      hash ^= latitude_.GetHashCode();
      hash ^= heading_.GetHashCode();
      hash ^= waveHeight_.GetHashCode();
      if (QueryLocation != global::vesselMotionService.LocationOnShipEnum.Unknown) hash ^= QueryLocation.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      portPropMotorPower_.WriteTo(output, _repeated_portPropMotorPower_codec);
      windSpeedRelative_.WriteTo(output, _repeated_windSpeedRelative_codec);
      latitude_.WriteTo(output, _repeated_latitude_codec);
      heading_.WriteTo(output, _repeated_heading_codec);
      waveHeight_.WriteTo(output, _repeated_waveHeight_codec);
      if (QueryLocation != global::vesselMotionService.LocationOnShipEnum.Unknown) {
        output.WriteRawTag(48);
        output.WriteEnum((int) QueryLocation);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      portPropMotorPower_.WriteTo(ref output, _repeated_portPropMotorPower_codec);
      windSpeedRelative_.WriteTo(ref output, _repeated_windSpeedRelative_codec);
      latitude_.WriteTo(ref output, _repeated_latitude_codec);
      heading_.WriteTo(ref output, _repeated_heading_codec);
      waveHeight_.WriteTo(ref output, _repeated_waveHeight_codec);
      if (QueryLocation != global::vesselMotionService.LocationOnShipEnum.Unknown) {
        output.WriteRawTag(48);
        output.WriteEnum((int) QueryLocation);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      size += portPropMotorPower_.CalculateSize(_repeated_portPropMotorPower_codec);
      size += windSpeedRelative_.CalculateSize(_repeated_windSpeedRelative_codec);
      size += latitude_.CalculateSize(_repeated_latitude_codec);
      size += heading_.CalculateSize(_repeated_heading_codec);
      size += waveHeight_.CalculateSize(_repeated_waveHeight_codec);
      if (QueryLocation != global::vesselMotionService.LocationOnShipEnum.Unknown) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) QueryLocation);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(MotionEstimateRequest other) {
      if (other == null) {
        return;
      }
      portPropMotorPower_.Add(other.portPropMotorPower_);
      windSpeedRelative_.Add(other.windSpeedRelative_);
      latitude_.Add(other.latitude_);
      heading_.Add(other.heading_);
      waveHeight_.Add(other.waveHeight_);
      if (other.QueryLocation != global::vesselMotionService.LocationOnShipEnum.Unknown) {
        QueryLocation = other.QueryLocation;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10:
          case 13: {
            portPropMotorPower_.AddEntriesFrom(input, _repeated_portPropMotorPower_codec);
            break;
          }
          case 18:
          case 21: {
            windSpeedRelative_.AddEntriesFrom(input, _repeated_windSpeedRelative_codec);
            break;
          }
          case 26:
          case 29: {
            latitude_.AddEntriesFrom(input, _repeated_latitude_codec);
            break;
          }
          case 34:
          case 37: {
            heading_.AddEntriesFrom(input, _repeated_heading_codec);
            break;
          }
          case 42:
          case 45: {
            waveHeight_.AddEntriesFrom(input, _repeated_waveHeight_codec);
            break;
          }
          case 48: {
            QueryLocation = (global::vesselMotionService.LocationOnShipEnum) input.ReadEnum();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10:
          case 13: {
            portPropMotorPower_.AddEntriesFrom(ref input, _repeated_portPropMotorPower_codec);
            break;
          }
          case 18:
          case 21: {
            windSpeedRelative_.AddEntriesFrom(ref input, _repeated_windSpeedRelative_codec);
            break;
          }
          case 26:
          case 29: {
            latitude_.AddEntriesFrom(ref input, _repeated_latitude_codec);
            break;
          }
          case 34:
          case 37: {
            heading_.AddEntriesFrom(ref input, _repeated_heading_codec);
            break;
          }
          case 42:
          case 45: {
            waveHeight_.AddEntriesFrom(ref input, _repeated_waveHeight_codec);
            break;
          }
          case 48: {
            QueryLocation = (global::vesselMotionService.LocationOnShipEnum) input.ReadEnum();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class MotionTrackingRequest : pb::IMessage<MotionTrackingRequest>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<MotionTrackingRequest> _parser = new pb::MessageParser<MotionTrackingRequest>(() => new MotionTrackingRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<MotionTrackingRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::vesselMotionService.VesselMotionServiceApiV1Reflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public MotionTrackingRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public MotionTrackingRequest(MotionTrackingRequest other) : this() {
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public MotionTrackingRequest Clone() {
      return new MotionTrackingRequest(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as MotionTrackingRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(MotionTrackingRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(MotionTrackingRequest other) {
      if (other == null) {
        return;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
        }
      }
    }
    #endif

  }

  public sealed partial class MotionResponse : pb::IMessage<MotionResponse>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<MotionResponse> _parser = new pb::MessageParser<MotionResponse>(() => new MotionResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<MotionResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::vesselMotionService.VesselMotionServiceApiV1Reflection.Descriptor.MessageTypes[2]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public MotionResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public MotionResponse(MotionResponse other) : this() {
      acceleration_ = other.acceleration_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public MotionResponse Clone() {
      return new MotionResponse(this);
    }

    /// <summary>Field number for the "acceleration" field.</summary>
    public const int AccelerationFieldNumber = 1;
    private static readonly pb::FieldCodec<float> _repeated_acceleration_codec
        = pb::FieldCodec.ForFloat(10);
    private readonly pbc::RepeatedField<float> acceleration_ = new pbc::RepeatedField<float>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<float> Acceleration {
      get { return acceleration_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as MotionResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(MotionResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!acceleration_.Equals(other.acceleration_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= acceleration_.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      acceleration_.WriteTo(output, _repeated_acceleration_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      acceleration_.WriteTo(ref output, _repeated_acceleration_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      size += acceleration_.CalculateSize(_repeated_acceleration_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(MotionResponse other) {
      if (other == null) {
        return;
      }
      acceleration_.Add(other.acceleration_);
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10:
          case 13: {
            acceleration_.AddEntriesFrom(input, _repeated_acceleration_codec);
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10:
          case 13: {
            acceleration_.AddEntriesFrom(ref input, _repeated_acceleration_codec);
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class MotionEvaluationResponse : pb::IMessage<MotionEvaluationResponse>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<MotionEvaluationResponse> _parser = new pb::MessageParser<MotionEvaluationResponse>(() => new MotionEvaluationResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<MotionEvaluationResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::vesselMotionService.VesselMotionServiceApiV1Reflection.Descriptor.MessageTypes[3]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public MotionEvaluationResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public MotionEvaluationResponse(MotionEvaluationResponse other) : this() {
      accelerationEstimate_ = other.accelerationEstimate_.Clone();
      accelerationActual_ = other.accelerationActual_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public MotionEvaluationResponse Clone() {
      return new MotionEvaluationResponse(this);
    }

    /// <summary>Field number for the "acceleration_estimate" field.</summary>
    public const int AccelerationEstimateFieldNumber = 1;
    private static readonly pb::FieldCodec<float> _repeated_accelerationEstimate_codec
        = pb::FieldCodec.ForFloat(10);
    private readonly pbc::RepeatedField<float> accelerationEstimate_ = new pbc::RepeatedField<float>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<float> AccelerationEstimate {
      get { return accelerationEstimate_; }
    }

    /// <summary>Field number for the "acceleration_actual" field.</summary>
    public const int AccelerationActualFieldNumber = 2;
    private static readonly pb::FieldCodec<float> _repeated_accelerationActual_codec
        = pb::FieldCodec.ForFloat(18);
    private readonly pbc::RepeatedField<float> accelerationActual_ = new pbc::RepeatedField<float>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<float> AccelerationActual {
      get { return accelerationActual_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as MotionEvaluationResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(MotionEvaluationResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!accelerationEstimate_.Equals(other.accelerationEstimate_)) return false;
      if(!accelerationActual_.Equals(other.accelerationActual_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= accelerationEstimate_.GetHashCode();
      hash ^= accelerationActual_.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      accelerationEstimate_.WriteTo(output, _repeated_accelerationEstimate_codec);
      accelerationActual_.WriteTo(output, _repeated_accelerationActual_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      accelerationEstimate_.WriteTo(ref output, _repeated_accelerationEstimate_codec);
      accelerationActual_.WriteTo(ref output, _repeated_accelerationActual_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      size += accelerationEstimate_.CalculateSize(_repeated_accelerationEstimate_codec);
      size += accelerationActual_.CalculateSize(_repeated_accelerationActual_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(MotionEvaluationResponse other) {
      if (other == null) {
        return;
      }
      accelerationEstimate_.Add(other.accelerationEstimate_);
      accelerationActual_.Add(other.accelerationActual_);
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10:
          case 13: {
            accelerationEstimate_.AddEntriesFrom(input, _repeated_accelerationEstimate_codec);
            break;
          }
          case 18:
          case 21: {
            accelerationActual_.AddEntriesFrom(input, _repeated_accelerationActual_codec);
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10:
          case 13: {
            accelerationEstimate_.AddEntriesFrom(ref input, _repeated_accelerationEstimate_codec);
            break;
          }
          case 18:
          case 21: {
            accelerationActual_.AddEntriesFrom(ref input, _repeated_accelerationActual_codec);
            break;
          }
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code
