package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RedundantAssignmentAlertIncident 
type RedundantAssignmentAlertIncident struct {
    UnifiedRoleManagementAlertIncident
}
// NewRedundantAssignmentAlertIncident instantiates a new RedundantAssignmentAlertIncident and sets the default values.
func NewRedundantAssignmentAlertIncident()(*RedundantAssignmentAlertIncident) {
    m := &RedundantAssignmentAlertIncident{
        UnifiedRoleManagementAlertIncident: *NewUnifiedRoleManagementAlertIncident(),
    }
    odataTypeValue := "#microsoft.graph.redundantAssignmentAlertIncident"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRedundantAssignmentAlertIncidentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRedundantAssignmentAlertIncidentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRedundantAssignmentAlertIncident(), nil
}
// GetAssigneeDisplayName gets the assigneeDisplayName property value. The assigneeDisplayName property
func (m *RedundantAssignmentAlertIncident) GetAssigneeDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("assigneeDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssigneeId gets the assigneeId property value. The assigneeId property
func (m *RedundantAssignmentAlertIncident) GetAssigneeId()(*string) {
    val, err := m.GetBackingStore().Get("assigneeId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssigneeUserPrincipalName gets the assigneeUserPrincipalName property value. The assigneeUserPrincipalName property
func (m *RedundantAssignmentAlertIncident) GetAssigneeUserPrincipalName()(*string) {
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
func (m *RedundantAssignmentAlertIncident) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["lastActivationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivationDateTime(val)
        }
        return nil
    }
    res["roleDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinitionId(val)
        }
        return nil
    }
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
// GetLastActivationDateTime gets the lastActivationDateTime property value. The lastActivationDateTime property
func (m *RedundantAssignmentAlertIncident) GetLastActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastActivationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRoleDefinitionId gets the roleDefinitionId property value. The roleDefinitionId property
func (m *RedundantAssignmentAlertIncident) GetRoleDefinitionId()(*string) {
    val, err := m.GetBackingStore().Get("roleDefinitionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleDisplayName gets the roleDisplayName property value. The roleDisplayName property
func (m *RedundantAssignmentAlertIncident) GetRoleDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("roleDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleTemplateId gets the roleTemplateId property value. The roleTemplateId property
func (m *RedundantAssignmentAlertIncident) GetRoleTemplateId()(*string) {
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
func (m *RedundantAssignmentAlertIncident) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteTimeValue("lastActivationDateTime", m.GetLastActivationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleDefinitionId", m.GetRoleDefinitionId())
        if err != nil {
            return err
        }
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
// SetAssigneeDisplayName sets the assigneeDisplayName property value. The assigneeDisplayName property
func (m *RedundantAssignmentAlertIncident) SetAssigneeDisplayName(value *string)() {
    err := m.GetBackingStore().Set("assigneeDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAssigneeId sets the assigneeId property value. The assigneeId property
func (m *RedundantAssignmentAlertIncident) SetAssigneeId(value *string)() {
    err := m.GetBackingStore().Set("assigneeId", value)
    if err != nil {
        panic(err)
    }
}
// SetAssigneeUserPrincipalName sets the assigneeUserPrincipalName property value. The assigneeUserPrincipalName property
func (m *RedundantAssignmentAlertIncident) SetAssigneeUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("assigneeUserPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActivationDateTime sets the lastActivationDateTime property value. The lastActivationDateTime property
func (m *RedundantAssignmentAlertIncident) SetLastActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastActivationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitionId sets the roleDefinitionId property value. The roleDefinitionId property
func (m *RedundantAssignmentAlertIncident) SetRoleDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("roleDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDisplayName sets the roleDisplayName property value. The roleDisplayName property
func (m *RedundantAssignmentAlertIncident) SetRoleDisplayName(value *string)() {
    err := m.GetBackingStore().Set("roleDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleTemplateId sets the roleTemplateId property value. The roleTemplateId property
func (m *RedundantAssignmentAlertIncident) SetRoleTemplateId(value *string)() {
    err := m.GetBackingStore().Set("roleTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// RedundantAssignmentAlertIncidentable 
type RedundantAssignmentAlertIncidentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertIncidentable
    GetAssigneeDisplayName()(*string)
    GetAssigneeId()(*string)
    GetAssigneeUserPrincipalName()(*string)
    GetLastActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleDefinitionId()(*string)
    GetRoleDisplayName()(*string)
    GetRoleTemplateId()(*string)
    SetAssigneeDisplayName(value *string)()
    SetAssigneeId(value *string)()
    SetAssigneeUserPrincipalName(value *string)()
    SetLastActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleDefinitionId(value *string)()
    SetRoleDisplayName(value *string)()
    SetRoleTemplateId(value *string)()
}
