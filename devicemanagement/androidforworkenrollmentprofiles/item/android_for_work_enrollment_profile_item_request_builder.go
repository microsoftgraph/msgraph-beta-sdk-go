package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5c8c19d77c389478909a5d5db267ec881c1ece4dbe130901468968c73c5b61f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidforworkenrollmentprofiles/item/createtoken"
    i9267d2002d296224b1f7223f20e956b623bc0d9efea1544156a9301683c48b4f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidforworkenrollmentprofiles/item/revoketoken"
)

// AndroidForWorkEnrollmentProfileItemRequestBuilder builds and executes requests for operations under \deviceManagement\androidForWorkEnrollmentProfiles\{androidForWorkEnrollmentProfile-id}
type AndroidForWorkEnrollmentProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AndroidForWorkEnrollmentProfileItemRequestBuilderDeleteOptions options for Delete
type AndroidForWorkEnrollmentProfileItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AndroidForWorkEnrollmentProfileItemRequestBuilderGetOptions options for Get
type AndroidForWorkEnrollmentProfileItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AndroidForWorkEnrollmentProfileItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AndroidForWorkEnrollmentProfileItemRequestBuilderGetQueryParameters android for Work enrollment profile entities.
type AndroidForWorkEnrollmentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AndroidForWorkEnrollmentProfileItemRequestBuilderPatchOptions options for Patch
type AndroidForWorkEnrollmentProfileItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidForWorkEnrollmentProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAndroidForWorkEnrollmentProfileItemRequestBuilderInternal instantiates a new AndroidForWorkEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAndroidForWorkEnrollmentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AndroidForWorkEnrollmentProfileItemRequestBuilder) {
    m := &AndroidForWorkEnrollmentProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidForWorkEnrollmentProfiles/{androidForWorkEnrollmentProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAndroidForWorkEnrollmentProfileItemRequestBuilder instantiates a new AndroidForWorkEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAndroidForWorkEnrollmentProfileItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AndroidForWorkEnrollmentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidForWorkEnrollmentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation android for Work enrollment profile entities.
func (m *AndroidForWorkEnrollmentProfileItemRequestBuilder) CreateDeleteRequestInformation(options *AndroidForWorkEnrollmentProfileItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation android for Work enrollment profile entities.
func (m *AndroidForWorkEnrollmentProfileItemRequestBuilder) CreateGetRequestInformation(options *AndroidForWorkEnrollmentProfileItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
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
// CreatePatchRequestInformation android for Work enrollment profile entities.
func (m *AndroidForWorkEnrollmentProfileItemRequestBuilder) CreatePatchRequestInformation(options *AndroidForWorkEnrollmentProfileItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
func (m *AndroidForWorkEnrollmentProfileItemRequestBuilder) CreateToken()(*i5c8c19d77c389478909a5d5db267ec881c1ece4dbe130901468968c73c5b61f0.CreateTokenRequestBuilder) {
    return i5c8c19d77c389478909a5d5db267ec881c1ece4dbe130901468968c73c5b61f0.NewCreateTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete android for Work enrollment profile entities.
func (m *AndroidForWorkEnrollmentProfileItemRequestBuilder) Delete(options *AndroidForWorkEnrollmentProfileItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get android for Work enrollment profile entities.
func (m *AndroidForWorkEnrollmentProfileItemRequestBuilder) Get(options *AndroidForWorkEnrollmentProfileItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidForWorkEnrollmentProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAndroidForWorkEnrollmentProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidForWorkEnrollmentProfile), nil
}
// Patch android for Work enrollment profile entities.
func (m *AndroidForWorkEnrollmentProfileItemRequestBuilder) Patch(options *AndroidForWorkEnrollmentProfileItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AndroidForWorkEnrollmentProfileItemRequestBuilder) RevokeToken()(*i9267d2002d296224b1f7223f20e956b623bc0d9efea1544156a9301683c48b4f.RevokeTokenRequestBuilder) {
    return i9267d2002d296224b1f7223f20e956b623bc0d9efea1544156a9301683c48b4f.NewRevokeTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
