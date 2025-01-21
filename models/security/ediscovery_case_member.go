package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type EdiscoveryCaseMember struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewEdiscoveryCaseMember instantiates a new EdiscoveryCaseMember and sets the default values.
func NewEdiscoveryCaseMember()(*EdiscoveryCaseMember) {
    m := &EdiscoveryCaseMember{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateEdiscoveryCaseMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEdiscoveryCaseMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryCaseMember(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *EdiscoveryCaseMember) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
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
func (m *EdiscoveryCaseMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["recipientType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecipientType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientType(val.(*RecipientType))
        }
        return nil
    }
    res["smtpAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmtpAddress(val)
        }
        return nil
    }
    return res
}
// GetRecipientType gets the recipientType property value. The recipientType property
// returns a *RecipientType when successful
func (m *EdiscoveryCaseMember) GetRecipientType()(*RecipientType) {
    val, err := m.GetBackingStore().Get("recipientType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RecipientType)
    }
    return nil
}
// GetSmtpAddress gets the smtpAddress property value. The smtpAddress property
// returns a *string when successful
func (m *EdiscoveryCaseMember) GetSmtpAddress()(*string) {
    val, err := m.GetBackingStore().Get("smtpAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EdiscoveryCaseMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetRecipientType() != nil {
        cast := (*m.GetRecipientType()).String()
        err = writer.WriteStringValue("recipientType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("smtpAddress", m.GetSmtpAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *EdiscoveryCaseMember) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetRecipientType sets the recipientType property value. The recipientType property
func (m *EdiscoveryCaseMember) SetRecipientType(value *RecipientType)() {
    err := m.GetBackingStore().Set("recipientType", value)
    if err != nil {
        panic(err)
    }
}
// SetSmtpAddress sets the smtpAddress property value. The smtpAddress property
func (m *EdiscoveryCaseMember) SetSmtpAddress(value *string)() {
    err := m.GetBackingStore().Set("smtpAddress", value)
    if err != nil {
        panic(err)
    }
}
type EdiscoveryCaseMemberable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetRecipientType()(*RecipientType)
    GetSmtpAddress()(*string)
    SetDisplayName(value *string)()
    SetRecipientType(value *RecipientType)()
    SetSmtpAddress(value *string)()
}
