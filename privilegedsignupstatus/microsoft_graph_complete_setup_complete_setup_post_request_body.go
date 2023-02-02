package privilegedsignupstatus

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// MicrosoftGraphCompleteSetupCompleteSetupPostRequestBody 
type MicrosoftGraphCompleteSetupCompleteSetupPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The tenantSetupInfo property
    tenantSetupInfo ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantSetupInfoable
}
// NewMicrosoftGraphCompleteSetupCompleteSetupPostRequestBody instantiates a new MicrosoftGraphCompleteSetupCompleteSetupPostRequestBody and sets the default values.
func NewMicrosoftGraphCompleteSetupCompleteSetupPostRequestBody()(*MicrosoftGraphCompleteSetupCompleteSetupPostRequestBody) {
    m := &MicrosoftGraphCompleteSetupCompleteSetupPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftGraphCompleteSetupCompleteSetupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftGraphCompleteSetupCompleteSetupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphCompleteSetupCompleteSetupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftGraphCompleteSetupCompleteSetupPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftGraphCompleteSetupCompleteSetupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["tenantSetupInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTenantSetupInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantSetupInfo(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantSetupInfoable))
        }
        return nil
    }
    return res
}
// GetTenantSetupInfo gets the tenantSetupInfo property value. The tenantSetupInfo property
func (m *MicrosoftGraphCompleteSetupCompleteSetupPostRequestBody) GetTenantSetupInfo()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantSetupInfoable) {
    return m.tenantSetupInfo
}
// Serialize serializes information the current object
func (m *MicrosoftGraphCompleteSetupCompleteSetupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MicrosoftGraphCompleteSetupCompleteSetupPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTenantSetupInfo sets the tenantSetupInfo property value. The tenantSetupInfo property
func (m *MicrosoftGraphCompleteSetupCompleteSetupPostRequestBody) SetTenantSetupInfo(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantSetupInfoable)() {
    m.tenantSetupInfo = value
}
