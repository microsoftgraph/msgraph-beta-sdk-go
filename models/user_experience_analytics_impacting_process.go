package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsImpactingProcess the user experience analytics top impacting process entity.
type UserExperienceAnalyticsImpactingProcess struct {
    Entity
    // The category of impacting process.
    category *string
    // The description of process.
    description *string
    // The unique identifier of the impacted device.
    deviceId *string
    // The impact value of the process. Valid values 0 to 1.79769313486232E+308
    impactValue *float64
    // The process name.
    processName *string
    // The publisher of the process.
    publisher *string
}
// NewUserExperienceAnalyticsImpactingProcess instantiates a new userExperienceAnalyticsImpactingProcess and sets the default values.
func NewUserExperienceAnalyticsImpactingProcess()(*UserExperienceAnalyticsImpactingProcess) {
    m := &UserExperienceAnalyticsImpactingProcess{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.userExperienceAnalyticsImpactingProcess";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserExperienceAnalyticsImpactingProcessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsImpactingProcessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsImpactingProcess(), nil
}
// GetCategory gets the category property value. The category of impacting process.
func (m *UserExperienceAnalyticsImpactingProcess) GetCategory()(*string) {
    return m.category
}
// GetDescription gets the description property value. The description of process.
func (m *UserExperienceAnalyticsImpactingProcess) GetDescription()(*string) {
    return m.description
}
// GetDeviceId gets the deviceId property value. The unique identifier of the impacted device.
func (m *UserExperienceAnalyticsImpactingProcess) GetDeviceId()(*string) {
    return m.deviceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsImpactingProcess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCategory)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["deviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceId)
    res["impactValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetImpactValue)
    res["processName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProcessName)
    res["publisher"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPublisher)
    return res
}
// GetImpactValue gets the impactValue property value. The impact value of the process. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsImpactingProcess) GetImpactValue()(*float64) {
    return m.impactValue
}
// GetProcessName gets the processName property value. The process name.
func (m *UserExperienceAnalyticsImpactingProcess) GetProcessName()(*string) {
    return m.processName
}
// GetPublisher gets the publisher property value. The publisher of the process.
func (m *UserExperienceAnalyticsImpactingProcess) GetPublisher()(*string) {
    return m.publisher
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsImpactingProcess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("impactValue", m.GetImpactValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("processName", m.GetProcessName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategory sets the category property value. The category of impacting process.
func (m *UserExperienceAnalyticsImpactingProcess) SetCategory(value *string)() {
    m.category = value
}
// SetDescription sets the description property value. The description of process.
func (m *UserExperienceAnalyticsImpactingProcess) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceId sets the deviceId property value. The unique identifier of the impacted device.
func (m *UserExperienceAnalyticsImpactingProcess) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetImpactValue sets the impactValue property value. The impact value of the process. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsImpactingProcess) SetImpactValue(value *float64)() {
    m.impactValue = value
}
// SetProcessName sets the processName property value. The process name.
func (m *UserExperienceAnalyticsImpactingProcess) SetProcessName(value *string)() {
    m.processName = value
}
// SetPublisher sets the publisher property value. The publisher of the process.
func (m *UserExperienceAnalyticsImpactingProcess) SetPublisher(value *string)() {
    m.publisher = value
}
