package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder provides operations to manage the windowsUpdateCatalogItems property of the microsoft.graph.deviceManagement entity.
type WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderGetQueryParameters a collection of windows update catalog items (fetaure updates item , quality updates item)
type WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderGetQueryParameters
}
// WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderInternal instantiates a new WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder and sets the default values.
func NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder) {
    m := &WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsUpdateCatalogItems/{windowsUpdateCatalogItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder instantiates a new WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder and sets the default values.
func NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsUpdateCatalogItems for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of windows update catalog items (fetaure updates item , quality updates item)
// returns a WindowsUpdateCatalogItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUpdateCatalogItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsUpdateCatalogItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUpdateCatalogItemable), nil
}
// Patch update the navigation property windowsUpdateCatalogItems in deviceManagement
// returns a WindowsUpdateCatalogItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUpdateCatalogItemable, requestConfiguration *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUpdateCatalogItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsUpdateCatalogItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUpdateCatalogItemable), nil
}
// ToDeleteRequestInformation delete navigation property windowsUpdateCatalogItems for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of windows update catalog items (fetaure updates item , quality updates item)
// returns a *RequestInformation when successful
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsUpdateCatalogItems in deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUpdateCatalogItemable, requestConfiguration *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder when successful
func (m *WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder) WithUrl(rawUrl string)(*WindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder) {
    return NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
