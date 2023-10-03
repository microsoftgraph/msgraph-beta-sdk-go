package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureRoleDefinition 
type AzureRoleDefinition struct {
    Entity
}
// NewAzureRoleDefinition instantiates a new azureRoleDefinition and sets the default values.
func NewAzureRoleDefinition()(*AzureRoleDefinition) {
    m := &AzureRoleDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAzureRoleDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureRoleDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureRoleDefinition(), nil
}
// GetAssignableScopes gets the assignableScopes property value. The assignableScopes property
func (m *AzureRoleDefinition) GetAssignableScopes()([]string) {
    val, err := m.GetBackingStore().Get("assignableScopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAzureRoleDefinitionType gets the azureRoleDefinitionType property value. The azureRoleDefinitionType property
func (m *AzureRoleDefinition) GetAzureRoleDefinitionType()(*AzureRoleDefinitionType) {
    val, err := m.GetBackingStore().Get("azureRoleDefinitionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AzureRoleDefinitionType)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *AzureRoleDefinition) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalId gets the externalId property value. The externalId property
func (m *AzureRoleDefinition) GetExternalId()(*string) {
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
func (m *AzureRoleDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignableScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAssignableScopes(res)
        }
        return nil
    }
    res["azureRoleDefinitionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAzureRoleDefinitionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureRoleDefinitionType(val.(*AzureRoleDefinitionType))
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
    return res
}
// Serialize serializes information the current object
func (m *AzureRoleDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignableScopes() != nil {
        err = writer.WriteCollectionOfStringValues("assignableScopes", m.GetAssignableScopes())
        if err != nil {
            return err
        }
    }
    if m.GetAzureRoleDefinitionType() != nil {
        cast := (*m.GetAzureRoleDefinitionType()).String()
        err = writer.WriteStringValue("azureRoleDefinitionType", &cast)
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
    return nil
}
// SetAssignableScopes sets the assignableScopes property value. The assignableScopes property
func (m *AzureRoleDefinition) SetAssignableScopes(value []string)() {
    err := m.GetBackingStore().Set("assignableScopes", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureRoleDefinitionType sets the azureRoleDefinitionType property value. The azureRoleDefinitionType property
func (m *AzureRoleDefinition) SetAzureRoleDefinitionType(value *AzureRoleDefinitionType)() {
    err := m.GetBackingStore().Set("azureRoleDefinitionType", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *AzureRoleDefinition) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalId sets the externalId property value. The externalId property
func (m *AzureRoleDefinition) SetExternalId(value *string)() {
    err := m.GetBackingStore().Set("externalId", value)
    if err != nil {
        panic(err)
    }
}
// AzureRoleDefinitionable 
type AzureRoleDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignableScopes()([]string)
    GetAzureRoleDefinitionType()(*AzureRoleDefinitionType)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    SetAssignableScopes(value []string)()
    SetAzureRoleDefinitionType(value *AzureRoleDefinitionType)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
}
