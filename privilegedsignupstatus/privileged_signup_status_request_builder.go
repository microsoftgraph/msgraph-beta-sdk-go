package privilegedsignupstatus

import (
    i505eff958ea7946e725489038a9fbe85c297626b944e12d9c6f68697f1a65258 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/issignedup"
    i576c2492be4604b88fea75e47108c8bd8db61a3de84b266733d2d7849bec46eb "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/cansignup"
    ib823f830e3ea0500a214b973cf170985ed37b35d907e2b81e06df8aaa8f66256 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/completesetup"
    ibe660ff6aa6da80e38b282109151294b9409fe46a692bb37484dd8a1705b5a94 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/signup"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// Builds and executes requests for operations under \privilegedSignupStatus
type PrivilegedSignupStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type PrivilegedSignupStatusRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrivilegedSignupStatusRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entities from privilegedSignupStatus
type PrivilegedSignupStatusRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// Options for Post
type PrivilegedSignupStatusRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedSignupStatus;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Builds and executes requests for operations under \privilegedSignupStatus\microsoft.graph.canSignUp()
func (m *PrivilegedSignupStatusRequestBuilder) CanSignUp()(*i576c2492be4604b88fea75e47108c8bd8db61a3de84b266733d2d7849bec46eb.CanSignUpRequestBuilder) {
    return i576c2492be4604b88fea75e47108c8bd8db61a3de84b266733d2d7849bec46eb.NewCanSignUpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedSignupStatusRequestBuilder) CompleteSetup()(*ib823f830e3ea0500a214b973cf170985ed37b35d907e2b81e06df8aaa8f66256.CompleteSetupRequestBuilder) {
    return ib823f830e3ea0500a214b973cf170985ed37b35d907e2b81e06df8aaa8f66256.NewCompleteSetupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new PrivilegedSignupStatusRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrivilegedSignupStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedSignupStatusRequestBuilder) {
    m := &PrivilegedSignupStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedSignupStatus{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new PrivilegedSignupStatusRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrivilegedSignupStatusRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedSignupStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedSignupStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Get entities from privilegedSignupStatus
// Parameters:
//  - options : Options for the request
func (m *PrivilegedSignupStatusRequestBuilder) CreateGetRequestInformation(options *PrivilegedSignupStatusRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Add new entity to privilegedSignupStatus
// Parameters:
//  - options : Options for the request
func (m *PrivilegedSignupStatusRequestBuilder) CreatePostRequestInformation(options *PrivilegedSignupStatusRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get entities from privilegedSignupStatus
// Parameters:
//  - options : Options for the request
func (m *PrivilegedSignupStatusRequestBuilder) Get(options *PrivilegedSignupStatusRequestBuilderGetOptions)(*PrivilegedSignupStatusResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedSignupStatusResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*PrivilegedSignupStatusResponse), nil
}
// Builds and executes requests for operations under \privilegedSignupStatus\microsoft.graph.isSignedUp()
func (m *PrivilegedSignupStatusRequestBuilder) IsSignedUp()(*i505eff958ea7946e725489038a9fbe85c297626b944e12d9c6f68697f1a65258.IsSignedUpRequestBuilder) {
    return i505eff958ea7946e725489038a9fbe85c297626b944e12d9c6f68697f1a65258.NewIsSignedUpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Add new entity to privilegedSignupStatus
// Parameters:
//  - options : Options for the request
func (m *PrivilegedSignupStatusRequestBuilder) Post(options *PrivilegedSignupStatusRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedSignupStatus, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrivilegedSignupStatus() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedSignupStatus), nil
}
func (m *PrivilegedSignupStatusRequestBuilder) SignUp()(*ibe660ff6aa6da80e38b282109151294b9409fe46a692bb37484dd8a1705b5a94.SignUpRequestBuilder) {
    return ibe660ff6aa6da80e38b282109151294b9409fe46a692bb37484dd8a1705b5a94.NewSignUpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
