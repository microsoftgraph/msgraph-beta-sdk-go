package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NoMfaOnRoleActivationAlertIncident 
type NoMfaOnRoleActivationAlertIncident struct {
    UnifiedRoleManagementAlertIncident
    // The OdataType property
    OdataType *string
}
// NewNoMfaOnRoleActivationAlertIncident instantiates a new noMfaOnRoleActivationAlertIncident and sets the default values.
func NewNoMfaOnRoleActivationAlertIncident()(*NoMfaOnRoleActivationAlertIncident) {
    m := &NoMfaOnRoleActivationAlertIncident{
        UnifiedRoleManagementAlertIncident: *NewUnifiedRoleManagementAlertIncident(),
    }
    odataTypeValue := "#microsoft.graph.noMfaOnRoleActivationAlertIncident"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateNoMfaOnRoleActivationAlertIncidentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNoMfaOnRoleActivationAlertIncidentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNoMfaOnRoleActivationAlertIncident(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NoMfaOnRoleActivationAlertIncident) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementAlertIncident.GetFieldDeserializers()
    res["roleDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDisplayName(val)
        }
        return nil
    }
    res["roleTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleTemplateId(val)
        }
        return nil
    }
    return res
}
// GetRoleDisplayName gets the roleDisplayName property value. The name of the Azure AD directory role.
func (m *NoMfaOnRoleActivationAlertIncident) GetRoleDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("roleDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleTemplateId gets the roleTemplateId property value. The globally unique identifier for a directory role.
func (m *NoMfaOnRoleActivationAlertIncident) GetRoleTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("roleTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *NoMfaOnRoleActivationAlertIncident) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementAlertIncident.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("roleDisplayName", m.GetRoleDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleTemplateId", m.GetRoleTemplateId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRoleDisplayName sets the roleDisplayName property value. The name of the Azure AD directory role.
func (m *NoMfaOnRoleActivationAlertIncident) SetRoleDisplayName(value *string)() {
    err := m.GetBackingStore().Set("roleDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleTemplateId sets the roleTemplateId property value. The globally unique identifier for a directory role.
func (m *NoMfaOnRoleActivationAlertIncident) SetRoleTemplateId(value *string)() {
    err := m.GetBackingStore().Set("roleTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// NoMfaOnRoleActivationAlertIncidentable 
type NoMfaOnRoleActivationAlertIncidentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertIncidentable
    GetRoleDisplayName()(*string)
    GetRoleTemplateId()(*string)
    SetRoleDisplayName(value *string)()
    SetRoleTemplateId(value *string)()
}
