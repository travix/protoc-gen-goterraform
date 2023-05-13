package mocks

import (
	mock "github.com/stretchr/testify/mock"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

// MockedFileDescriptor is an autogenerated mock type for the FileDescriptor type.
type MockedFileDescriptor struct {
	protoreflect.FieldDescriptor
	mock.Mock
}

type MockedFileDescriptor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockedFileDescriptor) EXPECT() *MockedFileDescriptor_Expecter {
	return &MockedFileDescriptor_Expecter{mock: &_m.Mock}
}

// Enums provides a mock function with given fields:.
func (_m *MockedFileDescriptor) Enums() protoreflect.EnumDescriptors {
	ret := _m.Called()

	var r0 protoreflect.EnumDescriptors
	if rf, ok := ret.Get(0).(func() protoreflect.EnumDescriptors); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protoreflect.EnumDescriptors)
		}
	}

	return r0
}

// MockedFileDescriptor_Enums_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Enums'.
type MockedFileDescriptor_Enums_Call struct {
	*mock.Call
}

// Enums is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) Enums() *MockedFileDescriptor_Enums_Call {
	return &MockedFileDescriptor_Enums_Call{Call: _e.mock.On("Enums")}
}

func (_c *MockedFileDescriptor_Enums_Call) Run(run func()) *MockedFileDescriptor_Enums_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_Enums_Call) Return(_a0 protoreflect.EnumDescriptors) *MockedFileDescriptor_Enums_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_Enums_Call) RunAndReturn(run func() protoreflect.EnumDescriptors) *MockedFileDescriptor_Enums_Call {
	_c.Call.Return(run)
	return _c
}

// Extensions provides a mock function with given fields:.
func (_m *MockedFileDescriptor) Extensions() protoreflect.ExtensionDescriptors {
	ret := _m.Called()

	var r0 protoreflect.ExtensionDescriptors
	if rf, ok := ret.Get(0).(func() protoreflect.ExtensionDescriptors); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protoreflect.ExtensionDescriptors)
		}
	}

	return r0
}

// MockedFileDescriptor_Extensions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Extensions'.
type MockedFileDescriptor_Extensions_Call struct {
	*mock.Call
}

// Extensions is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) Extensions() *MockedFileDescriptor_Extensions_Call {
	return &MockedFileDescriptor_Extensions_Call{Call: _e.mock.On("Extensions")}
}

func (_c *MockedFileDescriptor_Extensions_Call) Run(run func()) *MockedFileDescriptor_Extensions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_Extensions_Call) Return(_a0 protoreflect.ExtensionDescriptors) *MockedFileDescriptor_Extensions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_Extensions_Call) RunAndReturn(run func() protoreflect.ExtensionDescriptors) *MockedFileDescriptor_Extensions_Call {
	_c.Call.Return(run)
	return _c
}

// FullName provides a mock function with given fields:.
func (_m *MockedFileDescriptor) FullName() protoreflect.FullName {
	ret := _m.Called()

	var r0 protoreflect.FullName
	if rf, ok := ret.Get(0).(func() protoreflect.FullName); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(protoreflect.FullName)
	}

	return r0
}

// MockedFileDescriptor_FullName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FullName'.
type MockedFileDescriptor_FullName_Call struct {
	*mock.Call
}

// FullName is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) FullName() *MockedFileDescriptor_FullName_Call {
	return &MockedFileDescriptor_FullName_Call{Call: _e.mock.On("FullName")}
}

func (_c *MockedFileDescriptor_FullName_Call) Run(run func()) *MockedFileDescriptor_FullName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_FullName_Call) Return(_a0 protoreflect.FullName) *MockedFileDescriptor_FullName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_FullName_Call) RunAndReturn(run func() protoreflect.FullName) *MockedFileDescriptor_FullName_Call {
	_c.Call.Return(run)
	return _c
}

// Imports provides a mock function with given fields:.
func (_m *MockedFileDescriptor) Imports() protoreflect.FileImports {
	ret := _m.Called()

	var r0 protoreflect.FileImports
	if rf, ok := ret.Get(0).(func() protoreflect.FileImports); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protoreflect.FileImports)
		}
	}

	return r0
}

// MockedFileDescriptor_Imports_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Imports'.
type MockedFileDescriptor_Imports_Call struct {
	*mock.Call
}

// Imports is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) Imports() *MockedFileDescriptor_Imports_Call {
	return &MockedFileDescriptor_Imports_Call{Call: _e.mock.On("Imports")}
}

func (_c *MockedFileDescriptor_Imports_Call) Run(run func()) *MockedFileDescriptor_Imports_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_Imports_Call) Return(_a0 protoreflect.FileImports) *MockedFileDescriptor_Imports_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_Imports_Call) RunAndReturn(run func() protoreflect.FileImports) *MockedFileDescriptor_Imports_Call {
	_c.Call.Return(run)
	return _c
}

