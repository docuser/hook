## extremely naive Docker Hub webhook tester

run it using `make`, then add `http://<hostip>/whatever` to a webhook

I get the following output:

(ie, the documentation for the webhook chain is incomplete)

```
docker run --rm -it -p 80:80 hook
+ exec app
BODY: {"push_data": {"pushed_at": 1417501124, "images": ["imagehash1", "imagehash2", "imagehash3"], "pusher": "svendowideit"}, "callback_url": "https://registry.hub.docker.com/u/svendowideit/testhook/hook/21402b11bee3ecbdif11i4fifg0242eg1100a1/", "repository": {"status": "Active", "description": "", "is_trusted": true, "full_description": null, "repo_url": "https://registry.hub.docker.com/u/svendowideit/testhook/", "owner": "svendowideit", "is_official": false, "is_private": true, "name": "testhook", "namespace": "svendowideit", "star_count": 0, "comment_count": 0, "date_created": 1417494799, "dockerfile": "#\n# BUILD\t\tdocker build -t svendowideit/apt-cacher .\n# RUN\t\tdocker run -d -p 3142:3142 -name apt-cacher-run apt-cacher\n#\n# and then you can run containers with:\n# \t\tdocker run -t -i -rm -e http_proxy http://192.168.1.2:3142/ debian bash\n#\nFROM\t\tubuntu\nMAINTAINER\tSvenDowideit@home.org.au\n\n\nVOLUME\t\t[\"/var/cache/apt-cacher-ng\"]\nRUN\t\tapt-get update ; apt-get install -yq apt-cacher-ng\n\nEXPOSE \t\t3142\nCMD\t\tchmod 777 /var/cache/apt-cacher-ng ; /etc/init.d/apt-cacher-ng start ; tail -f /var/log/apt-cacher-ng/*\n", "repo_name": "svendowideit/testhook"}}
-- push_data: images is an array:
0 imagehash1
1 imagehash2
2 imagehash3
pusher is string svendowideit
pushed_at is float64 1.417501124e+09

callback_url is string https://registry.hub.docker.com/u/svendowideit/testhook/hook/21402b11bee3ecbdif11i4fifg0242eg1100a1/
-- repository: date_created is float64 1.417494799e+09
full_description is nil
star_count is float64 0
namespace is string svendowideit
comment_count is float64 0
description is string 
is_private is bool true
name is string testhook
is_trusted is bool true
repo_url is string https://registry.hub.docker.com/u/svendowideit/testhook/
is_official is bool false
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

repo_name is string svendowideit/testhook
status is string Active
owner is string svendowideit

callbackURL:  https://registry.hub.docker.com/u/svendowideit/testhook/hook/21402b11bee3ecbdif11i4fifg0242eg1100a1/
SUCCESS callback:  &{400 BAD REQUEST 400 HTTP/1.1 1 1 map[Server:[nginx] Date:[Tue, 02 Dec 2014 06:18:44 GMT] Content-Type:[application/json] Vary:[Cookie] X-Frame-Options:[SAMEORIGIN] Strict-Transport-Security:[max-age=31536000]] 0xc208260dc0 -1 [chunked] true map[] 0xc208026340 0xc20821a8a0}
```
