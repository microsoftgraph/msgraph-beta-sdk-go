package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkDisplayConfiguration 
type TeamworkDisplayConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    configuredDisplays []TeamworkConfiguredPeripheral;
    // 
    displayCount *int32;
    // 
    inBuiltDisplayScreenConfiguration *TeamworkDisplayScreenConfiguration;
    // 
    isContentDuplicationAllowed *bool;
    // 
    isDualDisplayModeEnabled *bool;
}
// NewTeamworkDisplayConfiguration instantiates a new teamworkDisplayConfiguration and sets the default values.
func NewTeamworkDisplayConfiguration()(*TeamworkDisplayConfiguration) {
    m := &TeamworkDisplayConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkDisplayConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConfiguredDisplays gets the configuredDisplays property value. 
func (m *TeamworkDisplayConfiguration) GetConfiguredDisplays()([]TeamworkConfiguredPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.configuredDisplays
    }
}
// GetDisplayCount gets the displayCount property value. 
func (m *TeamworkDisplayConfiguration) GetDisplayCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.displayCount
    }
}
// GetInBuiltDisplayScreenConfiguration gets the inBuiltDisplayScreenConfiguration property value. 
func (m *TeamworkDisplayConfiguration) GetInBuiltDisplayScreenConfiguration()(*TeamworkDisplayScreenConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.inBuiltDisplayScreenConfiguration
    }
}
// GetIsContentDuplicationAllowed gets the isContentDuplicationAllowed property value. 
func (m *TeamworkDisplayConfiguration) GetIsContentDuplicationAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isContentDuplicationAllowed
    }
}
// GetIsDualDisplayModeEnabled gets the isDualDisplayModeEnabled property value. 
func (m *TeamworkDisplayConfiguration) GetIsDualDisplayModeEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDualDisplayModeEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDisplayConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["configuredDisplays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkConfiguredPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkConfiguredPeripheral, len(val))
            for i, v := range val {
                res[i] = *(v.(*TeamworkConfiguredPeripheral))
            }
            m.SetConfiguredDisplays(res)
        }
        return nil
    }
    res["displayCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayCount(val)
        }
        return nil
    }
    res["inBuiltDisplayScreenConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkDisplayScreenConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInBuiltDisplayScreenConfiguration(val.(*TeamworkDisplayScreenConfiguration))
        }
        return nil
    }
    res["isContentDuplicationAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsContentDuplicationAllowed(val)
        }
        return nil
    }
    res["isDualDisplayModeEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDualDisplayModeEnabled(val)
        }
        return nil
    }
    return res
}
func (m *TeamworkDisplayConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkDisplayConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConfiguredDisplays()))
        for i, v := range m.GetConfiguredDisplays() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("configuredDisplays", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("displayCount", m.GetDisplayCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("inBuiltDisplayScreenConfiguration", m.GetInBuiltDisplayScreenConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isContentDuplicationAllowed", m.GetIsContentDuplicationAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDualDisplayModeEnabled", m.GetIsDualDisplayModeEnabled())
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
func (m *TeamworkDisplayConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConfiguredDisplays sets the configuredDisplays property value. 
func (m *TeamworkDisplayConfiguration) SetConfiguredDisplays(value []TeamworkConfiguredPeripheral)() {
    if m != nil {
        m.configuredDisplays = value
    }
}
// SetDisplayCount sets the displayCount property value. 
func (m *TeamworkDisplayConfiguration) SetDisplayCount(value *int32)() {
    if m != nil {
        m.displayCount = value
    }
}
// SetInBuiltDisplayScreenConfiguration sets the inBuiltDisplayScreenConfiguration property value. 
func (m *TeamworkDisplayConfiguration) SetInBuiltDisplayScreenConfiguration(value *TeamworkDisplayScreenConfiguration)() {
    if m != nil {
        m.inBuiltDisplayScreenConfiguration = value
    }
}
// SetIsContentDuplicationAllowed sets the isContentDuplicationAllowed property value. 
func (m *TeamworkDisplayConfiguration) SetIsContentDuplicationAllowed(value *bool)() {
    if m != nil {
        m.isContentDuplicationAllowed = value
    }
}
// SetIsDualDisplayModeEnabled sets the isDualDisplayModeEnabled property value. 
func (m *TeamworkDisplayConfiguration) SetIsDualDisplayModeEnabled(value *bool)() {
    if m != nil {
        m.isDualDisplayModeEnabled = value
    }
}
