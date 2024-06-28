// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: weather.proto

package weather

import (
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

type Time struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Time) Reset() {
	*x = Time{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Time) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Time) ProtoMessage() {}

func (x *Time) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Time.ProtoReflect.Descriptor instead.
func (*Time) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{0}
}

func (x *Time) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type Weather struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weather    string `protobuf:"bytes,1,opt,name=weather,proto3" json:"weather,omitempty"`
	Tempratura int32  `protobuf:"varint,2,opt,name=tempratura,proto3" json:"tempratura,omitempty"`
	Damp       int32  `protobuf:"varint,3,opt,name=damp,proto3" json:"damp,omitempty"`
	Wind       int32  `protobuf:"varint,4,opt,name=wind,proto3" json:"wind,omitempty"`
	Time       string `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Weather) Reset() {
	*x = Weather{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Weather) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weather) ProtoMessage() {}

func (x *Weather) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weather.ProtoReflect.Descriptor instead.
func (*Weather) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{1}
}

func (x *Weather) GetWeather() string {
	if x != nil {
		return x.Weather
	}
	return ""
}

func (x *Weather) GetTempratura() int32 {
	if x != nil {
		return x.Tempratura
	}
	return 0
}

func (x *Weather) GetDamp() int32 {
	if x != nil {
		return x.Damp
	}
	return 0
}

func (x *Weather) GetWind() int32 {
	if x != nil {
		return x.Wind
	}
	return 0
}

func (x *Weather) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type Day struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Day string `protobuf:"bytes,1,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *Day) Reset() {
	*x = Day{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Day) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Day) ProtoMessage() {}

func (x *Day) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Day.ProtoReflect.Descriptor instead.
func (*Day) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{2}
}

func (x *Day) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_weather_proto protoreflect.FileDescriptor

var file_weather_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x22, 0x1a, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0x7f, 0x0a, 0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6d,
	0x70, 0x72, 0x61, 0x74, 0x75, 0x72, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74,
	0x65, 0x6d, 0x70, 0x72, 0x61, 0x74, 0x75, 0x72, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x77, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x77, 0x69, 0x6e,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x22, 0x20,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x32, 0xb9, 0x01, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x10, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12,
	0x0c, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x64, 0x61, 0x79, 0x1a, 0x10, 0x2e,
	0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12,
	0x3b, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x77, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x12, 0x5a, 0x10,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weather_proto_rawDescOnce sync.Once
	file_weather_proto_rawDescData = file_weather_proto_rawDesc
)

func file_weather_proto_rawDescGZIP() []byte {
	file_weather_proto_rawDescOnce.Do(func() {
		file_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_proto_rawDescData)
	})
	return file_weather_proto_rawDescData
}

var file_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_weather_proto_goTypes = []interface{}{
	(*Time)(nil),    // 0: weather.time
	(*Weather)(nil), // 1: weather.weather
	(*Day)(nil),     // 2: weather.day
	(*Status)(nil),  // 3: weather.status
}
var file_weather_proto_depIdxs = []int32{
	0, // 0: weather.WeatherService.GetCurrentWeather:input_type -> weather.time
	2, // 1: weather.WeatherService.GetWeatherForecast:input_type -> weather.day
	1, // 2: weather.WeatherService.ReportWeatherCondition:input_type -> weather.weather
	1, // 3: weather.WeatherService.GetCurrentWeather:output_type -> weather.weather
	1, // 4: weather.WeatherService.GetWeatherForecast:output_type -> weather.weather
	3, // 5: weather.WeatherService.ReportWeatherCondition:output_type -> weather.status
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_weather_proto_init() }
func file_weather_proto_init() {
	if File_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weather_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Time); i {
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
		file_weather_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Weather); i {
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
		file_weather_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Day); i {
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
		file_weather_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weather_proto_goTypes,
		DependencyIndexes: file_weather_proto_depIdxs,
		MessageInfos:      file_weather_proto_msgTypes,
	}.Build()
	File_weather_proto = out.File
	file_weather_proto_rawDesc = nil
	file_weather_proto_goTypes = nil
	file_weather_proto_depIdxs = nil
}
