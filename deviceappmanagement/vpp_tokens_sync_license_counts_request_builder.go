package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VppTokensSyncLicenseCountsRequestBuilder provides operations to call the syncLicenseCounts method.
type VppTokensSyncLicenseCountsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VppTokensSyncLicenseCountsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VppTokensSyncLicenseCountsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVppTokensSyncLicenseCountsRequestBuilderInternal instantiates a new SyncLicenseCountsRequestBuilder and sets the default values.
func NewVppTokensSyncLicenseCountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VppTokensSyncLicenseCountsRequestBuilder) {
    m := &VppTokensSyncLicenseCountsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/vppTokens/syncLicenseCounts", pathParameters),
    }
    return m
}
// NewVppTokensSyncLicenseCountsRequestBuilder instantiates a new SyncLicenseCountsRequestBuilder and sets the default values.
func NewVppTokensSyncLicenseCountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VppTokensSyncLicenseCountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVppTokensSyncLicenseCountsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action syncLicenseCounts
func (m *VppTokensSyncLicenseCountsRequestBuilder) Post(ctx context.Context, requestConfiguration *VppTokensSyncLicenseCountsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action syncLicenseCounts
func (m *VppTokensSyncLicenseCountsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VppTokensSyncLicenseCountsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VppTokensSyncLicenseCountsRequestBuilder) WithUrl(rawUrl string)(*VppTokensSyncLicenseCountsRequestBuilder) {
    return NewVppTokensSyncLicenseCountsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
