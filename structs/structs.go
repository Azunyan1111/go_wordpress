package structs

type Post struct{
	Title string 		`json:"title,omitempty,omitempty"`
	Content string		`json:"content,omitempty"`
	Excerpt string    	`json:"excerpt,omitempty"`
	DataGmt string		`json:"date_gmt,omitempty"`
	Categories []int	`json:"categories,omitempty"`
	Tags []int 		`json:"tags,omitempty"`
	Status  string		`json:"status,omitempty"`
}