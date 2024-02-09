package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type RoleManagement struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewRoleManagement instantiates a new RoleManagement and sets the default values.
func NewRoleManagement()(*RoleManagement) {
    m := &RoleManagement{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRoleManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRoleManagementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleManagement(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RoleManagement) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *RoleManagement) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCloudPC gets the cloudPC property value. The cloudPC property
// returns a RbacApplicationMultipleable when successful
func (m *RoleManagement) GetCloudPC()(RbacApplicationMultipleable) {
    val, err := m.GetBackingStore().Get("cloudPC")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RbacApplicationMultipleable)
    }
    return nil
}
// GetDeviceManagement gets the deviceManagement property value. The RbacApplication for Device Management
// returns a RbacApplicationMultipleable when successful
func (m *RoleManagement) GetDeviceManagement()(RbacApplicationMultipleable) {
    val, err := m.GetBackingStore().Get("deviceManagement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RbacApplicationMultipleable)
    }
    return nil
}
// GetDirectory gets the directory property value. The directory property
// returns a RbacApplicationable when successful
func (m *RoleManagement) GetDirectory()(RbacApplicationable) {
    val, err := m.GetBackingStore().Get("directory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RbacApplicationable)
    }
    return nil
}
// GetEnterpriseApps gets the enterpriseApps property value. The enterpriseApps property
// returns a []RbacApplicationable when successful
func (m *RoleManagement) GetEnterpriseApps()([]RbacApplicationable) {
    val, err := m.GetBackingStore().Get("enterpriseApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RbacApplicationable)
    }
    return nil
}
// GetEntitlementManagement gets the entitlementManagement property value. The RbacApplication for Entitlement Management
// returns a RbacApplicationable when successful
func (m *RoleManagement) GetEntitlementManagement()(RbacApplicationable) {
    val, err := m.GetBackingStore().Get("entitlementManagement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RbacApplicationable)
    }
    return nil
}
// GetExchange gets the exchange property value. The exchange property
// returns a UnifiedRbacApplicationable when successful
func (m *RoleManagement) GetExchange()(UnifiedRbacApplicationable) {
    val, err := m.GetBackingStore().Get("exchange")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UnifiedRbacApplicationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RoleManagement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cloudPC"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRbacApplicationMultipleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPC(val.(RbacApplicationMultipleable))
        }
        return nil
    }
    res["deviceManagement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRbacApplicationMultipleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManagement(val.(RbacApplicationMultipleable))
        }
        return nil
    }
    res["directory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRbacApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectory(val.(RbacApplicationable))
        }
        return nil
    }
    res["enterpriseApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRbacApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RbacApplicationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RbacApplicationable)
                }
            }
            m.SetEnterpriseApps(res)
        }
        return nil
    }
    res["entitlementManagement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRbacApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntitlementManagement(val.(RbacApplicationable))
        }
        return nil
    }
    res["exchange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRbacApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchange(val.(UnifiedRbacApplicationable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *RoleManagement) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RoleManagement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cloudPC", m.GetCloudPC())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceManagement", m.GetDeviceManagement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("directory", m.GetDirectory())
        if err != nil {
            return err
        }
    }
    if m.GetEnterpriseApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnterpriseApps()))
        for i, v := range m.GetEnterpriseApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("enterpriseApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("entitlementManagement", m.GetEntitlementManagement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("exchange", m.GetExchange())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleManagement) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *RoleManagement) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCloudPC sets the cloudPC property value. The cloudPC property
func (m *RoleManagement) SetCloudPC(value RbacApplicationMultipleable)() {
    err := m.GetBackingStore().Set("cloudPC", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceManagement sets the deviceManagement property value. The RbacApplication for Device Management
func (m *RoleManagement) SetDeviceManagement(value RbacApplicationMultipleable)() {
    err := m.GetBackingStore().Set("deviceManagement", value)
    if err != nil {
        panic(err)
    }
}
// SetDirectory sets the directory property value. The directory property
func (m *RoleManagement) SetDirectory(value RbacApplicationable)() {
    err := m.GetBackingStore().Set("directory", value)
    if err != nil {
        panic(err)
    }
}
// SetEnterpriseApps sets the enterpriseApps property value. The enterpriseApps property
func (m *RoleManagement) SetEnterpriseApps(value []RbacApplicationable)() {
    err := m.GetBackingStore().Set("enterpriseApps", value)
    if err != nil {
        panic(err)
    }
}
// SetEntitlementManagement sets the entitlementManagement property value. The RbacApplication for Entitlement Management
func (m *RoleManagement) SetEntitlementManagement(value RbacApplicationable)() {
    err := m.GetBackingStore().Set("entitlementManagement", value)
    if err != nil {
        panic(err)
    }
}
// SetExchange sets the exchange property value. The exchange property
func (m *RoleManagement) SetExchange(value UnifiedRbacApplicationable)() {
    err := m.GetBackingStore().Set("exchange", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RoleManagement) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type RoleManagementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCloudPC()(RbacApplicationMultipleable)
    GetDeviceManagement()(RbacApplicationMultipleable)
    GetDirectory()(RbacApplicationable)
    GetEnterpriseApps()([]RbacApplicationable)
    GetEntitlementManagement()(RbacApplicationable)
    GetExchange()(UnifiedRbacApplicationable)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCloudPC(value RbacApplicationMultipleable)()
    SetDeviceManagement(value RbacApplicationMultipleable)()
    SetDirectory(value RbacApplicationable)()
    SetEnterpriseApps(value []RbacApplicationable)()
    SetEntitlementManagement(value RbacApplicationable)()
    SetExchange(value UnifiedRbacApplicationable)()
    SetOdataType(value *string)()
}
