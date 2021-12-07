package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExpeditedWindowsQualityUpdateSettings 
type ExpeditedWindowsQualityUpdateSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The number of days after installation that forced reboot will happen.
    daysUntilForcedReboot *int32;
    // The release date to identify a quality update.
    qualityUpdateRelease *string;
}
// NewExpeditedWindowsQualityUpdateSettings instantiates a new expeditedWindowsQualityUpdateSettings and sets the default values.
func NewExpeditedWindowsQualityUpdateSettings()(*ExpeditedWindowsQualityUpdateSettings) {
    m := &ExpeditedWindowsQualityUpdateSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExpeditedWindowsQualityUpdateSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDaysUntilForcedReboot gets the daysUntilForcedReboot property value. The number of days after installation that forced reboot will happen.
func (m *ExpeditedWindowsQualityUpdateSettings) GetDaysUntilForcedReboot()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.daysUntilForcedReboot
    }
}
// GetQualityUpdateRelease gets the qualityUpdateRelease property value. The release date to identify a quality update.
func (m *ExpeditedWindowsQualityUpdateSettings) GetQualityUpdateRelease()(*string) {
    if m == nil {
        return nil
    } else {
        return m.qualityUpdateRelease
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExpeditedWindowsQualityUpdateSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["daysUntilForcedReboot"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDaysUntilForcedReboot(val)
        }
        return nil
    }
    res["qualityUpdateRelease"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdateRelease(val)
        }
        return nil
    }
    return res
}
func (m *ExpeditedWindowsQualityUpdateSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExpeditedWindowsQualityUpdateSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("daysUntilForcedReboot", m.GetDaysUntilForcedReboot())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("qualityUpdateRelease", m.GetQualityUpdateRelease())
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
func (m *ExpeditedWindowsQualityUpdateSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDaysUntilForcedReboot sets the daysUntilForcedReboot property value. The number of days after installation that forced reboot will happen.
func (m *ExpeditedWindowsQualityUpdateSettings) SetDaysUntilForcedReboot(value *int32)() {
    if m != nil {
        m.daysUntilForcedReboot = value
    }
}
// SetQualityUpdateRelease sets the qualityUpdateRelease property value. The release date to identify a quality update.
func (m *ExpeditedWindowsQualityUpdateSettings) SetQualityUpdateRelease(value *string)() {
    if m != nil {
        m.qualityUpdateRelease = value
    }
}
