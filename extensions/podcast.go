package ext

// PodcastFeedExtension represents the podcast namespace extension for feeds
type PodcastFeedExtension struct {
	Locked           *PodcastLocked          `json:"locked,omitempty"`
	Funding          []*PodcastFunding       `json:"funding,omitempty"`
	GUID             string                  `json:"guid,omitempty"`
	License          *PodcastLicense         `json:"license,omitempty"`
	Location         *PodcastLocation        `json:"location,omitempty"`
	Medium           string                  `json:"medium,omitempty"`
	Images           *PodcastImages          `json:"images,omitempty"`
	Value            *PodcastValue           `json:"value,omitempty"`
	Block            []*PodcastBlock         `json:"block,omitempty"`
	Txt              []*PodcastTxt           `json:"txt,omitempty"`
	PodRoll          []*PodcastPodRoll       `json:"podRoll,omitempty"`
	UpdateFrequency  *PodcastUpdateFrequency `json:"updateFrequency,omitempty"`
	Podping          []*PodcastPodping       `json:"podping,omitempty"`
	Chat             []*PodcastChat          `json:"chat,omitempty"`
	Publisher        []*PodcastPublisher     `json:"publisher,omitempty"`
	Trailer          []*PodcastTrailer       `json:"trailer,omitempty"`
	LiveItems        []*PodcastLiveItem      `json:"liveItems,omitempty"`
}

// PodcastItemExtension represents the podcast namespace extension for items
type PodcastItemExtension struct {
	Transcript       []*PodcastTranscript      `json:"transcript,omitempty"`
	Chapters         []*PodcastChapters        `json:"chapters,omitempty"`
	Soundbite        []*PodcastSoundbite       `json:"soundbite,omitempty"`
	Person           []*PodcastPerson          `json:"person,omitempty"`
	Location         *PodcastLocation          `json:"location,omitempty"`
	Season           *PodcastSeason            `json:"season,omitempty"`
	Episode          string                    `json:"episode,omitempty"`
	AlternateEnclosure []*PodcastAlternateEnclosure `json:"alternateEnclosure,omitempty"`
	Value            *PodcastValue             `json:"value,omitempty"`
	Images           *PodcastImages            `json:"images,omitempty"`
	SocialInteract   []*PodcastSocialInteract  `json:"socialInteract,omitempty"`
	ContentLink      []*PodcastContentLink     `json:"contentLink,omitempty"`
	Txt              []*PodcastTxt             `json:"txt,omitempty"`
	ValueTimeSplit   []*PodcastValueTimeSplit  `json:"valueTimeSplit,omitempty"`
}

// PodcastLocked represents the podcast:locked tag
type PodcastLocked struct {
	Owner string `json:"owner,omitempty"`
	Value string `json:"value,omitempty"`
}

// PodcastTranscript represents the podcast:transcript tag
type PodcastTranscript struct {
	URL    string `json:"url,omitempty"`
	Type   string `json:"type,omitempty"`
	Language string `json:"language,omitempty"`
	Rel    string `json:"rel,omitempty"`
}

// PodcastFunding represents the podcast:funding tag
type PodcastFunding struct {
	URL   string `json:"url,omitempty"`
	Value string `json:"value,omitempty"`
}

// PodcastChapters represents the podcast:chapters tag
type PodcastChapters struct {
	URL  string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
}

// PodcastSoundbite represents the podcast:soundbite tag
type PodcastSoundbite struct {
	StartTime string `json:"startTime,omitempty"`
	Duration  string `json:"duration,omitempty"`
	Value     string `json:"value,omitempty"`
}

// PodcastPerson represents the podcast:person tag
type PodcastPerson struct {
	Name  string `json:"name,omitempty"`
	Role  string `json:"role,omitempty"`
	Group string `json:"group,omitempty"`
	Img   string `json:"img,omitempty"`
	Href  string `json:"href,omitempty"`
}

// PodcastLocation represents the podcast:location tag
type PodcastLocation struct {
	Geo string `json:"geo,omitempty"`
	OSM string `json:"osm,omitempty"`
	Value string `json:"value,omitempty"`
}

