package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVppAppAssignedLicense iOS Volume Purchase Program license assignment. This class does not support Create, Delete, or Update.
type IosVppAppAssignedLicense struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewIosVppAppAssignedLicense instantiates a new iosVppAppAssignedLicense and sets the default values.
func NewIosVppAppAssignedLicense()(*IosVppAppAssignedLicense) {
    m := &IosVppAppAssignedLicense{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIosVppAppAssignedLicenseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosVppAppAssignedLicenseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.iosVppAppAssignedDeviceLicense":
                        return NewIosVppAppAssignedDeviceLicense(), nil
                    case "#microsoft.graph.iosVppAppAssignedUserLicense":
                        return NewIosVppAppAssignedUserLicense(), nil
                }
            }
        }
    }
    return NewIosVppAppAssignedLicense(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosVppAppAssignedLicense) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["userEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEmailAddress(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetUserEmailAddress gets the userEmailAddress property value. The user email address.
func (m *IosVppAppAssignedLicense) GetUserEmailAddress()(*string) {
    val, err := m.GetBackingStore().Get("userEmailAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. The user ID.
func (m *IosVppAppAssignedLicense) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserName gets the userName property value. The user name.
func (m *IosVppAppAssignedLicense) GetUserName()(*string) {
    val, err := m.GetBackingStore().Get("userName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name.
func (m *IosVppAppAssignedLicense) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosVppAppAssignedLicense) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userEmailAddress", m.GetUserEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserEmailAddress sets the userEmailAddress property value. The user email address.
func (m *IosVppAppAssignedLicense) SetUserEmailAddress(value *string)() {
    err := m.GetBackingStore().Set("userEmailAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The user ID.
func (m *IosVppAppAssignedLicense) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserName sets the userName property value. The user name.
func (m *IosVppAppAssignedLicense) SetUserName(value *string)() {
    err := m.GetBackingStore().Set("userName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name.
func (m *IosVppAppAssignedLicense) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// IosVppAppAssignedLicenseable 
type IosVppAppAssignedLicenseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUserEmailAddress()(*string)
    GetUserId()(*string)
    GetUserName()(*string)
    GetUserPrincipalName()(*string)
    SetUserEmailAddress(value *string)()
    SetUserId(value *string)()
    SetUserName(value *string)()
    SetUserPrincipalName(value *string)()
}
