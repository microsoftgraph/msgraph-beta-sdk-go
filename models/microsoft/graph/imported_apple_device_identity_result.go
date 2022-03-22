package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ImportedAppleDeviceIdentityResult 
type ImportedAppleDeviceIdentityResult struct {
    ImportedAppleDeviceIdentity
    // Status of imported device identity
    status *bool;
}
// NewImportedAppleDeviceIdentityResult instantiates a new importedAppleDeviceIdentityResult and sets the default values.
func NewImportedAppleDeviceIdentityResult()(*ImportedAppleDeviceIdentityResult) {
    m := &ImportedAppleDeviceIdentityResult{
        ImportedAppleDeviceIdentity: *NewImportedAppleDeviceIdentity(),
    }
    return m
}
// CreateImportedAppleDeviceIdentityResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedAppleDeviceIdentityResultFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewImportedAppleDeviceIdentityResult(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportedAppleDeviceIdentityResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ImportedAppleDeviceIdentity.GetFieldDeserializers()
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. Status of imported device identity
func (m *ImportedAppleDeviceIdentityResult) GetStatus()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Serialize serializes information the current object
func (m *ImportedAppleDeviceIdentityResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ImportedAppleDeviceIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetStatus sets the status property value. Status of imported device identity
func (m *ImportedAppleDeviceIdentityResult) SetStatus(value *bool)() {
    if m != nil {
        m.status = value
    }
}
