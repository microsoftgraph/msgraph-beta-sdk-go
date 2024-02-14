package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Mention struct {
    Entity
}
// NewMention instantiates a new Mention and sets the default values.
func NewMention()(*Mention) {
    m := &Mention{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMentionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMentionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMention(), nil
}
// GetApplication gets the application property value. The name of the application where the mention is created. Optional. Not used and defaulted as null for message.
// returns a *string when successful
func (m *Mention) GetApplication()(*string) {
    val, err := m.GetBackingStore().Get("application")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetClientReference gets the clientReference property value. A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for message.
// returns a *string when successful
func (m *Mention) GetClientReference()(*string) {
    val, err := m.GetBackingStore().Get("clientReference")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. The email information of the user who made the mention.
// returns a EmailAddressable when successful
func (m *Mention) GetCreatedBy()(EmailAddressable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EmailAddressable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time that the mention is created on the client.
// returns a *Time when successful
func (m *Mention) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeepLink gets the deepLink property value. A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for message.
// returns a *string when successful
func (m *Mention) GetDeepLink()(*string) {
    val, err := m.GetBackingStore().Get("deepLink")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Mention) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["application"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val)
        }
        return nil
    }
    res["clientReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientReference(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(EmailAddressable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deepLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeepLink(val)
        }
        return nil
    }
    res["mentioned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMentioned(val.(EmailAddressable))
        }
        return nil
    }
    res["mentionText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMentionText(val)
        }
        return nil
    }
    res["serverCreatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerCreatedDateTime(val)
        }
        return nil
    }
    return res
}
// GetMentioned gets the mentioned property value. The mentioned property
// returns a EmailAddressable when successful
func (m *Mention) GetMentioned()(EmailAddressable) {
    val, err := m.GetBackingStore().Get("mentioned")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EmailAddressable)
    }
    return nil
}
// GetMentionText gets the mentionText property value. Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property of the message instead.
// returns a *string when successful
func (m *Mention) GetMentionText()(*string) {
    val, err := m.GetBackingStore().Get("mentionText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServerCreatedDateTime gets the serverCreatedDateTime property value. The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message.
// returns a *Time when successful
func (m *Mention) GetServerCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("serverCreatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Mention) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientReference", m.GetClientReference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deepLink", m.GetDeepLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mentioned", m.GetMentioned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mentionText", m.GetMentionText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("serverCreatedDateTime", m.GetServerCreatedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplication sets the application property value. The name of the application where the mention is created. Optional. Not used and defaulted as null for message.
func (m *Mention) SetApplication(value *string)() {
    err := m.GetBackingStore().Set("application", value)
    if err != nil {
        panic(err)
    }
}
// SetClientReference sets the clientReference property value. A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for message.
func (m *Mention) SetClientReference(value *string)() {
    err := m.GetBackingStore().Set("clientReference", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. The email information of the user who made the mention.
func (m *Mention) SetCreatedBy(value EmailAddressable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time that the mention is created on the client.
func (m *Mention) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeepLink sets the deepLink property value. A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for message.
func (m *Mention) SetDeepLink(value *string)() {
    err := m.GetBackingStore().Set("deepLink", value)
    if err != nil {
        panic(err)
    }
}
// SetMentioned sets the mentioned property value. The mentioned property
func (m *Mention) SetMentioned(value EmailAddressable)() {
    err := m.GetBackingStore().Set("mentioned", value)
    if err != nil {
        panic(err)
    }
}
// SetMentionText sets the mentionText property value. Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property of the message instead.
func (m *Mention) SetMentionText(value *string)() {
    err := m.GetBackingStore().Set("mentionText", value)
    if err != nil {
        panic(err)
    }
}
// SetServerCreatedDateTime sets the serverCreatedDateTime property value. The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message.
func (m *Mention) SetServerCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("serverCreatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
type Mentionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplication()(*string)
    GetClientReference()(*string)
    GetCreatedBy()(EmailAddressable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeepLink()(*string)
    GetMentioned()(EmailAddressable)
    GetMentionText()(*string)
    GetServerCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetApplication(value *string)()
    SetClientReference(value *string)()
    SetCreatedBy(value EmailAddressable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeepLink(value *string)()
    SetMentioned(value EmailAddressable)()
    SetMentionText(value *string)()
    SetServerCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
