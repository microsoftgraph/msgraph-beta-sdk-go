package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody 
type AndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The tokenValidityInSeconds property
    tokenValidityInSeconds *int32
}
// NewAndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody instantiates a new AndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody and sets the default values.
func NewAndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody()(*AndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody) {
    m := &AndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["tokenValidityInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenValidityInSeconds(val)
        }
        return nil
    }
    return res
}
// GetTokenValidityInSeconds gets the tokenValidityInSeconds property value. The tokenValidityInSeconds property
func (m *AndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody) GetTokenValidityInSeconds()(*int32) {
    return m.tokenValidityInSeconds
}
// Serialize serializes information the current object
func (m *AndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("tokenValidityInSeconds", m.GetTokenValidityInSeconds())
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
func (m *AndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTokenValidityInSeconds sets the tokenValidityInSeconds property value. The tokenValidityInSeconds property
func (m *AndroidForWorkEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBody) SetTokenValidityInSeconds(value *int32)() {
    m.tokenValidityInSeconds = value
}
