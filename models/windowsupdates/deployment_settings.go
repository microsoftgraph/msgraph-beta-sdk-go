package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DeploymentSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeploymentSettings instantiates a new DeploymentSettings and sets the default values.
func NewDeploymentSettings()(*DeploymentSettings) {
    m := &DeploymentSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeploymentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeploymentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeploymentSettings) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *DeploymentSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetContentApplicability gets the contentApplicability property value. Settings for governing whether content is applicable to a device.
// returns a ContentApplicabilitySettingsable when successful
func (m *DeploymentSettings) GetContentApplicability()(ContentApplicabilitySettingsable) {
    val, err := m.GetBackingStore().Get("contentApplicability")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ContentApplicabilitySettingsable)
    }
    return nil
}
// GetExpedite gets the expedite property value. Settings for governing whether updates should be expedited.
// returns a ExpediteSettingsable when successful
func (m *DeploymentSettings) GetExpedite()(ExpediteSettingsable) {
    val, err := m.GetBackingStore().Get("expedite")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ExpediteSettingsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeploymentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentApplicability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentApplicabilitySettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentApplicability(val.(ContentApplicabilitySettingsable))
        }
        return nil
    }
    res["expedite"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExpediteSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpedite(val.(ExpediteSettingsable))
        }
        return nil
    }
    res["monitoring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMonitoringSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonitoring(val.(MonitoringSettingsable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateScheduleSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(ScheduleSettingsable))
        }
        return nil
    }
    res["userExperience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperience(val.(UserExperienceSettingsable))
        }
        return nil
    }
    return res
}
// GetMonitoring gets the monitoring property value. Settings for governing conditions to monitor and automated actions to take.
// returns a MonitoringSettingsable when successful
func (m *DeploymentSettings) GetMonitoring()(MonitoringSettingsable) {
    val, err := m.GetBackingStore().Get("monitoring")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MonitoringSettingsable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DeploymentSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSchedule gets the schedule property value. Settings for governing how and when the content is rolled out.
// returns a ScheduleSettingsable when successful
func (m *DeploymentSettings) GetSchedule()(ScheduleSettingsable) {
    val, err := m.GetBackingStore().Get("schedule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ScheduleSettingsable)
    }
    return nil
}
// GetUserExperience gets the userExperience property value. Settings for governing end user update experience.
// returns a UserExperienceSettingsable when successful
func (m *DeploymentSettings) GetUserExperience()(UserExperienceSettingsable) {
    val, err := m.GetBackingStore().Get("userExperience")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeploymentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentApplicability", m.GetContentApplicability())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("expedite", m.GetExpedite())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("monitoring", m.GetMonitoring())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("schedule", m.GetSchedule())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("userExperience", m.GetUserExperience())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeploymentSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetContentApplicability sets the contentApplicability property value. Settings for governing whether content is applicable to a device.
func (m *DeploymentSettings) SetContentApplicability(value ContentApplicabilitySettingsable)() {
    err := m.GetBackingStore().Set("contentApplicability", value)
    if err != nil {
        panic(err)
    }
}
// SetExpedite sets the expedite property value. Settings for governing whether updates should be expedited.
func (m *DeploymentSettings) SetExpedite(value ExpediteSettingsable)() {
    err := m.GetBackingStore().Set("expedite", value)
    if err != nil {
        panic(err)
    }
}
// SetMonitoring sets the monitoring property value. Settings for governing conditions to monitor and automated actions to take.
func (m *DeploymentSettings) SetMonitoring(value MonitoringSettingsable)() {
    err := m.GetBackingStore().Set("monitoring", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeploymentSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSchedule sets the schedule property value. Settings for governing how and when the content is rolled out.
func (m *DeploymentSettings) SetSchedule(value ScheduleSettingsable)() {
    err := m.GetBackingStore().Set("schedule", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperience sets the userExperience property value. Settings for governing end user update experience.
func (m *DeploymentSettings) SetUserExperience(value UserExperienceSettingsable)() {
    err := m.GetBackingStore().Set("userExperience", value)
    if err != nil {
        panic(err)
    }
}
type DeploymentSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetContentApplicability()(ContentApplicabilitySettingsable)
    GetExpedite()(ExpediteSettingsable)
    GetMonitoring()(MonitoringSettingsable)
    GetOdataType()(*string)
    GetSchedule()(ScheduleSettingsable)
    GetUserExperience()(UserExperienceSettingsable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetContentApplicability(value ContentApplicabilitySettingsable)()
    SetExpedite(value ExpediteSettingsable)()
    SetMonitoring(value MonitoringSettingsable)()
    SetOdataType(value *string)()
    SetSchedule(value ScheduleSettingsable)()
    SetUserExperience(value UserExperienceSettingsable)()
}
