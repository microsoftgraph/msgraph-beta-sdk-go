package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BasicFilter struct {
    Filter
}
// NewBasicFilter instantiates a new BasicFilter and sets the default values.
func NewBasicFilter()(*BasicFilter) {
    m := &BasicFilter{
        Filter: *NewFilter(),
    }
    odataTypeValue := "#microsoft.graph.industryData.basicFilter"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateBasicFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBasicFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBasicFilter(), nil
}
// GetAttribute gets the attribute property value. The attribute property
// returns a *FilterOptions when successful
func (m *BasicFilter) GetAttribute()(*FilterOptions) {
    val, err := m.GetBackingStore().Get("attribute")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*FilterOptions)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BasicFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Filter.GetFieldDeserializers()
    res["attribute"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFilterOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttribute(val.(*FilterOptions))
        }
        return nil
    }
    res["in"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetIn(res)
        }
        return nil
    }
    return res
}
// GetIn gets the in property value. The condition to filter with.
// returns a []string when successful
func (m *BasicFilter) GetIn()([]string) {
    val, err := m.GetBackingStore().Get("in")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BasicFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Filter.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttribute() != nil {
        cast := (*m.GetAttribute()).String()
        err = writer.WriteStringValue("attribute", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIn() != nil {
        err = writer.WriteCollectionOfStringValues("in", m.GetIn())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttribute sets the attribute property value. The attribute property
func (m *BasicFilter) SetAttribute(value *FilterOptions)() {
    err := m.GetBackingStore().Set("attribute", value)
    if err != nil {
        panic(err)
    }
}
// SetIn sets the in property value. The condition to filter with.
func (m *BasicFilter) SetIn(value []string)() {
    err := m.GetBackingStore().Set("in", value)
    if err != nil {
        panic(err)
    }
}
type BasicFilterable interface {
    Filterable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttribute()(*FilterOptions)
    GetIn()([]string)
    SetAttribute(value *FilterOptions)()
    SetIn(value []string)()
}
