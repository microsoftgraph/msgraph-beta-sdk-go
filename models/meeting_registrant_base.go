package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingRegistrantBase provides operations to manage the collection of accessReview entities.
type MeetingRegistrantBase struct {
    Entity
    // A unique web URL for the registrant to join the meeting. Read-only.
    joinWebUrl *string
}
// NewMeetingRegistrantBase instantiates a new meetingRegistrantBase and sets the default values.
func NewMeetingRegistrantBase()(*MeetingRegistrantBase) {
    m := &MeetingRegistrantBase{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.meetingRegistrantBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMeetingRegistrantBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingRegistrantBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.externalMeetingRegistrant":
                        return NewExternalMeetingRegistrant(), nil
                    case "#microsoft.graph.meetingRegistrant":
                        return NewMeetingRegistrant(), nil
                }
            }
        }
    }
    return NewMeetingRegistrantBase(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingRegistrantBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["joinWebUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJoinWebUrl)
    return res
}
// GetJoinWebUrl gets the joinWebUrl property value. A unique web URL for the registrant to join the meeting. Read-only.
func (m *MeetingRegistrantBase) GetJoinWebUrl()(*string) {
    return m.joinWebUrl
}
// Serialize serializes information the current object
func (m *MeetingRegistrantBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetJoinWebUrl sets the joinWebUrl property value. A unique web URL for the registrant to join the meeting. Read-only.
func (m *MeetingRegistrantBase) SetJoinWebUrl(value *string)() {
    m.joinWebUrl = value
}
