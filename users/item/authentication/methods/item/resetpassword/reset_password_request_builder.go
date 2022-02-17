package resetpassword

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ResetPasswordRequestBuilder builds and executes requests for operations under \users\{user-id}\authentication\methods\{authenticationMethod-id}\microsoft.graph.resetPassword
type ResetPasswordRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ResetPasswordRequestBuilderPostOptions options for Post
type ResetPasswordRequestBuilderPostOptions struct {
    // 
    Body *ResetPasswordRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ResetPasswordResponse union type wrapper for classes passwordResetResponse
type ResetPasswordResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type passwordResetResponse
    passwordResetResponse *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordResetResponse;
}
// NewResetPasswordResponse instantiates a new resetPasswordResponse and sets the default values.
func NewResetPasswordResponse()(*ResetPasswordResponse) {
    m := &ResetPasswordResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResetPasswordResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPasswordResetResponse gets the passwordResetResponse property value. Union type representation for type passwordResetResponse
func (m *ResetPasswordResponse) GetPasswordResetResponse()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordResetResponse) {
    if m == nil {
        return nil
    } else {
        return m.passwordResetResponse
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResetPasswordResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["passwordResetResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPasswordResetResponse() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordResetResponse(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordResetResponse))
        }
        return nil
    }
    return res
}
func (m *ResetPasswordResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ResetPasswordResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("passwordResetResponse", m.GetPasswordResetResponse())
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
func (m *ResetPasswordResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPasswordResetResponse sets the passwordResetResponse property value. Union type representation for type passwordResetResponse
func (m *ResetPasswordResponse) SetPasswordResetResponse(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordResetResponse)() {
    if m != nil {
        m.passwordResetResponse = value
    }
}
// NewResetPasswordRequestBuilderInternal instantiates a new ResetPasswordRequestBuilder and sets the default values.
func NewResetPasswordRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResetPasswordRequestBuilder) {
    m := &ResetPasswordRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/authentication/methods/{authenticationMethod_id}/microsoft.graph.resetPassword";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewResetPasswordRequestBuilder instantiates a new ResetPasswordRequestBuilder and sets the default values.
func NewResetPasswordRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResetPasswordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResetPasswordRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action resetPassword
func (m *ResetPasswordRequestBuilder) CreatePostRequestInformation(options *ResetPasswordRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action resetPassword
func (m *ResetPasswordRequestBuilder) Post(options *ResetPasswordRequestBuilderPostOptions)(*ResetPasswordResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResetPasswordResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ResetPasswordResponse), nil
}
