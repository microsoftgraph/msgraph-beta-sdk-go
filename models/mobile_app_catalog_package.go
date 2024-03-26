package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppCatalogPackage mobileAppCatalogPackage is an abstract type that application catalog package entities derive from. A mobileAppCatalogPackage entity contains information about an application catalog package that can be deployed to Intune-managed devices.
type MobileAppCatalogPackage struct {
    Entity
}
// NewMobileAppCatalogPackage instantiates a new MobileAppCatalogPackage and sets the default values.
func NewMobileAppCatalogPackage()(*MobileAppCatalogPackage) {
    m := &MobileAppCatalogPackage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppCatalogPackageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMobileAppCatalogPackageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.win32MobileAppCatalogPackage":
                        return NewWin32MobileAppCatalogPackage(), nil
                }
            }
        }
    }
    return NewMobileAppCatalogPackage(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MobileAppCatalogPackage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["productDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductDisplayName(val)
        }
        return nil
    }
    res["productId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductId(val)
        }
        return nil
    }
    res["publisherDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherDisplayName(val)
        }
        return nil
    }
    res["versionDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionDisplayName(val)
        }
        return nil
    }
    return res
}
// GetProductDisplayName gets the productDisplayName property value. The name of the product (example: "Fabrikam for Business"). Returned by default. Read-only. Supports: $filter, $search, $select. This property is read-only.
// returns a *string when successful
func (m *MobileAppCatalogPackage) GetProductDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("productDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProductId gets the productId property value. The identifier of a specific product irrespective of version, or other attributes. Read-only. Returned by default. Supports: $filter, $select. This property is read-only.
// returns a *string when successful
func (m *MobileAppCatalogPackage) GetProductId()(*string) {
    val, err := m.GetBackingStore().Get("productId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublisherDisplayName gets the publisherDisplayName property value. The name of the application catalog package publisher (example: "Fabrikam"). Returned by default. Read-only. Supports $filter, $search, $select. This property is read-only.
// returns a *string when successful
func (m *MobileAppCatalogPackage) GetPublisherDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("publisherDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVersionDisplayName gets the versionDisplayName property value. The name of the product version (example: "1.2203.156"). Returned by default. Read-only. Supports: $filter, $search, $select. This property is read-only.
// returns a *string when successful
func (m *MobileAppCatalogPackage) GetVersionDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("versionDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MobileAppCatalogPackage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SetProductDisplayName sets the productDisplayName property value. The name of the product (example: "Fabrikam for Business"). Returned by default. Read-only. Supports: $filter, $search, $select. This property is read-only.
func (m *MobileAppCatalogPackage) SetProductDisplayName(value *string)() {
    err := m.GetBackingStore().Set("productDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetProductId sets the productId property value. The identifier of a specific product irrespective of version, or other attributes. Read-only. Returned by default. Supports: $filter, $select. This property is read-only.
func (m *MobileAppCatalogPackage) SetProductId(value *string)() {
    err := m.GetBackingStore().Set("productId", value)
    if err != nil {
        panic(err)
    }
}
// SetPublisherDisplayName sets the publisherDisplayName property value. The name of the application catalog package publisher (example: "Fabrikam"). Returned by default. Read-only. Supports $filter, $search, $select. This property is read-only.
func (m *MobileAppCatalogPackage) SetPublisherDisplayName(value *string)() {
    err := m.GetBackingStore().Set("publisherDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetVersionDisplayName sets the versionDisplayName property value. The name of the product version (example: "1.2203.156"). Returned by default. Read-only. Supports: $filter, $search, $select. This property is read-only.
func (m *MobileAppCatalogPackage) SetVersionDisplayName(value *string)() {
    err := m.GetBackingStore().Set("versionDisplayName", value)
    if err != nil {
        panic(err)
    }
}
type MobileAppCatalogPackageable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetProductDisplayName()(*string)
    GetProductId()(*string)
    GetPublisherDisplayName()(*string)
    GetVersionDisplayName()(*string)
    SetProductDisplayName(value *string)()
    SetProductId(value *string)()
    SetPublisherDisplayName(value *string)()
    SetVersionDisplayName(value *string)()
}
