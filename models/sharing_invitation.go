package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharingInvitation 
type SharingInvitation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The email address provided for the recipient of the sharing invitation. Read-only.
    email *string
    // Provides information about who sent the invitation that created this permission, if that information is available. Read-only.
    invitedBy IdentitySetable
    // The redeemedBy property
    redeemedBy *string
    // If true the recipient of the invitation needs to sign in in order to access the shared item. Read-only.
    signInRequired *bool
}
// NewSharingInvitation instantiates a new sharingInvitation and sets the default values.
func NewSharingInvitation()(*SharingInvitation) {
    m := &SharingInvitation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSharingInvitationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharingInvitationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharingInvitation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharingInvitation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEmail gets the email property value. The email address provided for the recipient of the sharing invitation. Read-only.
func (m *SharingInvitation) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharingInvitation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["invitedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["redeemedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedeemedBy(val)
        }
        return nil
    }
    res["signInRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInRequired(val)
        }
        return nil
    }
    return res
}
// GetInvitedBy gets the invitedBy property value. Provides information about who sent the invitation that created this permission, if that information is available. Read-only.
func (m *SharingInvitation) GetInvitedBy()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.invitedBy
    }
}
// GetRedeemedBy gets the redeemedBy property value. The redeemedBy property
func (m *SharingInvitation) GetRedeemedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.redeemedBy
    }
}
// GetSignInRequired gets the signInRequired property value. If true the recipient of the invitation needs to sign in in order to access the shared item. Read-only.
func (m *SharingInvitation) GetSignInRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.signInRequired
    }
}
// Serialize serializes information the current object
func (m *SharingInvitation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("invitedBy", m.GetInvitedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("redeemedBy", m.GetRedeemedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("signInRequired", m.GetSignInRequired())
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
func (m *SharingInvitation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEmail sets the email property value. The email address provided for the recipient of the sharing invitation. Read-only.
func (m *SharingInvitation) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
// SetInvitedBy sets the invitedBy property value. Provides information about who sent the invitation that created this permission, if that information is available. Read-only.
func (m *SharingInvitation) SetInvitedBy(value IdentitySetable)() {
    if m != nil {
        m.invitedBy = value
    }
}
// SetRedeemedBy sets the redeemedBy property value. The redeemedBy property
func (m *SharingInvitation) SetRedeemedBy(value *string)() {
    if m != nil {
        m.redeemedBy = value
    }
}
// SetSignInRequired sets the signInRequired property value. If true the recipient of the invitation needs to sign in in order to access the shared item. Read-only.
func (m *SharingInvitation) SetSignInRequired(value *bool)() {
    if m != nil {
        m.signInRequired = value
    }
}
