package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleSuccessStatistics 
type RoleSuccessStatistics struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The permanentFail property
    permanentFail *int64
    // The permanentSuccess property
    permanentSuccess *int64
    // The removeFail property
    removeFail *int64
    // The removeSuccess property
    removeSuccess *int64
    // The roleId property
    roleId *string
    // The roleName property
    roleName *string
    // The temporaryFail property
    temporaryFail *int64
    // The temporarySuccess property
    temporarySuccess *int64
    // The unknownFail property
    unknownFail *int64
}
// NewRoleSuccessStatistics instantiates a new roleSuccessStatistics and sets the default values.
func NewRoleSuccessStatistics()(*RoleSuccessStatistics) {
    m := &RoleSuccessStatistics{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.roleSuccessStatistics";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRoleSuccessStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleSuccessStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleSuccessStatistics(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleSuccessStatistics) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleSuccessStatistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["permanentFail"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetPermanentFail)
    res["permanentSuccess"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetPermanentSuccess)
    res["removeFail"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetRemoveFail)
    res["removeSuccess"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetRemoveSuccess)
    res["roleId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRoleId)
    res["roleName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRoleName)
    res["temporaryFail"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetTemporaryFail)
    res["temporarySuccess"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetTemporarySuccess)
    res["unknownFail"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetUnknownFail)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RoleSuccessStatistics) GetOdataType()(*string) {
    return m.odataType
}
// GetPermanentFail gets the permanentFail property value. The permanentFail property
func (m *RoleSuccessStatistics) GetPermanentFail()(*int64) {
    return m.permanentFail
}
// GetPermanentSuccess gets the permanentSuccess property value. The permanentSuccess property
func (m *RoleSuccessStatistics) GetPermanentSuccess()(*int64) {
    return m.permanentSuccess
}
// GetRemoveFail gets the removeFail property value. The removeFail property
func (m *RoleSuccessStatistics) GetRemoveFail()(*int64) {
    return m.removeFail
}
// GetRemoveSuccess gets the removeSuccess property value. The removeSuccess property
func (m *RoleSuccessStatistics) GetRemoveSuccess()(*int64) {
    return m.removeSuccess
}
// GetRoleId gets the roleId property value. The roleId property
func (m *RoleSuccessStatistics) GetRoleId()(*string) {
    return m.roleId
}
// GetRoleName gets the roleName property value. The roleName property
func (m *RoleSuccessStatistics) GetRoleName()(*string) {
    return m.roleName
}
// GetTemporaryFail gets the temporaryFail property value. The temporaryFail property
func (m *RoleSuccessStatistics) GetTemporaryFail()(*int64) {
    return m.temporaryFail
}
// GetTemporarySuccess gets the temporarySuccess property value. The temporarySuccess property
func (m *RoleSuccessStatistics) GetTemporarySuccess()(*int64) {
    return m.temporarySuccess
}
// GetUnknownFail gets the unknownFail property value. The unknownFail property
func (m *RoleSuccessStatistics) GetUnknownFail()(*int64) {
    return m.unknownFail
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
func (m *RoleSuccessStatistics) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RoleSuccessStatistics) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPermanentFail sets the permanentFail property value. The permanentFail property
func (m *RoleSuccessStatistics) SetPermanentFail(value *int64)() {
    m.permanentFail = value
}
// SetPermanentSuccess sets the permanentSuccess property value. The permanentSuccess property
func (m *RoleSuccessStatistics) SetPermanentSuccess(value *int64)() {
    m.permanentSuccess = value
}
// SetRemoveFail sets the removeFail property value. The removeFail property
func (m *RoleSuccessStatistics) SetRemoveFail(value *int64)() {
    m.removeFail = value
}
// SetRemoveSuccess sets the removeSuccess property value. The removeSuccess property
func (m *RoleSuccessStatistics) SetRemoveSuccess(value *int64)() {
    m.removeSuccess = value
}
// SetRoleId sets the roleId property value. The roleId property
func (m *RoleSuccessStatistics) SetRoleId(value *string)() {
    m.roleId = value
}
// SetRoleName sets the roleName property value. The roleName property
func (m *RoleSuccessStatistics) SetRoleName(value *string)() {
    m.roleName = value
}
// SetTemporaryFail sets the temporaryFail property value. The temporaryFail property
func (m *RoleSuccessStatistics) SetTemporaryFail(value *int64)() {
    m.temporaryFail = value
}
// SetTemporarySuccess sets the temporarySuccess property value. The temporarySuccess property
func (m *RoleSuccessStatistics) SetTemporarySuccess(value *int64)() {
    m.temporarySuccess = value
}
// SetUnknownFail sets the unknownFail property value. The unknownFail property
func (m *RoleSuccessStatistics) SetUnknownFail(value *int64)() {
    m.unknownFail = value
}
