package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkSpeakerConfiguration 
type TeamworkSpeakerConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The defaultCommunicationSpeaker property
    defaultCommunicationSpeaker TeamworkPeripheralable
    // The defaultSpeaker property
    defaultSpeaker TeamworkPeripheralable
    // True if the communication speaker is optional. Used to compute the health state if the communication speaker is not optional.
    isCommunicationSpeakerOptional *bool
    // True if the configured speaker is optional. Used to compute the health state if the speaker is not optional.
    isSpeakerOptional *bool
    // The speakers property
    speakers []TeamworkPeripheralable
}
// NewTeamworkSpeakerConfiguration instantiates a new teamworkSpeakerConfiguration and sets the default values.
func NewTeamworkSpeakerConfiguration()(*TeamworkSpeakerConfiguration) {
    m := &TeamworkSpeakerConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkSpeakerConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkSpeakerConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkSpeakerConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkSpeakerConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultCommunicationSpeaker gets the defaultCommunicationSpeaker property value. The defaultCommunicationSpeaker property
func (m *TeamworkSpeakerConfiguration) GetDefaultCommunicationSpeaker()(TeamworkPeripheralable) {
    if m == nil {
        return nil
    } else {
        return m.defaultCommunicationSpeaker
    }
}
// GetDefaultSpeaker gets the defaultSpeaker property value. The defaultSpeaker property
func (m *TeamworkSpeakerConfiguration) GetDefaultSpeaker()(TeamworkPeripheralable) {
    if m == nil {
        return nil
    } else {
        return m.defaultSpeaker
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkSpeakerConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultCommunicationSpeaker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultCommunicationSpeaker(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["defaultSpeaker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultSpeaker(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["isCommunicationSpeakerOptional"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCommunicationSpeakerOptional(val)
        }
        return nil
    }
    res["isSpeakerOptional"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSpeakerOptional(val)
        }
        return nil
    }
    res["speakers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkPeripheralable, len(val))
            for i, v := range val {
                res[i] = v.(TeamworkPeripheralable)
            }
            m.SetSpeakers(res)
        }
        return nil
    }
    return res
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
// GetSpeakers gets the speakers property value. The speakers property
func (m *TeamworkSpeakerConfiguration) GetSpeakers()([]TeamworkPeripheralable) {
    if m == nil {
        return nil
    } else {
        return m.speakers
    }
}
// Serialize serializes information the current object
func (m *TeamworkSpeakerConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSpeakers()))
        for i, v := range m.GetSpeakers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
// SetDefaultCommunicationSpeaker sets the defaultCommunicationSpeaker property value. The defaultCommunicationSpeaker property
func (m *TeamworkSpeakerConfiguration) SetDefaultCommunicationSpeaker(value TeamworkPeripheralable)() {
    if m != nil {
        m.defaultCommunicationSpeaker = value
    }
}
// SetDefaultSpeaker sets the defaultSpeaker property value. The defaultSpeaker property
func (m *TeamworkSpeakerConfiguration) SetDefaultSpeaker(value TeamworkPeripheralable)() {
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
// SetSpeakers sets the speakers property value. The speakers property
func (m *TeamworkSpeakerConfiguration) SetSpeakers(value []TeamworkPeripheralable)() {
    if m != nil {
        m.speakers = value
    }
}
