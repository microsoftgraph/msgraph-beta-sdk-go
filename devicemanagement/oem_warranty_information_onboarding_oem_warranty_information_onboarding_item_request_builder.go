package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder provides operations to manage the oemWarrantyInformationOnboarding property of the microsoft.graph.deviceManagement entity.
type OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetQueryParameters list of OEM Warranty Statuses
type OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetQueryParameters
}
// OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderInternal instantiates a new OemWarrantyInformationOnboardingItemRequestBuilder and sets the default values.
func NewOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) {
    m := &OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/oemWarrantyInformationOnboarding/{oemWarrantyInformationOnboarding%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder instantiates a new OemWarrantyInformationOnboardingItemRequestBuilder and sets the default values.
func NewOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property oemWarrantyInformationOnboarding for deviceManagement
func (m *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Disable provides operations to call the disable method.
func (m *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) Disable()(*OemWarrantyInformationOnboardingItemDisableRequestBuilder) {
    return NewOemWarrantyInformationOnboardingItemDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Enable provides operations to call the enable method.
func (m *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) Enable()(*OemWarrantyInformationOnboardingItemEnableRequestBuilder) {
    return NewOemWarrantyInformationOnboardingItemEnableRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get list of OEM Warranty Statuses
func (m *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OemWarrantyInformationOnboardingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOemWarrantyInformationOnboardingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OemWarrantyInformationOnboardingable), nil
}
// Patch update the navigation property oemWarrantyInformationOnboarding in deviceManagement
func (m *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OemWarrantyInformationOnboardingable, requestConfiguration *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OemWarrantyInformationOnboardingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOemWarrantyInformationOnboardingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OemWarrantyInformationOnboardingable), nil
}
// ToDeleteRequestInformation delete navigation property oemWarrantyInformationOnboarding for deviceManagement
func (m *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation list of OEM Warranty Statuses
func (m *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property oemWarrantyInformationOnboarding in deviceManagement
func (m *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OemWarrantyInformationOnboardingable, requestConfiguration *OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
