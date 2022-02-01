package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkPeripheralsHealth 
type TeamworkPeripheralsHealth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The health details about the communication speaker.
    communicationSpeakerHealth *TeamworkPeripheralHealth;
    // The health details about the content camera.
    contentCameraHealth *TeamworkPeripheralHealth;
    // The health details about displays.
    displayHealthCollection []TeamworkPeripheralHealth;
    // The health details about the microphone.
    microphoneHealth *TeamworkPeripheralHealth;
    // The health details about the room camera.
    roomCameraHealth *TeamworkPeripheralHealth;
    // The health details about the speaker.
    speakerHealth *TeamworkPeripheralHealth;
}
// NewTeamworkPeripheralsHealth instantiates a new teamworkPeripheralsHealth and sets the default values.
func NewTeamworkPeripheralsHealth()(*TeamworkPeripheralsHealth) {
    m := &TeamworkPeripheralsHealth{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkPeripheralsHealth) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCommunicationSpeakerHealth gets the communicationSpeakerHealth property value. The health details about the communication speaker.
func (m *TeamworkPeripheralsHealth) GetCommunicationSpeakerHealth()(*TeamworkPeripheralHealth) {
    if m == nil {
        return nil
    } else {
        return m.communicationSpeakerHealth
    }
}
// GetContentCameraHealth gets the contentCameraHealth property value. The health details about the content camera.
func (m *TeamworkPeripheralsHealth) GetContentCameraHealth()(*TeamworkPeripheralHealth) {
    if m == nil {
        return nil
    } else {
        return m.contentCameraHealth
    }
}
// GetDisplayHealthCollection gets the displayHealthCollection property value. The health details about displays.
func (m *TeamworkPeripheralsHealth) GetDisplayHealthCollection()([]TeamworkPeripheralHealth) {
    if m == nil {
        return nil
    } else {
        return m.displayHealthCollection
    }
}
// GetMicrophoneHealth gets the microphoneHealth property value. The health details about the microphone.
func (m *TeamworkPeripheralsHealth) GetMicrophoneHealth()(*TeamworkPeripheralHealth) {
    if m == nil {
        return nil
    } else {
        return m.microphoneHealth
    }
}
// GetRoomCameraHealth gets the roomCameraHealth property value. The health details about the room camera.
func (m *TeamworkPeripheralsHealth) GetRoomCameraHealth()(*TeamworkPeripheralHealth) {
    if m == nil {
        return nil
    } else {
        return m.roomCameraHealth
    }
}
// GetSpeakerHealth gets the speakerHealth property value. The health details about the speaker.
func (m *TeamworkPeripheralsHealth) GetSpeakerHealth()(*TeamworkPeripheralHealth) {
    if m == nil {
        return nil
    } else {
        return m.speakerHealth
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkPeripheralsHealth) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["communicationSpeakerHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheralHealth() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommunicationSpeakerHealth(val.(*TeamworkPeripheralHealth))
        }
        return nil
    }
    res["contentCameraHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheralHealth() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCameraHealth(val.(*TeamworkPeripheralHealth))
        }
        return nil
    }
    res["displayHealthCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheralHealth() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkPeripheralHealth, len(val))
            for i, v := range val {
                res[i] = *(v.(*TeamworkPeripheralHealth))
            }
            m.SetDisplayHealthCollection(res)
        }
        return nil
    }
    res["microphoneHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheralHealth() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophoneHealth(val.(*TeamworkPeripheralHealth))
        }
        return nil
    }
    res["roomCameraHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheralHealth() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoomCameraHealth(val.(*TeamworkPeripheralHealth))
        }
        return nil
    }
    res["speakerHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheralHealth() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeakerHealth(val.(*TeamworkPeripheralHealth))
        }
        return nil
    }
    return res
}
func (m *TeamworkPeripheralsHealth) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkPeripheralsHealth) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("communicationSpeakerHealth", m.GetCommunicationSpeakerHealth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("contentCameraHealth", m.GetContentCameraHealth())
        if err != nil {
            return err
        }
    }
    if m.GetDisplayHealthCollection() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDisplayHealthCollection()))
        for i, v := range m.GetDisplayHealthCollection() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("displayHealthCollection", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("microphoneHealth", m.GetMicrophoneHealth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("roomCameraHealth", m.GetRoomCameraHealth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("speakerHealth", m.GetSpeakerHealth())
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
func (m *TeamworkPeripheralsHealth) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCommunicationSpeakerHealth sets the communicationSpeakerHealth property value. The health details about the communication speaker.
func (m *TeamworkPeripheralsHealth) SetCommunicationSpeakerHealth(value *TeamworkPeripheralHealth)() {
    if m != nil {
        m.communicationSpeakerHealth = value
    }
}
// SetContentCameraHealth sets the contentCameraHealth property value. The health details about the content camera.
func (m *TeamworkPeripheralsHealth) SetContentCameraHealth(value *TeamworkPeripheralHealth)() {
    if m != nil {
        m.contentCameraHealth = value
    }
}
// SetDisplayHealthCollection sets the displayHealthCollection property value. The health details about displays.
func (m *TeamworkPeripheralsHealth) SetDisplayHealthCollection(value []TeamworkPeripheralHealth)() {
    if m != nil {
        m.displayHealthCollection = value
    }
}
// SetMicrophoneHealth sets the microphoneHealth property value. The health details about the microphone.
func (m *TeamworkPeripheralsHealth) SetMicrophoneHealth(value *TeamworkPeripheralHealth)() {
    if m != nil {
        m.microphoneHealth = value
    }
}
// SetRoomCameraHealth sets the roomCameraHealth property value. The health details about the room camera.
func (m *TeamworkPeripheralsHealth) SetRoomCameraHealth(value *TeamworkPeripheralHealth)() {
    if m != nil {
        m.roomCameraHealth = value
    }
}
// SetSpeakerHealth sets the speakerHealth property value. The health details about the speaker.
func (m *TeamworkPeripheralsHealth) SetSpeakerHealth(value *TeamworkPeripheralHealth)() {
    if m != nil {
        m.speakerHealth = value
    }
}
