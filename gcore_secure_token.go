package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

func gethash(client_id string, video_id string, secret string, expires int64) string {
	hash_body := fmt.Sprintf("%s_%s_%s_%d_", client_id, video_id, secret, expires)   // set of unique parameters of video

	md5sum := md5.Sum([]byte(hash_body))                   // get MD5 hash from parameters of video
	hash_md5 := base64.StdEncoding.EncodeToString(md5sum[:])

	hash_md5 = strings.Replace(hash_md5, "+", "-", -1)     // preparation for use in URL
	hash_md5 = strings.Replace(hash_md5, "/", "_", -1)
	hash_md5 = strings.Replace(hash_md5, "=", "", -1)
	return hash_md5
}

func main() {
	client_id := "2675"        // enter your account ID here
	secret := ""               // enter your secret key from CDN-resource here

	//VOD
	video_slug := "3dk4NsRt6vWsffEr"                      // enter your video's slug here
	expires := time.Now().UTC().Unix() + 24*60*60         // 24 hours, unixtime in seconds

	token := gethash(client_id, video_slug, secret, expires)
	fmt.Printf("https://demo-protected.gvideo.io/videos/%s_%s/%s/%d/master.m3u8 \n", client_id, video_slug, token, expires)

	//LIVE
	stream_id := "201693"                                 // enter your live stream id here
	expires = time.Now().UTC().Unix() + 24*60*60          // 24 hours, unixtime in seconds

	token = gethash(client_id, stream_id, secret, expires)
	fmt.Printf("https://demo-protected.gvideo.io/cmaf/%s_%s/%s/%d/master.m3u8 \n", client_id, stream_id, token, expires)
}
