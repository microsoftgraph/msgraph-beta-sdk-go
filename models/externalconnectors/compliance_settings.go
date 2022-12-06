package externalconnectors

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComplianceSettings 
type ComplianceSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Enables the developer to define the appearance of the content and configure conditions that dictate when the template should be displayed. Maximum of two eDiscovery result templates per connection.
    eDiscoveryResultTemplates []DisplayTemplateable
    // The OdataType property
    odataType *string
}
// NewComplianceSettings instantiates a new complianceSettings and sets the default values.
func NewComplianceSettings()(*ComplianceSettings) {
    m := &ComplianceSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateComplianceSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComplianceSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComplianceSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComplianceSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEDiscoveryResultTemplates gets the eDiscoveryResultTemplates property value. Enables the developer to define the appearance of the content and configure conditions that dictate when the template should be displayed. Maximum of two eDiscovery result templates per connection.
func (m *ComplianceSettings) GetEDiscoveryResultTemplates()([]DisplayTemplateable) {
    return m.eDiscoveryResultTemplates
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComplianceSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["eDiscoveryResultTemplates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDisplayTemplateFromDiscriminatorValue , m.SetEDiscoveryResultTemplates)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ComplianceSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ComplianceSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEDiscoveryResultTemplates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEDiscoveryResultTemplates())
        err := writer.WriteCollectionOfObjectValues("eDiscoveryResultTemplates", cast)
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
func (m *ComplianceSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEDiscoveryResultTemplates sets the eDiscoveryResultTemplates property value. Enables the developer to define the appearance of the content and configure conditions that dictate when the template should be displayed. Maximum of two eDiscovery result templates per connection.
func (m *ComplianceSettings) SetEDiscoveryResultTemplates(value []DisplayTemplateable)() {
    m.eDiscoveryResultTemplates = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ComplianceSettings) SetOdataType(value *string)() {
    m.odataType = value
}
