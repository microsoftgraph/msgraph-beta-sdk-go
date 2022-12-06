package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i6cd0dd643a8853c5e36c664f3266e8bc15e43dfd7da99ab8ec908ba5a2e2846c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/unpublish"
    i71a00748ec1b3f1916f08f1378ac75be568e315e8c45671703ce0a998a523535 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/basetypes"
    i8b97072a604d4b89bbdcd18ca9b769ea9b5f892512bf132888c5ceb0d874996d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/columnpositions"
    i8f46b19091d8fc007c8204da63e2fecf63e3e12658f88e306d5afdce4ccf783c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/associatewithhubsites"
    ia5eb9dbcf222805d1a39da295590fdc1cdff5f38e71fbe02b0cbde0ac79badf8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/ispublished"
    ic33bb574f58874bf5559129e9d7f446f758bfc0baa4cc59067b93d1138da4756 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/columnlinks"
    ic77a6b884e9bdc94d36c4085a09bbcf0676b01c87b3bc44e07282c5c2f0de7d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/publish"
    id33ae53bb415f03b269f46aa3901afb4aa844cd415a3baca9ba770f3f695c744 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/copytodefaultcontentlocation"
    ida787a0eecc9cbdceedc5fa328132abca55da74c1475c3cd90d33f515a955a58 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/base"
    idb904522da356683cdbc2ac70d390d7692262d927ea7a3752d44744d20e4c435 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/columns"
    i5c3b82c0971f80566e2f27fb191b318b1e3605714d556179aab3bf7a231cc5e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/columnlinks/item"
    i881e5cfd2288828ecac5b4440b76a4eba0683212edb64e32af0744d8d4e2c0e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/columns/item"
    iab2728a4c36bbe5d4b7f8a74e0d7c74f5bb30814b9618b8c70571b67c6bd0c75 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/basetypes/item"
    ib8b83821cbc0899cc9656c33cc49eb4adc13be44554b10215afcaaf8eae6327a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item/columnpositions/item"
)

// ContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ContentTypeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ContentTypeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContentTypeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ContentTypeItemRequestBuilderGetQueryParameters get contentTypes from groups
type ContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ContentTypeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContentTypeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ContentTypeItemRequestBuilderGetQueryParameters
}
// ContentTypeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContentTypeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssociateWithHubSites provides operations to call the associateWithHubSites method.
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i8f46b19091d8fc007c8204da63e2fecf63e3e12658f88e306d5afdce4ccf783c.AssociateWithHubSitesRequestBuilder) {
    return i8f46b19091d8fc007c8204da63e2fecf63e3e12658f88e306d5afdce4ccf783c.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base provides operations to manage the base property of the microsoft.graph.contentType entity.
func (m *ContentTypeItemRequestBuilder) Base()(*ida787a0eecc9cbdceedc5fa328132abca55da74c1475c3cd90d33f515a955a58.BaseRequestBuilder) {
    return ida787a0eecc9cbdceedc5fa328132abca55da74c1475c3cd90d33f515a955a58.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i71a00748ec1b3f1916f08f1378ac75be568e315e8c45671703ce0a998a523535.BaseTypesRequestBuilder) {
    return i71a00748ec1b3f1916f08f1378ac75be568e315e8c45671703ce0a998a523535.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*iab2728a4c36bbe5d4b7f8a74e0d7c74f5bb30814b9618b8c70571b67c6bd0c75.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return iab2728a4c36bbe5d4b7f8a74e0d7c74f5bb30814b9618b8c70571b67c6bd0c75.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*ic33bb574f58874bf5559129e9d7f446f758bfc0baa4cc59067b93d1138da4756.ColumnLinksRequestBuilder) {
    return ic33bb574f58874bf5559129e9d7f446f758bfc0baa4cc59067b93d1138da4756.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i5c3b82c0971f80566e2f27fb191b318b1e3605714d556179aab3bf7a231cc5e3.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return i5c3b82c0971f80566e2f27fb191b318b1e3605714d556179aab3bf7a231cc5e3.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i8b97072a604d4b89bbdcd18ca9b769ea9b5f892512bf132888c5ceb0d874996d.ColumnPositionsRequestBuilder) {
    return i8b97072a604d4b89bbdcd18ca9b769ea9b5f892512bf132888c5ceb0d874996d.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*ib8b83821cbc0899cc9656c33cc49eb4adc13be44554b10215afcaaf8eae6327a.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return ib8b83821cbc0899cc9656c33cc49eb4adc13be44554b10215afcaaf8eae6327a.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns provides operations to manage the columns property of the microsoft.graph.contentType entity.
func (m *ContentTypeItemRequestBuilder) Columns()(*idb904522da356683cdbc2ac70d390d7692262d927ea7a3752d44744d20e4c435.ColumnsRequestBuilder) {
    return idb904522da356683cdbc2ac70d390d7692262d927ea7a3752d44744d20e4c435.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.contentType entity.
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i881e5cfd2288828ecac5b4440b76a4eba0683212edb64e32af0744d8d4e2c0e1.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i881e5cfd2288828ecac5b4440b76a4eba0683212edb64e32af0744d8d4e2c0e1.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/drives/{drive%2Did}/list/contentTypes/{contentType%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypeItemRequestBuilder instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToDefaultContentLocation provides operations to call the copyToDefaultContentLocation method.
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*id33ae53bb415f03b269f46aa3901afb4aa844cd415a3baca9ba770f3f695c744.CopyToDefaultContentLocationRequestBuilder) {
    return id33ae53bb415f03b269f46aa3901afb4aa844cd415a3baca9ba770f3f695c744.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for groups
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ContentTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get contentTypes from groups
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ContentTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property contentTypes in groups
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *ContentTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property contentTypes for groups
func (m *ContentTypeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ContentTypeItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get contentTypes from groups
func (m *ContentTypeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ContentTypeItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*ia5eb9dbcf222805d1a39da295590fdc1cdff5f38e71fbe02b0cbde0ac79badf8.IsPublishedRequestBuilder) {
    return ia5eb9dbcf222805d1a39da295590fdc1cdff5f38e71fbe02b0cbde0ac79badf8.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in groups
func (m *ContentTypeItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *ContentTypeItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable), nil
}
// Publish provides operations to call the publish method.
func (m *ContentTypeItemRequestBuilder) Publish()(*ic77a6b884e9bdc94d36c4085a09bbcf0676b01c87b3bc44e07282c5c2f0de7d1.PublishRequestBuilder) {
    return ic77a6b884e9bdc94d36c4085a09bbcf0676b01c87b3bc44e07282c5c2f0de7d1.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish provides operations to call the unpublish method.
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i6cd0dd643a8853c5e36c664f3266e8bc15e43dfd7da99ab8ec908ba5a2e2846c.UnpublishRequestBuilder) {
    return i6cd0dd643a8853c5e36c664f3266e8bc15e43dfd7da99ab8ec908ba5a2e2846c.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
