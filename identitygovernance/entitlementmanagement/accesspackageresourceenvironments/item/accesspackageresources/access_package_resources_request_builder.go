package accesspackageresources

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i3f7642279120e0d794d7e3f12f09d33b860300b7436ef36b2cf11a810c4a1522 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresourceenvironments/item/accesspackageresources/ref"
)

// AccessPackageResourcesRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageResourceEnvironments\{accessPackageResourceEnvironment-id}\accessPackageResources
type AccessPackageResourcesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageResourcesRequestBuilderGetOptions options for Get
type AccessPackageResourcesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageResourcesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageResourcesRequestBuilderGetQueryParameters read-only. Required.
type AccessPackageResourcesRequestBuilderGetQueryParameters struct {
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
// NewAccessPackageResourcesRequestBuilderInternal instantiates a new AccessPackageResourcesRequestBuilder and sets the default values.
func NewAccessPackageResourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageResourcesRequestBuilder) {
    m := &AccessPackageResourcesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResourceEnvironments/{accessPackageResourceEnvironment_id}/accessPackageResources{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageResourcesRequestBuilder instantiates a new AccessPackageResourcesRequestBuilder and sets the default values.
func NewAccessPackageResourcesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageResourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageResourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation read-only. Required.
func (m *AccessPackageResourcesRequestBuilder) CreateGetRequestInformation(options *AccessPackageResourcesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get read-only. Required.
func (m *AccessPackageResourcesRequestBuilder) Get(options *AccessPackageResourcesRequestBuilderGetOptions)(*AccessPackageResourcesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourcesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*AccessPackageResourcesResponse), nil
}
func (m *AccessPackageResourcesRequestBuilder) Ref()(*i3f7642279120e0d794d7e3f12f09d33b860300b7436ef36b2cf11a810c4a1522.RefRequestBuilder) {
    return i3f7642279120e0d794d7e3f12f09d33b860300b7436ef36b2cf11a810c4a1522.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
