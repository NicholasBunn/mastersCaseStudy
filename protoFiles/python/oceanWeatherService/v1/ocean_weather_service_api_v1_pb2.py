# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ocean_weather_service_api_v1.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='ocean_weather_service_api_v1.proto',
  package='oceanWeatherServiceAPI.v1',
  syntax='proto3',
  serialized_options=b'ZRrouteAnalysisAggregator/proto/v1/generated/oceanWeatherService;oceanWeatherService',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\"ocean_weather_service_api_v1.proto\x12\x19oceanWeatherServiceAPI.v1\"W\n\x1dOceanWeatherPredictionRequest\x12\x10\n\x08latitude\x18\x01 \x03(\x02\x12\x11\n\tlongitude\x18\x02 \x03(\x02\x12\x11\n\tunix_time\x18\x03 \x03(\x01\"\x98\x01\n\x1aOceanWeatherHistoryRequest\x12\x10\n\x08latitude\x18\x01 \x03(\x02\x12\x11\n\tlongitude\x18\x02 \x03(\x02\x12\x11\n\tunix_time\x18\x03 \x03(\x01\x12\x42\n\x0f\x61rchive_service\x18\x04 \x01(\x0e\x32).oceanWeatherServiceAPI.v1.ArchiveService\"\xec\x01\n\x1fOceanWeatherInformationResponse\x12\x11\n\tunix_time\x18\x01 \x03(\x01\x12\x16\n\x0ewind_direction\x18\x02 \x03(\x02\x12\x12\n\nwind_speed\x18\x03 \x03(\x02\x12\x17\n\x0f\x62\x65\x61ufort_number\x18\x04 \x03(\r\x12\x17\n\x0fswell_direction\x18\x05 \x03(\x02\x12\x13\n\x0bwave_length\x18\x06 \x03(\x02\x12\x14\n\x0cswell_height\x18\x07 \x03(\x02\x12\x17\n\x0fswell_frequency\x18\x08 \x03(\x02\x12\x14\n\x0cswell_period\x18\t \x03(\x02*I\n\x0e\x41rchiveService\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0e\n\nSTORMGLASS\x10\x01\x12\x08\n\x04\x45RA5\x10\x02\x12\x10\n\x0cOBSERVATIONS\x10\x03\x32\xb1\x02\n\x13OceanWeatherService\x12\x8e\x01\n\x16OceanWeatherPrediction\x12\x38.oceanWeatherServiceAPI.v1.OceanWeatherPredictionRequest\x1a:.oceanWeatherServiceAPI.v1.OceanWeatherInformationResponse\x12\x88\x01\n\x13OceanWeatherHistory\x12\x35.oceanWeatherServiceAPI.v1.OceanWeatherHistoryRequest\x1a:.oceanWeatherServiceAPI.v1.OceanWeatherInformationResponseBTZRrouteAnalysisAggregator/proto/v1/generated/oceanWeatherService;oceanWeatherServiceb\x06proto3'
)

_ARCHIVESERVICE = _descriptor.EnumDescriptor(
  name='ArchiveService',
  full_name='oceanWeatherServiceAPI.v1.ArchiveService',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='STORMGLASS', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ERA5', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='OBSERVATIONS', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=548,
  serialized_end=621,
)
_sym_db.RegisterEnumDescriptor(_ARCHIVESERVICE)

ArchiveService = enum_type_wrapper.EnumTypeWrapper(_ARCHIVESERVICE)
UNKNOWN = 0
STORMGLASS = 1
ERA5 = 2
OBSERVATIONS = 3



_OCEANWEATHERPREDICTIONREQUEST = _descriptor.Descriptor(
  name='OceanWeatherPredictionRequest',
  full_name='oceanWeatherServiceAPI.v1.OceanWeatherPredictionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='latitude', full_name='oceanWeatherServiceAPI.v1.OceanWeatherPredictionRequest.latitude', index=0,
      number=1, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='longitude', full_name='oceanWeatherServiceAPI.v1.OceanWeatherPredictionRequest.longitude', index=1,
      number=2, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='unix_time', full_name='oceanWeatherServiceAPI.v1.OceanWeatherPredictionRequest.unix_time', index=2,
      number=3, type=1, cpp_type=5, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=65,
  serialized_end=152,
)


_OCEANWEATHERHISTORYREQUEST = _descriptor.Descriptor(
  name='OceanWeatherHistoryRequest',
  full_name='oceanWeatherServiceAPI.v1.OceanWeatherHistoryRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='latitude', full_name='oceanWeatherServiceAPI.v1.OceanWeatherHistoryRequest.latitude', index=0,
      number=1, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='longitude', full_name='oceanWeatherServiceAPI.v1.OceanWeatherHistoryRequest.longitude', index=1,
      number=2, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='unix_time', full_name='oceanWeatherServiceAPI.v1.OceanWeatherHistoryRequest.unix_time', index=2,
      number=3, type=1, cpp_type=5, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='archive_service', full_name='oceanWeatherServiceAPI.v1.OceanWeatherHistoryRequest.archive_service', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=155,
  serialized_end=307,
)


