package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleDeviceFeaturesConfigurationBase apple device features configuration profile.
type AppleDeviceFeaturesConfigurationBase struct {
    DeviceConfiguration
    // An array of AirPrint printers that should always be shown. This collection can contain a maximum of 500 elements.
    airPrintDestinations []AirPrintDestinationable
}
// NewAppleDeviceFeaturesConfigurationBase instantiates a new appleDeviceFeaturesConfigurationBase and sets the default values.
func NewAppleDeviceFeaturesConfigurationBase()(*AppleDeviceFeaturesConfigurationBase) {
    m := &AppleDeviceFeaturesConfigurationBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.appleDeviceFeaturesConfigurationBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAppleDeviceFeaturesConfigurationBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleDeviceFeaturesConfigurationBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.iosDeviceFeaturesConfiguration":
                        return NewIosDeviceFeaturesConfiguration(), nil
                    case "#microsoft.graph.macOSDeviceFeaturesConfiguration":
                        return NewMacOSDeviceFeaturesConfiguration(), nil
                }
            }
        }
    }
    return NewAppleDeviceFeaturesConfigurationBase(), nil
}
// GetAirPrintDestinations gets the airPrintDestinations property value. An array of AirPrint printers that should always be shown. This collection can contain a maximum of 500 elements.
func (m *AppleDeviceFeaturesConfigurationBase) GetAirPrintDestinations()([]AirPrintDestinationable) {
    return m.airPrintDestinations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleDeviceFeaturesConfigurationBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["airPrintDestinations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAirPrintDestinationFromDiscriminatorValue , m.SetAirPrintDestinations)
    return res
}
// Serialize serializes information the current object
func (m *AppleDeviceFeaturesConfigurationBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAirPrintDestinations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAirPrintDestinations())
        err = writer.WriteCollectionOfObjectValues("airPrintDestinations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAirPrintDestinations sets the airPrintDestinations property value. An array of AirPrint printers that should always be shown. This collection can contain a maximum of 500 elements.
func (m *AppleDeviceFeaturesConfigurationBase) SetAirPrintDestinations(value []AirPrintDestinationable)() {
    m.airPrintDestinations = value
}
