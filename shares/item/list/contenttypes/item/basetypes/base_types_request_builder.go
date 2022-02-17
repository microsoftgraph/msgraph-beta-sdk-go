package basetypes

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i9c57d4a5bbd1c322e28c23af64ded3abcadfa773abee6d0369f670242ba8fc9d "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/basetypes/ref"
    i9e014eff9740d7fa21092d10246680328c4c30a3afed1acc8a87674c3028059c "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/basetypes/addcopyfromcontenttypehub"
    iae8a7dc2c5d118710f055378fd4243ba0d11393ac1a3efd544936f2df44c4cee "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/basetypes/addcopy"
    ic8211cdd8c444cd128154d4724df286bd102c90d01822ec82e2c1b2d4444e255 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/basetypes/getcompatiblehubcontenttypes"
)

// BaseTypesRequestBuilder builds and executes requests for operations under \shares\{sharedDriveItem-id}\list\contentTypes\{contentType-id}\baseTypes
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
func (m *BaseTypesRequestBuilder) AddCopy()(*iae8a7dc2c5d118710f055378fd4243ba0d11393ac1a3efd544936f2df44c4cee.AddCopyRequestBuilder) {
    return iae8a7dc2c5d118710f055378fd4243ba0d11393ac1a3efd544936f2df44c4cee.NewAddCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTypesRequestBuilder) AddCopyFromContentTypeHub()(*i9e014eff9740d7fa21092d10246680328c4c30a3afed1acc8a87674c3028059c.AddCopyFromContentTypeHubRequestBuilder) {
    return i9e014eff9740d7fa21092d10246680328c4c30a3afed1acc8a87674c3028059c.NewAddCopyFromContentTypeHubRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBaseTypesRequestBuilderInternal instantiates a new BaseTypesRequestBuilder and sets the default values.
func NewBaseTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTypesRequestBuilder) {
    m := &BaseTypesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}/list/contentTypes/{contentType_id}/baseTypes{?top,skip,search,filter,count,orderby,select,expand}";
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
// GetCompatibleHubContentTypes builds and executes requests for operations under \shares\{sharedDriveItem-id}\list\contentTypes\{contentType-id}\baseTypes\microsoft.graph.getCompatibleHubContentTypes()
func (m *BaseTypesRequestBuilder) GetCompatibleHubContentTypes()(*ic8211cdd8c444cd128154d4724df286bd102c90d01822ec82e2c1b2d4444e255.GetCompatibleHubContentTypesRequestBuilder) {
    return ic8211cdd8c444cd128154d4724df286bd102c90d01822ec82e2c1b2d4444e255.NewGetCompatibleHubContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTypesRequestBuilder) Ref()(*i9c57d4a5bbd1c322e28c23af64ded3abcadfa773abee6d0369f670242ba8fc9d.RefRequestBuilder) {
    return i9c57d4a5bbd1c322e28c23af64ded3abcadfa773abee6d0369f670242ba8fc9d.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