// Index provides a mock function with given fields:.
func (_m *MockedFileDescriptor) Index() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockedFileDescriptor_Index_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Index'.
type MockedFileDescriptor_Index_Call struct {
	*mock.Call
}

// Index is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) Index() *MockedFileDescriptor_Index_Call {
	return &MockedFileDescriptor_Index_Call{Call: _e.mock.On("Index")}
}

func (_c *MockedFileDescriptor_Index_Call) Run(run func()) *MockedFileDescriptor_Index_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_Index_Call) Return(_a0 int) *MockedFileDescriptor_Index_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_Index_Call) RunAndReturn(run func() int) *MockedFileDescriptor_Index_Call {
	_c.Call.Return(run)
	return _c
}

// IsPlaceholder provides a mock function with given fields:.
func (_m *MockedFileDescriptor) IsPlaceholder() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockedFileDescriptor_IsPlaceholder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsPlaceholder'.
type MockedFileDescriptor_IsPlaceholder_Call struct {
	*mock.Call
}

// IsPlaceholder is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) IsPlaceholder() *MockedFileDescriptor_IsPlaceholder_Call {
	return &MockedFileDescriptor_IsPlaceholder_Call{Call: _e.mock.On("IsPlaceholder")}
}

func (_c *MockedFileDescriptor_IsPlaceholder_Call) Run(run func()) *MockedFileDescriptor_IsPlaceholder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_IsPlaceholder_Call) Return(_a0 bool) *MockedFileDescriptor_IsPlaceholder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_IsPlaceholder_Call) RunAndReturn(run func() bool) *MockedFileDescriptor_IsPlaceholder_Call {
	_c.Call.Return(run)
	return _c
}

// Messages provides a mock function with given fields:.
func (_m *MockedFileDescriptor) Messages() protoreflect.MessageDescriptors {
	ret := _m.Called()

	var r0 protoreflect.MessageDescriptors
	if rf, ok := ret.Get(0).(func() protoreflect.MessageDescriptors); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protoreflect.MessageDescriptors)
		}
	}

	return r0
}

// MockedFileDescriptor_Messages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Messages'.
type MockedFileDescriptor_Messages_Call struct {
	*mock.Call
}

// Messages is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) Messages() *MockedFileDescriptor_Messages_Call {
	return &MockedFileDescriptor_Messages_Call{Call: _e.mock.On("Messages")}
}

func (_c *MockedFileDescriptor_Messages_Call) Run(run func()) *MockedFileDescriptor_Messages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_Messages_Call) Return(_a0 protoreflect.MessageDescriptors) *MockedFileDescriptor_Messages_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_Messages_Call) RunAndReturn(run func() protoreflect.MessageDescriptors) *MockedFileDescriptor_Messages_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:.
func (_m *MockedFileDescriptor) Name() protoreflect.Name {
	ret := _m.Called()

	var r0 protoreflect.Name
	if rf, ok := ret.Get(0).(func() protoreflect.Name); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(protoreflect.Name)
	}

	return r0
}

// MockedFileDescriptor_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'.
type MockedFileDescriptor_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) Name() *MockedFileDescriptor_Name_Call {
	return &MockedFileDescriptor_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockedFileDescriptor_Name_Call) Run(run func()) *MockedFileDescriptor_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_Name_Call) Return(_a0 protoreflect.Name) *MockedFileDescriptor_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_Name_Call) RunAndReturn(run func() protoreflect.Name) *MockedFileDescriptor_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Options provides a mock function with given fields:.
func (_m *MockedFileDescriptor) Options() protoreflect.ProtoMessage {
	ret := _m.Called()

	var r0 protoreflect.ProtoMessage
	if rf, ok := ret.Get(0).(func() protoreflect.ProtoMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protoreflect.ProtoMessage)
		}
	}

	return r0
}

// MockedFileDescriptor_Options_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Options'.
type MockedFileDescriptor_Options_Call struct {
	*mock.Call
}

// Options is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) Options() *MockedFileDescriptor_Options_Call {
	return &MockedFileDescriptor_Options_Call{Call: _e.mock.On("Options")}
}

func (_c *MockedFileDescriptor_Options_Call) Run(run func()) *MockedFileDescriptor_Options_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_Options_Call) Return(_a0 protoreflect.ProtoMessage) *MockedFileDescriptor_Options_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_Options_Call) RunAndReturn(run func() protoreflect.ProtoMessage) *MockedFileDescriptor_Options_Call {
	_c.Call.Return(run)
	return _c
}

