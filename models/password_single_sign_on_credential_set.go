package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PasswordSingleSignOnCredentialSet 
type PasswordSingleSignOnCredentialSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A list of credential objects that define the complete sign in flow.
    credentials []Credentialable
    // The ID of the user or group this credential set belongs to.
    id *string
    // The OdataType property
    odataType *string
}
// NewPasswordSingleSignOnCredentialSet instantiates a new passwordSingleSignOnCredentialSet and sets the default values.
func NewPasswordSingleSignOnCredentialSet()(*PasswordSingleSignOnCredentialSet) {
    m := &PasswordSingleSignOnCredentialSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.passwordSingleSignOnCredentialSet";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePasswordSingleSignOnCredentialSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePasswordSingleSignOnCredentialSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPasswordSingleSignOnCredentialSet(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordSingleSignOnCredentialSet) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCredentials gets the credentials property value. A list of credential objects that define the complete sign in flow.
func (m *PasswordSingleSignOnCredentialSet) GetCredentials()([]Credentialable) {
    return m.credentials
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordSingleSignOnCredentialSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["credentials"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCredentialFromDiscriminatorValue , m.SetCredentials)
    res["id"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetId gets the id property value. The ID of the user or group this credential set belongs to.
func (m *PasswordSingleSignOnCredentialSet) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PasswordSingleSignOnCredentialSet) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *PasswordSingleSignOnCredentialSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCredentials() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCredentials())
        err := writer.WriteCollectionOfObjectValues("credentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *PasswordSingleSignOnCredentialSet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCredentials sets the credentials property value. A list of credential objects that define the complete sign in flow.
func (m *PasswordSingleSignOnCredentialSet) SetCredentials(value []Credentialable)() {
    m.credentials = value
}
// SetId sets the id property value. The ID of the user or group this credential set belongs to.
func (m *PasswordSingleSignOnCredentialSet) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PasswordSingleSignOnCredentialSet) SetOdataType(value *string)() {
    m.odataType = value
}
