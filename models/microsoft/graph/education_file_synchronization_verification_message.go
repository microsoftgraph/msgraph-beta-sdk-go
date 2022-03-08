package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationFileSynchronizationVerificationMessage provides operations to call the start method.
type EducationFileSynchronizationVerificationMessage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Detailed information about the message type.
    description *string;
    // 
    fileName *string;
    // Type of the message. Possible values are: error, warning, information.
    type_escaped *string;
}
// NewEducationFileSynchronizationVerificationMessage instantiates a new educationFileSynchronizationVerificationMessage and sets the default values.
func NewEducationFileSynchronizationVerificationMessage()(*EducationFileSynchronizationVerificationMessage) {
    m := &EducationFileSynchronizationVerificationMessage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationFileSynchronizationVerificationMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationFileSynchronizationVerificationMessageFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEducationFileSynchronizationVerificationMessage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationFileSynchronizationVerificationMessage) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. Detailed information about the message type.
func (m *EducationFileSynchronizationVerificationMessage) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationFileSynchronizationVerificationMessage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["fileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. 
func (m *EducationFileSynchronizationVerificationMessage) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// GetType gets the type property value. Type of the message. Possible values are: error, warning, information.
func (m *EducationFileSynchronizationVerificationMessage) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *EducationFileSynchronizationVerificationMessage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EducationFileSynchronizationVerificationMessage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *EducationFileSynchronizationVerificationMessage) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. Detailed information about the message type.
func (m *EducationFileSynchronizationVerificationMessage) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetFileName sets the fileName property value. 
func (m *EducationFileSynchronizationVerificationMessage) SetFileName(value *string)() {
    if m != nil {
        m.fileName = value
    }
}
// SetType sets the type property value. Type of the message. Possible values are: error, warning, information.
func (m *EducationFileSynchronizationVerificationMessage) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
