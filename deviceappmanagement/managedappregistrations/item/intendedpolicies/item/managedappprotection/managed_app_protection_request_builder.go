package managedappprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ic4c47e30236fb70b5ce4e6c72d1bff76bb2987e814ba97bd5e1c766d4d964fb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item/managedappprotection/targetapps"
)

// managedAppProtectionRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedAppRegistrations\{managedAppRegistration-id}\intendedPolicies\{managedAppPolicy-id}\microsoft.graph.managedAppProtection
type ManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NewManagedAppProtectionRequestBuilderInternal instantiates a new ManagedAppProtectionRequestBuilder and sets the default values.
func NewManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppProtectionRequestBuilder) {
    m := &ManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration_id}/intendedPolicies/{managedAppPolicy_id}/microsoft.graph.managedAppProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedAppProtectionRequestBuilder instantiates a new ManagedAppProtectionRequestBuilder and sets the default values.
func NewManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagedAppProtectionRequestBuilder) TargetApps()(*ic4c47e30236fb70b5ce4e6c72d1bff76bb2987e814ba97bd5e1c766d4d964fb7.TargetAppsRequestBuilder) {
    return ic4c47e30236fb70b5ce4e6c72d1bff76bb2987e814ba97bd5e1c766d4d964fb7.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
