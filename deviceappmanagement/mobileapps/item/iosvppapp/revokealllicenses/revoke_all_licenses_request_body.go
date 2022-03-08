package revokealllicenses

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RevokeAllLicensesRequestBody provides operations to call the revokeAllLicenses method.
type RevokeAllLicensesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    notifyManagedDevices *bool;
}
// NewRevokeAllLicensesRequestBody instantiates a new revokeAllLicensesRequestBody and sets the default values.
func NewRevokeAllLicensesRequestBody()(*RevokeAllLicensesRequestBody) {
    m := &RevokeAllLicensesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRevokeAllLicensesRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRevokeAllLicensesRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRevokeAllLicensesRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RevokeAllLicensesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RevokeAllLicensesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    return res
}
// GetNotifyManagedDevices gets the notifyManagedDevices property value. 
func (m *RevokeAllLicensesRequestBody) GetNotifyManagedDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notifyManagedDevices
    }
}
func (m *RevokeAllLicensesRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RevokeAllLicensesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("notifyManagedDevices", m.GetNotifyManagedDevices())
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
func (m *RevokeAllLicensesRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNotifyManagedDevices sets the notifyManagedDevices property value. 
func (m *RevokeAllLicensesRequestBody) SetNotifyManagedDevices(value *bool)() {
    if m != nil {
        m.notifyManagedDevices = value
    }
}
