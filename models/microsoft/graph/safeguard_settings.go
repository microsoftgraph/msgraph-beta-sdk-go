package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SafeguardSettings 
type SafeguardSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // List of safeguards to ignore per device.
    disabledSafeguardProfiles []SafeguardProfile;
}
// NewSafeguardSettings instantiates a new safeguardSettings and sets the default values.
func NewSafeguardSettings()(*SafeguardSettings) {
    m := &SafeguardSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SafeguardSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisabledSafeguardProfiles gets the disabledSafeguardProfiles property value. List of safeguards to ignore per device.
func (m *SafeguardSettings) GetDisabledSafeguardProfiles()([]SafeguardProfile) {
    if m == nil {
        return nil
    } else {
        return m.disabledSafeguardProfiles
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SafeguardSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["disabledSafeguardProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSafeguardProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SafeguardProfile, len(val))
            for i, v := range val {
                res[i] = *(v.(*SafeguardProfile))
            }
            m.SetDisabledSafeguardProfiles(res)
        }
        return nil
    }
    return res
}
func (m *SafeguardSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SafeguardSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDisabledSafeguardProfiles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDisabledSafeguardProfiles()))
        for i, v := range m.GetDisabledSafeguardProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("disabledSafeguardProfiles", cast)
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
func (m *SafeguardSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisabledSafeguardProfiles sets the disabledSafeguardProfiles property value. List of safeguards to ignore per device.
func (m *SafeguardSettings) SetDisabledSafeguardProfiles(value []SafeguardProfile)() {
    if m != nil {
        m.disabledSafeguardProfiles = value
    }
}
