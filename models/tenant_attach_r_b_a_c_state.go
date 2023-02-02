package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantAttachRBACState represents result of GetState API.
type TenantAttachRBACState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether the tenant is enabled for Tenant Attach with role management.  TRUE if enabled, FALSE if the Tenant Attach with rolemanagement is disabled.
    enabled *bool
    // The OdataType property
    odataType *string
}
// NewTenantAttachRBACState instantiates a new tenantAttachRBACState and sets the default values.
func NewTenantAttachRBACState()(*TenantAttachRBACState) {
    m := &TenantAttachRBACState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTenantAttachRBACStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantAttachRBACStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantAttachRBACState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantAttachRBACState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. Indicates whether the tenant is enabled for Tenant Attach with role management.  TRUE if enabled, FALSE if the Tenant Attach with rolemanagement is disabled.
func (m *TenantAttachRBACState) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantAttachRBACState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TenantAttachRBACState) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TenantAttachRBACState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *TenantAttachRBACState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. Indicates whether the tenant is enabled for Tenant Attach with role management.  TRUE if enabled, FALSE if the Tenant Attach with rolemanagement is disabled.
func (m *TenantAttachRBACState) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TenantAttachRBACState) SetOdataType(value *string)() {
    m.odataType = value
}
