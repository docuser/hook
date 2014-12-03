## extremely naive Docker Hub webhook tester

run it using `make`, then add `http://<hostip>/whatever` to a webhook

I get the following output:


```
docker run --rm -it -p 80:80 hook
+ exec app
{
  "callback_url": "https://registry.hub.docker.com/u/svendowideit/testhook/hook/2141b4h1e4gjfibec111i4dcff0242eg110014/",
  "push_data": {
    "images": [
      "imagehash1",
      "imagehash2",
      "imagehash3"
    ],
    "pushed_at": 1.417565871e+09,
    "pusher": "svendowideit"
  },
  "repository": {
    "comment_count": 0,
    "date_created": 1.417494799e+09,
    "description": "",
    "dockerfile": "#\n# BUILD\u0009\u0009docker build -t svendowideit/apt-cacher .\n# RUN\u0009\u0009docker run -d -p 3142:3142 -name apt-cacher-run apt-cacher\n#\n# and then you can run containers with:\n# \u0009\u0009docker run -t -i -rm -e http_proxy http://192.168.1.2:3142/ debian bash\n#\nFROM\u0009\u0009ubuntu\nMAINTAINER\u0009SvenDowideit@home.org.au\n\n\nVOLUME\u0009\u0009[\"/var/cache/apt-cacher-ng\"]\nRUN\u0009\u0009apt-get update ; apt-get install -yq apt-cacher-ng\n\nEXPOSE \u0009\u00093142\nCMD\u0009\u0009chmod 777 /var/cache/apt-cacher-ng ; /etc/init.d/apt-cacher-ng start ; tail -f /var/log/apt-cacher-ng/*\n",
    "full_description": null,
    "is_official": false,
    "is_private": true,
    "is_trusted": true,
    "name": "testhook",
    "namespace": "svendowideit",
    "owner": "svendowideit",
    "repo_name": "svendowideit/testhook",
    "repo_url": "https://registry.hub.docker.com/u/svendowideit/testhook/",
    "star_count": 0,
    "status": "Active"
  }
}callback to  https://registry.hub.docker.com/u/svendowideit/testhook/hook/2141b4h1e4gjfibec111i4dcff0242eg110014/
"OK"
```

once you know the callback_url, you can just as easily use curl :)

```
$ curl --data '{"state":"success"}' https://registry.hub.docker.com/u/svendowideit/testhook/hook/21402b11bee3ecb1i4fifg0242eg1100a1/
"OK"
```
