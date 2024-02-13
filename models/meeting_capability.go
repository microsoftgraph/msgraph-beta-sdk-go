package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type MeetingCapability struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMeetingCapability instantiates a new MeetingCapability and sets the default values.
func NewMeetingCapability()(*MeetingCapability) {
    m := &MeetingCapability{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMeetingCapabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMeetingCapabilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingCapability(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MeetingCapability) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAllowAnonymousUsersToDialOut gets the allowAnonymousUsersToDialOut property value. Indicates whether anonymous users dialout is allowed in a meeting.
// returns a *bool when successful
func (m *MeetingCapability) GetAllowAnonymousUsersToDialOut()(*bool) {
    val, err := m.GetBackingStore().Get("allowAnonymousUsersToDialOut")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowAnonymousUsersToStartMeeting gets the allowAnonymousUsersToStartMeeting property value. Indicates whether anonymous users are allowed to start a meeting.
// returns a *bool when successful
func (m *MeetingCapability) GetAllowAnonymousUsersToStartMeeting()(*bool) {
    val, err := m.GetBackingStore().Get("allowAnonymousUsersToStartMeeting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAutoAdmittedUsers gets the autoAdmittedUsers property value. The autoAdmittedUsers property
// returns a *AutoAdmittedUsersType when successful
func (m *MeetingCapability) GetAutoAdmittedUsers()(*AutoAdmittedUsersType) {
    val, err := m.GetBackingStore().Get("autoAdmittedUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AutoAdmittedUsersType)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *MeetingCapability) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MeetingCapability) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowAnonymousUsersToDialOut"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAnonymousUsersToDialOut(val)
        }
        return nil
    }
    res["allowAnonymousUsersToStartMeeting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAnonymousUsersToStartMeeting(val)
        }
        return nil
    }
    res["autoAdmittedUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAutoAdmittedUsersType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoAdmittedUsers(val.(*AutoAdmittedUsersType))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *MeetingCapability) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MeetingCapability) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowAnonymousUsersToDialOut", m.GetAllowAnonymousUsersToDialOut())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowAnonymousUsersToStartMeeting", m.GetAllowAnonymousUsersToStartMeeting())
        if err != nil {
            return err
        }
    }
    if m.GetAutoAdmittedUsers() != nil {
        cast := (*m.GetAutoAdmittedUsers()).String()
        err := writer.WriteStringValue("autoAdmittedUsers", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingCapability) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowAnonymousUsersToDialOut sets the allowAnonymousUsersToDialOut property value. Indicates whether anonymous users dialout is allowed in a meeting.
func (m *MeetingCapability) SetAllowAnonymousUsersToDialOut(value *bool)() {
    err := m.GetBackingStore().Set("allowAnonymousUsersToDialOut", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowAnonymousUsersToStartMeeting sets the allowAnonymousUsersToStartMeeting property value. Indicates whether anonymous users are allowed to start a meeting.
func (m *MeetingCapability) SetAllowAnonymousUsersToStartMeeting(value *bool)() {
    err := m.GetBackingStore().Set("allowAnonymousUsersToStartMeeting", value)
    if err != nil {
        panic(err)
    }
}
// SetAutoAdmittedUsers sets the autoAdmittedUsers property value. The autoAdmittedUsers property
func (m *MeetingCapability) SetAutoAdmittedUsers(value *AutoAdmittedUsersType)() {
    err := m.GetBackingStore().Set("autoAdmittedUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MeetingCapability) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MeetingCapability) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type MeetingCapabilityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAnonymousUsersToDialOut()(*bool)
    GetAllowAnonymousUsersToStartMeeting()(*bool)
    GetAutoAdmittedUsers()(*AutoAdmittedUsersType)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    SetAllowAnonymousUsersToDialOut(value *bool)()
    SetAllowAnonymousUsersToStartMeeting(value *bool)()
    SetAutoAdmittedUsers(value *AutoAdmittedUsersType)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
}
