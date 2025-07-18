// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TlsInspectionWebCategoryDestination struct {
    TlsInspectionDestination
}
// NewTlsInspectionWebCategoryDestination instantiates a new TlsInspectionWebCategoryDestination and sets the default values.
func NewTlsInspectionWebCategoryDestination()(*TlsInspectionWebCategoryDestination) {
    m := &TlsInspectionWebCategoryDestination{
        TlsInspectionDestination: *NewTlsInspectionDestination(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.tlsInspectionWebCategoryDestination"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTlsInspectionWebCategoryDestinationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTlsInspectionWebCategoryDestinationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTlsInspectionWebCategoryDestination(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TlsInspectionWebCategoryDestination) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TlsInspectionDestination.GetFieldDeserializers()
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetValues gets the values property value. A collection of web category names to match against. This collection cannot be empty or null.
// returns a []string when successful
func (m *TlsInspectionWebCategoryDestination) GetValues()([]string) {
    val, err := m.GetBackingStore().Get("values")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TlsInspectionWebCategoryDestination) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TlsInspectionDestination.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValues() != nil {
        err = writer.WriteCollectionOfStringValues("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValues sets the values property value. A collection of web category names to match against. This collection cannot be empty or null.
func (m *TlsInspectionWebCategoryDestination) SetValues(value []string)() {
    err := m.GetBackingStore().Set("values", value)
    if err != nil {
        panic(err)
    }
}
type TlsInspectionWebCategoryDestinationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TlsInspectionDestinationable
    GetValues()([]string)
    SetValues(value []string)()
}
