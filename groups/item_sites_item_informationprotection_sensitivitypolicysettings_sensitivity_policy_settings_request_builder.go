package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder provides operations to manage the sensitivityPolicySettings property of the microsoft.graph.informationProtection entity.
type ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetQueryParameters get sensitivityPolicySettings from groups
type ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetQueryParameters
}
// ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderInternal instantiates a new ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) {
    m := &ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/sensitivityPolicySettings{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder instantiates a new ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sensitivityPolicySettings for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get sensitivityPolicySettings from groups
// returns a SensitivityPolicySettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityPolicySettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSensitivityPolicySettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityPolicySettingsable), nil
}
// Patch update the navigation property sensitivityPolicySettings in groups
// returns a SensitivityPolicySettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityPolicySettingsable, requestConfiguration *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityPolicySettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSensitivityPolicySettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityPolicySettingsable), nil
}
// ToDeleteRequestInformation delete navigation property sensitivityPolicySettings for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get sensitivityPolicySettings from groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property sensitivityPolicySettings in groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityPolicySettingsable, requestConfiguration *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder) {
    return NewItemSitesItemInformationprotectionSensitivitypolicysettingsSensitivityPolicySettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
