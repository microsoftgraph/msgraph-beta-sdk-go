package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserSubject struct {
    ConditionalAccessWhatIfSubject
}
// NewUserSubject instantiates a new UserSubject and sets the default values.
func NewUserSubject()(*UserSubject) {
    m := &UserSubject{
        ConditionalAccessWhatIfSubject: *NewConditionalAccessWhatIfSubject(),
    }
    odataTypeValue := "#microsoft.graph.userSubject"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUserSubjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserSubjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSubject(), nil
}
// GetExternalTenantId gets the externalTenantId property value. The externalTenantId property
// returns a *string when successful
func (m *UserSubject) GetExternalTenantId()(*string) {
    val, err := m.GetBackingStore().Get("externalTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalUserType gets the externalUserType property value. The externalUserType property
// returns a *ConditionalAccessGuestOrExternalUserTypes when successful
func (m *UserSubject) GetExternalUserType()(*ConditionalAccessGuestOrExternalUserTypes) {
    val, err := m.GetBackingStore().Get("externalUserType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConditionalAccessGuestOrExternalUserTypes)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserSubject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ConditionalAccessWhatIfSubject.GetFieldDeserializers()
    res["externalTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalTenantId(val)
        }
        return nil
    }
    res["externalUserType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessGuestOrExternalUserTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUserType(val.(*ConditionalAccessGuestOrExternalUserTypes))
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetUserId gets the userId property value. The userId property
// returns a *string when successful
func (m *UserSubject) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserSubject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ConditionalAccessWhatIfSubject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("externalTenantId", m.GetExternalTenantId())
        if err != nil {
            return err
        }
    }
    if m.GetExternalUserType() != nil {
        cast := (*m.GetExternalUserType()).String()
        err = writer.WriteStringValue("externalUserType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExternalTenantId sets the externalTenantId property value. The externalTenantId property
func (m *UserSubject) SetExternalTenantId(value *string)() {
    err := m.GetBackingStore().Set("externalTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalUserType sets the externalUserType property value. The externalUserType property
func (m *UserSubject) SetExternalUserType(value *ConditionalAccessGuestOrExternalUserTypes)() {
    err := m.GetBackingStore().Set("externalUserType", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The userId property
func (m *UserSubject) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
type UserSubjectable interface {
    ConditionalAccessWhatIfSubjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExternalTenantId()(*string)
    GetExternalUserType()(*ConditionalAccessGuestOrExternalUserTypes)
    GetUserId()(*string)
    SetExternalTenantId(value *string)()
    SetExternalUserType(value *ConditionalAccessGuestOrExternalUserTypes)()
    SetUserId(value *string)()
}
