package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthorizationSystemResource struct {
    Entity
}
// NewAuthorizationSystemResource instantiates a new AuthorizationSystemResource and sets the default values.
func NewAuthorizationSystemResource()(*AuthorizationSystemResource) {
    m := &AuthorizationSystemResource{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthorizationSystemResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthorizationSystemResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.awsAuthorizationSystemResource":
                        return NewAwsAuthorizationSystemResource(), nil
                    case "#microsoft.graph.azureAuthorizationSystemResource":
                        return NewAzureAuthorizationSystemResource(), nil
                    case "#microsoft.graph.gcpAuthorizationSystemResource":
                        return NewGcpAuthorizationSystemResource(), nil
                }
            }
        }
    }
    return NewAuthorizationSystemResource(), nil
}
// GetAuthorizationSystem gets the authorizationSystem property value. The authorization system that the resource exists in.
// returns a AuthorizationSystemable when successful
func (m *AuthorizationSystemResource) GetAuthorizationSystem()(AuthorizationSystemable) {
    val, err := m.GetBackingStore().Get("authorizationSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The name of the resource. Read-only. Supports $filter (eq,contains).
// returns a *string when successful
func (m *AuthorizationSystemResource) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalId gets the externalId property value. The ID of the resource as defined by the authorization system provider. Read-only. Supports $filter (eq).
// returns a *string when successful
func (m *AuthorizationSystemResource) GetExternalId()(*string) {
    val, err := m.GetBackingStore().Get("externalId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthorizationSystemResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authorizationSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizationSystem(val.(AuthorizationSystemable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["resourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceType(val)
        }
        return nil
    }
    return res
}
// GetResourceType gets the resourceType property value. The type of the resource. Read-only. Supports $filter (eq).
// returns a *string when successful
func (m *AuthorizationSystemResource) GetResourceType()(*string) {
    val, err := m.GetBackingStore().Get("resourceType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuthorizationSystemResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("authorizationSystem", m.GetAuthorizationSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceType", m.GetResourceType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthorizationSystem sets the authorizationSystem property value. The authorization system that the resource exists in.
func (m *AuthorizationSystemResource) SetAuthorizationSystem(value AuthorizationSystemable)() {
    err := m.GetBackingStore().Set("authorizationSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name of the resource. Read-only. Supports $filter (eq,contains).
func (m *AuthorizationSystemResource) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalId sets the externalId property value. The ID of the resource as defined by the authorization system provider. Read-only. Supports $filter (eq).
func (m *AuthorizationSystemResource) SetExternalId(value *string)() {
    err := m.GetBackingStore().Set("externalId", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceType sets the resourceType property value. The type of the resource. Read-only. Supports $filter (eq).
func (m *AuthorizationSystemResource) SetResourceType(value *string)() {
    err := m.GetBackingStore().Set("resourceType", value)
    if err != nil {
        panic(err)
    }
}
type AuthorizationSystemResourceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizationSystem()(AuthorizationSystemable)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetResourceType()(*string)
    SetAuthorizationSystem(value AuthorizationSystemable)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetResourceType(value *string)()
}
