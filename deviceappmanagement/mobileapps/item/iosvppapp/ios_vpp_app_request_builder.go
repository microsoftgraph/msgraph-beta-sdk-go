package iosvppapp

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i24b5eee907f3023ad6b32b470f7d6d36d9dbb47745fa42f740f7ec20c352f5db "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/iosvppapp/revokedevicelicense"
    i6e18af1428f7b694ac5766c2ccdc52630011ab10de0b02134747274ffb915696 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/iosvppapp/revokealllicenses"
    ieba7d49970f8ba88a7065c6eb5ffd1bee021c90edc2d1e7786356454063de65b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/iosvppapp/revokeuserlicense"
)

type IosVppAppRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func NewIosVppAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosVppAppRequestBuilder) {
    m := &IosVppAppRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceAppManagement/mobileApps/{mobileApp_id}/microsoft.graph.iosVppApp";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewIosVppAppRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosVppAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosVppAppRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *IosVppAppRequestBuilder) RevokeAllLicenses()(*i6e18af1428f7b694ac5766c2ccdc52630011ab10de0b02134747274ffb915696.RevokeAllLicensesRequestBuilder) {
    return i6e18af1428f7b694ac5766c2ccdc52630011ab10de0b02134747274ffb915696.NewRevokeAllLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosVppAppRequestBuilder) RevokeDeviceLicense()(*i24b5eee907f3023ad6b32b470f7d6d36d9dbb47745fa42f740f7ec20c352f5db.RevokeDeviceLicenseRequestBuilder) {
    return i24b5eee907f3023ad6b32b470f7d6d36d9dbb47745fa42f740f7ec20c352f5db.NewRevokeDeviceLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosVppAppRequestBuilder) RevokeUserLicense()(*ieba7d49970f8ba88a7065c6eb5ffd1bee021c90edc2d1e7786356454063de65b.RevokeUserLicenseRequestBuilder) {
    return ieba7d49970f8ba88a7065c6eb5ffd1bee021c90edc2d1e7786356454063de65b.NewRevokeUserLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
