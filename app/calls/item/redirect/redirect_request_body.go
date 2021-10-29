package redirect

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type RedirectRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    callbackUri *string;
    // 
    maskCallee *bool;
    // 
    maskCaller *bool;
    // 
    targetDisposition *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CallDisposition;
    // 
    targets []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InvitationParticipantInfo;
    // 
    timeout *int32;
}
// Instantiates a new redirectRequestBody and sets the default values.
func NewRedirectRequestBody()(*RedirectRequestBody) {
    m := &RedirectRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedirectRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the callbackUri property value. 
func (m *RedirectRequestBody) GetCallbackUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callbackUri
    }
}
// Gets the maskCallee property value. 
func (m *RedirectRequestBody) GetMaskCallee()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.maskCallee
    }
}
// Gets the maskCaller property value. 
func (m *RedirectRequestBody) GetMaskCaller()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.maskCaller
    }
}
// Gets the targetDisposition property value. 
func (m *RedirectRequestBody) GetTargetDisposition()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CallDisposition) {
    if m == nil {
        return nil
    } else {
        return m.targetDisposition
    }
}
// Gets the targets property value. 
func (m *RedirectRequestBody) GetTargets()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InvitationParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.targets
    }
}
// Gets the timeout property value. 
func (m *RedirectRequestBody) GetTimeout()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.timeout
    }
}
// The deserialization information for the current model
func (m *RedirectRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["callbackUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCallbackUri(val)
        return nil
    }
    res["maskCallee"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMaskCallee(val)
        return nil
    }
    res["maskCaller"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMaskCaller(val)
        return nil
    }
    res["targetDisposition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseCallDisposition)
        if err != nil {
            return err
        }
        cast := val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CallDisposition)
        m.SetTargetDisposition(&cast)
        return nil
    }
    res["targets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewInvitationParticipantInfo() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InvitationParticipantInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InvitationParticipantInfo))
        }
        m.SetTargets(res)
        return nil
    }
    res["timeout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTimeout(val)
        return nil
    }
    return res
}
func (m *RedirectRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RedirectRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("maskCallee", m.GetMaskCallee())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("maskCaller", m.GetMaskCaller())
        if err != nil {
            return err
        }
    }
    if m.GetTargetDisposition() != nil {
        cast := m.GetTargetDisposition().String()
        err := writer.WriteStringValue("targetDisposition", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTargets()))
        for i, v := range m.GetTargets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("targets", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("timeout", m.GetTimeout())
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
func (m *RedirectRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the callbackUri property value. 
// Parameters:
//  - value : Value to set for the callbackUri property.
func (m *RedirectRequestBody) SetCallbackUri(value *string)() {
    m.callbackUri = value
}
// Sets the maskCallee property value. 
// Parameters:
//  - value : Value to set for the maskCallee property.
func (m *RedirectRequestBody) SetMaskCallee(value *bool)() {
    m.maskCallee = value
}
// Sets the maskCaller property value. 
// Parameters:
//  - value : Value to set for the maskCaller property.
func (m *RedirectRequestBody) SetMaskCaller(value *bool)() {
    m.maskCaller = value
}
// Sets the targetDisposition property value. 
// Parameters:
//  - value : Value to set for the targetDisposition property.
func (m *RedirectRequestBody) SetTargetDisposition(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CallDisposition)() {
    m.targetDisposition = value
}
// Sets the targets property value. 
// Parameters:
//  - value : Value to set for the targets property.
func (m *RedirectRequestBody) SetTargets(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InvitationParticipantInfo)() {
    m.targets = value
}
// Sets the timeout property value. 
// Parameters:
//  - value : Value to set for the timeout property.
func (m *RedirectRequestBody) SetTimeout(value *int32)() {
    m.timeout = value
}
