package transitivereports

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i93e6244a8f9ec960e02c3fce9c31eb4d0ee1153ccd00153e3ece9dbf137c389a "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivereports/ref"
)

// TransitiveReportsRequestBuilder builds and executes requests for operations under \contacts\{orgContact-id}\transitiveReports
type TransitiveReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TransitiveReportsRequestBuilderGetOptions options for Get
type TransitiveReportsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TransitiveReportsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TransitiveReportsRequestBuilderGetQueryParameters the transitive reports for a contact. Read-only.
type TransitiveReportsRequestBuilderGetQueryParameters struct {
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
// NewTransitiveReportsRequestBuilderInternal instantiates a new TransitiveReportsRequestBuilder and sets the default values.
func NewTransitiveReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TransitiveReportsRequestBuilder) {
    m := &TransitiveReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts/{orgContact_id}/transitiveReports{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTransitiveReportsRequestBuilder instantiates a new TransitiveReportsRequestBuilder and sets the default values.
func NewTransitiveReportsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TransitiveReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTransitiveReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the transitive reports for a contact. Read-only.
func (m *TransitiveReportsRequestBuilder) CreateGetRequestInformation(options *TransitiveReportsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the transitive reports for a contact. Read-only.
func (m *TransitiveReportsRequestBuilder) Get(options *TransitiveReportsRequestBuilderGetOptions)(*TransitiveReportsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTransitiveReportsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*TransitiveReportsResponse), nil
}
func (m *TransitiveReportsRequestBuilder) Ref()(*i93e6244a8f9ec960e02c3fce9c31eb4d0ee1153ccd00153e3ece9dbf137c389a.RefRequestBuilder) {
    return i93e6244a8f9ec960e02c3fce9c31eb4d0ee1153ccd00153e3ece9dbf137c389a.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
