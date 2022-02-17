package basetypes

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2902651c3bd6366958f3764b6177a24fe04ba5d32595c02c4aae8c0aac87e82c "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/basetypes/ref"
    i42462c7bfc418a1df15dc35cdf6deca2e311436a9206ca181ee56bbd94fa6ac4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/basetypes/addcopy"
    i94092ec39f60c080c646aa53176869d6ca585a42ad3331f7eb14711c9701d973 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/basetypes/getcompatiblehubcontenttypes"
    i9e28af25b3d66be13a9842b471ff9e170a7453faeab20631252c224e21e9c745 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/basetypes/addcopyfromcontenttypehub"
)

// BaseTypesRequestBuilder builds and executes requests for operations under \drives\{drive-id}\list\contentTypes\{contentType-id}\baseTypes
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
func (m *BaseTypesRequestBuilder) AddCopy()(*i42462c7bfc418a1df15dc35cdf6deca2e311436a9206ca181ee56bbd94fa6ac4.AddCopyRequestBuilder) {
    return i42462c7bfc418a1df15dc35cdf6deca2e311436a9206ca181ee56bbd94fa6ac4.NewAddCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTypesRequestBuilder) AddCopyFromContentTypeHub()(*i9e28af25b3d66be13a9842b471ff9e170a7453faeab20631252c224e21e9c745.AddCopyFromContentTypeHubRequestBuilder) {
    return i9e28af25b3d66be13a9842b471ff9e170a7453faeab20631252c224e21e9c745.NewAddCopyFromContentTypeHubRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBaseTypesRequestBuilderInternal instantiates a new BaseTypesRequestBuilder and sets the default values.
func NewBaseTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTypesRequestBuilder) {
    m := &BaseTypesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/list/contentTypes/{contentType_id}/baseTypes{?top,skip,search,filter,count,orderby,select,expand}";
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
// GetCompatibleHubContentTypes builds and executes requests for operations under \drives\{drive-id}\list\contentTypes\{contentType-id}\baseTypes\microsoft.graph.getCompatibleHubContentTypes()
func (m *BaseTypesRequestBuilder) GetCompatibleHubContentTypes()(*i94092ec39f60c080c646aa53176869d6ca585a42ad3331f7eb14711c9701d973.GetCompatibleHubContentTypesRequestBuilder) {
    return i94092ec39f60c080c646aa53176869d6ca585a42ad3331f7eb14711c9701d973.NewGetCompatibleHubContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTypesRequestBuilder) Ref()(*i2902651c3bd6366958f3764b6177a24fe04ba5d32595c02c4aae8c0aac87e82c.RefRequestBuilder) {
    return i2902651c3bd6366958f3764b6177a24fe04ba5d32595c02c4aae8c0aac87e82c.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
