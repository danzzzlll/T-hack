// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: tinkoff/cloud/tts/v1/tts.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AudioEncoding int32

const (
	AudioEncoding_ENCODING_UNSPECIFIED AudioEncoding = 0  // <i>Invalid value.</i>
	AudioEncoding_LINEAR16             AudioEncoding = 1  // Raw PCM with signed integer 16-bit linear audio samples.
	AudioEncoding_ALAW                 AudioEncoding = 8  // Raw PCM with A-law mapped 8-bit audio samples (aka PCMA).
	AudioEncoding_RAW_OPUS             AudioEncoding = 11 // Opus frames packed into Protobuf messages.<br/><b>NOTE</b>: each Opus frame is packed into separate message with `audio_content` field. I. e., you can't just concatenate encoded Opus frames and push it as a single chunk into Opus decoder. Also although Opus is sample rate agnostic, estimated duration of synthesized audio is calculated in samples of specified sample rate.<br/><b>NOTE</b>: Not supported in the Synthesize method.
)

// Enum value maps for AudioEncoding.
var (
	AudioEncoding_name = map[int32]string{
		0:  "ENCODING_UNSPECIFIED",
		1:  "LINEAR16",
		8:  "ALAW",
		11: "RAW_OPUS",
	}
	AudioEncoding_value = map[string]int32{
		"ENCODING_UNSPECIFIED": 0,
		"LINEAR16":             1,
		"ALAW":                 8,
		"RAW_OPUS":             11,
	}
)

func (x AudioEncoding) Enum() *AudioEncoding {
	p := new(AudioEncoding)
	*p = x
	return p
}

func (x AudioEncoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AudioEncoding) Descriptor() protoreflect.EnumDescriptor {
	return file_tinkoff_cloud_tts_v1_tts_proto_enumTypes[0].Descriptor()
}

func (AudioEncoding) Type() protoreflect.EnumType {
	return &file_tinkoff_cloud_tts_v1_tts_proto_enumTypes[0]
}

func (x AudioEncoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AudioEncoding.Descriptor instead.
func (AudioEncoding) EnumDescriptor() ([]byte, []int) {
	return file_tinkoff_cloud_tts_v1_tts_proto_rawDescGZIP(), []int{0}
}

type Voice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // Voice name.
}

func (x *Voice) Reset() {
	*x = Voice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Voice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Voice) ProtoMessage() {}

func (x *Voice) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Voice.ProtoReflect.Descriptor instead.
func (*Voice) Descriptor() ([]byte, []int) {
	return file_tinkoff_cloud_tts_v1_tts_proto_rawDescGZIP(), []int{0}
}

func (x *Voice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListVoicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListVoicesRequest) Reset() {
	*x = ListVoicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVoicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVoicesRequest) ProtoMessage() {}

func (x *ListVoicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVoicesRequest.ProtoReflect.Descriptor instead.
func (*ListVoicesRequest) Descriptor() ([]byte, []int) {
	return file_tinkoff_cloud_tts_v1_tts_proto_rawDescGZIP(), []int{1}
}

type ListVoicesResponses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Voices []*Voice `protobuf:"bytes,1,rep,name=voices,proto3" json:"voices,omitempty"` // Array of voices.
}

func (x *ListVoicesResponses) Reset() {
	*x = ListVoicesResponses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVoicesResponses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVoicesResponses) ProtoMessage() {}

func (x *ListVoicesResponses) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVoicesResponses.ProtoReflect.Descriptor instead.
func (*ListVoicesResponses) Descriptor() ([]byte, []int) {
	return file_tinkoff_cloud_tts_v1_tts_proto_rawDescGZIP(), []int{2}
}

func (x *ListVoicesResponses) GetVoices() []*Voice {
	if x != nil {
		return x.Voices
	}
	return nil
}

type SynthesisInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"` // Raw text to synthesize.
	Ssml string `protobuf:"bytes,2,opt,name=ssml,proto3" json:"ssml,omitempty"` // The SSML document to synthesize.
}

