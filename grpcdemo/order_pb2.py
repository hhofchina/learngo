# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: order.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import service as _service
from google.protobuf import service_reflection
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='order.proto',
  package='order',
  syntax='proto3',
  serialized_pb=_b('\n\x0border.proto\x12\x05order\"\x9d\x01\n\x08OrderReq\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0f\n\x07user_id\x18\x66 \x01(\x05\x12\x0e\n\x06pro_id\x18g \x01(\x05\x12\x0e\n\x06wds_id\x18h \x01(\x05\x12\x0e\n\x06rew_id\x18i \x01(\x05\x12\x0f\n\x07md_play\x18m \x01(\t\x12\x12\n\naddress_id\x18n \x01(\x05\x12\x1f\n\x07\x61\x64\x64ress\x18o \x01(\x0b\x32\x0e.order.Address\"\xb6\x01\n\x08OrderRsp\x12+\n\x06orders\x18\x01 \x03(\x0b\x32\x1b.order.OrderRsp.OrderResult\x1a}\n\x0bOrderResult\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0f\n\x07user_id\x18\x02 \x01(\x05\x12\x0e\n\x06pro_id\x18\x03 \x01(\x05\x12\x0e\n\x06status\x18\x04 \x01(\x05\x12\x0e\n\x06\x61mount\x18\x05 \x01(\x02\x12\x0e\n\x06if_pay\x18\x06 \x01(\x05\x12\x11\n\tgame_code\x18\x14 \x01(\t\"\x87\x01\n\x07\x41\x64\x64ress\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x10\n\x08province\x18\x03 \x01(\t\x12\x0c\n\x04\x63ity\x18\x04 \x01(\t\x12\x0e\n\x06\x64\x65tail\x18\x05 \x01(\t\x12\x0e\n\x06mobile\x18\x06 \x01(\t\x12\r\n\x05\x65mail\x18\x07 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x08 \x01(\t\"8\n\x08StockReq\x12\x0e\n\x06pro_id\x18\x01 \x01(\x05\x12\x0c\n\x04lock\x18\x04 \x01(\x05\x12\x0e\n\x06unlock\x18\x05 \x01(\x05\"V\n\x08StockRsp\x12\x0e\n\x06pro_id\x18\x01 \x01(\x05\x12\r\n\x05stock\x18\x02 \x01(\x05\x12\r\n\x05total\x18\x65 \x01(\x05\x12\x0c\n\x04lock\x18\x66 \x01(\x05\x12\x0e\n\x06unlock\x18g \x01(\x05\"W\n\x08QueueReq\x12\x0f\n\x07user_id\x18\x01 \x01(\x05\x12\x0e\n\x06pro_id\x18\x02 \x01(\x05\x12\x0b\n\x03qid\x18\x65 \x01(\x05\x12\x0e\n\x06weight\x18\x66 \x01(\x02\x12\r\n\x05total\x18g \x01(\x05\"H\n\x08QueueRsp\x12\x0b\n\x03qid\x18\x01 \x01(\x05\x12\x0f\n\x07user_id\x18\x02 \x01(\x05\x12\x0e\n\x06pro_id\x18\x03 \x01(\x05\x12\x0e\n\x06weight\x18\x04 \x01(\x02\x32\xc5\x03\n\x0cOrderService\x12,\n\x08GetOrder\x12\x0f.order.OrderReq\x1a\x0f.order.OrderRsp\x12\x31\n\tListOrder\x12\x0f.order.OrderReq\x1a\x0f.order.OrderRsp(\x01\x30\x01\x12,\n\x08NewOrder\x12\x0f.order.OrderReq\x1a\x0f.order.OrderRsp\x12/\n\x0b\x43\x61ncelOrder\x12\x0f.order.OrderReq\x1a\x0f.order.OrderRsp\x12\x30\n\nCheckStock\x12\x0f.order.StockReq\x1a\x0f.order.StockRsp\"\x00\x12/\n\tLockStock\x12\x0f.order.StockReq\x1a\x0f.order.StockRsp\"\x00\x12\x31\n\x0bUnlockStock\x12\x0f.order.StockReq\x1a\x0f.order.StockRsp\"\x00\x12/\n\tQueuePush\x12\x0f.order.QueueReq\x1a\x0f.order.QueueRsp\"\x00\x12.\n\x08QueuePop\x12\x0f.order.QueueReq\x1a\x0f.order.QueueRsp\"\x00\x42#\n\x0c\x63om.md.orderB\nOrderProtoP\x01Z\x02md\x90\x01\x01\x62\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_ORDERREQ = _descriptor.Descriptor(
  name='OrderReq',
  full_name='order.OrderReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='order.OrderReq.id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='order.OrderReq.user_id', index=1,
      number=102, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='pro_id', full_name='order.OrderReq.pro_id', index=2,
      number=103, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='wds_id', full_name='order.OrderReq.wds_id', index=3,
      number=104, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='rew_id', full_name='order.OrderReq.rew_id', index=4,
      number=105, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='md_play', full_name='order.OrderReq.md_play', index=5,
      number=109, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='address_id', full_name='order.OrderReq.address_id', index=6,
      number=110, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='address', full_name='order.OrderReq.address', index=7,
      number=111, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=23,
  serialized_end=180,
)


