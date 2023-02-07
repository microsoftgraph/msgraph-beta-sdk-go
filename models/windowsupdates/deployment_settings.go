package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentSettings 
type DeploymentSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Settings for governing whether content is applicable to a device.
    contentApplicability ContentApplicabilitySettingsable
    // Settings for governing whether updates should be expedited.
    expedite ExpediteSettingsable
    // Settings for governing conditions to monitor and automated actions to take.
    monitoring MonitoringSettingsable
    // The OdataType property
    odataType *string
    // Settings for governing how and when the content is rolled out.
    schedule ScheduleSettingsable
    // Settings for governing end user update experience.
    userExperience UserExperienceSettingsable
}
// NewDeploymentSettings instantiates a new deploymentSettings and sets the default values.
func NewDeploymentSettings()(*DeploymentSettings) {
    m := &DeploymentSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeploymentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeploymentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentApplicability gets the contentApplicability property value. Settings for governing whether content is applicable to a device.
func (m *DeploymentSettings) GetContentApplicability()(ContentApplicabilitySettingsable) {
    return m.contentApplicability
}
// GetExpedite gets the expedite property value. Settings for governing whether updates should be expedited.
func (m *DeploymentSettings) GetExpedite()(ExpediteSettingsable) {
    return m.expedite
}
// GetFieldDeserializers the deserialization information for the current model
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
func (m *DeploymentSettings) GetMonitoring()(MonitoringSettingsable) {
    return m.monitoring
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeploymentSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetSchedule gets the schedule property value. Settings for governing how and when the content is rolled out.
func (m *DeploymentSettings) GetSchedule()(ScheduleSettingsable) {
    return m.schedule
}
// GetUserExperience gets the userExperience property value. Settings for governing end user update experience.
func (m *DeploymentSettings) GetUserExperience()(UserExperienceSettingsable) {
    return m.userExperience
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentApplicability sets the contentApplicability property value. Settings for governing whether content is applicable to a device.
func (m *DeploymentSettings) SetContentApplicability(value ContentApplicabilitySettingsable)() {
    m.contentApplicability = value
}
// SetExpedite sets the expedite property value. Settings for governing whether updates should be expedited.
func (m *DeploymentSettings) SetExpedite(value ExpediteSettingsable)() {
    m.expedite = value
}
// SetMonitoring sets the monitoring property value. Settings for governing conditions to monitor and automated actions to take.
func (m *DeploymentSettings) SetMonitoring(value MonitoringSettingsable)() {
    m.monitoring = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeploymentSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSchedule sets the schedule property value. Settings for governing how and when the content is rolled out.
func (m *DeploymentSettings) SetSchedule(value ScheduleSettingsable)() {
    m.schedule = value
}
// SetUserExperience sets the userExperience property value. Settings for governing end user update experience.
func (m *DeploymentSettings) SetUserExperience(value UserExperienceSettingsable)() {
    m.userExperience = value
}
