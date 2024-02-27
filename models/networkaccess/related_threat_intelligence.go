package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RelatedThreatIntelligence struct {
    RelatedResource
}
// NewRelatedThreatIntelligence instantiates a new RelatedThreatIntelligence and sets the default values.
func NewRelatedThreatIntelligence()(*RelatedThreatIntelligence) {
    m := &RelatedThreatIntelligence{
        RelatedResource: *NewRelatedResource(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.relatedThreatIntelligence"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRelatedThreatIntelligenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRelatedThreatIntelligenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelatedThreatIntelligence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RelatedThreatIntelligence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RelatedResource.GetFieldDeserializers()
    res["threatCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatCount(val)
        }
        return nil
    }
    return res
}
// GetThreatCount gets the threatCount property value. The threatCount property
// returns a *int64 when successful
func (m *RelatedThreatIntelligence) GetThreatCount()(*int64) {
    val, err := m.GetBackingStore().Get("threatCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RelatedThreatIntelligence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RelatedResource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("threatCount", m.GetThreatCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetThreatCount sets the threatCount property value. The threatCount property
func (m *RelatedThreatIntelligence) SetThreatCount(value *int64)() {
    err := m.GetBackingStore().Set("threatCount", value)
    if err != nil {
        panic(err)
    }
}
type RelatedThreatIntelligenceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RelatedResourceable
    GetThreatCount()(*int64)
    SetThreatCount(value *int64)()
}
