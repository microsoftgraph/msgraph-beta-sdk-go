package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/security"
)

// ContentLabel 
type ContentLabel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    assignmentMethod *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AssignmentMethod;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    sensitivityLabelId *string;
}
// NewContentLabel instantiates a new contentLabel and sets the default values.
func NewContentLabel()(*ContentLabel) {
    m := &ContentLabel{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentLabel) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignmentMethod gets the assignmentMethod property value. 
func (m *ContentLabel) GetAssignmentMethod()(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AssignmentMethod) {
    if m == nil {
        return nil
    } else {
        return m.assignmentMethod
    }
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *ContentLabel) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetSensitivityLabelId gets the sensitivityLabelId property value. 
func (m *ContentLabel) GetSensitivityLabelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sensitivityLabelId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentLabel) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignmentMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ParseAssignmentMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentMethod(val.(*i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AssignmentMethod))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["sensitivityLabelId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivityLabelId(val)
        }
        return nil
    }
    return res
}
func (m *ContentLabel) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ContentLabel) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAssignmentMethod() != nil {
        cast := (*m.GetAssignmentMethod()).String()
        err := writer.WriteStringValue("assignmentMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sensitivityLabelId", m.GetSensitivityLabelId())
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
func (m *ContentLabel) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignmentMethod sets the assignmentMethod property value. 
func (m *ContentLabel) SetAssignmentMethod(value *i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.AssignmentMethod)() {
    if m != nil {
        m.assignmentMethod = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *ContentLabel) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetSensitivityLabelId sets the sensitivityLabelId property value. 
func (m *ContentLabel) SetSensitivityLabelId(value *string)() {
    if m != nil {
        m.sensitivityLabelId = value
    }
}
