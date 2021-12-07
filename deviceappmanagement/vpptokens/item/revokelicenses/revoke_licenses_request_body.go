package revokelicenses

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RevokeLicensesRequestBody 
type RevokeLicensesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    notifyManagedDevices *bool;
    // 
    revokeUntrackedLicenses *bool;
}
// NewRevokeLicensesRequestBody instantiates a new revokeLicensesRequestBody and sets the default values.
func NewRevokeLicensesRequestBody()(*RevokeLicensesRequestBody) {
    m := &RevokeLicensesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RevokeLicensesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetNotifyManagedDevices gets the notifyManagedDevices property value. 
func (m *RevokeLicensesRequestBody) GetNotifyManagedDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notifyManagedDevices
    }
}
// GetRevokeUntrackedLicenses gets the revokeUntrackedLicenses property value. 
func (m *RevokeLicensesRequestBody) GetRevokeUntrackedLicenses()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.revokeUntrackedLicenses
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RevokeLicensesRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNotifyManagedDevices sets the notifyManagedDevices property value. 
func (m *RevokeLicensesRequestBody) SetNotifyManagedDevices(value *bool)() {
    if m != nil {
        m.notifyManagedDevices = value
    }
}
// SetRevokeUntrackedLicenses sets the revokeUntrackedLicenses property value. 
func (m *RevokeLicensesRequestBody) SetRevokeUntrackedLicenses(value *bool)() {
    if m != nil {
        m.revokeUntrackedLicenses = value
    }
}
