package createpasswordsinglesignoncredentials

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type CreatePasswordSingleSignOnCredentialsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type CreatePasswordSingleSignOnCredentialsResponse struct {
    additionalData map[string]interface{};
    passwordSingleSignOnCredentialSet *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordSingleSignOnCredentialSet;
}
func NewCreatePasswordSingleSignOnCredentialsResponse()(*CreatePasswordSingleSignOnCredentialsResponse) {
    m := &CreatePasswordSingleSignOnCredentialsResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CreatePasswordSingleSignOnCredentialsResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CreatePasswordSingleSignOnCredentialsResponse) GetPasswordSingleSignOnCredentialSet()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordSingleSignOnCredentialSet) {
    if m == nil {
        return nil
    } else {
        return m.passwordSingleSignOnCredentialSet
    }
}
func (m *CreatePasswordSingleSignOnCredentialsResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["passwordSingleSignOnCredentialSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPasswordSingleSignOnCredentialSet() })
        if err != nil {
            return err
        }
        m.SetPasswordSingleSignOnCredentialSet(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordSingleSignOnCredentialSet))
        return nil
    }
    return res
}
func (m *CreatePasswordSingleSignOnCredentialsResponse) IsNil()(bool) {
    return m == nil
}
func (m *CreatePasswordSingleSignOnCredentialsResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("passwordSingleSignOnCredentialSet", m.GetPasswordSingleSignOnCredentialSet())
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
func (m *CreatePasswordSingleSignOnCredentialsResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CreatePasswordSingleSignOnCredentialsResponse) SetPasswordSingleSignOnCredentialSet(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordSingleSignOnCredentialSet)() {
    m.passwordSingleSignOnCredentialSet = value
}
func NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreatePasswordSingleSignOnCredentialsRequestBuilder) {
    m := &CreatePasswordSingleSignOnCredentialsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/servicePrincipals/{servicePrincipal_id}/microsoft.graph.createPasswordSingleSignOnCredentials";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewCreatePasswordSingleSignOnCredentialsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreatePasswordSingleSignOnCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *CreatePasswordSingleSignOnCredentialsRequestBuilder) CreatePostRequestInformation(body *CreatePasswordSingleSignOnCredentialsRequestBody, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *CreatePasswordSingleSignOnCredentialsRequestBuilder) Post(body *CreatePasswordSingleSignOnCredentialsRequestBody, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*CreatePasswordSingleSignOnCredentialsResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCreatePasswordSingleSignOnCredentialsResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*CreatePasswordSingleSignOnCredentialsResponse), nil
}
