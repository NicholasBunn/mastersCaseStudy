# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: power_train_service_api_v1.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='power_train_service_api_v1.proto',
  package='powerTrainServiceAPI.v1',
  syntax='proto3',
  serialized_options=b'Z\'powerTrainService/v1/;powerTrainService',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n power_train_service_api_v1.proto\x12\x17powerTrainServiceAPI.v1\"\xec\x02\n\x19PowerTrainEstimateRequest\x12\x11\n\tunix_time\x18\x01 \x03(\x01\x12\x1d\n\x15port_prop_motor_speed\x18\x02 \x03(\x02\x12\x1d\n\x15stbd_prop_motor_speed\x18\x03 \x03(\x02\x12\x1c\n\x14propeller_pitch_port\x18\x04 \x03(\x02\x12\x1c\n\x14propeller_pitch_stbd\x18\x05 \x03(\x02\x12\x0b\n\x03sog\x18\x06 \x03(\x02\x12\x1f\n\x17wind_direction_relative\x18\x07 \x03(\x02\x12\x12\n\nwind_speed\x18\x08 \x03(\x02\x12\x17\n\x0f\x62\x65\x61ufort_number\x18\t \x03(\r\x12\x16\n\x0ewave_direction\x18\n \x03(\x02\x12\x13\n\x0bwave_length\x18\x0b \x03(\x02\x12:\n\nmodel_type\x18\x0c \x01(\x0e\x32&.powerTrainServiceAPI.v1.ModelTypeEnum\"\x16\n\x14PowerTrackingRequest\"b\n\x15PowerEstimateResponse\x12\x11\n\tunix_time\x18\x01 \x03(\x01\x12\x16\n\x0epower_estimate\x18\x02 \x03(\x02\x12\x1e\n\x16power_estimate_average\x18\x03 \x01(\x02\"\x8c\x01\n\x14\x43ostEstimateResponse\x12\x11\n\tunix_time\x18\x01 \x03(\x01\x12\x16\n\x0epower_estimate\x18\x02 \x03(\x02\x12\x15\n\rcost_estimate\x18\x03 \x03(\x02\x12\x12\n\ntotal_cost\x18\x04 \x01(\x02\x12\x1e\n\x16power_estimate_average\x18\x05 \x01(\x02\"Z\n\x17PowerEvaluationResponse\x12\x11\n\tunix_time\x18\x01 \x03(\x01\x12\x16\n\x0epower_estimate\x18\x02 \x03(\x02\x12\x14\n\x0cpower_actual\x18\x03 \x03(\x02\"-\n\x15PowerTrackingResponse\x12\x14\n\x0cpower_actual\x18\x01 \x03(\x02*4\n\rModelTypeEnum\x12\x0b\n\x07UNKNOWN\x10\x00\x12\r\n\tOPENWATER\x10\x01\x12\x07\n\x03ICE\x10\x02\x32\xec\x03\n\x11PowerTrainService\x12s\n\rPowerEstimate\x12\x32.powerTrainServiceAPI.v1.PowerTrainEstimateRequest\x1a..powerTrainServiceAPI.v1.PowerEstimateResponse\x12q\n\x0c\x43ostEstimate\x12\x32.powerTrainServiceAPI.v1.PowerTrainEstimateRequest\x1a-.powerTrainServiceAPI.v1.CostEstimateResponse\x12n\n\rPowerTracking\x12-.powerTrainServiceAPI.v1.PowerTrackingRequest\x1a..powerTrainServiceAPI.v1.PowerTrackingResponse\x12\x7f\n\x17PowerEstimateEvaluation\x12\x32.powerTrainServiceAPI.v1.PowerTrainEstimateRequest\x1a\x30.powerTrainServiceAPI.v1.PowerEvaluationResponseB)Z\'powerTrainService/v1/;powerTrainServiceb\x06proto3'
)

_MODELTYPEENUM = _descriptor.EnumDescriptor(
  name='ModelTypeEnum',
  full_name='powerTrainServiceAPI.v1.ModelTypeEnum',
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
      name='OPENWATER', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ICE', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=834,
  serialized_end=886,
)
_sym_db.RegisterEnumDescriptor(_MODELTYPEENUM)

ModelTypeEnum = enum_type_wrapper.EnumTypeWrapper(_MODELTYPEENUM)
UNKNOWN = 0
OPENWATER = 1
ICE = 2



