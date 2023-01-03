package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Windows 
type Windows struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
    updates Updatesable
}
// NewWindows instantiates a new Windows and sets the default values.
func NewWindows()(*Windows) {
    m := &Windows{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateWindowsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["updates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUpdatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdates(val.(Updatesable))
        }
        return nil
    }
    return res
}
// GetUpdates gets the updates property value. Entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
func (m *Windows) GetUpdates()(Updatesable) {
    return m.updates
}
// Serialize serializes information the current object
func (m *Windows) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("updates", m.GetUpdates())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUpdates sets the updates property value. Entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
func (m *Windows) SetUpdates(value Updatesable)() {
    m.updates = value
}
