package homerealmdiscoverypolicies

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i411df080e4cf6f60ba796f18554d6cb880c43f585e0de4a61e3e92b4a9e5dbc6 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/homerealmdiscoverypolicies/count"
)

// HomeRealmDiscoveryPoliciesRequestBuilder provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.servicePrincipal entity.
type HomeRealmDiscoveryPoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// HomeRealmDiscoveryPoliciesRequestBuilderGetOptions options for Get
type HomeRealmDiscoveryPoliciesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *HomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// HomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters the homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
type HomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters struct {
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
// NewHomeRealmDiscoveryPoliciesRequestBuilderInternal instantiates a new HomeRealmDiscoveryPoliciesRequestBuilder and sets the default values.
func NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*HomeRealmDiscoveryPoliciesRequestBuilder) {
    m := &HomeRealmDiscoveryPoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}/homeRealmDiscoveryPolicies{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewHomeRealmDiscoveryPoliciesRequestBuilder instantiates a new HomeRealmDiscoveryPoliciesRequestBuilder and sets the default values.
func NewHomeRealmDiscoveryPoliciesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*HomeRealmDiscoveryPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *HomeRealmDiscoveryPoliciesRequestBuilder) Count()(*i411df080e4cf6f60ba796f18554d6cb880c43f585e0de4a61e3e92b4a9e5dbc6.CountRequestBuilder) {
    return i411df080e4cf6f60ba796f18554d6cb880c43f585e0de4a61e3e92b4a9e5dbc6.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
func (m *HomeRealmDiscoveryPoliciesRequestBuilder) CreateGetRequestInformation(options *HomeRealmDiscoveryPoliciesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
func (m *HomeRealmDiscoveryPoliciesRequestBuilder) Get(options *HomeRealmDiscoveryPoliciesRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.HomeRealmDiscoveryPolicyCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateHomeRealmDiscoveryPolicyCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.HomeRealmDiscoveryPolicyCollectionResponseable), nil
}
