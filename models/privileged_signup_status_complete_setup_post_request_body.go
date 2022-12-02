package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedSignupStatusCompleteSetupPostRequestBody provides operations to call the completeSetup method.
type PrivilegedSignupStatusCompleteSetupPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The tenantSetupInfo property
    tenantSetupInfo TenantSetupInfoable
}
// NewPrivilegedSignupStatusCompleteSetupPostRequestBody instantiates a new PrivilegedSignupStatusCompleteSetupPostRequestBody and sets the default values.
func NewPrivilegedSignupStatusCompleteSetupPostRequestBody()(*PrivilegedSignupStatusCompleteSetupPostRequestBody) {
    m := &PrivilegedSignupStatusCompleteSetupPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrivilegedSignupStatusCompleteSetupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedSignupStatusCompleteSetupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedSignupStatusCompleteSetupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrivilegedSignupStatusCompleteSetupPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedSignupStatusCompleteSetupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["tenantSetupInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTenantSetupInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantSetupInfo(val.(TenantSetupInfoable))
        }
        return nil
    }
    return res
}
// GetTenantSetupInfo gets the tenantSetupInfo property value. The tenantSetupInfo property
func (m *PrivilegedSignupStatusCompleteSetupPostRequestBody) GetTenantSetupInfo()(TenantSetupInfoable) {
    return m.tenantSetupInfo
}
// Serialize serializes information the current object
func (m *PrivilegedSignupStatusCompleteSetupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("tenantSetupInfo", m.GetTenantSetupInfo())
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
func (m *PrivilegedSignupStatusCompleteSetupPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetTenantSetupInfo sets the tenantSetupInfo property value. The tenantSetupInfo property
func (m *PrivilegedSignupStatusCompleteSetupPostRequestBody) SetTenantSetupInfo(value TenantSetupInfoable)() {
    m.tenantSetupInfo = value
}
