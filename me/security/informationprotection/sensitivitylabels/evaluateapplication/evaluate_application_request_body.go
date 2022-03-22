package evaluateapplication

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/security"
)

// EvaluateApplicationRequestBody provides operations to call the evaluateApplication method.
type EvaluateApplicationRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    contentInfo i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ContentInfoable;
    // 
    labelingOptions i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.LabelingOptionsable;
}
// NewEvaluateApplicationRequestBody instantiates a new evaluateApplicationRequestBody and sets the default values.
func NewEvaluateApplicationRequestBody()(*EvaluateApplicationRequestBody) {
    m := &EvaluateApplicationRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEvaluateApplicationRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEvaluateApplicationRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEvaluateApplicationRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateApplicationRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentInfo gets the contentInfo property value. 
func (m *EvaluateApplicationRequestBody) GetContentInfo()(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ContentInfoable) {
    if m == nil {
        return nil
    } else {
        return m.contentInfo
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateApplicationRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.CreateContentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ContentInfoable))
        }
        return nil
    }
    res["labelingOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.CreateLabelingOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelingOptions(val.(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.LabelingOptionsable))
        }
        return nil
    }
    return res
}
// GetLabelingOptions gets the labelingOptions property value. 
func (m *EvaluateApplicationRequestBody) GetLabelingOptions()(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.LabelingOptionsable) {
    if m == nil {
        return nil
    } else {
        return m.labelingOptions
    }
}
// Serialize serializes information the current object
func (m *EvaluateApplicationRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentInfo", m.GetContentInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("labelingOptions", m.GetLabelingOptions())
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
func (m *EvaluateApplicationRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentInfo sets the contentInfo property value. 
func (m *EvaluateApplicationRequestBody) SetContentInfo(value i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ContentInfoable)() {
    if m != nil {
        m.contentInfo = value
    }
}
// SetLabelingOptions sets the labelingOptions property value. 
func (m *EvaluateApplicationRequestBody) SetLabelingOptions(value i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.LabelingOptionsable)() {
    if m != nil {
        m.labelingOptions = value
    }
}
