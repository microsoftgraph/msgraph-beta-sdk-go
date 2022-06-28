package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharePointIdentity 
type SharePointIdentity struct {
    Identity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The sign in name of the SharePoint identity.
    loginName *string
}
// NewSharePointIdentity instantiates a new SharePointIdentity and sets the default values.
func NewSharePointIdentity()(*SharePointIdentity) {
    m := &SharePointIdentity{
        Identity: *NewIdentity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSharePointIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharePointIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharePointIdentity(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharePointIdentity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharePointIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["loginName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginName(val)
        }
        return nil
    }
    return res
}
// GetLoginName gets the loginName property value. The sign in name of the SharePoint identity.
func (m *SharePointIdentity) GetLoginName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.loginName
    }
}
// Serialize serializes information the current object
func (m *SharePointIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("loginName", m.GetLoginName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharePointIdentity) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLoginName sets the loginName property value. The sign in name of the SharePoint identity.
func (m *SharePointIdentity) SetLoginName(value *string)() {
    if m != nil {
        m.loginName = value
    }
}
