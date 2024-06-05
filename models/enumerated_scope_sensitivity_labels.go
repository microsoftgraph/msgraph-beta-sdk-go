package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnumeratedScopeSensitivityLabels struct {
    ScopeSensitivityLabels
}
// NewEnumeratedScopeSensitivityLabels instantiates a new EnumeratedScopeSensitivityLabels and sets the default values.
func NewEnumeratedScopeSensitivityLabels()(*EnumeratedScopeSensitivityLabels) {
    m := &EnumeratedScopeSensitivityLabels{
        ScopeSensitivityLabels: *NewScopeSensitivityLabels(),
    }
    odataTypeValue := "#microsoft.graph.enumeratedScopeSensitivityLabels"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEnumeratedScopeSensitivityLabelsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnumeratedScopeSensitivityLabelsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnumeratedScopeSensitivityLabels(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnumeratedScopeSensitivityLabels) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ScopeSensitivityLabels.GetFieldDeserializers()
    res["sensitivityLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSensitivityLabels(res)
        }
        return nil
    }
    return res
}
// GetSensitivityLabels gets the sensitivityLabels property value. The sensitivity labels that are applicable to the scope type and have been preapproved. Required.
// returns a []string when successful
func (m *EnumeratedScopeSensitivityLabels) GetSensitivityLabels()([]string) {
    val, err := m.GetBackingStore().Get("sensitivityLabels")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EnumeratedScopeSensitivityLabels) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ScopeSensitivityLabels.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSensitivityLabels() != nil {
        err = writer.WriteCollectionOfStringValues("sensitivityLabels", m.GetSensitivityLabels())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSensitivityLabels sets the sensitivityLabels property value. The sensitivity labels that are applicable to the scope type and have been preapproved. Required.
func (m *EnumeratedScopeSensitivityLabels) SetSensitivityLabels(value []string)() {
    err := m.GetBackingStore().Set("sensitivityLabels", value)
    if err != nil {
        panic(err)
    }
}
type EnumeratedScopeSensitivityLabelsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ScopeSensitivityLabelsable
    GetSensitivityLabels()([]string)
    SetSensitivityLabels(value []string)()
}
