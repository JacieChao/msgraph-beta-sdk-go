package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder provides operations to manage the internalSponsors property of the microsoft.graph.connectedOrganization entity.
type EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetQueryParameters retrieve a list of a connectedOrganization's internal sponsors.  The internal sponsors are a set of users who can approve requests on behalf of other users from that connected organization. This API is available in the following national cloud deployments.
type EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.connectedOrganizations.item.internalSponsors.item collection
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderInternal instantiates a new InternalSponsorsRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) {
    m := &EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/internalSponsors{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder instantiates a new InternalSponsorsRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) Count()(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsCountRequestBuilder) {
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of a connectedOrganization's internal sponsors.  The internal sponsors are a set of users who can approve requests on behalf of other users from that connected organization. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/connectedorganization-list-internalsponsors?view=graph-rest-1.0
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) Ref()(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsRefRequestBuilder) {
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation retrieve a list of a connectedOrganization's internal sponsors.  The internal sponsors are a set of users who can approve requests on behalf of other users from that connected organization. This API is available in the following national cloud deployments.
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder) {
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
