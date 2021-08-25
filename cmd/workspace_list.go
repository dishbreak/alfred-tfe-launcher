package main

import (
	"context"

	"github.com/dishbreak/terraform-cloud-launcher/lib"
	"github.com/dishbreak/terraform-cloud-launcher/models"
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

func (wl *workspaceLister) FetchWorkspaces() ([]models.ListItem, error) {
	items := make([]models.ListItem, 0)
	for nextPage := 1; nextPage != 0; {
		workspaceList, err := wl.List(context.Background(), "nerdwallet", paginate(nextPage))
		if err != nil {
			return nil, err
		}
		for _, workspace := range workspaceList.Items {
			items = append(items, models.ListItem{
				Title:    workspace.Name,
				Subtitle: workspace.Description,
				Arg:      workspace.Name,
				Valid:    true,
			})
		}
		nextPage = workspaceList.NextPage
	}

	return items, nil
}

func (w *WorkspaceListCmd) Run(ctx *Context) error {
	resp := models.NewScriptResponse()

	client, err := lib.NewTfeClient()
	if err != nil {
		resp.SetError(err)
		return err
	}

	lister := &workspaceLister{
		workspaceClient: client.Workspaces,
	}

	items, err := lister.FetchWorkspaces()
	if err != nil {
		resp.SetError(err)
		return err
	}

	for _, item := range items {
		resp.AddItem(item)
	}

	resp.SendFeedback()

	return nil
}
