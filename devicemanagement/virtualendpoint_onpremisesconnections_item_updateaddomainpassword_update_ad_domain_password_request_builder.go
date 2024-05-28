package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder provides operations to call the updateAdDomainPassword method.
type VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilderInternal instantiates a new VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder and sets the default values.
func NewVirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder) {
    m := &VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/onPremisesConnections/{cloudPcOnPremisesConnection%2Did}/updateAdDomainPassword", pathParameters),
    }
    return m
}
// NewVirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder instantiates a new VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder and sets the default values.
func NewVirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the Active Directory domain password for a cloudPcOnPremisesConnection object. This API is supported when the type of the cloudPcOnPremisesConnection object is hybridAzureADJoin.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpconpremisesconnection-updateaddomainpassword?view=graph-rest-beta
func (m *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder) Post(ctx context.Context, body VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBodyable, requestConfiguration *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation update the Active Directory domain password for a cloudPcOnPremisesConnection object. This API is supported when the type of the cloudPcOnPremisesConnection object is hybridAzureADJoin.
// returns a *RequestInformation when successful
func (m *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordPostRequestBodyable, requestConfiguration *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder when successful
func (m *VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder) {
    return NewVirtualendpointOnpremisesconnectionsItemUpdateaddomainpasswordUpdateAdDomainPasswordRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
