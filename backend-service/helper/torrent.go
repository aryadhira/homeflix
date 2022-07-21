package helper

import (
	"homeflix/models"
	"net/url"
)

func EncodeUrl(urls string) string {
	query := urls
	return url.QueryEscape(query)
}

func GenerateEncodedTrackerUrl() string {
	urltrackers := []string{
		"udp://open.demonii.com:1337/announce",
		// "udp://tracker.openbittorrent.com:80",
		// "udp://tracker.coppersurfer.tk:6969",
		"udp://glotorrents.pw:6969/announce",
		"udp://tracker.opentrackr.org:1337/announce",
		"udp://torrent.gresille.org:80/announce",
		// "udp://p4p.arenabg.com:1337",
		// "udp://tracker.leechers-paradise.org:6969",
	}

	strreturn := ""
	lastdataindex := len(urltrackers) - 1
	for i, each := range urltrackers {
		encodedtracker := EncodeUrl(each)

		if i < lastdataindex {
			strreturn = strreturn + "&tr=" + encodedtracker + "&tr="
		} else {
			strreturn = strreturn + "&tr=" + encodedtracker
		}

	}
	// use tracker for webtorrent
	// strreturn = "&tr=udp%3A%2F%2Fexplodie.org%3A6969&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.empire-js.us%3A1337&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337&tr=wss%3A%2F%2Ftracker.btorrent.xyz&tr=wss%3A%2F%2Ftracker.fastcast.nz&tr=wss%3A%2F%2Ftracker.openwebtorrent.com"

	return strreturn
}

func CreateMagnetUrl(movie models.Movie, torrentidx int) string {
	titlelong := movie.TitleLong
	torrents := movie.Torrents
	torrentdata := torrents[torrentidx]
	hash := torrentdata.Hash
	quality := torrentdata.Quality

	urltoencode := titlelong + " [" + quality + "] " + "[YTS.MX]"
	encodedurl := EncodeUrl(urltoencode)
	trackerurl := GenerateEncodedTrackerUrl()
	magneturl := "magnet:?xt=urn:btih:" + hash + "&dn=" + encodedurl + trackerurl

	return magneturl
}

func GetTorrentByQuality(torrents []models.Torrent, _type string, _qlty string) int {
	idx := 0
	for i, each := range torrents {
		if each.Quality == _qlty && each.Type == _type {
			idx = i
		}
	}
	return idx
}

func GenerateMagnetUrl(title_long string, hash string, _qlty string) string {
	urltoencode := title_long + " [" + _qlty + "] " + "[YTS.MX]"
	encodedurl := EncodeUrl(urltoencode)
	trackerurl := GenerateEncodedTrackerUrl()
	magneturl := "magnet:?xt=urn:btih:" + hash + "&dn=" + encodedurl + trackerurl

	return magneturl
}
