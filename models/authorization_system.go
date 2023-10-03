package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthorizationSystem 
type AuthorizationSystem struct {
    Entity
}
// NewAuthorizationSystem instantiates a new authorizationSystem and sets the default values.
func NewAuthorizationSystem()(*AuthorizationSystem) {
    m := &AuthorizationSystem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthorizationSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthorizationSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.awsAuthorizationSystem":
                        return NewAwsAuthorizationSystem(), nil
                    case "#microsoft.graph.azureAuthorizationSystem":
                        return NewAzureAuthorizationSystem(), nil
                    case "#microsoft.graph.gcpAuthorizationSystem":
                        return NewGcpAuthorizationSystem(), nil
                }
            }
        }
    }
    return NewAuthorizationSystem(), nil
}
// GetAuthorizationSystemId gets the authorizationSystemId property value. The authorizationSystemId property
func (m *AuthorizationSystem) GetAuthorizationSystemId()(*string) {
    val, err := m.GetBackingStore().Get("authorizationSystemId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAuthorizationSystemName gets the authorizationSystemName property value. The authorizationSystemName property
func (m *AuthorizationSystem) GetAuthorizationSystemName()(*string) {
    val, err := m.GetBackingStore().Get("authorizationSystemName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAuthorizationSystemType gets the authorizationSystemType property value. The authorizationSystemType property
func (m *AuthorizationSystem) GetAuthorizationSystemType()(*string) {
    val, err := m.GetBackingStore().Get("authorizationSystemType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDataCollectionInfo gets the dataCollectionInfo property value. The dataCollectionInfo property
func (m *AuthorizationSystem) GetDataCollectionInfo()(DataCollectionInfoable) {
    val, err := m.GetBackingStore().Get("dataCollectionInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DataCollectionInfoable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthorizationSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authorizationSystemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizationSystemId(val)
        }
        return nil
    }
    res["authorizationSystemName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizationSystemName(val)
        }
        return nil
    }
    res["authorizationSystemType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizationSystemType(val)
        }
        return nil
    }
    res["dataCollectionInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDataCollectionInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataCollectionInfo(val.(DataCollectionInfoable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AuthorizationSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("authorizationSystemId", m.GetAuthorizationSystemId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("authorizationSystemName", m.GetAuthorizationSystemName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("authorizationSystemType", m.GetAuthorizationSystemType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dataCollectionInfo", m.GetDataCollectionInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthorizationSystemId sets the authorizationSystemId property value. The authorizationSystemId property
func (m *AuthorizationSystem) SetAuthorizationSystemId(value *string)() {
    err := m.GetBackingStore().Set("authorizationSystemId", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthorizationSystemName sets the authorizationSystemName property value. The authorizationSystemName property
func (m *AuthorizationSystem) SetAuthorizationSystemName(value *string)() {
    err := m.GetBackingStore().Set("authorizationSystemName", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthorizationSystemType sets the authorizationSystemType property value. The authorizationSystemType property
func (m *AuthorizationSystem) SetAuthorizationSystemType(value *string)() {
    err := m.GetBackingStore().Set("authorizationSystemType", value)
    if err != nil {
        panic(err)
    }
}
// SetDataCollectionInfo sets the dataCollectionInfo property value. The dataCollectionInfo property
func (m *AuthorizationSystem) SetDataCollectionInfo(value DataCollectionInfoable)() {
    err := m.GetBackingStore().Set("dataCollectionInfo", value)
    if err != nil {
        panic(err)
    }
}
// AuthorizationSystemable 
type AuthorizationSystemable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizationSystemId()(*string)
    GetAuthorizationSystemName()(*string)
    GetAuthorizationSystemType()(*string)
    GetDataCollectionInfo()(DataCollectionInfoable)
    SetAuthorizationSystemId(value *string)()
    SetAuthorizationSystemName(value *string)()
    SetAuthorizationSystemType(value *string)()
    SetDataCollectionInfo(value DataCollectionInfoable)()
}
