package recentplans

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i32bb13a0096ebe13f36b06bcec6b204cfd1b137d6a7e383561311fbbed92c693 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/recentplans/ref"
)

// RecentPlansRequestBuilder builds and executes requests for operations under \me\planner\recentPlans
type RecentPlansRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RecentPlansRequestBuilderGetOptions options for Get
type RecentPlansRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RecentPlansRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RecentPlansRequestBuilderGetQueryParameters read-only. Nullable. Returns the plannerPlans that have been recently viewed by the user in apps that support recent plans.
type RecentPlansRequestBuilderGetQueryParameters struct {
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
// NewRecentPlansRequestBuilderInternal instantiates a new RecentPlansRequestBuilder and sets the default values.
func NewRecentPlansRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RecentPlansRequestBuilder) {
    m := &RecentPlansRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/planner/recentPlans{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRecentPlansRequestBuilder instantiates a new RecentPlansRequestBuilder and sets the default values.
func NewRecentPlansRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RecentPlansRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRecentPlansRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation read-only. Nullable. Returns the plannerPlans that have been recently viewed by the user in apps that support recent plans.
func (m *RecentPlansRequestBuilder) CreateGetRequestInformation(options *RecentPlansRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get read-only. Nullable. Returns the plannerPlans that have been recently viewed by the user in apps that support recent plans.
func (m *RecentPlansRequestBuilder) Get(options *RecentPlansRequestBuilderGetOptions)(*RecentPlansResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecentPlansResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*RecentPlansResponse), nil
}
func (m *RecentPlansRequestBuilder) Ref()(*i32bb13a0096ebe13f36b06bcec6b204cfd1b137d6a7e383561311fbbed92c693.RefRequestBuilder) {
    return i32bb13a0096ebe13f36b06bcec6b204cfd1b137d6a7e383561311fbbed92c693.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
