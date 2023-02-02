package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CorsConfiguration_v2 
type CorsConfiguration_v2 struct {
    Entity
    // The request headers that the origin domain may specify on the CORS request. The wildcard character * indicates that any header beginning with the specified prefix is allowed.
    allowedHeaders []string
    // The HTTP request methods that the origin domain may use for a CORS request.
    allowedMethods []string
    // The origin domains that are permitted to make a request against the service via CORS. The origin domain is the domain from which the request originates. The origin must be an exact case-sensitive match with the origin that the user agent sends to the service.
    allowedOrigins []string
    // The maximum amount of time that a browser should cache the response to the preflight OPTIONS request.
    maxAgeInSeconds *int32
    // Resource within the application segment for which CORS permissions are granted. / grants permission for the whole app segment.
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
// GetAllowedHeaders gets the allowedHeaders property value. The request headers that the origin domain may specify on the CORS request. The wildcard character * indicates that any header beginning with the specified prefix is allowed.
func (m *CorsConfiguration_v2) GetAllowedHeaders()([]string) {
    return m.allowedHeaders
}
// GetAllowedMethods gets the allowedMethods property value. The HTTP request methods that the origin domain may use for a CORS request.
func (m *CorsConfiguration_v2) GetAllowedMethods()([]string) {
    return m.allowedMethods
}
// GetAllowedOrigins gets the allowedOrigins property value. The origin domains that are permitted to make a request against the service via CORS. The origin domain is the domain from which the request originates. The origin must be an exact case-sensitive match with the origin that the user agent sends to the service.
func (m *CorsConfiguration_v2) GetAllowedOrigins()([]string) {
    return m.allowedOrigins
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CorsConfiguration_v2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedHeaders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedHeaders(res)
        }
        return nil
    }
    res["allowedMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedMethods(res)
        }
        return nil
    }
    res["allowedOrigins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedOrigins(res)
        }
        return nil
    }
    res["maxAgeInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxAgeInSeconds(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    return res
}
// GetMaxAgeInSeconds gets the maxAgeInSeconds property value. The maximum amount of time that a browser should cache the response to the preflight OPTIONS request.
func (m *CorsConfiguration_v2) GetMaxAgeInSeconds()(*int32) {
    return m.maxAgeInSeconds
}
// GetResource gets the resource property value. Resource within the application segment for which CORS permissions are granted. / grants permission for the whole app segment.
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
// SetAllowedHeaders sets the allowedHeaders property value. The request headers that the origin domain may specify on the CORS request. The wildcard character * indicates that any header beginning with the specified prefix is allowed.
func (m *CorsConfiguration_v2) SetAllowedHeaders(value []string)() {
    m.allowedHeaders = value
}
// SetAllowedMethods sets the allowedMethods property value. The HTTP request methods that the origin domain may use for a CORS request.
func (m *CorsConfiguration_v2) SetAllowedMethods(value []string)() {
    m.allowedMethods = value
}
// SetAllowedOrigins sets the allowedOrigins property value. The origin domains that are permitted to make a request against the service via CORS. The origin domain is the domain from which the request originates. The origin must be an exact case-sensitive match with the origin that the user agent sends to the service.
func (m *CorsConfiguration_v2) SetAllowedOrigins(value []string)() {
    m.allowedOrigins = value
}
// SetMaxAgeInSeconds sets the maxAgeInSeconds property value. The maximum amount of time that a browser should cache the response to the preflight OPTIONS request.
func (m *CorsConfiguration_v2) SetMaxAgeInSeconds(value *int32)() {
    m.maxAgeInSeconds = value
}
// SetResource sets the resource property value. Resource within the application segment for which CORS permissions are granted. / grants permission for the whole app segment.
func (m *CorsConfiguration_v2) SetResource(value *string)() {
    m.resource = value
}
