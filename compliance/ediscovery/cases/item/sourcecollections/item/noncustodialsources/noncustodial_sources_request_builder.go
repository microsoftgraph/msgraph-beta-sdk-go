package noncustodialsources

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i127ed970f4cf70dae4e3a9bc9ae48112e65996c5d47c39e609c0b55f44fc2ecb "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/noncustodialsources/removehold"
    i2b53e8d32167b7b8d9dbd842472314407904e34f8d9feeb85454bc41d81e5372 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/noncustodialsources/ref"
    ia78d53705d4e84211433969df7a4fd05268f2735ae29bd62c0e068728055ed5f "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/noncustodialsources/applyhold"
)

// NoncustodialSourcesRequestBuilder builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\sourceCollections\{sourceCollection-id}\noncustodialSources
type NoncustodialSourcesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NoncustodialSourcesRequestBuilderGetOptions options for Get
type NoncustodialSourcesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *NoncustodialSourcesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NoncustodialSourcesRequestBuilderGetQueryParameters noncustodialDataSource sources that are included in the sourceCollection
type NoncustodialSourcesRequestBuilderGetQueryParameters struct {
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
func (m *NoncustodialSourcesRequestBuilder) ApplyHold()(*ia78d53705d4e84211433969df7a4fd05268f2735ae29bd62c0e068728055ed5f.ApplyHoldRequestBuilder) {
    return ia78d53705d4e84211433969df7a4fd05268f2735ae29bd62c0e068728055ed5f.NewApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewNoncustodialSourcesRequestBuilderInternal instantiates a new NoncustodialSourcesRequestBuilder and sets the default values.
func NewNoncustodialSourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NoncustodialSourcesRequestBuilder) {
    m := &NoncustodialSourcesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/sourceCollections/{sourceCollection_id}/noncustodialSources{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewNoncustodialSourcesRequestBuilder instantiates a new NoncustodialSourcesRequestBuilder and sets the default values.
func NewNoncustodialSourcesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NoncustodialSourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNoncustodialSourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation noncustodialDataSource sources that are included in the sourceCollection
func (m *NoncustodialSourcesRequestBuilder) CreateGetRequestInformation(options *NoncustodialSourcesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get noncustodialDataSource sources that are included in the sourceCollection
func (m *NoncustodialSourcesRequestBuilder) Get(options *NoncustodialSourcesRequestBuilderGetOptions)(*NoncustodialSourcesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNoncustodialSourcesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*NoncustodialSourcesResponse), nil
}
func (m *NoncustodialSourcesRequestBuilder) Ref()(*i2b53e8d32167b7b8d9dbd842472314407904e34f8d9feeb85454bc41d81e5372.RefRequestBuilder) {
    return i2b53e8d32167b7b8d9dbd842472314407904e34f8d9feeb85454bc41d81e5372.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *NoncustodialSourcesRequestBuilder) RemoveHold()(*i127ed970f4cf70dae4e3a9bc9ae48112e65996c5d47c39e609c0b55f44fc2ecb.RemoveHoldRequestBuilder) {
    return i127ed970f4cf70dae4e3a9bc9ae48112e65996c5d47c39e609c0b55f44fc2ecb.NewRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
