package redirect

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// RedirectRequestBody 
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
// NewRedirectRequestBody instantiates a new redirectRequestBody and sets the default values.
func NewRedirectRequestBody()(*RedirectRequestBody) {
    m := &RedirectRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedirectRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCallbackUri gets the callbackUri property value. 
func (m *RedirectRequestBody) GetCallbackUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callbackUri
    }
}
// GetMaskCallee gets the maskCallee property value. 
func (m *RedirectRequestBody) GetMaskCallee()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.maskCallee
    }
}
// GetMaskCaller gets the maskCaller property value. 
func (m *RedirectRequestBody) GetMaskCaller()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.maskCaller
    }
}
// GetTargetDisposition gets the targetDisposition property value. 
func (m *RedirectRequestBody) GetTargetDisposition()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CallDisposition) {
    if m == nil {
        return nil
    } else {
        return m.targetDisposition
    }
}
// GetTargets gets the targets property value. 
func (m *RedirectRequestBody) GetTargets()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InvitationParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.targets
    }
}
// GetTimeout gets the timeout property value. 
func (m *RedirectRequestBody) GetTimeout()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.timeout
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RedirectRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["callbackUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackUri(val)
        }
        return nil
    }
    res["maskCallee"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaskCallee(val)
        }
        return nil
    }
    res["maskCaller"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaskCaller(val)
        }
        return nil
    }
    res["targetDisposition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseCallDisposition)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDisposition(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CallDisposition))
        }
        return nil
    }
    res["targets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewInvitationParticipantInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InvitationParticipantInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InvitationParticipantInfo))
            }
            m.SetTargets(res)
        }
        return nil
    }
    res["timeout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeout(val)
        }
        return nil
    }
    return res
}
func (m *RedirectRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetTargetDisposition()).String()
        err := writer.WriteStringValue("targetDisposition", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargets() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedirectRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCallbackUri sets the callbackUri property value. 
func (m *RedirectRequestBody) SetCallbackUri(value *string)() {
    if m != nil {
        m.callbackUri = value
    }
}
// SetMaskCallee sets the maskCallee property value. 
func (m *RedirectRequestBody) SetMaskCallee(value *bool)() {
    if m != nil {
        m.maskCallee = value
    }
}
// SetMaskCaller sets the maskCaller property value. 
func (m *RedirectRequestBody) SetMaskCaller(value *bool)() {
    if m != nil {
        m.maskCaller = value
    }
}
// SetTargetDisposition sets the targetDisposition property value. 
func (m *RedirectRequestBody) SetTargetDisposition(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CallDisposition)() {
    if m != nil {
        m.targetDisposition = value
    }
}
// SetTargets sets the targets property value. 
func (m *RedirectRequestBody) SetTargets(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InvitationParticipantInfo)() {
    if m != nil {
        m.targets = value
    }
}
// SetTimeout sets the timeout property value. 
func (m *RedirectRequestBody) SetTimeout(value *int32)() {
    if m != nil {
        m.timeout = value
    }
}
