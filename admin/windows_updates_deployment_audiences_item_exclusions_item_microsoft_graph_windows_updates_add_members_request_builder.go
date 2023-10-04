package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder provides operations to call the addMembers method.
type WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderInternal instantiates a new MicrosoftGraphWindowsUpdatesAddMembersRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) {
    m := &WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deploymentAudiences/{deploymentAudience%2Did}/exclusions/{updatableAsset%2Did}/microsoft.graph.windowsUpdates.addMembers", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder instantiates a new MicrosoftGraphWindowsUpdatesAddMembersRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add members to an updatableAssetGroup. You can add azureADDevice resources as members, but may not add updatableAssetGroup resources as members. Adding an Azure AD device as a member of an updatable asset group automatically creates an azureADDevice object, if it does not already exist. You can also use the method addMembersById to add members. This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-updatableassetgroup-addmembers?view=graph-rest-1.0
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) Post(ctx context.Context, body WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation add members to an updatableAssetGroup. You can add azureADDevice resources as members, but may not add updatableAssetGroup resources as members. Adding an Azure AD device as a member of an updatable asset group automatically creates an azureADDevice object, if it does not already exist. You can also use the method addMembersById to add members. This API is supported in the following national cloud deployments.
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder) {
    return NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
