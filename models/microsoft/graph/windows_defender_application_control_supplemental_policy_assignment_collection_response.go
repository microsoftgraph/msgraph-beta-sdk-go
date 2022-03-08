package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse provides operations to manage the assignments property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
type WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    nextLink *string;
    // 
    value []WindowsDefenderApplicationControlSupplementalPolicyAssignmentable;
}
// NewWindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse instantiates a new WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse()(*WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["@odata.nextLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextLink(val)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsDefenderApplicationControlSupplementalPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDefenderApplicationControlSupplementalPolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsDefenderApplicationControlSupplementalPolicyAssignmentable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetNextLink gets the @odata.nextLink property value. 
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse) GetNextLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nextLink
    }
}
// GetValue gets the value property value. 
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse) GetValue()([]WindowsDefenderApplicationControlSupplementalPolicyAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.nextLink", m.GetNextLink())
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNextLink sets the @odata.nextLink property value. 
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse) SetNextLink(value *string)() {
    if m != nil {
        m.nextLink = value
    }
}
// SetValue sets the value property value. 
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignmentCollectionResponse) SetValue(value []WindowsDefenderApplicationControlSupplementalPolicyAssignmentable)() {
    if m != nil {
        m.value = value
    }
}
