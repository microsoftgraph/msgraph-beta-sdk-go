package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingCollectionInstance 
type DeviceManagementConfigurationChoiceSettingCollectionInstance struct {
    DeviceManagementConfigurationSettingInstance
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Choice setting collection value
    choiceSettingCollectionValue []DeviceManagementConfigurationChoiceSettingValueable
}
// NewDeviceManagementConfigurationChoiceSettingCollectionInstance instantiates a new DeviceManagementConfigurationChoiceSettingCollectionInstance and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingCollectionInstance()(*DeviceManagementConfigurationChoiceSettingCollectionInstance) {
    m := &DeviceManagementConfigurationChoiceSettingCollectionInstance{
        DeviceManagementConfigurationSettingInstance: *NewDeviceManagementConfigurationSettingInstance(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingCollectionInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingCollectionInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingCollectionInstance(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstance) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetChoiceSettingCollectionValue gets the choiceSettingCollectionValue property value. Choice setting collection value
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstance) GetChoiceSettingCollectionValue()([]DeviceManagementConfigurationChoiceSettingValueable) {
    if m == nil {
        return nil
    } else {
        return m.choiceSettingCollectionValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstance.GetFieldDeserializers()
    res["choiceSettingCollectionValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationChoiceSettingValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationChoiceSettingValueable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationChoiceSettingValueable)
            }
            m.SetChoiceSettingCollectionValue(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChoiceSettingCollectionValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChoiceSettingCollectionValue()))
        for i, v := range m.GetChoiceSettingCollectionValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("choiceSettingCollectionValue", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstance) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetChoiceSettingCollectionValue sets the choiceSettingCollectionValue property value. Choice setting collection value
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstance) SetChoiceSettingCollectionValue(value []DeviceManagementConfigurationChoiceSettingValueable)() {
    if m != nil {
        m.choiceSettingCollectionValue = value
    }
}
