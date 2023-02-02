package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilder provides operations to call the getSharedUseLicenseUsageReport method.
type VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilderInternal instantiates a new GetSharedUseLicenseUsageReportRequestBuilder and sets the default values.
func NewVirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilder) {
    m := &VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/reports/microsoft.graph.getSharedUseLicenseUsageReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewVirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilder instantiates a new GetSharedUseLicenseUsageReportRequestBuilder and sets the default values.
func NewVirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get a usage report on shared-use licenses, such as **servicePlanId**, **licenseCount**, and **claimedLicenseCount**, for real-time, 7 days, or 28 days trend.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/cloudpcreports-getshareduselicenseusagereport?view=graph-rest-1.0
func (m *VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilder) Post(ctx context.Context, body VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportPostRequestBodyable, requestConfiguration *VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation get a usage report on shared-use licenses, such as **servicePlanId**, **licenseCount**, and **claimedLicenseCount**, for real-time, 7 days, or 28 days trend.
func (m *VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportPostRequestBodyable, requestConfiguration *VirtualEndpointReportsMicrosoftGraphGetSharedUseLicenseUsageReportGetSharedUseLicenseUsageReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
