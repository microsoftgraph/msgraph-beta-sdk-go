package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CustomExtensionHandlerInstance 
type CustomExtensionHandlerInstance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    customExtensionId *string;
    // 
    externalCorrelationId *string;
    // 
    stage *AccessPackageCustomExtensionStage;
    // 
    status *AccessPackageCustomExtensionHandlerStatus;
}
// NewCustomExtensionHandlerInstance instantiates a new customExtensionHandlerInstance and sets the default values.
func NewCustomExtensionHandlerInstance()(*CustomExtensionHandlerInstance) {
    m := &CustomExtensionHandlerInstance{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomExtensionHandlerInstance) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCustomExtensionId gets the customExtensionId property value. 
func (m *CustomExtensionHandlerInstance) GetCustomExtensionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customExtensionId
    }
}
// GetExternalCorrelationId gets the externalCorrelationId property value. 
func (m *CustomExtensionHandlerInstance) GetExternalCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalCorrelationId
    }
}
// GetStage gets the stage property value. 
func (m *CustomExtensionHandlerInstance) GetStage()(*AccessPackageCustomExtensionStage) {
    if m == nil {
        return nil
    } else {
        return m.stage
    }
}
// GetStatus gets the status property value. 
func (m *CustomExtensionHandlerInstance) GetStatus()(*AccessPackageCustomExtensionHandlerStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomExtensionHandlerInstance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["customExtensionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomExtensionId(val)
        }
        return nil
    }
    res["externalCorrelationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalCorrelationId(val)
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageCustomExtensionHandlerStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AccessPackageCustomExtensionHandlerStatus))
        }
        return nil
    }
    return res
}
func (m *CustomExtensionHandlerInstance) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CustomExtensionHandlerInstance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("customExtensionId", m.GetCustomExtensionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalCorrelationId", m.GetExternalCorrelationId())
        if err != nil {
            return err
        }
    }
    if m.GetStage() != nil {
        cast := (*m.GetStage()).String()
        err := writer.WriteStringValue("stage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *CustomExtensionHandlerInstance) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCustomExtensionId sets the customExtensionId property value. 
func (m *CustomExtensionHandlerInstance) SetCustomExtensionId(value *string)() {
    if m != nil {
        m.customExtensionId = value
    }
}
// SetExternalCorrelationId sets the externalCorrelationId property value. 
func (m *CustomExtensionHandlerInstance) SetExternalCorrelationId(value *string)() {
    if m != nil {
        m.externalCorrelationId = value
    }
}
// SetStage sets the stage property value. 
func (m *CustomExtensionHandlerInstance) SetStage(value *AccessPackageCustomExtensionStage)() {
    if m != nil {
        m.stage = value
    }
}
// SetStatus sets the status property value. 
func (m *CustomExtensionHandlerInstance) SetStatus(value *AccessPackageCustomExtensionHandlerStatus)() {
    if m != nil {
        m.status = value
    }
}
