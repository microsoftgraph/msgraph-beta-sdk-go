package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RoleSuccessStatistics provides operations to call the completeSetup method.
type RoleSuccessStatistics struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    permanentFail *int32;
    // 
    permanentSuccess *int32;
    // 
    removeFail *int32;
    // 
    removeSuccess *int32;
    // 
    roleId *string;
    // 
    roleName *string;
    // 
    temporaryFail *int32;
    // 
    temporarySuccess *int32;
    // 
    unknownFail *int32;
}
// NewRoleSuccessStatistics instantiates a new roleSuccessStatistics and sets the default values.
func NewRoleSuccessStatistics()(*RoleSuccessStatistics) {
    m := &RoleSuccessStatistics{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRoleSuccessStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleSuccessStatisticsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRoleSuccessStatistics(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleSuccessStatistics) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleSuccessStatistics) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["permanentFail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermanentFail(val)
        }
        return nil
    }
    res["permanentSuccess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermanentSuccess(val)
        }
        return nil
    }
    res["removeFail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveFail(val)
        }
        return nil
    }
    res["removeSuccess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveSuccess(val)
        }
        return nil
    }
    res["roleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleId(val)
        }
        return nil
    }
    res["roleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleName(val)
        }
        return nil
    }
    res["temporaryFail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemporaryFail(val)
        }
        return nil
    }
    res["temporarySuccess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemporarySuccess(val)
        }
        return nil
    }
    res["unknownFail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
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
// GetPermanentFail gets the permanentFail property value. 
func (m *RoleSuccessStatistics) GetPermanentFail()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.permanentFail
    }
}
// GetPermanentSuccess gets the permanentSuccess property value. 
func (m *RoleSuccessStatistics) GetPermanentSuccess()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.permanentSuccess
    }
}
// GetRemoveFail gets the removeFail property value. 
func (m *RoleSuccessStatistics) GetRemoveFail()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.removeFail
    }
}
// GetRemoveSuccess gets the removeSuccess property value. 
func (m *RoleSuccessStatistics) GetRemoveSuccess()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.removeSuccess
    }
}
// GetRoleId gets the roleId property value. 
func (m *RoleSuccessStatistics) GetRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleId
    }
}
// GetRoleName gets the roleName property value. 
func (m *RoleSuccessStatistics) GetRoleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleName
    }
}
// GetTemporaryFail gets the temporaryFail property value. 
func (m *RoleSuccessStatistics) GetTemporaryFail()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.temporaryFail
    }
}
// GetTemporarySuccess gets the temporarySuccess property value. 
func (m *RoleSuccessStatistics) GetTemporarySuccess()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.temporarySuccess
    }
}
// GetUnknownFail gets the unknownFail property value. 
func (m *RoleSuccessStatistics) GetUnknownFail()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownFail
    }
}
func (m *RoleSuccessStatistics) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RoleSuccessStatistics) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("permanentFail", m.GetPermanentFail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("permanentSuccess", m.GetPermanentSuccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("removeFail", m.GetRemoveFail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("removeSuccess", m.GetRemoveSuccess())
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
        err := writer.WriteInt32Value("temporaryFail", m.GetTemporaryFail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("temporarySuccess", m.GetTemporarySuccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unknownFail", m.GetUnknownFail())
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
    if m != nil {
        m.additionalData = value
    }
}
// SetPermanentFail sets the permanentFail property value. 
func (m *RoleSuccessStatistics) SetPermanentFail(value *int32)() {
    if m != nil {
        m.permanentFail = value
    }
}
// SetPermanentSuccess sets the permanentSuccess property value. 
func (m *RoleSuccessStatistics) SetPermanentSuccess(value *int32)() {
    if m != nil {
        m.permanentSuccess = value
    }
}
// SetRemoveFail sets the removeFail property value. 
func (m *RoleSuccessStatistics) SetRemoveFail(value *int32)() {
    if m != nil {
        m.removeFail = value
    }
}
// SetRemoveSuccess sets the removeSuccess property value. 
func (m *RoleSuccessStatistics) SetRemoveSuccess(value *int32)() {
    if m != nil {
        m.removeSuccess = value
    }
}
// SetRoleId sets the roleId property value. 
func (m *RoleSuccessStatistics) SetRoleId(value *string)() {
    if m != nil {
        m.roleId = value
    }
}
// SetRoleName sets the roleName property value. 
func (m *RoleSuccessStatistics) SetRoleName(value *string)() {
    if m != nil {
        m.roleName = value
    }
}
// SetTemporaryFail sets the temporaryFail property value. 
func (m *RoleSuccessStatistics) SetTemporaryFail(value *int32)() {
    if m != nil {
        m.temporaryFail = value
    }
}
// SetTemporarySuccess sets the temporarySuccess property value. 
func (m *RoleSuccessStatistics) SetTemporarySuccess(value *int32)() {
    if m != nil {
        m.temporarySuccess = value
    }
}
// SetUnknownFail sets the unknownFail property value. 
func (m *RoleSuccessStatistics) SetUnknownFail(value *int32)() {
    if m != nil {
        m.unknownFail = value
    }
}
