package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder provides operations to manage the appLogCollectionRequests property of the microsoft.graph.mobileAppTroubleshootingEvent entity.
type ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters indicates collection of App Log Upload Request.
type ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters
}
// ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderInternal instantiates a new ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder and sets the default values.
func NewItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) {
    m := &ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent%2Did}/appLogCollectionRequests/{appLogCollectionRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder instantiates a new ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder and sets the default values.
func NewItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDownloadUrl provides operations to call the createDownloadUrl method.
// returns a *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder when successful
func (m *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) CreateDownloadUrl()(*ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    return NewItemMobileapptroubleshootingeventsItemApplogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property appLogCollectionRequests for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get indicates collection of App Log Upload Request.
// returns a AppLogCollectionRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppLogCollectionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable), nil
}
// Patch update the navigation property appLogCollectionRequests in users
// returns a AppLogCollectionRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable, requestConfiguration *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppLogCollectionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable), nil
}
// ToDeleteRequestInformation delete navigation property appLogCollectionRequests for users
// returns a *RequestInformation when successful
func (m *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation indicates collection of App Log Upload Request.
// returns a *RequestInformation when successful
func (m *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property appLogCollectionRequests in users
// returns a *RequestInformation when successful
func (m *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable, requestConfiguration *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder when successful
func (m *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) WithUrl(rawUrl string)(*ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder) {
    return NewItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
