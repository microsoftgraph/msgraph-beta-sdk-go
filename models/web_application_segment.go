package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebApplicationSegment 
type WebApplicationSegment struct {
    ApplicationSegment
}
// NewWebApplicationSegment instantiates a new WebApplicationSegment and sets the default values.
func NewWebApplicationSegment()(*WebApplicationSegment) {
    m := &WebApplicationSegment{
        ApplicationSegment: *NewApplicationSegment(),
    }
    odataTypeValue := "#microsoft.graph.webApplicationSegment"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWebApplicationSegmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebApplicationSegmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebApplicationSegment(), nil
}
// GetAlternateUrl gets the alternateUrl property value. If you're configuring a traffic manager in front of multiple App Proxy application segments, this property contains the user-friendly URL that will point to the traffic manager.
func (m *WebApplicationSegment) GetAlternateUrl()(*string) {
    val, err := m.GetBackingStore().Get("alternateUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCorsConfigurations gets the corsConfigurations property value. A collection of CORS Rule definitions for a particular application segment.
func (m *WebApplicationSegment) GetCorsConfigurations()([]CorsConfiguration_v2able) {
    val, err := m.GetBackingStore().Get("corsConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CorsConfiguration_v2able)
    }
    return nil
}
// GetExternalUrl gets the externalUrl property value. The published external URL for the application segment; for example, https://intranet.contoso.com/.
func (m *WebApplicationSegment) GetExternalUrl()(*string) {
    val, err := m.GetBackingStore().Get("externalUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebApplicationSegment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ApplicationSegment.GetFieldDeserializers()
    res["alternateUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternateUrl(val)
        }
        return nil
    }
    res["corsConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCorsConfiguration_v2FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CorsConfiguration_v2able, len(val))
            for i, v := range val {
                res[i] = v.(CorsConfiguration_v2able)
            }
            m.SetCorsConfigurations(res)
        }
        return nil
    }
    res["externalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalUrl(val)
        }
        return nil
    }
    res["internalUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalUrl(val)
        }
        return nil
    }
    return res
}
// GetInternalUrl gets the internalUrl property value. The internal URL of the application segment; for example, https://intranet/.
func (m *WebApplicationSegment) GetInternalUrl()(*string) {
    val, err := m.GetBackingStore().Get("internalUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCorsConfigurations()))
        for i, v := range m.GetCorsConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
// SetAlternateUrl sets the alternateUrl property value. If you're configuring a traffic manager in front of multiple App Proxy application segments, this property contains the user-friendly URL that will point to the traffic manager.
func (m *WebApplicationSegment) SetAlternateUrl(value *string)() {
    err := m.GetBackingStore().Set("alternateUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCorsConfigurations sets the corsConfigurations property value. A collection of CORS Rule definitions for a particular application segment.
func (m *WebApplicationSegment) SetCorsConfigurations(value []CorsConfiguration_v2able)() {
    err := m.GetBackingStore().Set("corsConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalUrl sets the externalUrl property value. The published external URL for the application segment; for example, https://intranet.contoso.com/.
func (m *WebApplicationSegment) SetExternalUrl(value *string)() {
    err := m.GetBackingStore().Set("externalUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetInternalUrl sets the internalUrl property value. The internal URL of the application segment; for example, https://intranet/.
func (m *WebApplicationSegment) SetInternalUrl(value *string)() {
    err := m.GetBackingStore().Set("internalUrl", value)
    if err != nil {
        panic(err)
    }
}
// WebApplicationSegmentable 
type WebApplicationSegmentable interface {
    ApplicationSegmentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlternateUrl()(*string)
    GetCorsConfigurations()([]CorsConfiguration_v2able)
    GetExternalUrl()(*string)
    GetInternalUrl()(*string)
    SetAlternateUrl(value *string)()
    SetCorsConfigurations(value []CorsConfiguration_v2able)()
    SetExternalUrl(value *string)()
    SetInternalUrl(value *string)()
}
