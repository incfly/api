// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/operator.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Status describes the current state of a component.
type IstioOperatorSpec_Status int32

const (
	// Component is not present.
	IstioOperatorSpec_NONE IstioOperatorSpec_Status = 0
	// Component is being updated to a different version.
	IstioOperatorSpec_UPDATING IstioOperatorSpec_Status = 1
	// Controller has started but not yet completed reconciliation loop for the component.
	IstioOperatorSpec_RECONCILING IstioOperatorSpec_Status = 2
	// Component is healthy.
	IstioOperatorSpec_HEALTHY IstioOperatorSpec_Status = 3
	// Component is in an error state.
	IstioOperatorSpec_ERROR IstioOperatorSpec_Status = 4
)

var IstioOperatorSpec_Status_name = map[int32]string{
	0: "NONE",
	1: "UPDATING",
	2: "RECONCILING",
	3: "HEALTHY",
	4: "ERROR",
}

var IstioOperatorSpec_Status_value = map[string]int32{
	"NONE":        0,
	"UPDATING":    1,
	"RECONCILING": 2,
	"HEALTHY":     3,
	"ERROR":       4,
}

func (x IstioOperatorSpec_Status) String() string {
	return proto.EnumName(IstioOperatorSpec_Status_name, int32(x))
}

func (IstioOperatorSpec_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_03e04c906468a7ff, []int{0, 0}
}

