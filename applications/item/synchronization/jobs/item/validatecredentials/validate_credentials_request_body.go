package validatecredentials

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ValidateCredentialsRequestBody provides operations to call the validateCredentials method.
type ValidateCredentialsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The applicationIdentifier property
    applicationIdentifier *string
    // The credentials property
    credentials []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSecretKeyStringValuePairable
    // The templateId property
    templateId *string
    // The useSavedCredentials property
    useSavedCredentials *bool
}
// NewValidateCredentialsRequestBody instantiates a new validateCredentialsRequestBody and sets the default values.
func NewValidateCredentialsRequestBody()(*ValidateCredentialsRequestBody) {
    m := &ValidateCredentialsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateValidateCredentialsRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateValidateCredentialsRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewValidateCredentialsRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidateCredentialsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicationIdentifier gets the applicationIdentifier property value. The applicationIdentifier property
func (m *ValidateCredentialsRequestBody) GetApplicationIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationIdentifier
    }
}
// GetCredentials gets the credentials property value. The credentials property
func (m *ValidateCredentialsRequestBody) GetCredentials()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSecretKeyStringValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.credentials
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidateCredentialsRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationIdentifier(val)
        }
        return nil
    }
    res["credentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSynchronizationSecretKeyStringValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSecretKeyStringValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSecretKeyStringValuePairable)
            }
            m.SetCredentials(res)
        }
        return nil
    }
    res["templateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    res["useSavedCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseSavedCredentials(val)
        }
        return nil
    }
    return res
}
// GetTemplateId gets the templateId property value. The templateId property
func (m *ValidateCredentialsRequestBody) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// GetUseSavedCredentials gets the useSavedCredentials property value. The useSavedCredentials property
func (m *ValidateCredentialsRequestBody) GetUseSavedCredentials()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.useSavedCredentials
    }
}
// Serialize serializes information the current object
func (m *ValidateCredentialsRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationIdentifier", m.GetApplicationIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCredentials()))
        for i, v := range m.GetCredentials() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("credentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("useSavedCredentials", m.GetUseSavedCredentials())
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
func (m *ValidateCredentialsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicationIdentifier sets the applicationIdentifier property value. The applicationIdentifier property
func (m *ValidateCredentialsRequestBody) SetApplicationIdentifier(value *string)() {
    if m != nil {
        m.applicationIdentifier = value
    }
}
// SetCredentials sets the credentials property value. The credentials property
func (m *ValidateCredentialsRequestBody) SetCredentials(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationSecretKeyStringValuePairable)() {
    if m != nil {
        m.credentials = value
    }
}
// SetTemplateId sets the templateId property value. The templateId property
func (m *ValidateCredentialsRequestBody) SetTemplateId(value *string)() {
    if m != nil {
        m.templateId = value
    }
}
// SetUseSavedCredentials sets the useSavedCredentials property value. The useSavedCredentials property
func (m *ValidateCredentialsRequestBody) SetUseSavedCredentials(value *bool)() {
    if m != nil {
        m.useSavedCredentials = value
    }
}
