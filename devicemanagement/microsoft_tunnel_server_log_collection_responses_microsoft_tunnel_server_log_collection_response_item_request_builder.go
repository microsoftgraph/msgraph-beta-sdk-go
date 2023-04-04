package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder provides operations to manage the microsoftTunnelServerLogCollectionResponses property of the microsoft.graph.deviceManagement entity.
type MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetQueryParameters collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
type MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetQueryParameters
}
// MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderInternal instantiates a new MicrosoftTunnelServerLogCollectionResponseItemRequestBuilder and sets the default values.
func NewMicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) {
    m := &MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelServerLogCollectionResponses/{microsoftTunnelServerLogCollectionResponse%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder instantiates a new MicrosoftTunnelServerLogCollectionResponseItemRequestBuilder and sets the default values.
func NewMicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDownloadUrl provides operations to call the createDownloadUrl method.
func (m *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) CreateDownloadUrl()(*MicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlRequestBuilder) {
    return NewMicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property microsoftTunnelServerLogCollectionResponses for deviceManagement
func (m *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// GenerateDownloadUrl provides operations to call the generateDownloadUrl method.
func (m *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) GenerateDownloadUrl()(*MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder) {
    return NewMicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
func (m *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerLogCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerLogCollectionResponseable, requestConfiguration *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerLogCollectionResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
func (m *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// ToPatchRequestInformation update the navigation property microsoftTunnelServerLogCollectionResponses in deviceManagement
func (m *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerLogCollectionResponseable, requestConfiguration *MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
