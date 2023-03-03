package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// RoleSuccessStatistics 
type RoleSuccessStatistics struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewRoleSuccessStatistics instantiates a new roleSuccessStatistics and sets the default values.
func NewRoleSuccessStatistics()(*RoleSuccessStatistics) {
    m := &RoleSuccessStatistics{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRoleSuccessStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleSuccessStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleSuccessStatistics(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleSuccessStatistics) GetAdditionalData()(map[string]any) {
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
func (m *RoleSuccessStatistics) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleSuccessStatistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["permanentFail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermanentFail(val)
        }
        return nil
    }
    res["permanentSuccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermanentSuccess(val)
        }
        return nil
    }
    res["removeFail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveFail(val)
        }
        return nil
    }
    res["removeSuccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveSuccess(val)
        }
        return nil
    }
    res["roleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleId(val)
        }
        return nil
    }
    res["roleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleName(val)
        }
        return nil
    }
    res["temporaryFail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemporaryFail(val)
        }
        return nil
    }
    res["temporarySuccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemporarySuccess(val)
        }
        return nil
    }
    res["unknownFail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownFail(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RoleSuccessStatistics) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPermanentFail gets the permanentFail property value. The permanentFail property
func (m *RoleSuccessStatistics) GetPermanentFail()(*int64) {
    val, err := m.GetBackingStore().Get("permanentFail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetPermanentSuccess gets the permanentSuccess property value. The permanentSuccess property
func (m *RoleSuccessStatistics) GetPermanentSuccess()(*int64) {
    val, err := m.GetBackingStore().Get("permanentSuccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetRemoveFail gets the removeFail property value. The removeFail property
func (m *RoleSuccessStatistics) GetRemoveFail()(*int64) {
    val, err := m.GetBackingStore().Get("removeFail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetRemoveSuccess gets the removeSuccess property value. The removeSuccess property
func (m *RoleSuccessStatistics) GetRemoveSuccess()(*int64) {
    val, err := m.GetBackingStore().Get("removeSuccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetRoleId gets the roleId property value. The roleId property
func (m *RoleSuccessStatistics) GetRoleId()(*string) {
    val, err := m.GetBackingStore().Get("roleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleName gets the roleName property value. The roleName property
func (m *RoleSuccessStatistics) GetRoleName()(*string) {
    val, err := m.GetBackingStore().Get("roleName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTemporaryFail gets the temporaryFail property value. The temporaryFail property
func (m *RoleSuccessStatistics) GetTemporaryFail()(*int64) {
    val, err := m.GetBackingStore().Get("temporaryFail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTemporarySuccess gets the temporarySuccess property value. The temporarySuccess property
func (m *RoleSuccessStatistics) GetTemporarySuccess()(*int64) {
    val, err := m.GetBackingStore().Get("temporarySuccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetUnknownFail gets the unknownFail property value. The unknownFail property
func (m *RoleSuccessStatistics) GetUnknownFail()(*int64) {
    val, err := m.GetBackingStore().Get("unknownFail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RoleSuccessStatistics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("permanentFail", m.GetPermanentFail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("permanentSuccess", m.GetPermanentSuccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("removeFail", m.GetRemoveFail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("removeSuccess", m.GetRemoveSuccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("roleId", m.GetRoleId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("roleName", m.GetRoleName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("temporaryFail", m.GetTemporaryFail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("temporarySuccess", m.GetTemporarySuccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("unknownFail", m.GetUnknownFail())
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
func (m *RoleSuccessStatistics) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *RoleSuccessStatistics) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RoleSuccessStatistics) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPermanentFail sets the permanentFail property value. The permanentFail property
func (m *RoleSuccessStatistics) SetPermanentFail(value *int64)() {
    err := m.GetBackingStore().Set("permanentFail", value)
    if err != nil {
        panic(err)
    }
}
// SetPermanentSuccess sets the permanentSuccess property value. The permanentSuccess property
func (m *RoleSuccessStatistics) SetPermanentSuccess(value *int64)() {
    err := m.GetBackingStore().Set("permanentSuccess", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoveFail sets the removeFail property value. The removeFail property
func (m *RoleSuccessStatistics) SetRemoveFail(value *int64)() {
    err := m.GetBackingStore().Set("removeFail", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoveSuccess sets the removeSuccess property value. The removeSuccess property
func (m *RoleSuccessStatistics) SetRemoveSuccess(value *int64)() {
    err := m.GetBackingStore().Set("removeSuccess", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleId sets the roleId property value. The roleId property
func (m *RoleSuccessStatistics) SetRoleId(value *string)() {
    err := m.GetBackingStore().Set("roleId", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleName sets the roleName property value. The roleName property
func (m *RoleSuccessStatistics) SetRoleName(value *string)() {
    err := m.GetBackingStore().Set("roleName", value)
    if err != nil {
        panic(err)
    }
}
// SetTemporaryFail sets the temporaryFail property value. The temporaryFail property
func (m *RoleSuccessStatistics) SetTemporaryFail(value *int64)() {
    err := m.GetBackingStore().Set("temporaryFail", value)
    if err != nil {
        panic(err)
    }
}
// SetTemporarySuccess sets the temporarySuccess property value. The temporarySuccess property
func (m *RoleSuccessStatistics) SetTemporarySuccess(value *int64)() {
    err := m.GetBackingStore().Set("temporarySuccess", value)
    if err != nil {
        panic(err)
    }
}
// SetUnknownFail sets the unknownFail property value. The unknownFail property
func (m *RoleSuccessStatistics) SetUnknownFail(value *int64)() {
    err := m.GetBackingStore().Set("unknownFail", value)
    if err != nil {
        panic(err)
    }
}
// RoleSuccessStatisticsable 
type RoleSuccessStatisticsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetPermanentFail()(*int64)
    GetPermanentSuccess()(*int64)
    GetRemoveFail()(*int64)
    GetRemoveSuccess()(*int64)
    GetRoleId()(*string)
    GetRoleName()(*string)
    GetTemporaryFail()(*int64)
    GetTemporarySuccess()(*int64)
    GetUnknownFail()(*int64)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetPermanentFail(value *int64)()
    SetPermanentSuccess(value *int64)()
    SetRemoveFail(value *int64)()
    SetRemoveSuccess(value *int64)()
    SetRoleId(value *string)()
    SetRoleName(value *string)()
    SetTemporaryFail(value *int64)()
    SetTemporarySuccess(value *int64)()
    SetUnknownFail(value *int64)()
}
