package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type LifecycleManagementSettings struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewLifecycleManagementSettings instantiates a new LifecycleManagementSettings and sets the default values.
func NewLifecycleManagementSettings()(*LifecycleManagementSettings) {
    m := &LifecycleManagementSettings{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateLifecycleManagementSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLifecycleManagementSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLifecycleManagementSettings(), nil
}
// GetEmailSettings gets the emailSettings property value. The emailSettings property
// returns a EmailSettingsable when successful
func (m *LifecycleManagementSettings) GetEmailSettings()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmailSettingsable) {
    val, err := m.GetBackingStore().Get("emailSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmailSettingsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LifecycleManagementSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emailSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmailSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSettings(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmailSettingsable))
        }
        return nil
    }
    res["workflowScheduleIntervalInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflowScheduleIntervalInHours(val)
        }
        return nil
    }
    return res
}
// GetWorkflowScheduleIntervalInHours gets the workflowScheduleIntervalInHours property value. The interval in hours at which all workflows running in the tenant should be scheduled for execution. This interval has a minimum value of 1 and a maximum value of 24. The default value is 3 hours.
// returns a *int32 when successful
func (m *LifecycleManagementSettings) GetWorkflowScheduleIntervalInHours()(*int32) {
    val, err := m.GetBackingStore().Get("workflowScheduleIntervalInHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *LifecycleManagementSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("emailSettings", m.GetEmailSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workflowScheduleIntervalInHours", m.GetWorkflowScheduleIntervalInHours())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmailSettings sets the emailSettings property value. The emailSettings property
func (m *LifecycleManagementSettings) SetEmailSettings(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmailSettingsable)() {
    err := m.GetBackingStore().Set("emailSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkflowScheduleIntervalInHours sets the workflowScheduleIntervalInHours property value. The interval in hours at which all workflows running in the tenant should be scheduled for execution. This interval has a minimum value of 1 and a maximum value of 24. The default value is 3 hours.
func (m *LifecycleManagementSettings) SetWorkflowScheduleIntervalInHours(value *int32)() {
    err := m.GetBackingStore().Set("workflowScheduleIntervalInHours", value)
    if err != nil {
        panic(err)
    }
}
type LifecycleManagementSettingsable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmailSettings()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmailSettingsable)
    GetWorkflowScheduleIntervalInHours()(*int32)
    SetEmailSettings(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmailSettingsable)()
    SetWorkflowScheduleIntervalInHours(value *int32)()
}
