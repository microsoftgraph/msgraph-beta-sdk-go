package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RichLongRunningOperation 
type RichLongRunningOperation struct {
    LongRunningOperation
    // Error due to which the operation failed.
    error *PublicError;
    // A value between 0 and 100 that indicates the progress of the operation.
    percentageComplete *int32;
    // A unique identifier for the result.
    resourceId *string;
    // Type of the operation.
    type_escaped *string;
}
// NewRichLongRunningOperation instantiates a new richLongRunningOperation and sets the default values.
func NewRichLongRunningOperation()(*RichLongRunningOperation) {
    m := &RichLongRunningOperation{
        LongRunningOperation: *NewLongRunningOperation(),
    }
    return m
}
// GetError gets the error property value. Error due to which the operation failed.
func (m *RichLongRunningOperation) GetError()(*PublicError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetPercentageComplete gets the percentageComplete property value. A value between 0 and 100 that indicates the progress of the operation.
func (m *RichLongRunningOperation) GetPercentageComplete()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.percentageComplete
    }
}
// GetResourceId gets the resourceId property value. A unique identifier for the result.
func (m *RichLongRunningOperation) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// GetType gets the type property value. Type of the operation.
func (m *RichLongRunningOperation) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RichLongRunningOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.LongRunningOperation.GetFieldDeserializers()
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublicError() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(*PublicError))
        }
        return nil
    }
    res["percentageComplete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentageComplete(val)
        }
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
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
func (m *RichLongRunningOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RichLongRunningOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.LongRunningOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("percentageComplete", m.GetPercentageComplete())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. Error due to which the operation failed.
func (m *RichLongRunningOperation) SetError(value *PublicError)() {
    if m != nil {
        m.error = value
    }
}
// SetPercentageComplete sets the percentageComplete property value. A value between 0 and 100 that indicates the progress of the operation.
func (m *RichLongRunningOperation) SetPercentageComplete(value *int32)() {
    if m != nil {
        m.percentageComplete = value
    }
}
// SetResourceId sets the resourceId property value. A unique identifier for the result.
func (m *RichLongRunningOperation) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
// SetType sets the type property value. Type of the operation.
func (m *RichLongRunningOperation) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
