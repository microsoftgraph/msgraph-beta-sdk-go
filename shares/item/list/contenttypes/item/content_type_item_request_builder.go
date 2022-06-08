package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0c5a6eb9fea66115bd69699c7bf3c14c8c56d6ddaf4d97c8afd618199d935eae "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/base"
    i2730267b865ade42ffde9c118a145ac5cbe3724b0429d2790c809e94a72e89fa "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/copytodefaultcontentlocation"
    i2da8c35e7e97f2ef273098f53bc6e776a22cb7d02b84bfe9900b6a3040d2cc9f "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/associatewithhubsites"
    i3ba10a504af14f129ece317956d8965ac8085afa2820b041d949da7eb57874d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/ispublished"
    i48be76cbf9aaa82f0511c875016a4ee52f28bd44fe4c1ddd7305aecf7b2a5bf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/columnlinks"
    i50b805823a2e1e06133a87ccbd948dbdae542783183fd4393f09b7971d5822ab "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/basetypes"
    i5b9a34041146ae573ad4ad3e977514d4015468ef377df8319306938b4570f0c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/columnpositions"
    i760311997236edb768f49330f26e7daf027dafc1e3608a33efd7d83824b07a14 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/publish"
    i91d1e71a788d2b90b0e38b793d7946b2b495724509c5446dfad0545167467897 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/unpublish"
    iad7db789331194a827d671c9f4c9c9db1a2a5947624f5e39904deb155bb43cce "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/columns"
    i36f46f4ddd3e4e31658a8596691f7c13fb36ceef9f5b000c1b72092e3c8bccb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/basetypes/item"
    ic32342812075059c3425a8c2f3f7c9d6cbe906851bdae63f19bf4b836ca3479b "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/columnpositions/item"
    ice1877e581ee39a1cf8b872c81041f0bab8f68c1e2db0210dcaa15a5cb798804 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/columns/item"
    icfa060aea76f8132231b7d0f12ab0a407b2c00f0d82270c0adce976e1034f6ff "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item/columnlinks/item"
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
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i2da8c35e7e97f2ef273098f53bc6e776a22cb7d02b84bfe9900b6a3040d2cc9f.AssociateWithHubSitesRequestBuilder) {
    return i2da8c35e7e97f2ef273098f53bc6e776a22cb7d02b84bfe9900b6a3040d2cc9f.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base the base property
func (m *ContentTypeItemRequestBuilder) Base()(*i0c5a6eb9fea66115bd69699c7bf3c14c8c56d6ddaf4d97c8afd618199d935eae.BaseRequestBuilder) {
    return i0c5a6eb9fea66115bd69699c7bf3c14c8c56d6ddaf4d97c8afd618199d935eae.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes the baseTypes property
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i50b805823a2e1e06133a87ccbd948dbdae542783183fd4393f09b7971d5822ab.BaseTypesRequestBuilder) {
    return i50b805823a2e1e06133a87ccbd948dbdae542783183fd4393f09b7971d5822ab.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i36f46f4ddd3e4e31658a8596691f7c13fb36ceef9f5b000c1b72092e3c8bccb0.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return i36f46f4ddd3e4e31658a8596691f7c13fb36ceef9f5b000c1b72092e3c8bccb0.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks the columnLinks property
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i48be76cbf9aaa82f0511c875016a4ee52f28bd44fe4c1ddd7305aecf7b2a5bf9.ColumnLinksRequestBuilder) {
    return i48be76cbf9aaa82f0511c875016a4ee52f28bd44fe4c1ddd7305aecf7b2a5bf9.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*icfa060aea76f8132231b7d0f12ab0a407b2c00f0d82270c0adce976e1034f6ff.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return icfa060aea76f8132231b7d0f12ab0a407b2c00f0d82270c0adce976e1034f6ff.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions the columnPositions property
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i5b9a34041146ae573ad4ad3e977514d4015468ef377df8319306938b4570f0c6.ColumnPositionsRequestBuilder) {
    return i5b9a34041146ae573ad4ad3e977514d4015468ef377df8319306938b4570f0c6.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*ic32342812075059c3425a8c2f3f7c9d6cbe906851bdae63f19bf4b836ca3479b.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return ic32342812075059c3425a8c2f3f7c9d6cbe906851bdae63f19bf4b836ca3479b.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns the columns property
func (m *ContentTypeItemRequestBuilder) Columns()(*iad7db789331194a827d671c9f4c9c9db1a2a5947624f5e39904deb155bb43cce.ColumnsRequestBuilder) {
    return iad7db789331194a827d671c9f4c9c9db1a2a5947624f5e39904deb155bb43cce.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*ice1877e581ee39a1cf8b872c81041f0bab8f68c1e2db0210dcaa15a5cb798804.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return ice1877e581ee39a1cf8b872c81041f0bab8f68c1e2db0210dcaa15a5cb798804.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem%2Did}/list/contentTypes/{contentType%2Did}{?%24select,%24expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i2730267b865ade42ffde9c118a145ac5cbe3724b0429d2790c809e94a72e89fa.CopyToDefaultContentLocationRequestBuilder) {
    return i2730267b865ade42ffde9c118a145ac5cbe3724b0429d2790c809e94a72e89fa.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for shares
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property contentTypes for shares
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ContentTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ContentTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property contentTypes in shares
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property contentTypes in shares
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *ContentTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property contentTypes for shares
func (m *ContentTypeItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property contentTypes for shares
func (m *ContentTypeItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ContentTypeItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ContentTypeItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentTypeFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i3ba10a504af14f129ece317956d8965ac8085afa2820b041d949da7eb57874d0.IsPublishedRequestBuilder) {
    return i3ba10a504af14f129ece317956d8965ac8085afa2820b041d949da7eb57874d0.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in shares
func (m *ContentTypeItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property contentTypes in shares
func (m *ContentTypeItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *ContentTypeItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Publish the publish property
func (m *ContentTypeItemRequestBuilder) Publish()(*i760311997236edb768f49330f26e7daf027dafc1e3608a33efd7d83824b07a14.PublishRequestBuilder) {
    return i760311997236edb768f49330f26e7daf027dafc1e3608a33efd7d83824b07a14.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i91d1e71a788d2b90b0e38b793d7946b2b495724509c5446dfad0545167467897.UnpublishRequestBuilder) {
    return i91d1e71a788d2b90b0e38b793d7946b2b495724509c5446dfad0545167467897.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
