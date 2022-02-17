package userflowidentityproviders

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i69f3e5cfc271fa5f16a77319cdecf0ad649be03e57660b497481f2902064d8f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/userflowidentityproviders/availableprovidertypes"
    ic5c1a88dc5fc77efaac004a516d8984d4cd2f0c4875f51c372a13b2debc9f05e "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item/userflowidentityproviders/ref"
)

// UserFlowIdentityProvidersRequestBuilder builds and executes requests for operations under \identity\b2xUserFlows\{b2xIdentityUserFlow-id}\userFlowIdentityProviders
type UserFlowIdentityProvidersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserFlowIdentityProvidersRequestBuilderGetOptions options for Get
type UserFlowIdentityProvidersRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserFlowIdentityProvidersRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserFlowIdentityProvidersRequestBuilderGetQueryParameters get userFlowIdentityProviders from identity
type UserFlowIdentityProvidersRequestBuilderGetQueryParameters struct {
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
// AvailableProviderTypes builds and executes requests for operations under \identity\b2xUserFlows\{b2xIdentityUserFlow-id}\userFlowIdentityProviders\microsoft.graph.availableProviderTypes()
func (m *UserFlowIdentityProvidersRequestBuilder) AvailableProviderTypes()(*i69f3e5cfc271fa5f16a77319cdecf0ad649be03e57660b497481f2902064d8f5.AvailableProviderTypesRequestBuilder) {
    return i69f3e5cfc271fa5f16a77319cdecf0ad649be03e57660b497481f2902064d8f5.NewAvailableProviderTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserFlowIdentityProvidersRequestBuilderInternal instantiates a new UserFlowIdentityProvidersRequestBuilder and sets the default values.
func NewUserFlowIdentityProvidersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserFlowIdentityProvidersRequestBuilder) {
    m := &UserFlowIdentityProvidersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow_id}/userFlowIdentityProviders{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserFlowIdentityProvidersRequestBuilder instantiates a new UserFlowIdentityProvidersRequestBuilder and sets the default values.
func NewUserFlowIdentityProvidersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserFlowIdentityProvidersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserFlowIdentityProvidersRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get userFlowIdentityProviders from identity
func (m *UserFlowIdentityProvidersRequestBuilder) CreateGetRequestInformation(options *UserFlowIdentityProvidersRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get userFlowIdentityProviders from identity
func (m *UserFlowIdentityProvidersRequestBuilder) Get(options *UserFlowIdentityProvidersRequestBuilderGetOptions)(*UserFlowIdentityProvidersResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserFlowIdentityProvidersResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*UserFlowIdentityProvidersResponse), nil
}
func (m *UserFlowIdentityProvidersRequestBuilder) Ref()(*ic5c1a88dc5fc77efaac004a516d8984d4cd2f0c4875f51c372a13b2debc9f05e.RefRequestBuilder) {
    return ic5c1a88dc5fc77efaac004a516d8984d4cd2f0c4875f51c372a13b2debc9f05e.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
