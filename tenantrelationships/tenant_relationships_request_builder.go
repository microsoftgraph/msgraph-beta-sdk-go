package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2a0d4734956a5e963566aae111af73e279c3d5379a451b3d2c44cd13b01a47a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants"
    i764e92d08cb9718848f3e0ba8dbe891a44a52ede8e667e46244db3f6c20da77e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadmincustomers"
    i83ea38918d7b5bf5d9afba525ad5cb3ea378a623d3e00cf8109b420010f0c284 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships"
    i8c09ab3d6b642d8b05be3a28136bd617177588be11ced0b67fe87bffb07179c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/findtenantinformationbytenantidwithtenantid"
    ie0323efc36fbbfde058302b5282d95ac9e55f5ae03e3970436289fe9f2589c3b "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/findtenantinformationbydomainnamewithdomainname"
    i794f862ecc921099af0c35b23c015b58bfdfa8b986158639b5bbfb2aacaea17e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadmincustomers/item"
    ia7f4fac9af2719d5ab9e4339a8b2f55b843d5b618e4117c5460ccbfca8c2f196 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item"
)

// TenantRelationshipsRequestBuilder provides operations to manage the tenantRelationship singleton.
type TenantRelationshipsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsRequestBuilderGetQueryParameters get tenantRelationships
type TenantRelationshipsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantRelationshipsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsRequestBuilderGetQueryParameters
}
// TenantRelationshipsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTenantRelationshipsRequestBuilderInternal instantiates a new TenantRelationshipsRequestBuilder and sets the default values.
func NewTenantRelationshipsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsRequestBuilder) {
    m := &TenantRelationshipsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantRelationshipsRequestBuilder instantiates a new TenantRelationshipsRequestBuilder and sets the default values.
func NewTenantRelationshipsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get tenantRelationships
func (m *TenantRelationshipsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update tenantRelationships
func (m *TenantRelationshipsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantRelationshipable, requestConfiguration *TenantRelationshipsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DelegatedAdminCustomers provides operations to manage the delegatedAdminCustomers property of the microsoft.graph.tenantRelationship entity.
func (m *TenantRelationshipsRequestBuilder) DelegatedAdminCustomers()(*i764e92d08cb9718848f3e0ba8dbe891a44a52ede8e667e46244db3f6c20da77e.DelegatedAdminCustomersRequestBuilder) {
    return i764e92d08cb9718848f3e0ba8dbe891a44a52ede8e667e46244db3f6c20da77e.NewDelegatedAdminCustomersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DelegatedAdminCustomersById provides operations to manage the delegatedAdminCustomers property of the microsoft.graph.tenantRelationship entity.
func (m *TenantRelationshipsRequestBuilder) DelegatedAdminCustomersById(id string)(*i794f862ecc921099af0c35b23c015b58bfdfa8b986158639b5bbfb2aacaea17e.DelegatedAdminCustomerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminCustomer%2Did"] = id
    }
    return i794f862ecc921099af0c35b23c015b58bfdfa8b986158639b5bbfb2aacaea17e.NewDelegatedAdminCustomerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DelegatedAdminRelationships provides operations to manage the delegatedAdminRelationships property of the microsoft.graph.tenantRelationship entity.
func (m *TenantRelationshipsRequestBuilder) DelegatedAdminRelationships()(*i83ea38918d7b5bf5d9afba525ad5cb3ea378a623d3e00cf8109b420010f0c284.DelegatedAdminRelationshipsRequestBuilder) {
    return i83ea38918d7b5bf5d9afba525ad5cb3ea378a623d3e00cf8109b420010f0c284.NewDelegatedAdminRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DelegatedAdminRelationshipsById provides operations to manage the delegatedAdminRelationships property of the microsoft.graph.tenantRelationship entity.
func (m *TenantRelationshipsRequestBuilder) DelegatedAdminRelationshipsById(id string)(*ia7f4fac9af2719d5ab9e4339a8b2f55b843d5b618e4117c5460ccbfca8c2f196.DelegatedAdminRelationshipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminRelationship%2Did"] = id
    }
    return ia7f4fac9af2719d5ab9e4339a8b2f55b843d5b618e4117c5460ccbfca8c2f196.NewDelegatedAdminRelationshipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FindTenantInformationByDomainNameWithDomainName provides operations to call the findTenantInformationByDomainName method.
func (m *TenantRelationshipsRequestBuilder) FindTenantInformationByDomainNameWithDomainName(domainName *string)(*ie0323efc36fbbfde058302b5282d95ac9e55f5ae03e3970436289fe9f2589c3b.FindTenantInformationByDomainNameWithDomainNameRequestBuilder) {
    return ie0323efc36fbbfde058302b5282d95ac9e55f5ae03e3970436289fe9f2589c3b.NewFindTenantInformationByDomainNameWithDomainNameRequestBuilderInternal(m.pathParameters, m.requestAdapter, domainName);
}
// FindTenantInformationByTenantIdWithTenantId provides operations to call the findTenantInformationByTenantId method.
func (m *TenantRelationshipsRequestBuilder) FindTenantInformationByTenantIdWithTenantId(tenantId *string)(*i8c09ab3d6b642d8b05be3a28136bd617177588be11ced0b67fe87bffb07179c4.FindTenantInformationByTenantIdWithTenantIdRequestBuilder) {
    return i8c09ab3d6b642d8b05be3a28136bd617177588be11ced0b67fe87bffb07179c4.NewFindTenantInformationByTenantIdWithTenantIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, tenantId);
}
// Get get tenantRelationships
func (m *TenantRelationshipsRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantRelationshipable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTenantRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantRelationshipable), nil
}
// ManagedTenants provides operations to manage the managedTenants property of the microsoft.graph.tenantRelationship entity.
func (m *TenantRelationshipsRequestBuilder) ManagedTenants()(*i2a0d4734956a5e963566aae111af73e279c3d5379a451b3d2c44cd13b01a47a0.ManagedTenantsRequestBuilder) {
    return i2a0d4734956a5e963566aae111af73e279c3d5379a451b3d2c44cd13b01a47a0.NewManagedTenantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update tenantRelationships
func (m *TenantRelationshipsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantRelationshipable, requestConfiguration *TenantRelationshipsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantRelationshipable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTenantRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantRelationshipable), nil
}
