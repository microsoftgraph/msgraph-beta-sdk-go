package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// OfficeConfiguration 
type OfficeConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOfficeConfiguration instantiates a new OfficeConfiguration and sets the default values.
func NewOfficeConfiguration()(*OfficeConfiguration) {
    m := &OfficeConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOfficeConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOfficeConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OfficeConfiguration) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *OfficeConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetClientConfigurations gets the clientConfigurations property value. The clientConfigurations property
func (m *OfficeConfiguration) GetClientConfigurations()([]OfficeClientConfigurationable) {
    val, err := m.GetBackingStore().Get("clientConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OfficeClientConfigurationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOfficeClientConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfficeClientConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(OfficeClientConfigurationable)
            }
            m.SetClientConfigurations(res)
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
    res["tenantCheckinStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOfficeClientCheckinStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfficeClientCheckinStatusable, len(val))
            for i, v := range val {
                res[i] = v.(OfficeClientCheckinStatusable)
            }
            m.SetTenantCheckinStatuses(res)
        }
        return nil
    }
    res["tenantUserCheckinSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOfficeUserCheckinSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantUserCheckinSummary(val.(OfficeUserCheckinSummaryable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OfficeConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantCheckinStatuses gets the tenantCheckinStatuses property value. The tenantCheckinStatuses property
func (m *OfficeConfiguration) GetTenantCheckinStatuses()([]OfficeClientCheckinStatusable) {
    val, err := m.GetBackingStore().Get("tenantCheckinStatuses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OfficeClientCheckinStatusable)
    }
    return nil
}
// GetTenantUserCheckinSummary gets the tenantUserCheckinSummary property value. The tenantUserCheckinSummary property
func (m *OfficeConfiguration) GetTenantUserCheckinSummary()(OfficeUserCheckinSummaryable) {
    val, err := m.GetBackingStore().Get("tenantUserCheckinSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OfficeUserCheckinSummaryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OfficeConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClientConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClientConfigurations()))
        for i, v := range m.GetClientConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("clientConfigurations", cast)
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
    if m.GetTenantCheckinStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenantCheckinStatuses()))
        for i, v := range m.GetTenantCheckinStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("tenantCheckinStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("tenantUserCheckinSummary", m.GetTenantUserCheckinSummary())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OfficeConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *OfficeConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetClientConfigurations sets the clientConfigurations property value. The clientConfigurations property
func (m *OfficeConfiguration) SetClientConfigurations(value []OfficeClientConfigurationable)() {
    err := m.GetBackingStore().Set("clientConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OfficeConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantCheckinStatuses sets the tenantCheckinStatuses property value. The tenantCheckinStatuses property
func (m *OfficeConfiguration) SetTenantCheckinStatuses(value []OfficeClientCheckinStatusable)() {
    err := m.GetBackingStore().Set("tenantCheckinStatuses", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantUserCheckinSummary sets the tenantUserCheckinSummary property value. The tenantUserCheckinSummary property
func (m *OfficeConfiguration) SetTenantUserCheckinSummary(value OfficeUserCheckinSummaryable)() {
    err := m.GetBackingStore().Set("tenantUserCheckinSummary", value)
    if err != nil {
        panic(err)
    }
}
// OfficeConfigurationable 
type OfficeConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetClientConfigurations()([]OfficeClientConfigurationable)
    GetOdataType()(*string)
    GetTenantCheckinStatuses()([]OfficeClientCheckinStatusable)
    GetTenantUserCheckinSummary()(OfficeUserCheckinSummaryable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetClientConfigurations(value []OfficeClientConfigurationable)()
    SetOdataType(value *string)()
    SetTenantCheckinStatuses(value []OfficeClientCheckinStatusable)()
    SetTenantUserCheckinSummary(value OfficeUserCheckinSummaryable)()
}
