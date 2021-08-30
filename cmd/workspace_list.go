package main

import (
	"context"

	"github.com/dishbreak/terraform-cloud-launcher/lib"
	"github.com/hashicorp/go-tfe"
)

type WorkspaceListCmd struct{}

func (w *WorkspaceListCmd) Help() string {
	return `
Return a list of workspaces in Alfred JSON format.
	`
}

func paginate(nextPage int) tfe.WorkspaceListOptions {
	return tfe.WorkspaceListOptions{
		ListOptions: tfe.ListOptions{
			PageNumber: nextPage,
		},
	}
}

type workspaceClient interface {
	List(context.Context, string, tfe.WorkspaceListOptions) (*tfe.WorkspaceList, error)
}

type workspaceLister struct {
	workspaceClient
}

func (wl *workspaceLister) FetchWorkspaces() ([]lib.ListItem, error) {
	items := make([]lib.ListItem, 0)
	for nextPage := 1; nextPage != 0; {
		workspaceList, err := wl.List(context.Background(), "nerdwallet", paginate(nextPage))
		if err != nil {
			return nil, err
		}
		for _, workspace := range workspaceList.Items {
			items = append(items, lib.ListItem{
				Title:    workspace.Name,
				Subtitle: workspace.Description,
				Arg:      "",
				Valid:    true,
				Variables: map[string]string{
					"workspace_id":   workspace.ID,
					"workspace_name": workspace.Name,
					"workspace":      workspace.SourceURL,
				},
			})
		}
		nextPage = workspaceList.NextPage
	}

	return items, nil
}

func (w *WorkspaceListCmd) Run(ctx *Context) error {
	resp := lib.NewScriptFilterResponse()
	defer lib.RecoverIfErr(resp)

	client, err := lib.NewTfeClient()
	if err != nil {
		panic(err)
	}

	settings, err := lib.NewSettings()
	if err != nil {
		panic(err)
	}

	if err := settings.Load(); err != nil {
		panic(err)
	}

	lister := &workspaceLister{
		workspaceClient: client.Workspaces,
	}

	wsCache, err := lib.NewCache("workspaces", settings.CacheTimeout, lister.FetchWorkspaces)
	if err != nil {
		panic(err)
	}

	items, err := wsCache.Get()
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		resp.AddItem(item)
	}

	resp.SendFeedback()

	return nil
}
