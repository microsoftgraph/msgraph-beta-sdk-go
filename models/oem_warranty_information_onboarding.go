package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OemWarrantyInformationOnboarding warranty status entity for a given OEM
type OemWarrantyInformationOnboarding struct {
    Entity
    // Specifies whether warranty API is available. This property is read-only.
    available *bool
    // Specifies whether warranty query is enabled for given OEM. This property is read-only.
    enabled *bool
    // OEM name. This property is read-only.
    oemName *string
}
// NewOemWarrantyInformationOnboarding instantiates a new oemWarrantyInformationOnboarding and sets the default values.
func NewOemWarrantyInformationOnboarding()(*OemWarrantyInformationOnboarding) {
    m := &OemWarrantyInformationOnboarding{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.oemWarrantyInformationOnboarding";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOemWarrantyInformationOnboardingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOemWarrantyInformationOnboardingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOemWarrantyInformationOnboarding(), nil
}
// GetAvailable gets the available property value. Specifies whether warranty API is available. This property is read-only.
func (m *OemWarrantyInformationOnboarding) GetAvailable()(*bool) {
    return m.available
}
// GetEnabled gets the enabled property value. Specifies whether warranty query is enabled for given OEM. This property is read-only.
func (m *OemWarrantyInformationOnboarding) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OemWarrantyInformationOnboarding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// GetOemName gets the oemName property value. OEM name. This property is read-only.
func (m *OemWarrantyInformationOnboarding) GetOemName()(*string) {
    return m.oemName
}
// Serialize serializes information the current object
func (m *OemWarrantyInformationOnboarding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
