package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.user entity.
type ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetQueryParameters the list of mobile app troubleshooting events for this user.
type ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetQueryParameters
}
// ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppLogCollectionRequests provides operations to manage the appLogCollectionRequests property of the microsoft.graph.mobileAppTroubleshootingEvent entity.
// returns a *ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestsRequestBuilder when successful
func (m *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) AppLogCollectionRequests()(*ItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestsRequestBuilder) {
    return NewItemMobileapptroubleshootingeventsItemApplogcollectionrequestsAppLogCollectionRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderInternal instantiates a new ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder and sets the default values.
func NewItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) {
    m := &ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder instantiates a new ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder and sets the default values.
func NewItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mobileAppTroubleshootingEvents for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of mobile app troubleshooting events for this user.
// returns a MobileAppTroubleshootingEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppTroubleshootingEventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppTroubleshootingEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppTroubleshootingEventable), nil
}
// Patch update the navigation property mobileAppTroubleshootingEvents in users
// returns a MobileAppTroubleshootingEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppTroubleshootingEventable, requestConfiguration *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppTroubleshootingEventable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppTroubleshootingEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppTroubleshootingEventable), nil
}
// ToDeleteRequestInformation delete navigation property mobileAppTroubleshootingEvents for users
// returns a *RequestInformation when successful
func (m *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of mobile app troubleshooting events for this user.
// returns a *RequestInformation when successful
func (m *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mobileAppTroubleshootingEvents in users
// returns a *RequestInformation when successful
func (m *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppTroubleshootingEventable, requestConfiguration *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder when successful
func (m *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder) {
    return NewItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
