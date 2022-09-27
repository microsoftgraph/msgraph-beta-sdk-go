package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsageRight provides operations to manage the collection of activityStatistics entities.
type UsageRight struct {
    Entity
    // Product id corresponding to the usage right.
    catalogId *string
    // Identifier of the service corresponding to the usage right.
    serviceIdentifier *string
    // The state property
    state *UsageRightState
}
// NewUsageRight instantiates a new usageRight and sets the default values.
func NewUsageRight()(*UsageRight) {
    m := &UsageRight{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.usageRight";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUsageRightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsageRightFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsageRight(), nil
}
// GetCatalogId gets the catalogId property value. Product id corresponding to the usage right.
func (m *UsageRight) GetCatalogId()(*string) {
    return m.catalogId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsageRight) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["catalogId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCatalogId)
    res["serviceIdentifier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetServiceIdentifier)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseUsageRightState , m.SetState)
    return res
}
// GetServiceIdentifier gets the serviceIdentifier property value. Identifier of the service corresponding to the usage right.
func (m *UsageRight) GetServiceIdentifier()(*string) {
    return m.serviceIdentifier
}
// GetState gets the state property value. The state property
func (m *UsageRight) GetState()(*UsageRightState) {
    return m.state
}
// Serialize serializes information the current object
func (m *UsageRight) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("catalogId", m.GetCatalogId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceIdentifier", m.GetServiceIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCatalogId sets the catalogId property value. Product id corresponding to the usage right.
func (m *UsageRight) SetCatalogId(value *string)() {
    m.catalogId = value
}
// SetServiceIdentifier sets the serviceIdentifier property value. Identifier of the service corresponding to the usage right.
func (m *UsageRight) SetServiceIdentifier(value *string)() {
    m.serviceIdentifier = value
}
// SetState sets the state property value. The state property
func (m *UsageRight) SetState(value *UsageRightState)() {
    m.state = value
}
