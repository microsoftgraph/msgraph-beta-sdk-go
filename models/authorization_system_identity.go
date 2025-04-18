// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthorizationSystemIdentity struct {
    Entity
}
// NewAuthorizationSystemIdentity instantiates a new AuthorizationSystemIdentity and sets the default values.
func NewAuthorizationSystemIdentity()(*AuthorizationSystemIdentity) {
    m := &AuthorizationSystemIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthorizationSystemIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthorizationSystemIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.awsAccessKey":
                        return NewAwsAccessKey(), nil
                    case "#microsoft.graph.awsEc2Instance":
                        return NewAwsEc2Instance(), nil
                    case "#microsoft.graph.awsGroup":
                        return NewAwsGroup(), nil
                    case "#microsoft.graph.awsIdentity":
                        return NewAwsIdentity(), nil
                    case "#microsoft.graph.awsLambda":
                        return NewAwsLambda(), nil
                    case "#microsoft.graph.awsRole":
                        return NewAwsRole(), nil
                    case "#microsoft.graph.awsUser":
                        return NewAwsUser(), nil
                    case "#microsoft.graph.azureGroup":
                        return NewAzureGroup(), nil
                    case "#microsoft.graph.azureIdentity":
                        return NewAzureIdentity(), nil
                    case "#microsoft.graph.azureManagedIdentity":
                        return NewAzureManagedIdentity(), nil
                    case "#microsoft.graph.azureServerlessFunction":
                        return NewAzureServerlessFunction(), nil
                    case "#microsoft.graph.azureServicePrincipal":
                        return NewAzureServicePrincipal(), nil
                    case "#microsoft.graph.azureUser":
                        return NewAzureUser(), nil
                    case "#microsoft.graph.gcpCloudFunction":
                        return NewGcpCloudFunction(), nil
                    case "#microsoft.graph.gcpGroup":
                        return NewGcpGroup(), nil
                    case "#microsoft.graph.gcpIdentity":
                        return NewGcpIdentity(), nil
                    case "#microsoft.graph.gcpServiceAccount":
                        return NewGcpServiceAccount(), nil
                    case "#microsoft.graph.gcpUser":
                        return NewGcpUser(), nil
                }
            }
        }
    }
    return NewAuthorizationSystemIdentity(), nil
}
// GetAuthorizationSystem gets the authorizationSystem property value. Navigation to the authorizationSystem object
// returns a AuthorizationSystemable when successful
func (m *AuthorizationSystemIdentity) GetAuthorizationSystem()(AuthorizationSystemable) {
    val, err := m.GetBackingStore().Get("authorizationSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The name of the identity. Read-only. Supports $filter and (eq,contains).
// returns a *string when successful
func (m *AuthorizationSystemIdentity) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalId gets the externalId property value. Unique ID of the identity within the external system. Read-only.
// returns a *string when successful
func (m *AuthorizationSystemIdentity) GetExternalId()(*string) {
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
func (m *AuthorizationSystemIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationSystemIdentitySourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(AuthorizationSystemIdentitySourceable))
        }
        return nil
    }
    return res
}
// GetSource gets the source property value. Represents details of the source of the identity.
// returns a AuthorizationSystemIdentitySourceable when successful
func (m *AuthorizationSystemIdentity) GetSource()(AuthorizationSystemIdentitySourceable) {
    val, err := m.GetBackingStore().Get("source")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemIdentitySourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuthorizationSystemIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthorizationSystem sets the authorizationSystem property value. Navigation to the authorizationSystem object
func (m *AuthorizationSystemIdentity) SetAuthorizationSystem(value AuthorizationSystemable)() {
    err := m.GetBackingStore().Set("authorizationSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name of the identity. Read-only. Supports $filter and (eq,contains).
func (m *AuthorizationSystemIdentity) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalId sets the externalId property value. Unique ID of the identity within the external system. Read-only.
func (m *AuthorizationSystemIdentity) SetExternalId(value *string)() {
    err := m.GetBackingStore().Set("externalId", value)
    if err != nil {
        panic(err)
    }
}
// SetSource sets the source property value. Represents details of the source of the identity.
func (m *AuthorizationSystemIdentity) SetSource(value AuthorizationSystemIdentitySourceable)() {
    err := m.GetBackingStore().Set("source", value)
    if err != nil {
        panic(err)
    }
}
type AuthorizationSystemIdentityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizationSystem()(AuthorizationSystemable)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetSource()(AuthorizationSystemIdentitySourceable)
    SetAuthorizationSystem(value AuthorizationSystemable)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetSource(value AuthorizationSystemIdentitySourceable)()
}