// PodcastSeason represents the podcast:season tag
type PodcastSeason struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// PodcastLicense represents the podcast:license tag
type PodcastLicense struct {
	URL   string `json:"url,omitempty"`
	Value string `json:"value,omitempty"`
}

// PodcastAlternateEnclosure represents the podcast:alternateEnclosure tag
type PodcastAlternateEnclosure struct {
	Type     string                 `json:"type,omitempty"`
	Length   string                 `json:"length,omitempty"`
	Bitrate  string                 `json:"bitrate,omitempty"`
	Height   string                 `json:"height,omitempty"`
	Default  string                 `json:"default,omitempty"`
	Title    string                 `json:"title,omitempty"`
	Sources  []*PodcastSource       `json:"sources,omitempty"`
	Integrity []*PodcastIntegrity   `json:"integrity,omitempty"`
}

// PodcastSource represents the podcast:source tag
type PodcastSource struct {
	URI string `json:"uri,omitempty"`
}

// PodcastIntegrity represents the podcast:integrity tag
type PodcastIntegrity struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// PodcastValue represents the podcast:value tag
type PodcastValue struct {
	Type      string                     `json:"type,omitempty"`
	Method    string                     `json:"method,omitempty"`
	Suggested string                     `json:"suggested,omitempty"`
	Recipients []*PodcastValueRecipient  `json:"recipients,omitempty"`
}

// PodcastValueRecipient represents the podcast:valueRecipient tag
type PodcastValueRecipient struct {
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Address  string `json:"address,omitempty"`
	Split    string `json:"split,omitempty"`
	Fee      string `json:"fee,omitempty"`
	Customkey string `json:"customkey,omitempty"`
	Customvalue string `json:"customvalue,omitempty"`
}

// PodcastImages represents the podcast:images tag
type PodcastImages struct {
	Srcset string `json:"srcset,omitempty"`
}

// PodcastLiveItem represents the podcast:liveItem tag
type PodcastLiveItem struct {
	Status  string `json:"status,omitempty"`
	Start   string `json:"start,omitempty"`
	End     string `json:"end,omitempty"`
	Content string `json:"content,omitempty"`
}

// PodcastContentLink represents the podcast:contentLink tag
type PodcastContentLink struct {
	Href  string `json:"href,omitempty"`
	Value string `json:"value,omitempty"`
}

// PodcastSocialInteract represents the podcast:socialInteract tag
type PodcastSocialInteract struct {
	URI        string `json:"uri,omitempty"`
	Protocol   string `json:"protocol,omitempty"`
	AccountID  string `json:"accountId,omitempty"`
	AccountURL string `json:"accountUrl,omitempty"`
	Priority   string `json:"priority,omitempty"`
}

// PodcastBlock represents the podcast:block tag
type PodcastBlock struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

// PodcastTxt represents the podcast:txt tag
type PodcastTxt struct {
	Purpose string `json:"purpose,omitempty"`
	Value   string `json:"value,omitempty"`
}

// PodcastPodRoll represents the podcast:podroll tag
type PodcastPodRoll struct {
	GUID  string `json:"guid,omitempty"`
	Value string `json:"value,omitempty"`
}

// PodcastUpdateFrequency represents the podcast:updateFrequency tag
type PodcastUpdateFrequency struct {
	Complete  string `json:"complete,omitempty"`
	Duration  string `json:"duration,omitempty"`
	Value     string `json:"value,omitempty"`
}

// PodcastPodping represents the podcast:podping tag
type PodcastPodping struct {
	Medium  string `json:"medium,omitempty"`
	Service string `json:"service,omitempty"`
}

// PodcastValueTimeSplit represents the podcast:valueTimeSplit tag
type PodcastValueTimeSplit struct {
	StartTime   string `json:"startTime,omitempty"`
	Duration    string `json:"duration,omitempty"`
	RemotePercentage string `json:"remotePercentage,omitempty"`
	RemoteStartTime string `json:"remoteStartTime,omitempty"`
	RemoteDuration string `json:"remoteDuration,omitempty"`
	RemoteItem string `json:"remoteItem,omitempty"`
}

