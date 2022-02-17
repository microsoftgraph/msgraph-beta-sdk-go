package validatepassword

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ValidatePasswordRequestBuilder builds and executes requests for operations under \users\microsoft.graph.validatePassword
type ValidatePasswordRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ValidatePasswordRequestBuilderPostOptions options for Post
type ValidatePasswordRequestBuilderPostOptions struct {
    // 
    Body *ValidatePasswordRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ValidatePasswordResponse union type wrapper for classes passwordValidationInformation
type ValidatePasswordResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type passwordValidationInformation
    passwordValidationInformation *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordValidationInformation;
}
// NewValidatePasswordResponse instantiates a new validatePasswordResponse and sets the default values.
func NewValidatePasswordResponse()(*ValidatePasswordResponse) {
    m := &ValidatePasswordResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidatePasswordResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPasswordValidationInformation gets the passwordValidationInformation property value. Union type representation for type passwordValidationInformation
func (m *ValidatePasswordResponse) GetPasswordValidationInformation()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordValidationInformation) {
    if m == nil {
        return nil
    } else {
        return m.passwordValidationInformation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidatePasswordResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["passwordValidationInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPasswordValidationInformation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordValidationInformation(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordValidationInformation))
        }
        return nil
    }
    return res
}
func (m *ValidatePasswordResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ValidatePasswordResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("passwordValidationInformation", m.GetPasswordValidationInformation())
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
func (m *ValidatePasswordResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPasswordValidationInformation sets the passwordValidationInformation property value. Union type representation for type passwordValidationInformation
func (m *ValidatePasswordResponse) SetPasswordValidationInformation(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordValidationInformation)() {
    if m != nil {
        m.passwordValidationInformation = value
    }
}
// NewValidatePasswordRequestBuilderInternal instantiates a new ValidatePasswordRequestBuilder and sets the default values.
func NewValidatePasswordRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ValidatePasswordRequestBuilder) {
    m := &ValidatePasswordRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/microsoft.graph.validatePassword";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewValidatePasswordRequestBuilder instantiates a new ValidatePasswordRequestBuilder and sets the default values.
func NewValidatePasswordRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ValidatePasswordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewValidatePasswordRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action validatePassword
func (m *ValidatePasswordRequestBuilder) CreatePostRequestInformation(options *ValidatePasswordRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action validatePassword
func (m *ValidatePasswordRequestBuilder) Post(options *ValidatePasswordRequestBuilderPostOptions)(*ValidatePasswordResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewValidatePasswordResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ValidatePasswordResponse), nil
}
