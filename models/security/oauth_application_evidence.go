package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OauthApplicationEvidence 
type OauthApplicationEvidence struct {
    AlertEvidence
    // Unique identifier of the application.
    appId *string
    // Name of the application.
    displayName *string
    // The unique identifier of the application object in Azure AD.
    objectId *string
    // The name of the application publisher.
    publisher *string
}
// NewOauthApplicationEvidence instantiates a new OauthApplicationEvidence and sets the default values.
func NewOauthApplicationEvidence()(*OauthApplicationEvidence) {
    m := &OauthApplicationEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.oauthApplicationEvidence";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOauthApplicationEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOauthApplicationEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOauthApplicationEvidence(), nil
}
// GetAppId gets the appId property value. Unique identifier of the application.
func (m *OauthApplicationEvidence) GetAppId()(*string) {
    return m.appId
}
// GetDisplayName gets the displayName property value. Name of the application.
func (m *OauthApplicationEvidence) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OauthApplicationEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["appId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppId)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["objectId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetObjectId)
    res["publisher"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPublisher)
    return res
}
// GetObjectId gets the objectId property value. The unique identifier of the application object in Azure AD.
func (m *OauthApplicationEvidence) GetObjectId()(*string) {
    return m.objectId
}
// GetPublisher gets the publisher property value. The name of the application publisher.
func (m *OauthApplicationEvidence) GetPublisher()(*string) {
    return m.publisher
}
// Serialize serializes information the current object
func (m *OauthApplicationEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("objectId", m.GetObjectId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. Unique identifier of the application.
func (m *OauthApplicationEvidence) SetAppId(value *string)() {
    m.appId = value
}
// SetDisplayName sets the displayName property value. Name of the application.
func (m *OauthApplicationEvidence) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetObjectId sets the objectId property value. The unique identifier of the application object in Azure AD.
func (m *OauthApplicationEvidence) SetObjectId(value *string)() {
    m.objectId = value
}
// SetPublisher sets the publisher property value. The name of the application publisher.
func (m *OauthApplicationEvidence) SetPublisher(value *string)() {
    m.publisher = value
}
