# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/messages.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import field_mask_pb2 as google_dot_protobuf_dot_field__mask__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14proto/messages.proto\x12\x10update.compactor\x1a google/protobuf/field_mask.proto\"^\n\x06Person\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\x05\x12*\n\x07\x61\x64\x64ress\x18\x03 \x01(\x0b\x32\x19.update.compactor.Address\x12\x0e\n\x06\x65mails\x18\x04 \x03(\t\"b\n\x07\x41\x64\x64ress\x12\x0f\n\x07street1\x18\x01 \x01(\t\x12\x0f\n\x07street2\x18\x02 \x01(\t\x12\x0c\n\x04\x63ity\x18\x03 \x01(\t\x12\x13\n\x0bpostal_code\x18\x04 \x01(\t\x12\x12\n\nstate_code\x18\x05 \x01(\t\"u\n\x0cUpdatePerson\x12\n\n\x02id\x18\x01 \x01(\x05\x12(\n\x06person\x18\x02 \x01(\x0b\x32\x18.update.compactor.Person\x12/\n\x0bupdate_mask\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.FieldMaskB.Z,github.com/jamessocol/update-compactor/protob\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.messages_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z,github.com/jamessocol/update-compactor/proto'
  _PERSON._serialized_start=76
  _PERSON._serialized_end=170
  _ADDRESS._serialized_start=172
  _ADDRESS._serialized_end=270
  _UPDATEPERSON._serialized_start=272
  _UPDATEPERSON._serialized_end=389
# @@protoc_insertion_point(module_scope)
