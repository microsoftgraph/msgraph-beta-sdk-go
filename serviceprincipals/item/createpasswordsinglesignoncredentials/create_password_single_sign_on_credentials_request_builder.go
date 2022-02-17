package createpasswordsinglesignoncredentials

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// CreatePasswordSingleSignOnCredentialsRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\microsoft.graph.createPasswordSingleSignOnCredentials
type CreatePasswordSingleSignOnCredentialsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CreatePasswordSingleSignOnCredentialsRequestBuilderPostOptions options for Post
type CreatePasswordSingleSignOnCredentialsRequestBuilderPostOptions struct {
    // 
    Body *CreatePasswordSingleSignOnCredentialsRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CreatePasswordSingleSignOnCredentialsResponse union type wrapper for classes passwordSingleSignOnCredentialSet
type CreatePasswordSingleSignOnCredentialsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type passwordSingleSignOnCredentialSet
    passwordSingleSignOnCredentialSet *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordSingleSignOnCredentialSet;
}
// NewCreatePasswordSingleSignOnCredentialsResponse instantiates a new createPasswordSingleSignOnCredentialsResponse and sets the default values.
func NewCreatePasswordSingleSignOnCredentialsResponse()(*CreatePasswordSingleSignOnCredentialsResponse) {
    m := &CreatePasswordSingleSignOnCredentialsResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreatePasswordSingleSignOnCredentialsResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPasswordSingleSignOnCredentialSet gets the passwordSingleSignOnCredentialSet property value. Union type representation for type passwordSingleSignOnCredentialSet
func (m *CreatePasswordSingleSignOnCredentialsResponse) GetPasswordSingleSignOnCredentialSet()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordSingleSignOnCredentialSet) {
    if m == nil {
        return nil
    } else {
        return m.passwordSingleSignOnCredentialSet
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreatePasswordSingleSignOnCredentialsResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["passwordSingleSignOnCredentialSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPasswordSingleSignOnCredentialSet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordSingleSignOnCredentialSet(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordSingleSignOnCredentialSet))
        }
        return nil
    }
    return res
}
func (m *CreatePasswordSingleSignOnCredentialsResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreatePasswordSingleSignOnCredentialsResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPasswordSingleSignOnCredentialSet sets the passwordSingleSignOnCredentialSet property value. Union type representation for type passwordSingleSignOnCredentialSet
func (m *CreatePasswordSingleSignOnCredentialsResponse) SetPasswordSingleSignOnCredentialSet(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordSingleSignOnCredentialSet)() {
    if m != nil {
        m.passwordSingleSignOnCredentialSet = value
    }
}
// NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal instantiates a new CreatePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreatePasswordSingleSignOnCredentialsRequestBuilder) {
    m := &CreatePasswordSingleSignOnCredentialsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}/microsoft.graph.createPasswordSingleSignOnCredentials";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreatePasswordSingleSignOnCredentialsRequestBuilder instantiates a new CreatePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewCreatePasswordSingleSignOnCredentialsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreatePasswordSingleSignOnCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action createPasswordSingleSignOnCredentials
func (m *CreatePasswordSingleSignOnCredentialsRequestBuilder) CreatePostRequestInformation(options *CreatePasswordSingleSignOnCredentialsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post invoke action createPasswordSingleSignOnCredentials
func (m *CreatePasswordSingleSignOnCredentialsRequestBuilder) Post(options *CreatePasswordSingleSignOnCredentialsRequestBuilderPostOptions)(*CreatePasswordSingleSignOnCredentialsResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCreatePasswordSingleSignOnCredentialsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*CreatePasswordSingleSignOnCredentialsResponse), nil
}
