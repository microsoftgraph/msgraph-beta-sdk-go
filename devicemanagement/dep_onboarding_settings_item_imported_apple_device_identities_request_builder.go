package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder provides operations to manage the importedAppleDeviceIdentities property of the microsoft.graph.depOnboardingSetting entity.
type DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderGetQueryParameters the imported Apple device identities.
type DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderGetQueryParameters
}
// DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByImportedAppleDeviceIdentityId provides operations to manage the importedAppleDeviceIdentities property of the microsoft.graph.depOnboardingSetting entity.
// returns a *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportedAppleDeviceIdentityItemRequestBuilder when successful
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) ByImportedAppleDeviceIdentityId(importedAppleDeviceIdentityId string)(*DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportedAppleDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if importedAppleDeviceIdentityId != "" {
        urlTplParams["importedAppleDeviceIdentity%2Did"] = importedAppleDeviceIdentityId
    }
    return NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportedAppleDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderInternal instantiates a new DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) {
    m := &DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/importedAppleDeviceIdentities{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder instantiates a new DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesCountRequestBuilder when successful
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) Count()(*DepOnboardingSettingsItemImportedAppleDeviceIdentitiesCountRequestBuilder) {
    return NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the imported Apple device identities.
// returns a ImportedAppleDeviceIdentityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) Get(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImportedAppleDeviceIdentityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityCollectionResponseable), nil
}
// ImportAppleDeviceIdentityList provides operations to call the importAppleDeviceIdentityList method.
// returns a *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder when successful
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) ImportAppleDeviceIdentityList()(*DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilder) {
    return NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportAppleDeviceIdentityListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to importedAppleDeviceIdentities for deviceManagement
// returns a ImportedAppleDeviceIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable, requestConfiguration *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImportedAppleDeviceIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable), nil
}
// ToGetRequestInformation the imported Apple device identities.
// returns a *RequestInformation when successful
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to importedAppleDeviceIdentities for deviceManagement
// returns a *RequestInformation when successful
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedAppleDeviceIdentityable, requestConfiguration *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder when successful
func (m *DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) WithUrl(rawUrl string)(*DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) {
    return NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