// Package provides a mock function with given fields:.
func (_m *MockedFileDescriptor) Package() protoreflect.FullName {
	ret := _m.Called()

	var r0 protoreflect.FullName
	if rf, ok := ret.Get(0).(func() protoreflect.FullName); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(protoreflect.FullName)
	}

	return r0
}

// MockedFileDescriptor_Package_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Package'.
type MockedFileDescriptor_Package_Call struct {
	*mock.Call
}

// Package is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) Package() *MockedFileDescriptor_Package_Call {
	return &MockedFileDescriptor_Package_Call{Call: _e.mock.On("Package")}
}

func (_c *MockedFileDescriptor_Package_Call) Run(run func()) *MockedFileDescriptor_Package_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_Package_Call) Return(_a0 protoreflect.FullName) *MockedFileDescriptor_Package_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_Package_Call) RunAndReturn(run func() protoreflect.FullName) *MockedFileDescriptor_Package_Call {
	_c.Call.Return(run)
	return _c
}

// Parent provides a mock function with given fields:.
func (_m *MockedFileDescriptor) Parent() protoreflect.Descriptor {
	ret := _m.Called()

	var r0 protoreflect.Descriptor
	if rf, ok := ret.Get(0).(func() protoreflect.Descriptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protoreflect.Descriptor)
		}
	}

	return r0
}

// MockedFileDescriptor_Parent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Parent'.
type MockedFileDescriptor_Parent_Call struct {
	*mock.Call
}

// Parent is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) Parent() *MockedFileDescriptor_Parent_Call {
	return &MockedFileDescriptor_Parent_Call{Call: _e.mock.On("Parent")}
}

func (_c *MockedFileDescriptor_Parent_Call) Run(run func()) *MockedFileDescriptor_Parent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_Parent_Call) Return(_a0 protoreflect.Descriptor) *MockedFileDescriptor_Parent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_Parent_Call) RunAndReturn(run func() protoreflect.Descriptor) *MockedFileDescriptor_Parent_Call {
	_c.Call.Return(run)
	return _c
}

// ParentFile provides a mock function with given fields:.
func (_m *MockedFileDescriptor) ParentFile() protoreflect.FileDescriptor {
	ret := _m.Called()

	var r0 protoreflect.FileDescriptor
	if rf, ok := ret.Get(0).(func() protoreflect.FileDescriptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protoreflect.FileDescriptor)
		}
	}

	return r0
}

// MockedFileDescriptor_ParentFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParentFile'.
type MockedFileDescriptor_ParentFile_Call struct {
	*mock.Call
}

// ParentFile is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) ParentFile() *MockedFileDescriptor_ParentFile_Call {
	return &MockedFileDescriptor_ParentFile_Call{Call: _e.mock.On("ParentFile")}
}

func (_c *MockedFileDescriptor_ParentFile_Call) Run(run func()) *MockedFileDescriptor_ParentFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_ParentFile_Call) Return(_a0 protoreflect.FileDescriptor) *MockedFileDescriptor_ParentFile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_ParentFile_Call) RunAndReturn(run func() protoreflect.FileDescriptor) *MockedFileDescriptor_ParentFile_Call {
	_c.Call.Return(run)
	return _c
}

// Path provides a mock function with given fields:.
func (_m *MockedFileDescriptor) Path() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockedFileDescriptor_Path_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Path'.
type MockedFileDescriptor_Path_Call struct {
	*mock.Call
}

// Path is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) Path() *MockedFileDescriptor_Path_Call {
	return &MockedFileDescriptor_Path_Call{Call: _e.mock.On("Path")}
}

func (_c *MockedFileDescriptor_Path_Call) Run(run func()) *MockedFileDescriptor_Path_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_Path_Call) Return(_a0 string) *MockedFileDescriptor_Path_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_Path_Call) RunAndReturn(run func() string) *MockedFileDescriptor_Path_Call {
	_c.Call.Return(run)
	return _c
}

// MockedFileDescriptor_ProtoInternal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProtoInternal'.
type MockedFileDescriptor_ProtoInternal_Call struct {
	*mock.Call
}

// ProtoInternal is a helper method to define mock.On call
//   - _a0 pragma.DoNotImplement
func (_e *MockedFileDescriptor_Expecter) ProtoInternal(_a0 interface{}) *MockedFileDescriptor_ProtoInternal_Call {
	return &MockedFileDescriptor_ProtoInternal_Call{Call: _e.mock.On("ProtoInternal", _a0)}
}

func (_c *MockedFileDescriptor_ProtoInternal_Call) Return() *MockedFileDescriptor_ProtoInternal_Call {
	_c.Call.Return()
	return _c
}

// ProtoType provides a mock function with given fields: _a0.
func (_m *MockedFileDescriptor) ProtoType(_a0 protoreflect.FileDescriptor) {
	_m.Called(_a0)
}

