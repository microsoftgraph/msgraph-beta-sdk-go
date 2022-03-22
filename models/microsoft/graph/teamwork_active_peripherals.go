package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkActivePeripherals 
type TeamworkActivePeripherals struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    communicationSpeaker TeamworkPeripheralable;
    // 
    contentCamera TeamworkPeripheralable;
    // 
    microphone TeamworkPeripheralable;
    // 
    roomCamera TeamworkPeripheralable;
    // 
    speaker TeamworkPeripheralable;
}
// NewTeamworkActivePeripherals instantiates a new teamworkActivePeripherals and sets the default values.
func NewTeamworkActivePeripherals()(*TeamworkActivePeripherals) {
    m := &TeamworkActivePeripherals{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkActivePeripheralsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkActivePeripheralsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamworkActivePeripherals(), nil
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
func (m *TeamworkActivePeripherals) GetCommunicationSpeaker()(TeamworkPeripheralable) {
    if m == nil {
        return nil
    } else {
        return m.communicationSpeaker
    }
}
// GetContentCamera gets the contentCamera property value. 
func (m *TeamworkActivePeripherals) GetContentCamera()(TeamworkPeripheralable) {
    if m == nil {
        return nil
    } else {
        return m.contentCamera
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkActivePeripherals) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["communicationSpeaker"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommunicationSpeaker(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["contentCamera"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCamera(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["microphone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophone(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["roomCamera"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoomCamera(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["speaker"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeaker(val.(TeamworkPeripheralable))
        }
        return nil
    }
    return res
}
// GetMicrophone gets the microphone property value. 
func (m *TeamworkActivePeripherals) GetMicrophone()(TeamworkPeripheralable) {
    if m == nil {
        return nil
    } else {
        return m.microphone
    }
}
// GetRoomCamera gets the roomCamera property value. 
func (m *TeamworkActivePeripherals) GetRoomCamera()(TeamworkPeripheralable) {
    if m == nil {
        return nil
    } else {
        return m.roomCamera
    }
}
// GetSpeaker gets the speaker property value. 
func (m *TeamworkActivePeripherals) GetSpeaker()(TeamworkPeripheralable) {
    if m == nil {
        return nil
    } else {
        return m.speaker
    }
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
func (m *TeamworkActivePeripherals) SetCommunicationSpeaker(value TeamworkPeripheralable)() {
    if m != nil {
        m.communicationSpeaker = value
    }
}
// SetContentCamera sets the contentCamera property value. 
func (m *TeamworkActivePeripherals) SetContentCamera(value TeamworkPeripheralable)() {
    if m != nil {
        m.contentCamera = value
    }
}
// SetMicrophone sets the microphone property value. 
func (m *TeamworkActivePeripherals) SetMicrophone(value TeamworkPeripheralable)() {
    if m != nil {
        m.microphone = value
    }
}
// SetRoomCamera sets the roomCamera property value. 
func (m *TeamworkActivePeripherals) SetRoomCamera(value TeamworkPeripheralable)() {
    if m != nil {
        m.roomCamera = value
    }
}
// SetSpeaker sets the speaker property value. 
func (m *TeamworkActivePeripherals) SetSpeaker(value TeamworkPeripheralable)() {
    if m != nil {
        m.speaker = value
    }
}
