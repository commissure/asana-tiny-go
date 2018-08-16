package asana

type Task struct {
	Assignee       interface{} `json:"assignee"`
	AssigneeStatus string      `json:"assignee_status"`
	Completed      bool        `json:"completed"`
	CompletedAt    interface{} `json:"completed_at"`
	CreatedAt      string      `json:"created_at"`
	DueAt          interface{} `json:"due_at"`
	DueOn          interface{} `json:"due_on"`
	Followers      []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"followers"`
	Hearted     bool          `json:"hearted"`
	Hearts      []interface{} `json:"hearts"`
	ID          int           `json:"id"`
	Liked       bool          `json:"liked"`
	Likes       []interface{} `json:"likes"`
	Memberships []struct {
		Project struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"project"`
		Section interface{} `json:"section"`
	} `json:"memberships"`
	ModifiedAt string      `json:"modified_at"`
	Name       string      `json:"name"`
	Notes      string      `json:"notes"`
	NumHearts  int         `json:"num_hearts"`
	NumLikes   int         `json:"num_likes"`
	Parent     interface{} `json:"parent"`
	Projects   []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"projects"`
	StartOn   interface{}   `json:"start_on"`
	Tags      []interface{} `json:"tags"`
	Workspace struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"workspace"`
}

type TaskRequest struct {
	Workspace   int64             `json:"workspace"`
	Name        string            `json:"name"`
	Notes       string            `json:"notes"`
	Followers   []*TaskFollower   `json:"followers,omitempty"`
	Memberships []*TaskMembership `json:"memberships,omitempty"`
}

type TaskMembership struct {
	Project int64 `json:"project"`
	Section int64 `json:"section,omitempty"`
}

type TaskFollower struct {
	ID int64 `json:"id"`
}
