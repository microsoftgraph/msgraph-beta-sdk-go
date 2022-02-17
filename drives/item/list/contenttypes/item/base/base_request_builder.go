package base

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4f502c834b91b81d422874ae1305b8e893af6a209b186bb0b3c34e6f9a881be8 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/base/copytodefaultcontentlocation"
    i84f249b4e581b95fe33c32cd89501d8cc15bd2c481feb5c5a4972994b1fb4fe7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/base/ispublished"
    ia9eee964ed819f4c6e5eb1324a5057bd19d2563b785350a514646643ccc875d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/base/publish"
    iac09efe71a32af377b410b0cc245313ce6f4fc4bf8976871727573634fdc771a "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/base/ref"
    ie2f7f3e758dde263e51067e614835b062b4527d6d3e8693946ad374be8aefdea "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/base/unpublish"
    ifaa531f1b6ae0c7bfc2907fdbd8db58a9ecf5ad87a62b85e6af848d60d955ec2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/base/associatewithhubsites"
)

// BaseRequestBuilder builds and executes requests for operations under \drives\{drive-id}\list\contentTypes\{contentType-id}\base
type BaseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BaseRequestBuilderGetOptions options for Get
type BaseRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BaseRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BaseRequestBuilderGetQueryParameters parent contentType from which this content type is derived.
type BaseRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
func (m *BaseRequestBuilder) AssociateWithHubSites()(*ifaa531f1b6ae0c7bfc2907fdbd8db58a9ecf5ad87a62b85e6af848d60d955ec2.AssociateWithHubSitesRequestBuilder) {
    return ifaa531f1b6ae0c7bfc2907fdbd8db58a9ecf5ad87a62b85e6af848d60d955ec2.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBaseRequestBuilderInternal instantiates a new BaseRequestBuilder and sets the default values.
func NewBaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    m := &BaseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/list/contentTypes/{contentType_id}/base{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBaseRequestBuilder instantiates a new BaseRequestBuilder and sets the default values.
func NewBaseRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBaseRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *BaseRequestBuilder) CopyToDefaultContentLocation()(*i4f502c834b91b81d422874ae1305b8e893af6a209b186bb0b3c34e6f9a881be8.CopyToDefaultContentLocationRequestBuilder) {
    return i4f502c834b91b81d422874ae1305b8e893af6a209b186bb0b3c34e6f9a881be8.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation parent contentType from which this content type is derived.
func (m *BaseRequestBuilder) CreateGetRequestInformation(options *BaseRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get parent contentType from which this content type is derived.
func (m *BaseRequestBuilder) Get(options *BaseRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewContentType() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType), nil
}
// IsPublished builds and executes requests for operations under \drives\{drive-id}\list\contentTypes\{contentType-id}\base\microsoft.graph.isPublished()
func (m *BaseRequestBuilder) IsPublished()(*i84f249b4e581b95fe33c32cd89501d8cc15bd2c481feb5c5a4972994b1fb4fe7.IsPublishedRequestBuilder) {
    return i84f249b4e581b95fe33c32cd89501d8cc15bd2c481feb5c5a4972994b1fb4fe7.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Publish()(*ia9eee964ed819f4c6e5eb1324a5057bd19d2563b785350a514646643ccc875d8.PublishRequestBuilder) {
    return ia9eee964ed819f4c6e5eb1324a5057bd19d2563b785350a514646643ccc875d8.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Ref()(*iac09efe71a32af377b410b0cc245313ce6f4fc4bf8976871727573634fdc771a.RefRequestBuilder) {
    return iac09efe71a32af377b410b0cc245313ce6f4fc4bf8976871727573634fdc771a.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Unpublish()(*ie2f7f3e758dde263e51067e614835b062b4527d6d3e8693946ad374be8aefdea.UnpublishRequestBuilder) {
    return ie2f7f3e758dde263e51067e614835b062b4527d6d3e8693946ad374be8aefdea.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
