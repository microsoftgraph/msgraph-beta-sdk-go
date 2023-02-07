package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody 
type ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Administrator configured device compliance state Enum
    complianceState *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministratorConfiguredDeviceComplianceState
    // The remediationUrl property
    remediationUrl *string
}
// NewComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody instantiates a new ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody()(*ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody) {
    m := &ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComplianceState gets the complianceState property value. Administrator configured device compliance state Enum
func (m *ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody) GetComplianceState()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministratorConfiguredDeviceComplianceState) {
    return m.complianceState
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["complianceState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseAdministratorConfiguredDeviceComplianceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceState(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministratorConfiguredDeviceComplianceState))
        }
        return nil
    }
    res["remediationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationUrl(val)
        }
        return nil
    }
    return res
}
// GetRemediationUrl gets the remediationUrl property value. The remediationUrl property
func (m *ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody) GetRemediationUrl()(*string) {
    return m.remediationUrl
}
// Serialize serializes information the current object
func (m *ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetComplianceState() != nil {
        cast := (*m.GetComplianceState()).String()
        err := writer.WriteStringValue("complianceState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remediationUrl", m.GetRemediationUrl())
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
func (m *ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComplianceState sets the complianceState property value. Administrator configured device compliance state Enum
func (m *ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody) SetComplianceState(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministratorConfiguredDeviceComplianceState)() {
    m.complianceState = value
}
// SetRemediationUrl sets the remediationUrl property value. The remediationUrl property
func (m *ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBody) SetRemediationUrl(value *string)() {
    m.remediationUrl = value
}
