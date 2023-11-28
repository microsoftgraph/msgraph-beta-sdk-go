package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AwsPolicy 
type AwsPolicy struct {
    Entity
}
// NewAwsPolicy instantiates a new awsPolicy and sets the default values.
func NewAwsPolicy()(*AwsPolicy) {
    m := &AwsPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAwsPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAwsPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsPolicy(), nil
}
// GetAwsPolicyType gets the awsPolicyType property value. The awsPolicyType property
func (m *AwsPolicy) GetAwsPolicyType()(*AwsPolicyType) {
    val, err := m.GetBackingStore().Get("awsPolicyType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AwsPolicyType)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the AWS policy. Read-only. Supports $filter and (eq,contains).
func (m *AwsPolicy) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalId gets the externalId property value. The base64 encoded identifier for the AWS policy as defined by AWS. Read-only. Alternate key. Supports $filter and eq.
func (m *AwsPolicy) GetExternalId()(*string) {
    val, err := m.GetBackingStore().Get("externalId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AwsPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["awsPolicyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAwsPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAwsPolicyType(val.(*AwsPolicyType))
        }
        return nil
    }
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
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AwsPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAwsPolicyType() != nil {
        cast := (*m.GetAwsPolicyType()).String()
        err = writer.WriteStringValue("awsPolicyType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAwsPolicyType sets the awsPolicyType property value. The awsPolicyType property
func (m *AwsPolicy) SetAwsPolicyType(value *AwsPolicyType)() {
    err := m.GetBackingStore().Set("awsPolicyType", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the AWS policy. Read-only. Supports $filter and (eq,contains).
func (m *AwsPolicy) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalId sets the externalId property value. The base64 encoded identifier for the AWS policy as defined by AWS. Read-only. Alternate key. Supports $filter and eq.
func (m *AwsPolicy) SetExternalId(value *string)() {
    err := m.GetBackingStore().Set("externalId", value)
    if err != nil {
        panic(err)
    }
}
// AwsPolicyable 
type AwsPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAwsPolicyType()(*AwsPolicyType)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    SetAwsPolicyType(value *AwsPolicyType)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
}
