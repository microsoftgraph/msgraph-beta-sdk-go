package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i072c7323df45db1f1ee748560abd200a43db5ebff2606b32cdad943a8247a3e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/publish"
    i2e24155317ddebf9f80799a7e7567793da2868054b1116675108e3b4674a0fc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/copytodefaultcontentlocation"
    i48e086ed64fcd59d03281d5ffa804a5c862c38e82c9234f5eceee38d4e99a30a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/associatewithhubsites"
    i5591abfce713804f2fb74b208efc6db0ee225ff736b8e9ac134955aad253bf7d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columnlinks"
    i69a0f0826d29adb7fdeff42634c0cfeaa9dfca0c790eb1d7972dc8275fdcd7bb "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/unpublish"
    i7bed78d5c7b182bb6b35c0cbd54072e0caa68f8aaee7b2c72ed1530d26e66641 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/base"
    i929fc0841e147d731d330fffac2b603f59d2e7ee09ba0bb76a82c060f484478d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/basetypes"
    ib19674cd06300cabca5bf2c5ea72a1bbc22118fa8aabb80e98be0a419c35e6cc "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/ispublished"
    icc9cf5371dacb763ed32f4832e68daf65622d2d8a427a8b01fd79665f253c4cc "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columns"
    ieabef4a072848f73eb7851ae8297bf4213dbcc6b84e915fef10453e9c720ea90 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columnpositions"
    i47b1c489561ff1669cae57a9dd84a767bbed7f82b77c0cdc7389dc73950c6be7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/basetypes/item"
    i49c22f7846eb325805a3ce05072fb08dfdd09048ded457ecb9796b7a866ec2dd "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columnpositions/item"
    ib1338a26b88df4781df2ab470814626fd7c5b61709da7294d63f28bea3263a37 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columns/item"
    if1bf402e4bfc313e6d11fabbebde44dffbe81a5aa7a12e657549b48143aefbb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columnlinks/item"
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i48e086ed64fcd59d03281d5ffa804a5c862c38e82c9234f5eceee38d4e99a30a.AssociateWithHubSitesRequestBuilder) {
    return i48e086ed64fcd59d03281d5ffa804a5c862c38e82c9234f5eceee38d4e99a30a.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base the base property
func (m *ContentTypeItemRequestBuilder) Base()(*i7bed78d5c7b182bb6b35c0cbd54072e0caa68f8aaee7b2c72ed1530d26e66641.BaseRequestBuilder) {
    return i7bed78d5c7b182bb6b35c0cbd54072e0caa68f8aaee7b2c72ed1530d26e66641.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes the baseTypes property
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i929fc0841e147d731d330fffac2b603f59d2e7ee09ba0bb76a82c060f484478d.BaseTypesRequestBuilder) {
    return i929fc0841e147d731d330fffac2b603f59d2e7ee09ba0bb76a82c060f484478d.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i47b1c489561ff1669cae57a9dd84a767bbed7f82b77c0cdc7389dc73950c6be7.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return i47b1c489561ff1669cae57a9dd84a767bbed7f82b77c0cdc7389dc73950c6be7.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks the columnLinks property
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i5591abfce713804f2fb74b208efc6db0ee225ff736b8e9ac134955aad253bf7d.ColumnLinksRequestBuilder) {
    return i5591abfce713804f2fb74b208efc6db0ee225ff736b8e9ac134955aad253bf7d.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*if1bf402e4bfc313e6d11fabbebde44dffbe81a5aa7a12e657549b48143aefbb0.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return if1bf402e4bfc313e6d11fabbebde44dffbe81a5aa7a12e657549b48143aefbb0.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions the columnPositions property
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*ieabef4a072848f73eb7851ae8297bf4213dbcc6b84e915fef10453e9c720ea90.ColumnPositionsRequestBuilder) {
    return ieabef4a072848f73eb7851ae8297bf4213dbcc6b84e915fef10453e9c720ea90.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i49c22f7846eb325805a3ce05072fb08dfdd09048ded457ecb9796b7a866ec2dd.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i49c22f7846eb325805a3ce05072fb08dfdd09048ded457ecb9796b7a866ec2dd.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns the columns property
func (m *ContentTypeItemRequestBuilder) Columns()(*icc9cf5371dacb763ed32f4832e68daf65622d2d8a427a8b01fd79665f253c4cc.ColumnsRequestBuilder) {
    return icc9cf5371dacb763ed32f4832e68daf65622d2d8a427a8b01fd79665f253c4cc.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*ib1338a26b88df4781df2ab470814626fd7c5b61709da7294d63f28bea3263a37.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return ib1338a26b88df4781df2ab470814626fd7c5b61709da7294d63f28bea3263a37.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/list/contentTypes/{contentType%2Did}{?%24select,%24expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i2e24155317ddebf9f80799a7e7567793da2868054b1116675108e3b4674a0fc1.CopyToDefaultContentLocationRequestBuilder) {
    return i2e24155317ddebf9f80799a7e7567793da2868054b1116675108e3b4674a0fc1.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for drive
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property contentTypes for drive
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
// CreatePatchRequestInformation update the navigation property contentTypes in drive
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property contentTypes in drive
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
// Delete delete navigation property contentTypes for drive
func (m *ContentTypeItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property contentTypes for drive
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
func (m *ContentTypeItemRequestBuilder) IsPublished()(*ib19674cd06300cabca5bf2c5ea72a1bbc22118fa8aabb80e98be0a419c35e6cc.IsPublishedRequestBuilder) {
    return ib19674cd06300cabca5bf2c5ea72a1bbc22118fa8aabb80e98be0a419c35e6cc.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in drive
func (m *ContentTypeItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property contentTypes in drive
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
func (m *ContentTypeItemRequestBuilder) Publish()(*i072c7323df45db1f1ee748560abd200a43db5ebff2606b32cdad943a8247a3e3.PublishRequestBuilder) {
    return i072c7323df45db1f1ee748560abd200a43db5ebff2606b32cdad943a8247a3e3.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i69a0f0826d29adb7fdeff42634c0cfeaa9dfca0c790eb1d7972dc8275fdcd7bb.UnpublishRequestBuilder) {
    return i69a0f0826d29adb7fdeff42634c0cfeaa9dfca0c790eb1d7972dc8275fdcd7bb.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
