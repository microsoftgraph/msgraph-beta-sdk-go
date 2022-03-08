package appconsentrequestsforapproval

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i53c3b6bd8769d2980538a78d12d5f10d790e2f3ee9f946ae1c82d2ae13c1304d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/appconsentrequestsforapproval/count"
    iff9cf0d04c2f0688104b77d1c264e925cb4c0845c53ab4fdb05464cd52ec1f18 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/appconsentrequestsforapproval/filterbycurrentuserwithon"
)

// AppConsentRequestsForApprovalRequestBuilder provides operations to manage the appConsentRequestsForApproval property of the microsoft.graph.user entity.
type AppConsentRequestsForApprovalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AppConsentRequestsForApprovalRequestBuilderGetOptions options for Get
type AppConsentRequestsForApprovalRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AppConsentRequestsForApprovalRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AppConsentRequestsForApprovalRequestBuilderGetQueryParameters get appConsentRequestsForApproval from me
type AppConsentRequestsForApprovalRequestBuilderGetQueryParameters struct {
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
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// AppConsentRequestsForApprovalRequestBuilderPostOptions options for Post
type AppConsentRequestsForApprovalRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppConsentRequestable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAppConsentRequestsForApprovalRequestBuilderInternal instantiates a new AppConsentRequestsForApprovalRequestBuilder and sets the default values.
func NewAppConsentRequestsForApprovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AppConsentRequestsForApprovalRequestBuilder) {
    m := &AppConsentRequestsForApprovalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/appConsentRequestsForApproval{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAppConsentRequestsForApprovalRequestBuilder instantiates a new AppConsentRequestsForApprovalRequestBuilder and sets the default values.
func NewAppConsentRequestsForApprovalRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AppConsentRequestsForApprovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppConsentRequestsForApprovalRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AppConsentRequestsForApprovalRequestBuilder) Count()(*i53c3b6bd8769d2980538a78d12d5f10d790e2f3ee9f946ae1c82d2ae13c1304d.CountRequestBuilder) {
    return i53c3b6bd8769d2980538a78d12d5f10d790e2f3ee9f946ae1c82d2ae13c1304d.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get appConsentRequestsForApproval from me
func (m *AppConsentRequestsForApprovalRequestBuilder) CreateGetRequestInformation(options *AppConsentRequestsForApprovalRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to appConsentRequestsForApproval for me
func (m *AppConsentRequestsForApprovalRequestBuilder) CreatePostRequestInformation(options *AppConsentRequestsForApprovalRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
func (m *AppConsentRequestsForApprovalRequestBuilder) FilterByCurrentUserWithOn(on *string)(*iff9cf0d04c2f0688104b77d1c264e925cb4c0845c53ab4fdb05464cd52ec1f18.FilterByCurrentUserWithOnRequestBuilder) {
    return iff9cf0d04c2f0688104b77d1c264e925cb4c0845c53ab4fdb05464cd52ec1f18.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
}
// Get get appConsentRequestsForApproval from me
func (m *AppConsentRequestsForApprovalRequestBuilder) Get(options *AppConsentRequestsForApprovalRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppConsentRequestCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateAppConsentRequestCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppConsentRequestCollectionResponseable), nil
}
// Post create new navigation property to appConsentRequestsForApproval for me
func (m *AppConsentRequestsForApprovalRequestBuilder) Post(options *AppConsentRequestsForApprovalRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppConsentRequestable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateAppConsentRequestFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppConsentRequestable), nil
}
