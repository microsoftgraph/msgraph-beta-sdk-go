package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppCallsItemChangeScreenSharingRolePostRequestBody provides operations to call the changeScreenSharingRole method.
type AppCallsItemChangeScreenSharingRolePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The role property
    role *ScreenSharingRole
}
// NewAppCallsItemChangeScreenSharingRolePostRequestBody instantiates a new AppCallsItemChangeScreenSharingRolePostRequestBody and sets the default values.
func NewAppCallsItemChangeScreenSharingRolePostRequestBody()(*AppCallsItemChangeScreenSharingRolePostRequestBody) {
    m := &AppCallsItemChangeScreenSharingRolePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppCallsItemChangeScreenSharingRolePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppCallsItemChangeScreenSharingRolePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppCallsItemChangeScreenSharingRolePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppCallsItemChangeScreenSharingRolePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppCallsItemChangeScreenSharingRolePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseScreenSharingRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*ScreenSharingRole))
        }
        return nil
    }
    return res
}
// GetRole gets the role property value. The role property
func (m *AppCallsItemChangeScreenSharingRolePostRequestBody) GetRole()(*ScreenSharingRole) {
    return m.role
}
// Serialize serializes information the current object
func (m *AppCallsItemChangeScreenSharingRolePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err := writer.WriteStringValue("role", &cast)
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
func (m *AppCallsItemChangeScreenSharingRolePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetRole sets the role property value. The role property
func (m *AppCallsItemChangeScreenSharingRolePostRequestBody) SetRole(value *ScreenSharingRole)() {
    m.role = value
}
