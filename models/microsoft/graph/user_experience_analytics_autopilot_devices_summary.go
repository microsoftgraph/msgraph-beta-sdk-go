package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsAutopilotDevicesSummary 
type UserExperienceAnalyticsAutopilotDevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The count of intune devices that are not autopilot registerd.
    devicesNotAutopilotRegistered *int32;
    // The count of intune devices not autopilot profile assigned.
    devicesWithoutAutopilotProfileAssigned *int32;
    // The count of windows 10 devices that are Intune and Comanaged.
    totalWindows10DevicesWithoutTenantAttached *int32;
}
// NewUserExperienceAnalyticsAutopilotDevicesSummary instantiates a new userExperienceAnalyticsAutopilotDevicesSummary and sets the default values.
func NewUserExperienceAnalyticsAutopilotDevicesSummary()(*UserExperienceAnalyticsAutopilotDevicesSummary) {
    m := &UserExperienceAnalyticsAutopilotDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDevicesNotAutopilotRegistered gets the devicesNotAutopilotRegistered property value. The count of intune devices that are not autopilot registerd.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetDevicesNotAutopilotRegistered()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.devicesNotAutopilotRegistered
    }
}
// GetDevicesWithoutAutopilotProfileAssigned gets the devicesWithoutAutopilotProfileAssigned property value. The count of intune devices not autopilot profile assigned.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetDevicesWithoutAutopilotProfileAssigned()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.devicesWithoutAutopilotProfileAssigned
    }
}
// GetTotalWindows10DevicesWithoutTenantAttached gets the totalWindows10DevicesWithoutTenantAttached property value. The count of windows 10 devices that are Intune and Comanaged.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetTotalWindows10DevicesWithoutTenantAttached()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalWindows10DevicesWithoutTenantAttached
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["devicesNotAutopilotRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicesNotAutopilotRegistered(val)
        }
        return nil
    }
    res["devicesWithoutAutopilotProfileAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicesWithoutAutopilotProfileAssigned(val)
        }
        return nil
    }
    res["totalWindows10DevicesWithoutTenantAttached"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalWindows10DevicesWithoutTenantAttached(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("devicesNotAutopilotRegistered", m.GetDevicesNotAutopilotRegistered())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("devicesWithoutAutopilotProfileAssigned", m.GetDevicesWithoutAutopilotProfileAssigned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalWindows10DevicesWithoutTenantAttached", m.GetTotalWindows10DevicesWithoutTenantAttached())
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
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDevicesNotAutopilotRegistered sets the devicesNotAutopilotRegistered property value. The count of intune devices that are not autopilot registerd.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) SetDevicesNotAutopilotRegistered(value *int32)() {
    if m != nil {
        m.devicesNotAutopilotRegistered = value
    }
}
// SetDevicesWithoutAutopilotProfileAssigned sets the devicesWithoutAutopilotProfileAssigned property value. The count of intune devices not autopilot profile assigned.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) SetDevicesWithoutAutopilotProfileAssigned(value *int32)() {
    if m != nil {
        m.devicesWithoutAutopilotProfileAssigned = value
    }
}
// SetTotalWindows10DevicesWithoutTenantAttached sets the totalWindows10DevicesWithoutTenantAttached property value. The count of windows 10 devices that are Intune and Comanaged.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) SetTotalWindows10DevicesWithoutTenantAttached(value *int32)() {
    if m != nil {
        m.totalWindows10DevicesWithoutTenantAttached = value
    }
}
