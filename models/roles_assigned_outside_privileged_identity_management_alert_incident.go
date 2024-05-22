package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident struct {
    UnifiedRoleManagementAlertIncident
}
// NewRolesAssignedOutsidePrivilegedIdentityManagementAlertIncident instantiates a new RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident and sets the default values.
func NewRolesAssignedOutsidePrivilegedIdentityManagementAlertIncident()(*RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) {
    m := &RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident{
        UnifiedRoleManagementAlertIncident: *NewUnifiedRoleManagementAlertIncident(),
    }
    odataTypeValue := "#microsoft.graph.rolesAssignedOutsidePrivilegedIdentityManagementAlertIncident"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRolesAssignedOutsidePrivilegedIdentityManagementAlertIncidentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRolesAssignedOutsidePrivilegedIdentityManagementAlertIncidentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRolesAssignedOutsidePrivilegedIdentityManagementAlertIncident(), nil
}
// GetAssigneeDisplayName gets the assigneeDisplayName property value. Display name of the subject that the incident applies to.
// returns a *string when successful
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) GetAssigneeDisplayName()(*string) {
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
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) GetAssigneeId()(*string) {
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
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) GetAssigneeUserPrincipalName()(*string) {
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
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) GetAssignmentCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetRoleDefinitionId gets the roleDefinitionId property value. The identifier for the directory role definition that's in scope of this incident.
// returns a *string when successful
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) GetRoleDefinitionId()(*string) {
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
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) GetRoleDisplayName()(*string) {
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
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) GetRoleTemplateId()(*string) {
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
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) SetAssigneeDisplayName(value *string)() {
    err := m.GetBackingStore().Set("assigneeDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAssigneeId sets the assigneeId property value. The identifier of the subject that the incident applies to.
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) SetAssigneeId(value *string)() {
    err := m.GetBackingStore().Set("assigneeId", value)
    if err != nil {
        panic(err)
    }
}
// SetAssigneeUserPrincipalName sets the assigneeUserPrincipalName property value. User principal name of the subject that the incident applies to. Applies to user principals.
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) SetAssigneeUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("assigneeUserPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentCreatedDateTime sets the assignmentCreatedDateTime property value. Date and time of assignment creation.
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) SetAssignmentCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("assignmentCreatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitionId sets the roleDefinitionId property value. The identifier for the directory role definition that's in scope of this incident.
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) SetRoleDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("roleDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDisplayName sets the roleDisplayName property value. The display name for the directory role.
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) SetRoleDisplayName(value *string)() {
    err := m.GetBackingStore().Set("roleDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleTemplateId sets the roleTemplateId property value. The globally unique identifier for the directory role.
func (m *RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) SetRoleTemplateId(value *string)() {
    err := m.GetBackingStore().Set("roleTemplateId", value)
    if err != nil {
        panic(err)
    }
}
type RolesAssignedOutsidePrivilegedIdentityManagementAlertIncidentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertIncidentable
    GetAssigneeDisplayName()(*string)
    GetAssigneeId()(*string)
    GetAssigneeUserPrincipalName()(*string)
    GetAssignmentCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleDefinitionId()(*string)
    GetRoleDisplayName()(*string)
    GetRoleTemplateId()(*string)
    SetAssigneeDisplayName(value *string)()
    SetAssigneeId(value *string)()
    SetAssigneeUserPrincipalName(value *string)()
    SetAssignmentCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleDefinitionId(value *string)()
    SetRoleDisplayName(value *string)()
    SetRoleTemplateId(value *string)()
}
