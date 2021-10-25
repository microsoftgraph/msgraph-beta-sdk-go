package incompatibleaccesspackages

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ia8a1f20bc9f473a476c8cca5fc08383ca89a4b88d195361c51400af2f500b035 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/incompatibleaccesspackages/ref"
    ic40c3381e5e05290b1bd0f5dc409c253d35cab481bd5b56307041bd50023d7bb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/incompatibleaccesspackages/filterbycurrentuserwithon"
    id4612aa41a2d3fba51bcaa77d5d43428a37eb65776baf7760a7f2556f9fc2884 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/incompatibleaccesspackages/search"
)

type IncompatibleAccessPackagesRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type IncompatibleAccessPackagesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escaped []string;
    Skip *int32;
    Top *int32;
}
func NewIncompatibleAccessPackagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IncompatibleAccessPackagesRequestBuilder) {
    m := &IncompatibleAccessPackagesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment_id}/accessPackage/incompatibleAccessPackages{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewIncompatibleAccessPackagesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IncompatibleAccessPackagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIncompatibleAccessPackagesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *IncompatibleAccessPackagesRequestBuilder) CreateGetRequestInformation(q func (value *IncompatibleAccessPackagesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(IncompatibleAccessPackagesRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *IncompatibleAccessPackagesRequestBuilder) FilterByCurrentUserWithOn(on *string)(*ic40c3381e5e05290b1bd0f5dc409c253d35cab481bd5b56307041bd50023d7bb.FilterByCurrentUserWithOnRequestBuilder) {
    return ic40c3381e5e05290b1bd0f5dc409c253d35cab481bd5b56307041bd50023d7bb.NewFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on);
}
func (m *IncompatibleAccessPackagesRequestBuilder) Get(q func (value *IncompatibleAccessPackagesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*IncompatibleAccessPackagesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIncompatibleAccessPackagesResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*IncompatibleAccessPackagesResponse), nil
}
func (m *IncompatibleAccessPackagesRequestBuilder) Ref()(*ia8a1f20bc9f473a476c8cca5fc08383ca89a4b88d195361c51400af2f500b035.RefRequestBuilder) {
    return ia8a1f20bc9f473a476c8cca5fc08383ca89a4b88d195361c51400af2f500b035.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IncompatibleAccessPackagesRequestBuilder) Search()(*id4612aa41a2d3fba51bcaa77d5d43428a37eb65776baf7760a7f2556f9fc2884.SearchRequestBuilder) {
    return id4612aa41a2d3fba51bcaa77d5d43428a37eb65776baf7760a7f2556f9fc2884.NewSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
