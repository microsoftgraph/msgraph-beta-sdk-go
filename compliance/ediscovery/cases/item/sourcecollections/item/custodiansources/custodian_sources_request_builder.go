package custodiansources

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i016cdc841da321df87a610d6d28946c91e7f54e0d34ef26597a57074b9a7b341 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/custodiansources/ref"
)

// CustodianSourcesRequestBuilder builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\sourceCollections\{sourceCollection-id}\custodianSources
type CustodianSourcesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CustodianSourcesRequestBuilderGetOptions options for Get
type CustodianSourcesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CustodianSourcesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CustodianSourcesRequestBuilderGetQueryParameters custodian sources that are included in the sourceCollection.
type CustodianSourcesRequestBuilderGetQueryParameters struct {
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
// NewCustodianSourcesRequestBuilderInternal instantiates a new CustodianSourcesRequestBuilder and sets the default values.
func NewCustodianSourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustodianSourcesRequestBuilder) {
    m := &CustodianSourcesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/sourceCollections/{sourceCollection_id}/custodianSources{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCustodianSourcesRequestBuilder instantiates a new CustodianSourcesRequestBuilder and sets the default values.
func NewCustodianSourcesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustodianSourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustodianSourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation custodian sources that are included in the sourceCollection.
func (m *CustodianSourcesRequestBuilder) CreateGetRequestInformation(options *CustodianSourcesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get custodian sources that are included in the sourceCollection.
func (m *CustodianSourcesRequestBuilder) Get(options *CustodianSourcesRequestBuilderGetOptions)(*CustodianSourcesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustodianSourcesResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*CustodianSourcesResponse), nil
}
func (m *CustodianSourcesRequestBuilder) Ref()(*i016cdc841da321df87a610d6d28946c91e7f54e0d34ef26597a57074b9a7b341.RefRequestBuilder) {
    return i016cdc841da321df87a610d6d28946c91e7f54e0d34ef26597a57074b9a7b341.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
