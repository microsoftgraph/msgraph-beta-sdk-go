package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
type EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters the embedded SIM activation code pools created by this account.
type EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters
}
// EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Assign()(*EmbeddedSIMActivationCodePoolsItemAssignRequestBuilder) {
    return NewEmbeddedSIMActivationCodePoolsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.embeddedSIMActivationCodePool entity.
func (m *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Assignments()(*EmbeddedSIMActivationCodePoolsItemAssignmentsRequestBuilder) {
    return NewEmbeddedSIMActivationCodePoolsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderInternal instantiates a new EmbeddedSIMActivationCodePoolItemRequestBuilder and sets the default values.
func NewEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) {
    m := &EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/embeddedSIMActivationCodePools/{embeddedSIMActivationCodePool%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder instantiates a new EmbeddedSIMActivationCodePoolItemRequestBuilder and sets the default values.
func NewEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property embeddedSIMActivationCodePools for deviceManagement
func (m *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceStates provides operations to manage the deviceStates property of the microsoft.graph.embeddedSIMActivationCodePool entity.
func (m *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) DeviceStates()(*EmbeddedSIMActivationCodePoolsItemDeviceStatesRequestBuilder) {
    return NewEmbeddedSIMActivationCodePoolsItemDeviceStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the embedded SIM activation code pools created by this account.
func (m *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable), nil
}
// Patch update the navigation property embeddedSIMActivationCodePools in deviceManagement
func (m *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, requestConfiguration *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable), nil
}
// ToDeleteRequestInformation delete navigation property embeddedSIMActivationCodePools for deviceManagement
func (m *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the embedded SIM activation code pools created by this account.
func (m *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property embeddedSIMActivationCodePools in deviceManagement
func (m *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, requestConfiguration *EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
