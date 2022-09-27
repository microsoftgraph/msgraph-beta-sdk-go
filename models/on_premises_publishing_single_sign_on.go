package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesPublishingSingleSignOn 
type OnPremisesPublishingSingleSignOn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
    kerberosSignOnSettings KerberosSignOnSettingsable
    // The OdataType property
    odataType *string
    // The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased.
    singleSignOnMode *SingleSignOnMode
}
// NewOnPremisesPublishingSingleSignOn instantiates a new onPremisesPublishingSingleSignOn and sets the default values.
func NewOnPremisesPublishingSingleSignOn()(*OnPremisesPublishingSingleSignOn) {
    m := &OnPremisesPublishingSingleSignOn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.onPremisesPublishingSingleSignOn";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOnPremisesPublishingSingleSignOnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesPublishingSingleSignOnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesPublishingSingleSignOn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesPublishingSingleSignOn) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesPublishingSingleSignOn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["kerberosSignOnSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateKerberosSignOnSettingsFromDiscriminatorValue , m.SetKerberosSignOnSettings)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["singleSignOnMode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSingleSignOnMode , m.SetSingleSignOnMode)
    return res
}
// GetKerberosSignOnSettings gets the kerberosSignOnSettings property value. The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
func (m *OnPremisesPublishingSingleSignOn) GetKerberosSignOnSettings()(KerberosSignOnSettingsable) {
    return m.kerberosSignOnSettings
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnPremisesPublishingSingleSignOn) GetOdataType()(*string) {
    return m.odataType
}
// GetSingleSignOnMode gets the singleSignOnMode property value. The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased.
func (m *OnPremisesPublishingSingleSignOn) GetSingleSignOnMode()(*SingleSignOnMode) {
    return m.singleSignOnMode
}
// Serialize serializes information the current object
func (m *OnPremisesPublishingSingleSignOn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("kerberosSignOnSettings", m.GetKerberosSignOnSettings())
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
    if m.GetSingleSignOnMode() != nil {
        cast := (*m.GetSingleSignOnMode()).String()
        err := writer.WriteStringValue("singleSignOnMode", &cast)
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
func (m *OnPremisesPublishingSingleSignOn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKerberosSignOnSettings sets the kerberosSignOnSettings property value. The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
func (m *OnPremisesPublishingSingleSignOn) SetKerberosSignOnSettings(value KerberosSignOnSettingsable)() {
    m.kerberosSignOnSettings = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesPublishingSingleSignOn) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSingleSignOnMode sets the singleSignOnMode property value. The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased.
func (m *OnPremisesPublishingSingleSignOn) SetSingleSignOnMode(value *SingleSignOnMode)() {
    m.singleSignOnMode = value
}
