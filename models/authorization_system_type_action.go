package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthorizationSystemTypeAction 
type AuthorizationSystemTypeAction struct {
    Entity
}
// NewAuthorizationSystemTypeAction instantiates a new authorizationSystemTypeAction and sets the default values.
func NewAuthorizationSystemTypeAction()(*AuthorizationSystemTypeAction) {
    m := &AuthorizationSystemTypeAction{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthorizationSystemTypeActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthorizationSystemTypeActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.awsAuthorizationSystemTypeAction":
                        return NewAwsAuthorizationSystemTypeAction(), nil
                    case "#microsoft.graph.azureAuthorizationSystemTypeAction":
                        return NewAzureAuthorizationSystemTypeAction(), nil
                    case "#microsoft.graph.gcpAuthorizationSystemTypeAction":
                        return NewGcpAuthorizationSystemTypeAction(), nil
                }
            }
        }
    }
    return NewAuthorizationSystemTypeAction(), nil
}
// GetActionType gets the actionType property value. The type of action allowed in the authorization system's service. The possible values are: delete, read, unknownFutureValue. Supports $filter and (eq).
func (m *AuthorizationSystemTypeAction) GetActionType()(*AuthorizationSystemActionType) {
    val, err := m.GetBackingStore().Get("actionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuthorizationSystemActionType)
    }
    return nil
}
// GetExternalId gets the externalId property value. The display name of an action. Read-only. Supports $filter and (eq).
func (m *AuthorizationSystemTypeAction) GetExternalId()(*string) {
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
func (m *AuthorizationSystemTypeAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthorizationSystemActionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionType(val.(*AuthorizationSystemActionType))
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
    res["resourceTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetResourceTypes(res)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthorizationSystemActionSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*AuthorizationSystemActionSeverity))
        }
        return nil
    }
    return res
}
// GetResourceTypes gets the resourceTypes property value. The resource types in the authorization system's service where the action can be performed. Supports $filter and (eq).
func (m *AuthorizationSystemTypeAction) GetResourceTypes()([]string) {
    val, err := m.GetBackingStore().Get("resourceTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSeverity gets the severity property value. The severity property
func (m *AuthorizationSystemTypeAction) GetSeverity()(*AuthorizationSystemActionSeverity) {
    val, err := m.GetBackingStore().Get("severity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuthorizationSystemActionSeverity)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuthorizationSystemTypeAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionType() != nil {
        cast := (*m.GetActionType()).String()
        err = writer.WriteStringValue("actionType", &cast)
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
    if m.GetResourceTypes() != nil {
        err = writer.WriteCollectionOfStringValues("resourceTypes", m.GetResourceTypes())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err = writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionType sets the actionType property value. The type of action allowed in the authorization system's service. The possible values are: delete, read, unknownFutureValue. Supports $filter and (eq).
func (m *AuthorizationSystemTypeAction) SetActionType(value *AuthorizationSystemActionType)() {
    err := m.GetBackingStore().Set("actionType", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalId sets the externalId property value. The display name of an action. Read-only. Supports $filter and (eq).
func (m *AuthorizationSystemTypeAction) SetExternalId(value *string)() {
    err := m.GetBackingStore().Set("externalId", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceTypes sets the resourceTypes property value. The resource types in the authorization system's service where the action can be performed. Supports $filter and (eq).
func (m *AuthorizationSystemTypeAction) SetResourceTypes(value []string)() {
    err := m.GetBackingStore().Set("resourceTypes", value)
    if err != nil {
        panic(err)
    }
}
// SetSeverity sets the severity property value. The severity property
func (m *AuthorizationSystemTypeAction) SetSeverity(value *AuthorizationSystemActionSeverity)() {
    err := m.GetBackingStore().Set("severity", value)
    if err != nil {
        panic(err)
    }
}
// AuthorizationSystemTypeActionable 
type AuthorizationSystemTypeActionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionType()(*AuthorizationSystemActionType)
    GetExternalId()(*string)
    GetResourceTypes()([]string)
    GetSeverity()(*AuthorizationSystemActionSeverity)
    SetActionType(value *AuthorizationSystemActionType)()
    SetExternalId(value *string)()
    SetResourceTypes(value []string)()
    SetSeverity(value *AuthorizationSystemActionSeverity)()
}
