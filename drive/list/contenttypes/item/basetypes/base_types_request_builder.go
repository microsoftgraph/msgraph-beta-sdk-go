package basetypes

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i7a4d66a397bda0a98fc6a006dd6af855dfbe2235b5de88b21c57e23b5db40bda "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/basetypes/getcompatiblehubcontenttypes"
    i89b4240539243c923c9e17e1d8a7018574ed21085c273fdbb32b24996be44f66 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/basetypes/ref"
    ic1281fad5e9c0c6bd0f84546975afaf8094ff6ce99db9feb709dc6cf43c44799 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/basetypes/addcopy"
    ic4a09d4ee40b63294a3784f6fab1f9daa68eba54e38e447292036ee57d3342b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/basetypes/addcopyfromcontenttypehub"
)

// BaseTypesRequestBuilder builds and executes requests for operations under \drive\list\contentTypes\{contentType-id}\baseTypes
type BaseTypesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BaseTypesRequestBuilderGetOptions options for Get
type BaseTypesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BaseTypesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BaseTypesRequestBuilderGetQueryParameters the collection of content types that are ancestors of this content type.
type BaseTypesRequestBuilderGetQueryParameters struct {
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
func (m *BaseTypesRequestBuilder) AddCopy()(*ic1281fad5e9c0c6bd0f84546975afaf8094ff6ce99db9feb709dc6cf43c44799.AddCopyRequestBuilder) {
    return ic1281fad5e9c0c6bd0f84546975afaf8094ff6ce99db9feb709dc6cf43c44799.NewAddCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTypesRequestBuilder) AddCopyFromContentTypeHub()(*ic4a09d4ee40b63294a3784f6fab1f9daa68eba54e38e447292036ee57d3342b4.AddCopyFromContentTypeHubRequestBuilder) {
    return ic4a09d4ee40b63294a3784f6fab1f9daa68eba54e38e447292036ee57d3342b4.NewAddCopyFromContentTypeHubRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBaseTypesRequestBuilderInternal instantiates a new BaseTypesRequestBuilder and sets the default values.
func NewBaseTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTypesRequestBuilder) {
    m := &BaseTypesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/list/contentTypes/{contentType_id}/baseTypes{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBaseTypesRequestBuilder instantiates a new BaseTypesRequestBuilder and sets the default values.
func NewBaseTypesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBaseTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the collection of content types that are ancestors of this content type.
func (m *BaseTypesRequestBuilder) CreateGetRequestInformation(options *BaseTypesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the collection of content types that are ancestors of this content type.
func (m *BaseTypesRequestBuilder) Get(options *BaseTypesRequestBuilderGetOptions)(*BaseTypesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBaseTypesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*BaseTypesResponse), nil
}
// GetCompatibleHubContentTypes builds and executes requests for operations under \drive\list\contentTypes\{contentType-id}\baseTypes\microsoft.graph.getCompatibleHubContentTypes()
func (m *BaseTypesRequestBuilder) GetCompatibleHubContentTypes()(*i7a4d66a397bda0a98fc6a006dd6af855dfbe2235b5de88b21c57e23b5db40bda.GetCompatibleHubContentTypesRequestBuilder) {
    return i7a4d66a397bda0a98fc6a006dd6af855dfbe2235b5de88b21c57e23b5db40bda.NewGetCompatibleHubContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTypesRequestBuilder) Ref()(*i89b4240539243c923c9e17e1d8a7018574ed21085c273fdbb32b24996be44f66.RefRequestBuilder) {
    return i89b4240539243c923c9e17e1d8a7018574ed21085c273fdbb32b24996be44f66.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