func (x *SynthesisInput) Reset() {
	*x = SynthesisInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynthesisInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynthesisInput) ProtoMessage() {}

func (x *SynthesisInput) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynthesisInput.ProtoReflect.Descriptor instead.
func (*SynthesisInput) Descriptor() ([]byte, []int) {
	return file_tinkoff_cloud_tts_v1_tts_proto_rawDescGZIP(), []int{3}
}

func (x *SynthesisInput) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SynthesisInput) GetSsml() string {
	if x != nil {
		return x.Ssml
	}
	return ""
}

type VoiceSelectionParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // Exact voice name.
}

func (x *VoiceSelectionParams) Reset() {
	*x = VoiceSelectionParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoiceSelectionParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceSelectionParams) ProtoMessage() {}

func (x *VoiceSelectionParams) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceSelectionParams.ProtoReflect.Descriptor instead.
func (*VoiceSelectionParams) Descriptor() ([]byte, []int) {
	return file_tinkoff_cloud_tts_v1_tts_proto_rawDescGZIP(), []int{4}
}

func (x *VoiceSelectionParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AudioConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioEncoding   AudioEncoding `protobuf:"varint,1,opt,name=audio_encoding,json=audioEncoding,proto3,enum=tinkoff.cloud.tts.v1.AudioEncoding" json:"audio_encoding,omitempty"` // Audio encoding. Specifies both container and codec. Must be specified explicitly.
	SampleRateHertz int32         `protobuf:"varint,5,opt,name=sample_rate_hertz,json=sampleRateHertz,proto3" json:"sample_rate_hertz,omitempty"`                                 // Sample rate of generated audio in Hertz. Must be specified explicitly. Only sample rates in inclusive range from 1000 Hz to 48000 Hz are supported.
}

func (x *AudioConfig) Reset() {
	*x = AudioConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioConfig) ProtoMessage() {}

func (x *AudioConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioConfig.ProtoReflect.Descriptor instead.
func (*AudioConfig) Descriptor() ([]byte, []int) {
	return file_tinkoff_cloud_tts_v1_tts_proto_rawDescGZIP(), []int{5}
}

func (x *AudioConfig) GetAudioEncoding() AudioEncoding {
	if x != nil {
		return x.AudioEncoding
	}
	return AudioEncoding_ENCODING_UNSPECIFIED
}

func (x *AudioConfig) GetSampleRateHertz() int32 {
	if x != nil {
		return x.SampleRateHertz
	}
	return 0
}

type SynthesizeSpeechRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input       *SynthesisInput       `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`                                // Input to synthesize.
	Voice       *VoiceSelectionParams `protobuf:"bytes,2,opt,name=voice,proto3" json:"voice,omitempty"`                                // Voice selection parameters.
	AudioConfig *AudioConfig          `protobuf:"bytes,3,opt,name=audio_config,json=audioConfig,proto3" json:"audio_config,omitempty"` // Audio configuration.  Must be specified explicitly.
}

func (x *SynthesizeSpeechRequest) Reset() {
	*x = SynthesizeSpeechRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynthesizeSpeechRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynthesizeSpeechRequest) ProtoMessage() {}

func (x *SynthesizeSpeechRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynthesizeSpeechRequest.ProtoReflect.Descriptor instead.
func (*SynthesizeSpeechRequest) Descriptor() ([]byte, []int) {
	return file_tinkoff_cloud_tts_v1_tts_proto_rawDescGZIP(), []int{6}
}

func (x *SynthesizeSpeechRequest) GetInput() *SynthesisInput {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *SynthesizeSpeechRequest) GetVoice() *VoiceSelectionParams {
	if x != nil {
		return x.Voice
	}
	return nil
}

func (x *SynthesizeSpeechRequest) GetAudioConfig() *AudioConfig {
	if x != nil {
		return x.AudioConfig
	}
	return nil
}

type SynthesizeSpeechResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioContent []byte `protobuf:"bytes,1,opt,name=audio_content,json=audioContent,proto3" json:"audio_content,omitempty"` // Whole synthesized audio.
}

