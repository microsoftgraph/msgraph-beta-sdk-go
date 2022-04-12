package overridecompliancestate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// OverrideComplianceStateRequestBody provides operations to call the overrideComplianceState method.
type OverrideComplianceStateRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The complianceState property
    complianceState *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministratorConfiguredDeviceComplianceState
    // The remediationUrl property
    remediationUrl *string
}
// NewOverrideComplianceStateRequestBody instantiates a new overrideComplianceStateRequestBody and sets the default values.
func NewOverrideComplianceStateRequestBody()(*OverrideComplianceStateRequestBody) {
    m := &OverrideComplianceStateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOverrideComplianceStateRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOverrideComplianceStateRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOverrideComplianceStateRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OverrideComplianceStateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetComplianceState gets the complianceState property value. The complianceState property
func (m *OverrideComplianceStateRequestBody) GetComplianceState()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministratorConfiguredDeviceComplianceState) {
    if m == nil {
        return nil
    } else {
        return m.complianceState
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OverrideComplianceStateRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *OverrideComplianceStateRequestBody) GetRemediationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediationUrl
    }
}
// Serialize serializes information the current object
func (m *OverrideComplianceStateRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OverrideComplianceStateRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetComplianceState sets the complianceState property value. The complianceState property
func (m *OverrideComplianceStateRequestBody) SetComplianceState(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministratorConfiguredDeviceComplianceState)() {
    if m != nil {
        m.complianceState = value
    }
}
// SetRemediationUrl sets the remediationUrl property value. The remediationUrl property
func (m *OverrideComplianceStateRequestBody) SetRemediationUrl(value *string)() {
    if m != nil {
        m.remediationUrl = value
    }
}
