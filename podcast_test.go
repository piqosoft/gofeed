package gofeed_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/piqosoft/gofeed"
	"github.com/stretchr/testify/assert"
)

func TestParseWithPodcastNamespace(t *testing.T) {
	// Open the example podcast namespace XML file
	f, err := os.Open(filepath.Join("testdata", "podcast", "example.xml"))
	assert.NoError(t, err)
	defer f.Close()

	// Parse the feed
	fp := gofeed.NewParser()
	feed, err := fp.Parse(f)
	assert.NoError(t, err)

	// Verify feed-level podcast namespace elements
	assert.NotNil(t, feed.PodcastExt, "PodcastExt should not be nil")
	
	// Test podcast:locked
	assert.NotNil(t, feed.PodcastExt.Locked, "Locked should not be nil")
	assert.Equal(t, "yes", feed.PodcastExt.Locked.Value)
	assert.Equal(t, "podcastowner@example.com", feed.PodcastExt.Locked.Owner)
	
	// Test podcast:guid
	assert.Equal(t, "y0ur-gu1d-g035-h3r3", feed.PodcastExt.GUID)
	
	// Test podcast:license
	assert.NotNil(t, feed.PodcastExt.License, "License should not be nil")
	assert.Equal(t, "my-podcast-license-v1", feed.PodcastExt.License.Value)
	assert.Equal(t, "https://example.org/mypodcastlicense/full.pdf", feed.PodcastExt.License.URL)
	
	// Test podcast:location
	assert.NotNil(t, feed.PodcastExt.Location, "Location should not be nil")
	assert.Equal(t, "Austin, TX", feed.PodcastExt.Location.Value)
	assert.Equal(t, "geo:30.2672,97.7431", feed.PodcastExt.Location.Geo)
	assert.Equal(t, "R113314", feed.PodcastExt.Location.OSM)
	
	// Test podcast:medium
	assert.Equal(t, "podcast", feed.PodcastExt.Medium)
	
	// Test podcast:block
	assert.NotNil(t, feed.PodcastExt.Block, "Block should not be nil")
	assert.GreaterOrEqual(t, len(feed.PodcastExt.Block), 3, "Block should have at least 3 entries")
	
	// Test podcast:funding
	assert.NotNil(t, feed.PodcastExt.Funding, "Funding should not be nil")
	assert.GreaterOrEqual(t, len(feed.PodcastExt.Funding), 1, "Funding should have at least 1 entry")
	assert.Equal(t, "Support the show!", feed.PodcastExt.Funding[0].Value)
	assert.Equal(t, "https://example.com/donate", feed.PodcastExt.Funding[0].URL)
	
	// Test podcast:value
	assert.NotNil(t, feed.PodcastExt.Value, "Value should not be nil")
	assert.Equal(t, "lightning", feed.PodcastExt.Value.Type)
	assert.Equal(t, "keysend", feed.PodcastExt.Value.Method)
	assert.Equal(t, "0.00000005000", feed.PodcastExt.Value.Suggested)
	
	// Test value recipients
	assert.NotNil(t, feed.PodcastExt.Value.Recipients, "Value Recipients should not be nil")
	assert.Equal(t, 2, len(feed.PodcastExt.Value.Recipients), "Value should have 2 recipients")
	assert.Equal(t, "podcaster", feed.PodcastExt.Value.Recipients[0].Name)
	assert.Equal(t, "node", feed.PodcastExt.Value.Recipients[0].Type)
	assert.Equal(t, "036557ea56b3b86f08be31bcd2557cae8021b0e3a9413f0c0e52625c6696972e57", feed.PodcastExt.Value.Recipients[0].Address)
	assert.Equal(t, "99", feed.PodcastExt.Value.Recipients[0].Split)
	
	// Check raw extension data for trailer and liveItem
	if ext, ok := feed.Extensions["podcast"]; ok {
		// Verify trailer raw extension
		if trailers, ok := ext["trailer"]; ok {
			assert.GreaterOrEqual(t, len(trailers), 1, "Raw trailer extensions should have at least 1 entry")
			trailer := trailers[0]
			assert.Equal(t, "Coming April 1st, 2021", trailer.Value)
			assert.Equal(t, "https://example.org/trailers/teaser", trailer.Attrs["url"])
			assert.Equal(t, "12345678", trailer.Attrs["length"])
			assert.Equal(t, "audio/mp3", trailer.Attrs["type"])
			assert.Equal(t, "Thu, 01 Apr 2021 08:00:00 EST", trailer.Attrs["pubdate"])
		} else {
			t.Error("No trailer entries found in raw extensions")
		}
		
		// Verify liveItem raw extension
		if liveItems, ok := ext["liveItem"]; ok {
			assert.GreaterOrEqual(t, len(liveItems), 1, "Raw liveItem extensions should have at least 1 entry")
			liveItem := liveItems[0]
			assert.Equal(t, "live", liveItem.Attrs["status"])
			assert.Equal(t, "2021-09-26T07:30:00.000-0600", liveItem.Attrs["start"])
			assert.Equal(t, "2021-09-26T09:30:00.000-0600", liveItem.Attrs["end"])
			
			// Check liveItem children
			assert.Contains(t, liveItem.Children, "title")
			assert.Contains(t, liveItem.Children, "description")
			assert.Contains(t, liveItem.Children, "link")
			assert.Contains(t, liveItem.Children, "guid")
			assert.Contains(t, liveItem.Children, "alternateEnclosure")
		} else {
			t.Error("No liveItem entries found in raw extensions")
		}
	} else {
		t.Error("No podcast extensions found")
	}
	
	// Test item-level podcast elements in first episode
	assert.GreaterOrEqual(t, len(feed.Items), 3, "Feed should have at least 3 items")
	
	item := feed.Items[0] // First item
	
	// Verify item has podcast extension
	assert.NotNil(t, item.PodcastExt, "Item PodcastExt should not be nil")
	
	// Test podcast:season
	assert.NotNil(t, item.PodcastExt.Season, "Season should not be nil")
	assert.Equal(t, "1", item.PodcastExt.Season.Value)
	assert.Equal(t, "Podcasting 2.0", item.PodcastExt.Season.Name)
	
	// Test podcast:episode
	assert.Equal(t, "3", item.PodcastExt.Episode)
	
	// Test podcast:chapters
	assert.NotNil(t, item.PodcastExt.Chapters, "Chapters should not be nil")
	assert.GreaterOrEqual(t, len(item.PodcastExt.Chapters), 1, "Chapters should have at least 1 entry")
	assert.Equal(t, "https://example.com/ep3_chapters.json", item.PodcastExt.Chapters[0].URL)
	assert.Equal(t, "application/json", item.PodcastExt.Chapters[0].Type)
	
	// Test podcast:soundbite
	assert.NotNil(t, item.PodcastExt.Soundbite, "Soundbite should not be nil")
	assert.GreaterOrEqual(t, len(item.PodcastExt.Soundbite), 1, "Soundbite should have at least 1 entry")
	assert.Equal(t, "33.833", item.PodcastExt.Soundbite[0].StartTime)
	assert.Equal(t, "60.0", item.PodcastExt.Soundbite[0].Duration)
	
	// Test podcast:transcript
	assert.NotNil(t, item.PodcastExt.Transcript, "Transcript should not be nil")
	assert.GreaterOrEqual(t, len(item.PodcastExt.Transcript), 1, "Transcript should have at least 1 entry")
	assert.Equal(t, "https://example.com/ep3/transcript.txt", item.PodcastExt.Transcript[0].URL)
	assert.Equal(t, "text/plain", item.PodcastExt.Transcript[0].Type)
	
	// Test podcast:person
	assert.NotNil(t, item.PodcastExt.Person, "Person should not be nil")
	assert.GreaterOrEqual(t, len(item.PodcastExt.Person), 3, "Person should have at least 3 entries")
	assert.Equal(t, "Adam Curry", item.PodcastExt.Person[0].Name)
	assert.Equal(t, "https://www.podchaser.com/creators/adam-curry-107ZzmWE5f", item.PodcastExt.Person[0].Href)
	assert.Equal(t, "https://example.com/images/adamcurry.jpg", item.PodcastExt.Person[0].Img)
	
	// Test second person with role
	assert.Equal(t, "Dave Jones", item.PodcastExt.Person[1].Name)
	assert.Equal(t, "guest", item.PodcastExt.Person[1].Role)
	
	// Test third person with group
	assert.Equal(t, "Becky Smith", item.PodcastExt.Person[2].Name)
	assert.Equal(t, "visuals", item.PodcastExt.Person[2].Group)
	assert.Equal(t, "cover art designer", item.PodcastExt.Person[2].Role)
	
	// Test podcast:alternateEnclosure
	assert.NotNil(t, item.PodcastExt.AlternateEnclosure, "AlternateEnclosure should not be nil")
	assert.GreaterOrEqual(t, len(item.PodcastExt.AlternateEnclosure), 5, "AlternateEnclosure should have at least 5 entries")
	
	// Check first alternateEnclosure
	assert.Equal(t, "audio/mpeg", item.PodcastExt.AlternateEnclosure[0].Type)
	assert.Equal(t, "43200000", item.PodcastExt.AlternateEnclosure[0].Length)
	assert.Equal(t, "128000", item.PodcastExt.AlternateEnclosure[0].Bitrate)
	assert.Equal(t, "true", item.PodcastExt.AlternateEnclosure[0].Default)
	assert.Equal(t, "Standard", item.PodcastExt.AlternateEnclosure[0].Title)
	
	// Check sources in first alternateEnclosure
	assert.NotNil(t, item.PodcastExt.AlternateEnclosure[0].Sources, "Sources should not be nil")
	assert.Equal(t, 2, len(item.PodcastExt.AlternateEnclosure[0].Sources), "Sources should have 2 entries")
	assert.Equal(t, "https://example.com/file-03.mp3", item.PodcastExt.AlternateEnclosure[0].Sources[0].URI)
	assert.Equal(t, "ipfs://someRandomMpegFile03", item.PodcastExt.AlternateEnclosure[0].Sources[1].URI)
	
	// Test video alternateEnclosure and verify it has some expected values
	videoEnc := item.PodcastExt.AlternateEnclosure[4]
	assert.Equal(t, "video/mp4", videoEnc.Type)
	assert.Equal(t, "Video version", videoEnc.Title)
	assert.Equal(t, "720", videoEnc.Height)
	
	// Verify integrity exists on video enclosure
	assert.NotNil(t, videoEnc.Integrity, "Integrity should not be nil")
	assert.GreaterOrEqual(t, len(videoEnc.Integrity), 1, "Integrity should have at least 1 entry")
	
	// Check one of the values
	if len(videoEnc.Integrity) > 0 {
		assert.Equal(t, "sri", videoEnc.Integrity[0].Type)
	}
	
	// Test podcast:value in item
	assert.NotNil(t, item.PodcastExt.Value, "Value should not be nil")
	assert.Equal(t, "lightning", item.PodcastExt.Value.Type)
	assert.Equal(t, "keysend", item.PodcastExt.Value.Method)
	
	// Check item value recipients
	assert.NotNil(t, item.PodcastExt.Value.Recipients, "Value Recipients should not be nil")
	assert.Equal(t, 3, len(item.PodcastExt.Value.Recipients), "Item value should have 3 recipients")
	
	// Test podcast:socialInteract
	assert.NotNil(t, item.PodcastExt.SocialInteract, "SocialInteract should not be nil")
	assert.GreaterOrEqual(t, len(item.PodcastExt.SocialInteract), 2, "SocialInteract should have at least 2 entries")
	assert.Equal(t, "https://podcastindex.social/web/@dave/108013847520053258", item.PodcastExt.SocialInteract[0].URI)
	assert.Equal(t, "activitypub", item.PodcastExt.SocialInteract[0].Protocol)
	assert.Equal(t, "@dave", item.PodcastExt.SocialInteract[0].AccountID)
	assert.Equal(t, "1", item.PodcastExt.SocialInteract[0].Priority)
}