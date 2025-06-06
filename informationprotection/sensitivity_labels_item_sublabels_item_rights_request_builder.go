// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SensitivityLabelsItemSublabelsItemRightsRequestBuilder provides operations to manage the rights property of the microsoft.graph.sensitivityLabel entity.
type SensitivityLabelsItemSublabelsItemRightsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SensitivityLabelsItemSublabelsItemRightsRequestBuilderGetQueryParameters get rights from informationProtection
type SensitivityLabelsItemSublabelsItemRightsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SensitivityLabelsItemSublabelsItemRightsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SensitivityLabelsItemSublabelsItemRightsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SensitivityLabelsItemSublabelsItemRightsRequestBuilderGetQueryParameters
}
// NewSensitivityLabelsItemSublabelsItemRightsRequestBuilderInternal instantiates a new SensitivityLabelsItemSublabelsItemRightsRequestBuilder and sets the default values.
func NewSensitivityLabelsItemSublabelsItemRightsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SensitivityLabelsItemSublabelsItemRightsRequestBuilder) {
    m := &SensitivityLabelsItemSublabelsItemRightsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/informationProtection/sensitivityLabels/{sensitivityLabel%2Did}/sublabels/{sensitivityLabel%2Did1}/rights{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewSensitivityLabelsItemSublabelsItemRightsRequestBuilder instantiates a new SensitivityLabelsItemSublabelsItemRightsRequestBuilder and sets the default values.
func NewSensitivityLabelsItemSublabelsItemRightsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SensitivityLabelsItemSublabelsItemRightsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSensitivityLabelsItemSublabelsItemRightsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get rights from informationProtection
// returns a UsageRightsIncludedable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SensitivityLabelsItemSublabelsItemRightsRequestBuilder) Get(ctx context.Context, requestConfiguration *SensitivityLabelsItemSublabelsItemRightsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsageRightsIncludedable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUsageRightsIncludedFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsageRightsIncludedable), nil
}
// ToGetRequestInformation get rights from informationProtection
// returns a *RequestInformation when successful
func (m *SensitivityLabelsItemSublabelsItemRightsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SensitivityLabelsItemSublabelsItemRightsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SensitivityLabelsItemSublabelsItemRightsRequestBuilder when successful
func (m *SensitivityLabelsItemSublabelsItemRightsRequestBuilder) WithUrl(rawUrl string)(*SensitivityLabelsItemSublabelsItemRightsRequestBuilder) {
    return NewSensitivityLabelsItemSublabelsItemRightsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
