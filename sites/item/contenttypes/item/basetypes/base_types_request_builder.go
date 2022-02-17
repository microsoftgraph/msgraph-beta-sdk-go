package basetypes

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1353dc1dfbbb16e821a2606b6de6cf4586e7d5667be330d8dac349b8ea9b8815 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/basetypes/getcompatiblehubcontenttypes"
    i3b7930a8ec3b2b0639ff62905f2bf67c64731ca16abe9f10959e7a77bd308045 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/basetypes/addcopyfromcontenttypehub"
    i5eb56b072586e1f7b99dfc9e2264b730e82a6f93d64de6579bbfdb68acdf8839 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/basetypes/addcopy"
    if7af84e45b3bdced0fec4830c03eba10bd0c52d30ff71cfb49de8ac66b288594 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/basetypes/ref"
)

// BaseTypesRequestBuilder builds and executes requests for operations under \sites\{site-id}\contentTypes\{contentType-id}\baseTypes
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
func (m *BaseTypesRequestBuilder) AddCopy()(*i5eb56b072586e1f7b99dfc9e2264b730e82a6f93d64de6579bbfdb68acdf8839.AddCopyRequestBuilder) {
    return i5eb56b072586e1f7b99dfc9e2264b730e82a6f93d64de6579bbfdb68acdf8839.NewAddCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTypesRequestBuilder) AddCopyFromContentTypeHub()(*i3b7930a8ec3b2b0639ff62905f2bf67c64731ca16abe9f10959e7a77bd308045.AddCopyFromContentTypeHubRequestBuilder) {
    return i3b7930a8ec3b2b0639ff62905f2bf67c64731ca16abe9f10959e7a77bd308045.NewAddCopyFromContentTypeHubRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBaseTypesRequestBuilderInternal instantiates a new BaseTypesRequestBuilder and sets the default values.
func NewBaseTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTypesRequestBuilder) {
    m := &BaseTypesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/contentTypes/{contentType_id}/baseTypes{?top,skip,search,filter,count,orderby,select,expand}";
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
// GetCompatibleHubContentTypes builds and executes requests for operations under \sites\{site-id}\contentTypes\{contentType-id}\baseTypes\microsoft.graph.getCompatibleHubContentTypes()
func (m *BaseTypesRequestBuilder) GetCompatibleHubContentTypes()(*i1353dc1dfbbb16e821a2606b6de6cf4586e7d5667be330d8dac349b8ea9b8815.GetCompatibleHubContentTypesRequestBuilder) {
    return i1353dc1dfbbb16e821a2606b6de6cf4586e7d5667be330d8dac349b8ea9b8815.NewGetCompatibleHubContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTypesRequestBuilder) Ref()(*if7af84e45b3bdced0fec4830c03eba10bd0c52d30ff71cfb49de8ac66b288594.RefRequestBuilder) {
    return if7af84e45b3bdced0fec4830c03eba10bd0c52d30ff71cfb49de8ac66b288594.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
