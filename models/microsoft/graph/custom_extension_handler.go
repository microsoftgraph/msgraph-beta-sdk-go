package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CustomExtensionHandler 
type CustomExtensionHandler struct {
    Entity
    // 
    customExtension *CustomAccessPackageWorkflowExtension;
    // 
    stage *AccessPackageCustomExtensionStage;
}
// NewCustomExtensionHandler instantiates a new customExtensionHandler and sets the default values.
func NewCustomExtensionHandler()(*CustomExtensionHandler) {
    m := &CustomExtensionHandler{
        Entity: *NewEntity(),
    }
    return m
}
// GetCustomExtension gets the customExtension property value. 
func (m *CustomExtensionHandler) GetCustomExtension()(*CustomAccessPackageWorkflowExtension) {
    if m == nil {
        return nil
    } else {
        return m.customExtension
    }
}
// GetStage gets the stage property value. 
func (m *CustomExtensionHandler) GetStage()(*AccessPackageCustomExtensionStage) {
    if m == nil {
        return nil
    } else {
        return m.stage
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomExtensionHandler) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customExtension"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomAccessPackageWorkflowExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtension(val.(*CustomAccessPackageWorkflowExtension))
        }
        return nil
    }
    res["stage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageCustomExtensionStage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStage(val.(*AccessPackageCustomExtensionStage))
        }
        return nil
    }
    return res
}
func (m *CustomExtensionHandler) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CustomExtensionHandler) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("customExtension", m.GetCustomExtension())
        if err != nil {
            return err
        }
    }
    if m.GetStage() != nil {
        cast := (*m.GetStage()).String()
        err = writer.WriteStringValue("stage", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomExtension sets the customExtension property value. 
func (m *CustomExtensionHandler) SetCustomExtension(value *CustomAccessPackageWorkflowExtension)() {
    if m != nil {
        m.customExtension = value
    }
}
// SetStage sets the stage property value. 
func (m *CustomExtensionHandler) SetStage(value *AccessPackageCustomExtensionStage)() {
    if m != nil {
        m.stage = value
    }
}
