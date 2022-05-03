package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0eae2228a06806e82320cf19a03d576b42621cc4a2dd8806ce532a57bfebb362 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/associatewithhubsites"
    i105f5ec8936f20cc343ef53f19692d9484bb0e2317204f4e0a83cb86af419abf "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/columnlinks"
    i20d9f259a492016a967dd6a10e660b98616fd7d4888ef6ee8111990cf5517382 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/unpublish"
    i6d9ef8ddda3a3e7ebcb0a0d568b23e622439e98ff5309351bc26797ae8d500cb "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/columnpositions"
    i7e506eefecd6b2935e76b82d621a57aa9ed417d501afb7d19965f3434a172010 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/basetypes"
    i7e8f4985115eafb5e1b9a5dcfae2a3c836415e064af553f7801cdbb55757ca11 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/base"
    i910a70f3e4a25581ff8d025ecc5ab7dd183edf0f6e7af5d6e64e49e265867068 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/copytodefaultcontentlocation"
    i943b393012430b44d1cf499bc2d8e50d28fa67a47c6be6689e906105bd5eb4e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/ispublished"
    if8b488d84d1b1415d1c88a73424e76e4a1a1efd4526e5164d0604809ea5c8d13 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/publish"
    ifd4412476ff102f7c92f4e4a139b13990d10cc2647ddfa6e976c2dd5a98653ce "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/columns"
    i88d2724fa2dada76f95cc3b7d1c2d9bd0f61469930ea96ff9d0dc112126c5a8e "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/columnlinks/item"
    i9e8f3bdc2a2da10dbbb9f032cac127a9f38b1132156d876cec226651df880b26 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/columns/item"
    ic5b939ddec677078eada74ae96425a6dd8ae52980fe6a4ce4ef57a15ff2c7be8 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/basetypes/item"
    ie3518f6b63ce04526457403d54384566536a4de177b91d352bd57d81171f4cd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/columnpositions/item"
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i0eae2228a06806e82320cf19a03d576b42621cc4a2dd8806ce532a57bfebb362.AssociateWithHubSitesRequestBuilder) {
    return i0eae2228a06806e82320cf19a03d576b42621cc4a2dd8806ce532a57bfebb362.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base the base property
func (m *ContentTypeItemRequestBuilder) Base()(*i7e8f4985115eafb5e1b9a5dcfae2a3c836415e064af553f7801cdbb55757ca11.BaseRequestBuilder) {
    return i7e8f4985115eafb5e1b9a5dcfae2a3c836415e064af553f7801cdbb55757ca11.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes the baseTypes property
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i7e506eefecd6b2935e76b82d621a57aa9ed417d501afb7d19965f3434a172010.BaseTypesRequestBuilder) {
    return i7e506eefecd6b2935e76b82d621a57aa9ed417d501afb7d19965f3434a172010.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*ic5b939ddec677078eada74ae96425a6dd8ae52980fe6a4ce4ef57a15ff2c7be8.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return ic5b939ddec677078eada74ae96425a6dd8ae52980fe6a4ce4ef57a15ff2c7be8.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks the columnLinks property
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i105f5ec8936f20cc343ef53f19692d9484bb0e2317204f4e0a83cb86af419abf.ColumnLinksRequestBuilder) {
    return i105f5ec8936f20cc343ef53f19692d9484bb0e2317204f4e0a83cb86af419abf.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i88d2724fa2dada76f95cc3b7d1c2d9bd0f61469930ea96ff9d0dc112126c5a8e.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return i88d2724fa2dada76f95cc3b7d1c2d9bd0f61469930ea96ff9d0dc112126c5a8e.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions the columnPositions property
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i6d9ef8ddda3a3e7ebcb0a0d568b23e622439e98ff5309351bc26797ae8d500cb.ColumnPositionsRequestBuilder) {
    return i6d9ef8ddda3a3e7ebcb0a0d568b23e622439e98ff5309351bc26797ae8d500cb.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*ie3518f6b63ce04526457403d54384566536a4de177b91d352bd57d81171f4cd7.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return ie3518f6b63ce04526457403d54384566536a4de177b91d352bd57d81171f4cd7.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns the columns property
func (m *ContentTypeItemRequestBuilder) Columns()(*ifd4412476ff102f7c92f4e4a139b13990d10cc2647ddfa6e976c2dd5a98653ce.ColumnsRequestBuilder) {
    return ifd4412476ff102f7c92f4e4a139b13990d10cc2647ddfa6e976c2dd5a98653ce.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i9e8f3bdc2a2da10dbbb9f032cac127a9f38b1132156d876cec226651df880b26.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i9e8f3bdc2a2da10dbbb9f032cac127a9f38b1132156d876cec226651df880b26.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/contentTypes/{contentType%2Did}{?%24select,%24expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i910a70f3e4a25581ff8d025ecc5ab7dd183edf0f6e7af5d6e64e49e265867068.CopyToDefaultContentLocationRequestBuilder) {
    return i910a70f3e4a25581ff8d025ecc5ab7dd183edf0f6e7af5d6e64e49e265867068.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property contentTypes for sites
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property contentTypes for sites
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
// CreateGetRequestInformationWithRequestConfiguration the collection of content types defined for this site.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the collection of content types defined for this site.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ContentTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property contentTypes in sites
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property contentTypes in sites
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
// DeleteWithResponseHandler delete navigation property contentTypes for sites
func (m *ContentTypeItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *ContentTypeItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property contentTypes for sites
func (m *ContentTypeItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *ContentTypeItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetWithResponseHandler the collection of content types defined for this site.
func (m *ContentTypeItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ContentTypeItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the collection of content types defined for this site.
func (m *ContentTypeItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ContentTypeItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
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
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i943b393012430b44d1cf499bc2d8e50d28fa67a47c6be6689e906105bd5eb4e3.IsPublishedRequestBuilder) {
    return i943b393012430b44d1cf499bc2d8e50d28fa67a47c6be6689e906105bd5eb4e3.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property contentTypes in sites
func (m *ContentTypeItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *ContentTypeItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property contentTypes in sites
func (m *ContentTypeItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *ContentTypeItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *ContentTypeItemRequestBuilder) Publish()(*if8b488d84d1b1415d1c88a73424e76e4a1a1efd4526e5164d0604809ea5c8d13.PublishRequestBuilder) {
    return if8b488d84d1b1415d1c88a73424e76e4a1a1efd4526e5164d0604809ea5c8d13.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i20d9f259a492016a967dd6a10e660b98616fd7d4888ef6ee8111990cf5517382.UnpublishRequestBuilder) {
    return i20d9f259a492016a967dd6a10e660b98616fd7d4888ef6ee8111990cf5517382.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
