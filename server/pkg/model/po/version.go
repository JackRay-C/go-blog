package po

type Version struct {
	ID              int64  `json:"id"`
	PostID          int64  `json:"post_id"`
	MarkdownContent string `json:"markdown_content"`
	HtmlContent     string `json:"html_content"`
	SubjectID       int64  `json:"subject_id"`
}
