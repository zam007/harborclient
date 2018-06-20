package sdk

// projects
type Project struct {
	Project_id           int             // Project ID
	Owner_id             int             // The owner ID of the project always means the creator of the project. ,
	Name                 string          // The name of the project. ,
	Creation_time        string          // The creation time of the project. ,
	Update_time          string          // The update time of the project. ,
	Deleted              int             // A deletion mark of the project (1 means it's deleted, 0 is not) ,
	Owner_name           string          // The owner name of the project. ,
	Togglable            bool            // Correspond to the UI about whether the project's publicity is updatable (for UI) ,
	Current_user_role_id int             // The role ID of the current user who triggered the API (for UI) ,
	Repo_count           int             // The number of the repositories under this project. ,
	Metadata             ProjectMetadata // The metadata of the project.
}

type ProjectMetadata struct {
	Public                                          int    // The public status of the project. ,
	Enable_content_trust                            bool   // Whether content trust is enabled or not. If it is enabled, user cann't pull unsigned images from this project. ,
	Prevent_vulnerable_images_from_running          bool   // Whether prevent the vulnerable images from running. ,
	Prevent_vulnerable_images_from_running_severity string // If the vulnerability is high than severity defined here, the images cann't be pulled. ,
	Automatically_scan_images_on_push               bool   // Whether scan images automatically when pushing.
}

var Projects []Project

// repo
type Repository struct {
	Id            int     // The ID of repository. ,
	Name          string  // The name of repository. ,
	Project_id    int     // The project ID of repository. ,
	Description   string  // The description of repository. ,
	Pull_count    int     // The pull count of repository. ,
	Star_count    int     // The star count of repository. ,
	Tags_count    int     // The tags count of repository. ,
	Labels        []Label // The label list. ,
	Creation_time string  // The creation time of repository. ,
	Update_time   string  // The update time of repository.
}
type Label struct {
	Id            int    // The ID of label. ,
	Name          string // The name of label. ,
	Description   string //The description of label. ,
	Color         string // The color of label. ,
	Scope         int    // The scope of label, g for global labels and p for project labels. ,
	Project_id    int    // The project ID if the label is a project label. ,
	Creation_time string // The creation time of label. ,
	Update_time   string // The update time of label.
}

var Repositorys []Repository

// tag
type DetailedTag struct {
	Digest         string         // The digest of the tag. ,
	Name           string         // The name of the tag. ,
	Size           int            // The size of the image. ,
	Architecture   string         // The architecture of the image. ,
	Os             string         // The os of the image. ,
	Docker_version string         // The version of docker which builds the image. ,
	Author         string         // The author of the image. ,
	Created        string         // The build time of the image. ,
	Signature      interface{}    // The signature of image, defined by RepoSignature. If it is null, the image is unsigned. ,
	Scan_overview  Inline_model_0 // The overview of the scan result. This is an optional property. ,
	Labels         []Label        // The label list.
}

type Inline_model_0 struct {
	Digest      string // The digest of the image. ,
	Scan_status string // The status of the scan job, it can be "pendnig", "running", "finished", "error". ,
	Job_id      int    // The ID of the job on jobservice to scan the image. ,
	Severity    int
	Details_key string         // The top layer name of this image in Clair, this is for calling Clair API to get the vulnerability list of this image. ,
	Components  Inline_Model_1 // The components overview of the image.
}

type Inline_Model_1 struct {
	Total   int                      // Total number of the components in this image. ,
	Summary []ComponentOverviewEntry // List of number of components of different severities.
}

type ComponentOverviewEntry struct {
	Severity int
	Count    int
}

var DetailedTags []DetailedTag
