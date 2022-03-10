package evaluateremoval

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/security"
)

// EvaluateRemovalRequestBody provides operations to call the evaluateRemoval method.
type EvaluateRemovalRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    contentInfo i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ContentInfoable;
    // 
    downgradeJustification i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.DowngradeJustificationable;
}
// NewEvaluateRemovalRequestBody instantiates a new evaluateRemovalRequestBody and sets the default values.
func NewEvaluateRemovalRequestBody()(*EvaluateRemovalRequestBody) {
    m := &EvaluateRemovalRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEvaluateRemovalRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEvaluateRemovalRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEvaluateRemovalRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateRemovalRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentInfo gets the contentInfo property value. 
func (m *EvaluateRemovalRequestBody) GetContentInfo()(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ContentInfoable) {
    if m == nil {
        return nil
    } else {
        return m.contentInfo
    }
}
// GetDowngradeJustification gets the downgradeJustification property value. 
func (m *EvaluateRemovalRequestBody) GetDowngradeJustification()(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.DowngradeJustificationable) {
    if m == nil {
        return nil
    } else {
        return m.downgradeJustification
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateRemovalRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["downgradeJustification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.CreateDowngradeJustificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDowngradeJustification(val.(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.DowngradeJustificationable))
        }
        return nil
    }
    return res
}
func (m *EvaluateRemovalRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EvaluateRemovalRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentInfo", m.GetContentInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("downgradeJustification", m.GetDowngradeJustification())
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
func (m *EvaluateRemovalRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentInfo sets the contentInfo property value. 
func (m *EvaluateRemovalRequestBody) SetContentInfo(value i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ContentInfoable)() {
    if m != nil {
        m.contentInfo = value
    }
}
// SetDowngradeJustification sets the downgradeJustification property value. 
func (m *EvaluateRemovalRequestBody) SetDowngradeJustification(value i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.DowngradeJustificationable)() {
    if m != nil {
        m.downgradeJustification = value
    }
}
