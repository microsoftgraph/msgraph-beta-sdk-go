package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceNotificationPolicy 
type GovernanceNotificationPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enabledTemplateTypes property
    enabledTemplateTypes []string
    // The notificationTemplates property
    notificationTemplates []GovernanceNotificationTemplateable
    // The OdataType property
    odataType *string
}
// NewGovernanceNotificationPolicy instantiates a new governanceNotificationPolicy and sets the default values.
func NewGovernanceNotificationPolicy()(*GovernanceNotificationPolicy) {
    m := &GovernanceNotificationPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGovernanceNotificationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceNotificationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernanceNotificationPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceNotificationPolicy) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnabledTemplateTypes gets the enabledTemplateTypes property value. The enabledTemplateTypes property
func (m *GovernanceNotificationPolicy) GetEnabledTemplateTypes()([]string) {
    return m.enabledTemplateTypes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceNotificationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabledTemplateTypes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetEnabledTemplateTypes)
    res["notificationTemplates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGovernanceNotificationTemplateFromDiscriminatorValue , m.SetNotificationTemplates)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetNotificationTemplates gets the notificationTemplates property value. The notificationTemplates property
func (m *GovernanceNotificationPolicy) GetNotificationTemplates()([]GovernanceNotificationTemplateable) {
    return m.notificationTemplates
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GovernanceNotificationPolicy) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *GovernanceNotificationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnabledTemplateTypes() != nil {
        err := writer.WriteCollectionOfStringValues("enabledTemplateTypes", m.GetEnabledTemplateTypes())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationTemplates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetNotificationTemplates())
        err := writer.WriteCollectionOfObjectValues("notificationTemplates", cast)
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
func (m *GovernanceNotificationPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnabledTemplateTypes sets the enabledTemplateTypes property value. The enabledTemplateTypes property
func (m *GovernanceNotificationPolicy) SetEnabledTemplateTypes(value []string)() {
    m.enabledTemplateTypes = value
}
// SetNotificationTemplates sets the notificationTemplates property value. The notificationTemplates property
func (m *GovernanceNotificationPolicy) SetNotificationTemplates(value []GovernanceNotificationTemplateable)() {
    m.notificationTemplates = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GovernanceNotificationPolicy) SetOdataType(value *string)() {
    m.odataType = value
}
