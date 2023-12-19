package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TenantAllowBlockListEntryResult 
type TenantAllowBlockListEntryResult struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTenantAllowBlockListEntryResult instantiates a new tenantAllowBlockListEntryResult and sets the default values.
func NewTenantAllowBlockListEntryResult()(*TenantAllowBlockListEntryResult) {
    m := &TenantAllowBlockListEntryResult{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTenantAllowBlockListEntryResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantAllowBlockListEntryResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantAllowBlockListEntryResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantAllowBlockListEntryResult) GetAdditionalData()(map[string]any) {
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
func (m *TenantAllowBlockListEntryResult) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEntryType gets the entryType property value. The tenant allow-block list entry type. The possible values are: url, fileHash, sender, recipient and unkownFutureValue.
func (m *TenantAllowBlockListEntryResult) GetEntryType()(*TenantAllowBlockListEntryResult_entryType) {
    val, err := m.GetBackingStore().Get("entryType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TenantAllowBlockListEntryResult_entryType)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. Specifies the date and time when the entry expires.
func (m *TenantAllowBlockListEntryResult) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantAllowBlockListEntryResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["entryType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTenantAllowBlockListEntryResult_entryType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntryType(val.(*TenantAllowBlockListEntryResult_entryType))
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["identity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTenantAllowBlockListEntryResult_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*TenantAllowBlockListEntryResult_status))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. Specifies the identity of the entry generated by the tenant allow block list system.
func (m *TenantAllowBlockListEntryResult) GetIdentity()(*string) {
    val, err := m.GetBackingStore().Get("identity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TenantAllowBlockListEntryResult) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. Specifies whether the tenant allow block list entry creation operation was successful. The possible values are: notStarted, running, succeeded, failed, skipped and unkownFutureValue.
func (m *TenantAllowBlockListEntryResult) GetStatus()(*TenantAllowBlockListEntryResult_status) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TenantAllowBlockListEntryResult_status)
    }
    return nil
}
// GetValue gets the value property value. Specifies the value of the created tenant allow block list entry.
func (m *TenantAllowBlockListEntryResult) GetValue()(*string) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TenantAllowBlockListEntryResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEntryType() != nil {
        cast := (*m.GetEntryType()).String()
        err := writer.WriteStringValue("entryType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identity", m.GetIdentity())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *TenantAllowBlockListEntryResult) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TenantAllowBlockListEntryResult) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEntryType sets the entryType property value. The tenant allow-block list entry type. The possible values are: url, fileHash, sender, recipient and unkownFutureValue.
func (m *TenantAllowBlockListEntryResult) SetEntryType(value *TenantAllowBlockListEntryResult_entryType)() {
    err := m.GetBackingStore().Set("entryType", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. Specifies the date and time when the entry expires.
func (m *TenantAllowBlockListEntryResult) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentity sets the identity property value. Specifies the identity of the entry generated by the tenant allow block list system.
func (m *TenantAllowBlockListEntryResult) SetIdentity(value *string)() {
    err := m.GetBackingStore().Set("identity", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TenantAllowBlockListEntryResult) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Specifies whether the tenant allow block list entry creation operation was successful. The possible values are: notStarted, running, succeeded, failed, skipped and unkownFutureValue.
func (m *TenantAllowBlockListEntryResult) SetStatus(value *TenantAllowBlockListEntryResult_status)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetValue sets the value property value. Specifies the value of the created tenant allow block list entry.
func (m *TenantAllowBlockListEntryResult) SetValue(value *string)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// TenantAllowBlockListEntryResultable 
type TenantAllowBlockListEntryResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEntryType()(*TenantAllowBlockListEntryResult_entryType)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIdentity()(*string)
    GetOdataType()(*string)
    GetStatus()(*TenantAllowBlockListEntryResult_status)
    GetValue()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEntryType(value *TenantAllowBlockListEntryResult_entryType)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIdentity(value *string)()
    SetOdataType(value *string)()
    SetStatus(value *TenantAllowBlockListEntryResult_status)()
    SetValue(value *string)()
}
