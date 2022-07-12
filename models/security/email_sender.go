package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmailSender 
type EmailSender struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The displayName property
    displayName *string
    // The domainName property
    domainName *string
    // The emailAddress property
    emailAddress *string
}
// NewEmailSender instantiates a new emailSender and sets the default values.
func NewEmailSender()(*EmailSender) {
    m := &EmailSender{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEmailSenderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailSenderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmailSender(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmailSender) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *EmailSender) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDomainName gets the domainName property value. The domainName property
func (m *EmailSender) GetDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domainName
    }
}
// GetEmailAddress gets the emailAddress property value. The emailAddress property
func (m *EmailSender) GetEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailSender) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EmailSender) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("emailAddress", m.GetEmailAddress())
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
func (m *EmailSender) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *EmailSender) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDomainName sets the domainName property value. The domainName property
func (m *EmailSender) SetDomainName(value *string)() {
    if m != nil {
        m.domainName = value
    }
}
// SetEmailAddress sets the emailAddress property value. The emailAddress property
func (m *EmailSender) SetEmailAddress(value *string)() {
    if m != nil {
        m.emailAddress = value
    }
}
