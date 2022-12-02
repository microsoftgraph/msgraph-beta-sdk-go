package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody provides operations to call the validateCredentials method.
type ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The applicationIdentifier property
    applicationIdentifier *string
    // The credentials property
    credentials []SynchronizationSecretKeyStringValuePairable
    // The templateId property
    templateId *string
    // The useSavedCredentials property
    useSavedCredentials *bool
}
// NewApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody instantiates a new ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody and sets the default values.
func NewApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody()(*ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody) {
    m := &ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApplicationIdentifier gets the applicationIdentifier property value. The applicationIdentifier property
func (m *ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody) GetApplicationIdentifier()(*string) {
    return m.applicationIdentifier
}
// GetCredentials gets the credentials property value. The credentials property
func (m *ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody) GetCredentials()([]SynchronizationSecretKeyStringValuePairable) {
    return m.credentials
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationSecretKeyStringValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationSecretKeyStringValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(SynchronizationSecretKeyStringValuePairable)
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
func (m *ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody) GetTemplateId()(*string) {
    return m.templateId
}
// GetUseSavedCredentials gets the useSavedCredentials property value. The useSavedCredentials property
func (m *ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody) GetUseSavedCredentials()(*bool) {
    return m.useSavedCredentials
}
// Serialize serializes information the current object
func (m *ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApplicationIdentifier sets the applicationIdentifier property value. The applicationIdentifier property
func (m *ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody) SetApplicationIdentifier(value *string)() {
    m.applicationIdentifier = value
}
// SetCredentials sets the credentials property value. The credentials property
func (m *ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody) SetCredentials(value []SynchronizationSecretKeyStringValuePairable)() {
    m.credentials = value
}
// SetTemplateId sets the templateId property value. The templateId property
func (m *ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody) SetTemplateId(value *string)() {
    m.templateId = value
}
// SetUseSavedCredentials sets the useSavedCredentials property value. The useSavedCredentials property
func (m *ApplicationsItemSynchronizationJobsValidateCredentialsPostRequestBody) SetUseSavedCredentials(value *bool)() {
    m.useSavedCredentials = value
}