_ORDERRSP_ORDERRESULT = _descriptor.Descriptor(
  name='OrderResult',
  full_name='order.OrderRsp.OrderResult',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='order.OrderRsp.OrderResult.id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='order.OrderRsp.OrderResult.user_id', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='pro_id', full_name='order.OrderRsp.OrderResult.pro_id', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='status', full_name='order.OrderRsp.OrderResult.status', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='amount', full_name='order.OrderRsp.OrderResult.amount', index=4,
      number=5, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='if_pay', full_name='order.OrderRsp.OrderResult.if_pay', index=5,
      number=6, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='game_code', full_name='order.OrderRsp.OrderResult.game_code', index=6,
      number=20, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=240,
  serialized_end=365,
)

_ORDERRSP = _descriptor.Descriptor(
  name='OrderRsp',
  full_name='order.OrderRsp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='orders', full_name='order.OrderRsp.orders', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[_ORDERRSP_ORDERRESULT, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=183,
  serialized_end=365,
)


_ADDRESS = _descriptor.Descriptor(
  name='Address',
  full_name='order.Address',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='order.Address.id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='order.Address.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='province', full_name='order.Address.province', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='city', full_name='order.Address.city', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='detail', full_name='order.Address.detail', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='mobile', full_name='order.Address.mobile', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='email', full_name='order.Address.email', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='description', full_name='order.Address.description', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=368,
  serialized_end=503,
)


_STOCKREQ = _descriptor.Descriptor(
  name='StockReq',
  full_name='order.StockReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pro_id', full_name='order.StockReq.pro_id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='lock', full_name='order.StockReq.lock', index=1,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='unlock', full_name='order.StockReq.unlock', index=2,
      number=5, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=505,
  serialized_end=561,
)


_STOCKRSP = _descriptor.Descriptor(
  name='StockRsp',
  full_name='order.StockRsp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pro_id', full_name='order.StockRsp.pro_id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='stock', full_name='order.StockRsp.stock', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='total', full_name='order.StockRsp.total', index=2,
      number=101, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='lock', full_name='order.StockRsp.lock', index=3,
      number=102, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='unlock', full_name='order.StockRsp.unlock', index=4,
      number=103, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=563,
  serialized_end=649,
)


_QUEUEREQ = _descriptor.Descriptor(
  name='QueueReq',
  full_name='order.QueueReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='order.QueueReq.user_id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='pro_id', full_name='order.QueueReq.pro_id', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='qid', full_name='order.QueueReq.qid', index=2,
      number=101, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='weight', full_name='order.QueueReq.weight', index=3,
      number=102, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='total', full_name='order.QueueReq.total', index=4,
      number=103, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=651,
  serialized_end=738,
)


_QUEUERSP = _descriptor.Descriptor(
  name='QueueRsp',
  full_name='order.QueueRsp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='qid', full_name='order.QueueRsp.qid', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='order.QueueRsp.user_id', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='pro_id', full_name='order.QueueRsp.pro_id', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='weight', full_name='order.QueueRsp.weight', index=3,
      number=4, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=740,
  serialized_end=812,
)

_ORDERREQ.fields_by_name['address'].message_type = _ADDRESS
_ORDERRSP_ORDERRESULT.containing_type = _ORDERRSP
_ORDERRSP.fields_by_name['orders'].message_type = _ORDERRSP_ORDERRESULT
DESCRIPTOR.message_types_by_name['OrderReq'] = _ORDERREQ
DESCRIPTOR.message_types_by_name['OrderRsp'] = _ORDERRSP
DESCRIPTOR.message_types_by_name['Address'] = _ADDRESS
DESCRIPTOR.message_types_by_name['StockReq'] = _STOCKREQ
DESCRIPTOR.message_types_by_name['StockRsp'] = _STOCKRSP
DESCRIPTOR.message_types_by_name['QueueReq'] = _QUEUEREQ
DESCRIPTOR.message_types_by_name['QueueRsp'] = _QUEUERSP

OrderReq = _reflection.GeneratedProtocolMessageType('OrderReq', (_message.Message,), dict(
  DESCRIPTOR = _ORDERREQ,
  __module__ = 'order_pb2'
  # @@protoc_insertion_point(class_scope:order.OrderReq)
  ))
_sym_db.RegisterMessage(OrderReq)

