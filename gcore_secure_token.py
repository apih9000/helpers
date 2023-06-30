import base64
from hashlib import md5
from time import time

def gethash(client_id, video_id, secret, expires):
    hash_body = "%s_%s_%s_%s_" % (client_id, video_id, secret, expires)   # set of unique parameters of video
    hash_md5 =  base64.b64encode(                                         # get MD5 hash from parameters of video
            md5(hash_body.encode()).digest()                            
        ).decode().replace("+", "-").replace("/", "_").replace("=", "")   # preparation for use in URL 
    return hash_md5

client_id = "2675"       # enter your account ID here
secret = ""              # enter your secret key from CDN-resource here

#VOD
video_slug = "3dk4NsRt6vWsffEr"     # enter your video's slug here
expires = int(time()) + 24*60*60    # 24 hours, unixtime in seconds

token = gethash(client_id, video_slug, secret, expires)
print(f"https://demo-protected.gvideo.io/videos/{client_id}_{video_slug}/{token}/{expires}/master.m3u8")

#LIVE
stream_id = "201693"                      # enter your live stream id here
expires = int(time()) + 24*60*60    # 24 hours, unixtime in seconds

token = gethash(client_id, stream_id, secret, expires)
print(f"https://demo-protected.gvideo.io/cmaf/{client_id}_{stream_id}/{token}/{expires}/master.m3u8")
