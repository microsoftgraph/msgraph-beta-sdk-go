package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SamlOrWsFedExternalDomainFederation 
type SamlOrWsFedExternalDomainFederation struct {
    SamlOrWsFedProvider
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Collection of domain names of the external organizations that the tenant is federating with. Supports $filter (eq).
    domains []ExternalDomainNameable
}
// NewSamlOrWsFedExternalDomainFederation instantiates a new SamlOrWsFedExternalDomainFederation and sets the default values.
func NewSamlOrWsFedExternalDomainFederation()(*SamlOrWsFedExternalDomainFederation) {
    m := &SamlOrWsFedExternalDomainFederation{
        SamlOrWsFedProvider: *NewSamlOrWsFedProvider(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSamlOrWsFedExternalDomainFederationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSamlOrWsFedExternalDomainFederationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSamlOrWsFedExternalDomainFederation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SamlOrWsFedExternalDomainFederation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDomains gets the domains property value. Collection of domain names of the external organizations that the tenant is federating with. Supports $filter (eq).
func (m *SamlOrWsFedExternalDomainFederation) GetDomains()([]ExternalDomainNameable) {
    if m == nil {
        return nil
    } else {
        return m.domains
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SamlOrWsFedExternalDomainFederation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SamlOrWsFedProvider.GetFieldDeserializers()
    res["domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExternalDomainNameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalDomainNameable, len(val))
            for i, v := range val {
                res[i] = v.(ExternalDomainNameable)
            }
            m.SetDomains(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *SamlOrWsFedExternalDomainFederation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SamlOrWsFedProvider.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDomains() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDomains()))
        for i, v := range m.GetDomains() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("domains", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SamlOrWsFedExternalDomainFederation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDomains sets the domains property value. Collection of domain names of the external organizations that the tenant is federating with. Supports $filter (eq).
func (m *SamlOrWsFedExternalDomainFederation) SetDomains(value []ExternalDomainNameable)() {
    if m != nil {
        m.domains = value
    }
}
