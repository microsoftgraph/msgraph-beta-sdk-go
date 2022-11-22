package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LoggedOnUser 
type LoggedOnUser struct {
    // User account name of the logged-on user.
    accountName *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // User account domain of the logged-on user.
    domainName *string
    // The OdataType property
    odataType *string
}
// NewLoggedOnUser instantiates a new loggedOnUser and sets the default values.
func NewLoggedOnUser()(*LoggedOnUser) {
    m := &LoggedOnUser{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLoggedOnUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLoggedOnUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLoggedOnUser(), nil
}
// GetAccountName gets the accountName property value. User account name of the logged-on user.
func (m *LoggedOnUser) GetAccountName()(*string) {
    return m.accountName
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LoggedOnUser) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDomainName gets the domainName property value. User account domain of the logged-on user.
func (m *LoggedOnUser) GetDomainName()(*string) {
    return m.domainName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LoggedOnUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAccountName)
    res["domainName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDomainName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LoggedOnUser) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *LoggedOnUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("accountName", m.GetAccountName())
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
// SetAccountName sets the accountName property value. User account name of the logged-on user.
func (m *LoggedOnUser) SetAccountName(value *string)() {
    m.accountName = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LoggedOnUser) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDomainName sets the domainName property value. User account domain of the logged-on user.
func (m *LoggedOnUser) SetDomainName(value *string)() {
    m.domainName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LoggedOnUser) SetOdataType(value *string)() {
    m.odataType = value
}
