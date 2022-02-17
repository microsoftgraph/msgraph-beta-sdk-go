package userflowidentityproviders

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i70f436469490863988a889d70e14c401468a62c692956cd713796e412d3d6b4e "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userflowidentityproviders/ref"
    i8ab7bf9c75474f49faf98dbb1be0651a168ad02f853ea1051c2d43fba4af072b "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item/userflowidentityproviders/availableprovidertypes"
)

// UserFlowIdentityProvidersRequestBuilder builds and executes requests for operations under \identity\b2cUserFlows\{b2cIdentityUserFlow-id}\userFlowIdentityProviders
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
// AvailableProviderTypes builds and executes requests for operations under \identity\b2cUserFlows\{b2cIdentityUserFlow-id}\userFlowIdentityProviders\microsoft.graph.availableProviderTypes()
func (m *UserFlowIdentityProvidersRequestBuilder) AvailableProviderTypes()(*i8ab7bf9c75474f49faf98dbb1be0651a168ad02f853ea1051c2d43fba4af072b.AvailableProviderTypesRequestBuilder) {
    return i8ab7bf9c75474f49faf98dbb1be0651a168ad02f853ea1051c2d43fba4af072b.NewAvailableProviderTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserFlowIdentityProvidersRequestBuilderInternal instantiates a new UserFlowIdentityProvidersRequestBuilder and sets the default values.
func NewUserFlowIdentityProvidersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserFlowIdentityProvidersRequestBuilder) {
    m := &UserFlowIdentityProvidersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow_id}/userFlowIdentityProviders{?top,skip,search,filter,count,orderby,select,expand}";
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
func (m *UserFlowIdentityProvidersRequestBuilder) Ref()(*i70f436469490863988a889d70e14c401468a62c692956cd713796e412d3d6b4e.RefRequestBuilder) {
    return i70f436469490863988a889d70e14c401468a62c692956cd713796e412d3d6b4e.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
