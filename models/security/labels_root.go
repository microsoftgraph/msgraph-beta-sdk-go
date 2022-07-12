package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// LabelsRoot 
type LabelsRoot struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The retentionLabels property
    retentionLabels []RetentionLabelable
}
// NewLabelsRoot instantiates a new labelsRoot and sets the default values.
func NewLabelsRoot()(*LabelsRoot) {
    m := &LabelsRoot{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateLabelsRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLabelsRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelsRoot(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LabelsRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["retentionLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRetentionLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RetentionLabelable, len(val))
            for i, v := range val {
                res[i] = v.(RetentionLabelable)
            }
            m.SetRetentionLabels(res)
        }
        return nil
    }
    return res
}
// GetRetentionLabels gets the retentionLabels property value. The retentionLabels property
func (m *LabelsRoot) GetRetentionLabels()([]RetentionLabelable) {
    if m == nil {
        return nil
    } else {
        return m.retentionLabels
    }
}
// Serialize serializes information the current object
func (m *LabelsRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRetentionLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRetentionLabels()))
        for i, v := range m.GetRetentionLabels() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("retentionLabels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRetentionLabels sets the retentionLabels property value. The retentionLabels property
func (m *LabelsRoot) SetRetentionLabels(value []RetentionLabelable)() {
    if m != nil {
        m.retentionLabels = value
    }
}
