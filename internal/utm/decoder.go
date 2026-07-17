package utm

import "net/http"

type UtmData struct {
	Source   string
	Medium   string
	Campaign string
	Content  string
	Term     string
}

func DecodeUtm(r *http.Request) *UtmData {
	utmData := &UtmData{}
	query := r.URL.Query()

	for key, value := range query {
		switch key {
		case "utm_source":
			utmData.Source = value[0]
			break
		case "utm_medium":
			utmData.Medium = value[0]
			break
		case "utm_campaign":
			utmData.Campaign = value[0]
			break
		case "utm_Conetnt":
			utmData.Content = value[0]
			break
		case "utm_term":
			utmData.Term = value[0]
			break
		}
	}

	return utmData
}
