package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder provides operations to call the windowsDefenderUpdateSignatures method.
type ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilderInternal instantiates a new ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder and sets the default values.
func NewComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder) {
    m := &ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/windowsDefenderUpdateSignatures", pathParameters),
    }
    return m
}
// NewComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder instantiates a new ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder and sets the default values.
func NewComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action windowsDefenderUpdateSignatures
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder) Post(ctx context.Context, requestConfiguration *ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action windowsDefenderUpdateSignatures
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder when successful
func (m *ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder) WithUrl(rawUrl string)(*ComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewComanageddevicesItemWindowsdefenderupdatesignaturesWindowsDefenderUpdateSignaturesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
