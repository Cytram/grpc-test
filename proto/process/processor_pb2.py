# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/process/processor.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/process/processor.proto',
  package='processor',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x1dproto/process/processor.proto\x12\tprocessor\"6\n\x13ProcessImageRequest\x12\x1f\n\x05image\x18\x01 \x01(\x0b\x32\x10.processor.Image\"$\n\x11ProcessImageReply\x12\x0f\n\x07\x63ontent\x18\x01 \x01(\x0c\"@\n\x05Image\x12\x0f\n\x07\x63ontent\x18\x01 \x01(\x0c\x12&\n\x06source\x18\x02 \x01(\x0b\x32\x16.processor.ImageSource\"\x1f\n\x0bImageSource\x12\x10\n\x08http_uri\x18\x01 \x01(\t2[\n\tProcessor\x12N\n\x0cProcessImage\x12\x1e.processor.ProcessImageRequest\x1a\x1c.processor.ProcessImageReply\"\x00\x62\x06proto3')
)




_PROCESSIMAGEREQUEST = _descriptor.Descriptor(
  name='ProcessImageRequest',
  full_name='processor.ProcessImageRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='image', full_name='processor.ProcessImageRequest.image', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=44,
  serialized_end=98,
)


_PROCESSIMAGEREPLY = _descriptor.Descriptor(
  name='ProcessImageReply',
  full_name='processor.ProcessImageReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='content', full_name='processor.ProcessImageReply.content', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=100,
  serialized_end=136,
)


_IMAGE = _descriptor.Descriptor(
  name='Image',
  full_name='processor.Image',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='content', full_name='processor.Image.content', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source', full_name='processor.Image.source', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=138,
  serialized_end=202,
)


_IMAGESOURCE = _descriptor.Descriptor(
  name='ImageSource',
  full_name='processor.ImageSource',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='http_uri', full_name='processor.ImageSource.http_uri', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=204,
  serialized_end=235,
)

_PROCESSIMAGEREQUEST.fields_by_name['image'].message_type = _IMAGE
_IMAGE.fields_by_name['source'].message_type = _IMAGESOURCE
DESCRIPTOR.message_types_by_name['ProcessImageRequest'] = _PROCESSIMAGEREQUEST
DESCRIPTOR.message_types_by_name['ProcessImageReply'] = _PROCESSIMAGEREPLY
DESCRIPTOR.message_types_by_name['Image'] = _IMAGE
DESCRIPTOR.message_types_by_name['ImageSource'] = _IMAGESOURCE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ProcessImageRequest = _reflection.GeneratedProtocolMessageType('ProcessImageRequest', (_message.Message,), {
  'DESCRIPTOR' : _PROCESSIMAGEREQUEST,
  '__module__' : 'proto.process.processor_pb2'
  # @@protoc_insertion_point(class_scope:processor.ProcessImageRequest)
  })
_sym_db.RegisterMessage(ProcessImageRequest)

ProcessImageReply = _reflection.GeneratedProtocolMessageType('ProcessImageReply', (_message.Message,), {
  'DESCRIPTOR' : _PROCESSIMAGEREPLY,
  '__module__' : 'proto.process.processor_pb2'
  # @@protoc_insertion_point(class_scope:processor.ProcessImageReply)
  })
_sym_db.RegisterMessage(ProcessImageReply)

Image = _reflection.GeneratedProtocolMessageType('Image', (_message.Message,), {
  'DESCRIPTOR' : _IMAGE,
  '__module__' : 'proto.process.processor_pb2'
  # @@protoc_insertion_point(class_scope:processor.Image)
  })
_sym_db.RegisterMessage(Image)

ImageSource = _reflection.GeneratedProtocolMessageType('ImageSource', (_message.Message,), {
  'DESCRIPTOR' : _IMAGESOURCE,
  '__module__' : 'proto.process.processor_pb2'
  # @@protoc_insertion_point(class_scope:processor.ImageSource)
  })
_sym_db.RegisterMessage(ImageSource)



_PROCESSOR = _descriptor.ServiceDescriptor(
  name='Processor',
  full_name='processor.Processor',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=237,
  serialized_end=328,
  methods=[
  _descriptor.MethodDescriptor(
    name='ProcessImage',
    full_name='processor.Processor.ProcessImage',
    index=0,
    containing_service=None,
    input_type=_PROCESSIMAGEREQUEST,
    output_type=_PROCESSIMAGEREPLY,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_PROCESSOR)

DESCRIPTOR.services_by_name['Processor'] = _PROCESSOR

# @@protoc_insertion_point(module_scope)
