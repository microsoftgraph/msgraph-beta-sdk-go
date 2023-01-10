package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder provides operations to manage the appLogCollectionRequests property of the microsoft.graph.mobileAppTroubleshootingEvent entity.
type ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters the collection property of AppLogUploadRequest.
type ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters
}
// ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderInternal instantiates a new AppLogCollectionRequestItemRequestBuilder and sets the default values.
func NewItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) {
    m := &ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent%2Did}/appLogCollectionRequests/{appLogCollectionRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder instantiates a new AppLogCollectionRequestItemRequestBuilder and sets the default values.
func NewItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDownloadUrl provides operations to call the createDownloadUrl method.
func (m *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) CreateDownloadUrl()(*ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) {
    return NewItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemCreateDownloadUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property appLogCollectionRequests for users
func (m *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get the collection property of AppLogUploadRequest.
func (m *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppLogCollectionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable), nil
}
// Patch update the navigation property appLogCollectionRequests in users
func (m *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable, requestConfiguration *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppLogCollectionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable), nil
}
// ToDeleteRequestInformation delete navigation property appLogCollectionRequests for users
func (m *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the collection property of AppLogUploadRequest.
func (m *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property appLogCollectionRequests in users
func (m *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable, requestConfiguration *ItemMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
