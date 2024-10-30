package publishers

type PublisherResponse struct {
	PublishersID            string `json:"publishers_id"`
	Name                    string `json:"name"`
	Address                 string `json:"address"`
	Phone                   string `json:"phone"`
	Email                   string `json:"email"`
	Website                 string `json:"website"`
	FoundedYear             int    `json:"founded_year"`
	Country                 string `json:"country"`
	ContactPerson1          string `json:"contact_person_1"`
	ContactPerson2          string `json:"contact_person_2"`
	Fax                     string `json:"fax"`
	SocialMediaFBLinks      string `json:"social_fb_links"`
	SocialMediaTwitterLinks string `json:"social_twitter_links"`
	SocialMediaWebLinks     string `json:"social_web_links"`
	JoinDate                string `json:"join_date"`
	EntryTime               string `json:"entry_time"`
}
