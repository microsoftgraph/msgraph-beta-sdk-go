package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ImportedDeviceIdentityResult 
type ImportedDeviceIdentityResult struct {
    ImportedDeviceIdentity
    // Status of imported device identity
    status *bool;
}
// NewImportedDeviceIdentityResult instantiates a new importedDeviceIdentityResult and sets the default values.
func NewImportedDeviceIdentityResult()(*ImportedDeviceIdentityResult) {
    m := &ImportedDeviceIdentityResult{
        ImportedDeviceIdentity: *NewImportedDeviceIdentity(),
    }
    return m
}
// CreateImportedDeviceIdentityResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedDeviceIdentityResultFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewImportedDeviceIdentityResult(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportedDeviceIdentityResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ImportedDeviceIdentity.GetFieldDeserializers()
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
func (m *ImportedDeviceIdentityResult) GetStatus()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Serialize serializes information the current object
func (m *ImportedDeviceIdentityResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ImportedDeviceIdentity.Serialize(writer)
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
func (m *ImportedDeviceIdentityResult) SetStatus(value *bool)() {
    if m != nil {
        m.status = value
    }
}