OrderRsp = _reflection.GeneratedProtocolMessageType('OrderRsp', (_message.Message,), dict(

  OrderResult = _reflection.GeneratedProtocolMessageType('OrderResult', (_message.Message,), dict(
    DESCRIPTOR = _ORDERRSP_ORDERRESULT,
    __module__ = 'order_pb2'
    # @@protoc_insertion_point(class_scope:order.OrderRsp.OrderResult)
    ))
  ,
  DESCRIPTOR = _ORDERRSP,
  __module__ = 'order_pb2'
  # @@protoc_insertion_point(class_scope:order.OrderRsp)
  ))
_sym_db.RegisterMessage(OrderRsp)
_sym_db.RegisterMessage(OrderRsp.OrderResult)

Address = _reflection.GeneratedProtocolMessageType('Address', (_message.Message,), dict(
  DESCRIPTOR = _ADDRESS,
  __module__ = 'order_pb2'
  # @@protoc_insertion_point(class_scope:order.Address)
  ))
_sym_db.RegisterMessage(Address)

StockReq = _reflection.GeneratedProtocolMessageType('StockReq', (_message.Message,), dict(
  DESCRIPTOR = _STOCKREQ,
  __module__ = 'order_pb2'
  # @@protoc_insertion_point(class_scope:order.StockReq)
  ))
_sym_db.RegisterMessage(StockReq)

StockRsp = _reflection.GeneratedProtocolMessageType('StockRsp', (_message.Message,), dict(
  DESCRIPTOR = _STOCKRSP,
  __module__ = 'order_pb2'
  # @@protoc_insertion_point(class_scope:order.StockRsp)
  ))
_sym_db.RegisterMessage(StockRsp)

QueueReq = _reflection.GeneratedProtocolMessageType('QueueReq', (_message.Message,), dict(
  DESCRIPTOR = _QUEUEREQ,
  __module__ = 'order_pb2'
  # @@protoc_insertion_point(class_scope:order.QueueReq)
  ))
_sym_db.RegisterMessage(QueueReq)

QueueRsp = _reflection.GeneratedProtocolMessageType('QueueRsp', (_message.Message,), dict(
  DESCRIPTOR = _QUEUERSP,
  __module__ = 'order_pb2'
  # @@protoc_insertion_point(class_scope:order.QueueRsp)
  ))
_sym_db.RegisterMessage(QueueRsp)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n\014com.md.orderB\nOrderProtoP\001Z\002md\220\001\001'))

_ORDERSERVICE = _descriptor.ServiceDescriptor(
  name='OrderService',
  full_name='order.OrderService',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=815,
  serialized_end=1268,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetOrder',
    full_name='order.OrderService.GetOrder',
    index=0,
    containing_service=None,
    input_type=_ORDERREQ,
    output_type=_ORDERRSP,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ListOrder',
    full_name='order.OrderService.ListOrder',
    index=1,
    containing_service=None,
    input_type=_ORDERREQ,
    output_type=_ORDERRSP,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='NewOrder',
    full_name='order.OrderService.NewOrder',
    index=2,
    containing_service=None,
    input_type=_ORDERREQ,
    output_type=_ORDERRSP,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='CancelOrder',
    full_name='order.OrderService.CancelOrder',
    index=3,
    containing_service=None,
    input_type=_ORDERREQ,
    output_type=_ORDERRSP,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='CheckStock',
    full_name='order.OrderService.CheckStock',
    index=4,
    containing_service=None,
    input_type=_STOCKREQ,
    output_type=_STOCKRSP,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='LockStock',
    full_name='order.OrderService.LockStock',
    index=5,
    containing_service=None,
    input_type=_STOCKREQ,
    output_type=_STOCKRSP,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='UnlockStock',
    full_name='order.OrderService.UnlockStock',
    index=6,
    containing_service=None,
    input_type=_STOCKREQ,
    output_type=_STOCKRSP,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='QueuePush',
    full_name='order.OrderService.QueuePush',
    index=7,
    containing_service=None,
    input_type=_QUEUEREQ,
    output_type=_QUEUERSP,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='QueuePop',
    full_name='order.OrderService.QueuePop',
    index=8,
    containing_service=None,
    input_type=_QUEUEREQ,
    output_type=_QUEUERSP,
    options=None,
  ),
])

OrderService = service_reflection.GeneratedServiceType('OrderService', (_service.Service,), dict(
  DESCRIPTOR = _ORDERSERVICE,
  __module__ = 'order_pb2'
  ))

OrderService_Stub = service_reflection.GeneratedServiceStubType('OrderService_Stub', (OrderService,), dict(
  DESCRIPTOR = _ORDERSERVICE,
  __module__ = 'order_pb2'
  ))


# @@protoc_insertion_point(module_scope)
