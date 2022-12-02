package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody provides operations to call the getMemberObjects method.
type MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The securityEnabledOnly property
    securityEnabledOnly *bool
}
// NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody instantiates a new MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody and sets the default values.
func NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody()(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) {
    m := &MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["securityEnabledOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityEnabledOnly(val)
        }
        return nil
    }
    return res
}
// GetSecurityEnabledOnly gets the securityEnabledOnly property value. The securityEnabledOnly property
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) GetSecurityEnabledOnly()(*bool) {
    return m.securityEnabledOnly
}
// Serialize serializes information the current object
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("securityEnabledOnly", m.GetSecurityEnabledOnly())
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
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetSecurityEnabledOnly sets the securityEnabledOnly property value. The securityEnabledOnly property
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBody) SetSecurityEnabledOnly(value *bool)() {
    m.securityEnabledOnly = value
}
