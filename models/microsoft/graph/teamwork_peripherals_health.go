package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkPeripheralsHealth provides operations to manage the teamwork singleton.
type TeamworkPeripheralsHealth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The health details about the communication speaker.
    communicationSpeakerHealth TeamworkPeripheralHealthable;
    // The health details about the content camera.
    contentCameraHealth TeamworkPeripheralHealthable;
    // The health details about displays.
    displayHealthCollection []TeamworkPeripheralHealthable;
    // The health details about the microphone.
    microphoneHealth TeamworkPeripheralHealthable;
    // The health details about the room camera.
    roomCameraHealth TeamworkPeripheralHealthable;
    // The health details about the speaker.
    speakerHealth TeamworkPeripheralHealthable;
}
// NewTeamworkPeripheralsHealth instantiates a new teamworkPeripheralsHealth and sets the default values.
func NewTeamworkPeripheralsHealth()(*TeamworkPeripheralsHealth) {
    m := &TeamworkPeripheralsHealth{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkPeripheralsHealthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkPeripheralsHealthFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamworkPeripheralsHealth(), nil
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
func (m *TeamworkPeripheralsHealth) GetCommunicationSpeakerHealth()(TeamworkPeripheralHealthable) {
    if m == nil {
        return nil
    } else {
        return m.communicationSpeakerHealth
    }
}
// GetContentCameraHealth gets the contentCameraHealth property value. The health details about the content camera.
func (m *TeamworkPeripheralsHealth) GetContentCameraHealth()(TeamworkPeripheralHealthable) {
    if m == nil {
        return nil
    } else {
        return m.contentCameraHealth
    }
}
// GetDisplayHealthCollection gets the displayHealthCollection property value. The health details about displays.
func (m *TeamworkPeripheralsHealth) GetDisplayHealthCollection()([]TeamworkPeripheralHealthable) {
    if m == nil {
        return nil
    } else {
        return m.displayHealthCollection
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkPeripheralsHealth) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["communicationSpeakerHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommunicationSpeakerHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    res["contentCameraHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCameraHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    res["displayHealthCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkPeripheralHealthable, len(val))
            for i, v := range val {
                res[i] = v.(TeamworkPeripheralHealthable)
            }
            m.SetDisplayHealthCollection(res)
        }
        return nil
    }
    res["microphoneHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophoneHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    res["roomCameraHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoomCameraHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    res["speakerHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeakerHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    return res
}
// GetMicrophoneHealth gets the microphoneHealth property value. The health details about the microphone.
func (m *TeamworkPeripheralsHealth) GetMicrophoneHealth()(TeamworkPeripheralHealthable) {
    if m == nil {
        return nil
    } else {
        return m.microphoneHealth
    }
}
// GetRoomCameraHealth gets the roomCameraHealth property value. The health details about the room camera.
func (m *TeamworkPeripheralsHealth) GetRoomCameraHealth()(TeamworkPeripheralHealthable) {
    if m == nil {
        return nil
    } else {
        return m.roomCameraHealth
    }
}
// GetSpeakerHealth gets the speakerHealth property value. The health details about the speaker.
func (m *TeamworkPeripheralsHealth) GetSpeakerHealth()(TeamworkPeripheralHealthable) {
    if m == nil {
        return nil
    } else {
        return m.speakerHealth
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *TeamworkPeripheralsHealth) SetCommunicationSpeakerHealth(value TeamworkPeripheralHealthable)() {
    if m != nil {
        m.communicationSpeakerHealth = value
    }
}
// SetContentCameraHealth sets the contentCameraHealth property value. The health details about the content camera.
func (m *TeamworkPeripheralsHealth) SetContentCameraHealth(value TeamworkPeripheralHealthable)() {
    if m != nil {
        m.contentCameraHealth = value
    }
}
// SetDisplayHealthCollection sets the displayHealthCollection property value. The health details about displays.
func (m *TeamworkPeripheralsHealth) SetDisplayHealthCollection(value []TeamworkPeripheralHealthable)() {
    if m != nil {
        m.displayHealthCollection = value
    }
}
// SetMicrophoneHealth sets the microphoneHealth property value. The health details about the microphone.
func (m *TeamworkPeripheralsHealth) SetMicrophoneHealth(value TeamworkPeripheralHealthable)() {
    if m != nil {
        m.microphoneHealth = value
    }
}
// SetRoomCameraHealth sets the roomCameraHealth property value. The health details about the room camera.
func (m *TeamworkPeripheralsHealth) SetRoomCameraHealth(value TeamworkPeripheralHealthable)() {
    if m != nil {
        m.roomCameraHealth = value
    }
}
// SetSpeakerHealth sets the speakerHealth property value. The health details about the speaker.
func (m *TeamworkPeripheralsHealth) SetSpeakerHealth(value TeamworkPeripheralHealthable)() {
    if m != nil {
        m.speakerHealth = value
    }
}
