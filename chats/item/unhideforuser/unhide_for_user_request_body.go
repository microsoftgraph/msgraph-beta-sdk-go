package unhideforuser

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// UnhideForUserRequestBody provides operations to call the unhideForUser method.
type UnhideForUserRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    tenantId *string;
    // 
    user i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkUserIdentityable;
}
// NewUnhideForUserRequestBody instantiates a new unhideForUserRequestBody and sets the default values.
func NewUnhideForUserRequestBody()(*UnhideForUserRequestBody) {
    m := &UnhideForUserRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUnhideForUserRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnhideForUserRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUnhideForUserRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnhideForUserRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnhideForUserRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["user"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateTeamworkUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkUserIdentityable))
        }
        return nil
    }
    return res
}
// GetTenantId gets the tenantId property value. 
func (m *UnhideForUserRequestBody) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetUser gets the user property value. 
func (m *UnhideForUserRequestBody) GetUser()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkUserIdentityable) {
    if m == nil {
        return nil
    } else {
        return m.user
    }
}
func (m *UnhideForUserRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UnhideForUserRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *UnhideForUserRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTenantId sets the tenantId property value. 
func (m *UnhideForUserRequestBody) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
// SetUser sets the user property value. 
func (m *UnhideForUserRequestBody) SetUser(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkUserIdentityable)() {
    if m != nil {
        m.user = value
    }
}