func (x *SynthesizeSpeechResponse) Reset() {
	*x = SynthesizeSpeechResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynthesizeSpeechResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynthesizeSpeechResponse) ProtoMessage() {}

func (x *SynthesizeSpeechResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynthesizeSpeechResponse.ProtoReflect.Descriptor instead.
func (*SynthesizeSpeechResponse) Descriptor() ([]byte, []int) {
	return file_tinkoff_cloud_tts_v1_tts_proto_rawDescGZIP(), []int{7}
}

func (x *SynthesizeSpeechResponse) GetAudioContent() []byte {
	if x != nil {
		return x.AudioContent
	}
	return nil
}

type StreamingSynthesizeSpeechResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioChunk []byte `protobuf:"bytes,1,opt,name=audio_chunk,json=audioChunk,proto3" json:"audio_chunk,omitempty"` // Chunk of synthesized audio: either samples for `LINEAR16` or single frame for `RAW_OPUS`.
}

func (x *StreamingSynthesizeSpeechResponse) Reset() {
	*x = StreamingSynthesizeSpeechResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingSynthesizeSpeechResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingSynthesizeSpeechResponse) ProtoMessage() {}

func (x *StreamingSynthesizeSpeechResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingSynthesizeSpeechResponse.ProtoReflect.Descriptor instead.
func (*StreamingSynthesizeSpeechResponse) Descriptor() ([]byte, []int) {
	return file_tinkoff_cloud_tts_v1_tts_proto_rawDescGZIP(), []int{8}
}

func (x *StreamingSynthesizeSpeechResponse) GetAudioChunk() []byte {
	if x != nil {
		return x.AudioChunk
	}
	return nil
}

var File_tinkoff_cloud_tts_v1_tts_proto protoreflect.FileDescriptor

var file_tinkoff_cloud_tts_v1_tts_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x74, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x74, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x05, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08,
	0x04, 0x10, 0x05, 0x52, 0x0e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x73, 0x52, 0x0b, 0x73, 0x73, 0x6d, 0x6c, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x52, 0x19, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x22, 0x28, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x4a, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x06,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74,
	0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x06, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x73, 0x22, 0x38, 0x0a, 0x0e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x73, 0x6d, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x73, 0x6d, 0x6c, 0x22, 0x52, 0x0a, 0x14, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08,
	0x03, 0x10, 0x04, 0x52, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x52, 0x0b, 0x73, 0x73, 0x6d, 0x6c, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22,
	0xbd, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x4a, 0x0a, 0x0e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66,
	0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x61, 0x75,
	0x64, 0x69, 0x6f, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x0a, 0x11, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x65, 0x72, 0x74, 0x7a,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x48, 0x65, 0x72, 0x74, 0x7a, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08,
	0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x52, 0x0d, 0x73, 0x70, 0x65, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x52, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x52,
	0x0e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x67, 0x61, 0x69, 0x6e, 0x5f, 0x64, 0x62, 0x22,
	0xdd, 0x01, 0x0a, 0x17, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x53, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x69, 0x6e,
	0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x40, 0x0a, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x52, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x61, 0x75, 0x64,
	0x69, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x74, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x3f, 0x0a, 0x18, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x53, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x44, 0x0a, 0x21, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x6e,
	0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x75, 0x64, 0x69,
	0x6f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x2a, 0xe4, 0x01, 0x0a, 0x0d, 0x41, 0x75, 0x64, 0x69, 0x6f,
	0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x4e, 0x43, 0x4f,
	0x44, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x49, 0x4e, 0x45, 0x41, 0x52, 0x31, 0x36, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x41, 0x4c, 0x41, 0x57, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x41,
	0x57, 0x5f, 0x4f, 0x50, 0x55, 0x53, 0x10, 0x0b, 0x22, 0x04, 0x08, 0x02, 0x10, 0x02, 0x22, 0x04,
	0x08, 0x03, 0x10, 0x03, 0x22, 0x04, 0x08, 0x04, 0x10, 0x04, 0x22, 0x04, 0x08, 0x05, 0x10, 0x05,
	0x22, 0x04, 0x08, 0x06, 0x10, 0x06, 0x22, 0x04, 0x08, 0x07, 0x10, 0x07, 0x22, 0x04, 0x08, 0x09,
	0x10, 0x09, 0x22, 0x04, 0x08, 0x0a, 0x10, 0x0a, 0x22, 0x04, 0x08, 0x0c, 0x10, 0x0c, 0x2a, 0x04,
	0x46, 0x4c, 0x41, 0x43, 0x2a, 0x05, 0x4d, 0x55, 0x4c, 0x41, 0x57, 0x2a, 0x03, 0x41, 0x4d, 0x52,
	0x2a, 0x06, 0x41, 0x4d, 0x52, 0x5f, 0x57, 0x42, 0x2a, 0x08, 0x4f, 0x47, 0x47, 0x5f, 0x4f, 0x50,
	0x55, 0x53, 0x2a, 0x16, 0x53, 0x50, 0x45, 0x45, 0x58, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x48,
	0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x2a, 0x09, 0x4c, 0x49, 0x4e, 0x45,
	0x41, 0x52, 0x33, 0x32, 0x46, 0x2a, 0x0a, 0x4f, 0x47, 0x47, 0x5f, 0x56, 0x4f, 0x52, 0x42, 0x49,
	0x53, 0x2a, 0x0a, 0x4d, 0x50, 0x45, 0x47, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x32, 0x9b, 0x03,
	0x0a, 0x0c, 0x54, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x12, 0x7d,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x27, 0x2e, 0x74,
	0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x74,
	0x73, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x8a, 0x01,
	0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x2d, 0x2e, 0x74,
	0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x53, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x74, 0x69,
	0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x53, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x74, 0x73, 0x3a, 0x73, 0x79, 0x6e,
	0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x7f, 0x0a, 0x13, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x2d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73,
	0x69, 0x7a, 0x65, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x37, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x53, 0x70, 0x65, 0x65, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x4e, 0x5a, 0x44, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x69, 0x6e, 0x6b, 0x6f, 0x66,
	0x66, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x6b, 0x69, 0x74, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74,
	0x69, 0x6e, 0x6b, 0x6f, 0x66, 0x66, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x54, 0x56, 0x4b, 0x53, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tinkoff_cloud_tts_v1_tts_proto_rawDescOnce sync.Once
	file_tinkoff_cloud_tts_v1_tts_proto_rawDescData = file_tinkoff_cloud_tts_v1_tts_proto_rawDesc
)

