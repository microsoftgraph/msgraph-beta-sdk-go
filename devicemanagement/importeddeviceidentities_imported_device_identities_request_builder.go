package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder provides operations to manage the importedDeviceIdentities property of the microsoft.graph.deviceManagement entity.
type ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderGetQueryParameters the imported device identities.
type ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderGetQueryParameters struct {
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
// ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderGetQueryParameters
}
// ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByImportedDeviceIdentityId provides operations to manage the importedDeviceIdentities property of the microsoft.graph.deviceManagement entity.
// returns a *ImporteddeviceidentitiesImportedDeviceIdentityItemRequestBuilder when successful
func (m *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder) ByImportedDeviceIdentityId(importedDeviceIdentityId string)(*ImporteddeviceidentitiesImportedDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if importedDeviceIdentityId != "" {
        urlTplParams["importedDeviceIdentity%2Did"] = importedDeviceIdentityId
    }
    return NewImporteddeviceidentitiesImportedDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderInternal instantiates a new ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder and sets the default values.
func NewImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder) {
    m := &ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/importedDeviceIdentities{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder instantiates a new ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder and sets the default values.
func NewImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ImporteddeviceidentitiesCountRequestBuilder when successful
func (m *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder) Count()(*ImporteddeviceidentitiesCountRequestBuilder) {
    return NewImporteddeviceidentitiesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the imported device identities.
// returns a ImportedDeviceIdentityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder) Get(ctx context.Context, requestConfiguration *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImportedDeviceIdentityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityCollectionResponseable), nil
}
// ImportDeviceIdentityList provides operations to call the importDeviceIdentityList method.
// returns a *ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder when successful
func (m *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder) ImportDeviceIdentityList()(*ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder) {
    return NewImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to importedDeviceIdentities for deviceManagement
// returns a ImportedDeviceIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable, requestConfiguration *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImportedDeviceIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable), nil
}
// SearchExistingIdentities provides operations to call the searchExistingIdentities method.
// returns a *ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder when successful
func (m *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder) SearchExistingIdentities()(*ImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilder) {
    return NewImporteddeviceidentitiesSearchexistingidentitiesSearchExistingIdentitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the imported device identities.
// returns a *RequestInformation when successful
func (m *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to importedDeviceIdentities for deviceManagement
// returns a *RequestInformation when successful
func (m *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable, requestConfiguration *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder when successful
func (m *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder) WithUrl(rawUrl string)(*ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder) {
    return NewImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
