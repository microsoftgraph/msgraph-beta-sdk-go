package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SafeguardSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    disabledSafeguardProfiles []SafeguardProfile;
}
// Instantiates a new safeguardSettings and sets the default values.
func NewSafeguardSettings()(*SafeguardSettings) {
    m := &SafeguardSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SafeguardSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the disabledSafeguardProfiles property value. 
func (m *SafeguardSettings) GetDisabledSafeguardProfiles()([]SafeguardProfile) {
    if m == nil {
        return nil
    } else {
        return m.disabledSafeguardProfiles
    }
}
// The deserialization information for the current model
func (m *SafeguardSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["disabledSafeguardProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSafeguardProfile() })
        if err != nil {
            return err
        }
        res := make([]SafeguardProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*SafeguardProfile))
        }
        m.SetDisabledSafeguardProfiles(res)
        return nil
    }
    return res
}
func (m *SafeguardSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SafeguardSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SafeguardSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the disabledSafeguardProfiles property value. 
// Parameters:
//  - value : Value to set for the disabledSafeguardProfiles property.
func (m *SafeguardSettings) SetDisabledSafeguardProfiles(value []SafeguardProfile)() {
    m.disabledSafeguardProfiles = value
}