// PodcastTrailer represents the podcast:trailer tag
type PodcastTrailer struct {
	URL     string `json:"url,omitempty"`
	PubDate string `json:"pubDate,omitempty"`
	Length  string `json:"length,omitempty"`
	Type    string `json:"type,omitempty"`
	Season  string `json:"season,omitempty"`
	Value   string `json:"value,omitempty"`
}

// PodcastChat represents the podcast:chat tag
type PodcastChat struct {
	Server   string `json:"server,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	URI      string `json:"uri,omitempty"`
	Space    string `json:"space,omitempty"`
	Platform string `json:"platform,omitempty"`
	Value    string `json:"value,omitempty"`
}

// PodcastPublisher represents the podcast:publisher tag
type PodcastPublisher struct {
	Name     string `json:"name,omitempty"`
	URL      string `json:"url,omitempty"`
	Img      string `json:"img,omitempty"`
	Value    string `json:"value,omitempty"`
}

// NewPodcastFeedExtension creates a new PodcastFeedExtension from a generic extension map
func NewPodcastFeedExtension(extensions map[string][]Extension) *PodcastFeedExtension {
	podcastExt := &PodcastFeedExtension{}
	
	// Parse locked
	if locked, ok := extensions["locked"]; ok && len(locked) > 0 {
		podcastExt.Locked = &PodcastLocked{
			Value: locked[0].Value,
		}
		if owner, ok := locked[0].Attrs["owner"]; ok {
			podcastExt.Locked.Owner = owner
		}
	}
	
	// Parse funding
	if funding, ok := extensions["funding"]; ok {
		podcastExt.Funding = []*PodcastFunding{}
		for _, f := range funding {
			fund := &PodcastFunding{
				Value: f.Value,
			}
			if url, ok := f.Attrs["url"]; ok {
				fund.URL = url
			}
			podcastExt.Funding = append(podcastExt.Funding, fund)
		}
	}
	
	// Parse guid
	if guid, ok := extensions["guid"]; ok && len(guid) > 0 {
		podcastExt.GUID = guid[0].Value
	}
	
	// Parse license
	if license, ok := extensions["license"]; ok && len(license) > 0 {
		podcastExt.License = &PodcastLicense{
			Value: license[0].Value,
		}
		if url, ok := license[0].Attrs["url"]; ok {
			podcastExt.License.URL = url
		}
	}
	
	// Parse location
	if location, ok := extensions["location"]; ok && len(location) > 0 {
		podcastExt.Location = &PodcastLocation{
			Value: location[0].Value,
		}
		if geo, ok := location[0].Attrs["geo"]; ok {
			podcastExt.Location.Geo = geo
		}
		if osm, ok := location[0].Attrs["osm"]; ok {
			podcastExt.Location.OSM = osm
		}
	}
	
	// Parse medium
	if medium, ok := extensions["medium"]; ok && len(medium) > 0 {
		podcastExt.Medium = medium[0].Value
	}
	
	// Parse block
	if block, ok := extensions["block"]; ok {
		podcastExt.Block = []*PodcastBlock{}
		for _, b := range block {
			blk := &PodcastBlock{
				Value: b.Value,
			}
			if id, ok := b.Attrs["id"]; ok {
				blk.ID = id
			}
			podcastExt.Block = append(podcastExt.Block, blk)
		}
	}
	
	// Parse images
	if images, ok := extensions["images"]; ok && len(images) > 0 {
		podcastExt.Images = &PodcastImages{}
		if srcset, ok := images[0].Attrs["srcset"]; ok {
			podcastExt.Images.Srcset = srcset
		}
	}
	
	// Parse value
	podcastExt.Value = parseValue(extensions)
	
	return podcastExt
}

// NewPodcastItemExtension creates a new PodcastItemExtension from a generic extension map
func NewPodcastItemExtension(extensions map[string][]Extension) *PodcastItemExtension {
	podcastExt := &PodcastItemExtension{}
	
	// Parse transcript
	if transcript, ok := extensions["transcript"]; ok {
		podcastExt.Transcript = []*PodcastTranscript{}
		for _, t := range transcript {
			trans := &PodcastTranscript{}
			if url, ok := t.Attrs["url"]; ok {
				trans.URL = url
			}
			if typ, ok := t.Attrs["type"]; ok {
				trans.Type = typ
			}
			if lang, ok := t.Attrs["language"]; ok {
				trans.Language = lang
			}
			if rel, ok := t.Attrs["rel"]; ok {
				trans.Rel = rel
			}
			podcastExt.Transcript = append(podcastExt.Transcript, trans)
		}
	}
	
	// Parse chapters
	if chapters, ok := extensions["chapters"]; ok {
		podcastExt.Chapters = []*PodcastChapters{}
		for _, c := range chapters {
			chap := &PodcastChapters{}
			if url, ok := c.Attrs["url"]; ok {
				chap.URL = url
			}
			if typ, ok := c.Attrs["type"]; ok {
				chap.Type = typ
			}
			podcastExt.Chapters = append(podcastExt.Chapters, chap)
		}
	}
	
	// Parse soundbite
	if soundbite, ok := extensions["soundbite"]; ok {
		podcastExt.Soundbite = []*PodcastSoundbite{}
		for _, s := range soundbite {
			sb := &PodcastSoundbite{
				Value: s.Value,
			}
			if startTime, ok := s.Attrs["startTime"]; ok {
				sb.StartTime = startTime
			}
			if duration, ok := s.Attrs["duration"]; ok {
				sb.Duration = duration
			}
			podcastExt.Soundbite = append(podcastExt.Soundbite, sb)
		}
	}
	
	// Parse season
	if season, ok := extensions["season"]; ok && len(season) > 0 {
		podcastExt.Season = &PodcastSeason{
			Value: season[0].Value,
		}
		if name, ok := season[0].Attrs["name"]; ok {
			podcastExt.Season.Name = name
		}
	}
	
	// Parse episode
	if episode, ok := extensions["episode"]; ok && len(episode) > 0 {
		podcastExt.Episode = episode[0].Value
	}
	
	// Parse person
	if person, ok := extensions["person"]; ok {
		podcastExt.Person = []*PodcastPerson{}
		for _, p := range person {
			pers := &PodcastPerson{
				Name: p.Value,
			}
			if role, ok := p.Attrs["role"]; ok {
				pers.Role = role
			}
			if group, ok := p.Attrs["group"]; ok {
				pers.Group = group
			}
			if img, ok := p.Attrs["img"]; ok {
				pers.Img = img
			}
			if href, ok := p.Attrs["href"]; ok {
				pers.Href = href
			}
			podcastExt.Person = append(podcastExt.Person, pers)
		}
	}
	
	// Parse location
	if location, ok := extensions["location"]; ok && len(location) > 0 {
		podcastExt.Location = &PodcastLocation{
			Value: location[0].Value,
		}
		if geo, ok := location[0].Attrs["geo"]; ok {
			podcastExt.Location.Geo = geo
		}
		if osm, ok := location[0].Attrs["osm"]; ok {
			podcastExt.Location.OSM = osm
		}
	}
	
	// Parse alternateEnclosure
	podcastExt.AlternateEnclosure = parseAlternateEnclosure(extensions)
	
	// Parse value
	podcastExt.Value = parseValue(extensions)
	
	// Parse images
	if images, ok := extensions["images"]; ok && len(images) > 0 {
		podcastExt.Images = &PodcastImages{}
		if srcset, ok := images[0].Attrs["srcset"]; ok {
			podcastExt.Images.Srcset = srcset
		}
	}
	
	// Parse socialInteract
	if socialInteract, ok := extensions["socialInteract"]; ok {
		podcastExt.SocialInteract = []*PodcastSocialInteract{}
		for _, s := range socialInteract {
			si := &PodcastSocialInteract{}
			if uri, ok := s.Attrs["uri"]; ok {
				si.URI = uri
			}
			if protocol, ok := s.Attrs["protocol"]; ok {
				si.Protocol = protocol
			}
			if accountID, ok := s.Attrs["accountId"]; ok {
				si.AccountID = accountID
			}
			if accountURL, ok := s.Attrs["accountUrl"]; ok {
				si.AccountURL = accountURL
			}
			if priority, ok := s.Attrs["priority"]; ok {
				si.Priority = priority
			}
			podcastExt.SocialInteract = append(podcastExt.SocialInteract, si)
		}
	}
	
	// Parse contentLink
	if contentLink, ok := extensions["contentLink"]; ok {
		podcastExt.ContentLink = []*PodcastContentLink{}
		for _, c := range contentLink {
			cl := &PodcastContentLink{
				Value: c.Value,
			}
			if href, ok := c.Attrs["href"]; ok {
				cl.Href = href
			}
			podcastExt.ContentLink = append(podcastExt.ContentLink, cl)
		}
	}
	
	return podcastExt
}

// Helper function to parse podcast:value
func parseValue(extensions map[string][]Extension) *PodcastValue {
	if value, ok := extensions["value"]; ok && len(value) > 0 {
		v := &PodcastValue{}
		
		if typ, ok := value[0].Attrs["type"]; ok {
			v.Type = typ
		}
		if method, ok := value[0].Attrs["method"]; ok {
			v.Method = method
		}
		if suggested, ok := value[0].Attrs["suggested"]; ok {
			v.Suggested = suggested
		}
		
		// Parse recipients
		if recipientEls, ok := value[0].Children["valueRecipient"]; ok {
			v.Recipients = []*PodcastValueRecipient{}
			
			for _, r := range recipientEls {
				recipient := &PodcastValueRecipient{}
				
				if name, ok := r.Attrs["name"]; ok {
					recipient.Name = name
				}
				if typ, ok := r.Attrs["type"]; ok {
					recipient.Type = typ
				}
				if address, ok := r.Attrs["address"]; ok {
					recipient.Address = address
				}
				if split, ok := r.Attrs["split"]; ok {
					recipient.Split = split
				}
				if fee, ok := r.Attrs["fee"]; ok {
					recipient.Fee = fee
				}
				if customkey, ok := r.Attrs["customkey"]; ok {
					recipient.Customkey = customkey
				}
				if customvalue, ok := r.Attrs["customvalue"]; ok {
					recipient.Customvalue = customvalue
				}
				
				v.Recipients = append(v.Recipients, recipient)
			}
		}
		
		return v
	}
	
	return nil
}

// Helper function to parse podcast:alternateEnclosure
func parseAlternateEnclosure(extensions map[string][]Extension) []*PodcastAlternateEnclosure {
	if altEnc, ok := extensions["alternateEnclosure"]; ok {
		result := []*PodcastAlternateEnclosure{}
		
		for _, a := range altEnc {
			ae := &PodcastAlternateEnclosure{}
			
			if typ, ok := a.Attrs["type"]; ok {
				ae.Type = typ
			}
			if length, ok := a.Attrs["length"]; ok {
				ae.Length = length
			}
			if bitrate, ok := a.Attrs["bitrate"]; ok {
				ae.Bitrate = bitrate
			}
			if height, ok := a.Attrs["height"]; ok {
				ae.Height = height
			}
			if def, ok := a.Attrs["default"]; ok {
				ae.Default = def
			}
			if title, ok := a.Attrs["title"]; ok {
				ae.Title = title
			}
			
			// Parse sources
			if sourceEls, ok := a.Children["source"]; ok {
				ae.Sources = []*PodcastSource{}
				
				for _, s := range sourceEls {
					source := &PodcastSource{}
					
					if uri, ok := s.Attrs["uri"]; ok {
						source.URI = uri
					}
					
					ae.Sources = append(ae.Sources, source)
				}
			}
			
			// Parse integrity
			if integrityEls, ok := a.Children["integrity"]; ok {
				ae.Integrity = []*PodcastIntegrity{}
				
				for _, i := range integrityEls {
					integrity := &PodcastIntegrity{
						Value: i.Value,
					}
					
					if typ, ok := i.Attrs["type"]; ok {
						integrity.Type = typ
					}
					
					ae.Integrity = append(ae.Integrity, integrity)
				}
			}
			
			result = append(result, ae)
		}
		
		return result
	}
	
	return nil
}