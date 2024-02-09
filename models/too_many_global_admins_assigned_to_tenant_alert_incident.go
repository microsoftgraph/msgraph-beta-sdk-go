package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TooManyGlobalAdminsAssignedToTenantAlertIncident struct {
    UnifiedRoleManagementAlertIncident
}
// NewTooManyGlobalAdminsAssignedToTenantAlertIncident instantiates a new TooManyGlobalAdminsAssignedToTenantAlertIncident and sets the default values.
func NewTooManyGlobalAdminsAssignedToTenantAlertIncident()(*TooManyGlobalAdminsAssignedToTenantAlertIncident) {
    m := &TooManyGlobalAdminsAssignedToTenantAlertIncident{
        UnifiedRoleManagementAlertIncident: *NewUnifiedRoleManagementAlertIncident(),
    }
    odataTypeValue := "#microsoft.graph.tooManyGlobalAdminsAssignedToTenantAlertIncident"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTooManyGlobalAdminsAssignedToTenantAlertIncidentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTooManyGlobalAdminsAssignedToTenantAlertIncidentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTooManyGlobalAdminsAssignedToTenantAlertIncident(), nil
}
// GetAssigneeDisplayName gets the assigneeDisplayName property value. Display name of the subject that the incident applies to.
// returns a *string when successful
func (m *TooManyGlobalAdminsAssignedToTenantAlertIncident) GetAssigneeDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("assigneeDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssigneeId gets the assigneeId property value. The identifier of the subject that the incident applies to.
// returns a *string when successful
func (m *TooManyGlobalAdminsAssignedToTenantAlertIncident) GetAssigneeId()(*string) {
    val, err := m.GetBackingStore().Get("assigneeId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssigneeUserPrincipalName gets the assigneeUserPrincipalName property value. User principal name of the subject that the incident applies to. Applies to user principals.
// returns a *string when successful
func (m *TooManyGlobalAdminsAssignedToTenantAlertIncident) GetAssigneeUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("assigneeUserPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TooManyGlobalAdminsAssignedToTenantAlertIncident) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementAlertIncident.GetFieldDeserializers()
    res["assigneeDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssigneeDisplayName(val)
        }
        return nil
    }
    res["assigneeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssigneeId(val)
        }
        return nil
    }
    res["assigneeUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssigneeUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TooManyGlobalAdminsAssignedToTenantAlertIncident) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementAlertIncident.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assigneeDisplayName", m.GetAssigneeDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assigneeId", m.GetAssigneeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assigneeUserPrincipalName", m.GetAssigneeUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssigneeDisplayName sets the assigneeDisplayName property value. Display name of the subject that the incident applies to.
func (m *TooManyGlobalAdminsAssignedToTenantAlertIncident) SetAssigneeDisplayName(value *string)() {
    err := m.GetBackingStore().Set("assigneeDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAssigneeId sets the assigneeId property value. The identifier of the subject that the incident applies to.
func (m *TooManyGlobalAdminsAssignedToTenantAlertIncident) SetAssigneeId(value *string)() {
    err := m.GetBackingStore().Set("assigneeId", value)
    if err != nil {
        panic(err)
    }
}
// SetAssigneeUserPrincipalName sets the assigneeUserPrincipalName property value. User principal name of the subject that the incident applies to. Applies to user principals.
func (m *TooManyGlobalAdminsAssignedToTenantAlertIncident) SetAssigneeUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("assigneeUserPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
type TooManyGlobalAdminsAssignedToTenantAlertIncidentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertIncidentable
    GetAssigneeDisplayName()(*string)
    GetAssigneeId()(*string)
    GetAssigneeUserPrincipalName()(*string)
    SetAssigneeDisplayName(value *string)()
    SetAssigneeId(value *string)()
    SetAssigneeUserPrincipalName(value *string)()
}
