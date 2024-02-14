package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SequentialActivationRenewalsAlertIncident struct {
    UnifiedRoleManagementAlertIncident
}
// NewSequentialActivationRenewalsAlertIncident instantiates a new SequentialActivationRenewalsAlertIncident and sets the default values.
func NewSequentialActivationRenewalsAlertIncident()(*SequentialActivationRenewalsAlertIncident) {
    m := &SequentialActivationRenewalsAlertIncident{
        UnifiedRoleManagementAlertIncident: *NewUnifiedRoleManagementAlertIncident(),
    }
    odataTypeValue := "#microsoft.graph.sequentialActivationRenewalsAlertIncident"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSequentialActivationRenewalsAlertIncidentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSequentialActivationRenewalsAlertIncidentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSequentialActivationRenewalsAlertIncident(), nil
}
// GetActivationCount gets the activationCount property value. The length of sequential activation of the same role.
// returns a *int32 when successful
func (m *SequentialActivationRenewalsAlertIncident) GetActivationCount()(*int32) {
    val, err := m.GetBackingStore().Get("activationCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAssigneeDisplayName gets the assigneeDisplayName property value. Display name of the subject that the incident applies to.
// returns a *string when successful
func (m *SequentialActivationRenewalsAlertIncident) GetAssigneeDisplayName()(*string) {
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
func (m *SequentialActivationRenewalsAlertIncident) GetAssigneeId()(*string) {
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
func (m *SequentialActivationRenewalsAlertIncident) GetAssigneeUserPrincipalName()(*string) {
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
func (m *SequentialActivationRenewalsAlertIncident) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementAlertIncident.GetFieldDeserializers()
    res["activationCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationCount(val)
        }
        return nil
    }
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
    res["sequenceEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSequenceEndDateTime(val)
        }
        return nil
    }
    res["sequenceStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSequenceStartDateTime(val)
        }
        return nil
    }
    return res
}
// GetRoleDefinitionId gets the roleDefinitionId property value. The identifier for the directory role definition that's in scope of this incident.
// returns a *string when successful
func (m *SequentialActivationRenewalsAlertIncident) GetRoleDefinitionId()(*string) {
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
func (m *SequentialActivationRenewalsAlertIncident) GetRoleDisplayName()(*string) {
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
func (m *SequentialActivationRenewalsAlertIncident) GetRoleTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("roleTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSequenceEndDateTime gets the sequenceEndDateTime property value. End date time of the sequential activation event.
// returns a *Time when successful
func (m *SequentialActivationRenewalsAlertIncident) GetSequenceEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("sequenceEndDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSequenceStartDateTime gets the sequenceStartDateTime property value. Start date time of the sequential activation event.
// returns a *Time when successful
func (m *SequentialActivationRenewalsAlertIncident) GetSequenceStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("sequenceStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SequentialActivationRenewalsAlertIncident) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementAlertIncident.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activationCount", m.GetActivationCount())
        if err != nil {
            return err
        }
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
    {
        err = writer.WriteTimeValue("sequenceEndDateTime", m.GetSequenceEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("sequenceStartDateTime", m.GetSequenceStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivationCount sets the activationCount property value. The length of sequential activation of the same role.
func (m *SequentialActivationRenewalsAlertIncident) SetActivationCount(value *int32)() {
    err := m.GetBackingStore().Set("activationCount", value)
    if err != nil {
        panic(err)
    }
}
// SetAssigneeDisplayName sets the assigneeDisplayName property value. Display name of the subject that the incident applies to.
func (m *SequentialActivationRenewalsAlertIncident) SetAssigneeDisplayName(value *string)() {
    err := m.GetBackingStore().Set("assigneeDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAssigneeId sets the assigneeId property value. The identifier of the subject that the incident applies to.
func (m *SequentialActivationRenewalsAlertIncident) SetAssigneeId(value *string)() {
    err := m.GetBackingStore().Set("assigneeId", value)
    if err != nil {
        panic(err)
    }
}
// SetAssigneeUserPrincipalName sets the assigneeUserPrincipalName property value. User principal name of the subject that the incident applies to. Applies to user principals.
func (m *SequentialActivationRenewalsAlertIncident) SetAssigneeUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("assigneeUserPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitionId sets the roleDefinitionId property value. The identifier for the directory role definition that's in scope of this incident.
func (m *SequentialActivationRenewalsAlertIncident) SetRoleDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("roleDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDisplayName sets the roleDisplayName property value. The display name for the directory role.
func (m *SequentialActivationRenewalsAlertIncident) SetRoleDisplayName(value *string)() {
    err := m.GetBackingStore().Set("roleDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleTemplateId sets the roleTemplateId property value. The globally unique identifier for the directory role.
func (m *SequentialActivationRenewalsAlertIncident) SetRoleTemplateId(value *string)() {
    err := m.GetBackingStore().Set("roleTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// SetSequenceEndDateTime sets the sequenceEndDateTime property value. End date time of the sequential activation event.
func (m *SequentialActivationRenewalsAlertIncident) SetSequenceEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("sequenceEndDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSequenceStartDateTime sets the sequenceStartDateTime property value. Start date time of the sequential activation event.
func (m *SequentialActivationRenewalsAlertIncident) SetSequenceStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("sequenceStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
type SequentialActivationRenewalsAlertIncidentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertIncidentable
    GetActivationCount()(*int32)
    GetAssigneeDisplayName()(*string)
    GetAssigneeId()(*string)
    GetAssigneeUserPrincipalName()(*string)
    GetRoleDefinitionId()(*string)
    GetRoleDisplayName()(*string)
    GetRoleTemplateId()(*string)
    GetSequenceEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSequenceStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetActivationCount(value *int32)()
    SetAssigneeDisplayName(value *string)()
    SetAssigneeId(value *string)()
    SetAssigneeUserPrincipalName(value *string)()
    SetRoleDefinitionId(value *string)()
    SetRoleDisplayName(value *string)()
    SetRoleTemplateId(value *string)()
    SetSequenceEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSequenceStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
