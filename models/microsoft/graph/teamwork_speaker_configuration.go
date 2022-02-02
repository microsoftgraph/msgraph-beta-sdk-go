package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkSpeakerConfiguration 
type TeamworkSpeakerConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    defaultCommunicationSpeaker *TeamworkPeripheral;
    // 
    defaultSpeaker *TeamworkPeripheral;
    // True if the communication speaker is optional. Used to compute the health state if the communication speaker is not optional.
    isCommunicationSpeakerOptional *bool;
    // True if the configured speaker is optional. Used to compute the health state if the speaker is not optional.
    isSpeakerOptional *bool;
    // 
    speakers []TeamworkPeripheral;
}
// NewTeamworkSpeakerConfiguration instantiates a new teamworkSpeakerConfiguration and sets the default values.
func NewTeamworkSpeakerConfiguration()(*TeamworkSpeakerConfiguration) {
    m := &TeamworkSpeakerConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkSpeakerConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultCommunicationSpeaker gets the defaultCommunicationSpeaker property value. 
func (m *TeamworkSpeakerConfiguration) GetDefaultCommunicationSpeaker()(*TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.defaultCommunicationSpeaker
    }
}
// GetDefaultSpeaker gets the defaultSpeaker property value. 
func (m *TeamworkSpeakerConfiguration) GetDefaultSpeaker()(*TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.defaultSpeaker
    }
}
// GetIsCommunicationSpeakerOptional gets the isCommunicationSpeakerOptional property value. True if the communication speaker is optional. Used to compute the health state if the communication speaker is not optional.
func (m *TeamworkSpeakerConfiguration) GetIsCommunicationSpeakerOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCommunicationSpeakerOptional
    }
}
// GetIsSpeakerOptional gets the isSpeakerOptional property value. True if the configured speaker is optional. Used to compute the health state if the speaker is not optional.
func (m *TeamworkSpeakerConfiguration) GetIsSpeakerOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSpeakerOptional
    }
}
// GetSpeakers gets the speakers property value. 
func (m *TeamworkSpeakerConfiguration) GetSpeakers()([]TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.speakers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkSpeakerConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["defaultCommunicationSpeaker"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultCommunicationSpeaker(val.(*TeamworkPeripheral))
        }
        return nil
    }
    res["defaultSpeaker"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultSpeaker(val.(*TeamworkPeripheral))
        }
        return nil
    }
    res["isCommunicationSpeakerOptional"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCommunicationSpeakerOptional(val)
        }
        return nil
    }
    res["isSpeakerOptional"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSpeakerOptional(val)
        }
        return nil
    }
    res["speakers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkPeripheral, len(val))
            for i, v := range val {
                res[i] = *(v.(*TeamworkPeripheral))
            }
            m.SetSpeakers(res)
        }
        return nil
    }
    return res
}
func (m *TeamworkSpeakerConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkSpeakerConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("defaultCommunicationSpeaker", m.GetDefaultCommunicationSpeaker())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("defaultSpeaker", m.GetDefaultSpeaker())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCommunicationSpeakerOptional", m.GetIsCommunicationSpeakerOptional())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSpeakerOptional", m.GetIsSpeakerOptional())
        if err != nil {
            return err
        }
    }
    if m.GetSpeakers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSpeakers()))
        for i, v := range m.GetSpeakers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("speakers", cast)
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
func (m *TeamworkSpeakerConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefaultCommunicationSpeaker sets the defaultCommunicationSpeaker property value. 
func (m *TeamworkSpeakerConfiguration) SetDefaultCommunicationSpeaker(value *TeamworkPeripheral)() {
    if m != nil {
        m.defaultCommunicationSpeaker = value
    }
}
// SetDefaultSpeaker sets the defaultSpeaker property value. 
func (m *TeamworkSpeakerConfiguration) SetDefaultSpeaker(value *TeamworkPeripheral)() {
    if m != nil {
        m.defaultSpeaker = value
    }
}
// SetIsCommunicationSpeakerOptional sets the isCommunicationSpeakerOptional property value. True if the communication speaker is optional. Used to compute the health state if the communication speaker is not optional.
func (m *TeamworkSpeakerConfiguration) SetIsCommunicationSpeakerOptional(value *bool)() {
    if m != nil {
        m.isCommunicationSpeakerOptional = value
    }
}
// SetIsSpeakerOptional sets the isSpeakerOptional property value. True if the configured speaker is optional. Used to compute the health state if the speaker is not optional.
func (m *TeamworkSpeakerConfiguration) SetIsSpeakerOptional(value *bool)() {
    if m != nil {
        m.isSpeakerOptional = value
    }
}
// SetSpeakers sets the speakers property value. 
func (m *TeamworkSpeakerConfiguration) SetSpeakers(value []TeamworkPeripheral)() {
    if m != nil {
        m.speakers = value
    }
}
