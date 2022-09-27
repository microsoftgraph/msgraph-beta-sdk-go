package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessGuestsOrExternalUsers 
type ConditionalAccessGuestsOrExternalUsers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The externalTenants property
    externalTenants ConditionalAccessExternalTenantsable
    // The guestOrExternalUserTypes property
    guestOrExternalUserTypes *ConditionalAccessGuestOrExternalUserTypes
    // The OdataType property
    odataType *string
}
// NewConditionalAccessGuestsOrExternalUsers instantiates a new conditionalAccessGuestsOrExternalUsers and sets the default values.
func NewConditionalAccessGuestsOrExternalUsers()(*ConditionalAccessGuestsOrExternalUsers) {
    m := &ConditionalAccessGuestsOrExternalUsers{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.conditionalAccessGuestsOrExternalUsers";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateConditionalAccessGuestsOrExternalUsersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessGuestsOrExternalUsersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessGuestsOrExternalUsers(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessGuestsOrExternalUsers) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExternalTenants gets the externalTenants property value. The externalTenants property
func (m *ConditionalAccessGuestsOrExternalUsers) GetExternalTenants()(ConditionalAccessExternalTenantsable) {
    return m.externalTenants
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessGuestsOrExternalUsers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["externalTenants"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateConditionalAccessExternalTenantsFromDiscriminatorValue , m.SetExternalTenants)
    res["guestOrExternalUserTypes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseConditionalAccessGuestOrExternalUserTypes , m.SetGuestOrExternalUserTypes)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetGuestOrExternalUserTypes gets the guestOrExternalUserTypes property value. The guestOrExternalUserTypes property
func (m *ConditionalAccessGuestsOrExternalUsers) GetGuestOrExternalUserTypes()(*ConditionalAccessGuestOrExternalUserTypes) {
    return m.guestOrExternalUserTypes
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConditionalAccessGuestsOrExternalUsers) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ConditionalAccessGuestsOrExternalUsers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("externalTenants", m.GetExternalTenants())
        if err != nil {
            return err
        }
    }
    if m.GetGuestOrExternalUserTypes() != nil {
        cast := (*m.GetGuestOrExternalUserTypes()).String()
        err := writer.WriteStringValue("guestOrExternalUserTypes", &cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessGuestsOrExternalUsers) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExternalTenants sets the externalTenants property value. The externalTenants property
func (m *ConditionalAccessGuestsOrExternalUsers) SetExternalTenants(value ConditionalAccessExternalTenantsable)() {
    m.externalTenants = value
}
// SetGuestOrExternalUserTypes sets the guestOrExternalUserTypes property value. The guestOrExternalUserTypes property
func (m *ConditionalAccessGuestsOrExternalUsers) SetGuestOrExternalUserTypes(value *ConditionalAccessGuestOrExternalUserTypes)() {
    m.guestOrExternalUserTypes = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessGuestsOrExternalUsers) SetOdataType(value *string)() {
    m.odataType = value
}
