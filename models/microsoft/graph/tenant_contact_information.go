package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TenantContactInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The email address for the contact. Optional
    email *string;
    // The name for the contact. Required.
    name *string;
    // The notes associated with the contact. Optional
    notes *string;
    // The phone number for the contact. Optional.
    phone *string;
    // The title for the contact. Required.
    title *string;
}
// Instantiates a new tenantContactInformation and sets the default values.
func NewTenantContactInformation()(*TenantContactInformation) {
    m := &TenantContactInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantContactInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the email property value. The email address for the contact. Optional
func (m *TenantContactInformation) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the name property value. The name for the contact. Required.
func (m *TenantContactInformation) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the notes property value. The notes associated with the contact. Optional
func (m *TenantContactInformation) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// Gets the phone property value. The phone number for the contact. Optional.
func (m *TenantContactInformation) GetPhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phone
    }
}
// Gets the title property value. The title for the contact. Required.
func (m *TenantContactInformation) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// The deserialization information for the current model
func (m *TenantContactInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["phone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val)
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
func (m *TenantContactInformation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TenantContactInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *TenantContactInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the email property value. The email address for the contact. Optional
// Parameters:
//  - value : Value to set for the email property.
func (m *TenantContactInformation) SetEmail(value *string)() {
    m.email = value
}
// Sets the name property value. The name for the contact. Required.
// Parameters:
//  - value : Value to set for the name property.
func (m *TenantContactInformation) SetName(value *string)() {
    m.name = value
}
// Sets the notes property value. The notes associated with the contact. Optional
// Parameters:
//  - value : Value to set for the notes property.
func (m *TenantContactInformation) SetNotes(value *string)() {
    m.notes = value
}
// Sets the phone property value. The phone number for the contact. Optional.
// Parameters:
//  - value : Value to set for the phone property.
func (m *TenantContactInformation) SetPhone(value *string)() {
    m.phone = value
}
// Sets the title property value. The title for the contact. Required.
// Parameters:
//  - value : Value to set for the title property.
func (m *TenantContactInformation) SetTitle(value *string)() {
    m.title = value
}
