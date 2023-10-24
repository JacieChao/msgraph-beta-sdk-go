package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder provides operations to manage the assignedUsers property of the microsoft.graph.cloudPcProvisioningPolicyAssignment entity.
type VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilderGetQueryParameters the assignment targeted users for the provisioning policy. This list of users is computed based on assignments, licenses, group memberships, and policies. This property is read-only. Supports$expand.
type VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilderGetQueryParameters
}
// NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder) {
    m := &VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/provisioningPolicies/{cloudPcProvisioningPolicy%2Did}/assignments/{cloudPcProvisioningPolicyAssignment%2Did}/assignedUsers/{user%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the assignment targeted users for the provisioning policy. This list of users is computed based on assignments, licenses, group memberships, and policies. This property is read-only. Supports$expand.
// Deprecated: The onPremisesConnectionId property is deprecated and will stop returning on July 30, 2023. as of 2023-03/onPremisesConnectionId on 2023-03-16 and will be removed 2023-07-30
func (m *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// MailboxSettings the mailboxSettings property
func (m *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder) MailboxSettings()(*VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersItemMailboxSettingsRequestBuilder) {
    return NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersItemMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
func (m *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder) ServiceProvisioningErrors()(*VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersItemServiceProvisioningErrorsRequestBuilder) {
    return NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersItemServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the assignment targeted users for the provisioning policy. This list of users is computed based on assignments, licenses, group memberships, and policies. This property is read-only. Supports$expand.
// Deprecated: The onPremisesConnectionId property is deprecated and will stop returning on July 30, 2023. as of 2023-03/onPremisesConnectionId on 2023-03-16 and will be removed 2023-07-30
func (m *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The onPremisesConnectionId property is deprecated and will stop returning on July 30, 2023. as of 2023-03/onPremisesConnectionId on 2023-03-16 and will be removed 2023-07-30
func (m *VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder) {
    return NewVirtualEndpointProvisioningPoliciesItemAssignmentsItemAssignedUsersUserItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
