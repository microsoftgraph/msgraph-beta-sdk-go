package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder provides operations to manage the microsoftTunnelServerLogCollectionResponses property of the microsoft.graph.deviceManagement entity.
type MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetQueryParameters collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
type MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetQueryParameters
}
// MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderInternal instantiates a new MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder and sets the default values.
func NewMicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) {
    m := &MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelServerLogCollectionResponses/{microsoftTunnelServerLogCollectionResponse%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder instantiates a new MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder and sets the default values.
func NewMicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDownloadUrl provides operations to call the createDownloadUrl method.
// returns a *MicrosofttunnelserverlogcollectionresponsesItemCreatedownloadurlCreateDownloadUrlRequestBuilder when successful
func (m *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) CreateDownloadUrl()(*MicrosofttunnelserverlogcollectionresponsesItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    return NewMicrosofttunnelserverlogcollectionresponsesItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property microsoftTunnelServerLogCollectionResponses for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// GenerateDownloadUrl provides operations to call the generateDownloadUrl method.
// returns a *MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder when successful
func (m *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) GenerateDownloadUrl()(*MicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilder) {
    return NewMicrosofttunnelserverlogcollectionresponsesItemGeneratedownloadurlGenerateDownloadUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
// returns a MicrosoftTunnelServerLogCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerLogCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelServerLogCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerLogCollectionResponseable), nil
}
// Patch update the navigation property microsoftTunnelServerLogCollectionResponses in deviceManagement
// returns a MicrosoftTunnelServerLogCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerLogCollectionResponseable, requestConfiguration *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerLogCollectionResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelServerLogCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerLogCollectionResponseable), nil
}
// ToDeleteRequestInformation delete navigation property microsoftTunnelServerLogCollectionResponses for deviceManagement
// returns a *RequestInformation when successful
func (m *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
// returns a *RequestInformation when successful
func (m *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property microsoftTunnelServerLogCollectionResponses in deviceManagement
// returns a *RequestInformation when successful
func (m *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerLogCollectionResponseable, requestConfiguration *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder when successful
func (m *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) WithUrl(rawUrl string)(*MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) {
    return NewMicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
