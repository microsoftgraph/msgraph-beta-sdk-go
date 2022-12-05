package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody provides operations to call the importResourceActions method.
type RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The format property
    format *string
    // The overwriteResourceNamespace property
    overwriteResourceNamespace *bool
    // The value property
    value *string
}
// NewRoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody instantiates a new RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody and sets the default values.
func NewRoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody()(*RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody) {
    m := &RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val)
        }
        return nil
    }
    res["overwriteResourceNamespace"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverwriteResourceNamespace(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. The format property
func (m *RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody) GetFormat()(*string) {
    return m.format
}
// GetOverwriteResourceNamespace gets the overwriteResourceNamespace property value. The overwriteResourceNamespace property
func (m *RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody) GetOverwriteResourceNamespace()(*bool) {
    return m.overwriteResourceNamespace
}
// GetValue gets the value property value. The value property
func (m *RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("overwriteResourceNamespace", m.GetOverwriteResourceNamespace())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFormat sets the format property value. The format property
func (m *RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody) SetFormat(value *string)() {
    m.format = value
}
// SetOverwriteResourceNamespace sets the overwriteResourceNamespace property value. The overwriteResourceNamespace property
func (m *RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody) SetOverwriteResourceNamespace(value *bool)() {
    m.overwriteResourceNamespace = value
}
// SetValue sets the value property value. The value property
func (m *RoleManagementDeviceManagementResourceNamespacesItemImportResourceActionsPostRequestBody) SetValue(value *string)() {
    m.value = value
}
