package managedappprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i8693f2d829b51c0a9cb7c1470270114f68b9728886ff3ebf3f5ceeab350fca8a "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedapppolicies/item/managedappprotection/targetapps"
)

type ManagedAppProtectionRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func NewManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppProtectionRequestBuilder) {
    m := &ManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceAppManagement/managedAppPolicies/{managedAppPolicy_id}/microsoft.graph.managedAppProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagedAppProtectionRequestBuilder) TargetApps()(*i8693f2d829b51c0a9cb7c1470270114f68b9728886ff3ebf3f5ceeab350fca8a.TargetAppsRequestBuilder) {
    return i8693f2d829b51c0a9cb7c1470270114f68b9728886ff3ebf3f5ceeab350fca8a.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
