TF_DIR=../../tensorflow/tensorflow
TF_FILES=$(shell find ${TF_DIR}/tensorflow/core/framework -name '*.proto')

proto: $(TF_FILES)
	mkdir -p proto
	protoc \
  -I ${TF_DIR} \
  --gogoslick_out=proto \
  $(TF_FILES)
	touch proto
