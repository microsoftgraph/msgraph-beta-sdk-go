package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SoftwareUpdateFilter struct {
    ContentFilter
}
// NewSoftwareUpdateFilter instantiates a new SoftwareUpdateFilter and sets the default values.
func NewSoftwareUpdateFilter()(*SoftwareUpdateFilter) {
    m := &SoftwareUpdateFilter{
        ContentFilter: *NewContentFilter(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.softwareUpdateFilter"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSoftwareUpdateFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSoftwareUpdateFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windowsUpdates.windowsUpdateFilter":
                        return NewWindowsUpdateFilter(), nil
                }
            }
        }
    }
    return NewSoftwareUpdateFilter(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SoftwareUpdateFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ContentFilter.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SoftwareUpdateFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ContentFilter.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type SoftwareUpdateFilterable interface {
    ContentFilterable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
