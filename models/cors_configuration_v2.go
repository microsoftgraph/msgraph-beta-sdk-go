package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CorsConfiguration_v2 provides operations to manage the collection of accessReview entities.
type CorsConfiguration_v2 struct {
    Entity
    // The allowedHeaders property
    allowedHeaders []string
    // The allowedMethods property
    allowedMethods []string
    // The allowedOrigins property
    allowedOrigins []string
    // The maxAgeInSeconds property
    maxAgeInSeconds *int32
    // The resource property
    resource *string
}
// NewCorsConfiguration_v2 instantiates a new corsConfiguration_v2 and sets the default values.
func NewCorsConfiguration_v2()(*CorsConfiguration_v2) {
    m := &CorsConfiguration_v2{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCorsConfiguration_v2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCorsConfiguration_v2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCorsConfiguration_v2(), nil
}
// GetAllowedHeaders gets the allowedHeaders property value. The allowedHeaders property
func (m *CorsConfiguration_v2) GetAllowedHeaders()([]string) {
    return m.allowedHeaders
}
// GetAllowedMethods gets the allowedMethods property value. The allowedMethods property
func (m *CorsConfiguration_v2) GetAllowedMethods()([]string) {
    return m.allowedMethods
}
// GetAllowedOrigins gets the allowedOrigins property value. The allowedOrigins property
func (m *CorsConfiguration_v2) GetAllowedOrigins()([]string) {
    return m.allowedOrigins
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CorsConfiguration_v2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedHeaders"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAllowedHeaders)
    res["allowedMethods"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAllowedMethods)
    res["allowedOrigins"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAllowedOrigins)
    res["maxAgeInSeconds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMaxAgeInSeconds)
    res["resource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetResource)
    return res
}
// GetMaxAgeInSeconds gets the maxAgeInSeconds property value. The maxAgeInSeconds property
func (m *CorsConfiguration_v2) GetMaxAgeInSeconds()(*int32) {
    return m.maxAgeInSeconds
}
// GetResource gets the resource property value. The resource property
func (m *CorsConfiguration_v2) GetResource()(*string) {
    return m.resource
}
// Serialize serializes information the current object
func (m *CorsConfiguration_v2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedHeaders() != nil {
        err = writer.WriteCollectionOfStringValues("allowedHeaders", m.GetAllowedHeaders())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedMethods() != nil {
        err = writer.WriteCollectionOfStringValues("allowedMethods", m.GetAllowedMethods())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedOrigins() != nil {
        err = writer.WriteCollectionOfStringValues("allowedOrigins", m.GetAllowedOrigins())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maxAgeInSeconds", m.GetMaxAgeInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedHeaders sets the allowedHeaders property value. The allowedHeaders property
func (m *CorsConfiguration_v2) SetAllowedHeaders(value []string)() {
    m.allowedHeaders = value
}
// SetAllowedMethods sets the allowedMethods property value. The allowedMethods property
func (m *CorsConfiguration_v2) SetAllowedMethods(value []string)() {
    m.allowedMethods = value
}
// SetAllowedOrigins sets the allowedOrigins property value. The allowedOrigins property
func (m *CorsConfiguration_v2) SetAllowedOrigins(value []string)() {
    m.allowedOrigins = value
}
// SetMaxAgeInSeconds sets the maxAgeInSeconds property value. The maxAgeInSeconds property
func (m *CorsConfiguration_v2) SetMaxAgeInSeconds(value *int32)() {
    m.maxAgeInSeconds = value
}
// SetResource sets the resource property value. The resource property
func (m *CorsConfiguration_v2) SetResource(value *string)() {
    m.resource = value
}
