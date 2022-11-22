package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebApplicationSegment 
type WebApplicationSegment struct {
    ApplicationSegment
    // The alternateUrl property
    alternateUrl *string
    // The corsConfigurations property
    corsConfigurations []CorsConfiguration_v2able
    // The externalUrl property
    externalUrl *string
    // The internalUrl property
    internalUrl *string
}
// NewWebApplicationSegment instantiates a new WebApplicationSegment and sets the default values.
func NewWebApplicationSegment()(*WebApplicationSegment) {
    m := &WebApplicationSegment{
        ApplicationSegment: *NewApplicationSegment(),
    }
    odataTypeValue := "#microsoft.graph.webApplicationSegment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWebApplicationSegmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebApplicationSegmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebApplicationSegment(), nil
}
// GetAlternateUrl gets the alternateUrl property value. The alternateUrl property
func (m *WebApplicationSegment) GetAlternateUrl()(*string) {
    return m.alternateUrl
}
// GetCorsConfigurations gets the corsConfigurations property value. The corsConfigurations property
func (m *WebApplicationSegment) GetCorsConfigurations()([]CorsConfiguration_v2able) {
    return m.corsConfigurations
}
// GetExternalUrl gets the externalUrl property value. The externalUrl property
func (m *WebApplicationSegment) GetExternalUrl()(*string) {
    return m.externalUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebApplicationSegment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ApplicationSegment.GetFieldDeserializers()
    res["alternateUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAlternateUrl)
    res["corsConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCorsConfiguration_v2FromDiscriminatorValue , m.SetCorsConfigurations)
    res["externalUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalUrl)
    res["internalUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInternalUrl)
    return res
}
// GetInternalUrl gets the internalUrl property value. The internalUrl property
func (m *WebApplicationSegment) GetInternalUrl()(*string) {
    return m.internalUrl
}
// Serialize serializes information the current object
func (m *WebApplicationSegment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ApplicationSegment.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("alternateUrl", m.GetAlternateUrl())
        if err != nil {
            return err
        }
    }
    if m.GetCorsConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCorsConfigurations())
        err = writer.WriteCollectionOfObjectValues("corsConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalUrl", m.GetExternalUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("internalUrl", m.GetInternalUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlternateUrl sets the alternateUrl property value. The alternateUrl property
func (m *WebApplicationSegment) SetAlternateUrl(value *string)() {
    m.alternateUrl = value
}
// SetCorsConfigurations sets the corsConfigurations property value. The corsConfigurations property
func (m *WebApplicationSegment) SetCorsConfigurations(value []CorsConfiguration_v2able)() {
    m.corsConfigurations = value
}
// SetExternalUrl sets the externalUrl property value. The externalUrl property
func (m *WebApplicationSegment) SetExternalUrl(value *string)() {
    m.externalUrl = value
}
// SetInternalUrl sets the internalUrl property value. The internalUrl property
func (m *WebApplicationSegment) SetInternalUrl(value *string)() {
    m.internalUrl = value
}