_POWERTRAINESTIMATEREQUEST = _descriptor.Descriptor(
  name='PowerTrainEstimateRequest',
  full_name='powerTrainServiceAPI.v1.PowerTrainEstimateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='unix_time', full_name='powerTrainServiceAPI.v1.PowerTrainEstimateRequest.unix_time', index=0,
      number=1, type=1, cpp_type=5, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='port_prop_motor_speed', full_name='powerTrainServiceAPI.v1.PowerTrainEstimateRequest.port_prop_motor_speed', index=1,
      number=2, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='stbd_prop_motor_speed', full_name='powerTrainServiceAPI.v1.PowerTrainEstimateRequest.stbd_prop_motor_speed', index=2,
      number=3, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='propeller_pitch_port', full_name='powerTrainServiceAPI.v1.PowerTrainEstimateRequest.propeller_pitch_port', index=3,
      number=4, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='propeller_pitch_stbd', full_name='powerTrainServiceAPI.v1.PowerTrainEstimateRequest.propeller_pitch_stbd', index=4,
      number=5, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sog', full_name='powerTrainServiceAPI.v1.PowerTrainEstimateRequest.sog', index=5,
      number=6, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='wind_direction_relative', full_name='powerTrainServiceAPI.v1.PowerTrainEstimateRequest.wind_direction_relative', index=6,
      number=7, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='wind_speed', full_name='powerTrainServiceAPI.v1.PowerTrainEstimateRequest.wind_speed', index=7,
      number=8, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='beaufort_number', full_name='powerTrainServiceAPI.v1.PowerTrainEstimateRequest.beaufort_number', index=8,
      number=9, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='wave_direction', full_name='powerTrainServiceAPI.v1.PowerTrainEstimateRequest.wave_direction', index=9,
      number=10, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='wave_length', full_name='powerTrainServiceAPI.v1.PowerTrainEstimateRequest.wave_length', index=10,
      number=11, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='model_type', full_name='powerTrainServiceAPI.v1.PowerTrainEstimateRequest.model_type', index=11,
      number=12, type=14, cpp_type=8, label=1,
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
  serialized_start=62,
  serialized_end=426,
)


_POWERTRACKINGREQUEST = _descriptor.Descriptor(
  name='PowerTrackingRequest',
  full_name='powerTrainServiceAPI.v1.PowerTrackingRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
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
  serialized_start=428,
  serialized_end=450,
)


_POWERESTIMATERESPONSE = _descriptor.Descriptor(
  name='PowerEstimateResponse',
  full_name='powerTrainServiceAPI.v1.PowerEstimateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='unix_time', full_name='powerTrainServiceAPI.v1.PowerEstimateResponse.unix_time', index=0,
      number=1, type=1, cpp_type=5, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='power_estimate', full_name='powerTrainServiceAPI.v1.PowerEstimateResponse.power_estimate', index=1,
      number=2, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='power_estimate_average', full_name='powerTrainServiceAPI.v1.PowerEstimateResponse.power_estimate_average', index=2,
      number=3, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
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
  serialized_start=452,
  serialized_end=550,
)


_COSTESTIMATERESPONSE = _descriptor.Descriptor(
  name='CostEstimateResponse',
  full_name='powerTrainServiceAPI.v1.CostEstimateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='unix_time', full_name='powerTrainServiceAPI.v1.CostEstimateResponse.unix_time', index=0,
      number=1, type=1, cpp_type=5, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='power_estimate', full_name='powerTrainServiceAPI.v1.CostEstimateResponse.power_estimate', index=1,
      number=2, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cost_estimate', full_name='powerTrainServiceAPI.v1.CostEstimateResponse.cost_estimate', index=2,
      number=3, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='total_cost', full_name='powerTrainServiceAPI.v1.CostEstimateResponse.total_cost', index=3,
      number=4, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='power_estimate_average', full_name='powerTrainServiceAPI.v1.CostEstimateResponse.power_estimate_average', index=4,
      number=5, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
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
  serialized_start=553,
  serialized_end=693,
)


_POWEREVALUATIONRESPONSE = _descriptor.Descriptor(
  name='PowerEvaluationResponse',
  full_name='powerTrainServiceAPI.v1.PowerEvaluationResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='unix_time', full_name='powerTrainServiceAPI.v1.PowerEvaluationResponse.unix_time', index=0,
      number=1, type=1, cpp_type=5, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='power_estimate', full_name='powerTrainServiceAPI.v1.PowerEvaluationResponse.power_estimate', index=1,
      number=2, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='power_actual', full_name='powerTrainServiceAPI.v1.PowerEvaluationResponse.power_actual', index=2,
      number=3, type=2, cpp_type=6, label=3,
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
  serialized_start=695,
  serialized_end=785,
)


_POWERTRACKINGRESPONSE = _descriptor.Descriptor(
  name='PowerTrackingResponse',
  full_name='powerTrainServiceAPI.v1.PowerTrackingResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='power_actual', full_name='powerTrainServiceAPI.v1.PowerTrackingResponse.power_actual', index=0,
      number=1, type=2, cpp_type=6, label=3,
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
  serialized_start=787,
  serialized_end=832,
)

