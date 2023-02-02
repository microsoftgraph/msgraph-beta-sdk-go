package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder provides operations to manage the appLogCollectionRequests property of the microsoft.graph.mobileAppTroubleshootingEvent entity.
type MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters the collection property of AppLogUploadRequest.
type MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetQueryParameters
}
// MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderInternal instantiates a new AppLogCollectionRequestItemRequestBuilder and sets the default values.
func NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, appLogCollectionRequestId *string)(*MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) {
    m := &MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent%2Did}/appLogCollectionRequests/{appLogCollectionRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if appLogCollectionRequestId != nil {
        urlTplParams["appLogCollectionRequest%2Did"] = *appLogCollectionRequestId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder instantiates a new AppLogCollectionRequestItemRequestBuilder and sets the default values.
func NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property appLogCollectionRequests for me
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection property of AppLogUploadRequest.
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppLogCollectionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable), nil
}
// MicrosoftGraphCreateDownloadUrl provides operations to call the createDownloadUrl method.
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) MicrosoftGraphCreateDownloadUrl()(*MobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemMicrosoftGraphCreateDownloadUrlCreateDownloadUrlRequestBuilder) {
    return NewMobileAppTroubleshootingEventsItemAppLogCollectionRequestsItemMicrosoftGraphCreateDownloadUrlCreateDownloadUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property appLogCollectionRequests in me
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppLogCollectionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable), nil
}
// ToDeleteRequestInformation delete navigation property appLogCollectionRequests for me
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property appLogCollectionRequests in me
func (m *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppLogCollectionRequestable, requestConfiguration *MobileAppTroubleshootingEventsItemAppLogCollectionRequestsAppLogCollectionRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
