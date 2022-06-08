package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i097bfeda18c323e4c2c8ad2a58999a280c8282d0e51cffc0c46c63a73351825f "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list"
    i1201c3f43119d05817263421e5f147bf42b57283cb9f554ef5b144a491f89625 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/items"
    i2b03e5349f14f4e882066fc04d344abd4c4a92b60832759bc6992a3d82e123ca "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/root"
    i5243f1e851b438da7e97e810d9e472d42d508e31f347cb52482c0b6e22a96dc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/listitem"
    i9bb99583bee5cdfb8243ddb4415954a793021828f7cadbd6bd2f195a9d7d4ea7 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/site"
    id98d79401b99ae85b3bc19d7bb389aa17f2bfff9077e3e86b4de0856908138fd "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/driveitem"
    ife9d8be655ee27d89f6d70f3a9fe6d45a7f877b7c963a8bc1ad46f7c7a9cfee6 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/permission"
    if3f442e91af227cf13d8f8602aede2bf3ce68223487cdb569bebba5512713828 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/items/item"
)

// SharedDriveItemItemRequestBuilder provides operations to manage the collection of sharedDriveItem entities.
type SharedDriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SharedDriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharedDriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SharedDriveItemItemRequestBuilderGetQueryParameters access a shared DriveItem or a collection of shared items by using a **shareId** or sharing URL. To use a sharing URL with this API, your app needs to transform the URL into a sharing token.
type SharedDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SharedDriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharedDriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SharedDriveItemItemRequestBuilderGetQueryParameters
}
// SharedDriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharedDriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSharedDriveItemItemRequestBuilderInternal instantiates a new SharedDriveItemItemRequestBuilder and sets the default values.
func NewSharedDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharedDriveItemItemRequestBuilder) {
    m := &SharedDriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSharedDriveItemItemRequestBuilder instantiates a new SharedDriveItemItemRequestBuilder and sets the default values.
func NewSharedDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharedDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharedDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from shares
func (m *SharedDriveItemItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete entity from shares
func (m *SharedDriveItemItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *SharedDriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation access a shared DriveItem or a collection of shared items by using a **shareId** or sharing URL. To use a sharing URL with this API, your app needs to transform the URL into a sharing token.
func (m *SharedDriveItemItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration access a shared DriveItem or a collection of shared items by using a **shareId** or sharing URL. To use a sharing URL with this API, your app needs to transform the URL into a sharing token.
func (m *SharedDriveItemItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SharedDriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in shares
func (m *SharedDriveItemItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedDriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update entity in shares
func (m *SharedDriveItemItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedDriveItemable, requestConfiguration *SharedDriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete entity from shares
func (m *SharedDriveItemItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete entity from shares
func (m *SharedDriveItemItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *SharedDriveItemItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// DriveItem the driveItem property
func (m *SharedDriveItemItemRequestBuilder) DriveItem()(*id98d79401b99ae85b3bc19d7bb389aa17f2bfff9077e3e86b4de0856908138fd.DriveItemRequestBuilder) {
    return id98d79401b99ae85b3bc19d7bb389aa17f2bfff9077e3e86b4de0856908138fd.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get access a shared DriveItem or a collection of shared items by using a **shareId** or sharing URL. To use a sharing URL with this API, your app needs to transform the URL into a sharing token.
func (m *SharedDriveItemItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedDriveItemable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler access a shared DriveItem or a collection of shared items by using a **shareId** or sharing URL. To use a sharing URL with this API, your app needs to transform the URL into a sharing token.
func (m *SharedDriveItemItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *SharedDriveItemItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedDriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSharedDriveItemFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedDriveItemable), nil
}
// Items the items property
func (m *SharedDriveItemItemRequestBuilder) Items()(*i1201c3f43119d05817263421e5f147bf42b57283cb9f554ef5b144a491f89625.ItemsRequestBuilder) {
    return i1201c3f43119d05817263421e5f147bf42b57283cb9f554ef5b144a491f89625.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.items.item collection
func (m *SharedDriveItemItemRequestBuilder) ItemsById(id string)(*if3f442e91af227cf13d8f8602aede2bf3ce68223487cdb569bebba5512713828.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return if3f442e91af227cf13d8f8602aede2bf3ce68223487cdb569bebba5512713828.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// List the list property
func (m *SharedDriveItemItemRequestBuilder) List()(*i097bfeda18c323e4c2c8ad2a58999a280c8282d0e51cffc0c46c63a73351825f.ListRequestBuilder) {
    return i097bfeda18c323e4c2c8ad2a58999a280c8282d0e51cffc0c46c63a73351825f.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem the listItem property
func (m *SharedDriveItemItemRequestBuilder) ListItem()(*i5243f1e851b438da7e97e810d9e472d42d508e31f347cb52482c0b6e22a96dc0.ListItemRequestBuilder) {
    return i5243f1e851b438da7e97e810d9e472d42d508e31f347cb52482c0b6e22a96dc0.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in shares
func (m *SharedDriveItemItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedDriveItemable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update entity in shares
func (m *SharedDriveItemItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedDriveItemable, requestConfiguration *SharedDriveItemItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Permission the permission property
func (m *SharedDriveItemItemRequestBuilder) Permission()(*ife9d8be655ee27d89f6d70f3a9fe6d45a7f877b7c963a8bc1ad46f7c7a9cfee6.PermissionRequestBuilder) {
    return ife9d8be655ee27d89f6d70f3a9fe6d45a7f877b7c963a8bc1ad46f7c7a9cfee6.NewPermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Root the root property
func (m *SharedDriveItemItemRequestBuilder) Root()(*i2b03e5349f14f4e882066fc04d344abd4c4a92b60832759bc6992a3d82e123ca.RootRequestBuilder) {
    return i2b03e5349f14f4e882066fc04d344abd4c4a92b60832759bc6992a3d82e123ca.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Site the site property
func (m *SharedDriveItemItemRequestBuilder) Site()(*i9bb99583bee5cdfb8243ddb4415954a793021828f7cadbd6bd2f195a9d7d4ea7.SiteRequestBuilder) {
    return i9bb99583bee5cdfb8243ddb4415954a793021828f7cadbd6bd2f195a9d7d4ea7.NewSiteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
