protoc --proto_path=./protos \
  --go_out=$import_replacements:./protos-cafes \
  $file_proto
