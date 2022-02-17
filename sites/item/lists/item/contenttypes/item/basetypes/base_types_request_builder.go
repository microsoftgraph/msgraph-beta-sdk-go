package basetypes

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0d5e866a8d2f206f204617c8afe6a61bd6e740095f4d6adcb003c9e5bfb86591 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/basetypes/getcompatiblehubcontenttypes"
    i1fb6c90038e174f9bda7e7123feeeafbc289c60224a9133d157b5943665d3c85 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/basetypes/addcopyfromcontenttypehub"
    i9b2fd694658511b1f9c73acd3c794477af0c97fce6079052b2f40bfc1363ad10 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/basetypes/ref"
    ie17d2b7dd4022f93c2c71a67b7a20a78bffcf7d8ccf29ebcee200a1107a650ad "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/item/basetypes/addcopy"
)

// BaseTypesRequestBuilder builds and executes requests for operations under \sites\{site-id}\lists\{list-id}\contentTypes\{contentType-id}\baseTypes
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
func (m *BaseTypesRequestBuilder) AddCopy()(*ie17d2b7dd4022f93c2c71a67b7a20a78bffcf7d8ccf29ebcee200a1107a650ad.AddCopyRequestBuilder) {
    return ie17d2b7dd4022f93c2c71a67b7a20a78bffcf7d8ccf29ebcee200a1107a650ad.NewAddCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTypesRequestBuilder) AddCopyFromContentTypeHub()(*i1fb6c90038e174f9bda7e7123feeeafbc289c60224a9133d157b5943665d3c85.AddCopyFromContentTypeHubRequestBuilder) {
    return i1fb6c90038e174f9bda7e7123feeeafbc289c60224a9133d157b5943665d3c85.NewAddCopyFromContentTypeHubRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBaseTypesRequestBuilderInternal instantiates a new BaseTypesRequestBuilder and sets the default values.
func NewBaseTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTypesRequestBuilder) {
    m := &BaseTypesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/lists/{list_id}/contentTypes/{contentType_id}/baseTypes{?top,skip,search,filter,count,orderby,select,expand}";
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
// GetCompatibleHubContentTypes builds and executes requests for operations under \sites\{site-id}\lists\{list-id}\contentTypes\{contentType-id}\baseTypes\microsoft.graph.getCompatibleHubContentTypes()
func (m *BaseTypesRequestBuilder) GetCompatibleHubContentTypes()(*i0d5e866a8d2f206f204617c8afe6a61bd6e740095f4d6adcb003c9e5bfb86591.GetCompatibleHubContentTypesRequestBuilder) {
    return i0d5e866a8d2f206f204617c8afe6a61bd6e740095f4d6adcb003c9e5bfb86591.NewGetCompatibleHubContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTypesRequestBuilder) Ref()(*i9b2fd694658511b1f9c73acd3c794477af0c97fce6079052b2f40bfc1363ad10.RefRequestBuilder) {
    return i9b2fd694658511b1f9c73acd3c794477af0c97fce6079052b2f40bfc1363ad10.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
