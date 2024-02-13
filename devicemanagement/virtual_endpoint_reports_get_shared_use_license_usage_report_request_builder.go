package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder provides operations to call the getSharedUseLicenseUsageReport method.
type VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilderInternal instantiates a new VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder) {
    m := &VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getSharedUseLicenseUsageReport", pathParameters),
    }
    return m
}
// NewVirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder instantiates a new VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get a usage report on shared-use licenses, such as servicePlanId, licenseCount, and claimedLicenseCount, for real-time, 7 days, or 28 days trend.
// Deprecated: The getSharedUseLicenseUsageReport API is deprecated and will stop returning on Oct 17, 2023. Please use getFrontlineReport instead. as of 2023-05/getSharedUseLicenseUsageReport
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcreports-getshareduselicenseusagereport?view=graph-rest-1.0
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder) Post(ctx context.Context, body VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToPostRequestInformation get a usage report on shared-use licenses, such as servicePlanId, licenseCount, and claimedLicenseCount, for real-time, 7 days, or 28 days trend.
// Deprecated: The getSharedUseLicenseUsageReport API is deprecated and will stop returning on Oct 17, 2023. Please use getFrontlineReport instead. as of 2023-05/getSharedUseLicenseUsageReport
// returns a *RequestInformation when successful
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointReportsGetSharedUseLicenseUsageReportPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The getSharedUseLicenseUsageReport API is deprecated and will stop returning on Oct 17, 2023. Please use getFrontlineReport instead. as of 2023-05/getSharedUseLicenseUsageReport
// returns a *VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder when successful
func (m *VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder) {
    return NewVirtualEndpointReportsGetSharedUseLicenseUsageReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
