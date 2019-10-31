package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Project ...
type Project struct {
	ProjectID    int64             `orm:"pk;auto;column(project_id)" json:"project_id"`
	OwnerID      int               `orm:"column(owner_id)" json:"owner_id"`
	Name         string            `orm:"column(name)" json:"name"`
	CreationTime time.Time         `orm:"column(creation_time)" json:"creation_time"`
	UpdateTime   time.Time         `orm:"column(update_time)" json:"update_time"`
	Deleted      int               `orm:"column(deleted)" json:"deleted"`
	OwnerName    string            `orm:"-" json:"owner_name"`
	Togglable    bool              `orm:"-" json:"togglable"`
	Role         int               `orm:"-" json:"current_user_role_id"`
	RepoCount    int64             `orm:"-" json:"repo_count"`
	Metadata     map[string]string `orm:"-" json:"metadata"`
}

// SearchRepository ...
type SearchRepository struct {

	// The ID of the project that the repository belongs to
	ProjectId int32 `json:"project_id,omitempty"`

	// The name of the project that the repository belongs to
	ProjectName string `json:"project_name,omitempty"`

	// The flag to indicate the publicity of the project that the repository belongs to
	ProjectPublic bool `json:"project_public,omitempty"`

	// The name of the repository
	RepositoryName string `json:"repository_name,omitempty"`
}

// Search ...
type Search struct {

	// Search results of the projects that matched the filter keywords.
	Projects     []Project          `json:"project,omitempty"`
	Repositories []SearchRepository `json:"repository,omitempty"`
}

func main() {
body := []byte(`{
	"project": [
		{
		"project_id": 1,
		"owner_id": 1,
		"name": "library",
		"creation_time": "2018-04-27T17:17:37+08:00",
		"update_time": "2018-04-27T17:17:37+08:00",
		"deleted": 0,
		"owner_name": "",
		"togglable": false,
		"current_user_role_id": 0,
		"repo_count": 0,
		"metadata": {
			"public": "true"
		}
		}
	],
	"repository": []
	}`)
var successPayload = new(Search)
err := json.Unmarshal(body, &successPayload)

if err != nil {
	fmt.Printf("Failed to load json object!, err %v\n", err)
} else {
	fmt.Printf("Load json %v\n", successPayload)
}
}
