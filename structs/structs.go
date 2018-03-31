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

type Cate struct {
	ID   int
	Name string
}

type C struct {
	Name string`json:"name"`
}

type Category struct {
	ID          int           `json:"id"`
	Count       int           `json:"count"`
	Description string        `json:"description"`
	Link        string        `json:"link"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Taxonomy    string        `json:"taxonomy"`
	Parent      int           `json:"parent"`
	Meta        []interface{} `json:"meta"`
	Links       struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		About []struct {
			Href string `json:"href"`
		} `json:"about"`
		WpPostType []struct {
			Href string `json:"href"`
		} `json:"wp:post_type"`
		Curies []struct {
			Name      string `json:"name"`
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
	} `json:"_links"`
}

type CateDb struct {
	Id          int    `gorm:"column:term_id"`
	Name     string    `gorm:"column:name"`
	Slug     string `gorm:"column:slug"`
	TermGroup     int `gorm:"column:term_group"`
}

type CateDbTaxonomy struct {
	TermTaxonomyId          int    `gorm:"column:term_taxonomy_id"`
	TermId     int    `gorm:"column:term_id"`
	Taxonomy     string `gorm:"column:taxonomy"`
	Description     string `gorm:"column:description"`
	Parent     int `gorm:"column:parent"`
	Count     int `gorm:"column:count"`
}

func (CateDb) TableName() string {
	return "wp_terms"
}

func (CateDbTaxonomy) TableName() string {
	return "wp_term_taxonomy"
}