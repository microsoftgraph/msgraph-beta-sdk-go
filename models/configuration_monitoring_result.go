// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConfigurationMonitoringResult struct {
    Entity
}
// NewConfigurationMonitoringResult instantiates a new ConfigurationMonitoringResult and sets the default values.
func NewConfigurationMonitoringResult()(*ConfigurationMonitoringResult) {
    m := &ConfigurationMonitoringResult{
        Entity: *NewEntity(),
    }
    return m
}
// CreateConfigurationMonitoringResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConfigurationMonitoringResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfigurationMonitoringResult(), nil
}
// GetDriftsCount gets the driftsCount property value. The driftsCount property
// returns a *int32 when successful
func (m *ConfigurationMonitoringResult) GetDriftsCount()(*int32) {
    val, err := m.GetBackingStore().Get("driftsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetErrorDetails gets the errorDetails property value. The errorDetails property
// returns a []ErrorDetailable when successful
func (m *ConfigurationMonitoringResult) GetErrorDetails()([]ErrorDetailable) {
    val, err := m.GetBackingStore().Get("errorDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ErrorDetailable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConfigurationMonitoringResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["driftsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriftsCount(val)
        }
        return nil
    }
    res["errorDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateErrorDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ErrorDetailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ErrorDetailable)
                }
            }
            m.SetErrorDetails(res)
        }
        return nil
    }
    res["monitorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonitorId(val)
        }
        return nil
    }
    res["runCompletionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunCompletionDateTime(val)
        }
        return nil
    }
    res["runInitiationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunInitiationDateTime(val)
        }
        return nil
    }
    res["runStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMonitorRunStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunStatus(val.(*MonitorRunStatus))
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetMonitorId gets the monitorId property value. The monitorId property
// returns a *string when successful
func (m *ConfigurationMonitoringResult) GetMonitorId()(*string) {
    val, err := m.GetBackingStore().Get("monitorId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRunCompletionDateTime gets the runCompletionDateTime property value. The runCompletionDateTime property
// returns a *Time when successful
func (m *ConfigurationMonitoringResult) GetRunCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("runCompletionDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRunInitiationDateTime gets the runInitiationDateTime property value. The runInitiationDateTime property
// returns a *Time when successful
func (m *ConfigurationMonitoringResult) GetRunInitiationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("runInitiationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRunStatus gets the runStatus property value. The runStatus property
// returns a *MonitorRunStatus when successful
func (m *ConfigurationMonitoringResult) GetRunStatus()(*MonitorRunStatus) {
    val, err := m.GetBackingStore().Get("runStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MonitorRunStatus)
    }
    return nil
}
// GetTenantId gets the tenantId property value. The tenantId property
// returns a *string when successful
func (m *ConfigurationMonitoringResult) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ConfigurationMonitoringResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRunStatus() != nil {
        cast := (*m.GetRunStatus()).String()
        err = writer.WriteStringValue("runStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDriftsCount sets the driftsCount property value. The driftsCount property
func (m *ConfigurationMonitoringResult) SetDriftsCount(value *int32)() {
    err := m.GetBackingStore().Set("driftsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorDetails sets the errorDetails property value. The errorDetails property
func (m *ConfigurationMonitoringResult) SetErrorDetails(value []ErrorDetailable)() {
    err := m.GetBackingStore().Set("errorDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetMonitorId sets the monitorId property value. The monitorId property
func (m *ConfigurationMonitoringResult) SetMonitorId(value *string)() {
    err := m.GetBackingStore().Set("monitorId", value)
    if err != nil {
        panic(err)
    }
}
// SetRunCompletionDateTime sets the runCompletionDateTime property value. The runCompletionDateTime property
func (m *ConfigurationMonitoringResult) SetRunCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("runCompletionDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRunInitiationDateTime sets the runInitiationDateTime property value. The runInitiationDateTime property
func (m *ConfigurationMonitoringResult) SetRunInitiationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("runInitiationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRunStatus sets the runStatus property value. The runStatus property
func (m *ConfigurationMonitoringResult) SetRunStatus(value *MonitorRunStatus)() {
    err := m.GetBackingStore().Set("runStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ConfigurationMonitoringResult) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
type ConfigurationMonitoringResultable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDriftsCount()(*int32)
    GetErrorDetails()([]ErrorDetailable)
    GetMonitorId()(*string)
    GetRunCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRunInitiationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRunStatus()(*MonitorRunStatus)
    GetTenantId()(*string)
    SetDriftsCount(value *int32)()
    SetErrorDetails(value []ErrorDetailable)()
    SetMonitorId(value *string)()
    SetRunCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRunInitiationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRunStatus(value *MonitorRunStatus)()
    SetTenantId(value *string)()
}
