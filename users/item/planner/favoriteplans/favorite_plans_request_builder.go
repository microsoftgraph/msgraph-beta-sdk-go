package favoriteplans

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ic123b10aaeed7f9ca1dd45407cc1fcc4fa72c511ff6d25f4e2b568a80be947be "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/favoriteplans/ref"
)

// FavoritePlansRequestBuilder builds and executes requests for operations under \users\{user-id}\planner\favoritePlans
type FavoritePlansRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// FavoritePlansRequestBuilderGetOptions options for Get
type FavoritePlansRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *FavoritePlansRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// FavoritePlansRequestBuilderGetQueryParameters read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
type FavoritePlansRequestBuilderGetQueryParameters struct {
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
// NewFavoritePlansRequestBuilderInternal instantiates a new FavoritePlansRequestBuilder and sets the default values.
func NewFavoritePlansRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FavoritePlansRequestBuilder) {
    m := &FavoritePlansRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/planner/favoritePlans{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFavoritePlansRequestBuilder instantiates a new FavoritePlansRequestBuilder and sets the default values.
func NewFavoritePlansRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FavoritePlansRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFavoritePlansRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
func (m *FavoritePlansRequestBuilder) CreateGetRequestInformation(options *FavoritePlansRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
func (m *FavoritePlansRequestBuilder) Get(options *FavoritePlansRequestBuilderGetOptions)(*FavoritePlansResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFavoritePlansResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*FavoritePlansResponse), nil
}
func (m *FavoritePlansRequestBuilder) Ref()(*ic123b10aaeed7f9ca1dd45407cc1fcc4fa72c511ff6d25f4e2b568a80be947be.RefRequestBuilder) {
    return ic123b10aaeed7f9ca1dd45407cc1fcc4fa72c511ff6d25f4e2b568a80be947be.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
