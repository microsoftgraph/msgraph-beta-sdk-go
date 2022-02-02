package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkTeamsClientConfiguration 
type TeamworkTeamsClientConfiguration struct {
    // The configuration of the Microsoft Teams client user account for a device.
    accountConfiguration *TeamworkAccountConfiguration;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The configuration of Microsoft Teams client features for a device.
    featuresConfiguration *TeamworkFeaturesConfiguration;
}
// NewTeamworkTeamsClientConfiguration instantiates a new teamworkTeamsClientConfiguration and sets the default values.
func NewTeamworkTeamsClientConfiguration()(*TeamworkTeamsClientConfiguration) {
    m := &TeamworkTeamsClientConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAccountConfiguration gets the accountConfiguration property value. The configuration of the Microsoft Teams client user account for a device.
func (m *TeamworkTeamsClientConfiguration) GetAccountConfiguration()(*TeamworkAccountConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.accountConfiguration
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkTeamsClientConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFeaturesConfiguration gets the featuresConfiguration property value. The configuration of Microsoft Teams client features for a device.
func (m *TeamworkTeamsClientConfiguration) GetFeaturesConfiguration()(*TeamworkFeaturesConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.featuresConfiguration
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkTeamsClientConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accountConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkAccountConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountConfiguration(val.(*TeamworkAccountConfiguration))
        }
        return nil
    }
    res["featuresConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkFeaturesConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeaturesConfiguration(val.(*TeamworkFeaturesConfiguration))
        }
        return nil
    }
    return res
}
func (m *TeamworkTeamsClientConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkTeamsClientConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accountConfiguration", m.GetAccountConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("featuresConfiguration", m.GetFeaturesConfiguration())
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
// SetAccountConfiguration sets the accountConfiguration property value. The configuration of the Microsoft Teams client user account for a device.
func (m *TeamworkTeamsClientConfiguration) SetAccountConfiguration(value *TeamworkAccountConfiguration)() {
    if m != nil {
        m.accountConfiguration = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkTeamsClientConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFeaturesConfiguration sets the featuresConfiguration property value. The configuration of Microsoft Teams client features for a device.
func (m *TeamworkTeamsClientConfiguration) SetFeaturesConfiguration(value *TeamworkFeaturesConfiguration)() {
    if m != nil {
        m.featuresConfiguration = value
    }
}
