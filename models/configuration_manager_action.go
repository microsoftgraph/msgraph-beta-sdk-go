package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConfigurationManagerAction parameter for action triggerConfigurationManagerAction
type ConfigurationManagerAction struct {
    // Action type on Configuration Manager client
    action *ConfigurationManagerActionType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewConfigurationManagerAction instantiates a new configurationManagerAction and sets the default values.
func NewConfigurationManagerAction()(*ConfigurationManagerAction) {
    m := &ConfigurationManagerAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConfigurationManagerActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfigurationManagerActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfigurationManagerAction(), nil
}
// GetAction gets the action property value. Action type on Configuration Manager client
func (m *ConfigurationManagerAction) GetAction()(*ConfigurationManagerActionType) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigurationManagerAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigurationManagerAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationManagerActionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*ConfigurationManagerActionType))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ConfigurationManagerAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
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
// SetAction sets the action property value. Action type on Configuration Manager client
func (m *ConfigurationManagerAction) SetAction(value *ConfigurationManagerActionType)() {
    if m != nil {
        m.action = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigurationManagerAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
