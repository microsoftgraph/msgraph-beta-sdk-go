package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4d2cb6355c05070b43e771667e3c932735e55500af7d9008d0667dc15158dbfd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/columnlinks"
    i584a99730fbd2ece5c6bfeeaf5256d826a03e37884ca6438f650d5236ca3485b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/ispublished"
    i6158e8fe936ce7a9232caf8ccfc3d8619a86e3d9493d7860a852c5b13e90ab25 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/basetypes"
    i854362c6353a95e813ebb2720a7c790bfb67cfb95b830527c59f7966139d71c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/associatewithhubsites"
    i97c73e8bccb2aeec40e8c91eaf704c572ef18478ca0f71947d7c3b651d907ea4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/copytodefaultcontentlocation"
    i9def72881302efcfa7f2e3cc87905ee20db183df30ca0f6e13d8da1a294d04e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/publish"
    ic2b61d41ac59272716a65a67f10b958477f6f42c3287e4324c941b6257334a50 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/columnpositions"
    idf8dd2eb707934e76443176905248e470c89a626df0ff62a74cc94ad0b884239 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/columns"
    ie8f8fe074cdc7e704e1945d78071313b03611e728052a551ae404ac364ab9d82 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/base"
    ieb6b07f19a05128118ad274cf509ece9882db22095de3b1dc2bec37b991bfc9a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/unpublish"
    i0180fc6328d895b6a5187ca5c60771e42333dc3b52a00fcfd8aba46845553404 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/columns/item"
    i63c74afb1978caaa67f675dda58750a522b0c0ed5def8dee9adfe882940e58ff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/columnlinks/item"
    i9de1ac96113c37c6db0105ac34c7abae8b4de298ef6731e0d312197e57b7706a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/basetypes/item"
    ie43a627d2d2c6d4d2ad9648e748314d9e7fddb4990ac1ada48f7463ff2422d00 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/columnpositions/item"
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
// ContentTypeItemRequestBuilderGetQueryParameters get contentTypes from users
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i854362c6353a95e813ebb2720a7c790bfb67cfb95b830527c59f7966139d71c3.AssociateWithHubSitesRequestBuilder) {
    return i854362c6353a95e813ebb2720a7c790bfb67cfb95b830527c59f7966139d71c3.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base the base property
func (m *ContentTypeItemRequestBuilder) Base()(*ie8f8fe074cdc7e704e1945d78071313b03611e728052a551ae404ac364ab9d82.BaseRequestBuilder) {
    return ie8f8fe074cdc7e704e1945d78071313b03611e728052a551ae404ac364ab9d82.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes the baseTypes property
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i6158e8fe936ce7a9232caf8ccfc3d8619a86e3d9493d7860a852c5b13e90ab25.BaseTypesRequestBuilder) {
    return i6158e8fe936ce7a9232caf8ccfc3d8619a86e3d9493d7860a852c5b13e90ab25.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i9de1ac96113c37c6db0105ac34c7abae8b4de298ef6731e0d312197e57b7706a.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return i9de1ac96113c37c6db0105ac34c7abae8b4de298ef6731e0d312197e57b7706a.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks the columnLinks property
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i4d2cb6355c05070b43e771667e3c932735e55500af7d9008d0667dc15158dbfd.ColumnLinksRequestBuilder) {
    return i4d2cb6355c05070b43e771667e3c932735e55500af7d9008d0667dc15158dbfd.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i63c74afb1978caaa67f675dda58750a522b0c0ed5def8dee9adfe882940e58ff.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return i63c74afb1978caaa67f675dda58750a522b0c0ed5def8dee9adfe882940e58ff.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions the columnPositions property
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*ic2b61d41ac59272716a65a67f10b958477f6f42c3287e4324c941b6257334a50.ColumnPositionsRequestBuilder) {
    return ic2b61d41ac59272716a65a67f10b958477f6f42c3287e4324c941b6257334a50.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*ie43a627d2d2c6d4d2ad9648e748314d9e7fddb4990ac1ada48f7463ff2422d00.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return ie43a627d2d2c6d4d2ad9648e748314d9e7fddb4990ac1ada48f7463ff2422d00.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns the columns property
func (m *ContentTypeItemRequestBuilder) Columns()(*idf8dd2eb707934e76443176905248e470c89a626df0ff62a74cc94ad0b884239.ColumnsRequestBuilder) {
    return idf8dd2eb707934e76443176905248e470c89a626df0ff62a74cc94ad0b884239.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i0180fc6328d895b6a5187ca5c60771e42333dc3b52a00fcfd8aba46845553404.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i0180fc6328d895b6a5187ca5c60771e42333dc3b52a00fcfd8aba46845553404.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}/list/contentTypes/{contentType%2Did}{?%24select,%24expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i97c73e8bccb2aeec40e8c91eaf704c572ef18478ca0f71947d7c3b651d907ea4.CopyToDefaultContentLocationRequestBuilder) {
    return i97c73e8bccb2aeec40e8c91eaf704c572ef18478ca0f71947d7c3b651d907ea4.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for users
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
// CreateGetRequestInformation get contentTypes from users
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
// CreatePatchRequestInformation update the navigation property contentTypes in users
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
// Delete delete navigation property contentTypes for users
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
// Get get contentTypes from users
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
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i584a99730fbd2ece5c6bfeeaf5256d826a03e37884ca6438f650d5236ca3485b.IsPublishedRequestBuilder) {
    return i584a99730fbd2ece5c6bfeeaf5256d826a03e37884ca6438f650d5236ca3485b.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in users
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
func (m *ContentTypeItemRequestBuilder) Publish()(*i9def72881302efcfa7f2e3cc87905ee20db183df30ca0f6e13d8da1a294d04e1.PublishRequestBuilder) {
    return i9def72881302efcfa7f2e3cc87905ee20db183df30ca0f6e13d8da1a294d04e1.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *ContentTypeItemRequestBuilder) Unpublish()(*ieb6b07f19a05128118ad274cf509ece9882db22095de3b1dc2bec37b991bfc9a.UnpublishRequestBuilder) {
    return ieb6b07f19a05128118ad274cf509ece9882db22095de3b1dc2bec37b991bfc9a.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
