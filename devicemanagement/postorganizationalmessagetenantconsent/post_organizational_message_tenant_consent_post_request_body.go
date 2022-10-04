package postorganizationalmessagetenantconsent

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PostOrganizationalMessageTenantConsentPostRequestBody provides operations to call the postOrganizationalMessageTenantConsent method.
type PostOrganizationalMessageTenantConsentPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The firstPartyMessageAllowed property
    firstPartyMessageAllowed *bool
}
// NewPostOrganizationalMessageTenantConsentPostRequestBody instantiates a new postOrganizationalMessageTenantConsentPostRequestBody and sets the default values.
func NewPostOrganizationalMessageTenantConsentPostRequestBody()(*PostOrganizationalMessageTenantConsentPostRequestBody) {
    m := &PostOrganizationalMessageTenantConsentPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePostOrganizationalMessageTenantConsentPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePostOrganizationalMessageTenantConsentPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPostOrganizationalMessageTenantConsentPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PostOrganizationalMessageTenantConsentPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PostOrganizationalMessageTenantConsentPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["firstPartyMessageAllowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetFirstPartyMessageAllowed)
    return res
}
// GetFirstPartyMessageAllowed gets the firstPartyMessageAllowed property value. The firstPartyMessageAllowed property
func (m *PostOrganizationalMessageTenantConsentPostRequestBody) GetFirstPartyMessageAllowed()(*bool) {
    return m.firstPartyMessageAllowed
}
// Serialize serializes information the current object
func (m *PostOrganizationalMessageTenantConsentPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("firstPartyMessageAllowed", m.GetFirstPartyMessageAllowed())
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
func (m *PostOrganizationalMessageTenantConsentPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFirstPartyMessageAllowed sets the firstPartyMessageAllowed property value. The firstPartyMessageAllowed property
func (m *PostOrganizationalMessageTenantConsentPostRequestBody) SetFirstPartyMessageAllowed(value *bool)() {
    m.firstPartyMessageAllowed = value
}
