package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkMicrophoneConfiguration 
type TeamworkMicrophoneConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    defaultMicrophone *TeamworkPeripheral;
    // True if the configured microphone is optional. False if the microphone is not optional and the health state of the device should be computed.
    isMicrophoneOptional *bool;
    // 
    microphones []TeamworkPeripheral;
}
// NewTeamworkMicrophoneConfiguration instantiates a new teamworkMicrophoneConfiguration and sets the default values.
func NewTeamworkMicrophoneConfiguration()(*TeamworkMicrophoneConfiguration) {
    m := &TeamworkMicrophoneConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkMicrophoneConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultMicrophone gets the defaultMicrophone property value. 
func (m *TeamworkMicrophoneConfiguration) GetDefaultMicrophone()(*TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.defaultMicrophone
    }
}
// GetIsMicrophoneOptional gets the isMicrophoneOptional property value. True if the configured microphone is optional. False if the microphone is not optional and the health state of the device should be computed.
func (m *TeamworkMicrophoneConfiguration) GetIsMicrophoneOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMicrophoneOptional
    }
}
// GetMicrophones gets the microphones property value. 
func (m *TeamworkMicrophoneConfiguration) GetMicrophones()([]TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.microphones
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkMicrophoneConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["defaultMicrophone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultMicrophone(val.(*TeamworkPeripheral))
        }
        return nil
    }
    res["isMicrophoneOptional"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMicrophoneOptional(val)
        }
        return nil
    }
    res["microphones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkPeripheral, len(val))
            for i, v := range val {
                res[i] = *(v.(*TeamworkPeripheral))
            }
            m.SetMicrophones(res)
        }
        return nil
    }
    return res
}
func (m *TeamworkMicrophoneConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkMicrophoneConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("defaultMicrophone", m.GetDefaultMicrophone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isMicrophoneOptional", m.GetIsMicrophoneOptional())
        if err != nil {
            return err
        }
    }
    if m.GetMicrophones() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMicrophones()))
        for i, v := range m.GetMicrophones() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("microphones", cast)
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
func (m *TeamworkMicrophoneConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefaultMicrophone sets the defaultMicrophone property value. 
func (m *TeamworkMicrophoneConfiguration) SetDefaultMicrophone(value *TeamworkPeripheral)() {
    if m != nil {
        m.defaultMicrophone = value
    }
}
// SetIsMicrophoneOptional sets the isMicrophoneOptional property value. True if the configured microphone is optional. False if the microphone is not optional and the health state of the device should be computed.
func (m *TeamworkMicrophoneConfiguration) SetIsMicrophoneOptional(value *bool)() {
    if m != nil {
        m.isMicrophoneOptional = value
    }
}
// SetMicrophones sets the microphones property value. 
func (m *TeamworkMicrophoneConfiguration) SetMicrophones(value []TeamworkPeripheral)() {
    if m != nil {
        m.microphones = value
    }
}