_POWERTRAINESTIMATEREQUEST.fields_by_name['model_type'].enum_type = _MODELTYPEENUM
DESCRIPTOR.message_types_by_name['PowerTrainEstimateRequest'] = _POWERTRAINESTIMATEREQUEST
DESCRIPTOR.message_types_by_name['PowerTrackingRequest'] = _POWERTRACKINGREQUEST
DESCRIPTOR.message_types_by_name['PowerEstimateResponse'] = _POWERESTIMATERESPONSE
DESCRIPTOR.message_types_by_name['CostEstimateResponse'] = _COSTESTIMATERESPONSE
DESCRIPTOR.message_types_by_name['PowerEvaluationResponse'] = _POWEREVALUATIONRESPONSE
DESCRIPTOR.message_types_by_name['PowerTrackingResponse'] = _POWERTRACKINGRESPONSE
DESCRIPTOR.enum_types_by_name['ModelTypeEnum'] = _MODELTYPEENUM
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

PowerTrainEstimateRequest = _reflection.GeneratedProtocolMessageType('PowerTrainEstimateRequest', (_message.Message,), {
  'DESCRIPTOR' : _POWERTRAINESTIMATEREQUEST,
  '__module__' : 'power_train_service_api_v1_pb2'
  # @@protoc_insertion_point(class_scope:powerTrainServiceAPI.v1.PowerTrainEstimateRequest)
  })
_sym_db.RegisterMessage(PowerTrainEstimateRequest)

PowerTrackingRequest = _reflection.GeneratedProtocolMessageType('PowerTrackingRequest', (_message.Message,), {
  'DESCRIPTOR' : _POWERTRACKINGREQUEST,
  '__module__' : 'power_train_service_api_v1_pb2'
  # @@protoc_insertion_point(class_scope:powerTrainServiceAPI.v1.PowerTrackingRequest)
  })
_sym_db.RegisterMessage(PowerTrackingRequest)

PowerEstimateResponse = _reflection.GeneratedProtocolMessageType('PowerEstimateResponse', (_message.Message,), {
  'DESCRIPTOR' : _POWERESTIMATERESPONSE,
  '__module__' : 'power_train_service_api_v1_pb2'
  # @@protoc_insertion_point(class_scope:powerTrainServiceAPI.v1.PowerEstimateResponse)
  })
_sym_db.RegisterMessage(PowerEstimateResponse)

CostEstimateResponse = _reflection.GeneratedProtocolMessageType('CostEstimateResponse', (_message.Message,), {
  'DESCRIPTOR' : _COSTESTIMATERESPONSE,
  '__module__' : 'power_train_service_api_v1_pb2'
  # @@protoc_insertion_point(class_scope:powerTrainServiceAPI.v1.CostEstimateResponse)
  })
_sym_db.RegisterMessage(CostEstimateResponse)

PowerEvaluationResponse = _reflection.GeneratedProtocolMessageType('PowerEvaluationResponse', (_message.Message,), {
  'DESCRIPTOR' : _POWEREVALUATIONRESPONSE,
  '__module__' : 'power_train_service_api_v1_pb2'
  # @@protoc_insertion_point(class_scope:powerTrainServiceAPI.v1.PowerEvaluationResponse)
  })
_sym_db.RegisterMessage(PowerEvaluationResponse)

PowerTrackingResponse = _reflection.GeneratedProtocolMessageType('PowerTrackingResponse', (_message.Message,), {
  'DESCRIPTOR' : _POWERTRACKINGRESPONSE,
  '__module__' : 'power_train_service_api_v1_pb2'
  # @@protoc_insertion_point(class_scope:powerTrainServiceAPI.v1.PowerTrackingResponse)
  })
_sym_db.RegisterMessage(PowerTrackingResponse)


DESCRIPTOR._options = None

_POWERTRAINSERVICE = _descriptor.ServiceDescriptor(
  name='PowerTrainService',
  full_name='powerTrainServiceAPI.v1.PowerTrainService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=889,
  serialized_end=1381,
  methods=[
  _descriptor.MethodDescriptor(
    name='PowerEstimate',
    full_name='powerTrainServiceAPI.v1.PowerTrainService.PowerEstimate',
    index=0,
    containing_service=None,
    input_type=_POWERTRAINESTIMATEREQUEST,
    output_type=_POWERESTIMATERESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CostEstimate',
    full_name='powerTrainServiceAPI.v1.PowerTrainService.CostEstimate',
    index=1,
    containing_service=None,
    input_type=_POWERTRAINESTIMATEREQUEST,
    output_type=_COSTESTIMATERESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='PowerTracking',
    full_name='powerTrainServiceAPI.v1.PowerTrainService.PowerTracking',
    index=2,
    containing_service=None,
    input_type=_POWERTRACKINGREQUEST,
    output_type=_POWERTRACKINGRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='PowerEstimateEvaluation',
    full_name='powerTrainServiceAPI.v1.PowerTrainService.PowerEstimateEvaluation',
    index=3,
    containing_service=None,
    input_type=_POWERTRAINESTIMATEREQUEST,
    output_type=_POWEREVALUATIONRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_POWERTRAINSERVICE)

DESCRIPTOR.services_by_name['PowerTrainService'] = _POWERTRAINSERVICE

# @@protoc_insertion_point(module_scope)
