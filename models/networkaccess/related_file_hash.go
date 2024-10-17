package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RelatedFileHash struct {
    RelatedResource
}
// NewRelatedFileHash instantiates a new RelatedFileHash and sets the default values.
func NewRelatedFileHash()(*RelatedFileHash) {
    m := &RelatedFileHash{
        RelatedResource: *NewRelatedResource(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.relatedFileHash"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRelatedFileHashFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRelatedFileHashFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelatedFileHash(), nil
}
// GetAlgorithm gets the algorithm property value. The algorithm property
// returns a *Algorithm when successful
func (m *RelatedFileHash) GetAlgorithm()(*Algorithm) {
    val, err := m.GetBackingStore().Get("algorithm")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Algorithm)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RelatedFileHash) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RelatedResource.GetFieldDeserializers()
    res["algorithm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlgorithm)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlgorithm(val.(*Algorithm))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a *string when successful
func (m *RelatedFileHash) GetValue()(*string) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RelatedFileHash) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RelatedResource.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlgorithm() != nil {
        cast := (*m.GetAlgorithm()).String()
        err = writer.WriteStringValue("algorithm", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlgorithm sets the algorithm property value. The algorithm property
func (m *RelatedFileHash) SetAlgorithm(value *Algorithm)() {
    err := m.GetBackingStore().Set("algorithm", value)
    if err != nil {
        panic(err)
    }
}
// SetValue sets the value property value. The value property
func (m *RelatedFileHash) SetValue(value *string)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type RelatedFileHashable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RelatedResourceable
    GetAlgorithm()(*Algorithm)
    GetValue()(*string)
    SetAlgorithm(value *Algorithm)()
    SetValue(value *string)()
}