// IstioOperatorSpec defines the desired installed state of Istio components.
// The spec is a used to define a customization of the default profile values that are supplied with each Istio release.
// Because the spec is a customization API, specifying an empty IstioOperatorSpec results in a default Istio
// component values.
type IstioOperatorSpec struct {
	// Path or name for the profile e.g.
	//     - minimal (looks in profiles dir for a file called minimal.yaml)
	//     - /tmp/istio/install/values/custom/custom-install.yaml (local file path)
	// default profile is used if this field is unset.
	Profile string `protobuf:"bytes,10,opt,name=profile,proto3" json:"profile,omitempty"`
	// Path for the install package. e.g.
	//     - /tmp/istio-installer/nightly (local file path)
	InstallPackagePath string `protobuf:"bytes,11,opt,name=install_package_path,json=installPackagePath,proto3" json:"install_package_path,omitempty"`
	// Root for docker image paths e.g. docker.io/istio
	Hub string `protobuf:"bytes,12,opt,name=hub,proto3" json:"hub,omitempty"`
	// Version tag for docker images e.g. 1.0.6
	Tag string `protobuf:"bytes,13,opt,name=tag,proto3" json:"tag,omitempty"`
	// Resource suffix is appended to all resources installed by each component. Used in upgrade scenarios where two
	// Istio control planes must exist in the same namespace.
	ResourceSuffix string `protobuf:"bytes,14,opt,name=resource_suffix,json=resourceSuffix,proto3" json:"resource_suffix,omitempty"`
	// Config used by control plane components internally.
	MeshConfig *MeshConfig `protobuf:"bytes,40,opt,name=mesh_config,json=meshConfig,proto3" json:"mesh_config,omitempty"`
	// Kubernetes resource settings, enablement and component-specific settings that are not internal to the
	// component.
	Components *IstioComponentSetSpec `protobuf:"bytes,50,opt,name=components,proto3" json:"components,omitempty"`
	// Overrides for default values.yaml. This is a validated pass-through to Helm templates.
	// See the Helm installation options for schema details: https://istio.io/docs/reference/config/installation-options/.
	// Anything that is available in IstioOperatorSpec should be set above rather than using the passthrough. This
	// includes Kubernetes resource settings for components in KubernetesResourcesSpec.
	Values map[string]interface{} `protobuf:"bytes,100,opt,name=values,proto3" json:"values,omitempty"`
	// Unvalidated overrides for default values.yaml. Used for custom templates where new parameters are added.
	UnvalidatedValues map[string]interface{} `protobuf:"bytes,101,opt,name=unvalidated_values,json=unvalidatedValues,proto3" json:"unvalidated_values,omitempty"`
	// Overall status of all components controlled by the operator.
	// - If all components have status NONE, overall status is NONE.
	// - If all components are HEALTHY, overall status is HEALTHY.
	// - If one or more components are RECONCILING and others are HEALTHY, overall status is RECONCILING.
	// - If one or more components are UPDATING and others are HEALTHY, overall status is UPDATING.
	// - If components are a mix of RECONCILING, UPDATING and HEALTHY, overall status is UPDATING.
	// - If any component is in ERROR state, overall status is ERROR.
	Status IstioOperatorSpec_Status `protobuf:"varint,200,opt,name=status,proto3,enum=istio.mesh.v1alpha1.IstioOperatorSpec_Status" json:"status,omitempty"`
	// Individual status of each component controlled by the operator. The map key is the name of the component.
	ComponentStatus      map[string]*IstioOperatorSpec_VersionStatus `protobuf:"bytes,201,rep,name=component_status,json=componentStatus,proto3" json:"component_status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *IstioOperatorSpec) Reset()         { *m = IstioOperatorSpec{} }
func (m *IstioOperatorSpec) String() string { return proto.CompactTextString(m) }
func (*IstioOperatorSpec) ProtoMessage()    {}
func (*IstioOperatorSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_03e04c906468a7ff, []int{0}
}

func (m *IstioOperatorSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IstioOperatorSpec.Unmarshal(m, b)
}
func (m *IstioOperatorSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IstioOperatorSpec.Marshal(b, m, deterministic)
}
func (m *IstioOperatorSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioOperatorSpec.Merge(m, src)
}
func (m *IstioOperatorSpec) XXX_Size() int {
	return xxx_messageInfo_IstioOperatorSpec.Size(m)
}
func (m *IstioOperatorSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioOperatorSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IstioOperatorSpec proto.InternalMessageInfo

func (m *IstioOperatorSpec) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *IstioOperatorSpec) GetInstallPackagePath() string {
	if m != nil {
		return m.InstallPackagePath
	}
	return ""
}

func (m *IstioOperatorSpec) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

func (m *IstioOperatorSpec) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *IstioOperatorSpec) GetResourceSuffix() string {
	if m != nil {
		return m.ResourceSuffix
	}
	return ""
}

func (m *IstioOperatorSpec) GetMeshConfig() *MeshConfig {
	if m != nil {
		return m.MeshConfig
	}
	return nil
}

func (m *IstioOperatorSpec) GetComponents() *IstioComponentSetSpec {
	if m != nil {
		return m.Components
	}
	return nil
}



func (m *IstioOperatorSpec) GetStatus() IstioOperatorSpec_Status {
	if m != nil {
		return m.Status
	}
	return IstioOperatorSpec_NONE
}

func (m *IstioOperatorSpec) GetComponentStatus() map[string]*IstioOperatorSpec_VersionStatus {
	if m != nil {
		return m.ComponentStatus
	}
	return nil
}

// VersionStatus is the status and version of a component.
type IstioOperatorSpec_VersionStatus struct {
	Version              string                   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Status               IstioOperatorSpec_Status `protobuf:"varint,2,opt,name=status,proto3,enum=istio.mesh.v1alpha1.IstioOperatorSpec_Status" json:"status,omitempty"`
	StatusString         string                   `protobuf:"bytes,3,opt,name=status_string,json=statusString,proto3" json:"status_string,omitempty"`
	Error                string                   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *IstioOperatorSpec_VersionStatus) Reset()         { *m = IstioOperatorSpec_VersionStatus{} }
func (m *IstioOperatorSpec_VersionStatus) String() string { return proto.CompactTextString(m) }
func (*IstioOperatorSpec_VersionStatus) ProtoMessage()    {}
func (*IstioOperatorSpec_VersionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_03e04c906468a7ff, []int{0, 0}
}

func (m *IstioOperatorSpec_VersionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IstioOperatorSpec_VersionStatus.Unmarshal(m, b)
}
func (m *IstioOperatorSpec_VersionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IstioOperatorSpec_VersionStatus.Marshal(b, m, deterministic)
}
func (m *IstioOperatorSpec_VersionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioOperatorSpec_VersionStatus.Merge(m, src)
}
func (m *IstioOperatorSpec_VersionStatus) XXX_Size() int {
	return xxx_messageInfo_IstioOperatorSpec_VersionStatus.Size(m)
}
func (m *IstioOperatorSpec_VersionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioOperatorSpec_VersionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_IstioOperatorSpec_VersionStatus proto.InternalMessageInfo

func (m *IstioOperatorSpec_VersionStatus) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *IstioOperatorSpec_VersionStatus) GetStatus() IstioOperatorSpec_Status {
	if m != nil {
		return m.Status
	}
	return IstioOperatorSpec_NONE
}

func (m *IstioOperatorSpec_VersionStatus) GetStatusString() string {
	if m != nil {
		return m.StatusString
	}
	return ""
}

func (m *IstioOperatorSpec_VersionStatus) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// This is required because synthetic type definition has file rather than package scope.

func init() {
	proto.RegisterEnum("istio.mesh.v1alpha1.IstioOperatorSpec_Status", IstioOperatorSpec_Status_name, IstioOperatorSpec_Status_value)
	proto.RegisterType((*IstioOperatorSpec)(nil), "istio.mesh.v1alpha1.IstioOperatorSpec")
	proto.RegisterMapType((map[string]*IstioOperatorSpec_VersionStatus)(nil), "istio.mesh.v1alpha1.IstioOperatorSpec.ComponentStatusEntry")
	proto.RegisterType((*IstioOperatorSpec_VersionStatus)(nil), "istio.mesh.v1alpha1.IstioOperatorSpec.VersionStatus")
}

func init() { proto.RegisterFile("mesh/v1alpha1/operator.proto", fileDescriptor_03e04c906468a7ff) }

var fileDescriptor_03e04c906468a7ff = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x25, 0xeb, 0xd6, 0x6d, 0x37, 0xfb, 0xc8, 0xcc, 0x24, 0x4c, 0x34, 0x44, 0x35, 0x1e, 0xa8,
	0x10, 0xa4, 0x2c, 0xf0, 0x80, 0xe0, 0x85, 0xd1, 0x05, 0xd6, 0x69, 0x6b, 0xa7, 0x74, 0x4c, 0x02,
	0x1e, 0x22, 0x2f, 0x73, 0x93, 0x68, 0x59, 0x1c, 0xd9, 0x4e, 0xb5, 0x3e, 0xf2, 0x7b, 0xf8, 0x23,
	0xf0, 0xaf, 0x50, 0xec, 0xb4, 0xb4, 0xa8, 0x93, 0x26, 0x78, 0xb3, 0xcf, 0xbd, 0xe7, 0xe4, 0x9e,
	0x93, 0x2b, 0xc3, 0xce, 0x35, 0x15, 0x71, 0x6b, 0xb8, 0x47, 0xd2, 0x3c, 0x26, 0x7b, 0x2d, 0x96,
	0x53, 0x4e, 0x24, 0xe3, 0x4e, 0xce, 0x99, 0x64, 0xe8, 0x7e, 0x22, 0x64, 0xc2, 0x9c, 0xb2, 0xc7,
	0x19, 0xf7, 0xd8, 0xf6, 0x2c, 0x25, 0x64, 0xd9, 0x20, 0x89, 0x34, 0xc1, 0x7e, 0xf4, 0x77, 0xed,
	0x3a, 0x67, 0x19, 0xcd, 0xa4, 0x2e, 0xef, 0x7e, 0x5f, 0x81, 0xad, 0x4e, 0x29, 0xd9, 0xab, 0xbe,
	0xd3, 0xcf, 0x69, 0x88, 0x30, 0x2c, 0xe7, 0x9c, 0x0d, 0x92, 0x94, 0x62, 0x68, 0x18, 0xcd, 0x55,
	0x7f, 0x7c, 0x45, 0x2f, 0x61, 0x3b, 0xc9, 0x84, 0x24, 0x69, 0x1a, 0xe4, 0x24, 0xbc, 0x22, 0x11,
	0x0d, 0x72, 0x22, 0x63, 0x6c, 0xaa, 0x36, 0x54, 0xd5, 0x4e, 0x75, 0xe9, 0x94, 0xc8, 0x18, 0x59,
	0x50, 0x8b, 0x8b, 0x0b, 0xbc, 0xa6, 0x1a, 0xca, 0x63, 0x89, 0x48, 0x12, 0xe1, 0x75, 0x8d, 0x48,
	0x12, 0xa1, 0xa7, 0xb0, 0xc9, 0xa9, 0x60, 0x05, 0x0f, 0x69, 0x20, 0x8a, 0xc1, 0x20, 0xb9, 0xc1,
	0x1b, 0xaa, 0xba, 0x31, 0x86, 0xfb, 0x0a, 0x45, 0xef, 0xc1, 0x2c, 0xfd, 0x04, 0xda, 0x22, 0x6e,
	0x36, 0x8c, 0xa6, 0xe9, 0x3e, 0x76, 0xe6, 0x84, 0xe2, 0x9c, 0x50, 0x11, 0xb7, 0x55, 0x9b, 0x0f,
	0xd7, 0x93, 0x33, 0x3a, 0x02, 0x98, 0x64, 0x20, 0xb0, 0xab, 0x04, 0x9e, 0xcd, 0x15, 0x50, 0xb1,
	0xb4, 0xc7, 0xbd, 0x7d, 0x2a, 0xcb, 0x68, 0xfc, 0x29, 0x36, 0x3a, 0x80, 0xfa, 0x90, 0xa4, 0x05,
	0x15, 0xf8, 0x52, 0xe9, 0x3c, 0x9f, 0xab, 0x73, 0x36, 0xca, 0xe9, 0x09, 0xc9, 0xfb, 0x92, 0x27,
	0x59, 0xd4, 0xc9, 0x24, 0xe5, 0x03, 0x12, 0x52, 0xd7, 0xaf, 0xb8, 0xe8, 0x1b, 0xa0, 0x22, 0x1b,
	0x92, 0x34, 0xb9, 0x24, 0x92, 0x5e, 0x06, 0x95, 0x22, 0xfd, 0x07, 0xc5, 0xad, 0x29, 0x9d, 0x73,
	0x2d, 0xfe, 0x11, 0xea, 0x42, 0x12, 0x59, 0x08, 0xfc, 0xd3, 0x68, 0x18, 0xcd, 0x0d, 0xf7, 0xc5,
	0xed, 0x5e, 0xa7, 0x57, 0xc0, 0xe9, 0x2b, 0x96, 0x5f, 0xb1, 0x51, 0x04, 0xd6, 0xc4, 0x78, 0x50,
	0x29, 0xfe, 0x32, 0x1a, 0xb5, 0xa6, 0xe9, 0xbe, 0xbb, 0xa3, 0xe2, 0x9f, 0x28, 0x15, 0xdd, 0xcb,
	0x24, 0x1f, 0xf9, 0x9b, 0xe1, 0x2c, 0x6a, 0xff, 0x30, 0x60, 0xfd, 0x9c, 0x72, 0x91, 0xb0, 0x4c,
	0x23, 0xe5, 0x32, 0x0e, 0x35, 0x80, 0x0d, 0xbd, 0x8c, 0xd5, 0x15, 0x79, 0x13, 0x73, 0x0b, 0xff,
	0xe3, 0xed, 0x09, 0xac, 0xeb, 0x53, 0x20, 0x54, 0xa2, 0xb8, 0xa6, 0x3e, 0xb3, 0xa6, 0x41, 0x9d,
	0x32, 0xda, 0x86, 0x25, 0xca, 0x39, 0xe3, 0x78, 0x51, 0x15, 0xf5, 0xc5, 0xbe, 0x81, 0xed, 0x79,
	0xb6, 0xca, 0x15, 0xbf, 0xa2, 0xa3, 0x6a, 0xde, 0xf2, 0x88, 0x8e, 0x60, 0x49, 0xfd, 0x59, 0x35,
	0xaa, 0xe9, 0xbe, 0xbe, 0xe3, 0xa8, 0x33, 0x51, 0xf8, 0x5a, 0xe2, 0xed, 0xc2, 0x1b, 0x63, 0xb7,
	0x03, 0xf5, 0x2a, 0x9f, 0x15, 0x58, 0xec, 0xf6, 0xba, 0x9e, 0x75, 0x0f, 0xad, 0xc1, 0xca, 0xe7,
	0xd3, 0x83, 0xfd, 0xb3, 0x4e, 0xf7, 0x93, 0x65, 0xa0, 0x4d, 0x30, 0x7d, 0xaf, 0xdd, 0xeb, 0xb6,
	0x3b, 0xc7, 0x25, 0xb0, 0x80, 0x4c, 0x58, 0x3e, 0xf4, 0xf6, 0x8f, 0xcf, 0x0e, 0xbf, 0x58, 0x35,
	0xb4, 0x0a, 0x4b, 0x9e, 0xef, 0xf7, 0x7c, 0x6b, 0x71, 0xf7, 0x21, 0x3c, 0xb8, 0x65, 0xa3, 0x3e,
	0xec, 0x7c, 0xb5, 0xf5, 0x9c, 0x09, 0x6b, 0x91, 0x3c, 0x69, 0xcd, 0x3c, 0x26, 0x17, 0x75, 0xf5,
	0x86, 0xbc, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x98, 0x76, 0x3c, 0xb3, 0x04, 0x00, 0x00,
}
