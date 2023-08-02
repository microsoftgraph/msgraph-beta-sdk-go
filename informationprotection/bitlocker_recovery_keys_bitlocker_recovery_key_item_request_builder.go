package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder provides operations to manage the recoveryKeys property of the microsoft.graph.bitlocker entity.
type BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a bitlockerRecoveryKey object.  By default, this operation does not return the key property that represents the actual recovery key. To include the key property in the response, use the $select OData query parameter. Including the $select query parameter triggers an Azure AD audit of the operation and generates an audit log. You can find the log in Azure AD audit logs under the KeyManagement category.
type BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetQueryParameters
}
// NewBitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderInternal instantiates a new BitlockerRecoveryKeyItemRequestBuilder and sets the default values.
func NewBitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder) {
    m := &BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/informationProtection/bitlocker/recoveryKeys/{bitlockerRecoveryKey%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewBitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder instantiates a new BitlockerRecoveryKeyItemRequestBuilder and sets the default values.
func NewBitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve the properties and relationships of a bitlockerRecoveryKey object.  By default, this operation does not return the key property that represents the actual recovery key. To include the key property in the response, use the $select OData query parameter. Including the $select query parameter triggers an Azure AD audit of the operation and generates an audit log. You can find the log in Azure AD audit logs under the KeyManagement category.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bitlockerrecoverykey-get?view=graph-rest-1.0
func (m *BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BitlockerRecoveryKeyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBitlockerRecoveryKeyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BitlockerRecoveryKeyable), nil
}
// ToGetRequestInformation retrieve the properties and relationships of a bitlockerRecoveryKey object.  By default, this operation does not return the key property that represents the actual recovery key. To include the key property in the response, use the $select OData query parameter. Including the $select query parameter triggers an Azure AD audit of the operation and generates an audit log. You can find the log in Azure AD audit logs under the KeyManagement category.
func (m *BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
