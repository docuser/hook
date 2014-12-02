## extremely naive Docker Hub webhook tester

run it using `make`, then add `http://<hostip>/whatever` to a webhook

I get the following output:

(ie, the documentation for the webhook chain is incomplete)

```
docker run --rm -it -p 80:80 hook
+ exec app
BODY: {"push_data": {"pushed_at": 1417504228, "images": ["imagehash1", "imagehash2", "imagehash3"], "pusher": "svendowideit"}, "callback_url": "https://registry.hub.docker.com/u/svendowideit/testhook/hook/2140334fb4jjaabdj211i4ccdj0242eg1100a4/", "repository": {"status": "Active", "description": "", "is_trusted": true, "full_description": null, "repo_url": "https://registry.hub.docker.com/u/svendowideit/testhook/", "owner": "svendowideit", "is_official": false, "is_private": true, "name": "testhook", "namespace": "svendowideit", "star_count": 0, "comment_count": 0, "date_created": 1417494799, "dockerfile": "#\n# BUILD\t\tdocker build -t svendowideit/apt-cacher .\n# RUN\t\tdocker run -d -p 3142:3142 -name apt-cacher-run apt-cacher\n#\n# and then you can run containers with:\n# \t\tdocker run -t -i -rm -e http_proxy http://192.168.1.2:3142/ debian bash\n#\nFROM\t\tubuntu\nMAINTAINER\tSvenDowideit@home.org.au\n\n\nVOLUME\t\t[\"/var/cache/apt-cacher-ng\"]\nRUN\t\tapt-get update ; apt-get install -yq apt-cacher-ng\n\nEXPOSE \t\t3142\nCMD\t\tchmod 777 /var/cache/apt-cacher-ng ; /etc/init.d/apt-cacher-ng start ; tail -f /var/log/apt-cacher-ng/*\n", "repo_name": "svendowideit/testhook"}}
-- push_data: images is an array:
0 imagehash1
1 imagehash2
2 imagehash3
pusher is string svendowideit
pushed_at is float64 1.417504228e+09

callback_url is string https://registry.hub.docker.com/u/svendowideit/testhook/hook/2140334fb4jjaabdj211i4ccdj0242eg1100a4/
-- repository: is_official is bool false
dockerfile is string #
# BUILD         docker build -t svendowideit/apt-cacher .
# RUN           docker run -d -p 3142:3142 -name apt-cacher-run apt-cacher
#
# and then you can run containers with:
#               docker run -t -i -rm -e http_proxy http://192.168.1.2:3142/ debian bash
#
FROM            ubuntu
MAINTAINER      SvenDowideit@home.org.au


VOLUME          ["/var/cache/apt-cacher-ng"]
RUN             apt-get update ; apt-get install -yq apt-cacher-ng

EXPOSE          3142
CMD             chmod 777 /var/cache/apt-cacher-ng ; /etc/init.d/apt-cacher-ng start ; tail -f /var/log/apt-cacher-ng/*

owner is string svendowideit
is_private is bool true
namespace is string svendowideit
description is string 
is_trusted is bool true
status is string Active
date_created is float64 1.417494799e+09
name is string testhook
star_count is float64 0
comment_count is float64 0
repo_name is string svendowideit/testhook
full_description is nil
repo_url is string https://registry.hub.docker.com/u/svendowideit/testhook/

callbackURL:  https://registry.hub.docker.com/u/svendowideit/testhook/hook/2140334fb4jjaabdj211i4ccdj0242eg1100a4/
SUCCESS callback:  &{400 BAD REQUEST 400 HTTP/1.1 1 1 map[X-Frame-Options:[SAMEORIGIN] Strict-Transport-Security:[max-age=31536000] Server:[nginx] Date:[Tue, 02 Dec 2014 07:10:28 GMT] Content-Type:[application/json] Vary:[Cookie]] 0xc20821c4c0 -1 [chunked] true map[] 0xc2080284e0 0xc208005920}
```
even sadder, using `curl` from some other computer succeeds. I wonder what my go code is doing wrong, and why there's no security.

```
$ curl --data '{"state":"success"}' https://registry.hub.docker.com/u/svendowideit/testhook/hook/21402b11bee3ecb1i4fifg0242eg1100a1/
"OK"
```
