package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserAccount 
type UserAccount struct {
    // The accountName property
    accountName *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The azureAdUserId property
    azureAdUserId *string
    // The domainName property
    domainName *string
    // The userPrincipalName property
    userPrincipalName *string
    // The userSid property
    userSid *string
}
// NewUserAccount instantiates a new userAccount and sets the default values.
func NewUserAccount()(*UserAccount) {
    m := &UserAccount{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserAccount(), nil
}
// GetAccountName gets the accountName property value. The accountName property
func (m *UserAccount) GetAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountName
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserAccount) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAzureAdUserId gets the azureAdUserId property value. The azureAdUserId property
func (m *UserAccount) GetAzureAdUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureAdUserId
    }
}
// GetDomainName gets the domainName property value. The domainName property
func (m *UserAccount) GetDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domainName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountName(val)
        }
        return nil
    }
    res["azureAdUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdUserId(val)
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
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["userSid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserSid(val)
        }
        return nil
    }
    return res
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *UserAccount) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetUserSid gets the userSid property value. The userSid property
func (m *UserAccount) GetUserSid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userSid
    }
}
// Serialize serializes information the current object
func (m *UserAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("accountName", m.GetAccountName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("azureAdUserId", m.GetAzureAdUserId())
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
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userSid", m.GetUserSid())
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
// SetAccountName sets the accountName property value. The accountName property
func (m *UserAccount) SetAccountName(value *string)() {
    if m != nil {
        m.accountName = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserAccount) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAzureAdUserId sets the azureAdUserId property value. The azureAdUserId property
func (m *UserAccount) SetAzureAdUserId(value *string)() {
    if m != nil {
        m.azureAdUserId = value
    }
}
// SetDomainName sets the domainName property value. The domainName property
func (m *UserAccount) SetDomainName(value *string)() {
    if m != nil {
        m.domainName = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *UserAccount) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetUserSid sets the userSid property value. The userSid property
func (m *UserAccount) SetUserSid(value *string)() {
    if m != nil {
        m.userSid = value
    }
}
