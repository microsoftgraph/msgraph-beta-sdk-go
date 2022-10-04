package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i01d884c5b597bb4bbb21ef6ddb793472b4e5cfac11eee17fa06a7c93e7af3961 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/unpublish"
    i0dbe0e79f8d2132817cfd4fd9c0ba48a22c1436a654ba3b515c9d78ea04153bf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/columnpositions"
    i15b535ee8c2ffa35c2be78c2e83049b8f87132a9d6c9095165aece4073c7e68b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/basetypes"
    i5072c1863bff5f51d828c097f140a9067fffcdd636bce59d31fd3cd41b86a7f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/base"
    i53102ef79a47ebbae7839e80aaa6802806794d37343b777c078fcbaad0a195c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/columns"
    iaf87d69143fb1b049d1e8643342885a68173ad08feea74278b1460b57f935ec4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/columnlinks"
    iba673aaa67a5f7556ea399c86c8c32b04c721c33b18e487c30de92eea73d517b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/publish"
    ibdbf84d3479e495735427bb0dcdf1fe8f04d07858d5d6a3a0f34ef4cd0ea5c17 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/copytodefaultcontentlocation"
    ic426e586ad0e202686674873cbdea7a8018e5b74cf6d66de10ed9573184cdd97 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/associatewithhubsites"
    ifa15559168b540626f112cbef09640515aab7cf666d590c01974f1056022fd94 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/ispublished"
    i0c5d02e0200c8a34f0952dda3d6bed322ad1bf731eaa3dc3c37e94f37cc44a12 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/columnpositions/item"
    i70ae4aef64d211900e5691042f2ac667a4a38c969183ce5df5980865df763382 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/columnlinks/item"
    i71496ba62ae8bd6cc9c2a77d0a541c08ebf273eea9ecf422f1040c706f9de5ad "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/columns/item"
    if47acb49fe28f5ed2e47bb6c2bd5cb847ac5f03b74ad8571b28a23def60220c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/contenttypes/item/basetypes/item"
)

// ContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.site entity.
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
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types defined for this site.
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
// AssociateWithHubSites the associateWithHubSites property
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*ic426e586ad0e202686674873cbdea7a8018e5b74cf6d66de10ed9573184cdd97.AssociateWithHubSitesRequestBuilder) {
    return ic426e586ad0e202686674873cbdea7a8018e5b74cf6d66de10ed9573184cdd97.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base the base property
func (m *ContentTypeItemRequestBuilder) Base()(*i5072c1863bff5f51d828c097f140a9067fffcdd636bce59d31fd3cd41b86a7f7.BaseRequestBuilder) {
    return i5072c1863bff5f51d828c097f140a9067fffcdd636bce59d31fd3cd41b86a7f7.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes the baseTypes property
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i15b535ee8c2ffa35c2be78c2e83049b8f87132a9d6c9095165aece4073c7e68b.BaseTypesRequestBuilder) {
    return i15b535ee8c2ffa35c2be78c2e83049b8f87132a9d6c9095165aece4073c7e68b.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*if47acb49fe28f5ed2e47bb6c2bd5cb847ac5f03b74ad8571b28a23def60220c4.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return if47acb49fe28f5ed2e47bb6c2bd5cb847ac5f03b74ad8571b28a23def60220c4.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks the columnLinks property
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*iaf87d69143fb1b049d1e8643342885a68173ad08feea74278b1460b57f935ec4.ColumnLinksRequestBuilder) {
    return iaf87d69143fb1b049d1e8643342885a68173ad08feea74278b1460b57f935ec4.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i70ae4aef64d211900e5691042f2ac667a4a38c969183ce5df5980865df763382.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return i70ae4aef64d211900e5691042f2ac667a4a38c969183ce5df5980865df763382.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions the columnPositions property
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i0dbe0e79f8d2132817cfd4fd9c0ba48a22c1436a654ba3b515c9d78ea04153bf.ColumnPositionsRequestBuilder) {
    return i0dbe0e79f8d2132817cfd4fd9c0ba48a22c1436a654ba3b515c9d78ea04153bf.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i0c5d02e0200c8a34f0952dda3d6bed322ad1bf731eaa3dc3c37e94f37cc44a12.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i0c5d02e0200c8a34f0952dda3d6bed322ad1bf731eaa3dc3c37e94f37cc44a12.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns the columns property
func (m *ContentTypeItemRequestBuilder) Columns()(*i53102ef79a47ebbae7839e80aaa6802806794d37343b777c078fcbaad0a195c1.ColumnsRequestBuilder) {
    return i53102ef79a47ebbae7839e80aaa6802806794d37343b777c078fcbaad0a195c1.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i71496ba62ae8bd6cc9c2a77d0a541c08ebf273eea9ecf422f1040c706f9de5ad.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i71496ba62ae8bd6cc9c2a77d0a541c08ebf273eea9ecf422f1040c706f9de5ad.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/contentTypes/{contentType%2Did}{?%24select,%24expand}";
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
// CopyToDefaultContentLocation the copyToDefaultContentLocation property
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*ibdbf84d3479e495735427bb0dcdf1fe8f04d07858d5d6a3a0f34ef4cd0ea5c17.CopyToDefaultContentLocationRequestBuilder) {
    return ibdbf84d3479e495735427bb0dcdf1fe8f04d07858d5d6a3a0f34ef4cd0ea5c17.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// CreateGetRequestInformation the collection of content types defined for this site.
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
// Get the collection of content types defined for this site.
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
func (m *ContentTypeItemRequestBuilder) IsPublished()(*ifa15559168b540626f112cbef09640515aab7cf666d590c01974f1056022fd94.IsPublishedRequestBuilder) {
    return ifa15559168b540626f112cbef09640515aab7cf666d590c01974f1056022fd94.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Publish the publish property
func (m *ContentTypeItemRequestBuilder) Publish()(*iba673aaa67a5f7556ea399c86c8c32b04c721c33b18e487c30de92eea73d517b.PublishRequestBuilder) {
    return iba673aaa67a5f7556ea399c86c8c32b04c721c33b18e487c30de92eea73d517b.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i01d884c5b597bb4bbb21ef6ddb793472b4e5cfac11eee17fa06a7c93e7af3961.UnpublishRequestBuilder) {
    return i01d884c5b597bb4bbb21ef6ddb793472b4e5cfac11eee17fa06a7c93e7af3961.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
