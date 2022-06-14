package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationOptionDefinition 
type DeviceManagementConfigurationOptionDefinition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // List of Settings that depends on this option
    dependedOnBy []DeviceManagementConfigurationSettingDependedOnByable
    // List of dependent settings for this option
    dependentOn []DeviceManagementConfigurationDependentOnable
    // Description of the option
    description *string
    // Friendly name of the option
    displayName *string
    // Help text of the option
    helpText *string
    // Identifier of option
    itemId *string
    // Name of the option
    name *string
    // Value of the option
    optionValue DeviceManagementConfigurationSettingValueable
}
// NewDeviceManagementConfigurationOptionDefinition instantiates a new deviceManagementConfigurationOptionDefinition and sets the default values.
func NewDeviceManagementConfigurationOptionDefinition()(*DeviceManagementConfigurationOptionDefinition) {
    m := &DeviceManagementConfigurationOptionDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationOptionDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationOptionDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationOptionDefinition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationOptionDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDependedOnBy gets the dependedOnBy property value. List of Settings that depends on this option
func (m *DeviceManagementConfigurationOptionDefinition) GetDependedOnBy()([]DeviceManagementConfigurationSettingDependedOnByable) {
    if m == nil {
        return nil
    } else {
        return m.dependedOnBy
    }
}
// GetDependentOn gets the dependentOn property value. List of dependent settings for this option
func (m *DeviceManagementConfigurationOptionDefinition) GetDependentOn()([]DeviceManagementConfigurationDependentOnable) {
    if m == nil {
        return nil
    } else {
        return m.dependentOn
    }
}
// GetDescription gets the description property value. Description of the option
func (m *DeviceManagementConfigurationOptionDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Friendly name of the option
func (m *DeviceManagementConfigurationOptionDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationOptionDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dependedOnBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDependedOnByFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingDependedOnByable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSettingDependedOnByable)
            }
            m.SetDependedOnBy(res)
        }
        return nil
    }
    res["dependentOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationDependentOnFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationDependentOnable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationDependentOnable)
            }
            m.SetDependentOn(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["helpText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHelpText(val)
        }
        return nil
    }
    res["itemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["optionValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSettingValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptionValue(val.(DeviceManagementConfigurationSettingValueable))
        }
        return nil
    }
    return res
}
// GetHelpText gets the helpText property value. Help text of the option
func (m *DeviceManagementConfigurationOptionDefinition) GetHelpText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.helpText
    }
}
// GetItemId gets the itemId property value. Identifier of option
func (m *DeviceManagementConfigurationOptionDefinition) GetItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemId
    }
}
// GetName gets the name property value. Name of the option
func (m *DeviceManagementConfigurationOptionDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetOptionValue gets the optionValue property value. Value of the option
func (m *DeviceManagementConfigurationOptionDefinition) GetOptionValue()(DeviceManagementConfigurationSettingValueable) {
    if m == nil {
        return nil
    } else {
        return m.optionValue
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationOptionDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDependedOnBy() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDependedOnBy()))
        for i, v := range m.GetDependedOnBy() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("dependedOnBy", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDependentOn() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDependentOn()))
        for i, v := range m.GetDependentOn() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("dependentOn", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("helpText", m.GetHelpText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("itemId", m.GetItemId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("optionValue", m.GetOptionValue())
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
func (m *DeviceManagementConfigurationOptionDefinition) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDependedOnBy sets the dependedOnBy property value. List of Settings that depends on this option
func (m *DeviceManagementConfigurationOptionDefinition) SetDependedOnBy(value []DeviceManagementConfigurationSettingDependedOnByable)() {
    if m != nil {
        m.dependedOnBy = value
    }
}
// SetDependentOn sets the dependentOn property value. List of dependent settings for this option
func (m *DeviceManagementConfigurationOptionDefinition) SetDependentOn(value []DeviceManagementConfigurationDependentOnable)() {
    if m != nil {
        m.dependentOn = value
    }
}
// SetDescription sets the description property value. Description of the option
func (m *DeviceManagementConfigurationOptionDefinition) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Friendly name of the option
func (m *DeviceManagementConfigurationOptionDefinition) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetHelpText sets the helpText property value. Help text of the option
func (m *DeviceManagementConfigurationOptionDefinition) SetHelpText(value *string)() {
    if m != nil {
        m.helpText = value
    }
}
// SetItemId sets the itemId property value. Identifier of option
func (m *DeviceManagementConfigurationOptionDefinition) SetItemId(value *string)() {
    if m != nil {
        m.itemId = value
    }
}
// SetName sets the name property value. Name of the option
func (m *DeviceManagementConfigurationOptionDefinition) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetOptionValue sets the optionValue property value. Value of the option
func (m *DeviceManagementConfigurationOptionDefinition) SetOptionValue(value DeviceManagementConfigurationSettingValueable)() {
    if m != nil {
        m.optionValue = value
    }
}
