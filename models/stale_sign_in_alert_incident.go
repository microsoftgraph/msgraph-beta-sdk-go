// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StaleSignInAlertIncident struct {
    UnifiedRoleManagementAlertIncident
}
// NewStaleSignInAlertIncident instantiates a new StaleSignInAlertIncident and sets the default values.
func NewStaleSignInAlertIncident()(*StaleSignInAlertIncident) {
    m := &StaleSignInAlertIncident{
        UnifiedRoleManagementAlertIncident: *NewUnifiedRoleManagementAlertIncident(),
    }
    odataTypeValue := "#microsoft.graph.staleSignInAlertIncident"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateStaleSignInAlertIncidentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStaleSignInAlertIncidentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStaleSignInAlertIncident(), nil
}
// GetAssigneeDisplayName gets the assigneeDisplayName property value. Display name of the subject that the incident applies to.
// returns a *string when successful
func (m *StaleSignInAlertIncident) GetAssigneeDisplayName()(*string) {
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
func (m *StaleSignInAlertIncident) GetAssigneeId()(*string) {
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
func (m *StaleSignInAlertIncident) GetAssigneeUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("assigneeUserPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssignmentCreatedDateTime gets the assignmentCreatedDateTime property value. Date and time of assignment creation.
// returns a *Time when successful
func (m *StaleSignInAlertIncident) GetAssignmentCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("assignmentCreatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StaleSignInAlertIncident) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["assignmentCreatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentCreatedDateTime(val)
        }
        return nil
    }
    res["lastSignInDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSignInDateTime(val)
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
// GetLastSignInDateTime gets the lastSignInDateTime property value. Date and time of last sign in.
// returns a *Time when successful
func (m *StaleSignInAlertIncident) GetLastSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastSignInDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRoleDefinitionId gets the roleDefinitionId property value. The identifier for the directory role definition that's in scope of this incident.
// returns a *string when successful
func (m *StaleSignInAlertIncident) GetRoleDefinitionId()(*string) {
    val, err := m.GetBackingStore().Get("roleDefinitionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleDisplayName gets the roleDisplayName property value. The display name for the directory role.
// returns a *string when successful
func (m *StaleSignInAlertIncident) GetRoleDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("roleDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleTemplateId gets the roleTemplateId property value. The globally unique identifier for the directory role.
// returns a *string when successful
func (m *StaleSignInAlertIncident) GetRoleTemplateId()(*string) {
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
func (m *StaleSignInAlertIncident) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("assignmentCreatedDateTime", m.GetAssignmentCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSignInDateTime", m.GetLastSignInDateTime())
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
// SetAssigneeDisplayName sets the assigneeDisplayName property value. Display name of the subject that the incident applies to.
func (m *StaleSignInAlertIncident) SetAssigneeDisplayName(value *string)() {
    err := m.GetBackingStore().Set("assigneeDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAssigneeId sets the assigneeId property value. The identifier of the subject that the incident applies to.
func (m *StaleSignInAlertIncident) SetAssigneeId(value *string)() {
    err := m.GetBackingStore().Set("assigneeId", value)
    if err != nil {
        panic(err)
    }
}
// SetAssigneeUserPrincipalName sets the assigneeUserPrincipalName property value. User principal name of the subject that the incident applies to. Applies to user principals.
func (m *StaleSignInAlertIncident) SetAssigneeUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("assigneeUserPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentCreatedDateTime sets the assignmentCreatedDateTime property value. Date and time of assignment creation.
func (m *StaleSignInAlertIncident) SetAssignmentCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("assignmentCreatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSignInDateTime sets the lastSignInDateTime property value. Date and time of last sign in.
func (m *StaleSignInAlertIncident) SetLastSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastSignInDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitionId sets the roleDefinitionId property value. The identifier for the directory role definition that's in scope of this incident.
func (m *StaleSignInAlertIncident) SetRoleDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("roleDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDisplayName sets the roleDisplayName property value. The display name for the directory role.
func (m *StaleSignInAlertIncident) SetRoleDisplayName(value *string)() {
    err := m.GetBackingStore().Set("roleDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleTemplateId sets the roleTemplateId property value. The globally unique identifier for the directory role.
func (m *StaleSignInAlertIncident) SetRoleTemplateId(value *string)() {
    err := m.GetBackingStore().Set("roleTemplateId", value)
    if err != nil {
        panic(err)
    }
}
type StaleSignInAlertIncidentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertIncidentable
    GetAssigneeDisplayName()(*string)
    GetAssigneeId()(*string)
    GetAssigneeUserPrincipalName()(*string)
    GetAssignmentCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleDefinitionId()(*string)
    GetRoleDisplayName()(*string)
    GetRoleTemplateId()(*string)
    SetAssigneeDisplayName(value *string)()
    SetAssigneeId(value *string)()
    SetAssigneeUserPrincipalName(value *string)()
    SetAssignmentCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleDefinitionId(value *string)()
    SetRoleDisplayName(value *string)()
    SetRoleTemplateId(value *string)()
}
