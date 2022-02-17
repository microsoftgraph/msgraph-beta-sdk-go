package base

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i179ccad7de63b64d143c44f451b3b444a3c2eb26e54302cac18aa30f72e6bd41 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/base/ispublished"
    i2cbd91266eb81c1429813f4f60f6906f932527164b8f0c92c40f0cad1f6a2f12 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/base/unpublish"
    i2cd33d8763cd9001f7770e7bb89d35dfea2e62b7e13410757dee22e2a6967ad4 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/base/ref"
    i6120d63ec050fcd2c66e9cb891d59795ad2a334fe1cfae904272cc459b519673 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/base/associatewithhubsites"
    ib907afd8ed21707b417d5bfe13836e9bf4a80ad38758c91f738c063ea6e586f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/base/publish"
    ic962a2130bbff27f74713633234a56af1fd111f8038fc7dc4683eb35ca8dda1b "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/base/copytodefaultcontentlocation"
)

// BaseRequestBuilder builds and executes requests for operations under \shares\{sharedDriveItem-id}\list\contentTypes\{contentType-id}\base
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
func (m *BaseRequestBuilder) AssociateWithHubSites()(*i6120d63ec050fcd2c66e9cb891d59795ad2a334fe1cfae904272cc459b519673.AssociateWithHubSitesRequestBuilder) {
    return i6120d63ec050fcd2c66e9cb891d59795ad2a334fe1cfae904272cc459b519673.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewBaseRequestBuilderInternal instantiates a new BaseRequestBuilder and sets the default values.
func NewBaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    m := &BaseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}/list/contentTypes/{contentType_id}/base{?select,expand}";
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
func (m *BaseRequestBuilder) CopyToDefaultContentLocation()(*ic962a2130bbff27f74713633234a56af1fd111f8038fc7dc4683eb35ca8dda1b.CopyToDefaultContentLocationRequestBuilder) {
    return ic962a2130bbff27f74713633234a56af1fd111f8038fc7dc4683eb35ca8dda1b.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// IsPublished builds and executes requests for operations under \shares\{sharedDriveItem-id}\list\contentTypes\{contentType-id}\base\microsoft.graph.isPublished()
func (m *BaseRequestBuilder) IsPublished()(*i179ccad7de63b64d143c44f451b3b444a3c2eb26e54302cac18aa30f72e6bd41.IsPublishedRequestBuilder) {
    return i179ccad7de63b64d143c44f451b3b444a3c2eb26e54302cac18aa30f72e6bd41.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Publish()(*ib907afd8ed21707b417d5bfe13836e9bf4a80ad38758c91f738c063ea6e586f7.PublishRequestBuilder) {
    return ib907afd8ed21707b417d5bfe13836e9bf4a80ad38758c91f738c063ea6e586f7.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Ref()(*i2cd33d8763cd9001f7770e7bb89d35dfea2e62b7e13410757dee22e2a6967ad4.RefRequestBuilder) {
    return i2cd33d8763cd9001f7770e7bb89d35dfea2e62b7e13410757dee22e2a6967ad4.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Unpublish()(*i2cbd91266eb81c1429813f4f60f6906f932527164b8f0c92c40f0cad1f6a2f12.UnpublishRequestBuilder) {
    return i2cbd91266eb81c1429813f4f60f6906f932527164b8f0c92c40f0cad1f6a2f12.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
