package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkActivePeripherals 
type TeamworkActivePeripherals struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    communicationSpeaker *TeamworkPeripheral;
    // 
    contentCamera *TeamworkPeripheral;
    // 
    microphone *TeamworkPeripheral;
    // 
    roomCamera *TeamworkPeripheral;
    // 
    speaker *TeamworkPeripheral;
}
// NewTeamworkActivePeripherals instantiates a new teamworkActivePeripherals and sets the default values.
func NewTeamworkActivePeripherals()(*TeamworkActivePeripherals) {
    m := &TeamworkActivePeripherals{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkActivePeripherals) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCommunicationSpeaker gets the communicationSpeaker property value. 
func (m *TeamworkActivePeripherals) GetCommunicationSpeaker()(*TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.communicationSpeaker
    }
}
// GetContentCamera gets the contentCamera property value. 
func (m *TeamworkActivePeripherals) GetContentCamera()(*TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.contentCamera
    }
}
// GetMicrophone gets the microphone property value. 
func (m *TeamworkActivePeripherals) GetMicrophone()(*TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.microphone
    }
}
// GetRoomCamera gets the roomCamera property value. 
func (m *TeamworkActivePeripherals) GetRoomCamera()(*TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.roomCamera
    }
}
// GetSpeaker gets the speaker property value. 
func (m *TeamworkActivePeripherals) GetSpeaker()(*TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.speaker
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkActivePeripherals) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["communicationSpeaker"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommunicationSpeaker(val.(*TeamworkPeripheral))
        }
        return nil
    }
    res["contentCamera"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCamera(val.(*TeamworkPeripheral))
        }
        return nil
    }
    res["microphone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophone(val.(*TeamworkPeripheral))
        }
        return nil
    }
    res["roomCamera"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoomCamera(val.(*TeamworkPeripheral))
        }
        return nil
    }
    res["speaker"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeaker(val.(*TeamworkPeripheral))
        }
        return nil
    }
    return res
}
func (m *TeamworkActivePeripherals) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkActivePeripherals) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("communicationSpeaker", m.GetCommunicationSpeaker())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("contentCamera", m.GetContentCamera())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("microphone", m.GetMicrophone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("roomCamera", m.GetRoomCamera())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("speaker", m.GetSpeaker())
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
func (m *TeamworkActivePeripherals) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCommunicationSpeaker sets the communicationSpeaker property value. 
func (m *TeamworkActivePeripherals) SetCommunicationSpeaker(value *TeamworkPeripheral)() {
    if m != nil {
        m.communicationSpeaker = value
    }
}
// SetContentCamera sets the contentCamera property value. 
func (m *TeamworkActivePeripherals) SetContentCamera(value *TeamworkPeripheral)() {
    if m != nil {
        m.contentCamera = value
    }
}
// SetMicrophone sets the microphone property value. 
func (m *TeamworkActivePeripherals) SetMicrophone(value *TeamworkPeripheral)() {
    if m != nil {
        m.microphone = value
    }
}
// SetRoomCamera sets the roomCamera property value. 
func (m *TeamworkActivePeripherals) SetRoomCamera(value *TeamworkPeripheral)() {
    if m != nil {
        m.roomCamera = value
    }
}
// SetSpeaker sets the speaker property value. 
func (m *TeamworkActivePeripherals) SetSpeaker(value *TeamworkPeripheral)() {
    if m != nil {
        m.speaker = value
    }
}
