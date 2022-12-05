package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody provides operations to call the getRoleScopeTagsById method.
type DeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The roleScopeTagIds property
    roleScopeTagIds []string
}
// NewDeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody instantiates a new DeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody and sets the default values.
func NewDeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody()(*DeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody) {
    m := &DeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    return res
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. The roleScopeTagIds property
func (m *DeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// Serialize serializes information the current object
func (m *DeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRoleScopeTagIds() != nil {
        err := writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
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
func (m *DeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. The roleScopeTagIds property
func (m *DeviceManagementRoleScopeTagsGetRoleScopeTagsByIdPostRequestBody) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
