package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DeviceManagementConfigurationSettingApplicability struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceManagementConfigurationSettingApplicability instantiates a new DeviceManagementConfigurationSettingApplicability and sets the default values.
func NewDeviceManagementConfigurationSettingApplicability()(*DeviceManagementConfigurationSettingApplicability) {
    m := &DeviceManagementConfigurationSettingApplicability{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceManagementConfigurationSettingApplicabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementConfigurationSettingApplicabilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.deviceManagementConfigurationApplicationSettingApplicability":
                        return NewDeviceManagementConfigurationApplicationSettingApplicability(), nil
                    case "#microsoft.graph.deviceManagementConfigurationExchangeOnlineSettingApplicability":
                        return NewDeviceManagementConfigurationExchangeOnlineSettingApplicability(), nil
                    case "#microsoft.graph.deviceManagementConfigurationWindowsSettingApplicability":
                        return NewDeviceManagementConfigurationWindowsSettingApplicability(), nil
                }
            }
        }
    }
    return NewDeviceManagementConfigurationSettingApplicability(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeviceManagementConfigurationSettingApplicability) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *DeviceManagementConfigurationSettingApplicability) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDescription gets the description property value. description of the setting
// returns a *string when successful
func (m *DeviceManagementConfigurationSettingApplicability) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceMode gets the deviceMode property value. Describes applicability for the mode the device is in
// returns a *DeviceManagementConfigurationDeviceMode when successful
func (m *DeviceManagementConfigurationSettingApplicability) GetDeviceMode()(*DeviceManagementConfigurationDeviceMode) {
    val, err := m.GetBackingStore().Get("deviceMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationDeviceMode)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagementConfigurationSettingApplicability) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["deviceMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationDeviceMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceMode(val.(*DeviceManagementConfigurationDeviceMode))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationPlatforms)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*DeviceManagementConfigurationPlatforms))
        }
        return nil
    }
    res["technologies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTechnologies)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnologies(val.(*DeviceManagementConfigurationTechnologies))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DeviceManagementConfigurationSettingApplicability) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPlatform gets the platform property value. Supported platform types.
// returns a *DeviceManagementConfigurationPlatforms when successful
func (m *DeviceManagementConfigurationSettingApplicability) GetPlatform()(*DeviceManagementConfigurationPlatforms) {
    val, err := m.GetBackingStore().Get("platform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationPlatforms)
    }
    return nil
}
// GetTechnologies gets the technologies property value. Describes which technology this setting can be deployed with
// returns a *DeviceManagementConfigurationTechnologies when successful
func (m *DeviceManagementConfigurationSettingApplicability) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    val, err := m.GetBackingStore().Get("technologies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationTechnologies)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingApplicability) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceMode() != nil {
        cast := (*m.GetDeviceMode()).String()
        err := writer.WriteStringValue("deviceMode", &cast)
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
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err := writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTechnologies() != nil {
        cast := (*m.GetTechnologies()).String()
        err := writer.WriteStringValue("technologies", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingApplicability) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeviceManagementConfigurationSettingApplicability) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDescription sets the description property value. description of the setting
func (m *DeviceManagementConfigurationSettingApplicability) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceMode sets the deviceMode property value. Describes applicability for the mode the device is in
func (m *DeviceManagementConfigurationSettingApplicability) SetDeviceMode(value *DeviceManagementConfigurationDeviceMode)() {
    err := m.GetBackingStore().Set("deviceMode", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingApplicability) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatform sets the platform property value. Supported platform types.
func (m *DeviceManagementConfigurationSettingApplicability) SetPlatform(value *DeviceManagementConfigurationPlatforms)() {
    err := m.GetBackingStore().Set("platform", value)
    if err != nil {
        panic(err)
    }
}
// SetTechnologies sets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationSettingApplicability) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    err := m.GetBackingStore().Set("technologies", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementConfigurationSettingApplicabilityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDescription()(*string)
    GetDeviceMode()(*DeviceManagementConfigurationDeviceMode)
    GetOdataType()(*string)
    GetPlatform()(*DeviceManagementConfigurationPlatforms)
    GetTechnologies()(*DeviceManagementConfigurationTechnologies)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDescription(value *string)()
    SetDeviceMode(value *DeviceManagementConfigurationDeviceMode)()
    SetOdataType(value *string)()
    SetPlatform(value *DeviceManagementConfigurationPlatforms)()
    SetTechnologies(value *DeviceManagementConfigurationTechnologies)()
}
