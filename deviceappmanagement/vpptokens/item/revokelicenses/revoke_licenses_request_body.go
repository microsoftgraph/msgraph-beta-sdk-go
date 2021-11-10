package revokelicenses

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RevokeLicensesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    notifyManagedDevices *bool;
    // 
    revokeUntrackedLicenses *bool;
}
// Instantiates a new revokeLicensesRequestBody and sets the default values.
func NewRevokeLicensesRequestBody()(*RevokeLicensesRequestBody) {
    m := &RevokeLicensesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RevokeLicensesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the notifyManagedDevices property value. 
func (m *RevokeLicensesRequestBody) GetNotifyManagedDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notifyManagedDevices
    }
}
// Gets the revokeUntrackedLicenses property value. 
func (m *RevokeLicensesRequestBody) GetRevokeUntrackedLicenses()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.revokeUntrackedLicenses
    }
}
// The deserialization information for the current model
func (m *RevokeLicensesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["notifyManagedDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotifyManagedDevices(val)
        }
        return nil
    }
    res["revokeUntrackedLicenses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRevokeUntrackedLicenses(val)
        }
        return nil
    }
    return res
}
func (m *RevokeLicensesRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RevokeLicensesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("notifyManagedDevices", m.GetNotifyManagedDevices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("revokeUntrackedLicenses", m.GetRevokeUntrackedLicenses())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *RevokeLicensesRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the notifyManagedDevices property value. 
// Parameters:
//  - value : Value to set for the notifyManagedDevices property.
func (m *RevokeLicensesRequestBody) SetNotifyManagedDevices(value *bool)() {
    m.notifyManagedDevices = value
}
// Sets the revokeUntrackedLicenses property value. 
// Parameters:
//  - value : Value to set for the revokeUntrackedLicenses property.
func (m *RevokeLicensesRequestBody) SetRevokeUntrackedLicenses(value *bool)() {
    m.revokeUntrackedLicenses = value
}
