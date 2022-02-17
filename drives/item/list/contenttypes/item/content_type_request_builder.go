package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i10fd39bdcc8bda0be99610b10b9660e1b0e04f54545563ba2f86aa5ee0e1eecf "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/ispublished"
    i1c1ef2c93981d039b9981f142914a8d8273a5c6a604d05ada12f520de3c3bc3a "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/base"
    i480eb8696c1f2f0778d89a9973d0e3553fb6e20bf68d8d7a209091c72ebfa6b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/associatewithhubsites"
    i6ec643542bff33f27de27396195458487402cfb21b2c257f89b7398dc5f7e148 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/copytodefaultcontentlocation"
    i73e62d430e2ef7842d7de0b849815c1c3fae9c06b6d4d4bd7d652ad10bb250d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/unpublish"
    i780d8b918fb7c93a6f80425157ff5fe5d6dac5fd3bf1389aceedc4883cda015a "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/columns"
    i939c16067629a808e874686b07dc16eff2aebaf0e2efb806672f28ca10a0b51e "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/publish"
    ib9753a53e6cc7e02241d16f2cb491afc520b1ba9ef8c30f532cda46a47d28464 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/columnpositions"
    ibb85c829f0bf47e2edb37d29ddc83348b385703321d85c4e9806fbad41f9f9ec "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/columnlinks"
    iea60a21f9c0699e1dad425f617fcb7e6742815b0de87074a71241f64c3d47618 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/basetypes"
    i74e76b37c2841dc75246821a35ea7fcedba7d931ad304b4164fa9ad54febf6f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/columns/item"
    ic60b95819e0e1f7cbc627d1cb7014b6bcf6a61251611f9c92082158251e62ed4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item/columnlinks/item"
)

// ContentTypeRequestBuilder builds and executes requests for operations under \drives\{drive-id}\list\contentTypes\{contentType-id}
type ContentTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContentTypeRequestBuilderDeleteOptions options for Delete
type ContentTypeRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeRequestBuilderGetOptions options for Get
type ContentTypeRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContentTypeRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContentTypeRequestBuilderPatchOptions options for Patch
type ContentTypeRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContentTypeRequestBuilder) AssociateWithHubSites()(*i480eb8696c1f2f0778d89a9973d0e3553fb6e20bf68d8d7a209091c72ebfa6b7.AssociateWithHubSitesRequestBuilder) {
    return i480eb8696c1f2f0778d89a9973d0e3553fb6e20bf68d8d7a209091c72ebfa6b7.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Base()(*i1c1ef2c93981d039b9981f142914a8d8273a5c6a604d05ada12f520de3c3bc3a.BaseRequestBuilder) {
    return i1c1ef2c93981d039b9981f142914a8d8273a5c6a604d05ada12f520de3c3bc3a.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) BaseTypes()(*iea60a21f9c0699e1dad425f617fcb7e6742815b0de87074a71241f64c3d47618.BaseTypesRequestBuilder) {
    return iea60a21f9c0699e1dad425f617fcb7e6742815b0de87074a71241f64c3d47618.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnLinks()(*ibb85c829f0bf47e2edb37d29ddc83348b385703321d85c4e9806fbad41f9f9ec.ColumnLinksRequestBuilder) {
    return ibb85c829f0bf47e2edb37d29ddc83348b385703321d85c4e9806fbad41f9f9ec.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeRequestBuilder) ColumnLinksById(id string)(*ic60b95819e0e1f7cbc627d1cb7014b6bcf6a61251611f9c92082158251e62ed4.ColumnLinkRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return ic60b95819e0e1f7cbc627d1cb7014b6bcf6a61251611f9c92082158251e62ed4.NewColumnLinkRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnPositions()(*ib9753a53e6cc7e02241d16f2cb491afc520b1ba9ef8c30f532cda46a47d28464.ColumnPositionsRequestBuilder) {
    return ib9753a53e6cc7e02241d16f2cb491afc520b1ba9ef8c30f532cda46a47d28464.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Columns()(*i780d8b918fb7c93a6f80425157ff5fe5d6dac5fd3bf1389aceedc4883cda015a.ColumnsRequestBuilder) {
    return i780d8b918fb7c93a6f80425157ff5fe5d6dac5fd3bf1389aceedc4883cda015a.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeRequestBuilder) ColumnsById(id string)(*i74e76b37c2841dc75246821a35ea7fcedba7d931ad304b4164fa9ad54febf6f1.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i74e76b37c2841dc75246821a35ea7fcedba7d931ad304b4164fa9ad54febf6f1.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeRequestBuilderInternal instantiates a new ContentTypeRequestBuilder and sets the default values.
func NewContentTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    m := &ContentTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/list/contentTypes/{contentType_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypeRequestBuilder instantiates a new ContentTypeRequestBuilder and sets the default values.
func NewContentTypeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContentTypeRequestBuilder) CopyToDefaultContentLocation()(*i6ec643542bff33f27de27396195458487402cfb21b2c257f89b7398dc5f7e148.CopyToDefaultContentLocationRequestBuilder) {
    return i6ec643542bff33f27de27396195458487402cfb21b2c257f89b7398dc5f7e148.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) CreateGetRequestInformation(options *ContentTypeRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) CreatePatchRequestInformation(options *ContentTypeRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Delete the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) Delete(options *ContentTypeRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) Get(options *ContentTypeRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType, error) {
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
// IsPublished builds and executes requests for operations under \drives\{drive-id}\list\contentTypes\{contentType-id}\microsoft.graph.isPublished()
func (m *ContentTypeRequestBuilder) IsPublished()(*i10fd39bdcc8bda0be99610b10b9660e1b0e04f54545563ba2f86aa5ee0e1eecf.IsPublishedRequestBuilder) {
    return i10fd39bdcc8bda0be99610b10b9660e1b0e04f54545563ba2f86aa5ee0e1eecf.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the collection of content types present in this list.
func (m *ContentTypeRequestBuilder) Patch(options *ContentTypeRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeRequestBuilder) Publish()(*i939c16067629a808e874686b07dc16eff2aebaf0e2efb806672f28ca10a0b51e.PublishRequestBuilder) {
    return i939c16067629a808e874686b07dc16eff2aebaf0e2efb806672f28ca10a0b51e.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Unpublish()(*i73e62d430e2ef7842d7de0b849815c1c3fae9c06b6d4d4bd7d652ad10bb250d4.UnpublishRequestBuilder) {
    return i73e62d430e2ef7842d7de0b849815c1c3fae9c06b6d4d4bd7d652ad10bb250d4.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
