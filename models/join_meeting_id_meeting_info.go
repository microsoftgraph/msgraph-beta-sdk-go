package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JoinMeetingIdMeetingInfo 
type JoinMeetingIdMeetingInfo struct {
    MeetingInfo
    // The ID used to join the meeting.
    joinMeetingId *string
    // The passcode used to join the meeting. Optional.
    passcode *string
}
// NewJoinMeetingIdMeetingInfo instantiates a new JoinMeetingIdMeetingInfo and sets the default values.
func NewJoinMeetingIdMeetingInfo()(*JoinMeetingIdMeetingInfo) {
    m := &JoinMeetingIdMeetingInfo{
        MeetingInfo: *NewMeetingInfo(),
    }
    odataTypeValue := "#microsoft.graph.joinMeetingIdMeetingInfo";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateJoinMeetingIdMeetingInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJoinMeetingIdMeetingInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJoinMeetingIdMeetingInfo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JoinMeetingIdMeetingInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MeetingInfo.GetFieldDeserializers()
    res["joinMeetingId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJoinMeetingId)
    res["passcode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPasscode)
    return res
}
// GetJoinMeetingId gets the joinMeetingId property value. The ID used to join the meeting.
func (m *JoinMeetingIdMeetingInfo) GetJoinMeetingId()(*string) {
    return m.joinMeetingId
}
// GetPasscode gets the passcode property value. The passcode used to join the meeting. Optional.
func (m *JoinMeetingIdMeetingInfo) GetPasscode()(*string) {
    return m.passcode
}
// Serialize serializes information the current object
func (m *JoinMeetingIdMeetingInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MeetingInfo.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("joinMeetingId", m.GetJoinMeetingId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("passcode", m.GetPasscode())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetJoinMeetingId sets the joinMeetingId property value. The ID used to join the meeting.
func (m *JoinMeetingIdMeetingInfo) SetJoinMeetingId(value *string)() {
    m.joinMeetingId = value
}
// SetPasscode sets the passcode property value. The passcode used to join the meeting. Optional.
func (m *JoinMeetingIdMeetingInfo) SetPasscode(value *string)() {
    m.passcode = value
}