func file_tinkoff_cloud_tts_v1_tts_proto_rawDescGZIP() []byte {
	file_tinkoff_cloud_tts_v1_tts_proto_rawDescOnce.Do(func() {
		file_tinkoff_cloud_tts_v1_tts_proto_rawDescData = protoimpl.X.CompressGZIP(file_tinkoff_cloud_tts_v1_tts_proto_rawDescData)
	})
	return file_tinkoff_cloud_tts_v1_tts_proto_rawDescData
}

var file_tinkoff_cloud_tts_v1_tts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tinkoff_cloud_tts_v1_tts_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tinkoff_cloud_tts_v1_tts_proto_goTypes = []interface{}{
	(AudioEncoding)(0),                        // 0: tinkoff.cloud.tts.v1.AudioEncoding
	(*Voice)(nil),                             // 1: tinkoff.cloud.tts.v1.Voice
	(*ListVoicesRequest)(nil),                 // 2: tinkoff.cloud.tts.v1.ListVoicesRequest
	(*ListVoicesResponses)(nil),               // 3: tinkoff.cloud.tts.v1.ListVoicesResponses
	(*SynthesisInput)(nil),                    // 4: tinkoff.cloud.tts.v1.SynthesisInput
	(*VoiceSelectionParams)(nil),              // 5: tinkoff.cloud.tts.v1.VoiceSelectionParams
	(*AudioConfig)(nil),                       // 6: tinkoff.cloud.tts.v1.AudioConfig
	(*SynthesizeSpeechRequest)(nil),           // 7: tinkoff.cloud.tts.v1.SynthesizeSpeechRequest
	(*SynthesizeSpeechResponse)(nil),          // 8: tinkoff.cloud.tts.v1.SynthesizeSpeechResponse
	(*StreamingSynthesizeSpeechResponse)(nil), // 9: tinkoff.cloud.tts.v1.StreamingSynthesizeSpeechResponse
}
var file_tinkoff_cloud_tts_v1_tts_proto_depIdxs = []int32{
	1, // 0: tinkoff.cloud.tts.v1.ListVoicesResponses.voices:type_name -> tinkoff.cloud.tts.v1.Voice
	0, // 1: tinkoff.cloud.tts.v1.AudioConfig.audio_encoding:type_name -> tinkoff.cloud.tts.v1.AudioEncoding
	4, // 2: tinkoff.cloud.tts.v1.SynthesizeSpeechRequest.input:type_name -> tinkoff.cloud.tts.v1.SynthesisInput
	5, // 3: tinkoff.cloud.tts.v1.SynthesizeSpeechRequest.voice:type_name -> tinkoff.cloud.tts.v1.VoiceSelectionParams
	6, // 4: tinkoff.cloud.tts.v1.SynthesizeSpeechRequest.audio_config:type_name -> tinkoff.cloud.tts.v1.AudioConfig
	2, // 5: tinkoff.cloud.tts.v1.TextToSpeech.ListVoices:input_type -> tinkoff.cloud.tts.v1.ListVoicesRequest
	7, // 6: tinkoff.cloud.tts.v1.TextToSpeech.Synthesize:input_type -> tinkoff.cloud.tts.v1.SynthesizeSpeechRequest
	7, // 7: tinkoff.cloud.tts.v1.TextToSpeech.StreamingSynthesize:input_type -> tinkoff.cloud.tts.v1.SynthesizeSpeechRequest
	3, // 8: tinkoff.cloud.tts.v1.TextToSpeech.ListVoices:output_type -> tinkoff.cloud.tts.v1.ListVoicesResponses
	8, // 9: tinkoff.cloud.tts.v1.TextToSpeech.Synthesize:output_type -> tinkoff.cloud.tts.v1.SynthesizeSpeechResponse
	9, // 10: tinkoff.cloud.tts.v1.TextToSpeech.StreamingSynthesize:output_type -> tinkoff.cloud.tts.v1.StreamingSynthesizeSpeechResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tinkoff_cloud_tts_v1_tts_proto_init() }
func file_tinkoff_cloud_tts_v1_tts_proto_init() {
	if File_tinkoff_cloud_tts_v1_tts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Voice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVoicesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVoicesResponses); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynthesisInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoiceSelectionParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynthesizeSpeechRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynthesizeSpeechResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tinkoff_cloud_tts_v1_tts_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingSynthesizeSpeechResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tinkoff_cloud_tts_v1_tts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tinkoff_cloud_tts_v1_tts_proto_goTypes,
		DependencyIndexes: file_tinkoff_cloud_tts_v1_tts_proto_depIdxs,
		EnumInfos:         file_tinkoff_cloud_tts_v1_tts_proto_enumTypes,
		MessageInfos:      file_tinkoff_cloud_tts_v1_tts_proto_msgTypes,
	}.Build()
	File_tinkoff_cloud_tts_v1_tts_proto = out.File
	file_tinkoff_cloud_tts_v1_tts_proto_rawDesc = nil
	file_tinkoff_cloud_tts_v1_tts_proto_goTypes = nil
	file_tinkoff_cloud_tts_v1_tts_proto_depIdxs = nil
}