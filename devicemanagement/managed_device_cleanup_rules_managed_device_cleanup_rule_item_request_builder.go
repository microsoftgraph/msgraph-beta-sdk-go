package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder provides operations to manage the managedDeviceCleanupRules property of the microsoft.graph.deviceManagement entity.
type ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderGetQueryParameters device cleanup rule V2
type ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderGetQueryParameters
}
// ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderInternal instantiates a new ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder and sets the default values.
func NewManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder) {
    m := &ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDeviceCleanupRules/{managedDeviceCleanupRule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder instantiates a new ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder and sets the default values.
func NewManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedDeviceCleanupRules for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get device cleanup rule V2
// returns a ManagedDeviceCleanupRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCleanupRuleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceCleanupRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCleanupRuleable), nil
}
// Patch update the navigation property managedDeviceCleanupRules in deviceManagement
// returns a ManagedDeviceCleanupRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCleanupRuleable, requestConfiguration *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCleanupRuleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceCleanupRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCleanupRuleable), nil
}
// ToDeleteRequestInformation delete navigation property managedDeviceCleanupRules for deviceManagement
// returns a *RequestInformation when successful
func (m *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/managedDeviceCleanupRules/{managedDeviceCleanupRule%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation device cleanup rule V2
// returns a *RequestInformation when successful
func (m *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedDeviceCleanupRules in deviceManagement
// returns a *RequestInformation when successful
func (m *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCleanupRuleable, requestConfiguration *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/managedDeviceCleanupRules/{managedDeviceCleanupRule%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder when successful
func (m *ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder) WithUrl(rawUrl string)(*ManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder) {
    return NewManagedDeviceCleanupRulesManagedDeviceCleanupRuleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