// MockedFileDescriptor_ProtoType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProtoType'.
type MockedFileDescriptor_ProtoType_Call struct {
	*mock.Call
}

// ProtoType is a helper method to define mock.On call
//   - _a0 protoreflect.FileDescriptor
func (_e *MockedFileDescriptor_Expecter) ProtoType(_a0 interface{}) *MockedFileDescriptor_ProtoType_Call {
	return &MockedFileDescriptor_ProtoType_Call{Call: _e.mock.On("ProtoType", _a0)}
}

func (_c *MockedFileDescriptor_ProtoType_Call) Run(run func(_a0 protoreflect.FileDescriptor)) *MockedFileDescriptor_ProtoType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(protoreflect.FileDescriptor))
	})
	return _c
}

func (_c *MockedFileDescriptor_ProtoType_Call) Return() *MockedFileDescriptor_ProtoType_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockedFileDescriptor_ProtoType_Call) RunAndReturn(run func(protoreflect.FileDescriptor)) *MockedFileDescriptor_ProtoType_Call {
	_c.Call.Return(run)
	return _c
}

// Services provides a mock function with given fields:.
func (_m *MockedFileDescriptor) Services() protoreflect.ServiceDescriptors {
	ret := _m.Called()

	var r0 protoreflect.ServiceDescriptors
	if rf, ok := ret.Get(0).(func() protoreflect.ServiceDescriptors); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protoreflect.ServiceDescriptors)
		}
	}

	return r0
}

// MockedFileDescriptor_Services_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Services'.
type MockedFileDescriptor_Services_Call struct {
	*mock.Call
}

// Services is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) Services() *MockedFileDescriptor_Services_Call {
	return &MockedFileDescriptor_Services_Call{Call: _e.mock.On("Services")}
}

func (_c *MockedFileDescriptor_Services_Call) Run(run func()) *MockedFileDescriptor_Services_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_Services_Call) Return(_a0 protoreflect.ServiceDescriptors) *MockedFileDescriptor_Services_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_Services_Call) RunAndReturn(run func() protoreflect.ServiceDescriptors) *MockedFileDescriptor_Services_Call {
	_c.Call.Return(run)
	return _c
}

// SourceLocations provides a mock function with given fields:.
func (_m *MockedFileDescriptor) SourceLocations() protoreflect.SourceLocations {
	ret := _m.Called()

	var r0 protoreflect.SourceLocations
	if rf, ok := ret.Get(0).(func() protoreflect.SourceLocations); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protoreflect.SourceLocations)
		}
	}

	return r0
}

// MockedFileDescriptor_SourceLocations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SourceLocations'.
type MockedFileDescriptor_SourceLocations_Call struct {
	*mock.Call
}

// SourceLocations is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) SourceLocations() *MockedFileDescriptor_SourceLocations_Call {
	return &MockedFileDescriptor_SourceLocations_Call{Call: _e.mock.On("SourceLocations")}
}

func (_c *MockedFileDescriptor_SourceLocations_Call) Run(run func()) *MockedFileDescriptor_SourceLocations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_SourceLocations_Call) Return(_a0 protoreflect.SourceLocations) *MockedFileDescriptor_SourceLocations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_SourceLocations_Call) RunAndReturn(run func() protoreflect.SourceLocations) *MockedFileDescriptor_SourceLocations_Call {
	_c.Call.Return(run)
	return _c
}

// Syntax provides a mock function with given fields:.
func (_m *MockedFileDescriptor) Syntax() protoreflect.Syntax {
	ret := _m.Called()

	var r0 protoreflect.Syntax
	if rf, ok := ret.Get(0).(func() protoreflect.Syntax); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(protoreflect.Syntax)
	}

	return r0
}

// MockedFileDescriptor_Syntax_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Syntax'.
type MockedFileDescriptor_Syntax_Call struct {
	*mock.Call
}

// Syntax is a helper method to define mock.On call.
func (_e *MockedFileDescriptor_Expecter) Syntax() *MockedFileDescriptor_Syntax_Call {
	return &MockedFileDescriptor_Syntax_Call{Call: _e.mock.On("Syntax")}
}

func (_c *MockedFileDescriptor_Syntax_Call) Run(run func()) *MockedFileDescriptor_Syntax_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockedFileDescriptor_Syntax_Call) Return(_a0 protoreflect.Syntax) *MockedFileDescriptor_Syntax_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockedFileDescriptor_Syntax_Call) RunAndReturn(run func() protoreflect.Syntax) *MockedFileDescriptor_Syntax_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockedFileDescriptor interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockedFileDescriptor creates a new instance of MockedFileDescriptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockedFileDescriptor(t mockConstructorTestingTNewMockedFileDescriptor) *MockedFileDescriptor {
	mock := &MockedFileDescriptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
