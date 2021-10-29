package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TargetPolicyEndpoints struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Use to filter the notification distribution to a specific platform or platforms. Valid values are Windows, iOS, Android and WebPush. By default, all push endpoint types (Windows, iOS, Android and WebPush) are enabled.
    platformTypes []string;
}
// Instantiates a new targetPolicyEndpoints and sets the default values.
func NewTargetPolicyEndpoints()(*TargetPolicyEndpoints) {
    m := &TargetPolicyEndpoints{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TargetPolicyEndpoints) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the platformTypes property value. Use to filter the notification distribution to a specific platform or platforms. Valid values are Windows, iOS, Android and WebPush. By default, all push endpoint types (Windows, iOS, Android and WebPush) are enabled.
func (m *TargetPolicyEndpoints) GetPlatformTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.platformTypes
    }
}
// The deserialization information for the current model
func (m *TargetPolicyEndpoints) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["platformTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetPlatformTypes(res)
        return nil
    }
    return res
}
func (m *TargetPolicyEndpoints) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TargetPolicyEndpoints) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("platformTypes", m.GetPlatformTypes())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TargetPolicyEndpoints) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the platformTypes property value. Use to filter the notification distribution to a specific platform or platforms. Valid values are Windows, iOS, Android and WebPush. By default, all push endpoint types (Windows, iOS, Android and WebPush) are enabled.
// Parameters:
//  - value : Value to set for the platformTypes property.
func (m *TargetPolicyEndpoints) SetPlatformTypes(value []string)() {
    m.platformTypes = value
}
