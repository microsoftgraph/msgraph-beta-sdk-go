package ediscovery

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Ediscoveryroot 
type Ediscoveryroot struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The cases property
    cases []Case_escapedable
}
// NewEdiscoveryroot instantiates a new ediscoveryroot and sets the default values.
func NewEdiscoveryroot()(*Ediscoveryroot) {
    m := &Ediscoveryroot{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.ediscovery.ediscoveryroot";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEdiscoveryrootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryrootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryroot(), nil
}
// GetCases gets the cases property value. The cases property
func (m *Ediscoveryroot) GetCases()([]Case_escapedable) {
    return m.cases
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Ediscoveryroot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cases"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCase_escapedFromDiscriminatorValue , m.SetCases)
    return res
}
// Serialize serializes information the current object
func (m *Ediscoveryroot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCases() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCases())
        err = writer.WriteCollectionOfObjectValues("cases", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCases sets the cases property value. The cases property
func (m *Ediscoveryroot) SetCases(value []Case_escapedable)() {
    m.cases = value
}
