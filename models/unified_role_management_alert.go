package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementAlert 
type UnifiedRoleManagementAlert struct {
    Entity
}
// NewUnifiedRoleManagementAlert instantiates a new unifiedRoleManagementAlert and sets the default values.
func NewUnifiedRoleManagementAlert()(*UnifiedRoleManagementAlert) {
    m := &UnifiedRoleManagementAlert{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleManagementAlertFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementAlertFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleManagementAlert(), nil
}
// GetAlertConfiguration gets the alertConfiguration property value. The alertConfiguration property
func (m *UnifiedRoleManagementAlert) GetAlertConfiguration()(UnifiedRoleManagementAlertConfigurationable) {
    val, err := m.GetBackingStore().Get("alertConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UnifiedRoleManagementAlertConfigurationable)
    }
    return nil
}
// GetAlertDefinition gets the alertDefinition property value. The alertDefinition property
func (m *UnifiedRoleManagementAlert) GetAlertDefinition()(UnifiedRoleManagementAlertDefinitionable) {
    val, err := m.GetBackingStore().Get("alertDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UnifiedRoleManagementAlertDefinitionable)
    }
    return nil
}
// GetAlertDefinitionId gets the alertDefinitionId property value. The alertDefinitionId property
func (m *UnifiedRoleManagementAlert) GetAlertDefinitionId()(*string) {
    val, err := m.GetBackingStore().Get("alertDefinitionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAlertIncidents gets the alertIncidents property value. The alertIncidents property
func (m *UnifiedRoleManagementAlert) GetAlertIncidents()([]UnifiedRoleManagementAlertIncidentable) {
    val, err := m.GetBackingStore().Get("alertIncidents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleManagementAlertIncidentable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementAlert) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRoleManagementAlertConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertConfiguration(val.(UnifiedRoleManagementAlertConfigurationable))
        }
        return nil
    }
    res["alertDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRoleManagementAlertDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertDefinition(val.(UnifiedRoleManagementAlertDefinitionable))
        }
        return nil
    }
    res["alertDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertDefinitionId(val)
        }
        return nil
    }
    res["alertIncidents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementAlertIncidentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementAlertIncidentable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleManagementAlertIncidentable)
            }
            m.SetAlertIncidents(res)
        }
        return nil
    }
    res["incidentCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncidentCount(val)
        }
        return nil
    }
    res["isActive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["lastScannedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastScannedDateTime(val)
        }
        return nil
    }
    res["scopeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopeId(val)
        }
        return nil
    }
    res["scopeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopeType(val)
        }
        return nil
    }
    return res
}
// GetIncidentCount gets the incidentCount property value. The incidentCount property
func (m *UnifiedRoleManagementAlert) GetIncidentCount()(*int32) {
    val, err := m.GetBackingStore().Get("incidentCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetIsActive gets the isActive property value. The isActive property
func (m *UnifiedRoleManagementAlert) GetIsActive()(*bool) {
    val, err := m.GetBackingStore().Get("isActive")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *UnifiedRoleManagementAlert) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastScannedDateTime gets the lastScannedDateTime property value. The lastScannedDateTime property
func (m *UnifiedRoleManagementAlert) GetLastScannedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastScannedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetScopeId gets the scopeId property value. The scopeId property
func (m *UnifiedRoleManagementAlert) GetScopeId()(*string) {
    val, err := m.GetBackingStore().Get("scopeId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetScopeType gets the scopeType property value. The scopeType property
func (m *UnifiedRoleManagementAlert) GetScopeType()(*string) {
    val, err := m.GetBackingStore().Get("scopeType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementAlert) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("alertConfiguration", m.GetAlertConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("alertDefinition", m.GetAlertDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("alertDefinitionId", m.GetAlertDefinitionId())
        if err != nil {
            return err
        }
    }
    if m.GetAlertIncidents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlertIncidents()))
        for i, v := range m.GetAlertIncidents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alertIncidents", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("incidentCount", m.GetIncidentCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastScannedDateTime", m.GetLastScannedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scopeId", m.GetScopeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scopeType", m.GetScopeType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertConfiguration sets the alertConfiguration property value. The alertConfiguration property
func (m *UnifiedRoleManagementAlert) SetAlertConfiguration(value UnifiedRoleManagementAlertConfigurationable)() {
    err := m.GetBackingStore().Set("alertConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertDefinition sets the alertDefinition property value. The alertDefinition property
func (m *UnifiedRoleManagementAlert) SetAlertDefinition(value UnifiedRoleManagementAlertDefinitionable)() {
    err := m.GetBackingStore().Set("alertDefinition", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertDefinitionId sets the alertDefinitionId property value. The alertDefinitionId property
func (m *UnifiedRoleManagementAlert) SetAlertDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("alertDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertIncidents sets the alertIncidents property value. The alertIncidents property
func (m *UnifiedRoleManagementAlert) SetAlertIncidents(value []UnifiedRoleManagementAlertIncidentable)() {
    err := m.GetBackingStore().Set("alertIncidents", value)
    if err != nil {
        panic(err)
    }
}
// SetIncidentCount sets the incidentCount property value. The incidentCount property
func (m *UnifiedRoleManagementAlert) SetIncidentCount(value *int32)() {
    err := m.GetBackingStore().Set("incidentCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIsActive sets the isActive property value. The isActive property
func (m *UnifiedRoleManagementAlert) SetIsActive(value *bool)() {
    err := m.GetBackingStore().Set("isActive", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *UnifiedRoleManagementAlert) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastScannedDateTime sets the lastScannedDateTime property value. The lastScannedDateTime property
func (m *UnifiedRoleManagementAlert) SetLastScannedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastScannedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetScopeId sets the scopeId property value. The scopeId property
func (m *UnifiedRoleManagementAlert) SetScopeId(value *string)() {
    err := m.GetBackingStore().Set("scopeId", value)
    if err != nil {
        panic(err)
    }
}
// SetScopeType sets the scopeType property value. The scopeType property
func (m *UnifiedRoleManagementAlert) SetScopeType(value *string)() {
    err := m.GetBackingStore().Set("scopeType", value)
    if err != nil {
        panic(err)
    }
}
// UnifiedRoleManagementAlertable 
type UnifiedRoleManagementAlertable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertConfiguration()(UnifiedRoleManagementAlertConfigurationable)
    GetAlertDefinition()(UnifiedRoleManagementAlertDefinitionable)
    GetAlertDefinitionId()(*string)
    GetAlertIncidents()([]UnifiedRoleManagementAlertIncidentable)
    GetIncidentCount()(*int32)
    GetIsActive()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastScannedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetScopeId()(*string)
    GetScopeType()(*string)
    SetAlertConfiguration(value UnifiedRoleManagementAlertConfigurationable)()
    SetAlertDefinition(value UnifiedRoleManagementAlertDefinitionable)()
    SetAlertDefinitionId(value *string)()
    SetAlertIncidents(value []UnifiedRoleManagementAlertIncidentable)()
    SetIncidentCount(value *int32)()
    SetIsActive(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastScannedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetScopeId(value *string)()
    SetScopeType(value *string)()
}
