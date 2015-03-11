package main

// auto generated from
// http://mervine.net/json2struct
// /highfive
type Results struct {
	IncompleteResults bool `json:"incomplete_results"`
	Items             []struct {
		GitURL     string `json:"git_url"`
		HtmlURL    string `json:"html_url"`
		Name       string `json:"name"`
		Path       string `json:"path"`
		Repository struct {
			ArchiveURL       string  `json:"archive_url"`
			AssigneesURL     string  `json:"assignees_url"`
			BlobsURL         string  `json:"blobs_url"`
			BranchesURL      string  `json:"branches_url"`
			CollaboratorsURL string  `json:"collaborators_url"`
			CommentsURL      string  `json:"comments_url"`
			CommitsURL       string  `json:"commits_url"`
			CompareURL       string  `json:"compare_url"`
			ContentsURL      string  `json:"contents_url"`
			ContributorsURL  string  `json:"contributors_url"`
			Description      string  `json:"description"`
			DownloadsURL     string  `json:"downloads_url"`
			EventsURL        string  `json:"events_url"`
			Fork             bool    `json:"fork"`
			ForksURL         string  `json:"forks_url"`
			FullName         string  `json:"full_name"`
			GitCommitsURL    string  `json:"git_commits_url"`
			GitRefsURL       string  `json:"git_refs_url"`
			GitTagsURL       string  `json:"git_tags_url"`
			HooksURL         string  `json:"hooks_url"`
			HtmlURL          string  `json:"html_url"`
			ID               float64 `json:"id"`
			IssueCommentURL  string  `json:"issue_comment_url"`
			IssueEventsURL   string  `json:"issue_events_url"`
			IssuesURL        string  `json:"issues_url"`
			KeysURL          string  `json:"keys_url"`
			LabelsURL        string  `json:"labels_url"`
			LanguagesURL     string  `json:"languages_url"`
			MergesURL        string  `json:"merges_url"`
			MilestonesURL    string  `json:"milestones_url"`
			Name             string  `json:"name"`
			NotificationsURL string  `json:"notifications_url"`
			Owner            struct {
				AvatarURL         string  `json:"avatar_url"`
				EventsURL         string  `json:"events_url"`
				FollowersURL      string  `json:"followers_url"`
				FollowingURL      string  `json:"following_url"`
				GistsURL          string  `json:"gists_url"`
				GravatarID        string  `json:"gravatar_id"`
				HtmlURL           string  `json:"html_url"`
				ID                float64 `json:"id"`
				Login             string  `json:"login"`
				OrganizationsURL  string  `json:"organizations_url"`
				ReceivedEventsURL string  `json:"received_events_url"`
				ReposURL          string  `json:"repos_url"`
				SiteAdmin         bool    `json:"site_admin"`
				StarredURL        string  `json:"starred_url"`
				SubscriptionsURL  string  `json:"subscriptions_url"`
				Type              string  `json:"type"`
				URL               string  `json:"url"`
			} `json:"owner"`
			Private         bool   `json:"private"`
			PullsURL        string `json:"pulls_url"`
			ReleasesURL     string `json:"releases_url"`
			StargazersURL   string `json:"stargazers_url"`
			StatusesURL     string `json:"statuses_url"`
			SubscribersURL  string `json:"subscribers_url"`
			SubscriptionURL string `json:"subscription_url"`
			TagsURL         string `json:"tags_url"`
			TeamsURL        string `json:"teams_url"`
			TreesURL        string `json:"trees_url"`
			URL             string `json:"url"`
		} `json:"repository"`
		Score       float64 `json:"score"`
		Sha         string  `json:"sha"`
		TextMatches []struct {
			Fragment string `json:"fragment"`
			Matches  []struct {
				Indices []float64 `json:"indices"`
				Text    string    `json:"text"`
			} `json:"matches"`
			ObjectType string `json:"object_type"`
			ObjectURL  string `json:"object_url"`
			Property   string `json:"property"`
		} `json:"text_matches"`
		URL string `json:"url"`
	} `json:"items"`
	TotalCount float64 `json:"total_count"`
}
