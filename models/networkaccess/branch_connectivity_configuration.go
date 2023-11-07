package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// BranchConnectivityConfiguration 
type BranchConnectivityConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewBranchConnectivityConfiguration instantiates a new branchConnectivityConfiguration and sets the default values.
func NewBranchConnectivityConfiguration()(*BranchConnectivityConfiguration) {
    m := &BranchConnectivityConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBranchConnectivityConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBranchConnectivityConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBranchConnectivityConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BranchConnectivityConfiguration) GetAdditionalData()(map[string]any) {
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
func (m *BranchConnectivityConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBranchId gets the branchId property value. Unique identifier or a specific reference assigned to a branchSite. Key.
func (m *BranchConnectivityConfiguration) GetBranchId()(*string) {
    val, err := m.GetBackingStore().Get("branchId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBranchName gets the branchName property value. Display name assigned to a branchSite.
func (m *BranchConnectivityConfiguration) GetBranchName()(*string) {
    val, err := m.GetBackingStore().Get("branchName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BranchConnectivityConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["branchId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranchId(val)
        }
        return nil
    }
    res["branchName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranchName(val)
        }
        return nil
    }
    res["links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectivityConfigurationLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectivityConfigurationLinkable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConnectivityConfigurationLinkable)
                }
            }
            m.SetLinks(res)
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
// GetLinks gets the links property value. List of connectivity configurations for deviceLink objects.
func (m *BranchConnectivityConfiguration) GetLinks()([]ConnectivityConfigurationLinkable) {
    val, err := m.GetBackingStore().Get("links")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ConnectivityConfigurationLinkable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BranchConnectivityConfiguration) GetOdataType()(*string) {
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
func (m *BranchConnectivityConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("branchId", m.GetBranchId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("branchName", m.GetBranchName())
        if err != nil {
            return err
        }
    }
    if m.GetLinks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLinks()))
        for i, v := range m.GetLinks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("links", cast)
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
func (m *BranchConnectivityConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *BranchConnectivityConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBranchId sets the branchId property value. Unique identifier or a specific reference assigned to a branchSite. Key.
func (m *BranchConnectivityConfiguration) SetBranchId(value *string)() {
    err := m.GetBackingStore().Set("branchId", value)
    if err != nil {
        panic(err)
    }
}
// SetBranchName sets the branchName property value. Display name assigned to a branchSite.
func (m *BranchConnectivityConfiguration) SetBranchName(value *string)() {
    err := m.GetBackingStore().Set("branchName", value)
    if err != nil {
        panic(err)
    }
}
// SetLinks sets the links property value. List of connectivity configurations for deviceLink objects.
func (m *BranchConnectivityConfiguration) SetLinks(value []ConnectivityConfigurationLinkable)() {
    err := m.GetBackingStore().Set("links", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BranchConnectivityConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// BranchConnectivityConfigurationable 
type BranchConnectivityConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBranchId()(*string)
    GetBranchName()(*string)
    GetLinks()([]ConnectivityConfigurationLinkable)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBranchId(value *string)()
    SetBranchName(value *string)()
    SetLinks(value []ConnectivityConfigurationLinkable)()
    SetOdataType(value *string)()
}
