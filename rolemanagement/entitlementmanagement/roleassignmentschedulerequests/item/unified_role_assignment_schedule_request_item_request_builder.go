package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0a27aefac5e284898c6f31f3cec3cc5ca84d4aa4003ae193a4f82ea4c728e6a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/targetschedule"
    i2a944771b22cb5ea490413365394d6c74a1f803aba545a6f9b42769ec7319d09 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/appscope"
    i44d378692bae641c92a2c30cbfa78e06815a2cd19b7e2e10b8b725b11d8fc1cf "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/principal"
    i644c4ca2b29a486fb1565141c4c7f306a02458cb796d329032b2e0f088479442 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/activatedusing"
    iae59854ee3116fd1456d39cebc6a857e33eb8deea59f76066fe538101bc0f1b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/roledefinition"
    ibdcb7e68595c8a6e7db32816c49d6e162f0cf44239d2990965abb46b25451846 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/directoryscope"
    id18c7cf991b9a22e0a73384bb3123a5de24ee17868f1f7b779a94f0a9c5a6a3a "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/cancel"
)

// UnifiedRoleAssignmentScheduleRequestItemRequestBuilder provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters requests for active role assignments to principals through PIM.
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing the activatedUsing property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ActivatedUsing()(*i644c4ca2b29a486fb1565141c4c7f306a02458cb796d329032b2e0f088479442.ActivatedUsingRequestBuilder) {
    return i644c4ca2b29a486fb1565141c4c7f306a02458cb796d329032b2e0f088479442.NewActivatedUsingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppScope the appScope property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) AppScope()(*i2a944771b22cb5ea490413365394d6c74a1f803aba545a6f9b42769ec7319d09.AppScopeRequestBuilder) {
    return i2a944771b22cb5ea490413365394d6c74a1f803aba545a6f9b42769ec7319d09.NewAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Cancel()(*id18c7cf991b9a22e0a73384bb3123a5de24ee17868f1f7b779a94f0a9c5a6a3a.CancelRequestBuilder) {
    return id18c7cf991b9a22e0a73384bb3123a5de24ee17868f1f7b779a94f0a9c5a6a3a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    m := &UnifiedRoleAssignmentScheduleRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilder instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformationWithRequestConfiguration requests for active role assignments to principals through PIM.
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration requests for active role assignments to principals through PIM.
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleRequestable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleRequestable, requestConfiguration *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeleteWithResponseHandler delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DirectoryScope the directoryScope property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) DirectoryScope()(*ibdcb7e68595c8a6e7db32816c49d6e162f0cf44239d2990965abb46b25451846.DirectoryScopeRequestBuilder) {
    return ibdcb7e68595c8a6e7db32816c49d6e162f0cf44239d2990965abb46b25451846.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithResponseHandler requests for active role assignments to principals through PIM.
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) GetWithResponseHandler(requestConfiguration *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleRequestable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler requests for active role assignments to principals through PIM.
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) GetWithResponseHandler(requestConfiguration *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleRequestable), nil
}
// PatchWithResponseHandler update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleRequestable, requestConfiguration *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentScheduleRequestable, requestConfiguration *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Principal the principal property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Principal()(*i44d378692bae641c92a2c30cbfa78e06815a2cd19b7e2e10b8b725b11d8fc1cf.PrincipalRequestBuilder) {
    return i44d378692bae641c92a2c30cbfa78e06815a2cd19b7e2e10b8b725b11d8fc1cf.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinition the roleDefinition property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) RoleDefinition()(*iae59854ee3116fd1456d39cebc6a857e33eb8deea59f76066fe538101bc0f1b0.RoleDefinitionRequestBuilder) {
    return iae59854ee3116fd1456d39cebc6a857e33eb8deea59f76066fe538101bc0f1b0.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TargetSchedule the targetSchedule property
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) TargetSchedule()(*i0a27aefac5e284898c6f31f3cec3cc5ca84d4aa4003ae193a4f82ea4c728e6a4.TargetScheduleRequestBuilder) {
    return i0a27aefac5e284898c6f31f3cec3cc5ca84d4aa4003ae193a4f82ea4c728e6a4.NewTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