_OCEANWEATHERINFORMATIONRESPONSE = _descriptor.Descriptor(
  name='OceanWeatherInformationResponse',
  full_name='oceanWeatherServiceAPI.v1.OceanWeatherInformationResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='unix_time', full_name='oceanWeatherServiceAPI.v1.OceanWeatherInformationResponse.unix_time', index=0,
      number=1, type=1, cpp_type=5, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='wind_direction', full_name='oceanWeatherServiceAPI.v1.OceanWeatherInformationResponse.wind_direction', index=1,
      number=2, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='wind_speed', full_name='oceanWeatherServiceAPI.v1.OceanWeatherInformationResponse.wind_speed', index=2,
      number=3, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='beaufort_number', full_name='oceanWeatherServiceAPI.v1.OceanWeatherInformationResponse.beaufort_number', index=3,
      number=4, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='swell_direction', full_name='oceanWeatherServiceAPI.v1.OceanWeatherInformationResponse.swell_direction', index=4,
      number=5, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='wave_length', full_name='oceanWeatherServiceAPI.v1.OceanWeatherInformationResponse.wave_length', index=5,
      number=6, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='swell_height', full_name='oceanWeatherServiceAPI.v1.OceanWeatherInformationResponse.swell_height', index=6,
      number=7, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='swell_frequency', full_name='oceanWeatherServiceAPI.v1.OceanWeatherInformationResponse.swell_frequency', index=7,
      number=8, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='swell_period', full_name='oceanWeatherServiceAPI.v1.OceanWeatherInformationResponse.swell_period', index=8,
      number=9, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=310,
  serialized_end=546,
)

_OCEANWEATHERHISTORYREQUEST.fields_by_name['archive_service'].enum_type = _ARCHIVESERVICE
DESCRIPTOR.message_types_by_name['OceanWeatherPredictionRequest'] = _OCEANWEATHERPREDICTIONREQUEST
DESCRIPTOR.message_types_by_name['OceanWeatherHistoryRequest'] = _OCEANWEATHERHISTORYREQUEST
DESCRIPTOR.message_types_by_name['OceanWeatherInformationResponse'] = _OCEANWEATHERINFORMATIONRESPONSE
DESCRIPTOR.enum_types_by_name['ArchiveService'] = _ARCHIVESERVICE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

OceanWeatherPredictionRequest = _reflection.GeneratedProtocolMessageType('OceanWeatherPredictionRequest', (_message.Message,), {
  'DESCRIPTOR' : _OCEANWEATHERPREDICTIONREQUEST,
  '__module__' : 'ocean_weather_service_api_v1_pb2'
  # @@protoc_insertion_point(class_scope:oceanWeatherServiceAPI.v1.OceanWeatherPredictionRequest)
  })
_sym_db.RegisterMessage(OceanWeatherPredictionRequest)

OceanWeatherHistoryRequest = _reflection.GeneratedProtocolMessageType('OceanWeatherHistoryRequest', (_message.Message,), {
  'DESCRIPTOR' : _OCEANWEATHERHISTORYREQUEST,
  '__module__' : 'ocean_weather_service_api_v1_pb2'
  # @@protoc_insertion_point(class_scope:oceanWeatherServiceAPI.v1.OceanWeatherHistoryRequest)
  })
_sym_db.RegisterMessage(OceanWeatherHistoryRequest)

OceanWeatherInformationResponse = _reflection.GeneratedProtocolMessageType('OceanWeatherInformationResponse', (_message.Message,), {
  'DESCRIPTOR' : _OCEANWEATHERINFORMATIONRESPONSE,
  '__module__' : 'ocean_weather_service_api_v1_pb2'
  # @@protoc_insertion_point(class_scope:oceanWeatherServiceAPI.v1.OceanWeatherInformationResponse)
  })
_sym_db.RegisterMessage(OceanWeatherInformationResponse)


DESCRIPTOR._options = None

_OCEANWEATHERSERVICE = _descriptor.ServiceDescriptor(
  name='OceanWeatherService',
  full_name='oceanWeatherServiceAPI.v1.OceanWeatherService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=624,
  serialized_end=929,
  methods=[
  _descriptor.MethodDescriptor(
    name='OceanWeatherPrediction',
    full_name='oceanWeatherServiceAPI.v1.OceanWeatherService.OceanWeatherPrediction',
    index=0,
    containing_service=None,
    input_type=_OCEANWEATHERPREDICTIONREQUEST,
    output_type=_OCEANWEATHERINFORMATIONRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='OceanWeatherHistory',
    full_name='oceanWeatherServiceAPI.v1.OceanWeatherService.OceanWeatherHistory',
    index=1,
    containing_service=None,
    input_type=_OCEANWEATHERHISTORYREQUEST,
    output_type=_OCEANWEATHERINFORMATIONRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_OCEANWEATHERSERVICE)

DESCRIPTOR.services_by_name['OceanWeatherService'] = _OCEANWEATHERSERVICE

# @@protoc_insertion_point(module_scope)
