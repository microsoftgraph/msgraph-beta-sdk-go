package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DriverUpdateFilter 
type DriverUpdateFilter struct {
    WindowsUpdateFilter
}
// NewDriverUpdateFilter instantiates a new DriverUpdateFilter and sets the default values.
func NewDriverUpdateFilter()(*DriverUpdateFilter) {
    m := &DriverUpdateFilter{
        WindowsUpdateFilter: *NewWindowsUpdateFilter(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.driverUpdateFilter"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDriverUpdateFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriverUpdateFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriverUpdateFilter(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriverUpdateFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsUpdateFilter.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DriverUpdateFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsUpdateFilter.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
