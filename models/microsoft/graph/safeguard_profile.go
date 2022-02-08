package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/windowsupdates"
)

// SafeguardProfile 
type SafeguardProfile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies the category of safeguards. The possible values are: likelyIssues, unknownFutureValue.
    category *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.SafeguardCategory;
}
// NewSafeguardProfile instantiates a new safeguardProfile and sets the default values.
func NewSafeguardProfile()(*SafeguardProfile) {
    m := &SafeguardProfile{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SafeguardProfile) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCategory gets the category property value. Specifies the category of safeguards. The possible values are: likelyIssues, unknownFutureValue.
func (m *SafeguardProfile) GetCategory()(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.SafeguardCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SafeguardProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.ParseSafeguardCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.SafeguardCategory))
        }
        return nil
    }
    return res
}
func (m *SafeguardProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SafeguardProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err := writer.WriteStringValue("category", &cast)
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
func (m *SafeguardProfile) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCategory sets the category property value. Specifies the category of safeguards. The possible values are: likelyIssues, unknownFutureValue.
func (m *SafeguardProfile) SetCategory(value *ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.SafeguardCategory)() {
    if m != nil {
        m.category = value
    }
}
