// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WindowsUpdateFilter struct {
    SoftwareUpdateFilter
}
// NewWindowsUpdateFilter instantiates a new WindowsUpdateFilter and sets the default values.
func NewWindowsUpdateFilter()(*WindowsUpdateFilter) {
    m := &WindowsUpdateFilter{
        SoftwareUpdateFilter: *NewSoftwareUpdateFilter(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.windowsUpdateFilter"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsUpdateFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsUpdateFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windowsUpdates.driverUpdateFilter":
                        return NewDriverUpdateFilter(), nil
                    case "#microsoft.graph.windowsUpdates.qualityUpdateFilter":
                        return NewQualityUpdateFilter(), nil
                    case "#microsoft.graph.windowsUpdates.remediationUpdateFilter":
                        return NewRemediationUpdateFilter(), nil
                }
            }
        }
    }
    return NewWindowsUpdateFilter(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsUpdateFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SoftwareUpdateFilter.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *WindowsUpdateFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SoftwareUpdateFilter.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type WindowsUpdateFilterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SoftwareUpdateFilterable
}
