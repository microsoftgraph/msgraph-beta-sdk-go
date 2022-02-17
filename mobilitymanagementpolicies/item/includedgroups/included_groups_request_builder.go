package includedgroups

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i44a47714785ba6f6fac24c940ef2e4dbf135b31a210d29741f4c475a8c816bd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/mobilitymanagementpolicies/item/includedgroups/ref"
)

// IncludedGroupsRequestBuilder builds and executes requests for operations under \mobilityManagementPolicies\{mobilityManagementPolicy-id}\includedGroups
type IncludedGroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IncludedGroupsRequestBuilderGetOptions options for Get
type IncludedGroupsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *IncludedGroupsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IncludedGroupsRequestBuilderGetQueryParameters azure AD groups under the scope of the mobility management application if appliesTo is selected
type IncludedGroupsRequestBuilderGetQueryParameters struct {
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
// NewIncludedGroupsRequestBuilderInternal instantiates a new IncludedGroupsRequestBuilder and sets the default values.
func NewIncludedGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IncludedGroupsRequestBuilder) {
    m := &IncludedGroupsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/mobilityManagementPolicies/{mobilityManagementPolicy_id}/includedGroups{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIncludedGroupsRequestBuilder instantiates a new IncludedGroupsRequestBuilder and sets the default values.
func NewIncludedGroupsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IncludedGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIncludedGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation azure AD groups under the scope of the mobility management application if appliesTo is selected
func (m *IncludedGroupsRequestBuilder) CreateGetRequestInformation(options *IncludedGroupsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get azure AD groups under the scope of the mobility management application if appliesTo is selected
func (m *IncludedGroupsRequestBuilder) Get(options *IncludedGroupsRequestBuilderGetOptions)(*IncludedGroupsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIncludedGroupsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*IncludedGroupsResponse), nil
}
func (m *IncludedGroupsRequestBuilder) Ref()(*i44a47714785ba6f6fac24c940ef2e4dbf135b31a210d29741f4c475a8c816bd1.RefRequestBuilder) {
    return i44a47714785ba6f6fac24c940ef2e4dbf135b31a210d29741f4c475a8c816bd1.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
