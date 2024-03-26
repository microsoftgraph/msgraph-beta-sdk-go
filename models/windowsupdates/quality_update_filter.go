package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type QualityUpdateFilter struct {
    WindowsUpdateFilter
}
// NewQualityUpdateFilter instantiates a new QualityUpdateFilter and sets the default values.
func NewQualityUpdateFilter()(*QualityUpdateFilter) {
    m := &QualityUpdateFilter{
        WindowsUpdateFilter: *NewWindowsUpdateFilter(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.qualityUpdateFilter"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateQualityUpdateFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateQualityUpdateFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewQualityUpdateFilter(), nil
}
// GetCadence gets the cadence property value. The cadence property
// returns a *QualityUpdateCadence when successful
func (m *QualityUpdateFilter) GetCadence()(*QualityUpdateCadence) {
    val, err := m.GetBackingStore().Get("cadence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*QualityUpdateCadence)
    }
    return nil
}
// GetClassification gets the classification property value. The classification property
// returns a *QualityUpdateClassification when successful
func (m *QualityUpdateFilter) GetClassification()(*QualityUpdateClassification) {
    val, err := m.GetBackingStore().Get("classification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*QualityUpdateClassification)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *QualityUpdateFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsUpdateFilter.GetFieldDeserializers()
    res["cadence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseQualityUpdateCadence)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCadence(val.(*QualityUpdateCadence))
        }
        return nil
    }
    res["classification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseQualityUpdateClassification)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val.(*QualityUpdateClassification))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *QualityUpdateFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsUpdateFilter.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCadence() != nil {
        cast := (*m.GetCadence()).String()
        err = writer.WriteStringValue("cadence", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetClassification() != nil {
        cast := (*m.GetClassification()).String()
        err = writer.WriteStringValue("classification", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCadence sets the cadence property value. The cadence property
func (m *QualityUpdateFilter) SetCadence(value *QualityUpdateCadence)() {
    err := m.GetBackingStore().Set("cadence", value)
    if err != nil {
        panic(err)
    }
}
// SetClassification sets the classification property value. The classification property
func (m *QualityUpdateFilter) SetClassification(value *QualityUpdateClassification)() {
    err := m.GetBackingStore().Set("classification", value)
    if err != nil {
        panic(err)
    }
}
type QualityUpdateFilterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsUpdateFilterable
    GetCadence()(*QualityUpdateCadence)
    GetClassification()(*QualityUpdateClassification)
    SetCadence(value *QualityUpdateCadence)()
    SetClassification(value *QualityUpdateClassification)()
}
