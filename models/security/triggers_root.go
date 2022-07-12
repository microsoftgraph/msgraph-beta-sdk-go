package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TriggersRoot 
type TriggersRoot struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The retentionEvents property
    retentionEvents []RetentionEventable
}
// NewTriggersRoot instantiates a new triggersRoot and sets the default values.
func NewTriggersRoot()(*TriggersRoot) {
    m := &TriggersRoot{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateTriggersRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTriggersRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTriggersRoot(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TriggersRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["retentionEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRetentionEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RetentionEventable, len(val))
            for i, v := range val {
                res[i] = v.(RetentionEventable)
            }
            m.SetRetentionEvents(res)
        }
        return nil
    }
    return res
}
// GetRetentionEvents gets the retentionEvents property value. The retentionEvents property
func (m *TriggersRoot) GetRetentionEvents()([]RetentionEventable) {
    if m == nil {
        return nil
    } else {
        return m.retentionEvents
    }
}
// Serialize serializes information the current object
func (m *TriggersRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRetentionEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRetentionEvents()))
        for i, v := range m.GetRetentionEvents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("retentionEvents", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRetentionEvents sets the retentionEvents property value. The retentionEvents property
func (m *TriggersRoot) SetRetentionEvents(value []RetentionEventable)() {
    if m != nil {
        m.retentionEvents = value
    }
}
